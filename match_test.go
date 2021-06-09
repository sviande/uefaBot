package main

import "testing"

func TestCompareNewInfos(t *testing.T) {
	current := make(map[string]MatchInfo)
	current["test"] = MatchInfo{
		ID:     "test",
		Status: MatchLive,
		AwayTeam: AwayTeam{
			InternationalName: "awayTeam",
		},
		HomeTeam: HomeTeam{
			InternationalName: "homeTeam",
		},
	}
	previous := make(map[string]MatchInfo)
	previous["test"] = MatchInfo{
		ID:     "test",
		Status: MatchUpcoming,
	}
	got := compareNewInfos(current, previous)
	expected := []MatchEvent{
		{START, "Match homeTeam : awayTeam started"},
	}

	if got[0].Event != expected[0].Event {
		t.Errorf("CompareNewInfos = %d; want %d", got[0].Event, expected[0].Event)
	}

	if got[0].Label != expected[0].Label {
		t.Errorf("CompareNewInfos = %s; want %s", got[0].Label, expected[0].Label)
	}
}
