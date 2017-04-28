package browsersteps

import (
	"fmt"
	"os"
	"testing"
	"time"

	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/DATA-DOG/godog"
	"github.com/tebeka/selenium"
)

func iWaitFor(amount int, unit string) error {
	u := time.Second
	fmt.Printf("Waiting for %d %s", amount, unit)
	time.Sleep(u * time.Duration(amount))
	return nil
}

func FeatureContext(s *godog.Suite) {
	// selenium.SetDebug(true)

	s.Step(`^I wait for (\d+) (milliseconds|millisecond|seconds|second)$`, iWaitFor)

	var server *httptest.Server
	bs, _ := NewBrowserSteps(s,
		selenium.Capabilities{"browserName": "chrome"},
		"")

	s.BeforeSuite(func() {
		server = httptest.NewServer(http.FileServer(http.Dir("./public")))
		u, _ := url.Parse(server.URL)
		bs.SetBaseURL(u)
	})

	s.AfterSuite(func() {
		if server != nil {
			server.Close()
			server = nil
		}
	})
}

func TestMain(m *testing.M) {
	status := godog.Run("browsersteps", FeatureContext)
	os.Exit(status)
}
