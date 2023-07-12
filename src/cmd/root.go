package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/oatsaysai/simple-core-bank/src/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "simple-core-bank",
	Short: "simple-core-bank",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "cfg/config.yaml", "config file (default is cfg/config.yaml)")
}

func initConfig() {

	viper.SetConfigFile(configFile)

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetDefault("Log.Level", "debug")
	viper.SetDefault("Log.Color", true)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("unable to read config: %v\n", err)
		os.Exit(1)
	}
}

func getLogger() (logger.Logger, error) {
	logLevel := viper.GetString("Log.Level")
	logLevel = logger.NormalizeLogLevel(logLevel)

	logColor := viper.GetBool("Log.Color")
	logJSON := viper.GetBool("Log.JSON")

	logger, err := logger.NewLogger(&logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logLevel,
		ConsoleJSONFormat: logJSON,
		Color:             logColor,
	}, logger.InstanceZapLogger)
	if err != nil {
		return nil, err
	}
	return logger, nil
}
