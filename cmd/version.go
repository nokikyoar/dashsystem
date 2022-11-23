/*
Copyright © 2022 kikyoar
*/
package cmd

import (
	"fmt"
	"yuhao.com/dashSystem/config"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本信息",
	Long:  `目前版本属于Beta版本，暂未广泛使用.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dashsystem version: ", config.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
