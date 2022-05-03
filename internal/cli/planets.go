package cli

import (
	"github.com/spf13/cobra"
	"github.com/yulio94/star-wars-cli/internal"
	"log"
)

func InitPlanetsCmd(repository internal.PlanetRepo) (peopleCmd *cobra.Command) {
	peopleCmd = &cobra.Command{
		Use:   "planets",
		Short: "Star Wars planets",
		Run:   runPlanetsFn(repository),
	}
	peopleCmd.Flags().StringP(id, idShort, "", idUsage)
	peopleCmd.Flags().BoolP(file, fileShort, false, fileUsage)
	peopleCmd.Flags().StringP(fileName, fileNameShort, defaultFileName, fileNameUsage)
	return
}

func runPlanetsFn(repository internal.PlanetRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fileName := ""

		givenFileName, _ := cmd.Flags().GetString("filename")
		if givenFileName != "" {
			fileName = givenFileName
		}

		log.Fatal(repository.SaveFile(fileName))
	}
}
