/*
Copyright © 2021 berger7 bojun.cbj@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Berger7/yunxiao-go-client/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yunxiao-go-client",
	Short: "An easy client for Yunxiao",
	Long: `An easy client for yunxiao (https://devops.aliyun.com). This command
	include project,repository,cicd systems.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yunxiao.yaml)")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.AK, "accesskey", "a", "", "aliyun aceesskey")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.SK, "secretkey", "s", "", "aliyun secretkey")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.Endpoint, "endpoint", "e", "", "aliyun endpoint")
	rootCmd.PersistentFlags().StringVarP(&config.Cfg.Action, "action", "c", "", "action")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".yunxiao.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		viper.Unmarshal(&config.Cfg)
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
