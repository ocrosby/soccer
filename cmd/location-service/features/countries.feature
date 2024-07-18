Feature: Countries
  As a user of the location service
  I want to manage country information
  So that I can maintain an accurate and comprehensive database of countries

  Scenario: Add a new country
    Given I am an authorized user
    When I submit a request to add a new country with the name "France" and code "FR"
    Then the country "France" with code "FR" should be added to the database

  Scenario: Retrieve a country by code
    Given I am an authorized user
    And the country with code "FR" exists in the database
    When I submit a request to retrieve a country with code "FR"
    Then I should receive the details of the country "France"

  Scenario: Update a country's name
    Given I am an authorized user
    And the country with code "FR" exists in the database
    When I submit a request to update the country with code "FR" to change its name to "Republic of France"
    Then the name of the country with code "FR" should be updated to "Republic of France" in the database

  Scenario: Delete a country
    Given I am an authorized user
    And the country with code "FR" exists in the database
    When I submit a request to delete the country with code "FR"
    Then the country with code "FR" should be removed from the database