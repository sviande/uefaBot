package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type MatchInfo struct {
	AwayTeam          AwayTeam     `json:"awayTeam"`
	BehindClosedDoors bool         `json:"behindClosedDoors"`
	Competition       Competition  `json:"competition"`
	Condition         Condition    `json:"condition"`
	FullTimeAt        time.Time    `json:"fullTimeAt"`
	Group             Group        `json:"group"`
	HomeTeam          HomeTeam     `json:"homeTeam"`
	ID                string       `json:"id"`
	KickOffTime       KickOffTime  `json:"kickOffTime"`
	LineupStatus      string       `json:"lineupStatus"`
	MatchAttendance   int          `json:"matchAttendance"`
	Matchday          Matchday     `json:"matchday"`
	PlayerEvents      PlayerEvents `json:"playerEvents,omitEmpty"`
	Referees          []Referees   `json:"referees"`
	Round             Round        `json:"round"`
	Score             Score        `json:"score,omitEmpty"`
	SeasonYear        string       `json:"seasonYear"`
	SessionNumber     int          `json:"sessionNumber"`
	Stadium           Stadium      `json:"stadium"`
	Status            string       `json:"status"`
	Type              string       `json:"type"`
	Winner            Winner       `json:"winner"`
}
type CountryName struct {
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
}
type DisplayName struct {
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
}
type DisplayOfficialName struct {
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
}
type DisplayTeamCode struct {
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
}
type AwayTeamTranslations struct {
	CountryName         CountryName         `json:"countryName"`
	DisplayName         DisplayName         `json:"displayName"`
	DisplayOfficialName DisplayOfficialName `json:"displayOfficialName"`
	DisplayTeamCode     DisplayTeamCode     `json:"displayTeamCode"`
}
type AwayTeam struct {
	AssociationLogoURL string               `json:"associationLogoUrl"`
	BigLogoURL         string               `json:"bigLogoUrl"`
	CountryCode        string               `json:"countryCode"`
	ID                 string               `json:"id"`
	IDProvider         string               `json:"idProvider"`
	InternationalName  string               `json:"internationalName"`
	IsPlaceHolder      bool                 `json:"isPlaceHolder"`
	LogoURL            string               `json:"logoUrl"`
	MediumLogoURL      string               `json:"mediumLogoUrl"`
	TeamCode           string               `json:"teamCode"`
	TeamTypeDetail     string               `json:"teamTypeDetail"`
	Translations       AwayTeamTranslations `json:"translations"`
	TypeIsNational     bool                 `json:"typeIsNational"`
	TypeTeam           string               `json:"typeTeam"`
}
type Images struct {
	FullLogo string `json:"FULL_LOGO"`
}
type MetaData struct {
	Name string `json:"name"`
}
type Name struct {
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
}
type QualifyingName struct {
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
}
type TournamentName struct {
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
}
type CompetionTranslations struct {
	Name           Name           `json:"name"`
	QualifyingName QualifyingName `json:"qualifyingName"`
	TournamentName TournamentName `json:"tournamentName"`
}
type Competition struct {
	Age          string                `json:"age"`
	Code         string                `json:"code"`
	ID           string                `json:"id"`
	Images       Images                `json:"images"`
	MetaData     MetaData              `json:"metaData"`
	Region       string                `json:"region"`
	Sex          string                `json:"sex"`
	SportsType   string                `json:"sportsType"`
	TeamCategory string                `json:"teamCategory"`
	Translations CompetionTranslations `json:"translations"`
	Type         string                `json:"type"`
}
type PitchConditionName struct {
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
}
type WeatherConditionName struct {
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
}
type ConditionTranslations struct {
	PitchConditionName   PitchConditionName   `json:"pitchConditionName"`
	WeatherConditionName WeatherConditionName `json:"weatherConditionName"`
}
type Condition struct {
	PitchCondition   string                `json:"pitchCondition"`
	Temperature      int                   `json:"temperature"`
	Translations     ConditionTranslations `json:"translations"`
	WeatherCondition string                `json:"weatherCondition"`
	WindSpeed        int                   `json:"windSpeed"`
}
type GroupMetaData struct {
	GroupName      string `json:"groupName"`
	GroupShortName string `json:"groupShortName"`
}
type ShortName struct {
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
}
type GroupTranslations struct {
	Name      Name      `json:"name"`
	ShortName ShortName `json:"shortName"`
}
type Group struct {
	CompetitionID string            `json:"competitionId"`
	ID            string            `json:"id"`
	MetaData      GroupMetaData     `json:"metaData"`
	Order         int               `json:"order"`
	Phase         string            `json:"phase"`
	RoundID       string            `json:"roundId"`
	SeasonYear    string            `json:"seasonYear"`
	Translations  GroupTranslations `json:"translations"`
	Type          string            `json:"type"`
}
type HomeTeamTranslations struct {
	CountryName         CountryName         `json:"countryName"`
	DisplayName         DisplayName         `json:"displayName"`
	DisplayOfficialName DisplayOfficialName `json:"displayOfficialName"`
	DisplayTeamCode     DisplayTeamCode     `json:"displayTeamCode"`
}
type HomeTeam struct {
	AssociationLogoURL string               `json:"associationLogoUrl"`
	BigLogoURL         string               `json:"bigLogoUrl"`
	CountryCode        string               `json:"countryCode"`
	ID                 string               `json:"id"`
	IDProvider         string               `json:"idProvider"`
	InternationalName  string               `json:"internationalName"`
	IsPlaceHolder      bool                 `json:"isPlaceHolder"`
	LogoURL            string               `json:"logoUrl"`
	MediumLogoURL      string               `json:"mediumLogoUrl"`
	TeamCode           string               `json:"teamCode"`
	TeamTypeDetail     string               `json:"teamTypeDetail"`
	Translations       HomeTeamTranslations `json:"translations"`
	TypeIsNational     bool                 `json:"typeIsNational"`
	TypeTeam           string               `json:"typeTeam"`
}
type KickOffTime struct {
	Date             string    `json:"date"`
	DateTime         time.Time `json:"dateTime"`
	UtcOffsetInHours int       `json:"utcOffsetInHours"`
}
type LongName struct {
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
}
type MatchdayTranslations struct {
	LongName LongName `json:"longName"`
	Name     Name     `json:"name"`
}
type Matchday struct {
	CompetitionID  string               `json:"competitionId"`
	DateFrom       string               `json:"dateFrom"`
	DateTo         string               `json:"dateTo"`
	ID             string               `json:"id"`
	LongName       string               `json:"longName"`
	Name           string               `json:"name"`
	RoundID        string               `json:"roundId"`
	SeasonYear     string               `json:"seasonYear"`
	SequenceNumber string               `json:"sequenceNumber"`
	Translations   MatchdayTranslations `json:"translations"`
	Type           string               `json:"type"`
}
type ScorersImages struct {
	PlayerCelebrating string `json:"PLAYER_CELEBRATING"`
}
type FieldPosition struct {
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
}
type FirstName struct {
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
}
type LastName struct {
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
}
type PlayerTranslations struct {
	CountryName   CountryName   `json:"countryName"`
	FieldPosition FieldPosition `json:"fieldPosition"`
	FirstName     FirstName     `json:"firstName"`
	LastName      LastName      `json:"lastName"`
	Name          Name          `json:"name"`
	ShortName     ShortName     `json:"shortName"`
}
type Player struct {
	Age                   string             `json:"age"`
	BirthDate             string             `json:"birthDate"`
	CountryCode           string             `json:"countryCode"`
	DetailedFieldPosition string             `json:"detailedFieldPosition"`
	FieldPosition         string             `json:"fieldPosition"`
	Height                int                `json:"height"`
	ID                    string             `json:"id"`
	ImageURL              string             `json:"imageUrl"`
	InternationalName     string             `json:"internationalName"`
	Translations          PlayerTranslations `json:"translations"`
	Weight                int                `json:"weight"`
}
type Time struct {
	Minute int `json:"minute"`
	Second int `json:"second"`
}
type TotalScore struct {
	Away int `json:"away"`
	Home int `json:"home"`
}
type Scorers struct {
	GoalType       string        `json:"goalType"`
	ID             string        `json:"id"`
	Images         ScorersImages `json:"images"`
	Player         Player        `json:"player"`
	TeamID         string        `json:"teamId"`
	TeamIDProvider string        `json:"teamIdProvider"`
	Time           Time          `json:"time"`
	TotalScore     TotalScore    `json:"totalScore"`
}
type PlayerEvents struct {
	Scorers []Scorers `json:"scorers"`
}
type RefereesImages struct {
	SmallSquare string `json:"SMALL_SQUARE"`
}
type Translations struct {
	CountryName CountryName `json:"countryName"`
	FirstName   FirstName   `json:"firstName"`
	LastName    LastName    `json:"lastName"`
	Name        Name        `json:"name"`
	ShortName   ShortName   `json:"shortName"`
}
type Person struct {
	CountryCode  string       `json:"countryCode"`
	ID           string       `json:"id"`
	Translations Translations `json:"translations"`
}
type RoleName struct {
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
}
type RefereesTranslations struct {
	RoleName RoleName `json:"roleName"`
}
type Referees struct {
	Images       RefereesImages       `json:"images"`
	Person       Person               `json:"person"`
	Role         string               `json:"role"`
	Translations RefereesTranslations `json:"translations"`
}
type RoundMetaData struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type Abbreviation struct {
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
}
type RoundTranslations struct {
	Abbreviation Abbreviation `json:"abbreviation"`
	Name         Name         `json:"name"`
	ShortName    ShortName    `json:"shortName"`
}
type Round struct {
	Active             bool              `json:"active"`
	CompetitionID      string            `json:"competitionId"`
	DateFrom           string            `json:"dateFrom"`
	DateTo             string            `json:"dateTo"`
	GroupCount         int               `json:"groupCount"`
	ID                 string            `json:"id"`
	MetaData           RoundMetaData     `json:"metaData"`
	Mode               string            `json:"mode"`
	ModeDetail         string            `json:"modeDetail"`
	OrderInCompetition int               `json:"orderInCompetition"`
	Phase              string            `json:"phase"`
	SeasonYear         string            `json:"seasonYear"`
	StadiumNameType    string            `json:"stadiumNameType"`
	Status             string            `json:"status"`
	TeamCount          int               `json:"teamCount"`
	Translations       RoundTranslations `json:"translations"`
}
type Regular struct {
	Away int `json:"away"`
	Home int `json:"home"`
}
type Total struct {
	Away int `json:"away"`
	Home int `json:"home"`
}
type Score struct {
	Regular Regular `json:"regular"`
	Total   Total   `json:"total"`
}
type CityTranslations struct {
	Name Name `json:"name"`
}
type City struct {
	CountryCode  string           `json:"countryCode"`
	ID           string           `json:"id"`
	Translations CityTranslations `json:"translations"`
}
type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type StadiumImages struct {
	MediumWide     string `json:"MEDIUM_WIDE"`
	LargeUltraWide string `json:"LARGE_ULTRA_WIDE"`
}
type MediaName struct {
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
}
type OfficialName struct {
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
}
type SpecialEventsName struct {
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
}
type SponsorName struct {
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
}
type StadiumTranslations struct {
	MediaName         MediaName         `json:"mediaName"`
	Name              Name              `json:"name"`
	OfficialName      OfficialName      `json:"officialName"`
	SpecialEventsName SpecialEventsName `json:"specialEventsName"`
	SponsorName       SponsorName       `json:"sponsorName"`
}
type Stadium struct {
	Address      string              `json:"address"`
	Capacity     int                 `json:"capacity"`
	City         City                `json:"city"`
	CountryCode  string              `json:"countryCode"`
	Geolocation  Geolocation         `json:"geolocation"`
	ID           string              `json:"id"`
	Images       StadiumImages       `json:"images"`
	OpeningDate  string              `json:"openingDate"`
	Translations StadiumTranslations `json:"translations"`
}
type TeamTranslations struct {
	CountryName         CountryName         `json:"countryName"`
	DisplayName         DisplayName         `json:"displayName"`
	DisplayOfficialName DisplayOfficialName `json:"displayOfficialName"`
	DisplayTeamCode     DisplayTeamCode     `json:"displayTeamCode"`
}
type Team struct {
	AssociationLogoURL string           `json:"associationLogoUrl"`
	BigLogoURL         string           `json:"bigLogoUrl"`
	CountryCode        string           `json:"countryCode"`
	ID                 string           `json:"id"`
	IDProvider         string           `json:"idProvider"`
	InternationalName  string           `json:"internationalName"`
	IsPlaceHolder      bool             `json:"isPlaceHolder"`
	LogoURL            string           `json:"logoUrl"`
	MediumLogoURL      string           `json:"mediumLogoUrl"`
	TeamCode           string           `json:"teamCode"`
	TeamTypeDetail     string           `json:"teamTypeDetail"`
	Translations       TeamTranslations `json:"translations"`
	TypeIsNational     bool             `json:"typeIsNational"`
	TypeTeam           string           `json:"typeTeam"`
}
type Match struct {
	Reason string `json:"reason"`
	Team   Team   `json:"team"`
}
type Winner struct {
	Match Match `json:"match"`
}

func fetchMatches() (map[string]MatchInfo, error) {
	params := url.Values{}
	today := time.Now()
	dateLayout := "2006-01-02"
	tomorrow := today.Add(24 * time.Hour)
	params.Add("fromDate", today.Format(dateLayout))
	params.Add("toDate", tomorrow.Format(dateLayout))
	params.Add("offset", "0")
	params.Add("limit", "100")
	resp, err := http.Get("https://match.uefa.com/v2/matches?" + params.Encode())

	mapMatches := make(map[string]MatchInfo)
	if err != nil {
		return mapMatches, err
	}

	defer resp.Body.Close()

	matches := []MatchInfo{}
	err = json.NewDecoder(resp.Body).Decode(&matches)
	if err != nil {
		return mapMatches, err
	}

	for _, match := range matches {
		mapMatches[match.ID] = match
	}

	return mapMatches, err
}

type EventType int

const (
	START = iota
	END   = iota
	GOAL  = iota
)

type MatchStatus int

const (
	MatchUpcoming = "UPCOMING"
	MatchLive     = "LIVE"
	MatchFinished = "FINISHED"
)

type MatchEvent struct {
	Event EventType
	Label string
}

func compareNewInfos(currentMap, previousMap map[string]MatchInfo) []MatchEvent {
	events := []MatchEvent{}

	for ID, currentMatch := range currentMap {
		previous, exists := previousMap[ID]
		if exists == false {
			continue
		}

		if previous.Status == MatchUpcoming && currentMatch.Status == MatchLive {
			label := fmt.Sprintf("Match %s : %s started", currentMatch.HomeTeam.InternationalName, currentMatch.AwayTeam.InternationalName)
			newEvent := MatchEvent{
				Event: START,
				Label: label,
			}
			events = append(events, newEvent)
		}

		if currentMatch.PlayerEvents.Scorers == nil {
			continue
		}

		previousScorers := 0
		if previous.PlayerEvents.Scorers != nil {
			previousScorers = len(previous.PlayerEvents.Scorers)
		}
		currentScorers := len(currentMatch.PlayerEvents.Scorers)

		if previousScorers == currentScorers {
			continue
		}

		newScorers := currentMatch.PlayerEvents.Scorers[previousScorers:currentScorers]

		for _, scorer := range newScorers {
			label := fmt.Sprintf("GOAL!!! %s %d:%d", scorer.Player.InternationalName, scorer.TotalScore.Home, scorer.TotalScore.Away)
			newEvent := MatchEvent{
				Event: GOAL,
				Label: label,
			}
			events = append(events, newEvent)
		}
	}

	return events
}
