package subtractTimeRanges_test

import (
	"testing"
	"time"
	. "timeRangesSubtractor/subtractTimeRanges"
)

func TestTimeAndType_IsType(t *testing.T) {
	timeAndTypeRs := TimeAndType{0, 0}
	if !timeAndTypeRs.IsType("Rs") {
		t.Error("Expected timeAndTypeRs.IsType(\"Rs\") to be true, got ", timeAndTypeRs.IsType("Rs"))
	}
	if timeAndTypeRs.IsType("Re") {
		t.Error("Expected timeAndTypeRs.IsType(\"Re\") to be false, got ", timeAndTypeRs.IsType("Re"))
	}
	timeAndTypeRe := TimeAndType{0, 1}
	if !timeAndTypeRe.IsType("Re") {
		t.Error("Expected timeAndTypeRe.IsType(\"Re\") to be true, got ", timeAndTypeRe.IsType("Re"))
	}
	if timeAndTypeRe.IsType("Rs") {
		t.Error("Expected timeAndTypeRe.IsType(\"Rs\") to be false, got ", timeAndTypeRe.IsType("Rs"))
	}
	timeAndTypeBs := TimeAndType{0, 2}
	if !timeAndTypeBs.IsType("Bs") {
		t.Error("Expected timeAndTypeBs.IsType(\"Bs\") to be true, got ", timeAndTypeBs.IsType("Bs"))
	}
	if timeAndTypeBs.IsType("Re") {
		t.Error("Expected timeAndTypeBs.IsType(\"Re\") to be false, got ", timeAndTypeBs.IsType("Re"))
	}
	timeAndTypeBe := TimeAndType{0, 3}
	if !timeAndTypeBe.IsType("Be") {
		t.Error("Expected timeAndTypeBe.IsType(\"Be\") to be true, got ", timeAndTypeBe.IsType("Be"))
	}
	if timeAndTypeBe.IsType("Bs") {
		t.Error("Expected timeAndTypeBe.IsType(\"Bs\") to be false, got ", timeAndTypeBe.IsType("Bs"))
	}
}

func TestConvertEntryTypeFromStringToInt(t *testing.T) {
	if ConvertEntryTypeFromStringToInt("Rs") != 0 {
		t.Error("Expected ConvertEntryTypeFromStringToInt(\"Rs\") to be false, got ",
			ConvertEntryTypeFromStringToInt("Rs"))
	}
	if ConvertEntryTypeFromStringToInt("Re") != 1 {
		t.Error("Expected ConvertEntryTypeFromStringToInt(\"Re\") to be false, got ",
			ConvertEntryTypeFromStringToInt("Re"))
	}
	if ConvertEntryTypeFromStringToInt("Bs") != 2 {
		t.Error("Expected ConvertEntryTypeFromStringToInt(\"Bs\") to be false, got ",
			ConvertEntryTypeFromStringToInt("Bs"))
	}
	if ConvertEntryTypeFromStringToInt("Be") != 3 {
		t.Error("Expected ConvertEntryTypeFromStringToInt(\"Be\") to be false, got ",
			ConvertEntryTypeFromStringToInt("Be"))
	}
}

func TestNewTimeAndType(t *testing.T) {
	timeNow := time.Now()
	newTimeAndType := NewTimeAndType(timeNow, "Rs")
	if newTimeAndType.TimeNano != timeNow.UnixNano() {
		t.Error("Expected a TimeAndType entity to be created")
	}
}