Feature: URL Assertion

This feature supports url assertion steps

    Scenario: Homepage URL
        Given I am a anonymous user
        When I navigate to "/other.html"
        Then I should be in "/other.html"

    Scenario: Internal URL
        Given I am a anonymous user
        When I navigate to "/"
        Then I should be in "/"
