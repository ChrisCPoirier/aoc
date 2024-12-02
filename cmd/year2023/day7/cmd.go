package day7

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day7",
	Short: "day7",
	Long:  `day7`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/year%s/%s/1.txt`, parent, command))

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b), false))
	logrus.Infof("score part2: %d", part1(string(b), true))
}

const (
	FIVE_OF_A_KIND  = 7
	FOUR_OF_A_KIND  = 6
	FULL_HOUSE      = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR        = 3
	ONE_PAIR        = 2
	HIGH_CARD       = 1
)

type hand struct {
	cards    string
	bid      int
	bestHand int
}

func GetBestHand(cards string) int {
	m := map[rune]int{}
	for _, r := range cards {
		m[r] += 1
	}

	if len(m) == 1 {
		return FIVE_OF_A_KIND
	}

	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return FOUR_OF_A_KIND
			}
		}
		return FULL_HOUSE
	}

	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	}

	if len(m) == 5 {
		return HIGH_CARD
	}

	return ONE_PAIR
}

var strength = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func part1(s string, jokersAreWild bool) int64 {
	var score int64 = 0

	if jokersAreWild {
		strength['J'] = -1
	} else {
		strength['J'] = 11
	}

	//32T3K 765
	hands := []hand{}
	for _, line := range strings.Split(s, "\n") {

		split := strings.Split(line, ` `)

		bid, err := strconv.Atoi(split[1])

		if err != nil {
			panic(err)
		}

		h := hand{cards: split[0], bid: bid}
		h.bestHand = GetBestHand(h.cards)

		if jokersAreWild && strings.Contains(h.cards, `J`) {
			h.bestHand = playWilds(h)
		}

		hands = append(hands, h)

	}

	fmt.Printf("%#v\n", hands)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].bestHand == hands[j].bestHand {
			for k := range hands[i].cards {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}

				return strength[hands[i].cards[k]] < strength[hands[j].cards[k]]

			}
		}
		return hands[i].bestHand < hands[j].bestHand
	})

	fmt.Printf("%#v\n", hands)

	for i, hand := range hands {
		score += int64((i + 1) * hand.bid)
	}

	return score
}

// T55J5, KTJJT, and QQQJA are now all four of a kind!
func playWilds(h hand) int {
	m := map[rune]int{}
	for _, r := range h.cards {
		m[r] += 1
	}

	if m['J'] >= 4 {
		return FIVE_OF_A_KIND
	}

	if m['J'] == 3 {
		if len(m) == 2 {
			return FIVE_OF_A_KIND
		}
		return FOUR_OF_A_KIND
	}

	if m['J'] == 2 {
		if h.bestHand == TWO_PAIR {
			return FOUR_OF_A_KIND
		}

		if h.bestHand == ONE_PAIR {
			return THREE_OF_A_KIND
		}

		if h.bestHand == FULL_HOUSE {
			return FIVE_OF_A_KIND
		}
	}

	if m['J'] == 1 {
		if h.bestHand == THREE_OF_A_KIND {
			return FOUR_OF_A_KIND
		}
		if h.bestHand == TWO_PAIR {
			return FULL_HOUSE
		}

		if h.bestHand == ONE_PAIR {
			return THREE_OF_A_KIND
		}

		if h.bestHand == FOUR_OF_A_KIND {
			return FIVE_OF_A_KIND
		}
	}

	return ONE_PAIR
}
