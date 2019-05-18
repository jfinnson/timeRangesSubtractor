package subtractTimeRanges

import (
	"sort"
	"time"
)

// struct to store time ranges. Both values are stored as time.Time
type TimeRange struct {
	Start time.Time  // Range start time
	End time.Time  // Range end time
}

// Subtracts two time range slices
func SubtractTimeRanges(timeRangesA []TimeRange, timeRangesB []TimeRange) []TimeRange {
	var timeRangesResult []TimeRange     // timeRangesA - timeRangesB
	var timeRangesCombined []TimeAndType // timeRangesA starts + timeRangesA ends + timeRangesB starts + timeRangesB ends

	// Combine both timeRangesA and timeRangesB into a single slice
	for _, timeRange := range timeRangesA {
		timeRangesCombined = append(timeRangesCombined, NewTimeAndType(timeRange.Start, "Rs"))
		timeRangesCombined = append(timeRangesCombined, NewTimeAndType(timeRange.End, "Re"))
	}
	for _, timeRange := range timeRangesB {
		timeRangesCombined = append(timeRangesCombined, NewTimeAndType(timeRange.Start, "Bs"))
		timeRangesCombined = append(timeRangesCombined, NewTimeAndType(timeRange.End, "Be"))
	}

	// Sort combined time range starts and ends (with types) by time. Earliest first. O(n * log^2(n))
	sort.Slice(timeRangesCombined, func(i, j int) bool {
		return timeRangesCombined[i].TimeNano < timeRangesCombined[j].TimeNano
	})

	timeRangesResult = extractTimeRanges(timeRangesCombined)

	return timeRangesResult
}

// Extract valid/positive time ranges from a combined list of ranges/positive and blockers/negative times
func extractTimeRanges(timeRangesCombined []TimeAndType) []TimeRange {
	var timeRangesResult []TimeRange

	// Determine time range result of timeRangesA - timeRangesB. O(n)
	// This is done by iterating over timeRangesCombined and extracting new ranges where a range from timeRangesA is
	// active and no ranges from timeRangesB are active.
	// New ranges are extracted when unblocked ranges are ended with either Re (range end) or Bs (block start)
	rangeActiveCount := 0  // Count active range depth. Rs increases count, Re decreases count.
	blockActiveCount := 0  // Count active block depth. Bs increases count, Be decreases count.
	var currentRangeStart int64 = 0
	for _, timeAndType := range timeRangesCombined {

		if timeAndType.IsType("Rs"){ // Range start
			rangeActiveCount ++ // One more range active
			// Set new range start if no blocks active
			if currentRangeStart == 0 && blockActiveCount == 0 {
				currentRangeStart = timeAndType.TimeNano
			}
		}else if timeAndType.IsType("Re"){ // Range end
			rangeActiveCount -- // One less range active
			if currentRangeStart != 0 {
				if currentRangeStart < timeAndType.TimeNano { // Check ensure range end is after range start
					// End found for time range. Cut range and add to result slice.
					timeRangesResult = append(timeRangesResult,
						TimeRange{time.Unix(0, currentRangeStart),
							time.Unix(0, timeAndType.TimeNano)})
				}
				currentRangeStart = 0  // Reset range start
			}
		}else if timeAndType.IsType("Bs"){ // Blocker start
			blockActiveCount ++ // One more block active
			if currentRangeStart != 0 {
				if currentRangeStart < timeAndType.TimeNano { // Check ensure range end is after range start
					// End found for time range. Cut range and add to result slice.
					timeRangesResult = append(timeRangesResult,
						TimeRange{time.Unix(0, currentRangeStart),
							time.Unix(0, timeAndType.TimeNano)})
				}
				currentRangeStart = 0  // Reset range start
			}
		}else if timeAndType.IsType("Be"){ // Blocker end
			blockActiveCount -- // One less block active
			if blockActiveCount == 0 && rangeActiveCount > 0 {
				// No more blocks active
				// Range from timeRangesA is still active
				// Set new currentRangeStart, for the next range to be added to the result
				currentRangeStart = timeAndType.TimeNano // Reset range start
			}
		}
	}

	return timeRangesResult
}