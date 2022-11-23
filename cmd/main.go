package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"tvf-saleshub-backend/cmd/api"
	"tvf-saleshub-backend/config"
	"tvf-saleshub-backend/internal/logger"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const configFile = "config.yaml"

// run server with CLI
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server CLI",
	Long:  "run server with CLI",
}

// init initializes the env and logger.
func init() {
	initEnv()
	configs := initConfig()
	logger := initLogger()

	apiCmd := api.NewServerCmd(configs, logger)
	rootCmd.AddCommand(apiCmd)
}

// main is the entry point for the run command.
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("run command has failed with error: %v\n", err)
		os.Exit(1)
	}
}

// initEnv loads the. env file
func initEnv() {
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("skip load .env file")
		return
	}
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("init env has failed failed with error: %v\n", err)
		os.Exit(1)
	}
}

// initLogger creates a new zap. Logger
func initLogger() *zap.Logger {
	return logger.NewLogger()
}

// initConfig initializes the config.
func initConfig() *config.Configurations {
	viper.SetConfigType("yaml")

	// Expand environment variables inside the config file
	b, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	expand := os.ExpandEnv(string(b))
	configReader := strings.NewReader(expand)

	viper.AutomaticEnv()

	if err := viper.ReadConfig(configReader); err != nil {
		fmt.Printf("read config has failed with error: %v\n", err)
		os.Exit(1)
	}

	configs := config.Configurations{}
	if err := viper.Unmarshal(&configs); err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	return &configs
}
