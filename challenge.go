package main

import (
	"fmt"
	"time"
	"timeRangesSubtractor/subtractTimeRanges"
)

func main() {

	// Basic Test using example 5 from the challenge's definition
	layout := "2006-01-02T15:04:05.000Z"
	timeRangeA1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeA1End, _ := time.Parse(layout, "2019-05-17T11:00:00.000Z")
	timeRangeA2Start, _ := time.Parse(layout, "2019-05-17T13:00:00.000Z")
	timeRangeA2End, _ := time.Parse(layout, "2019-05-17T15:00:00.000Z")

	timeRangeB1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeB1End, _ := time.Parse(layout, "2019-05-17T09:15:00.000Z")
	timeRangeB2Start, _ := time.Parse(layout, "2019-05-17T10:00:00.000Z")
	timeRangeB2End, _ := time.Parse(layout, "2019-05-17T10:15:00.000Z")
	timeRangeB3Start, _ := time.Parse(layout, "2019-05-17T12:30:00.000Z")
	timeRangeB3End, _ := time.Parse(layout, "2019-05-17T16:00:00.000Z")

	// Create time rangeA and rangeB
	timeRangesA := []subtractTimeRanges.TimeRange{
		{timeRangeA1Start, timeRangeA1End},
		{timeRangeA2Start, timeRangeA2End}}
	timeRangesB := []subtractTimeRanges.TimeRange{
		{timeRangeB1Start, timeRangeB1End},
		{timeRangeB2Start, timeRangeB2End},
		{timeRangeB3Start, timeRangeB3End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = subtractTimeRanges.SubtractTimeRanges(timeRangesA, timeRangesB)

	// Print out resulting time range.
	for idx, timeRange := range timeRangeResult {
		fmt.Println(fmt.Sprintf("%d: %s - %s", idx, timeRange.Start.UTC(), timeRange.End.UTC()))
	}
}
