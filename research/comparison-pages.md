# Comparison-page research (RobusTest vs X / alternatives)

> ## ⚠️ CORRECTIONS FROM CODE INVENTORY (2026-07-16) — READ BEFORE WRITING ANY PAGE
>
> A capability inventory of the `tv`, `AppVitals`, and `nerve` repos contradicts
> three RobusTest-side claims this document was built on. Competitor facts below are
> unaffected; **our own** claims need fixing. Per the standing rule that capability
> facts come from code, not marketing:
>
> 1. ~~**Android TV is NOT backed by code.**~~ **WITHDRAWN 2026-07-16** — product
>    confirms **Android TV and Fire TV are just Android** and ride the generic ADB
>    path, which is exactly why grep finds no dedicated code. Coverage claim stands.
>    Likewise **Tizen/webOS are live via tvnode** (wired to nerve, working for manual
>    and automation) — nerve/neuron's install-only `device/tizen`,`device/webos` stubs
>    are the **legacy** path and misled the inventory. Full matrix: `blog-topics.md` §2b.
> 2. **"Always-on" = on-by-default-per-session, NOT scheduled** (product-clarified
>    2026-07-16). `grep -ril "cron|scheduler|recurring"` returns **zero files across
>    nerve's 448 Go files** — there is no scheduler, so **never imply 24/7 monitoring**
>    (pCloudy genuinely markets that; we'd lose). What we actually do: **every manual
>    session starts full monitoring automatically — performance, app logs AND device
>    logs — with no flag and no instrumentation**; automation gets a per-run flag.
>    That IS a real difference: BrowserStack's manual flow requires clicking *"Start
>    performance profiling"*, and Sauce's Device Vitals records only "during the test
>    session".
>    → **Ship**: *"monitoring is on by default for every session — perf, app logs,
>    device logs — nothing to enable."*
>    → **Don't ship**: "always-on/continuous monitoring."
>    Plus the rarer differentiator: statistically rigorous **build-over-build
>    regression detection** (`nerve/auth/appvitals/stats.go` — Mann-Whitney U with tie
>    correction + Cliff's delta + MDE, three-filter verdict in `buildCompare.go`).
> 3. ~~**Do not claim air-gapped deployment.**~~ **WITHDRAWN 2026-07-16** — product
>    confirms **air-gapped deployments exist in the field**, and the code supports it:
>    **SMTP is a first-class email option** (`util/email/email.go` — `gomail.NewDialer`
>    with customer host/port/creds; Sendgrid is merely the default), licence checks are
>    a local Mongo lookup with **no callback**, and GridFS keeps artifacts in-network.
>    The earlier "no SMTP path" finding was an inventory error. **We can counter-claim
>    air-gap against pCloudy.**
>
> Also: **Roku is real but vendored** — `nerve/platform/roku/` is wired (4,699 LOC)
> but its ECP client is Roku, Inc.'s Apache-2.0 code. "We support Roku" is fine;
> "we built a Roku ECP client" is not.

Deep research run 2026-07-14 (two verified workflow rounds + targeted gap-fill agents).
All competitor facts were live-fetched and adversarially verified on 2026-07-14 unless
noted. **Re-verify vendor pricing/claims shortly before publishing — pages change.**
Verification votes shown as e.g. "3–0" (three independent verifiers, none refuted).

---

## 1. Legal ground rules (verified 3–0, primary sources)

**United States — nominative fair use** permits naming competitors in comparative
advertising. Three-prong test (*New Kids on the Block v. News America Publishing*,
971 F.2d 302, 9th Cir. 1992):

1. **Necessity** — product not identifiable without using the mark
2. **Minimal use** — only as much of the mark as reasonably necessary
3. **No implied sponsorship/endorsement**

- Word mark only, never logos/distinctive lettering ("may compare to Coca-Cola or
  Coke, but not use Coca-Cola's distinctive lettering"). 2024 cases (*Xfinity Mobile
  v. Globalgurutech*, *Yelp v. ReviewVio*) show defenses failing on logo use and
  missing disclaimers.
- Circuit split exists (3rd Circuit treats it as a factor, not a replacement test) —
  plain-text names + disclaimer is the prudent rule everywhere.
- Sources: INTA fair-use fact sheet (inta.org), Dykema comparative-advertising primer
  (dykema.com PDF).

**India — Trade Marks Act, 1999**:

- **§30(1)** permits using a competitor's registered mark to identify their
  goods/services if in accordance with "honest practices" and not taking unfair
  advantage / harming the mark's repute (*Havells v. Amritanshu*, Del HC 2015;
  *Horlicks v. Heinz*, Del HC 2019).
- **§29(8)** is the boundary: unfair advantage + dishonest practices, damage to
  distinctive character, or harm to reputation = infringement.
- **Puffery vs disparagement** (*Reckitt & Colman v. M.P. Ramchandran*, Cal HC 1999,
  restated in *Horlicks v. Zydus* 2020): you may claim your product is better/best,
  but may **not say a competitor's product is bad — even if true**. Later benches
  narrowed the puffery latitude: superiority claims need a reasonable factual basis.
- ASCI code adds a self-regulatory layer: transparent, fact-based, not misleading.
- Trademark infringement in India carries criminal exposure too (§103: up to 3 years;
  §104: fines) — disparaging pages are not merely a civil risk.
- Sources: Indian Kanoon statute text, Anand & Anand and India Law Offices analyses.

**EU note** (if pages target EU buyers): honest commercial practices + no impression
of commercial connection.

**Operating rules for our pages** (Dykema primer, verbatim guidance):
- Competitor names in plain text; **no competitor logos** anywhere.
- "Not affiliated with or endorsed by X" disclaimer near brand references.
- **Verify every competitor fact against the competitor's own published materials and
  retain supporting documentation** (screenshot/archive each cited page with date).
- Assert RobusTest strengths; never rubbish rivals. Attribute contestable vendor
  claims as quotes ("HeadSpin describes its on-prem setup as 'the industry's only…'")
  rather than calling them false.
- Date-stamp facts ("as of July 2026") — required for honesty and it helps AI
  citability (freshness).

---

## 2. Verified competitor facts (citable, as of 2026-07-14)

### BrowserStack
- **App Automate (real-device automation) pricing** [3–0]: "Device Cloud" $199/mo
  annual for 1 parallel ($399/2, $999/5; monthly billing $249/$499/$1249). "Device
  Cloud Pro" $249/mo/1 parallel ($499/2, $1249/5). >5 parallels = contact sales.
  Source: browserstack.com/pricing (prices confirmed in raw HTML data attributes).
- **App Live (manual) pricing** [3–0]: Team $150/mo for 5 users, Team Pro $249/mo for
  5 users, Individual $39/mo. Source: browserstack.com/pricing?product=app-live.
- **Smart TV testing is alpha, request-access only** [3–0]: docs verbatim "Access to
  the smart TV alpha release on your BrowserStack account. To request access, contact
  support." Exactly 4 devices: Fire TV Stick 4K (Android 7.1), Nvidia Shield TV Pro
  2019 (Android 11), Apple TV 4K (tvOS 26), Roku TV (RokuOS v15+).
  **No Samsung Tizen, no LG webOS** anywhere in product page, docs, or the full App
  Automate device list [3–0 across three primary sources]. Roku: network logs/local
  testing need support enablement. Phrase as "alpha / early access" — their marketing
  page says "Available Now", the docs contradict it.
  Sources: browserstack.com/test-on-smart-tv, /docs/app-automate/smart-tv/getting-started,
  /list-of-browsers-and-platforms/app_automate.
- **Capterra complaint themes** [3–0]: overall 4.6/5 across 775 reviews (must concede
  this — sentiment is strongly positive). Most-cited negative: **pricing escalation
  with parallel sessions** ("pricing is the one honest gripe", costs "climbs fast").
  Also: interactive lag ("big lag between my mouse click & the remote device
  responding" — Yash D., Apr 2026), 15–20s session startup (Sejal Z., May 2026),
  peak-hour queues for popular devices. Praised strengths to concede: device/browser
  coverage, CI/CD integrations (Selenium/Playwright/Appium).
  Source: capterra.com/p/162900/BrowserStack/reviews/ ("most cited" reflects sampled
  pages).
- **Performance testing: per-session profiling only** (agent-verified 2026-07-14):
  "App Performance" is a profiling layer on App Live/App Automate sessions — manual
  flow is explicit start/stop ("Click Start performance profiling… Stop profiling"),
  automated flow is the `appProfiling: true` capability captured per Appium session.
  No scheduler, no always-on monitoring documented. Caveat: marketing claims
  "auto-detect performance regressions", so CI-scheduled suites can approximate
  build-over-build trends — phrase as "profiling tied to test sessions; no built-in
  scheduled/always-on monitoring." Sources: browserstack.com/app-performance,
  /docs/app-live/app-performance-testing.
- **Private devices: yes, but BrowserStack-hosted — no on-prem** (agent-verified
  2026-07-14): "Custom Device Lab" (formerly "Private Devices", launched Mar 2025,
  renamed by mid-2026) is Enterprise-only — "devices exclusively allocated to your
  organization", "reserved on our cloud", dedicated host machines available.
  Hosting: BrowserStack's own data centers. Historical wrinkle: the Nov 2025
  /private-devices page advertised "deployed on-premises or in the cloud" (Wayback
  20251113212615), but the current pages contain zero occurrences of "premise" —
  the on-prem wording was removed. **Safe claim**: "BrowserStack's dedicated-device
  offering is hosted exclusively in BrowserStack's own data centers — its current
  pages advertise no deployment on customer premises." **Unsafe**: "BrowserStack has
  no private/dedicated option" (false) or "never offered on-prem" (their own page
  said otherwise in Nov 2025; sales-negotiated options can't be ruled out).
  BrowserStack Local is a network tunnel to their cloud devices, not on-prem.
  Sources: browserstack.com/custom-device-lab (+ docs/custom-device-lab/overview),
  Mar 2025 press release, web.archive.org snapshots Jun+Nov 2025.

### Sauce Labs
- **Pricing** [3–0]: Real Device Cloud $199/mo annual ($249 monthly) per parallel,
  vendor-stated "unlimited users / unlimited testing minutes" (flag as vendor-stated).
  Virtual Device Cloud $149/mo annual ($199 monthly). Source: saucelabs.com/pricing.
  (A claimed $39/mo live plan was **refuted 0–3** — do not cite.)
- **"Private Devices" = Sauce-hosted dedicated pool, not on-prem** [3–0]: Enterprise
  contact-sales only; product page: "a dedicated pool of real devices … only
  accessible to members of your organization"; docs indicate AWS-region hosting. No
  customer-premises option in published portfolio (hosting-location wording is
  partially inferred — phrase carefully: "hosted by Sauce Labs; no on-premises option
  advertised").
- **No smart TV / OTT device testing** [3–0]: platform page lists desktop browsers +
  "7500+ iOS and Android real devices" only; no Tizen/webOS/Android TV/Roku/Fire
  TV/Apple TV anywhere. (Footnote: TestFairy app-distribution changelog mentions tvOS
  IPA distribution — distribution, not testing.)
  Sources: saucelabs.com/platform/supported-browsers-devices,
  /products/private-devices-real-device-cloud.

- **Performance testing: per-session only** (agent-verified 2026-07-14): "Device
  Vitals" records app performance "during the test session" (Live and Automated);
  Sauce Performance covers web page-load only, not native apps on devices. No
  continuous/scheduled on-device monitoring documented. Source:
  docs.saucelabs.com/mobile-apps/features/mobile-app-diagnostics/device-vitals/.
- **Capterra complaint themes** (agent-verified verbatim 2026-07-14): overall
  **4.4/5 but only 32 reviews** — disclose the thin sample when comparing ratings.
  Cost at scale: "It becomes costly when multiple instances need to be executed for
  automation purposes" (Analia F., Oct 2024). Sluggish real devices: "insanely slow
  and make for a very jittery experience" (anonymous verified reviewer, Feb 2021 —
  use only if anonymous quotes acceptable). Support: "Do not provide service which
  was paid for, even on yearly enterprise subscription" (Maria R., Jul 2025).
  Strengths to concede: breadth of real devices, real+virtual in one cloud.
  Source: capterra.com/p/149089/Sauce-Labs/reviews/ (note: product ID 149089).

### TestMu AI (formerly LambdaTest)
- **Rebrand** [3–0]: lambdatest.com/pricing 301-redirects to testmuai.com/pricing/;
  page branded "TestMu AI (Formerly LambdaTest)", announced 2026-01-12. Use both
  names on our pages (readers still search "LambdaTest").
- **Real-device pricing** [3–0]: live/manual "Real Device Plus Live" $39/mo;
  automation "Real Device Plus Automation Cloud" $199/mo — per parallel, billed
  annually, session-based not per-user. Lower tiers: Virtual Live $15, Web Automation
  $29 (Linux)/$79 (desktop), Virtual Automation $139. Vendor-stated "10,000+ real
  Android and iOS devices". Source: testmuai.com/pricing/ (confirmed in JSON-LD).
- **Smart TV testing: documented but gated** (agent-verified 2026-07-14): docs pages
  exist for Fire TV (1 model: Fire TV Stick, OS 7), Roku (Roku Ultra/Express, OS 11),
  Apple TV (Apple TV/4K, tvOS 18), Android TV (Nvidia Shield only, Android 11) on
  real devices — but **every TV platform requires contacting support to enable**, and
  Android TV verbatim "is available only with Private Device Plan". The dedicated
  smart-TV product page 404s after the rebrand, and TV is absent from their headline
  "10,000+ devices" real-device-cloud page. **No Samsung Tizen or LG webOS.**
  Sources: testmuai.com/support/docs/appium-firetv/ (-rokutv, -appletv, -androidtv),
  testmuai.com/real-device-cloud/.
- **Performance testing: per-session only** (agent-verified 2026-07-14): "App
  Performance Analytics" / "App Profiling Insights" are per-session Appium profiling
  ("review per-session metrics"); their own docs recommend external RUM for "ongoing
  performance optimization" — no scheduled monitoring product. Source:
  testmuai.com/support/docs/insights-app-profiling/.
- **Capterra complaint themes** (agent-verified verbatim 2026-07-14): overall 4.6/5
  across 546 reviews (Capterra listing renamed "TestMu AI", same /LambdaTest/ URL).
  Recurring theme is **peak-hour degradation**: "some browsers or devices can be a
  bit slow to start during peak times" (Abu S., Dec 2025); "Test execution could get
  a bit slow during peak hours" (Dasari B., Apr 2025); "Sometimes there is a waiting
  time for popular real devices during peak hours" (Saumil D., Jan 2026). Device
  availability: "real device availability has been inconsistent, especially for new
  iPhones and Ultra devices" and "Devices appear selectable but cannot be launched on
  paid plan when needed" (both Eshani N., Feb 2026, a 2-star review). Pricing:
  "By the time we got to procuring, it had already disappeared and now the base price
  is higher" (Rohan V., May 2026); "The pricing felt a bit steep, especially for
  smaller teams" (Dasari B., Apr 2025 — inside a 5-star review). Support complaints
  exist (Rohan V.) but support is more often praised — frame honestly. Most
  complaints sit inside otherwise 4–5-star reviews; disclose that.
  Sources: capterra.com/p/170732/LambdaTest/reviews/ (pages 1–2).

### AWS Device Farm
- **Pricing** [3–0]: $0.17/device-minute pay-as-you-go (first 1,000 min free is a
  **limited-time promo**; standard trial 250 min — footnote this). Unmetered:
  $250/device-slot/month, slots per usage type (automation vs remote access) and per
  family (Android/iOS). Source: aws.amazon.com/device-farm/pricing/.
- **"Private Devices" are in Amazon's data center, not yours** [3–0/2–1]: from
  $200/mo, deployed "in an Amazon data center… exclusive to your AWS account",
  **us-west-2 only**, setup via contact with the Device Farm team. No
  customer-premises option. Source: AWS developer guide, working-with-private-devices.
- **No TV devices at all** (agent-verified 2026-07-14): docs device-support page and
  FAQ cover "Android and iOS devices" / "mobile and tablet devices" only; "Fire OS"
  in their fleet description means Fire *tablets*, not Fire TV. No Fire TV/Android
  TV/Apple TV/Roku/Tizen/webOS anywhere in docs. Caveat: authoritative device list is
  console-gated — absence established from developer guide + FAQ.
  Sources: docs.aws.amazon.com/devicefarm/latest/developerguide/devices.html,
  aws.amazon.com/device-farm/faqs/.

### Firebase Test Lab
- **Pricing/quotas** [3–0]: Spark (free): 15 runs/day total (10 virtual + 5 physical;
  60/30 min daily caps). Blaze: same free daily allotment, then $5/hr physical,
  $1/hr virtual, per-minute rounded up. Source:
  firebase.google.com/docs/test-lab/usage-quotas-pricing (docs updated 2026-07-10).
- Positioning note: it's a CI test-execution service, not an interactive device
  cloud; no manual live testing.
- **TV: emulator-only and deprecated** (agent-verified 2026-07-14): the only TV
  entries in the device catalog are two Google TV *emulators*, both in the
  deprecated table ("Planned Removal Date 2025-05-31", successor "available soon").
  Never physical TV hardware; no Fire TV/Roku/Apple TV/Tizen/webOS.
  Source: firebase.google.com/docs/test-lab/android/available-testing-devices.

### Kobiton
- **Self-serve pricing** [3–0]: Startup "starting at $83/mo" (annual-billing
  asterisk; 500 device-min/mo), Accelerate $399/mo (3,000 min), Scale $9,000/yr
  (7,500 min) — all **public cloud only**, metered by device minutes.
- **On-prem only via Enterprise contact-sales** [3–0]: "Dedicated devices & fully
  on-prem devices" listed solely under Enterprise ("Let's Talk"). Perf/accessibility
  validations "Limited" in lower tiers. Source: kobiton.com/pricing/.

### pCloudy
- **No published base-plan prices** [2–1, medium confidence]: plans structured by
  parallels (1–15+, quote above that), 20% annual discount; only dollar figures on
  the page are add-ons — App Performance Experience $239/mo, Test Automation Agent
  $200/mo, Visual Testing $80/mo, Accessibility $23/mo. Source:
  pcloudy.com/pricing-packages/.
- **The only competitor in this set with true customer-premises on-prem**
  (vendor-stated) [3–0]: "Lab in a Box" appliance — "Run the full Pcloudy real device
  cloud inside your own data center… air-gapped builds, your devices, your network";
  "No cloud relay. No external servers." Plus single-tenant hosted Private Cloud
  (SOC 2 Type II, banking/fintech positioning), quote-only. Attribute as
  vendor-stated. Sources: pcloudy.com/lab-in-a-box-on-premise-testing/,
  /private-device-cloud/. **This is the competitor to take most seriously on the
  on-prem story.**
- **Smart TV: marketing claim only, no verifiable devices** (agent-verified
  2026-07-14): media/entertainment page claims "Validate apps on popular streaming
  hardware including Roku, Apple TV, Chromecast, and Fire TV devices", but the
  public device list (pcloudy.com/list-of-devices/) contains **only
  Android/iOS/Tablets — zero TV devices** — and docs.pcloudy.com has no TV how-to,
  model, or OS doc. Flag as vendor-stated with no public inventory behind it.
- **Continuous monitoring EXISTS — don't blanket-claim differentiation**
  (agent-verified 2026-07-14): "Digital Experience Monitoring" page markets "24/7 AI
  Synthetic Monitoring" — "Continuously test and monitor user flows on real devices"
  with performance trends. The $239/mo "App Performance Experience" add-on may or
  may not be this product (ambiguous); core perf docs are session-based; scheduling
  mechanics undocumented. Differentiate on specifics (device-metric trends per
  build, scheduling control, data ownership), not existence.
  Sources: pcloudy.com/synthetic-monitoring/, /pricing/.

### HeadSpin
- **Performance depth (closest competitor here)** [3–0]: perf platform spans
  Android/iOS/tablets/desktop/OTT/Smart TVs; vendor-stated "130+ KPIs across AV QoE,
  App, Device & Network Performance" incl. VMAF/UVQ; AI issue detection; devices in
  "50+ global locations". Attribute KPI count as vendor-stated.
- **Contestable superlative** [3–0]: markets "the industry's only enterprise-grade
  on-premise setup" — RobusTest, Kobiton and pCloudy all offer on-prem. Quote and
  attribute; don't call it false. Source: headspin.io/solutions/performance-testing.
- **TV OS coverage: BROAD — do not claim a HeadSpin TV gap** (agent grep-verified
  against raw HTML 2026-07-14; the earlier 0–3 refutation was itself wrong):
  smart-tv-testing page lists "Samsung (Tizen), LG (WebOS), and Vizio" and (FAQ)
  "LG webOS and Android TV"; global-device-infrastructure page adds "Roku, Nvidia
  Shield, Apple TV, Chromecast, and Fire TV Stick". Beyond marketing, HeadSpin's
  GitHub org actively maintains open-source Appium drivers for **Roku, Samsung Tizen
  TV, and LG webOS** (all pushed within a week of 2026-07-14) — engineering-grade
  evidence of real support. Their public docs device list is login-gated, so exact
  device inventory is not publicly verifiable. PlayStation/Xbox absent everywhere.
  HeadSpin is the one competitor that genuinely matches RobusTest's TV-platform
  breadth — position on deployment/data-ownership/price instead of TV coverage.
  Sources: headspin.io/solutions/smart-tv-testing, /global-device-infrastructure,
  /solutions/av-testing, github.com/headspinio (appium-roku-driver,
  appium-tizen-tv-driver, appium-lg-webos-driver).

### Deployment-model summary (the on-prem battlefield)
| Vendor | On-prem (customer premises) | Private/dedicated (vendor-hosted) |
|---|---|---|
| BrowserStack | none advertised today (on-prem wording removed after Nov 2025) | Custom Device Lab, Enterprise-only, BrowserStack DCs |
| Sauce Labs | none advertised | Enterprise contact-sales (Sauce/AWS-hosted) |
| TestMu AI | none found | "Private Device Plan" exists (referenced in TV docs, contact sales) |
| AWS Device Farm | none | $200+/mo, Amazon DC, us-west-2 only |
| Firebase Test Lab | none | none |
| Kobiton | Enterprise contact-sales | Enterprise contact-sales |
| pCloudy | **yes — "Lab in a Box"** (vendor-stated, incl. air-gapped) | single-tenant private cloud, quote-only |
| HeadSpin | claims on-prem ("industry's only" — contestable) | dedicated cloud devices |

---

## 3. Comparison-page craft (UNVERIFIED — blog/practitioner sources; treat as informed opinion)

These claims came from marketing blogs (backstageseo.com, getpassionfruit.com,
growthner.com, poweredbysearch.com) and were extracted but fell below the
verification budget. Directionally consistent with each other and with our verified
GEO research (stats/quotes/citations lift, structured content):

**Format & scope**
- One dedicated page per competitor (`/compare/robustest-vs-x`), starting with top ~5
  competitors; roundup "alternatives" pages are complementary (list yourself first +
  3–5 legitimate alternatives). Libraries of 20–50 pages claimed to compound best.
- Both formats have best-in-class examples: per-competitor (Duda vs Squarespace,
  ActiveCampaign vs Mailchimp, Chili Piper vs Calendly, Avoma vs Gong) and
  multi-competitor (Monday.com tabbed page, Unbounce single table).

**Page anatomy (recurring pattern across sources)**
1. "RobusTest vs X" title, /compare/ URL
2. Summary comparison table **above the fold** (one source claims 44.2% of LLM
   references come from the first 30% of a document — unverified)
3. Feature table: 8–12 rows, checkmarks **with contextual notes**, not bare ✓/✗
4. 3–5 qualitative differentiator sections (self-contained H2/H3s an LLM can quote
   in isolation)
5. **"When X is the better choice" section** — honest concession; self-qualifies
   prospects (Chili Piper's "If you just need basic scheduling, Calendly is probably
   enough")
6. Pricing section with real numbers + citations
7. Social proof (named customers — we have JioHotstar, Airtel, Swiggy, …)
8. FAQ (5–7 genuinely asked questions)
9. High-contrast CTA (trial/demo/migration)

**Tone & performance claims (all unverified)**
- Honest comparison > attack page: converts better, and neutral tone allegedly
  increases LLM citation ("LLMs are designed to avoid biased sources"). Attack pages
  also carry the India disparagement risk (§1 — that part IS verified).
- Claimed conversion benchmarks: comparison/BOFU pages ~5–10% visitor→lead vs 1–2%
  general organic; one practitioner reports 1,000–2,000 visits/mo per mature page.
- Claimed AI-citation data (PresenceAI, unverified): comparison matrices 61% average
  citation rate; structured comparisons +47% vs prose; 65% of AI-bot crawls target
  content updated in last 12 months → keep pages fresh and date-stamped.
- Every claim backed by a data point/named source/link — this one IS verified as our
  GEO strategy (KDD 2024) and as legal best practice (Dykema).
- One practitioner (50+ comparison pages): legal blowback rare — one C&D, one
  correction request. Anecdote, not data.

---

## 4. Gap-fill status (round-3 agents, all completed 2026-07-14)

- [x] Smart TV support: TestMu AI (documented, gated), AWS (none), Firebase
      (deprecated emulators only), pCloudy (marketing only) — folded into §2
- [x] HeadSpin's actual current TV-platform list — broad, real; folded into §2
- [x] Continuous vs per-session performance monitoring — BrowserStack/Sauce/TestMu
      per-session only; pCloudy has a 24/7 synthetic-monitoring product; folded into §2
- [x] Capterra complaint quotes: Sauce Labs, TestMu AI/LambdaTest — done, folded into §2
- [x] BrowserStack private/dedicated/on-prem posture — Custom Device Lab exists,
      BrowserStack-hosted only; folded into §2

Remaining before publishing any page: re-verify every cited vendor page on the day
of publication (screenshot/archive it), and eyeball the load-bearing pCloudy
synthetic-monitoring and TestMu TV docs pages in a browser (WebFetch quotes were
extracted via a summarizer for those two agents).

### Smart-TV coverage summary (all verified/agent-checked 2026-07-14)
| Vendor | Real-device TV testing | Tizen | webOS | Android TV | Roku | Fire TV | Apple TV |
|---|---|---|---|---|---|---|---|
| RobusTest | yes | yes (native SDB driver) | yes (native SSAP driver) | yes (= Android/ADB) | yes (full: live screen + automation) | yes (= Android/ADB) | yes (tvOS) |
| BrowserStack | alpha, request-access | no | no | 1 device | 1 device | 1 device | 1 device |
| HeadSpin | yes (drivers on GitHub) | yes | yes | yes | yes | yes | yes |
| TestMu AI | docs exist, contact-gated | no | no | private plan only | 2 models | 1 model | 2 models |
| Sauce Labs | none | – | – | – | – | – | – |
| AWS Device Farm | none | – | – | – | – | – | – |
| Firebase Test Lab | none (deprecated emulators) | – | – | – | – | – | – |
| pCloudy | marketing claim, no public devices | – | – | – | – | – | – |
| Kobiton | not researched | | | | | | |
| **Suitest** (specialist) | yes (+ simulators) | yes | yes | yes | yes | yes | yes |
| **Device.io** (ex-TV Labs) | yes | yes | yes | yes | yes | yes | yes |
| **stb-tester** | yes (HDMI capture) | n/a — capture-based | n/a | n/a | n/a | n/a | n/a |
| **Witbe** | yes (robot-based) | yes | yes | yes | yes | yes | yes |

Suitest also covers **Xbox + PlayStation + HbbTV/Freeview Play**, which nobody else
here does (including us).

Headline citable fact, **corrected 2026-07-16 and now heavily scoped**: among the
eight **general-purpose device clouds** researched here, only HeadSpin and RobusTest
cover Samsung Tizen and LG webOS on real hardware; BrowserStack (alpha) and TestMu AI
(gated) cover subsets of the streaming-box platforms; the rest have nothing.

> ⚠️ **RETRACTED — do not publish this as "only HeadSpin and RobusTest."** This
> research set scoped to general device clouds and missed the **TV/OTT specialists**,
> who beat us on breadth. Verified 2026-07-16 — see §10 for the full profile.
> TV platform breadth is **table stakes**, not our differentiator. Position on the
> protocol layer (native SDB/SSAP drivers), deployment, and price instead. Also:
> **HeadSpin authors the Appium Roku/Tizen/webOS drivers** and is Appium's first
> Development Partner — never claim Appium-TV-tooling authority.

---

## 5. Refuted claims — do NOT publish
- ~~"HeadSpin's smart TV page lists Samsung Tizen, LG webOS, Vizio" (0–3)~~ — the
  refutation itself was WRONG; grep-verified against raw HTML 2026-07-14 that the
  page does list them. See HeadSpin section. (Lesson: 0–3 votes can still be wrong.)
- "Sauce Labs has a $39/mo live testing plan" (0–3)
- "BrowserStack offers no on-prem/private/dedicated option" as verified from the
  pricing page alone (0–3 on scoping — needs broader research, see §4)

## 6. Comparison page inventory (my synthesis, not research output)

Priority = build order. "Ready" = enough verified facts in §2 to write the page today.

### Tier A — head-to-head `/compare/robustest-vs-<x>` (8 pages)

| # | Page | Our honest angle | Must concede | Ready? |
|---|---|---|---|---|
| A1 | **vs BrowserStack** | TV is a 4-device request-access alpha, no Tizen/webOS; price climbs per parallel ($199→$999 at 5); dedicated devices live in *their* DCs; profiling is per-session | 4.6/5 across 775 reviews; widest device/browser coverage; best CI/CD integrations | yes |
| A2 | **vs HeadSpin** | No published pricing at all; docs/device list login-gated; "industry's only enterprise-grade on-premise setup" is contestable (quote it) | **Real TV parity** — Tizen/webOS/Roku drivers are open-source and active. Do NOT claim a TV gap | yes |
| A3 | **vs pCloudy** | TV/OTT is a marketing claim with zero TV devices in their public device list; no published base pricing | Genuine air-gapped on-prem ("Lab in a Box") **and** a 24/7 synthetic-monitoring product — hardest page to write honestly | yes |
| A4 | **vs Sauce Labs** | No TV testing of any kind; private devices are Sauce-hosted Enterprise-only; Device Vitals is per-session | Mature enterprise tooling; disclose their 4.4/5 is only 32 reviews — thin sample cuts both ways | yes |
| A5 | **vs LambdaTest / TestMu AI** | TV docs exist but every platform is contact-gated (Android TV needs Private Device Plan); no Tizen/webOS; smart-TV product page 404s post-rebrand; peak-hour queues in 2026 reviews | 4.6/5 across 546 reviews; cheapest entry ($39 live); title must carry both names | yes |
| A6 | **vs Kobiton** | On-prem locked behind Enterprise "Let's Talk"; device-minute metering; perf validations "Limited" below Enterprise | Transparent self-serve pricing ($83–$399) | yes |
| A7 | **vs AWS Device Farm** | No TV; private devices sit in an Amazon DC, us-west-2 only; $0.17/device-min meters unpredictably | Unbeatable if you're all-in on AWS; no vendor to add | yes |
| A8 | **vs Firebase Test Lab** | CI-only, no interactive session; TV = two *deprecated emulators*; Android/iOS only | **Free** (15 runs/day). Lead with "when Firebase is enough" — the strongest concession page we can write | yes |

### Tier B — roundups `/compare/<x>-alternatives` (us first, then 3–5 honest others)
- B1 **BrowserStack alternatives** — highest-volume query in the category
- B2 **HeadSpin alternatives** — buyers here are price/opacity-driven; we fit
- B3 **Sauce Labs alternatives**
- B4 **LambdaTest / TestMu AI alternatives** — rebrand churn creates fresh search demand
- B5 **Kobiton alternatives**
- B6 **pCloudy alternatives**
- B7 **AWS Device Farm alternatives**
- B8 **Firebase Test Lab alternatives**

### Tier C — category roundups (our strongest ground; not competitor-branded)
- C1 **Best smart TV / OTT app testing platforms (2026)** — the §4 coverage table *is*
  this page. Only HeadSpin and us cover Tizen + webOS on real hardware
- C2 **Best on-premise / self-hosted device cloud options** — real field is us,
  pCloudy, Kobiton, HeadSpin. Everyone else is out by their own docs
- C3 **Mobile device cloud pricing compared (2026)** — we have verified list prices
  for all 8 vendors; highest citability per §3 (table + numbers + sources)

### Tier D — decision/intent pages (comparison-shaped, no competitor in the title)
- D1 On-premise vs cloud device labs — what actually leaves your network
- D2 Real devices vs emulators/simulators — honest, incl. where emulators win
- D3 Build vs buy a device lab — TCO
- D4 Device-minutes vs parallel-sessions vs seats — billing models decoded
  (BrowserStack/Sauce/TestMu bill parallels; Kobiton bills minutes; AWS bills both)

### Tier E — competitor-vs-competitor (rank on queries we're not in; mention us once)
- E1 BrowserStack vs Sauce Labs · E2 BrowserStack vs LambdaTest/TestMu AI ·
  E3 BrowserStack vs HeadSpin · E4 AWS Device Farm vs Firebase Test Lab
- Lower priority, highest neutrality bar — an E page that reads as an ad fails both
  the citation goal and the India honest-practices test (§1).

---

## 10. The TV/OTT specialists — the competitive set we actually missed (verified 2026-07-16)

The eight vendors in §2 are **general device clouds**. They are not who we lose TV
deals to. These are.

### Suitest (suite.st) — broadest matrix, published pricing, zero content
- **Coverage** (homepage, verbatim): "HbbTV, Freeview Play, LG webOS, Roku, Samsung
  Tizen, Sky, Vega, VIDAA, Vizio, Xfinity, Xumo, Xbox (One, Series X/S), PlayStation
  4/5, Android (TV and mobile), Apple TV (tvOS), iPhone and iPad". **Exceeds ours** —
  consoles and HbbTV are theirs alone. Note the *public lab* is narrower than the
  *supported list*: "Android, Fire OS, LG webOS, Roku, whaleOS and Samsung Tizen".
- **Deployment — the precise wording matters**: primary model is **bring-your-own
  device** — customer's TVs, customer's site, connected via Suitest control units
  (SuitestDrive, WingBox, CandyBox, Raspberry Pi) — with a **SaaS control plane**.
  A "private device lab" line exists only as a *sales-contact offer inside a security
  FAQ*; **no documented on-prem server install, no air-gap**. So: "Suitest has no
  on-prem" is **false**; "Suitest offers documented on-premise deployment" is **also
  false**. Honest framing: **devices on-prem, control plane in cloud.**
  ✅ **Confirmed 2026-07-16: RobusTest self-hosts the FULL stack** — control plane,
  devices and storage all inside the customer's DC, air-gap capable. **This is a real
  and citable gap against Suitest** (and against everyone except pCloudy/stb-tester).
  Best single citation, verbatim from `suite.st/docs/devices/`: *"Control local and
  remote devices with equal ease… You can also benefit from Suitest public device lab
  or use Suitest-managed devices…"* — all three of their models in one sentence.
  Also on every Suitest page: *"Patent granted in the US, the EU and China"* (an
  unverified self-claim — attribute, don't repeat as fact) and *"© 2016-{YEAR}"*,
  which makes them **a decade-old incumbent** vs Device.io's 3-person 2023 startup.
- **Pricing — published**: Free €0 · Single from €49/mo · Team from €149/mo · Sweet
  €494 · Super Sweet €935 · The Sweetest €1673 · Enterprise custom. Paid tiers say
  "Unlimited devices" (consistent with BYO-device).
- **QoE — explicitly disclaimed, verbatim**: "No, we are considering to add more
  video-related features, but **QoS/QoE is not our main focus**." Video assertions are
  property-level only — "Video length, Video position, Video state, Video URL" — read
  "directly from the platform's video object bypassing all upper layers".
- **Automation**: codeless visual editor + JS snippets + Appium support
  (`the.suite.st`). `suitest-js-api` npm v4.2.3, published 2026-06-16, **31,288
  downloads/30d** — actively maintained.
- **Scale**: **103 customer logos** incl. Sky, DAZN, Crunchyroll, Rakuten TV, Globo,
  Viaplay, PBS, Telekom, Showmax, Mediaset. **Applause, Testronic and Testlio are
  Suitest *customers*** — Suitest is infrastructure beneath the services labs.
- **Content: no blog.** Docs + a Canny changelog only. **Editorial vacuum confirmed.**
- Founding date/funding **not disclosed** — do not publish.

### Device.io (⚠️ formerly TV Labs / tvlabs.ai) — cloud-only, our cleanest contrast
- **⚠️ Rebrand in progress**: device.io is live ("The world's best brands build apps
  for TV, mobile and web with Device.io"), operated by TV Labs Ltd — but tvlabs.ai is
  **also still live**, still footers "© 2026 TV Labs Ltd", with different pricing and
  logos, and never mentions Device.io. **Lead with Device.io, note the TV Labs origin.**
- **Coverage**: Tizen, webOS, Android TV, Fire TV, Google TV, Roku, Vizio SmartCast,
  iOS, tvOS, Viera (+ TiVo, Vega OS on device.io). Real devices. **No consoles/HbbTV.**
- **Deployment: cloud-only. No on-prem, no data residency, no self-host anywhere in
  docs** — "Eliminate hardware… without buying a single TV", "a fully-hosted,
  cloud-based testing solution". Enterprise tier offers "Dedicated devices (exclusive
  to your org)". Region default: New York. **This is our cleanest on-prem contrast.**
- **Security**: "TV Labs is **engaged for** SOC 2 Type 1 & Type 2 compliance" — i.e.
  **not yet certified**. Quote exactly; it reads stronger than it is.
- **Pricing**: device.io "Teams" **$300/user/mo, min 3 seats** (≈$900/mo floor),
  credit-based; Enterprise custom. Legacy tvlabs.ai page shows no dollar figures.
- **Automation — Appium-first, technically credible**: proxy at `appium.tvlabs.ai:4723`,
  documents exact driver versions (roku 2.7.2, lg-webos 0.5.2, tizen-tv 0.18.4,
  xcuitest 9.9.1, uiautomator2 4.2.4, smartcast 1.0.4). `@tvlabs/wdio-service` npm
  **23,990 downloads/30d**; officially listed by WebdriverIO.
- **Performance/QoE: marketing only** — "app start time and content load speed" appears
  in copy but `docs.tvlabs.ai/platform/performance` **404s**. Only documented perf
  feature is `tvlabs:log_network`. Treat as unsubstantiated.
- **Content: none.** device.io/blog → 404. No blog on either site.
- **Scale**: **3 people** (Baldwin/Lucia/Weustenfeld, all co-founders). Founded 2023
  and funding **unverified** (Crunchbase blocked) — do not state as fact. But
  device.io names HBO Max, ITVX, Mercado Libre, CNN, AMC, FOX, Peacock, Sky.

### Others in the real competitive set
- **stb-tester** (stb-tester.com) — HDMI capture hardware + Python API, **genuinely
  on-prem**. ⚠️ **We are NOT the only on-prem TV lab** — do not claim that.
- **Witbe** (witbe.net) — robot-based real-device testing **and** live QoE monitoring.
  Closest full-scope rival on the QoE axis.
- **TestingBot** — smart TV testing, narrower. **Bitmovin Stream Lab** — tests streams,
  not app UI.
- **HeadSpin** — ⚠️ acquired by PartnerOne (2024) after its founder's fraud conviction;
  decay signals (Jonathan Lipps departed to Sony, HeadSpin University closed). This is
  **strategic intel only — never a talking point.** Using it on a page would be
  disparagement (§1) and legally reckless. Relevance: their Appium TV driver
  stewardship may lapse; re-check in 12 months.
- **Services labs, NOT rivals** (several are Suitest customers): Resillion, Accedo,
  Qvest, Sofia Digital, Norigin Media, Applause, QualityLogic.

### 🚨 Names that would be errors if published
- **"Accepto" does not exist** — no such TV-testing company (likely a garbling of
  Accedo). Drop it.
- **Eurofins Digital Testing** — domain dead; divested 2022, **rebranded to Resillion**.
- **Testronic** — games QA only now; its digital TV arm went to Eurofins → Resillion in
  2015. Listing it as a TV rival is wrong (it's a Suitest customer).
- **Rokuality, TestTrakt** — dead (no DNS) despite still ranking in search.
- S3 Group/StormTest → Accenture (2015); Mobile Labs → Kobiton; Spirent → Keysight;
  Perfecto has **no** TV offering.
- Applause/QualityLogic pages were Cloudflare-403'd — check in a browser first.

### Cross-cutting notes
- **Cannibalization risk**: devicelab.dev already publishes `kobiton-vs-browserstack`
  and `mobile-device-cloud-pricing-2025-comparison`. A6/C3/E-tier overlap those.
  Decide per page which property owns the query before writing — don't run both.
- **Kobiton TV support was never researched** — fill before A6/C1 ship.
- Every Tier A page needs its own "when X is the better choice" section (§3) — that
  section is what makes the page citable *and* keeps us inside India's
  no-disparagement line (§1).
