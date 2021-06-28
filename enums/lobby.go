package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type Lobby string

const (
	AnyLobby   Lobby = "any"
	leaguePair Lobby = "league_pair"
	leagueTeam Lobby = "league_team"
	ranked     Lobby = "gachi"
	private    Lobby = "private"
	regular    Lobby = "regular"
	fesSolo    Lobby = "fes_solo"
	fesTeam    Lobby = "fes_team"
)

func (l *Lobby) UnmarshalJSON(b []byte) error {
	type L Lobby
	r := (*L)(l)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *l {
	case AnyLobby, leagueTeam, leaguePair, ranked, private, regular, fesSolo, fesTeam:
		return nil
	}
	return errors.New("Invalid enums.Lobby. Got: " + fmt.Sprint(*l))
}

func GetLobby() []Lobby {
	return []Lobby{
		AnyLobby, leaguePair, leagueTeam, ranked, private, regular, fesSolo, fesTeam,
	}
}

func (l Lobby) GetDisplay(printer *message.Printer) string {
	switch l {
	case AnyLobby:
		return printer.Sprintf("Any Lobby")
	case leagueTeam:
		return printer.Sprintf("League Battle (Team)")
	case leaguePair:
		return printer.Sprintf("League Battle (Pair)")
	case ranked:
		return printer.Sprintf("Ranked Battle")
	case private:
		return printer.Sprintf("Private Battle")
	case regular:
		return printer.Sprintf("Normal Battle")
	case fesSolo:
		return printer.Sprintf("Splatfest Solo/Pro")
	case fesTeam:
		return printer.Sprintf("Splatfest Team/Normal")
	}
	return ""
}
