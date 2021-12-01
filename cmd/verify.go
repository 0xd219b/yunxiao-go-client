/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/Berger7/yunxiao-go-client/pkg/verify"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify your configuration",
	Long: `Verify can check your configuration is valid or not.

We store your configuration in a file named .yunxiao.yaml 
in the current user's home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting verify yunxiao client necessary configuration...")
		verify.RunVerify()
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
