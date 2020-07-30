package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

type Options struct {
	// The cli app's name.
	// Required: true
	name string

	// The app's description.
	description string
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}
