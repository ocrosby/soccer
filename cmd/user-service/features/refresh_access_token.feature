Feature: Refresh Access Token

  Scenario: Refresh access token with valid refresh token
    Given the user has a valid refresh token
    When the user requests to refresh the access token
    Then the server should respond with a new access token

  Scenario: Refresh access token with invalid refresh token
    Given the user has an invalid refresh token
    When the user attempts to refresh the access token
    Then the server should respond with an error