package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

//Match struct
type Match struct {
	AwayTeam struct {
		AssociationLogoURL string `json:"associationLogoUrl"`
		BigLogoURL         string `json:"bigLogoUrl"`
		CountryCode        string `json:"countryCode"`
		ID                 string `json:"id"`
		IDProvider         string `json:"idProvider"`
		InternationalName  string `json:"internationalName"`
		IsPlaceHolder      bool   `json:"isPlaceHolder"`
		LogoURL            string `json:"logoUrl"`
		MediumLogoURL      string `json:"mediumLogoUrl"`
		TeamCode           string `json:"teamCode"`
		TeamTypeDetail     string `json:"teamTypeDetail"`
		Translations       struct {
			CountryName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"countryName"`
			DisplayName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"displayName"`
			DisplayOfficialName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"displayOfficialName"`
			DisplayTeamCode struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"displayTeamCode"`
		} `json:"translations"`
		TypeIsNational bool   `json:"typeIsNational"`
		TypeTeam       string `json:"typeTeam"`
	} `json:"awayTeam"`
	BehindClosedDoors bool `json:"behindClosedDoors"`
	Competition       struct {
		Age    string `json:"age"`
		Code   string `json:"code"`
		ID     string `json:"id"`
		Images struct {
			FullLogo string `json:"FULL_LOGO"`
		} `json:"images"`
		MetaData struct {
			Name string `json:"name"`
		} `json:"metaData"`
		Region       string `json:"region"`
		Sex          string `json:"sex"`
		SportsType   string `json:"sportsType"`
		TeamCategory string `json:"teamCategory"`
		Translations struct {
			Name struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"name"`
			QualifyingName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"qualifyingName"`
			TournamentName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"tournamentName"`
		} `json:"translations"`
		Type string `json:"type"`
	} `json:"competition"`
	Group struct {
		CompetitionID string `json:"competitionId"`
		ID            string `json:"id"`
		MetaData      struct {
			GroupName      string `json:"groupName"`
			GroupShortName string `json:"groupShortName"`
		} `json:"metaData"`
		Order        int    `json:"order"`
		Phase        string `json:"phase"`
		RoundID      string `json:"roundId"`
		SeasonYear   string `json:"seasonYear"`
		Translations struct {
			Name struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"name"`
			ShortName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"shortName"`
		} `json:"translations"`
		Type string `json:"type"`
	} `json:"group"`
	HomeTeam struct {
		AssociationLogoURL string `json:"associationLogoUrl"`
		BigLogoURL         string `json:"bigLogoUrl"`
		CountryCode        string `json:"countryCode"`
		ID                 string `json:"id"`
		IDProvider         string `json:"idProvider"`
		InternationalName  string `json:"internationalName"`
		IsPlaceHolder      bool   `json:"isPlaceHolder"`
		LogoURL            string `json:"logoUrl"`
		MediumLogoURL      string `json:"mediumLogoUrl"`
		TeamCode           string `json:"teamCode"`
		TeamTypeDetail     string `json:"teamTypeDetail"`
		Translations       struct {
			CountryName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"countryName"`
			DisplayName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"displayName"`
			DisplayOfficialName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"displayOfficialName"`
			DisplayTeamCode struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"displayTeamCode"`
		} `json:"translations"`
		TypeIsNational bool   `json:"typeIsNational"`
		TypeTeam       string `json:"typeTeam"`
	} `json:"homeTeam"`
	ID          string `json:"id"`
	KickOffTime struct {
		Date             string    `json:"date"`
		DateTime         time.Time `json:"dateTime"`
		UtcOffsetInHours int       `json:"utcOffsetInHours"`
	} `json:"kickOffTime"`
	LineupStatus string `json:"lineupStatus"`
	MatchNumber  int    `json:"matchNumber"`
	Matchday     struct {
		CompetitionID  string `json:"competitionId"`
		DateFrom       string `json:"dateFrom"`
		DateTo         string `json:"dateTo"`
		ID             string `json:"id"`
		LongName       string `json:"longName"`
		Name           string `json:"name"`
		RoundID        string `json:"roundId"`
		SeasonYear     string `json:"seasonYear"`
		SequenceNumber string `json:"sequenceNumber"`
		Translations   struct {
			LongName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"longName"`
			Name struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"name"`
		} `json:"translations"`
		Type string `json:"type"`
	} `json:"matchday"`
	PlayerEvents struct {
		Scorers []struct {
			GoalType string `json:"goalType"`
			ID       string `json:"id"`
			Images   struct {
				PlayerCelebrating string `json:"PLAYER_CELEBRATING"`
			} `json:"images"`
			Player struct {
				Age                   string `json:"age"`
				BirthDate             string `json:"birthDate"`
				CountryCode           string `json:"countryCode"`
				DetailedFieldPosition string `json:"detailedFieldPosition"`
				FieldPosition         string `json:"fieldPosition"`
				Height                int64  `json:"height"`
				ID                    string `json:"id"`
				ImageURL              string `json:"imageUrl"`
				InternationalName     string `json:"internationalName"`
				Translations          struct {
					CountryName struct {
						Az string `json:"AZ"`
						Da string `json:"DA"`
						De string `json:"DE"`
						En string `json:"EN"`
						Es string `json:"ES"`
						Fr string `json:"FR"`
						Hu string `json:"HU"`
						It string `json:"IT"`
						Ja string `json:"JA"`
						Nl string `json:"NL"`
						Pt string `json:"PT"`
						Ro string `json:"RO"`
						Ru string `json:"RU"`
						Zh string `json:"ZH"`
					} `json:"countryName"`
					FieldPosition struct {
						Az string `json:"AZ"`
						Da string `json:"DA"`
						De string `json:"DE"`
						En string `json:"EN"`
						Es string `json:"ES"`
						Fr string `json:"FR"`
						Hu string `json:"HU"`
						It string `json:"IT"`
						Ja string `json:"JA"`
						Nl string `json:"NL"`
						Pt string `json:"PT"`
						Ro string `json:"RO"`
						Ru string `json:"RU"`
						Zh string `json:"ZH"`
					} `json:"fieldPosition"`
					FirstName struct {
						Az string `json:"AZ"`
						Da string `json:"DA"`
						De string `json:"DE"`
						En string `json:"EN"`
						Es string `json:"ES"`
						Fr string `json:"FR"`
						Hu string `json:"HU"`
						It string `json:"IT"`
						Ja string `json:"JA"`
						Nl string `json:"NL"`
						Pt string `json:"PT"`
						Ro string `json:"RO"`
						Ru string `json:"RU"`
						Zh string `json:"ZH"`
					} `json:"firstName"`
					LastName struct {
						Az string `json:"AZ"`
						Da string `json:"DA"`
						De string `json:"DE"`
						En string `json:"EN"`
						Es string `json:"ES"`
						Fr string `json:"FR"`
						Hu string `json:"HU"`
						It string `json:"IT"`
						Ja string `json:"JA"`
						Nl string `json:"NL"`
						Pt string `json:"PT"`
						Ro string `json:"RO"`
						Ru string `json:"RU"`
						Zh string `json:"ZH"`
					} `json:"lastName"`
					Name struct {
						Az string `json:"AZ"`
						Da string `json:"DA"`
						De string `json:"DE"`
						En string `json:"EN"`
						Es string `json:"ES"`
						Fr string `json:"FR"`
						Hu string `json:"HU"`
						It string `json:"IT"`
						Ja string `json:"JA"`
						Nl string `json:"NL"`
						Pt string `json:"PT"`
						Ro string `json:"RO"`
						Ru string `json:"RU"`
						Zh string `json:"ZH"`
					} `json:"name"`
					ShortName struct {
						Az string `json:"AZ"`
						Da string `json:"DA"`
						De string `json:"DE"`
						En string `json:"EN"`
						Es string `json:"ES"`
						Fr string `json:"FR"`
						Hu string `json:"HU"`
						It string `json:"IT"`
						Ja string `json:"JA"`
						Nl string `json:"NL"`
						Pt string `json:"PT"`
						Ro string `json:"RO"`
						Ru string `json:"RU"`
						Zh string `json:"ZH"`
					} `json:"shortName"`
				} `json:"translations"`
				Weight int64 `json:"weight"`
			} `json:"player"`
			TeamID         string `json:"teamId"`
			TeamIDProvider string `json:"teamIdProvider"`
			Time           struct {
				Minute int64 `json:"minute"`
				Second int64 `json:"second"`
			} `json:"time"`
			TotalScore struct {
				Away int64 `json:"away"`
				Home int64 `json:"home"`
			} `json:"totalScore"`
		} `json:"scorers"`
	} `json:"playerEvents"`
	Referees []struct {
		Images struct {
			SmallSquare string `json:"SMALL_SQUARE"`
		} `json:"images"`
		Person struct {
			CountryCode  string `json:"countryCode"`
			ID           string `json:"id"`
			Translations struct {
				CountryName struct {
					En string `json:"EN"`
					Fr string `json:"FR"`
					De string `json:"DE"`
					Es string `json:"ES"`
					Pt string `json:"PT"`
					It string `json:"IT"`
					Ru string `json:"RU"`
					Ja string `json:"JA"`
					Zh string `json:"ZH"`
					Hu string `json:"HU"`
					Ro string `json:"RO"`
					Da string `json:"DA"`
					Nl string `json:"NL"`
					Az string `json:"AZ"`
				} `json:"countryName"`
				FirstName struct {
					En string `json:"EN"`
					Fr string `json:"FR"`
					De string `json:"DE"`
					Es string `json:"ES"`
					Pt string `json:"PT"`
					It string `json:"IT"`
					Ru string `json:"RU"`
					Ja string `json:"JA"`
					Zh string `json:"ZH"`
					Hu string `json:"HU"`
					Ro string `json:"RO"`
					Da string `json:"DA"`
					Nl string `json:"NL"`
					Az string `json:"AZ"`
				} `json:"firstName"`
				LastName struct {
					En string `json:"EN"`
					Fr string `json:"FR"`
					De string `json:"DE"`
					Es string `json:"ES"`
					Pt string `json:"PT"`
					It string `json:"IT"`
					Ru string `json:"RU"`
					Ja string `json:"JA"`
					Zh string `json:"ZH"`
					Hu string `json:"HU"`
					Ro string `json:"RO"`
					Da string `json:"DA"`
					Nl string `json:"NL"`
					Az string `json:"AZ"`
				} `json:"lastName"`
				Name struct {
					En string `json:"EN"`
					Fr string `json:"FR"`
					De string `json:"DE"`
					Es string `json:"ES"`
					Pt string `json:"PT"`
					It string `json:"IT"`
					Ru string `json:"RU"`
					Ja string `json:"JA"`
					Zh string `json:"ZH"`
					Hu string `json:"HU"`
					Ro string `json:"RO"`
					Da string `json:"DA"`
					Nl string `json:"NL"`
					Az string `json:"AZ"`
				} `json:"name"`
				ShortName struct {
					En string `json:"EN"`
					Fr string `json:"FR"`
					De string `json:"DE"`
					Es string `json:"ES"`
					Pt string `json:"PT"`
					It string `json:"IT"`
					Ru string `json:"RU"`
					Ja string `json:"JA"`
					Zh string `json:"ZH"`
					Hu string `json:"HU"`
					Ro string `json:"RO"`
					Da string `json:"DA"`
					Nl string `json:"NL"`
					Az string `json:"AZ"`
				} `json:"shortName"`
			} `json:"translations"`
		} `json:"person"`
		Role         string `json:"role"`
		Translations struct {
			RoleName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"roleName"`
		} `json:"translations"`
	} `json:"referees"`
	Round struct {
		Active        bool   `json:"active"`
		CompetitionID string `json:"competitionId"`
		DateFrom      string `json:"dateFrom"`
		DateTo        string `json:"dateTo"`
		GroupCount    int    `json:"groupCount"`
		ID            string `json:"id"`
		MetaData      struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"metaData"`
		Mode               string `json:"mode"`
		ModeDetail         string `json:"modeDetail"`
		OrderInCompetition int    `json:"orderInCompetition"`
		Phase              string `json:"phase"`
		SeasonYear         string `json:"seasonYear"`
		SecondaryType      string `json:"secondaryType"`
		StadiumNameType    string `json:"stadiumNameType"`
		Status             string `json:"status"`
		TeamCount          int    `json:"teamCount"`
		Translations       struct {
			Abbreviation struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"abbreviation"`
			Name struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"name"`
			ShortName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"shortName"`
		} `json:"translations"`
	} `json:"round"`
	SeasonYear    string `json:"seasonYear"`
	SessionNumber int    `json:"sessionNumber"`
	Stadium       struct {
		Address  string `json:"address"`
		Capacity int    `json:"capacity"`
		City     struct {
			CountryCode  string `json:"countryCode"`
			ID           string `json:"id"`
			Translations struct {
				Name struct {
					En string `json:"EN"`
					Fr string `json:"FR"`
					De string `json:"DE"`
					Es string `json:"ES"`
					Pt string `json:"PT"`
					It string `json:"IT"`
					Ru string `json:"RU"`
					Ja string `json:"JA"`
					Zh string `json:"ZH"`
					Hu string `json:"HU"`
					Ro string `json:"RO"`
					Da string `json:"DA"`
					Nl string `json:"NL"`
					Az string `json:"AZ"`
				} `json:"name"`
			} `json:"translations"`
		} `json:"city"`
		CountryCode string `json:"countryCode"`
		Geolocation struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"geolocation"`
		ID     string `json:"id"`
		Images struct {
			MediumWide     string `json:"MEDIUM_WIDE"`
			LargeUltraWide string `json:"LARGE_ULTRA_WIDE"`
		} `json:"images"`
		OpeningDate  string `json:"openingDate"`
		Translations struct {
			MediaName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"mediaName"`
			Name struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"name"`
			OfficialName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"officialName"`
			SpecialEventsName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"specialEventsName"`
			SponsorName struct {
				En string `json:"EN"`
				Fr string `json:"FR"`
				De string `json:"DE"`
				Es string `json:"ES"`
				Pt string `json:"PT"`
				It string `json:"IT"`
				Ru string `json:"RU"`
				Ja string `json:"JA"`
				Zh string `json:"ZH"`
				Hu string `json:"HU"`
				Ro string `json:"RO"`
				Da string `json:"DA"`
				Nl string `json:"NL"`
				Az string `json:"AZ"`
			} `json:"sponsorName"`
		} `json:"translations"`
	} `json:"stadium"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

func fetchMatches() ([]Match, error) {
	params := url.Values{}
	params.Add("fromDate", "2021-06-09")
	params.Add("toDate", "2021-06-10")
	params.Add("offset", "0")
	params.Add("limit", "100")
	resp, err := http.Get("https://match.uefa.com/v2/matches?" + params.Encode())

	matches := []Match{}
	if err != nil {
		return matches, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&matches)
	if err != nil {
		return matches, err
	}

	return matches, err
}

type EventType int

const (
	START = iota
	END   = iota
	GOAL  = iota
)

type MatchEvent struct {
	Event EventType
	Label string
}

func compareNewInfos(current, previous []Match) []MatchEvent {
	return []MatchEvent{
		{START, "Match start"},
	}
}
