Feature: College Player Information Retrieval
    In order to get information about a college player
    As a college Soccer fan
    I want to be able to retrieve information about a specific player

  Scenario: Get information for a specific player
    Given I have the player ID "12345"
    When I request information for the player
    Then I should receive the player's information
