Feature: Scroll

Scroll the page

    Scenario: Scroll Top and Bottom
        Given I am a anonymous user
        And I navigate to "/long.html"
        When I scroll to end of page
        Then I should see "#home_link" css selector
        And I should not see "#bottom_link" css selector
        And I scroll to top of page
        And I should see "#bottom_link" css selector
        And I should not see "#home_link" css selector

    Scenario: Scroll Top and Bottom
        Given I am a anonymous user
        And I navigate to "/long.html"
        When I scroll to "nulla-quis" id
        Then I should see "nulla-quis" id

