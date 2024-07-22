Feature: Access Protected Resource

  Scenario: Access resource with valid access token
    Given the user has a valid access token
    When the user requests a protected resource with the access token
    Then the server should allow access to the resource

  Scenario: Access resource with expired access token
    Given the user has an expired access token
    When the user requests a protected resource with the expired access token
    Then the server should deny access and respond with an error