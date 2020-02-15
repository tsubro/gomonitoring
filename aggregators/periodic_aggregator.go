package aggregators


type PeriodicAggregatorImpl struct {

	agg map[string]int
}

func (a *PeriodicAggregatorImpl) Aggregate() {

	a.agg = make(map[string]int)
	// logs := context.GetLogs()

	// for _, val := range logs {
	// 	a.agg[val] = a.agg[val]+1
	// }

}
