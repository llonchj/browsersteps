package browsersteps

import (
	"errors"
	"fmt"

	"github.com/DATA-DOG/godog"
)

func (b *BrowserSteps) iNavigateTo(browseURL string) error {
	u, err := b.GetURL(browseURL)
	if err != nil {
		return err
	}
	return b.GetWebDriver().Get(u.String())
}

func (b *BrowserSteps) iNavigateForward() error {
	return b.GetWebDriver().Forward()
}

func (b *BrowserSteps) iNavigateBack() error {
	return b.GetWebDriver().Back()
}

func (b *BrowserSteps) iReloadPage() error {
	return b.GetWebDriver().Refresh()
}

func (b *BrowserSteps) iSwitchToNewWindow() error {
	wh, _ := b.GetWebDriver().WindowHandles()
	return b.GetWebDriver().SwitchWindow(wh[len(wh)-1])
}

func (b *BrowserSteps) iSwitchToMainWindow() error {
	wh, _ := b.GetWebDriver().WindowHandles()
	return b.GetWebDriver().SwitchWindow(wh[0])
}

func (b *BrowserSteps) iCloseNewWindow() error {
	wh, _ := b.GetWebDriver().WindowHandles()
	return b.GetWebDriver().CloseWindow(wh[len(wh)-1])
}

func (b *BrowserSteps) iSwitchToPreviousWindow() error {
	current, _ := b.GetWebDriver().CurrentWindowHandle()
	wh, _ := b.GetWebDriver().WindowHandles()
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
	current, _ := b.GetWebDriver().CurrentWindowHandle()
	wh, _ := b.GetWebDriver().WindowHandles()
	for i := 0; i < len(wh); i++ {
		b.GetWebDriver().SwitchWindow(wh[i])
		switch what {
		case "url":
			URL, _ := b.GetWebDriver().CurrentURL()
			if URL == match {
				return nil
			}
		case "title":
			title, _ := b.GetWebDriver().Title()
			if title == match {
				return nil
			}
		default:
			return fmt.Errorf("Invalid property. Found '%s', Expected 'url' or 'title'", what)
		}
	}
	b.GetWebDriver().SwitchWindow(current)
	return errors.New("Abnormal exception: No window found")
}

func (b *BrowserSteps) iResizeBrowserWindowTo(w, h int) error {
	current, _ := b.GetWebDriver().CurrentWindowHandle()
	return b.GetWebDriver().ResizeWindow(current, w, h)
}

func (b *BrowserSteps) iMaximizeResizeBrowserWindow() error {
	current, _ := b.GetWebDriver().CurrentWindowHandle()
	return b.GetWebDriver().MaximizeWindow(current)
}

func (b *BrowserSteps) buildNavigationSteps(s *godog.Suite) {
	s.Step(`^I navigate to "([^"]*)"$`, b.iNavigateTo)
	s.Step(`^I navigate forward$`, b.iNavigateForward)
	s.Step(`^I navigate back$`, b.iNavigateBack)

	s.Step(`^I reload the page$`, b.iReloadPage)
	s.Step(`^I refresh page$`, b.iReloadPage)

	s.Step(`^I switch to new window$`, b.iSwitchToNewWindow)
	s.Step(`^I switch to main window$`, b.iSwitchToMainWindow)
	s.Step(`^I switch to previous window$`, b.iSwitchToPreviousWindow)
	s.Step(`^I switch to window having (title|url) "(.*?)"$`, b.iSwitchToWindowHaving)
	s.Step(`^I close new window$`, b.iCloseNewWindow)

	s.Step(`^I resize browser window size to width (\d+) and height (\d+)$`, b.iResizeBrowserWindowTo)
	s.Step(`^I resize browser window size to (\d+)x(\d+)$`, b.iResizeBrowserWindowTo)
	s.Step(`^I maximize browser window$`, b.iMaximizeResizeBrowserWindow)
}
