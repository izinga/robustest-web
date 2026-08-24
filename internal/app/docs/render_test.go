package docs

import (
	"strings"
	"testing"
)

func TestDocPathFromLink(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"plain page", "admin/healthpage.md", "admin/healthpage"},
		{"leading slash", "/admin/healthpage.md", "admin/healthpage"},
		{"dot slash", "./healthpage.md", "healthpage"},
		{"parent segment collapses", "../mitm-rules-samples.md", "mitm-rules-samples"},
		{"nested parent segment", "testing/../hub/hubroku.md", "hub/hubroku"},
		{"readme is root", "README.md", ""},
		{"empty is root", "", ""},
		{"https is not a doc path", "https://developer.roku.com/a/b.md", ""},
		{"http is not a doc path", "http://example.com/x.md", ""},
		{"protocol relative is not a doc path", "//example.com/x.md", ""},
		{"mailto is not a doc path", "mailto:hello@robustest.com", ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := docPathFromLink(tc.in); got != tc.want {
				t.Errorf("docPathFromLink(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestRenderPageLinkRewriting(t *testing.T) {
	tests := []struct {
		name, src, want string
	}{
		{
			"internal link becomes docs route",
			"[a](admin/healthpage.md)",
			`href="/docs/admin/healthpage"`,
		},
		{
			"parent-relative link does not leak dot dot",
			"[a](../mitm-rules-samples.md)",
			`href="/docs/mitm-rules-samples"`,
		},
		{
			"external md link is left intact",
			"[a](https://developer.roku.com/en-gb/docs/x/javascript-library.md)",
			`href="https://developer.roku.com/en-gb/docs/x/javascript-library.md"`,
		},
		{
			"fragment is preserved",
			"[a](admin/healthpage.md#status)",
			`href="/docs/admin/healthpage#status"`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			page, err := renderPage([]byte("# T\n\n"+tc.src+"\n"), "x")
			if err != nil {
				t.Fatalf("renderPage: %v", err)
			}
			if got := string(page.Content); !strings.Contains(got, tc.want) {
				t.Errorf("rendered %q\n got: %s\nwant substring: %s", tc.src, got, tc.want)
			}
		})
	}
}

func TestMetaDescription(t *testing.T) {
	tests := []struct{ name, src, want string }{
		{
			"first paragraph wins",
			"# Live View\n\nLive View streams the device screen to your browser in real time.\n\nMore text.",
			"Live View streams the device screen to your browser in real time.",
		},
		{
			"skips code fence and heading",
			"# X\n\n```bash\nrun me\n```\n\nActual prose here.",
			"Actual prose here.",
		},
		{
			"falls back to title on link-list page",
			"# Adding New Devices\n\n- [Android](a.md)\n- [iOS](b.md)\n",
			"RobusTest documentation: Adding New Devices",
		},
		{
			"strips inline markdown and links",
			"# T\n\nUse **RobusTest Connect** and see [the guide](x.md) for setup.",
			"Use RobusTest Connect and see the guide for setup.",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := metaDescription(tc.src, titleOf(tc.src)); got != tc.want {
				t.Errorf("got  %q\nwant %q", got, tc.want)
			}
		})
	}
}

func TestMetaDescriptionLength(t *testing.T) {
	long := "# T\n\n" + strings.Repeat("word ", 100)
	got := metaDescription(long, "T")
	if len([]rune(got)) > 155 {
		t.Errorf("description too long: %d runes", len([]rune(got)))
	}
	if !strings.HasSuffix(got, "…") {
		t.Errorf("expected ellipsis on truncated description, got %q", got)
	}
}

// titleOf pulls the first h1 the way renderPage does, for test convenience.
func titleOf(src string) string {
	for _, line := range strings.Split(src, "\n") {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}
	return "Documentation"
}
