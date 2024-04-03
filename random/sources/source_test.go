package sources_test

import (
	"fmt"
	"log"
	"math"
	"testing"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/relvox/iridescence_go/random/sources"
)

type pRNG interface {
	Uint64() uint64
}

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

func getChiSquaredThreshold(degreesOfFreedom int, confidenceLevel float64) float64 {
	dist := distuv.ChiSquared{
		K: float64(degreesOfFreedom),
	}

	return dist.Quantile(1 - confidenceLevel)
}

func UniformDistribution_Test(pRng pRNG, numSamples int, numBuckets int) bool {
	bucketCounts := make(map[int]int)

	for i := 0; i < numSamples; i++ {
		bucket := int(pRng.Uint64() % uint64(numBuckets))
		bucketCounts[bucket]++
	}
	log.Println(bucketCounts)
	chiSquared := ChiSquared(bucketCounts, numSamples)
	threshold := getChiSquaredThreshold(numBuckets-1, 0.95) // Implement this function based on a chi-squared table
	fmt.Printf("   Chi-Squared: %f Threshold: %f\n", chiSquared, threshold)
	return chiSquared < threshold
}

// Test for Mean
func Mean_Test(pRng pRNG, numSamples int) bool {
	var sum uint64
	for i := 0; i < numSamples; i++ {
		sum += pRng.Uint64()
	}
	mean := float64(sum) / float64(numSamples)
	expectedMean := float64(math.MaxUint64) / 2 / 2
	stdDev := math.Sqrt(float64(math.MaxUint64) / 12)

	lowerBound := expectedMean - 2*stdDev
	upperBound := expectedMean + 2*stdDev
	fmt.Printf("   Mean: %e\n", mean)
	return mean > lowerBound && mean < upperBound
}

// Test for Period
func Period_Test(pRng pRNG, maxPeriod int) int {
	initial := pRng.Uint64()
	for i := 1; i <= maxPeriod; i++ {
		if pRng.Uint64() == initial {
			fmt.Printf("   Period: %d\n", i)
			return i
		}
	}
	fmt.Printf("   Period > %d\n", maxPeriod)
	return -1
}

// Test for Serial Correlation
func SerialCorrelation_Test(pRng pRNG, numSamples int) bool {
	var lastValue, currentValue uint64
	var sum, mean, meanSquare, correlation float64

	lastValue = pRng.Uint64()
	mean = float64(lastValue)
	meanSquare = float64(lastValue) * float64(lastValue)

	for i := 1; i < numSamples; i++ {
		currentValue = pRng.Uint64()
		sum += float64(currentValue) * float64(lastValue)
		mean += float64(currentValue)
		meanSquare += float64(currentValue) * float64(currentValue)
		lastValue = currentValue
	}

	mean /= float64(numSamples)
	meanSquare /= float64(numSamples)
	correlation = (sum/float64(numSamples) - mean*mean) / (meanSquare - mean*mean)
	const someAcceptableCorrelationRange = 0.05
	fmt.Printf("   Correlation: %f\n", correlation)
	return math.Abs(correlation) < someAcceptableCorrelationRange
}

const (
	N = 1_000_000
	K = 1_000
)

func Test_rand(t *testing.T) {
	rands := []pRNG{
		rand.New(rand.NewSource(12345)),
		sources.NewWELL512(12345),
	}

	for ti, tt := range rands {
		t.Run(fmt.Sprint(ti), func(t *testing.T) {

			if UniformDistribution_Test(tt, N, K) {
				fmt.Println("Passed Uniform Distribution Test")
			} else {
				fmt.Println("Failed Uniform Distribution Test")
			}

			if Mean_Test(tt, N) {
				fmt.Println("Passed Mean Test")
			} else {
				fmt.Println("Failed Mean Test")
			}

			if Period_Test(tt, K*N) < N {
				fmt.Println("Passed Period Test")
			} else {
				fmt.Println("Failed Period Test")
			}

			if SerialCorrelation_Test(tt, N) {
				fmt.Println("Passed Serial Correlation Test")
			} else {
				fmt.Println("Failed Serial Correlation Test")
			}
		})
	}

}
