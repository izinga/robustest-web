package handler

// docsRedirects maps docs pages that have been retired to the page that now
// carries their content, anchored at the section they became.
//
// These are 301s, not the catch-all bounce to /docs: a permanent redirect to
// the specific replacement passes the old URL's accumulated ranking on to the
// new one and keeps existing bookmarks and support links working. Sending them
// all to /docs instead reads to a crawler as a soft 404 and throws that away.
//
// Keys are the path under /docs/ (no leading slash), matching what the docs
// handler receives. Values are full site paths.
var docsRedirects = map[string]string{
	// Ten prerequisite install pages became sections of one setup guide.
	"installation/installandroid":          "/docs/installation/installation_guides#android-sdk",
	"installation/installjava":             "/docs/installation/installation_guides#java",
	"installation/installpython":           "/docs/installation/installation_guides#python",
	"installation/installpythonvirtualenv": "/docs/installation/installation_guides#python-virtual-environment",
	"installation/installnodejsappium":     "/docs/installation/installation_guides#nodejs-and-appium",
	"installation/installxcode":            "/docs/installation/installation_guides#xcode",
	"installation/installhomebrew":         "/docs/installation/installation_guides#homebrew-and-ios-device-tools",
	"installation/installidevicelocation":  "/docs/installation/installation_guides#idevicelocation",
	"installation/elevatedpowershell":      "/docs/installation/installation_guides#elevated-powershell",
	"installation/setup_https":             "/docs/installation/installation_guides#https-for-your-robustest-server",

	// Three Report Portal pages became one; two of them shared a table verbatim.
	"integration/pushdatatoreportportal":         "/docs/integration/integration_reportportal#conditional-sending",
	"integration/testdatacollectionreportportal": "/docs/integration/integration_reportportal#test-data-collection-entries",

	// Two hub stubs were link lists the sidebar already covered.
	"hub/hub_api":        "/docs/hub/robustesthub",
	"hub/hub_frameworks": "/docs/hub/robustesthub",

	"hub/hubseleniumrandombrowser": "/docs/hub/hubselenium#randomizing-browser-selection",
	"reference/downloads":          "/docs/integration/tunnelClientConfig#downloading-the-tunnel-client",
}
