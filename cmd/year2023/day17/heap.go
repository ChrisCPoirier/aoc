package day17

type cityBlock struct {
	heatLoss        int
	row             int
	column          int
	rowDirection    int
	columnDirection int
	count           int
	index           int
}

type PriorityQueue []*cityBlock

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	cityBlock := x.(*cityBlock)
	cityBlock.index = n
	*pq = append(*pq, cityBlock)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	cityBlock := old[n-1]
	old[n-1] = nil
	cityBlock.index = -1
	*pq = old[0 : n-1]
	return cityBlock
}
