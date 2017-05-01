package browsersteps

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/tebeka/selenium"
)

/*ByOption is the regular expression for allowed selenium.By* in step
definitions*/
const ByOption = `(id|xpath|link text|partial link text|name|tag name|class name|css selector)`

//GetWebDriver returns the webdriver
func (b *BrowserSteps) GetWebDriver() selenium.WebDriver {
	return b.wd
}

//GetURL returns a absolute url given a absolute or relative URL
func (b *BrowserSteps) GetURL(URL string) (*url.URL, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("Malformed URL: %s", URL)
	}
	if !u.IsAbs() {
		if b.URL == nil {
			return nil, errors.New("Using a relative URL without a base URL defined. Invoke BrowserSteps.SetBaseURL")
		}
		m := b.URL.ResolveReference(u)
		return m, nil
	}
	return u, nil
}
