Feature: Google Search

@remote
    Scenario: Submit a form
        Given I navigate to "https://google.com/"
        When I write "Google" to "q" name
        And I click "btnG" name
        And I wait for 2 seconds
        Then I should see "Google" in "Google" link text
