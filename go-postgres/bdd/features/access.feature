Feature: access policy
  Scenario: administrator log in
    Given I have an administrator account with login "admin"
    When I make log in with password "admin"
    Then I have "god permissions" with credentials "admin":"admin"

  Scenario: reader log in
    Given I have an reader account with login "guest"
    When I make log in with password "guest"
    Then I have "read only permissions" with credentials "guest":"guest"

  Scenario: redactor log in
    Given I have an redactor account with login "user"
    When I make log in with password "user"
    Then I have "limited crud" with credentials "user":"user"

