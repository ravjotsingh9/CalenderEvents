package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dustinkirkland/golang-petname"
)

const (
	SEED = 94608000
)

var (
	max       = 10
	words     = flag.Int("words", 2, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
)

type event struct {
	Start time.Time `json: "start"`
	End   time.Time `json: "end"`
}

type events struct {
	eventsMap map[string]event `json: "events"`
}

func randomTimestamp(seed int64) time.Time {
	randomTime := rand.Int63n(time.Now().Unix()-seed) + seed

	randomNow := time.Unix(randomTime, 0)

	return randomNow
}
func (o *events) init() {
	o.eventsMap = make(map[string]event)
}

func main() {
	// Create a map of events
	var eventList events
	eventList.init()

	// Get the args
	if len(os.Args) > 2 {
		fmt.Println("\nUsage: \n\t./generateData [Required Number of Events]\n")
		return
	}
	if len(os.Args) == 2 {
		var err error
		arg := os.Args[1]
		max, err = strconv.Atoi(arg)
		if err != nil {
			fmt.Println(err)
		}
	}
	// Generate event entries
	i := 0
	for i < max {
		ret := randomTimestamp(SEED)
		pet := petname.Generate(*words, *separator)
		eventList.eventsMap[pet] = event{ret, randomTimestamp(ret.Unix())}

		i++
	}

	// Open a file to dump the events
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Marshall the data
	data, err := json.MarshalIndent(eventList.eventsMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the json to the file
	err = ioutil.WriteFile("data.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("./data.txt created with " + strconv.Itoa(max) + " entries.")
}
