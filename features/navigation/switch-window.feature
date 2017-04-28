Feature: Switch Browser Window

This feature supports switch browser windows

    Scenario: Navigate to switch window
        Given I am a anonymous user
        When I navigate to "/other.html"
        And I click "long page" link text
        And I switch to new window
        And I should be in "/long.html"
        And I switch to main window
        Then I should be in "/other.html"
        And I switch to new window
        Then I should be in "/long.html"
        And I close new window
        And I switch to main window
        Then I should be in "/other.html"

    Scenario: Navigate to direct window having url
        Given I am a anonymous user
        When I navigate to "/other.html"
        And I click "long page" link text
        And I switch to new window
        Then I should be in "/long.html"
        And I switch to window having title "Other"
        Then I should be in "/other.html"

    Scenario: Navigate to direct window having url
        Given I am a anonymous user
        When I navigate to "/other.html"
        And I click "long page" link text
        And I switch to new window
        Then I should be in "/long.html"
        And I switch to window having url "/other.html"
        Then I should be in "/other.html"
