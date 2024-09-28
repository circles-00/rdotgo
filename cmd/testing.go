package cmd

import (
	"rdotgo/cmd/utils"

	"github.com/spf13/cobra"
)

// testingCmd represents the testing command
var testingCmd = &cobra.Command{
	Use:   "testing",
	Short: "Just a testing command for development",
	Long:  `Just a testing command for development`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.SetupUbuntuMachine()
	},
}

func init() {
	rootCmd.AddCommand(testingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
