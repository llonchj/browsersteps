package browsersteps

import (
	"errors"
	"fmt"

	"github.com/DATA-DOG/godog"
)

const (
	scrollEndScript = `window.scrollTo(0, document.body.scrollHeight);`
	scrollTopScript = `window.scrollTo(0, 0);`
)

func (b *BrowserSteps) iNavigateTo(browseURL string) error {
	u, err := b.GetURL(browseURL)
	if err != nil {
		return err
	}
	return b.GetWebDriver().Get(u.String())
}

func (b *BrowserSteps) iSwitchToNewWindow() error {
	wh, err := b.GetWebDriver().WindowHandles()
	if err != nil {
		return err
	}
	return b.GetWebDriver().SwitchWindow(wh[len(wh)-1])
}

func (b *BrowserSteps) iSwitchToMainWindow() error {
	wh, err := b.GetWebDriver().WindowHandles()
	if err != nil {
		return err
	}
	return b.GetWebDriver().SwitchWindow(wh[0])
}

func (b *BrowserSteps) iCloseNewWindow() error {
	wh, err := b.GetWebDriver().WindowHandles()
	if err != nil {
		return err
	}
	return b.GetWebDriver().CloseWindow(wh[len(wh)-1])
}

func (b *BrowserSteps) iSwitchToPreviousWindow() error {
	current, err := b.GetWebDriver().CurrentWindowHandle()
	if err != nil {
		return err
	}
	wh, err := b.GetWebDriver().WindowHandles()
	if err != nil {
		return err
	}
	if len(wh) < 2 {
		return errors.New("No previous window")
	}
	for i := 1; i < len(wh); i++ {
		if current == wh[i] {
			return b.GetWebDriver().CloseWindow(wh[i-1])
		}
	}
	return errors.New("Abnormal exception: No window found")
}

func (b *BrowserSteps) iSwitchToWindowHaving(what, match string) error {
	if what == "url" {
		u, err := b.GetURL(match)
		if err != nil {
			return err
		}
		match = u.String()
	}
	current, err := b.GetWebDriver().CurrentWindowHandle()
	if err != nil {
		return err
	}
	wh, err := b.GetWebDriver().WindowHandles()
	if err != nil {
		return err
	}
	for i := 0; i < len(wh); i++ {
		err := b.GetWebDriver().SwitchWindow(wh[i])
		if err != nil {
			return err
		}
		switch what {
		case "url":
			URL, err := b.GetWebDriver().CurrentURL()
			if err != nil {
				return err
			}
			if URL == match {
				return nil
			}
		case "title":
			title, err := b.GetWebDriver().Title()
			if err != nil {
				return err
			}
			if title == match {
				return nil
			}
		default:
			return fmt.Errorf("Invalid property. Found '%s', Expected 'url' or 'title'", what)
		}
	}
	err = b.GetWebDriver().SwitchWindow(current)
	if err != nil {
		return err
	}
	return errors.New("Abnormal exception: No window found")
}

func (b *BrowserSteps) iResizeBrowserWindowTo(w, h int) error {
	current, err := b.GetWebDriver().CurrentWindowHandle()
	if err != nil {
		return err
	}
	return b.GetWebDriver().ResizeWindow(current, w, h)
}

func (b *BrowserSteps) iMaximizeResizeBrowserWindow() error {
	current, err := b.GetWebDriver().CurrentWindowHandle()
	if err != nil {
		return err
	}
	return b.GetWebDriver().MaximizeWindow(current)
}

func (b *BrowserSteps) iScrollTo(where string) error {
	var script string
	switch where {
	case "top":
		script = scrollTopScript
	case "end":
		script = scrollEndScript
	default:
		return fmt.Errorf("Invalid scroll direction. Got: '%s', allowed: 'top' or 'end'", where)
	}
	_, err := b.GetWebDriver().ExecuteScript(script, []interface{}{})
	return err
}

func (b *BrowserSteps) iScrollToElement(selector, by string) error {
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}
	pt, err := element.LocationInView()
	if err != nil {
		return err
	}
	_, err = b.GetWebDriver().ExecuteScript(
		fmt.Sprintf("window.scrollTo(%d, %d);", pt.X, pt.Y), []interface{}{})
	return err
}

func (b *BrowserSteps) buildNavigationSteps(s *godog.Suite) {
	s.Step(`^I navigate to "([^"]*)"$`, b.iNavigateTo)
	s.Step(`^I navigate forward$`, func() error { return b.GetWebDriver().Forward() })
	s.Step(`^I navigate back$`, func() error { return b.GetWebDriver().Back() })

	s.Step(`^I (reload|refresh)( the)? page$`, func() error { return b.GetWebDriver().Refresh() })

	s.Step(`^I switch to new window$`, b.iSwitchToNewWindow)
	s.Step(`^I switch to main window$`, b.iSwitchToMainWindow)
	s.Step(`^I switch to previous window$`, b.iSwitchToPreviousWindow)
	s.Step(`^I switch to window having (title|url) "(.*?)"$`, b.iSwitchToWindowHaving)
	s.Step(`^I close new window$`, b.iCloseNewWindow)

	s.Step(`^I resize browser window size to width (\d+) and height (\d+)$`, b.iResizeBrowserWindowTo)
	s.Step(`^I resize browser window size to (\d+)x(\d+)$`, b.iResizeBrowserWindowTo)
	s.Step(`^I maximize browser window$`, b.iMaximizeResizeBrowserWindow)

	s.Step(`^I scroll to "([^"]*)" `+ByOption+`$`, b.iScrollToElement)
	s.Step(`^I scroll to (top|end) of page$`, b.iScrollTo)

}
