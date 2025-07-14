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
		want player.Player
	}{
		{
			NewPlayerArgs{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5},
			player.Player{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
		},
		{
			NewPlayerArgs{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1},
			player.Player{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
		},
	}
	for _, tc := range tests {
		got := player.NewPlayer(tc.give.Name, tc.give.Goals, tc.give.Misses, tc.give.Assists)
		if got != tc.want {
			t.Fatalf("%s", formatError(player.NewPlayer, tc.give, got, tc.want))
		}
	}
}

func TestSortGoals(t *testing.T) {
	type Test struct {
		give []player.Player
		want []player.Player
	}
	tests := []Test{
		{
			give: []player.Player{
				player.Player{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				player.Player{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.75},
			},
			want: []player.Player{
				player.Player{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				player.Player{Name: "John", Goals: 2, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 1.75},
			},
		},
		{
			give: []player.Player{
				player.Player{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				player.Player{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.75},
			},
			want: []player.Player{
				player.Player{Name: "Ronaldo", Goals: 5, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Kirill", Goals: 4, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Zakhar", Goals: 3, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "Andrey", Goals: 2, Misses: 0, Assists: 5, Rating: 4.5},
				player.Player{Name: "Denchik", Goals: 1, Misses: 2, Assists: 1, Rating: 1.75},
				player.Player{Name: "John", Goals: 0, Misses: 2, Assists: 1, Rating: 1.75},
			},
		},
	}

	for _, tc := range tests {
		got := player.SortGoals(tc.give)
		for i, gv := range got {
			if tc.want[i] != gv {
				t.Fatalf("%s", formatError(player.SortGoals, tc.give, gv, tc.want[i]))
			}
		}
	}
}
