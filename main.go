package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
)

func beep() {
	err := beeep.Beep(beeep.DefaultFreq, 5000)
	if err != nil {
		panic(err)
	}
}

func getStartTime() time.Time {
	roundTo, _ := time.ParseDuration("15m")
	addDuration, _ := time.ParseDuration("13m")
	subDuration, _ := time.ParseDuration("2m")
	var startTime time.Time
	if time.Now().Round(roundTo).After(time.Now()) {
		startTime = time.Now().Round(roundTo).Add(-subDuration)
	} else {
		startTime = time.Now().Round(roundTo).Add(addDuration)
	}
	fmt.Println(startTime.Local())
	return startTime
}

func runCronJobs() {
	s := gocron.NewScheduler(time.Now().Location())
	s.Every(15).Minutes().StartAt(getStartTime()).Do(beep)
	// s.Cron("0 13,28,43,58 9-17 ? * MON,TUE,WED,THU,FRI *").Do(beep)
	s.StartBlocking()
}

func main() {
	runCronJobs()
}
