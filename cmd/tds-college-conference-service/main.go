package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tds-college-conference-service",
	Short: "TDS College Conference Service",
	Long:  `A service for managing TDS College Conferences.`,
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Starts the TopDrawerSoccer College Conference Service server.`,
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		initConfig(env)

		// Read the PORT variable using Viper
		port := viper.GetString("PORT")
		if port == "" {
			log.Fatal("PORT must be set in .env or as an environment variable")
		}

		r := mux.NewRouter()

		// Register the health check handler
		r.HandleFunc("/health", HealthCheckHandler)

		// Existing handlers
		r.HandleFunc("/", HomeHandler)

		err := http.ListenAndServe(":"+port, r)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func initConfig(env string) {
	// Fallback to GO_ENV if the --env flag is not provided
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	// Default to development if no environment is specified
	if env == "" {
		env = "dev"
	}

	configFile := fmt.Sprintf("env.%s", env)

	viper.SetConfigType("env")                                         // Specify the type of the configuration file
	viper.SetConfigName(configFile)                                    // Name of the config file (without extension)
	viper.AddConfigPath("./cmd/tds-college-conference-service/config") // Path to look for the config file in

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Automatically override values with environment variables if they exist
	viper.AutomaticEnv()
}

func main() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("env", "e", "dev", "Specify the environment")

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
