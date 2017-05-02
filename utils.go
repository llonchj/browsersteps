package browsersteps

import (
	"errors"
	"fmt"
	"image"
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
		return b.URL.ResolveReference(u), nil
	}
	return u, nil
}

//GetCurrentWindowInnerSize returns window inner size
func (b *BrowserSteps) GetCurrentWindowInnerSize() (*selenium.Size, error) {
	widthI, err := b.GetWebDriver().ExecuteScript(`return window.innerWidth || 
		document.documentElement.clientWidth || document.body.clientWidth;
`, nil)
	if err != nil {
		return nil, err
	}
	width, ok := widthI.(float64)
	if !ok {
		return nil, fmt.Errorf("Can not cast width %+v to a float64", widthI)
	}
	heightI, err := b.GetWebDriver().ExecuteScript(`return window.innerHeight || 
		document.documentElement.clientHeight || document.body.clientHeight`, nil)
	if err != nil {
		return nil, err
	}
	height, ok := heightI.(float64)
	if !ok {
		return nil, fmt.Errorf("Can not cast height %+v to a float64", heightI)
	}
	return &selenium.Size{Width: int(width), Height: int(height)}, nil
}

//GetCurrentWindowScroll returns window scroll
func (b *BrowserSteps) GetCurrentWindowScroll() (*selenium.Point, error) {
	scrollXI, err := b.GetWebDriver().ExecuteScript(`return window.scrollX || 
		document.body.scrollLeft || document.documentElement.scrollLeft`, nil)
	if err != nil {
		return nil, err
	}
	scrollX, ok := scrollXI.(float64)
	if !ok {
		return nil, fmt.Errorf("Can not cast scrollX %+v to a float64", scrollXI)
	}
	scrollYI, err := b.GetWebDriver().ExecuteScript(`return window.scrollY ||
		document.body.scrollTop || document.documentElement.scrollTop`, nil)
	if err != nil {
		return nil, err
	}
	scrollY, ok := scrollYI.(float64)
	if !ok {
		return nil, fmt.Errorf("Can not cast scrollY %+v to a float64", scrollYI)
	}
	return &selenium.Point{X: int(scrollX), Y: int(scrollY)}, nil
}

//GetCurrentWindowViewport returns window scroll
func (b *BrowserSteps) GetCurrentWindowViewport() (image.Rectangle, error) {
	windowSize, err := b.GetCurrentWindowInnerSize()
	if err != nil {
		return image.Rectangle{}, err
	}
	scrollSize, err := b.GetCurrentWindowScroll()
	if err != nil {
		return image.Rectangle{}, err
	}
	return image.Rect(scrollSize.X, scrollSize.Y,
		scrollSize.X+windowSize.Width, scrollSize.Y+windowSize.Height), nil
}

//GetElementRect returns the element rectangle
func (b *BrowserSteps) GetElementRect(element selenium.WebElement) (image.Rectangle, error) {
	location, err := element.Location()
	if err != nil {
		return image.Rectangle{}, err
	}
	size, err := element.Size()
	if err != nil {
		return image.Rectangle{}, err
	}
	return image.Rect(location.X, location.Y,
		location.X+size.Width, location.Y+size.Height), nil
}
