package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yulio94/star-wars-cli/internal"
	"log"
	"strconv"
)

func InitStarshipsCmd(starshipsRepo internal.StarshipsRepo) (peopleCmd *cobra.Command) {
	peopleCmd = &cobra.Command{
		Use:   "starships",
		Short: "Star Wars starships",
		Run:   runStarshipsFn(starshipsRepo),
	}

	peopleCmd.Flags().StringP(id, idShort, "", idUsage)
	peopleCmd.Flags().BoolP(file, fileShort, false, fileUsage)
	peopleCmd.Flags().StringP(fileName, fileNameShort, defaultFileName, fileNameUsage)
	return
}

func runStarshipsFn(starshipsRepo internal.StarshipsRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		fileName := ""

		givenId, _ := cmd.Flags().GetString(id)
		if givenId != "" {
			id, _ := strconv.Atoi(givenId)
			values, err := starshipsRepo.GetStarship(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(values)
		} else {
			values, err := starshipsRepo.GetStarships()
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
