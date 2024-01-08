package intervals

import "sort"

func FindCover[T Number](intervals [][2]T) []T {
	switch {
	case len(intervals) <= 10_000:
		return FindCover2(intervals)
	case len(intervals) <= 100_000:
		return FindCover3(intervals)
	default:
		return FindCover1(intervals)
	}
}

func FindCover1[T Number](intervals [][2]T) []T {
	if len(intervals) == 0 {
		return []T{}
	}

	// Sort intervals by their starting point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result []T
	start, end := intervals[0][0], intervals[0][1]

	for _, interval := range intervals {
		if interval[0]-1 <= end {
			// Overlapping intervals, move the end if needed
			if interval[1] > end {
				end = interval[1]
			}
		} else {
			// Non-overlapping interval, add the previous interval and reset start and end
			result = append(result, start, end)
			start, end = interval[0], interval[1]
		}
	}

	// Add the last interval
	result = append(result, start, end)
	return result
}

func FindCover2[T Number](intervals [][2]T) []T {
	if len(intervals) == 0 || len(intervals) >= 20_000 {
		return nil
	}

	// Custom insertion sort for potentially better performance in nearly-sorted data
	for i := 1; i < len(intervals); i++ {
		j := i
		for j > 0 && intervals[j][0] < intervals[j-1][0] {
			intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
			j--
		}
	}

	// In-place merge intervals
	idx := 0
	for _, interval := range intervals {
		if interval[0]-1 <= intervals[idx][1] {
			if interval[1] > intervals[idx][1] {
				intervals[idx][1] = interval[1]
			}
		} else {
			idx++
			intervals[idx] = interval
		}
	}

	// Flatten the intervals into a single slice
	result := make([]T, 0, 2*(idx+1))
	for i := 0; i <= idx; i++ {
		result = append(result, intervals[i][0], intervals[i][1])
	}

	return result
}

func FindCover3[T Number](intervals [][2]T) []T {
	if len(intervals) == 0 {
		return nil
	}

	// Sort intervals by their starting point
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// In-place merge intervals
	idx := 0
	for _, interval := range intervals {
		if interval[0]-1 <= intervals[idx][1] {
			if interval[1] > intervals[idx][1] {
				intervals[idx][1] = interval[1]
			}
		} else {
			idx++
			intervals[idx] = interval
		}
	}

	// Flatten the intervals into a single slice
	result := make([]T, 0, 2*(idx+1))
	for i := 0; i <= idx; i++ {
		result = append(result, intervals[i][0], intervals[i][1])
	}

	return result
}

func FindCover4[T Number](intervals [][2]T) []T {
	if len(intervals) == 0 {
		return []T{}
	}

	// Sort intervals by their starting point
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result []T
	start, end := intervals[0][0], intervals[0][1]

	for _, interval := range intervals {
		if interval[0]-1 <= end {
			// Overlapping intervals, move the end if needed
			if interval[1] > end {
				end = interval[1]
			}
		} else {
			// Non-overlapping interval, add the previous interval and reset start and end
			result = append(result, start, end)
			start, end = interval[0], interval[1]
		}
	}

	// Add the last interval
	result = append(result, start, end)
	return result
}
