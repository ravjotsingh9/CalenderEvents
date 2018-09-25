package main

import (
	"fmt"
	"testing"
	"time"
)

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
