Feature: Alert Assertion

This feature supports alert assertion steps

    Scenario: Alert URL
        Given I am a anonymous user
        When I navigate to "/"
        And I click "Alert" link text
        Then I should see alert text as "This is an alert message"
        And I accept alert

