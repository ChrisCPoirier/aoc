package grid

var DIR_LEFT = []int{0, -1}
var DIR_RIGHT = []int{0, 1}
var DIR_UP = []int{-1, 0}
var DIR_DOWN = []int{1, 0}

var DIR_ALL = [][][]int{
	{DIR_UP},
	{DIR_UP, DIR_LEFT},
	{DIR_UP, DIR_RIGHT},
	{DIR_DOWN},
	{DIR_DOWN, DIR_LEFT},
	{DIR_DOWN, DIR_RIGHT},
	{DIR_LEFT},
	{DIR_RIGHT},
}

var DIR_X = [][][]int{
	{DIR_UP, DIR_LEFT},
	{DIR_UP, DIR_RIGHT},
	{DIR_DOWN, DIR_LEFT},
	{DIR_DOWN, DIR_RIGHT},
}

var DIR_CROSS = [][]int{
	DIR_UP,
	DIR_DOWN,
	DIR_LEFT,
	DIR_RIGHT,
}
