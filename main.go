package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func monitorMatches(previousMatches map[string]MatchInfo) map[string]MatchInfo {
	currentMatches, err := fetchMatches()
	if err != nil {
		log.Printf("Monitor failed, %v", err)
	}

	newEvents := compareNewInfos(previousMatches, currentMatches)
	processEvents(newEvents)

	return currentMatches
}

func processEvents(events []MatchEvent) {
	if len(events) == 0 {
		log.Println("Nothing new")
		return
	}

	for _, event := range events {
		log.Printf("New event %d: label: %s\n", event.Event, event.Label)
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	previousMatches := make(map[string]MatchInfo)

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				previousMatches = monitorMatches(previousMatches)
			}
		}
	}()

	<-c
	log.Println("Exit")
}
