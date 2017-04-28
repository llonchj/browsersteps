Feature: Navigate

This feature supports navigation to URL

    Scenario: Navigate
        Given I am a anonymous user
        When I navigate to "/other.html"
        Then I should be in "/other.html"
