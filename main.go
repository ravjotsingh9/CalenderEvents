package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"time"

	"github.com/tatsushid/go-prettytable"
)

const (
	FILEPATH = "./testData/data.txt"
)

type event struct {
	Start time.Time `json: "start"`
	End   time.Time `json:: "end"`
}

type events struct {
	eventsMap map[string]event `json: "events"`
}

type eventEntry struct {
	Name  string    `json: "name"`
	Start time.Time `json: "start"`
	End   time.Time `json:: "end"`
}

var (
	defaultFormat string
)

// tells if t1 is before t2
func compareTime(t1, t2 time.Time) bool {
	return t2.After(t1) && t1.Before(t2)
}

func printEvents(eventList map[string]event, str string) {

	fmt.Println("\n" + str + "\n")

	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "EVENT NAME"},
		{Header: "START TIME"},
		{Header: "END TIME"},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "

	for e, val := range eventList {
		tbl.AddRow(e, val.Start.Format(defaultFormat), val.End.Format(defaultFormat))
	}
	if len(eventList) < 1 {
		tbl.Print()
		fmt.Println("No Rows found.")
		return
	}
	tbl.Print()
}

// Reads file
func readFile(path string) ([]eventEntry, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var eventJSON events
	json.Unmarshal(dat, &eventJSON.eventsMap)
	i := 0
	var eventList = make([]eventEntry, len(eventJSON.eventsMap))
	for e, val := range eventJSON.eventsMap {
		var tmp eventEntry
		tmp.Name = e
		tmp.Start = val.Start
		tmp.End = val.End
		eventList[i] = tmp
		i++
	}
	return eventList, nil
}

/*
 * Given a sequence of events, each having a start and end time, write a program that will return the sequence of all pairs of overlapping events.
 *
 */

func findConflicts(eventList []eventEntry) (map[string]event, error) {
	overlappedEvents := make(map[string]event)
	if len(eventList) < 1 {
		return nil, errors.New("Error: Invalid input data.")
	}
	if len(eventList) == 1 {
		return nil, nil
	}

	// Sort the array
	sort.Slice(eventList[:], func(i, j int) bool {
		return compareTime(eventList[i].Start, eventList[j].Start)
	})

	// Identify overlaps
	for i := 1; i < len(eventList); i++ {
		if !(compareTime(eventList[i-1].End, eventList[i].Start) || eventList[i-1].End.Equal(eventList[i].Start)) {
			if _, ok := overlappedEvents[eventList[i-1].Name]; !ok {
				overlappedEvents[eventList[i-1].Name] = event{eventList[i-1].Start, eventList[i-1].End}
			}
			if _, ok := overlappedEvents[eventList[i].Name]; !ok {
				overlappedEvents[eventList[i].Name] = event{eventList[i].Start, eventList[i].End}
			}
		}
	}
	return overlappedEvents, nil

}

func main() {

	// Data format and output map
	defaultFormat = "2006-01-02 15:04:05"

	// Read the file and receive data in an Array
	eventList, err := readFile(FILEPATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	overlappedEvents, err := findConflicts(eventList)
	if err != nil {
		fmt.Println(err)
		return
	}

	printEvents(overlappedEvents, "------------ Overlapped Events ------------")
}
