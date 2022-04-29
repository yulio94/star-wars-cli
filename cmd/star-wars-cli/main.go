package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{Use: "starwars"}
	rootCmd.AddCommand()
	log.Fatal(rootCmd.Execute())
}
