package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yulio94/star-wars-cli/internal"
	"log"
	"strconv"
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

		givenId, _ := cmd.Flags().GetString(id)
		if givenId != "" {
			id, _ := strconv.Atoi(givenId)
			values, err := repository.GetPlanet(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(values)
		} else {
			values, err := repository.GetPlanets()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(values)
		}

		givenFileName, _ := cmd.Flags().GetString("fileName")
		if givenFileName != "" {
			fileName = givenFileName
		}

		fmt.Println("File fileName: ", fileName)
	}

}
