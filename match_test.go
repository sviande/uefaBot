package main

import "testing"

func TestCompareNewInfosMatchLive(t *testing.T) {
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
		{START, "🏆Match homeTeam : awayTeam started"},
	}

	if got[0].Event != expected[0].Event {
		t.Errorf("CompareNewInfos = %d; want %d", got[0].Event, expected[0].Event)
	}

	if got[0].Label != expected[0].Label {
		t.Errorf("CompareNewInfos = %s; want %s", got[0].Label, expected[0].Label)
	}
}

func TestCompareNewInfosMatchPlayerScore(t *testing.T) {
	current := make(map[string]MatchInfo)
	current["test"] = MatchInfo{
		ID:     "test",
		Status: MatchLive,
		HomeTeam: HomeTeam{
			InternationalName: "hometeam",
		},
		AwayTeam: AwayTeam{
			InternationalName: "awayteam",
		},
		PlayerEvents: PlayerEvents{
			Scorers: []Scorers{
				{
					Player: Player{
						InternationalName: "testPlayer",
					},
					TotalScore: TotalScore{
						Away: 0,
						Home: 1,
					},
				},
			},
		},
	}
	previous := make(map[string]MatchInfo)
	previous["test"] = MatchInfo{
		ID:           "test",
		Status:       MatchLive,
		PlayerEvents: PlayerEvents{},
	}
	got := compareNewInfos(current, previous)
	expected := []MatchEvent{
		{GOAL, "⚽ testPlayer hometeam 1:0 awayteam"},
	}

	if got[0].Event != expected[0].Event {
		t.Errorf("CompareNewInfos = %d; want %d", got[0].Event, expected[0].Event)
	}

	if got[0].Label != expected[0].Label {
		t.Errorf("CompareNewInfos = %s; want %s", got[0].Label, expected[0].Label)
	}
}
