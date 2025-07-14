package player

import "slices"

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func calculateRating(goals, misses, assists int) float64 {
	var r float64
	switch {
	case misses == 0:
		r = float64(goals) + float64(assists)/2.0
	default:
		r = (float64(goals) + float64(assists)/2.0) / float64(misses)
	}
	return r
}

func NewPlayer(name string, goals, misses, assists int) Player {
	r := calculateRating(goals, misses, assists)
	return Player{Name: name, Goals: goals, Misses: misses, Assists: assists, Rating: r}
}

func SortGoals(players []Player) []Player {
	return goalsSort(players)
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, byGoalsDesc)
	return players
}

func byGoalsDesc(a, e Player) int {
	switch {
	case a.Goals < e.Goals:
		return 1
	case a.Goals > e.Goals:
		return -1
	default:
		return 0
	}
}
