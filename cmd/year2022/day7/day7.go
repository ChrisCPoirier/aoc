package day7

import (
	"os"
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
		execute()
	},
}

func execute() {
	b, err := os.ReadFile(`data/day7-1.txt`)

	if err != nil {
		logrus.Fatal(err)
	}

	root := buildTree(string(b))
	logrus.Infof("Day 7 part 1: %d", aoc(root))
	logrus.Infof("Day 7 part 2: %d", aoc2(root))
}

const updateSize = 30000000
const fileSystemSize = 70000000

func aoc2(root *Directory) int {
	neededSpace := updateSize - (fileSystemSize - root.Size())
	return smallestDirAtOrAboveSize(root, neededSpace).Size()
}

func smallestDirAtOrAboveSize(d *Directory, s int) *Directory {
	var smallest *Directory

	if d.Size() >= s {
		smallest = d
	}

	for _, dir := range d.Directories() {
		smallDir := smallestDirAtOrAboveSize(dir, s)
		if smallDir != nil && smallDir.Size() < smallest.Size() {
			smallest = smallDir
		}
	}
	return smallest
}

func buildTree(s string) *Directory {
	root := &Directory{
		name: `/`,
	}

	lines := strings.Split(s, "\n")

	current := root
	for _, line := range lines {
		switch line {
		case `$ ls`:
			continue
		case `$ cd /`:
			current = root
			continue
		case `$ cd ..`:
			current = current.Parent()
			continue
		}

		if strings.Contains(line, `$ cd `) {
			subDirs := current.Directories()
			var newDir *Directory
			dirName := strings.Replace(line, `$ cd `, ``, 1)

			for _, sub := range subDirs {
				if sub.Name() == dirName {
					newDir = sub
					break
				}
			}

			if newDir != nil {
				current = newDir
				continue
			}
			logrus.Fatalf(`you messed up no directory exists with name %s`, dirName)
		}
		if line[0] == '$' {
			logrus.Fatal(`you messed up. Found operation with no action taken`)
		}

		if current.children == nil {
			current.children = []Node{}
		}

		if line[0:3] == `dir` {
			name := line[4:]

			current.children = append(current.children,
				&Directory{
					parent: current,
					name:   name,
				},
			)
			continue
		}

		s := strings.Split(line, ` `)

		size, _ := strconv.Atoi(s[0])
		name := s[1]

		current.children = append(current.children,
			&File{
				parent: current,
				name:   name,
				size:   size,
			},
		)

	}
	return root
}

func aoc(root *Directory) int {

	total := 0

	for _, dir := range getDirectoriesUnderSize(root, 100000) {
		total += dir.Size()
	}
	return total
}

func getDirectoriesUnderSize(d *Directory, s int) []*Directory {
	u := []*Directory{}
	if d.Size() <= s {
		u = append(u, d)
	}

	dirs := d.Directories()

	for _, sub := range dirs {
		under := getDirectoriesUnderSize(sub, s)
		if len(under) > 0 {
			u = append(u, under...)
		}
	}

	return u
}

type File struct {
	parent *Directory
	name   string
	size   int
}

func (f *File) Name() string {
	return f.name
}
func (f *File) Size() int {
	return f.size
}

func (f *File) Parent() *Directory {
	return f.parent
}

type Directory struct {
	parent   *Directory
	name     string
	size     int
	children []Node
}

func (d *Directory) Parent() *Directory {
	return d.parent
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	if d.size == 0 {
		for _, child := range d.children {
			d.size += child.Size()
		}
	}
	return d.size
}

func (d *Directory) Directories() []*Directory {
	directories := []*Directory{}

	for _, child := range d.children {
		switch d := child.(type) {
		case *Directory:
			directories = append(directories, d)
		}
	}
	return directories
}

type Node interface {
	Parent() *Directory
	Name() string
	Size() int
}
