package browsersteps

import (
	"context"
	"fmt"
	"time"

	"github.com/DATA-DOG/godog"
)

func (b *BrowserSteps) iShouldBeIn(browseURL string) error {
	u, err := b.GetURL(browseURL)
	if err != nil {
		return err
	}

	currentURL, err := b.GetWebDriver().CurrentURL()
	if err != nil {
		return err
	}

	if currentURL != u.String() {
		return fmt.Errorf("URL does not match. Expected '%s', Found '%s'",
			u.String(), currentURL)
	}
	return nil
}

func (b *BrowserSteps) iShouldSeePageTitleAs(expectedTitle string) error {
	title, err := b.GetWebDriver().Title()
	if err != nil {
		return err
	}

	if expectedTitle != title {
		return fmt.Errorf("Title does not match. Expected '%s', Found '%s'",
			expectedTitle, title)
	}
	return nil
}

func (b *BrowserSteps) iShouldSeeIn(expectedText, selector, by string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for {
		select {
		case <-time.After(50 * time.Millisecond):
			element, err := b.GetWebDriver().FindElement(by, selector)
			if err != nil {
				break
			}
			gotText, err := element.Text()
			if err != nil {
				break
			}
			if expectedText != gotText {
				break
			}
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (b *BrowserSteps) iShouldSee(selector, by string) error {
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}

	elementRect, err := b.GetElementRect(element)
	if err != nil {
		return err
	}
	viewportRect, err := b.GetCurrentWindowViewport()
	if err != nil {
		return err
	}
	if !elementRect.In(viewportRect) {
		return fmt.Errorf("Element '%s' %s not in the Window area", selector, by)
	}
	return nil
}

func (b *BrowserSteps) iShouldNotSee(selector, by string) error {
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}

	elementRect, err := b.GetElementRect(element)
	if err != nil {
		return err
	}
	viewportRect, err := b.GetCurrentWindowViewport()
	if err != nil {
		return err
	}

	if elementRect.In(viewportRect) {
		return fmt.Errorf("Element '%s' %s in the Window area", selector, by)
	}
	return nil
}

func (b *BrowserSteps) iShouldSeeAlertAs(expectedText string) error {
	gotText, err := b.GetWebDriver().AlertText()
	if err != nil {
		return err
	}
	if expectedText != gotText {
		return fmt.Errorf("Alert Text Mismatch. Got '%s', Expected '%s'", gotText, expectedText)
	}
	return nil
}

func (b *BrowserSteps) buildAssertionSteps(s *godog.Suite) {
	s.Step(`^I should be in "([^"]*)"$`, b.iShouldBeIn)
	s.Step(`^I should see page title as "(.*?)"$`, b.iShouldSeePageTitleAs)

	s.Step(`^I should see "([^"]*)" in "([^"]*)" `+ByOption+`$`, b.iShouldSeeIn)

	s.Step(`^I should see "([^"]*)" `+ByOption+`$`, b.iShouldSee)
	s.Step(`^I should not see "([^"]*)" `+ByOption+`$`, b.iShouldNotSee)

	s.Step(`I should see alert text as "(.*?)"$`, b.iShouldSeeAlertAs)
}
