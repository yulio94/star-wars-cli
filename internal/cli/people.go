package cli

import (
	"github.com/spf13/cobra"
	"github.com/yulio94/star-wars-cli/internal"
	"log"
)

func InitPeopleCmd(repository internal.PeopleRepo) (peopleCmd *cobra.Command) {
	peopleCmd = &cobra.Command{
		Use:   "people",
		Short: "Star Wars people",
		Run:   runPeopleFn(repository),
	}

	peopleCmd.Flags().StringP(id, idShort, "", idUsage)
	peopleCmd.Flags().BoolP(file, fileShort, false, fileUsage)
	peopleCmd.Flags().StringP(fileName, fileNameShort, defaultFileName, fileNameUsage)
	return
}

func runPeopleFn(repository internal.PeopleRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fileName := ""

		givenFileName, _ := cmd.Flags().GetString("filename")
		if givenFileName != "" {
			fileName = givenFileName
		}

		log.Fatal(repository.SaveFile(fileName))
	}
}
