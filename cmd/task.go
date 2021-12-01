/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/Berger7/yunxiao-go-client/config"
	"github.com/Berger7/yunxiao-go-client/pkg/task"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Get task information",
	Long:  `An easy way to get task.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
		task.RunTask()
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.Task.TaskID, "task_id", "i", "", "task id")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.Task.Status, "task_status", "u", "", "task status")
}
