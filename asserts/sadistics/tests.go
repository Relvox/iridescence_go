package sadistics

func ChiSquared(bucketCounts map[int]int, totalItems int) float64 {
	numBuckets := len(bucketCounts)
	expectedCount := float64(totalItems) / float64(numBuckets)
	var chiSquaredSum float64

	for _, count := range bucketCounts {
		observedCount := float64(count)
		chiSquaredSum += (observedCount - expectedCount) * (observedCount - expectedCount) / expectedCount
	}

	return chiSquaredSum
}

