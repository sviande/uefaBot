package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func monitorMatches(previousMatches []Match) []Match {
	currentMatches, err := fetchMatches()
	if err != nil {
		log.Printf("Monitor failed, %v", err)
	}

	newEvents := compareNewInfos(previousMatches, currentMatches)
	log.Printf("%v", newEvents)

	return currentMatches
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	previousMatches := []Match{}

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
