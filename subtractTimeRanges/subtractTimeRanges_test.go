package subtractTimeRanges_test

import (
	"testing"
	"time"
	. "timeRangesSubtractor/subtractTimeRanges"
)

// Test using example 1 from the challenge's definition
func TestSubtractTimeRanges_example1(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	timeRangeA1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeA1End, _ := time.Parse(layout, "2019-05-17T10:00:00.000Z")

	timeRangeB1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeB1End, _ := time.Parse(layout, "2019-05-17T09:30:00.000Z")

	timeRangesA := []TimeRange{{timeRangeA1Start, timeRangeA1End}}
	timeRangesB := []TimeRange{{timeRangeB1Start, timeRangeB1End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = SubtractTimeRanges(timeRangesA, timeRangesB)

	if len(timeRangeResult) != 1 {
		t.Error("Expected timeRangeResult to have 1 range, got ", len(timeRangeResult) )
	}

	if timeRangeResult[0].Start.UTC().String() != "2019-05-17 09:30:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:30:00 +0000 UTC, got ",
			timeRangeResult[0].Start.UTC().String() )
	}
	if timeRangeResult[0].End.UTC().String() != "2019-05-17 10:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 10:00:00 +0000 UTC, got ",
			timeRangeResult[0].End.UTC().String() )
	}
}

// Test using example 2 from the challenge's definition
func TestSubtractTimeRanges_example2(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	timeRangeA1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeA1End, _ := time.Parse(layout, "2019-05-17T10:00:00.000Z")

	timeRangeB1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeB1End, _ := time.Parse(layout, "2019-05-17T10:00:00.000Z")

	timeRangesA := []TimeRange{{timeRangeA1Start, timeRangeA1End}}
	timeRangesB := []TimeRange{{timeRangeB1Start, timeRangeB1End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = SubtractTimeRanges(timeRangesA, timeRangesB)

	if len(timeRangeResult) != 0 {
		t.Error("Expected timeRangeResult to have 0 ranges, got ", len(timeRangeResult) )
	}
}

// Test using example 3 from the challenge's definition
func TestSubtractTimeRanges_example3(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	timeRangeA1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeA1End, _ := time.Parse(layout, "2019-05-17T09:30:00.000Z")

	timeRangeB1Start, _ := time.Parse(layout, "2019-05-17T09:30:00.000Z")
	timeRangeB1End, _ := time.Parse(layout, "2019-05-17T15:00:00.000Z")

	timeRangesA := []TimeRange{{timeRangeA1Start, timeRangeA1End}}
	timeRangesB := []TimeRange{{timeRangeB1Start, timeRangeB1End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = SubtractTimeRanges(timeRangesA, timeRangesB)

	if len(timeRangeResult) != 1 {
		t.Error("Expected timeRangeResult to have 1 range, got ", len(timeRangeResult) )
	}

	if timeRangeResult[0].Start.UTC().String() != "2019-05-17 09:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:00:00 +0000 UTC, got ",
			timeRangeResult[0].Start.UTC().String() )
	}
	if timeRangeResult[0].End.UTC().String() != "2019-05-17 09:30:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:30:00 +0000 UTC, got ",
			timeRangeResult[0].End.UTC().String() )
	}
}

// Test using example 4 from the challenge's definition
func TestSubtractTimeRanges_example4(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	timeRangeA1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeA1End, _ := time.Parse(layout, "2019-05-17T09:30:00.000Z")
	timeRangeA2Start, _ := time.Parse(layout, "2019-05-17T10:00:00.000Z")
	timeRangeA2End, _ := time.Parse(layout, "2019-05-17T10:30:00.000Z")

	timeRangeB1Start, _ := time.Parse(layout, "2019-05-17T09:15:00.000Z")
	timeRangeB1End, _ := time.Parse(layout, "2019-05-17T10:15:00.000Z")

	timeRangesA := []TimeRange{
		{timeRangeA1Start, timeRangeA1End},
		{timeRangeA2Start, timeRangeA2End}}
	timeRangesB := []TimeRange{{timeRangeB1Start, timeRangeB1End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = SubtractTimeRanges(timeRangesA, timeRangesB)

	if len(timeRangeResult) != 2 {
		t.Error("Expected timeRangeResult to have 2 ranges, got ", len(timeRangeResult) )
	}

	if timeRangeResult[0].Start.UTC().String() != "2019-05-17 09:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:00:00 +0000 UTC, got ",
			timeRangeResult[0].Start.UTC().String() )
	}
	if timeRangeResult[0].End.UTC().String() != "2019-05-17 09:15:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:15:00 +0000 UTC, got ",
			timeRangeResult[0].End.UTC().String() )
	}
	if timeRangeResult[1].Start.UTC().String() != "2019-05-17 10:15:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 09:00:00 +0000 UTC, got ",
			timeRangeResult[1].Start.UTC().String() )
	}
	if timeRangeResult[1].End.UTC().String() != "2019-05-17 10:30:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 10:30:00 +0000 UTC, got ",
			timeRangeResult[1].End.UTC().String() )
	}
}

// Test using example 5 from the challenge's definition
func TestSubtractTimeRanges_example5(t *testing.T) {
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

	timeRangesA := []TimeRange{
		{timeRangeA1Start, timeRangeA1End},
		{timeRangeA2Start, timeRangeA2End}}
	timeRangesB := []TimeRange{
		{timeRangeB1Start, timeRangeB1End},
		{timeRangeB2Start, timeRangeB2End},
		{timeRangeB3Start, timeRangeB3End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = SubtractTimeRanges(timeRangesA, timeRangesB)

	if len(timeRangeResult) != 2 {
		t.Error("Expected timeRangeResult to have 2 ranges, got ", len(timeRangeResult) )
	}

	if timeRangeResult[0].Start.UTC().String() != "2019-05-17 09:15:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:15:00 +0000 UTC, got ",
			timeRangeResult[0].Start.UTC().String() )
	}
	if timeRangeResult[0].End.UTC().String() != "2019-05-17 10:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 10:00:00 +0000 UTC, got ",
			timeRangeResult[0].End.UTC().String() )
	}
	if timeRangeResult[1].Start.UTC().String() != "2019-05-17 10:15:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 10:15:00 +0000 UTC, got ",
			timeRangeResult[1].Start.UTC().String() )
	}
	if timeRangeResult[1].End.UTC().String() != "2019-05-17 11:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 11:00:00 +0000 UTC, got ",
			timeRangeResult[1].End.UTC().String() )
	}
}

// Test rangesA - rangesB when rangesB removes multiple "slices" of the first range in rangesA
// This test ensures that the approach used in SubtractTimeRanges can divide up a single time ranges into
// smaller ranges without issues.
func TestSubtractTimeRanges_multiple_nested_example(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	timeRangeA1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeA1End, _ := time.Parse(layout, "2019-05-17T11:00:00.000Z")
	timeRangeA2Start, _ := time.Parse(layout, "2019-05-17T13:00:00.000Z")
	timeRangeA2End, _ := time.Parse(layout, "2019-05-17T15:00:00.000Z")

	timeRangeB1Start, _ := time.Parse(layout, "2019-05-17T09:01:00.000Z")
	timeRangeB1End, _ := time.Parse(layout, "2019-05-17T09:02:00.000Z")
	timeRangeB2Start, _ := time.Parse(layout, "2019-05-17T09:02:20.000Z")
	timeRangeB2End, _ := time.Parse(layout, "2019-05-17T09:02:40.000Z")
	timeRangeB3Start, _ := time.Parse(layout, "2019-05-17T09:05:00.000Z")
	timeRangeB3End, _ := time.Parse(layout, "2019-05-17T09:05:10.000Z")

	timeRangesA := []TimeRange{
		{timeRangeA1Start, timeRangeA1End},
		{timeRangeA2Start, timeRangeA2End}}
	timeRangesB := []TimeRange{
		{timeRangeB1Start, timeRangeB1End},
		{timeRangeB2Start, timeRangeB2End},
		{timeRangeB3Start, timeRangeB3End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = SubtractTimeRanges(timeRangesA, timeRangesB)

	if len(timeRangeResult) != 5 {
		t.Error("Expected timeRangeResult to have 5 ranges, got ", len(timeRangeResult) )
	}

	if timeRangeResult[0].Start.UTC().String() != "2019-05-17 09:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:00:00 +0000 UTC, got ",
			timeRangeResult[0].Start.UTC().String() )
	}
	if timeRangeResult[0].End.UTC().String() != "2019-05-17 09:01:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:01:00 +0000 UTC, got ",
			timeRangeResult[0].End.UTC().String() )
	}
	if timeRangeResult[1].Start.UTC().String() != "2019-05-17 09:02:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 09:02:00 +0000 UTC, got ",
			timeRangeResult[1].Start.UTC().String() )
	}
	if timeRangeResult[1].End.UTC().String() != "2019-05-17 09:02:20 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 09:20:00 +0000 UTC, got ",
			timeRangeResult[1].End.UTC().String() )
	}
	if timeRangeResult[2].Start.UTC().String() != "2019-05-17 09:02:40 +0000 UTC" {
		t.Error("Expected timeRangeResult 3 Start to be 2019-05-17 09:40:00 +0000 UTC, got ",
			timeRangeResult[2].Start.UTC().String() )
	}
	if timeRangeResult[2].End.UTC().String() != "2019-05-17 09:05:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 3 Start to be 2019-05-17 09:05:00 +0000 UTC, got ",
			timeRangeResult[2].End.UTC().String() )
	}
	if timeRangeResult[3].Start.UTC().String() != "2019-05-17 09:05:10 +0000 UTC" {
		t.Error("Expected timeRangeResult 4 Start to be 2019-05-17 09:05:10 +0000 UTC, got ",
			timeRangeResult[3].Start.UTC().String() )
	}
	if timeRangeResult[3].End.UTC().String() != "2019-05-17 11:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 4 Start to be 2019-05-17 11:00:00 +0000 UTC, got ",
			timeRangeResult[3].End.UTC().String() )
	}
	if timeRangeResult[4].Start.UTC().String() != "2019-05-17 13:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 5 Start to be 2019-05-17 13:00:00 +0000 UTC, got ",
			timeRangeResult[4].Start.UTC().String() )
	}
	if timeRangeResult[4].End.UTC().String() != "2019-05-17 15:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 5 Start to be 2019-05-17 15:00:00 +0000 UTC, got ",
			timeRangeResult[4].End.UTC().String() )
	}
}

// Test time ranges subtraction when there is overlap "within" ranges A or "within" ranges B
// This test ensures that they "counting" approach used in SubtractTimeRanges works as intended,
// only allowing the creation of a range if all blocking ranges are passed.
func TestSubtractTimeRanges_multiple_layers(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	timeRangeA1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeA1End, _ := time.Parse(layout, "2019-05-17T11:00:00.000Z")
	timeRangeA2Start, _ := time.Parse(layout, "2019-05-17T08:00:00.000Z")
	timeRangeA2End, _ := time.Parse(layout, "2019-05-17T10:00:00.000Z")

	timeRangeB1Start, _ := time.Parse(layout, "2019-05-17T09:00:00.000Z")
	timeRangeB1End, _ := time.Parse(layout, "2019-05-17T09:15:00.000Z")
	timeRangeB2Start, _ := time.Parse(layout, "2019-05-17T10:00:00.000Z")
	timeRangeB2End, _ := time.Parse(layout, "2019-05-17T10:15:00.000Z")
	timeRangeB3Start, _ := time.Parse(layout, "2019-05-17T09:10:00.000Z")
	timeRangeB3End, _ := time.Parse(layout, "2019-05-17T09:25:00.000Z")

	timeRangesA := []TimeRange{
		{timeRangeA1Start, timeRangeA1End},
		{timeRangeA2Start, timeRangeA2End}}
	timeRangesB := []TimeRange{
		{timeRangeB1Start, timeRangeB1End},
		{timeRangeB2Start, timeRangeB2End},
		{timeRangeB3Start, timeRangeB3End}}

	// Subtract time ranges. (timeRangesA - timeRangesB)
	var timeRangeResult = SubtractTimeRanges(timeRangesA, timeRangesB)

	if len(timeRangeResult) != 3 {
		t.Error("Expected timeRangeResult to have 3 ranges, got ", len(timeRangeResult) )
	}

	if timeRangeResult[0].Start.UTC().String() != "2019-05-17 08:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 08:00:00 +0000 UTC, got ",
			timeRangeResult[0].Start.UTC().String() )
	}
	if timeRangeResult[0].End.UTC().String() != "2019-05-17 09:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 1 Start to be 2019-05-17 09:00:00 +0000 UTC, got ",
			timeRangeResult[0].End.UTC().String() )
	}
	if timeRangeResult[1].Start.UTC().String() != "2019-05-17 09:25:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 09:25:00 +0000 UTC, got ",
			timeRangeResult[1].Start.UTC().String() )
	}
	if timeRangeResult[1].End.UTC().String() != "2019-05-17 10:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 10:00:00 +0000 UTC, got ",
			timeRangeResult[1].End.UTC().String() )
	}
	if timeRangeResult[2].Start.UTC().String() != "2019-05-17 10:15:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 10:15:00 +0000 UTC, got ",
			timeRangeResult[2].Start.UTC().String() )
	}
	if timeRangeResult[2].End.UTC().String() != "2019-05-17 11:00:00 +0000 UTC" {
		t.Error("Expected timeRangeResult 2 Start to be 2019-05-17 11:00:00 +0000 UTC, got ",
			timeRangeResult[2].End.UTC().String() )
	}
}