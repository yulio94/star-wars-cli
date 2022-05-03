package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

type CobraFn func(cmd *cobra.Command, args []string)

const csvExtension string = ".csv"

var defaultFileName = fmt.Sprintf("%s.%s", time.Now(), csvExtension)

const (
	id            = "id"
	idShort       = "i"
	idUsage       = "id to get from api."
	file          = "file"
	fileShort     = "f"
	fileUsage     = "file to read from"
	fileName      = "fileName"
	fileNameShort = "fn"
	fileNameUsage = "Name to use in the csv file."
)
