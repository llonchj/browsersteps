Feature: Reload browser

This feature supports reload step

    Scenario: Reload the page
        Given I navigate to "/"
        When I reload the page
        Then I should be in "/"

    Scenario: Refresh the page
        Given I navigate to "/"
        When I refresh page
        Then I should be in "/"
