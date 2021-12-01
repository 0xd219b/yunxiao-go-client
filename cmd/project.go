/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/Berger7/yunxiao-go-client/config"
	"github.com/Berger7/yunxiao-go-client/pkg/project"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "an easy tools collection of project",
	Long: `You can use this command to manage your project, include project's 
issue and so on.

Have fun!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
		project.RunProject()
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.Project.ProjectID, "project_id", "p", "", "aliyun aceesskey")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.Project.OrgID, "org_id", "o", "", "aliyun secretkey")
}
