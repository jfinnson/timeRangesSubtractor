package subtractTimeRanges

import "time"

// Type to store both a nano timestamp from a time range and the type of stamp it represents.
// EntryType will be used to identify what type of action should be taken when determining new time ranges.
type TimeAndType struct {
	TimeNano int64
	// Represents if the time is:
	// 0: +Range start (Rs)
	// 1: +Range end (Re)
	// 2: -Blocker start (Bs)
	// 3: -Blocker end (Be)
	EntryType int
}

// Convert string representation of "entryType" to int
func ConvertEntryTypeFromStringToInt(entryTypeString string) int {
	var entryType int
	switch entryTypeString {
	case "Rs":
		entryType = 0
	case "Re":
		entryType = 1
	case "Bs":
		entryType = 2
	case "Be":
		entryType = 3
	}

	return entryType
}

// Convenience function to determine if timeAndType's entity type is the same as the string type passed
func (timeAndType TimeAndType) IsType(typeToCompare string) bool {
	return timeAndType.EntryType == ConvertEntryTypeFromStringToInt(typeToCompare)
}

// Create TimeAndType from time.Time and "entry type" (Rs, Re, Bs, or Be)
func NewTimeAndType(timeEntity time.Time, entryTypeString string) TimeAndType {
	var newTimeAndType TimeAndType
	var entryType int

	entryType = ConvertEntryTypeFromStringToInt(entryTypeString)
	newTimeAndType = TimeAndType{timeEntity.UnixNano(), entryType}

	return newTimeAndType
}
