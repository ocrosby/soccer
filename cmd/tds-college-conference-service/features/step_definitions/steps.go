package step_definitions

import "github.com/cucumber/godog"

func iAccessTheRootEndpoint() error {
	return godog.ErrPending
}

func iAmAUser() error {
	return godog.ErrPending
}

func iShouldSeeAWelcomeMessage() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I access the root endpoint$`, iAccessTheRootEndpoint)
	ctx.Step(`^I am a user$`, iAmAUser)
	ctx.Step(`^I should see a welcome message$`, iShouldSeeAWelcomeMessage)
}
