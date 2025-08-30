package structures

type FrequencyPair struct {
	Index     int
	Word      string
	Frequency int32
}

type MyHeap []*FrequencyPair

func (pq MyHeap) Len() int { return len(pq) }

func (pq MyHeap) Less(i, j int) bool {
	if pq[i].Frequency != pq[j].Frequency {
		return pq[i].Frequency > pq[j].Frequency
	} else {
		return pq[i].Word < pq[j].Word
	}
}

func (pq MyHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index, pq[j].Index = i, j
}

func (pq *MyHeap) Push(x any) {
	frequencyPair := x.(*FrequencyPair)
	frequencyPair.Index = len(*pq)
	*pq = append(*pq, frequencyPair)
}

func (pq *MyHeap) Pop() any {
	temp, n := *pq, len(*pq)
	lastItem := temp[n-1]
	temp[n-1] = nil
	*pq = temp[0 : n-1]
	return lastItem
}

func (pq MyHeap) Peek() any {
	if pq.Len() == 0 {
		return nil
	}
	return *pq[0]
}

func (pq MyHeap) IsEmpty() bool {
	return len(pq) == 0
}
