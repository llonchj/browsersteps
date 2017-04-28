package browsersteps

import "github.com/DATA-DOG/godog"

func (b *BrowserSteps) iAcceptAlert() error {
	return b.GetWebDriver().AcceptAlert()
}

func (b *BrowserSteps) iDismissAlert() error {
	return b.GetWebDriver().DismissAlert()
}

func (b *BrowserSteps) buildProcessSteps(s *godog.Suite) {
	s.Step(`^I accept alert$`, b.iAcceptAlert)
	s.Step(`^I dismiss alert$`, b.iDismissAlert)

}
