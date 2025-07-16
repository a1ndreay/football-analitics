package player

import (
	"fmt"
	"slices"
	"strings"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p Player) String() string {
	return fmt.Sprintf("Name - %s, Goals - %d, Missles - %d, Assists - %d, Rating - %.1f", p.Name, p.Goals, p.Misses, p.Assists, p.Rating)
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

func SortRating(players []Player) []Player {
	return ratingSort(players)
}

func SortGoalsAndMissles(players []Player) []Player {
	return gmSort(players)
}

func goalsSort(players []Player) []Player {
	slices.SortFunc(players, byGoalsDesc)
	return players
}

func ratingSort(players []Player) []Player {
	slices.SortFunc(players, byRatingDesc)
	return players
}

func gmSort(players []Player) []Player {
	slices.SortFunc(players, byGoalsAndMisslesDesc)
	return players
}

func byGoalsDesc(a, e Player) int {
	switch {
	case a.Goals < e.Goals:
		return 1
	case a.Goals > e.Goals:
		return -1
	default:
		return byName(a, e)
	}
}

func byRatingDesc(a, e Player) int {
	switch {
	case a.Rating < e.Rating:
		return 1
	case a.Rating > e.Rating:
		return -1
	default:
		return byName(a, e)
	}
}

func byGoalsAndMisslesDesc(a, e Player) int {
	if a.Misses != 0 && e.Misses != 0 {
		switch {
		case float64(a.Goals)/float64(a.Misses) < float64(e.Goals)/float64(e.Misses):
			return 1
		case float64(a.Goals)/float64(a.Misses) > float64(e.Goals)/float64(e.Misses):
			return -1
		default:
			return byName(a, e)
		}
	} else if e.Misses == 0 {
		switch {
		case float64(a.Goals)/float64(a.Misses) < float64(e.Goals):
			return 1
		case float64(a.Goals)/float64(a.Misses) > float64(e.Goals):
			return -1
		default:
			return byName(a, e)
		}
	} else if a.Misses == 0 {
		switch {
		case float64(a.Goals) < float64(e.Goals)/float64(e.Misses):
			return 1
		case float64(a.Goals) > float64(e.Goals)/float64(e.Misses):
			return -1
		default:
			return byName(a, e)
		}
	} else {
		switch {
		case float64(a.Goals) < float64(e.Goals):
			return 1
		case float64(a.Goals) > float64(e.Goals):
			return -1
		default:
			return byName(a, e)
		}
	}
}

func byName(a, e Player) int {
	switch {
	case strings.Compare(a.Name, e.Name) < 0:
		return -1
	case strings.Compare(a.Name, e.Name) > 0:
		return 1
	default:
		return 0
	}
}
