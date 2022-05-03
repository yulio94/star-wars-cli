package main

import (
	"github.com/spf13/cobra"
	"github.com/yulio94/star-wars-cli/internal/cli"
	"github.com/yulio94/star-wars-cli/internal/storage/swapi"
	"log"
)

func main() {
	var peopleRepo = swapi.NewSwapiPeopleRepo()
	var starshipsRepo = swapi.NewSwapiStarshipsRepo()
	var planetsRepo = swapi.NewSwapiPlanetsRepo()

	rootCmd := &cobra.Command{Use: "starwars"}
	rootCmd.AddCommand(cli.InitPeopleCmd(peopleRepo))
	rootCmd.AddCommand(cli.InitStarshipsCmd(starshipsRepo))
	rootCmd.AddCommand(cli.InitPlanetsCmd(planetsRepo))
	log.Fatal(rootCmd.Execute())
}
