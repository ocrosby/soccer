Feature: Obtain Authorization Code

  Scenario: User requests authorization code
    Given the user has navigated to the authorization URL
    When the user authorizes the application
    Then the authorization server should redirect to the callback URL with an authorization code
