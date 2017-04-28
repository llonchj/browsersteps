package browsersteps

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/tebeka/selenium"

	"github.com/llonchj/browsersteps"
)

func iWaitFor(amount int, unit string) error {
	u := time.Second
	fmt.Printf("Waiting for %d %s", amount, unit)
	time.Sleep(u * time.Duration(amount))
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I wait for (\d+) (milliseconds|millisecond|seconds|second)$`, iWaitFor)

	browsersteps.NewBrowserSteps(s,
		selenium.Capabilities{"browserName": "chrome"},
		"")
}

func TestMain(m *testing.M) {
	status := godog.Run("example", FeatureContext)
	os.Exit(status)
}
