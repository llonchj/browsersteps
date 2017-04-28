Feature: Navigate back

This feature supports navigate back step

    Scenario: Navigate back
        Given I navigate to "/"
        And I navigate to "/other.html"
        When I navigate back
        Then I should be in "/"
