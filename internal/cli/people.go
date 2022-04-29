package cli

import (
	"github.com/spf13/cobra"
	"log"
)

type CobraFn func(cmd *cobra.Command, args []string)

func InitPeopleCmd(repository StarWarsRepo) (peopleCmd *cobra.Command) {
	peopleCmd = &cobra.Command{
		Use:   "people",
		Short: "Star Wars commands",
		Run:   runStarWarsFn(repository),
	}

	peopleCmd.Flags().StringP("file-name", "f", "", "File name")
	return
}

func runPeopleFn(repository StarWarsRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fileName := ""

		givenFileName, _ := cmd.Flags().GetString("filename")
		if givenFileName != "" {
			fileName = givenFileName
		}

		log.Fatal(repository.SaveFile(fileName))
	}
}
