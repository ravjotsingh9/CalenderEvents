package main

import (
	"fmt"
	"testing"
	"time"
	"os"
	"errors"
)

const (
	testDataCorrectPath = "testData/test_data_correct.json"
	testDataEmptyPath   = "testData/test_data_empty.json"
	testDataInvalidPath = "testData/test_data_invalid.json"
	nonExistentFilePath = "testData/non_existent_file.json"
)

// Helper function to parse time and fail test on error
func parseTime(t *testing.T, timeStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		t.Fatalf("Error parsing time string '%s': %v", timeStr, err)
	}
	return parsedTime
}

// TestReadFile_CorrectData tests readFile with correctly formatted JSON data.
func TestReadFile_CorrectData(t *testing.T) {
	eventList, err := readFile(testDataCorrectPath)
	if err != nil {
		t.Fatalf("readFile failed with error: %v", err)
	}

	if len(eventList) != 3 {
		t.Fatalf("expected 3 events, got %d", len(eventList))
	}

	expectedEvents := []eventEntry{
		{
			Name:  "test-event-1",
			Start: parseTime(t, "2023-01-01T10:00:00Z"),
			End:   parseTime(t, "2023-01-01T11:00:00Z"),
		},
		{
			Name:  "test-event-2",
			Start: parseTime(t, "2023-01-01T10:30:00Z"),
			End:   parseTime(t, "2023-01-01T11:30:00Z"),
		},
		{
			Name:  "test-event-3",
			Start: parseTime(t, "2023-01-01T12:00:00Z"),
			End:   parseTime(t, "2023-01-01T13:00:00Z"),
		},
	}

	for _, expected := range expectedEvents {
		found := false
		for _, actual := range eventList {
			if actual.Name == expected.Name {
				found = true
				if !actual.Start.Equal(expected.Start) || !actual.End.Equal(expected.End) {
					t.Errorf("event %s: expected start %v, end %v; got start %v, end %v",
						expected.Name, expected.Start, expected.End, actual.Start, actual.End)
				}
				break
			}
		}
		if !found {
			t.Errorf("expected event %s not found in parsed list", expected.Name)
		}
	}
}

// TestReadFile_EmptyData tests readFile with an empty JSON object.
func TestReadFile_EmptyData(t *testing.T) {
	eventList, err := readFile(testDataEmptyPath)
	if err != nil {
		t.Fatalf("readFile failed with error: %v", err)
	}

	if len(eventList) != 0 {
		t.Fatalf("expected 0 events, got %d", len(eventList))
	}
}

// TestReadFile_InvalidData tests readFile with invalid JSON data.
func TestReadFile_InvalidData(t *testing.T) {
	_, err := readFile(testDataInvalidPath)
	if err == nil {
		t.Fatalf("readFile should have failed for invalid JSON, but it did not")
	}
}

// TestReadFile_FileNotFound tests readFile with a non-existent file.
func TestReadFile_FileNotFound(t *testing.T) {
	_, err := readFile(nonExistentFilePath)
	if err == nil {
		t.Fatalf("readFile should have failed for a non-existent file, but it did not")
	}
	if !os.IsNotExist(errors.Unwrap(err)) { // check for underlying os.ErrNotExist
		t.Fatalf("expected a file not found error, got %v", err)
	}
}

// NO overlapping
func Test_ShouldReturnNoOverlappedEvent_WhenNoOverlapping(t *testing.T) {
	// Data format and output map
	defaultFormat := "2006-01-02 15:04:05"

	time1, _ := time.Parse(defaultFormat, "2018-01-02 01:04:05")
	time2, _ := time.Parse(defaultFormat, "2018-01-02 02:04:05")
	time3, _ := time.Parse(defaultFormat, "2018-01-02 03:04:05")
	time4, _ := time.Parse(defaultFormat, "2018-01-02 04:04:05")
	time5, _ := time.Parse(defaultFormat, "2018-01-02 05:04:05")
	time6, _ := time.Parse(defaultFormat, "2018-01-02 06:04:05")
	time7, _ := time.Parse(defaultFormat, "2018-01-02 07:04:05")
	time8, _ := time.Parse(defaultFormat, "2018-01-02 08:04:05")
	time9, _ := time.Parse(defaultFormat, "2018-01-02 09:04:05")
	time10, _ := time.Parse(defaultFormat, "2018-01-02 10:04:05")

	// Read the file and receive data in an Array
	eventList := []eventEntry{
		{
			"Running",
			time1,
			time2,
		},
		{
			"walking",
			time3,
			time4,
		},
		{
			"waiting",
			time5,
			time6,
		},
		{
			"watching",
			time7,
			time8,
		},
		{
			"playing",
			time9,
			time10,
		},
	}

	overlappedEvents, err := findConflicts(eventList)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(overlappedEvents) != 0 {
		t.Fail()
	}
}

// Same end and start time
func Test_ShouldReturnNoOverlappedEvent_WhenEndAndStartTimesOfTwoEventsAreSame(t *testing.T) {
	// Data format and output map
	defaultFormat := "2006-01-02 15:04:05"

	time1, _ := time.Parse(defaultFormat, "2018-01-02 01:04:05")
	time2, _ := time.Parse(defaultFormat, "2018-01-02 02:04:05")
	//time3, _ := time.Parse(defaultFormat, "2018-01-02 03:04:05")
	time4, _ := time.Parse(defaultFormat, "2018-01-02 04:04:05")
	time5, _ := time.Parse(defaultFormat, "2018-01-02 05:04:05")
	time6, _ := time.Parse(defaultFormat, "2018-01-02 06:04:05")
	time7, _ := time.Parse(defaultFormat, "2018-01-02 07:04:05")
	time8, _ := time.Parse(defaultFormat, "2018-01-02 08:04:05")
	time9, _ := time.Parse(defaultFormat, "2018-01-02 09:04:05")
	time10, _ := time.Parse(defaultFormat, "2018-01-02 10:04:05")

	// Read the file and receive data in an Array
	eventList := []eventEntry{
		{
			"Running",
			time1,
			time2,
		},
		{
			"walking",
			time2,
			time4,
		},
		{
			"waiting",
			time5,
			time6,
		},
		{
			"watching",
			time7,
			time8,
		},
		{
			"playing",
			time9,
			time10,
		},
	}

	overlappedEvents, err := findConflicts(eventList)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(overlappedEvents) != 0 {
		t.Fail()
	}
}

// One overlap
func Test_ShouldReturnTwoOverlappedEvent_WhenTwoEventsOverlapped(t *testing.T) {
	// Data format and output map
	defaultFormat := "2006-01-02 15:04:05"

	time1, _ := time.Parse(defaultFormat, "2018-01-02 01:04:05")
	time2, _ := time.Parse(defaultFormat, "2018-01-02 02:04:05")
	time3, _ := time.Parse(defaultFormat, "2018-01-02 03:04:05")
	time4, _ := time.Parse(defaultFormat, "2018-01-02 04:04:05")
	time5, _ := time.Parse(defaultFormat, "2018-01-02 05:04:05")
	time6, _ := time.Parse(defaultFormat, "2018-01-02 06:04:05")
	time7, _ := time.Parse(defaultFormat, "2018-01-02 07:04:05")
	time8, _ := time.Parse(defaultFormat, "2018-01-02 08:04:05")
	time9, _ := time.Parse(defaultFormat, "2018-01-02 09:04:05")
	time10, _ := time.Parse(defaultFormat, "2018-01-02 10:04:05")

	// Read the file and receive data in an Array
	eventList := []eventEntry{
		{
			"Running",
			time1,
			time3,
		},
		{
			"walking",
			time2,
			time4,
		},
		{
			"waiting",
			time5,
			time6,
		},
		{
			"watching",
			time7,
			time8,
		},
		{
			"playing",
			time9,
			time10,
		},
	}

	overlappedEvents, err := findConflicts(eventList)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, ok1 := overlappedEvents["Running"]
	_, ok2 := overlappedEvents["walking"]
	if !((len(overlappedEvents) == 2) && (ok1) && (ok2)) {
		t.Fail()
	}
}

// three overlap
func Test_ShouldReturnTwoOverlappedEvent_WhenThreeEventsOverlapped(t *testing.T) {
	// Data format and output map
	defaultFormat := "2006-01-02 15:04:05"

	time1, _ := time.Parse(defaultFormat, "2018-01-02 01:04:05")
	time2, _ := time.Parse(defaultFormat, "2018-01-02 02:04:05")
	time3, _ := time.Parse(defaultFormat, "2018-01-02 03:04:05")
	time4, _ := time.Parse(defaultFormat, "2018-01-02 04:04:05")
	time5, _ := time.Parse(defaultFormat, "2018-01-02 05:04:05")
	time6, _ := time.Parse(defaultFormat, "2018-01-02 06:04:05")
	time7, _ := time.Parse(defaultFormat, "2018-01-02 07:04:05")
	time8, _ := time.Parse(defaultFormat, "2018-01-02 08:04:05")
	time9, _ := time.Parse(defaultFormat, "2018-01-02 09:04:05")
	time10, _ := time.Parse(defaultFormat, "2018-01-02 10:04:05")

	// Read the file and receive data in an Array
	eventList := []eventEntry{
		{
			"Running",
			time1,
			time5,
		},
		{
			"walking",
			time2,
			time3,
		},
		{
			"waiting",
			time4,
			time6,
		},
		{
			"watching",
			time7,
			time8,
		},
		{
			"playing",
			time9,
			time10,
		},
	}

	overlappedEvents, err := findConflicts(eventList)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, ok1 := overlappedEvents["Running"]
	_, ok2 := overlappedEvents["walking"]
	_, ok3 := overlappedEvents["waiting"]
	if !((len(overlappedEvents) == 3) && (ok1) && (ok2) && (ok3)) {
		t.Fail()
	}
}
