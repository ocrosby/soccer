Feature: Coaching Changes
  As a Soccer enthusiast
  I want to know the changes in coaching staff of DI programs
  So that I can be aware of the state of college Soccer

  Scenario: All Coaching Changes
    When I request all coaching changes
    Then there should be no errors
    And the response should be successful
    And there should be both male and female change results

  Scenario: Male Coaching Changes
    When I request male coaching changes
    Then there should be no errorsswagger generate spec -o ./services/tds-coaching-change-service/swagger.yaml --scan-models ./services/tds-coaching-change-service/
    And the response should be successful
    And all change results should be male

  Scenario: Female Coaching Changes
    When I request female coaching changes
    Then there should be no errors
    And the response should be successful
    And all change results should be female
