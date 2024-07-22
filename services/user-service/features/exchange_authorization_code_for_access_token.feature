Feature: Exchange Authorization Code for Access Token

  Scenario: Exchange valid authorization code for an access token
    Given the user has a valid authorization code
    When the user exchanges the authorization code for an access token
    Then the server should respond with an access token

  Scenario: Exchange invalid authorization code
    Given the user has an invalid authorization code
    When the user attempts to exchange the authorization code for an access token
    Then the server should respond with an error