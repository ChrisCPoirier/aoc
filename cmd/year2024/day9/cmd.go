package day9

import (
	"aoc/cmd/common"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Short: "day9",
	Long:  `day9`,
	Use:   "day9",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	common.Run(parent, command, 1, part1)
	common.Run(parent, command, 2, part2)
}

type block struct {
	id   int
	size int
}

func toBlocks(s []string) []block {
	id := 0
	blocks := []block{}
	for i, size := range common.AsInts(s) {
		if i%2 == 0 {
			b := block{
				id:   id,
				size: size,
			}
			id++
			blocks = append(blocks, b)
			continue
		}
		b := block{
			id:   -1,
			size: size,
		}
		blocks = append(blocks, b)
	}
	return blocks
}

func part1(s []byte) int {
	in := strings.Split(string(s), ``)
	blocks := toBlocks(in)

	for {
		move := blocks[len(blocks)-1]
		blocks = blocks[0 : len(blocks)-2]

		for i := 0; i < len(blocks); i++ {
			next := blocks[i]
			if move.size == 0 {
				break
			}

			if next.id != -1 {
				continue
			}

			if next.size < move.size {
				next.id = move.id
				blocks[i] = next
				move.size -= next.size
				continue
			}

			next.id = move.id

			newFree := block{
				id:   -1,
				size: next.size - move.size,
			}

			next.size = move.size

			blocks[i] = next
			if newFree.size > 0 {
				blocks = slices.Insert(blocks, i+1, newFree)
			}
			move.size = 0
			break
		}

		if move.size > 0 {
			blocks = append(blocks, move)
			break
		}
	}

	return score(blocks)
}

func part2(s []byte) int {
	in := strings.Split(string(s), ``)

	blocks := toBlocks(in)

	curID := -1
	for {
		j := getNextBlockIndex(blocks, curID)
		curID = blocks[j].id
		if j == 0 {
			break
		}

		for i := 0; i < j; i++ {
			if blocks[i].id != -1 {
				continue
			}

			if blocks[i].size < blocks[j].size {
				continue
			}

			newFree := block{
				id:   -1,
				size: blocks[i].size - blocks[j].size,
			}

			blocks[i].id = blocks[j].id
			blocks[i].size = blocks[j].size

			blocks[j].id = -1

			if newFree.size > 0 {
				blocks = slices.Insert(blocks, i+1, newFree)
			}

			break
		}
	}

	return score(blocks)
}

func getNextBlockIndex(blocks []block, curID int) int {
	if curID == -1 {
		return len(blocks) - 1
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		block := blocks[i]

		if block.id == curID-1 {
			return i
		}

	}
	return 0
}

func score(blocks []block) int {
	score := 0
	pos := 0
	for _, block := range blocks {
		if block.id == -1 {
			pos += block.size

			continue
		}

		for range block.size {
			score += pos * block.id
			pos++
		}
	}
	return score
}
