package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yulio94/star-wars-cli/internal"
	"log"
	"strconv"
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

		givenId, _ := cmd.Flags().GetString(id)
		if givenId != "" {
			id, _ := strconv.Atoi(givenId)
			values, err := repository.GetPerson(id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(values)
		} else {
			values, err := repository.GetPeople()
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
