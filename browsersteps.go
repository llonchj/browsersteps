package browsersteps

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/DATA-DOG/godog"
	"github.com/tebeka/selenium"
)

/*BrowserSteps represents a WebDriver context to run the Scenarios*/
type BrowserSteps struct {
	wd             selenium.WebDriver
	Capabilities   selenium.Capabilities
	DefaultURL     string
	URL            *url.URL
	ScreenshotPath string
}

/*SetBaseURL sets the absolute URL used to complete relative URLs*/
func (b *BrowserSteps) SetBaseURL(url *url.URL) error {
	if !url.IsAbs() {
		return errors.New("BaseURL must be absolute")
	}
	b.URL = url
	return nil
}

func (b *BrowserSteps) iAmAnAnonymousUser() error {
	return b.GetWebDriver().DeleteAllCookies()
}

func (b *BrowserSteps) iWriteTo(text, selector, by string) error {
	// Click the element
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}

	err = element.Clear()
	if err != nil {
		return err
	}
	return element.SendKeys(text)
}

func (b *BrowserSteps) iClick(selector, by string) error {
	// Submit the element
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}
	return element.Click()
}

func (b *BrowserSteps) iSubmit(selector, by string) error {
	// Submit the element
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}
	return element.Submit()
}

func (b *BrowserSteps) iMoveTo(selector, by string) error {
	element, err := b.GetWebDriver().FindElement(by, selector)
	if err != nil {
		return err
	}
	return element.MoveTo(0, 0)
}

//BeforeScenario is executed before each scenario
func (b *BrowserSteps) BeforeScenario(a interface{}) {
	var err error
	b.wd, err = selenium.NewRemote(b.Capabilities, b.DefaultURL)
	if err != nil {
		log.Panic(err)
	}
}

//AfterScenario is executed after each scenario
func (b *BrowserSteps) AfterScenario(a interface{}, err error) {
	if err != nil && b.ScreenshotPath != "" {
		filename := fmt.Sprintf("FAILED STEP - %s.png", err.Error())

		buff, err := b.GetWebDriver().Screenshot()
		if err != nil {
			fmt.Printf("Error %+v\n", err)
		}

		if _, err := os.Stat(b.ScreenshotPath); os.IsNotExist(err) {
			os.MkdirAll(b.ScreenshotPath, 0755)
		}
		pathname := filepath.Join(b.ScreenshotPath, filename)
		ioutil.WriteFile(pathname, buff, 0644)
	}
	b.GetWebDriver().Quit()
}

func (b *BrowserSteps) buildSteps(s *godog.Suite) {
	b.buildNavigationSteps(s)
	b.buildAssertionSteps(s)
	b.buildProcessSteps(s)

	s.Step(`^I am a anonymous user$`, b.iAmAnAnonymousUser)

	s.Step(`^I write "([^"]*)" to "([^"]*)" `+ByOption+`$`, b.iWriteTo)
	s.Step(`^I click "([^"]*)" `+ByOption+`$`, b.iClick)
	s.Step(`^I submit "([^"]*)" `+ByOption+`$`, b.iSubmit)

	s.Step(`^I move to "([^"]*)" `+ByOption+`$`, b.iMoveTo)

}

//NewBrowserSteps starts a new BrowserSteps instance.
func NewBrowserSteps(s *godog.Suite, cap selenium.Capabilities, defaultURL string) *BrowserSteps {
	bs := &BrowserSteps{Capabilities: cap, DefaultURL: defaultURL, ScreenshotPath: os.Getenv("SCREENSHOT_PATH")}
	bs.buildSteps(s)

	s.BeforeScenario(bs.BeforeScenario)
	s.AfterScenario(bs.AfterScenario)

	return bs
}
