package structures

type FrequencyPair struct {
	index     int
	word      string
	frequency int32
}

type MyHeap []*FrequencyPair

func (pq MyHeap) Len() int           { return len(pq) }
func (pq MyHeap) Less(i, j int) bool { return pq[i].frequency > pq[j].frequency }
func (pq MyHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}
