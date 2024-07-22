package cmd

import (
	"fmt"
	"github.com/ocrosby/soccer/pkg/common"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var bddCmd = &cobra.Command{
	Use:   "bdd [service]",
	Short: "Runs BDD tests for a service",
	Long:  `Runs BDD tests for the specified service.`,
	Args:  cobra.ExactArgs(1), // Requires exactly one argument: the service name
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]

		if serviceName == "*" {
			err := runAllBDDTests()
			if err != nil {
				fmt.Printf("Error running BDD tests: %v\n", err)
				os.Exit(2)
			}
		} else {
			err := runBDDTests(serviceName)
			if err != nil {
				fmt.Printf("Error running BDD tests for service %s: %v\n", serviceName, err)
				os.Exit(1)
			}
		}
	},
}

func runBDDTests(serviceName string) error {
	// Get the current working directory
	workingDirectory, _ := os.Getwd()

	// Run the BDD tests for the specified service
	fmt.Printf("Running BDD tests for service %s...\n", serviceName)

	// Add code here to run the BDD tests for the specified service
	featuresDirectory := fmt.Sprintf("./services/%s/features", serviceName)
	featuresDirectory = filepath.Join(workingDirectory, featuresDirectory)

	if _, err := os.Stat(featuresDirectory); os.IsNotExist(err) {
		fmt.Printf("The working directory is %s\n", workingDirectory)
		fmt.Printf("No BDD tests found for service %s\n", serviceName)
		fmt.Printf("The features directory %s does not exist\n", featuresDirectory)
		return nil
	}

	command := fmt.Sprintf("go test -tags=godog %s", featuresDirectory)

	if err := common.ExecuteCommand(command); err != nil {
		return err
	}

	return nil
}

func runAllBDDTests() error {
	// Run the BDD tests for all services
	fmt.Println("Running all BDD tests...")
	services, err := common.ServiceNames("services")

	if err != nil {
		fmt.Printf("Error getting service names: %v\n", err)
		os.Exit(1)
	}

	for _, serviceName := range services {
		if err = runBDDTests(serviceName); err != nil {
			fmt.Printf("Error running BDD tests for service %s: %v\n", serviceName, err)
			os.Exit(1)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(bddCmd) // Add the 'bdd' command to the root command
}
