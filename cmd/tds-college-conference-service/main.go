package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ocrosby/soccer/pkg/common"
	"github.com/ocrosby/soccer/pkg/tds"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

// HomeHandler responds to HTTP GET requests on the root "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	numberOfBytesWritten, err := fmt.Fprintln(w, "Welcome to the Gorilla Mux Service!")
	if err != nil {
		format := "Internal Server Error. Bytes written before error: %d"
		errMsg := fmt.Sprintf(format, numberOfBytesWritten)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
}

// HealthCheckHandler responds to HTTP GET requests on the "/health" endpoint
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	numberOfBytesWritten, err := fmt.Fprintln(w, "OK")
	if err != nil {
		format := "Internal Server Error. Bytes written before error: %d"
		errMsg := fmt.Sprintf(format, numberOfBytesWritten)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}
}

// GetConferencesHandler responds to HTTP GET requests on the "/conferences" endpoint
func GetConferencesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err          error
		division     common.Division
		gender       common.Gender
		conferences  []tds.Conference
		jsonResponse []byte
	)

	// Retrieve query parameters
	queryParams := r.URL.Query()

	// Retrieve the "gender" query parameter
	genderString := queryParams.Get("gender") // returns an empty string if not specified

	// Retrieve the "division" query parameter
	divisionString := queryParams.Get("division") // returns an empty string if not specified

	if gender, err = common.StringToGender(genderString); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if division, err = common.StringToDivision(divisionString); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if conferences, err = tds.Conferences(gender, division); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a JSON response
	if jsonResponse, err = json.Marshal(conferences); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	if _, err = w.Write(jsonResponse); err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

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

		// Existing handlers
		r.HandleFunc("/", HomeHandler)

		// Register the health check handler
		r.HandleFunc("/health", HealthCheckHandler)

		// Register the conferences handler
		r.HandleFunc("/conferences", GetConferencesHandler).Methods("GET")

		log.Printf("Server starting on http://localhost:%s\n", port)
		log.Printf("Test the API at http://localhost:%s/conferences?gender=female&division=di\n", port)

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
