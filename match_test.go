package main

import "testing"

func TestCompareNewInfos(t *testing.T) {
	current := []Match{}
	previous := []Match{}
	got := compareNewInfos(current, previous)
	expected := []MatchEvent{
		{START, "Match start"},
	}

	if got[0].Event != expected[0].Event {
		t.Errorf("CompareNewInfos = %d; want %d", got[0].Event, expected[0].Event)
	}

	if got[0].Label != expected[0].Label {
		t.Errorf("CompareNewInfos = %s; want %s", got[0].Label, expected[0].Label)
	}
}
