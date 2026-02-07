package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	TeamName    string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, pointsTeam1 int, team2 string, pointsTeam2 int) {
	if pointsTeam1 > pointsTeam2 {
		l.Wins[team1] += 1
	} else if pointsTeam1 < pointsTeam2 {
		l.Wins[team2] += 1
	}
}

func (l League) Ranking() []string {
	teamNames := make([]string, len(l.Teams))
	for i, team := range l.Teams {
		teamNames[i] = team.TeamName
	}
	// teamNamesをl.Wins[teamName]の降順でソート
	sort.Slice(teamNames, func(i, j int) bool {
		return l.Wins[teamNames[i]] > l.Wins[teamNames[j]]
	})
	return teamNames
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, s := range r.Ranking() {
		io.WriteString(w, s+"\n")
	}
}

func main() {
	// 5つのサンプルTeamからなる[]Teamsインスタンスを用意
	teams := []Team{
		{TeamName: "Tigers", PlayerNames: []string{"Alice", "Bob"}},
		{TeamName: "Lions", PlayerNames: []string{"Carol", "Dave"}},
		{TeamName: "Eagles", PlayerNames: []string{"Eve", "Frank"}},
		{TeamName: "Hawks", PlayerNames: []string{"Grace", "Heidi"}},
		{TeamName: "Wolves", PlayerNames: []string{"Ivan", "Judy"}},
	}

	// Leagueインスタンスを用意
	league := League{
		Teams: teams,
		Wins:  make(map[string]int),
	}

	// 4チームについてだけMatchResultで勝敗を記録、あえて1チームは記録を残さない
	league.MatchResult("Tigers", 3, "Lions", 1)  // Tigers win
	league.MatchResult("Tigers", 3, "Lions", 1)  // Tigers win
	league.MatchResult("Tigers", 3, "Lions", 1)  // Tigers win
	league.MatchResult("Tigers", 3, "Lions", 1)  // Tigers win
	league.MatchResult("Eagles", 2, "Hawks", 4)  // Hawks win
	league.MatchResult("Eagles", 2, "Hawks", 4)  // Hawks win
	league.MatchResult("Eagles", 2, "Hawks", 4)  // Hawks win
	league.MatchResult("Tigers", 2, "Eagles", 3) // Eagles win
	league.MatchResult("Tigers", 2, "Eagles", 3) // Eagles win
	league.MatchResult("Lions", 5, "Hawks", 3)   // Lions win
	// "Wolves" だけMatchResultに出さない

	RankPrinter(league, os.Stdout)
}
