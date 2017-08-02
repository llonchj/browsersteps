Feature: Should See Assertion

This feature supports visible/invisible assertion steps

    Scenario: Should See Element
        Given I am a anonymous user
        When I navigate to "/js.html"
        And I submit "submit" id
        Then I should see "Done!" in "p" id
