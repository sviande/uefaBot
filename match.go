package main

import "time"

//Match struct
type Match []struct {
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
