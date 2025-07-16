package test

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/a1ndreay/football-analitics/internal/player"
)

type NewPlayerArgs struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
}

func formatError(testFunc interface{}, give, got, want any) string {
	return fmt.Sprintf(
		"\nerror at %s:\n\tassert\t%v,\n\tgot\t%v,\n\texpected\t%v",
		runtime.FuncForPC(reflect.ValueOf(testFunc).Pointer()).Name(),
		give,
		got,
		want,
	)
}

func TestNewPlayer(t *testing.T) {
	tests := []struct {
		give NewPlayerArgs
		want *player.Player
	}{
		{
			NewPlayerArgs{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5},
			&player.Player{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
		},
		{
			NewPlayerArgs{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1},
			&player.Player{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
		},
	}
	for _, tc := range tests {
		got, _ := player.NewPlayer(tc.give.Name, tc.give.Goals, tc.give.Misses, tc.give.Assists)
		if got != tc.want {
			t.Fatalf("%s", formatError(player.NewPlayer, tc.give, got, tc.want))
		}
	}
}

type SortTest struct {
	give []player.Player
	want []player.Player
}

var (
	SortTestsAssert_TestSortGoals = []SortTest{
		{
			give: []player.Player{
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75},
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75},
				{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.25},
			},
			want: []player.Player{
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.25},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75},
			},
		},
		{
			give: []player.Player{
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.75},
			},
			want: []player.Player{
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.75},
			},
		},
	}
	SortTestsAssert_TestSortRating = []SortTest{
		{
			give: []player.Player{
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75},
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75},
				{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.25},
			},
			want: []player.Player{
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.25},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75},
			},
		},
		{
			give: []player.Player{
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75},
				{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 2.25},
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75},
				{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.25},
			},
			want: []player.Player{
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75},
				{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 2.25},
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.25},
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75},
			},
		},
	}
	SortTestsAssert_TestSortGoalsAndMissles = []SortTest{
		{
			give: []player.Player{
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},   //2
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},  //1.5
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75}, //0.5
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75}, //2.5
				{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.25},    //1
			},
			want: []player.Player{
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75}, //2.5
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},   //2
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},  //1.5
				{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.25},    //1
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75}, //0.5
			},
		},
		{
			give: []player.Player{
				{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 2.25},  //2
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},   //2
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},  //1.5
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75}, //0.5
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75}, //2.5
				{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.25},    //0
			},
			want: []player.Player{
				{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 2.75}, //2.5
				{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},   //2
				{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 2.25},  //2
				{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},  //1.5
				{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 0.75}, //0.5
				{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.25},    //0
			},
		},
	}
)

func TestSortGoals(t *testing.T) {
	for _, tc := range SortTestsAssert_TestSortGoals {
		got := player.SortGoals(tc.give)
		for i, gv := range got {
			if tc.want[i] != gv {
				t.Fatalf("%s", formatError(player.SortGoals, tc.give, gv, tc.want[i]))
			}
		}
	}
}

func TestSortRating(t *testing.T) {
	for _, tc := range SortTestsAssert_TestSortRating {
		got := player.SortRating(tc.give)
		for i, gv := range got {
			if tc.want[i] != gv {
				t.Fatalf("%s", formatError(player.SortGoals, tc.give, gv, tc.want[i]))
			}
		}
	}
}

func TestSortGoalsAndMissles(t *testing.T) {
	for _, tc := range SortTestsAssert_TestSortGoalsAndMissles {
		got := player.SortGoalsAndMissles(tc.give)
		for i, gv := range got {
			if tc.want[i] != gv {
				t.Fatalf("%s", formatError(player.SortGoals, tc.give, gv, tc.want[i]))
			}
		}
	}
}
