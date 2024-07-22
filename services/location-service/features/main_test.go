//go:build godog
// +build godog

package features

import (
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/ocrosby/soccer/services/location-service/features/steps"
)

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "pretty",
		Paths:     []string{"."},               // This includes all feature files in the directory
		Randomize: time.Now().UTC().UnixNano(), // Optional: randomizes scenario execution order
		Output:    colors.Colored(os.Stdout),
	}

	status := godog.TestSuite{
		Name:                 "Countries Feature",
		TestSuiteInitializer: steps.InitializeCountriesTestSuite,
		ScenarioInitializer:  steps.InitializeCountriesScenario,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
