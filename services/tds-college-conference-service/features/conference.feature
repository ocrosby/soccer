Feature: TopDrawerSoccer College Conference
  In order to be informed about NCAA Soccer
  As a user
  I need to be able to access college conference information

  Scenario: Health Check
    Given I am a user
    When I access the root endpoint
    Then I should see a welcome message
