package utils

import (
	"sort"
	"time"
)

func Timestamp(dt time.Time) string {
	return dt.Format("2006_01_02_15_04_05")
}

func TimestampNow() string {
	return time.Now().Format("2006_01_02_15_04_05")
}

type FrameStats struct {
	Durations []time.Duration
}

type DurationReport struct {
	Min    time.Duration
	Bot1p  time.Duration
	Bot10p time.Duration
	Avg    time.Duration
	Top10p time.Duration
	Top1p  time.Duration
	Max    time.Duration
}

func (fs *FrameStats) AddFrame(duration time.Duration) {
	fs.Durations = append(fs.Durations, duration)
}

func (fs *FrameStats) Report() DurationReport {
	if len(fs.Durations) == 0 {
		return DurationReport{}
	}

	sort.Slice(fs.Durations, func(i, j int) bool { return fs.Durations[i] < fs.Durations[j] })

	return DurationReport{
		Min:    fs.Durations[0],
		Bot1p:  percentile(fs.Durations, 0.01),
		Bot10p: percentile(fs.Durations, 0.10),
		Avg:    average(fs.Durations),
		Top10p: percentile(fs.Durations, 0.90),
		Top1p:  percentile(fs.Durations, 0.99),
		Max:    fs.Durations[len(fs.Durations)-1],
	}
}

func percentile(durations []time.Duration, p float64) time.Duration {
	index := int(p * float64(len(durations)))
	return durations[index]
}

// average calculates the average duration.
func average(durations []time.Duration) time.Duration {
	var sum time.Duration
	for _, d := range durations {
		sum += d
	}
	return sum / time.Duration(len(durations))
}
