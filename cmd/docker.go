package cmd

import (
	"fmt"
	"github.com/ocrosby/soccer/pkg/common"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// dockerCmd represents the docker subcommand
var dockerCmd = &cobra.Command{
	Use:   "docker [service]",
	Short: "Builds a Docker image for a service",
	Long:  `Builds a Docker image for the specified service, assuming there is a Dockerfile in the service's directory.`,
	Args:  cobra.ExactArgs(1), // Requires exactly one argument: the service name
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]

		if serviceName == "*" {
			err := buildAllDockerImages()
			if err != nil {
				fmt.Printf("Error building Docker images: %v\n", err)
				os.Exit(2)
			}
		} else {
			err := buildDockerImage(serviceName)
			if err != nil {
				fmt.Printf("Error building Docker image for service %s: %v\n", serviceName, err)
				os.Exit(1)
			}
		}
	},
}

func buildAllDockerImages() error {
	var (
		serviceName        string
		failedServices     []string
		successfulServices []string
	)

	fmt.Println("Building all services...")
	services, err := common.ServiceNames("services")

	if err != nil {
		fmt.Printf("Error getting service names: %v\n", err)
		os.Exit(1)
	}

	// Build Docker images for all services
	for _, serviceName = range services {
		if err = buildDockerImage(serviceName); err != nil {
			failedServices = append(failedServices, serviceName)
		} else {
			successfulServices = append(successfulServices, serviceName)
		}
	}

	failureCount := len(failedServices)
	if failureCount > 0 {
		if len(successfulServices) > 0 {
			for _, serviceName = range successfulServices {
				fmt.Printf("Docker image built successfully for service: %s\n", serviceName)
			}
		} else {
			fmt.Println("No Docker images were built successfully.")
		}

		for _, serviceName = range failedServices {
			fmt.Printf("Failed to build Docker image for service: %s\n", serviceName)
		}

		return fmt.Errorf("Failed to build %d Docker images", failureCount)
	} else {
		if len(successfulServices) > 0 {
			fmt.Println("All Docker images built successfully.")
		} else {
			fmt.Println("No services found to build.")
		}
	}

	return nil
}

func buildDockerImage(serviceName string) error {
	fmt.Printf("Building Docker image for service: %s\n", serviceName)

	// Assuming Dockerfile is located in services/<serviceName>/Dockerfile
	dockerfilePath := filepath.Join("services", serviceName, "Dockerfile")
	// contextPath := filepath.Join("services", serviceName)
	contextPath := "."

	// Check if the Dockerfile exists
	if _, err := os.Stat(dockerfilePath); os.IsNotExist(err) {
		return fmt.Errorf("Dockerfile does not exist in service %s", serviceName)
	}

	// Build the Docker image
	command := fmt.Sprintf("docker build --build-arg SERVICE_NAME=%s -t %s-image -f %s %s", serviceName, serviceName, dockerfilePath, contextPath)
	if err := common.ExecuteCommand(command); err != nil {
		return err
	}

	fmt.Println("Docker image built successfully.")

	// Prune dangling images
	command = "docker image prune -f"
	if err := common.ExecuteCommand(command); err != nil {
		return err
	}

	fmt.Println("Dangling images pruned successfully.")

	return nil
}

func init() {
	buildCmd.AddCommand(dockerCmd) // Add docker as a subcommand to build
}
