package main

import (
	"github.com/spf13/cobra"

	"github.com/PI-Victor/groundhog/pkg/localsys"
)

var (
	localPath string
)

var runCmd = &cobra.Command{
	Use:   "path",
	Short: "Rename your picture",
	Run: func(cmd *cobra.Command, args []string) {
		localsys.Run(localPath)
	},
}

func main() {
	runCmd.Execute()
}

func init() {
	runCmd.Flags().StringVar(&localPath, "path", "", "Specify the path to be searched")
}
