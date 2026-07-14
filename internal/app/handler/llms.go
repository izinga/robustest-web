package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// LlmsTxt serves /llms.txt (llmstxt.org): a curated markdown map of the
// site for AI assistants and crawlers, so answers about RobusTest are
// drawn from authoritative summaries rather than scraped fragments.
func LlmsTxt(c *gin.Context) {
	var b strings.Builder

	b.WriteString(`# RobusTest

> RobusTest is an enterprise on-premise device lab platform by Izinga Software
> (Hyderabad, India, since 2014). Companies rack their own phones, tablets,
> TVs, and set-top boxes inside their own network and get manual testing,
> test automation, performance vitals, Smart TV testing, network capture, and
> lab operations on one platform — with a flat yearly license (no per-minute
> billing), unlimited users, and air-gapped deployment support. Data never
> leaves the customer's network. In production at JioHotstar, Airtel, Swiggy,
> BookMyShow, Yodlee, OnMobile, and others.

Key facts: live device control at 10–20 ms latency streamed to the browser
(screen and audio); performance vitals (FPS, jank, CPU, memory, battery,
thermal) captured on every session by default with no SDK; Smart TV coverage
includes Samsung Tizen, LG webOS, Roku, Apple TV, Android TV, Fire TV,
Chromecast, plus cable boxes and game consoles over HDMI capture; network
capture records HTTP(S) traffic as HAR on every session. RobusTest is
sales-led and fully managed (hardware, software, and support delivered);
the same team's self-serve SaaS alternative is DeviceLab (devicelab.dev).

## Main pages

`)
	pageDescs := []struct{ path, title, desc string }{
		{"/", "Home", "What RobusTest is: the enterprise device lab on your premises"},
		{"/features", "Platform overview", "All seven capabilities and how they fit together"},
		{"/platform/manual-testing", "Manual testing", "Live device control from the browser: touch, GPS, shell, logs, vitals"},
		{"/platform/test-automation", "Test automation", "Appium hub, Espresso, XCUITest, Selenium grid, parallel runs, CI API"},
		{"/platform/performance-testing", "Performance testing", "No-SDK vitals on every session; build-over-build regression statistics"},
		{"/platform/tv-testing", "Smart TV & OTT testing", "Tizen, webOS, Roku, Apple TV, Android/Fire TV, consoles; model-year support matrix"},
		{"/platform/network-capture", "Network capture", "Automatic HAR capture, HTTPS inspection, gRPC/protobuf decoding, mocking"},
		{"/platform/device-lab", "Device lab operations", "Health scoring, booking, smart power control, iOS MDM, build library"},
		{"/platform/integrations", "Integrations & enterprise", "Google/Microsoft SSO, JIRA, ReportPortal, Slack, InfluxDB, CI/CD API"},
		{"/enterprise", "For enterprise teams", "Industry scenarios: streaming/OTT, banking & fintech, telecom"},
		{"/partners", "For testing services companies", "Partner model: run RobusTest labs at your site, deliver testing to clients"},
		{"/pricing", "Pricing", "Flat yearly license by device count; unlimited users and minutes; hardware included"},
		{"/security", "Security", "On-premise architecture, data ownership policy, air-gap, SSO, device certificates"},
		{"/about", "About", "Izinga Software, Hyderabad; building device labs since 2014"},
		{"/contact", "Contact", "Book a demo or request a lab quote"},
	}
	for _, p := range pageDescs {
		fmt.Fprintf(&b, "- [%s](https://robustest.com%s): %s\n", p.title, p.path, p.desc)
	}

	if docsStore != nil && docsStore.Ready() {
		b.WriteString("\n## Documentation\n\n")
		b.WriteString("Product documentation for RobusTest users:\n\n")
		section := ""
		for _, entry := range docsStore.Index() {
			if entry.Path == "" {
				continue
			}
			if entry.Section != section {
				section = entry.Section
				fmt.Fprintf(&b, "\n### %s\n\n", section)
			}
			fmt.Fprintf(&b, "- [%s](https://robustest.com/docs/%s)\n", entry.Title, entry.Path)
		}
	}

	b.WriteString(`
## Optional

- [maestro-runner](https://github.com/devicelab-dev/maestro-runner): open-source Maestro alternative by the same team (Apache 2.0)
- [DeviceLab](https://devicelab.dev): the team's self-serve SaaS device lab (software only, you run it yourself)
`)

	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.Header("Cache-Control", "public, max-age=3600")
	c.String(http.StatusOK, b.String())
}
