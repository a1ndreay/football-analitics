package cmd

import (
	"fmt"
	"strconv"

	"github.com/a1ndreay/football-analitics/internal/player"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newPlayerCmd)
}

var newPlayerCmd = &cobra.Command{
	Use:   "add",
	Short: "add new player",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		g, _ := strconv.Atoi(args[1])
		m, _ := strconv.Atoi(args[2])
		a, _ := strconv.Atoi(args[3])
		fmt.Println(player.NewPlayer(args[0], g, m, a))
	},
}
