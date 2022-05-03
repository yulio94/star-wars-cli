package cli

import (
	"github.com/spf13/cobra"
	"github.com/yulio94/star-wars-cli/internal"
	"log"
	"strconv"
)

func InitStarshipsCmd(starshipsRepo internal.StarshipsRepo, storageRepo storage.Repo) (peopleCmd *cobra.Command) {
	peopleCmd = &cobra.Command{
		Use:   "starships",
		Short: "Star Wars starships",
		Run:   runStarshipsFn(starshipsRepo, storageRepo),
	}

	peopleCmd.Flags().StringP(id, idShort, "", idUsage)
	peopleCmd.Flags().BoolP(file, fileShort, false, fileUsage)
	peopleCmd.Flags().StringP(fileName, fileNameShort, defaultFileName, fileNameUsage)
	return
}

func runStarshipsFn(starshipsRepo internal.StarshipsRepo, storageRepo storage.Repo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		name := ""

		givenId, _ := cmd.Flags().GetString(id)
		if givenId != "" {
			id, _ := strconv.Atoi(givenId)
			values, err := starshipsRepo.GetStarship(id)
			log.Fatal(storageRepo.Save(values, name))
		} else {
			values, err := starshipsRepo.GetStarships()
			log.Fatal(storageRepo.Save(values, name))
		}

		givenFileName, _ := cmd.Flags().GetString("fileName")
		if givenFileName != "" {
			name = givenFileName
		}
	}
}
