Feature: Navigate forward

This feature supports navigate forward step

    Scenario: Navigate forward
        Given I navigate to "/"
        And I navigate to "/other.html"
        And I navigate back
        Then I should be in "/"
        When I navigate forward
        Then I should be in "/other.html"
