package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func updateMatches(competitionID string, matchEventChan chan MatchEvent, previousMatches map[string]MatchInfo) map[string]MatchInfo {
	currentMatches, err := fetchMatches(competitionID)
	if err != nil {
		log.Printf("Fetch failed, %v", err)
		return previousMatches
	}

	newEvents := compareNewInfos(currentMatches, previousMatches)
	for _, event := range newEvents {
		matchEventChan <- event
	}

	return currentMatches
}

func monitorMatches(matchEventChan chan MatchEvent, competitionID string) {
	ticker := time.NewTicker(1 * time.Minute)
	previousMatches := make(map[string]MatchInfo)
	previousMatches = updateMatches(competitionID, matchEventChan, previousMatches)
	for {
		select {
		case <-ticker.C:
			previousMatches = updateMatches(competitionID, matchEventChan, previousMatches)
		}
	}
}

func sendSlackMessages(slack Slack, events chan MatchEvent) {
	for event := range events {
		log.Printf("Event %d: %s\n", event.Event, event.Label)
		slack.sendMessage(event.Label)
	}
}

func main() {
	webHookURL := os.Getenv("WEBHOOK_URL")
	if webHookURL == "" {
		log.Fatal("Missing env var WEBHOOK_URL")
	}

	competitionID := ""
	flag.StringVar(&competitionID, "competitionId", "", "UEFA competionId to monitor. Find all competions at https://comp.uefa.com/v1/competitions")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	slack := Slack{
		hookURL: webHookURL,
	}
	slack.sendMessage("🔥Bot is up")

	matchEventChan := make(chan MatchEvent, 10)
	go monitorMatches(matchEventChan, competitionID)
	go sendSlackMessages(slack, matchEventChan)

	<-c
	close(matchEventChan)
	log.Println("Exit")
}
