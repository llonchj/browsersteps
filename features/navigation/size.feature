Feature: Size Browser

This feature supports sizing steps

    Scenario: Maximize
        Given I navigate to "/"
        When I maximize browser window
        Then I should be in "/"

    Scenario: Resize
        Given I navigate to "/"
        When I resize browser window size to 640x480
        When I resize browser window size to 800x600
        Then I should be in "/"
