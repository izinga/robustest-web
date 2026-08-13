# Blog program research (robustest.com)

# ★ THE LIST — every topic, with its evidence (2026-07-16)

**How to check any row**: demand numbers are Google autocomplete suggestion counts —
`curl -s "https://suggestqueries.google.com/complete/search?client=firefox&q=QUERY"`.
A capped **10** = meaningful volume; **0–4** = none. Capability claims cite a file you
can grep. **Nothing here rests on my judgement alone — every row is checkable.**

> **Honest accounting**: an earlier draft of this doc listed 33 topics. Only 8 rested on
> evidence; 25 were inference of exactly the kind that produced ~14 wrong claims earlier
> in this session (all of them *negative* claims from single incomplete probes). The
> list below is what survives after measuring instead of inferring. **Two of my own
> topics were killed by the probe** (FD leaks = 0, SQLite = 2) and **the TV cluster was
> resurrected by it** (`webos developer mode` = 10, `tizen developer mode` = 10).

> ## ⚠️ PRIOR-ART AUDIT RESULT (2026-07-16) — the novelty is in FAILURE MODES, not mechanisms
>
> An adversarial audit ("assume 'nobody has this' is wrong and prove it") audited 7 of
> 21 topics. **The headline: it killed our strongest claimed differentiator.**
>
> **🛑 DEAD — Mann-Whitney U + Cliff's delta for perf regression (do NOT write as novel).**
> I called this "genuinely unmatched" and "nobody publishes a non-parametric
> significance test for mobile perf." **Both false.**
> - **Mozilla's PerfCompare ships MWU as its DEFAULT test** — verbatim: *"The default
>   statistical technique for PerfCompare is Mann-Whitney-U, recently changed from
>   Student-T"* — paired with Cliff's delta. (Migrated Mar 2026; t-test slated for removal.)
> - **Laaber/Scheuner/Leitner, EMSE 2019** use the *exact* pairing as "state-of-the-art",
>   with **our own verbatim rationale**: *"Due to the non-normality… of performance data,
>   we choose the Wilcoxon rank-sum test… we utilize Cliff's Delta as a measure of effect
>   size."* Established prior art since 2019, citing Bulej et al. 2017.
> - **MobileUPReg (arXiv 2509.16864, 2025)** — a *mobile* paper — uses it as the
>   **baseline to beat**. That closes the "but nobody did it for mobile" escape hatch.
> - Sources: firefox-source-docs PerfCompare; link.springer.com/article/10.1007/s10664-019-09681-1; arxiv.org/pdf/2509.16864.
>
> **🟡 webOS mechanisms are ALREADY PUBLISHED.** SSAP, `prisoner`@:9922, the ~1000h Dev
> Mode expiry and `ResetDevModeSession.dev` auto-renewal all exist as **working code in
> the canonical webosbrew org**, with third-party writeups on top. → The field guide
> must be re-angled: **not "how webOS dev mode works" (published) but "what breaks and
> how it fails silently" (not published).**
>
> **🟢 PARTIALLY COVERED — novelty survives in the failure modes**: the :3000/:3001
> hang, the Magic-Remote second-socket silent drop, and the distributed-CA-divergence
> framing for SCEP. **Mechanisms public; failure modes not.**
>
> **🟢 SDB wire protocol (topic 1)** — the only audited topic still leaning genuinely
> novel. But the evidence only proves *documentation neglect*, not that a correct public
> spec is absent. Don't overclaim.
>
> **⬜ 14 of 21 topics were never audited** — no verdict on the SAKEP ceiling, webOS 26
> `/tmp`, permessage-deflate, WoL, CDP-on-TV, DRM black frames, FD leaks, gfxinfo,
> `am start -W`, DTX/sysmontap, watermarking, wsp/NAT, OpenSTF limits, mitm. Treat as
> **unknown**, not novel.
>
> **THE STRATEGIC CONCLUSION — and it maps exactly onto category 6 ("what not to do"):**
> **mechanisms are documented; failure modes are not.** Nobody has written the bug.
> Stop pitching "how X works" (someone published it) and write "how X lied to us"
> (nobody did). That is where every surviving novelty in this audit sits — and it is
> the one genre a competitor cannot produce without having been burned.

## TIER B — BUSINESS / IDEA. No code. Category 1 + 7. **A 12-year operator's point of view.**

Not how-to, not tech, not derived from anything measurable. Opinion pieces only a
survivor can write. ⚠️ = collides with a devicelab.dev post; deliberate or drop.

| # | Post | The argument |
|---|---|---|
| B1 | **"Everyone went to the cloud. We didn't."** | ✅ **ARMED, but thinner than hoped.** Best case: **37signals — cloud bill from a $3.2M/yr run-rate to $1.3M/yr (2024); ~$700k of Dell hardware fully repaid out of H2-2023 savings alone** (DHH verbatim; always write *"37signals reports…"* — self-reported, gross, colo/power offsets undisclosed; the $1.3M is a 2024 snapshot, their S3 contract ended summer 2025). 🛑 **DO NOT use the "83% of enterprises are repatriating" stat** — it's a **Barclays CIO *sentiment* survey** counting CIOs who ***plan* to move *at least some* workloads**. Even a16z's Martin Casado, the person most invested in the narrative, describes it in those deflationary terms. Publishing it would be the misread. 🛑 **Sovereignty does NOT support this thesis either** — Cisco's Jan 2026 benchmark (5,200 respondents, double-blind): 81% face localization demand, but **85% call it added cost/complexity/risk and 83% want it harmonized away**. Honest thesis: *repatriation is real at the anecdote level; the mass-exodus number is a misread* — which is a **better** post than the triumphalist version |
| B2 | **"This market eats companies."** — ⭐ **the best-armed post on this list** | ✅ **A billion-dollar securities fraud is on the DOJ record inside our exact category.** DOJ's sentencing release, drawing on the plea agreement, describes HeadSpin verbatim: *"HeadSpin provided clients with software tools and access to remote devices to test mobile applications"* — founder **Manish Lachwani sentenced to 18 months, 19 Apr 2024**; mechanism was textbook ARR inflation. Bracketed by the category's datable **peak: BrowserStack's $200M Series B at $4B, 16 Jun 2021, BOND-led — still its high-water mark five years on** (no later round/valuation exists). Plus **Perfecto→Perforce closed 3 Dec 2018**. Also: OpenSTF abandoned; Rokuality/TestTrakt dead DNS; S3 Group→Accenture; Mobile Labs→Kobiton; Eurofins→Resillion; Testronic exited TV; LambdaTest→TestMu; TV Labs→Device.io. 🛑 **Kill "TPG bought Sauce Labs"** — unsupported; TPG first invested 2019 and "significantly increased" its stake 9 Mar 2021 with **no amount, valuation, stake %, or fund named**; the words buyout/majority/controlling appear nowhere. ⚠️ **Tone**: HeadSpin is *one entry in a pattern*, not the villain. Court record ≠ licence to attack (§1 India). Cite the DOJ URL (justice.gov 403s fetchers; verified via Wayback) |
| B3 | **"The requirement most testing partners can't answer."** | On-prem is appearing in procurement. **Name nobody, quote no RFP.** Suitest = control plane in their cloud; Device.io = cloud-only — both non-compliant on one line |
| B4 | **"Why service companies are stuck reselling someone else's cloud."** | 🛑 **UNARMED — do not write yet.** The research area returned **zero** surviving evidence. My "Applause/Testronic/Testlio are Suitest customers" claim was **refuted 0-3**: it's a vendor-controlled logo wall — self-reported, undated, and it does **not** show those firms *resell* Suitest rather than merely use it internally, which is the actual hypothesis. (One genuine thread: Testlio's own platform page says it aggregates devices from AWS + unnamed third parties + crowd rather than operating an owned lab — but that too was voted down.) **Either evidence it properly or write it as an opinion piece that claims no data** |
| B5 | **"'Simulators are good enough' — the argument the industry is having with itself."** | 🟡 **PARTIALLY ARMED — and honest about the seam.** Vendor-neutral backing exists: **ICSME 2024 (arXiv 2408.01810), 197 device-specific compatibility issues across 94 OSS repos — 72% OEM-caused functionality breaks, 24% OEM-only features absent from stock Android.** Plus a hostile witness: **Google's own ADPF docs** concede *"The same thermalHeadroom value may be mapped to a certain thermalStatus on one device model but a different thermalStatus on another"* — and Google ships both the Emulator and Firebase Test Lab. ⚠️ **BUT: neither source mentions emulators at all.** Both emulator conclusions are *deductions*, and **both emulator-extension claims put to a vote were refuted**. Argue only from **OEM customization being definitionally absent from a stock image** — never claim a study measured emulator miss rates. ⚠️ **And the Reddit argument is not rebutted, only bounded**: nothing found addresses **cost or parallelism**, which is the axis it actually runs on. Concede that axis outright — it's true, and conceding it is what makes the rest credible |
| B6 | **"What testing at Indian scale teaches you."** | Low-end Android, hostile networks, hundreds of millions of users. Most Western vendors have never seen this. **No customer named** |
| B7 | **"Why device labs get built, then quietly abandoned."** | The ops burden nobody budgets for. We've watched it happen for 12 years — this is why "build it yourself" usually fails, and it's an honest argument *against* part of our own pitch |
| B8 | **"Testing is a cost centre — until the app is the business."** | The exec framing. When the app *is* the product, QA stops being overhead |
| B9 | **"What happens to your unreleased app after you upload it."** ⚠️ | BrowserStack documents it plainly: video 30 days, logs 60 days. Not a gotcha — they disclose it. Most buyers have simply never done the maths |
| B10 | **"Why we install it ourselves."** | No self-serve on-prem trial. A service, not a download — and why that's a choice, not a gap |
| B11 | **"What we got wrong in 12 years."** | Only you can write it. It will out-read everything else here |

**⚠️ Collision check against devicelab.dev's 105**: B9 overlaps *"Cloud Device Labs:
The Hidden Security Risks"* / *"Test Mobile Apps Without Uploading to the Cloud"*.
Avoid entirely: *"Parallel Slot Tax"*, *"Build vs Buy Device Lab"*, *"Why Fintech Teams
Can't Use BrowserStack"*, *"Why Fortune 500 Teams Build Private Device Labs"* — all
already published there. **B1–B8, B10, B11 are clean.**

---

## TIER 0 — THE STORY. 12 years of decisions. **Nobody else can write a word of these.**

Not reverse-engineered from queries. Nobody has searched for any of it. Each one is a
decision RobusTest actually made, visible in the artifacts — categories **1 (business)**
and **7 (why we did)**. This is the blog; the rest are footnotes.

| # | Post | The evidence it rests on (all already established) |
|---|---|---|
| S1 | **"Everyone went to the cloud. We didn't."** — the 12-year on-prem bet: unfashionable for a decade, now the only thing that passes a security review. Written from conviction, not hindsight | BrowserStack/Sauce/AWS/Firebase/TestMu are **cloud-only**; BrowserStack quietly *deleted* its on-prem wording between Nov 2025 and Jul 2026 (Wayback). We built full-stack on-prem + air-gap the whole time |
| S2 | **"This market eats companies."** — a decade of the OTT/device-testing graveyard, from public record. We're still here | OpenSTF abandoned; Rokuality + TestTrakt = dead DNS; S3 Group/StormTest → Accenture (2015); Mobile Labs → Kobiton; Eurofins DT → Resillion; Testronic exited TV; HeadSpin → PartnerOne fire sale; LambdaTest → TestMu (Jan 2026); TV Labs → Device.io. **All public record — history, not disparagement (§1)** |
| S3 | **"The requirement most testing partners can't answer."** — on-prem device labs are now appearing in procurement; the pattern, **named to nobody**. ⚠️ **NEVER name the prospect or quote an RFP** — that burns the deal and breaches confidence. The market-trend framing is stronger anyway: it's a claim about the market, not gossip about a tender | Suitest = control plane in *their* cloud → non-compliant. Device.io = cloud-only → non-compliant. **The category has no name, so it has no search volume — and the demand is real anyway** |
| S4 | ~~"Both sides of India's biggest streaming merger"~~ — **CUT. The names were the story; without them there is no post.** Replacement: **"What testing at Indian scale actually teaches you"** — low-end Android, hostile networks, hundreds of millions of users, a market most Western vendors don't understand. No customer named | 12 years of Indian-scale operations. ⚠️ Confirm what's publishable before drafting — [[robustest-customers]] lists names cleared for the *website*, which is not the same as clearing a *story* about them |
| S5 | **"Why we wrote our own Tizen and webOS drivers instead of using Appium's."** | HeadSpin **authors** appium-tizen-tv/lg-webos/roku drivers and is Appium's first Development Partner. We went a layer below and implemented SDB + SSAP natively. Draft exists: `blog-native-drivers.md` |
| S6 | **"Where we stopped."** — reimplementing SDB from sdbd source, correcting the canonical-but-wrong public adb protocol doc, then hitting SAKEP over `A_ENCR` in a proprietary `_TV_REL` plugin and **publishing the wall** | `tv/internal/driver/tizenos/PROTOCOL.md`. **Publishing your own ceiling is the most credible thing on this list** |
| S7 | **"We built our own SDB and borrowed Roku's ECP."** — build vs borrow, and being honest about which is which | Native Tizen/webOS drivers (ours) vs `platform/roku/ecpClient` (Roku Inc., Apache-2.0). go-ios for iOS. **Credit is the point** |
| S8 | **"Monitoring is on by default. There's no button."** — why every manual session records perf + app logs + device logs with no flag | BrowserStack's manual flow requires clicking *"Start performance profiling"*; Sauce records "during the test session". A product philosophy, not a feature |
| S9 | ⚠️ **"We chose statistics over thresholds."** — **RE-ANGLE OR DROP.** The *method* is dead as novelty (below). Only possible survivor: the **product** decision — Mozilla runs MWU on *their own* CI; the open question is whether any device-cloud *vendor* ships a significance test to customers. **Do not assert that until checked** | `nerve/auth/appvitals/stats.go` — but see the ⚠️ prior-art block |
| S10 | **"Why there's an MDM inside our test lab."** — nobody puts a CA in a device farm; here's what forced it | `auth/scep/depot.go` — micromdm/scep ships filesystem/BoltDB depots only → per-replica CA divergence behind an LB |
| S11 | **"We install it ourselves."** — why on-prem has no self-serve trial, and why that's a choice rather than a gap | White-glove, product-confirmed. Inverts the "no Dockerfile" reading |
| S12 | **"The camera sees what the framebuffer can't."** — DRM renders black; the panel is ground truth; Suitest reads the video object's *properties* and reports what the player *believes* | Suitest verbatim: assertions are "Video length, Video position, Video state, Video URL", read "bypassing all upper layers". Drafts exist: `blog-camera-webrtc.md`, `blog-ok-not-worked.md` |
| S13 | **"What we got wrong in 12 years."** | Only you can write it — and it will out-read every other post here |

**Distribution reality**: none of these have search demand and it doesn't matter. They
travel by being forwarded — into an RFP response, a security review, a board deck, a
LinkedIn feed — and by being cited when someone asks an AI to describe the problem
rather than name the category. Judge them on that, never on sessions.

---

## TIER 1 — measured demand + our authority + no incumbent. **Write these.**

| # | Topic | Demand (checkable) | Authority (grep-able) |
|---|---|---|---|
| 1 | **How to actually measure app startup time: cold vs warm vs hot** | `app startup time` = **10** | `cold/warm/hot_start_ms`, `AppVitals/pkg/android/startup.go` |
| 2 | **Jank detection on Android: what `dumpsys gfxinfo framestats` really tells you** | `android jank` = **10** (→ *jank detection*, *janky frames*, *jankstats*) | `jank_rate`, `frozen_rate`, `frame_p95`, `slow_frames`; `pkg/android/gfxinfo.go` |
| 3 | **webOS developer mode: the parts that break your tooling** ← **DRAFT EXISTS** | `webos developer mode` = **10** (→ *login*, *account*, *user id*) | `tv/internal/driver/webos/devmode.go` — dev-mode expiry uninstalls your apps |
| 4 | **Tizen developer mode + the cert-permit flow** ← **draft partially covers** | `tizen developer mode` = **10** (→ *tv*, *apps*, *2025*) | `tv/internal/driver/tizenos/`, `docs/resign-tizen-app.md` |
| 5 | **Capturing app traffic with mitmproxy on Android** | `mitmproxy android` = **10** (→ *termux*, *apk*, *emulator*); `charles proxy android` = **10** | `nerve/auth/mitm/` (packaged with node, customer-run) |
| 6 | **Reading Android memory without the profiler UI** | `android memory profiler` = **10** (→ *memory leak*, *others*) | `pkg/android/memory.go` (`dumpsys meminfo`) |
| 7 | **Outgrowing OpenSTF/DeviceFarmer: iOS, TVs, scale** | `openstf` = **10** (→ *openstf ios*); `device farmer` = **10** (→ *docker*, *github*) | our multi-protocol control plane |
| 8 | **SCEP for a device fleet: one CA, many replicas** | `scep server` = **10** (→ *open source*, *invalid response*) | `nerve/auth/scep/depot.go` |

## TIER 2 — already written. Publish after a claim review. (`tv/content/`)
| # | Draft | Status |
|---|---|---|
| 9 | `blog-webos26-field-guide.md` | ✅ **has measured demand** (row 3) |
| 10 | `blog-native-drivers.md` (own Tizen/webOS drivers) | 🔵 inferred — but unmatched authority |
| 11 | `blog-camera-webrtc.md` (DRM black frames) | 🔵 inferred |
| 12 | `blog-ok-not-worked.md` (Magic-Remote silent drop) | 🔵 inferred |
| 13 | `blog-tv-performance.md` | 🔵 inferred — **must fix TV-type specificity (§2e)** |
| 14 | `blog-business-testing.md` | 🔵 inferred |

## TIER 3 — no measured demand. **Sales enablement / AI-citation only. Judge accordingly.**
| # | Topic | Why keep it |
|---|---|---|
| 15 | **Which device clouds can run in your data centre? (2026)** | 5 of 8 can't. All facts vendor-cited. The buyer conversation |
| 16 | **What actually leaves your network** (BrowserStack: video 30d, logs 60d, own docs) | Security-review ammunition |
| 17 | **"Did this build get slower?" Mann-Whitney U + Cliff's delta** | `nerve/auth/appvitals/stats.go` — genuinely unmatched |
| 18 | **What "on-prem" actually means** (full stack vs Suitest's devices-only vs cloud) | Taxonomy that reframes the evaluation |
| 19 | **Where your test artifacts live** (S3/GCS/**GridFS** = never leaves) | `auth/storage/fileStorage/` |
| 20 | **Your device lab is behind a firewall. Good.** (nodes dial out, no inbound ports) | `nerve/wsp/`, `tunnel/` |
| 21 | **Your MDM, your CA, air-gapped** | `auth/scep/` + product-confirmed |
| 22 | **Watermarking every frame a tester sees** | `tv/cmd/tvnode/watermark.go` (`screen watermark`=10 but generic intent) |
| 23 | **What RBI/SEBI/IRDAI actually require** — **never DPDP** | Maintained page; refresh Nov 2026 / May 2027 |
| 24 | **Five artifact formats for one "TV app"** (.zip/.ipk/.wgt/.ipa/.apk) | Real fragmentation we've lived |
| 25 | **A TV's browser is frozen at manufacture** (CDP 1.3 floor) | `tv/internal/driver/cdp/client.go` |
| 26 | **TVs can't turn themselves on** (WoL) | `tv/wol.go` (`wake on lan tv`=10 but *consumer* intent) |
| 27 | **"We install it" — white-glove on-prem** | Product-confirmed; reframes the missing installer |

## ✂️ ~~KILLED BY MEASUREMENT~~ — **VOID. Ignore this whole section.**

> **Retracted 2026-07-16.** Everything below was cut on autocomplete zeros. **Autocomplete
> shows only popular prefixes — Google does not suggest rare queries at all**, so a zero
> cannot tell "nobody searches this" from "a few hundred qualified people search it in
> phrasings too varied to autocomplete." In a niche B2B category that's *all* the demand.
> robustest.com's own Search Console, one day after launch, showed real queries arriving
> that autocomplete would never have surfaced. **Nothing here is cut. Write any of it if
> the material is good.**

| ~~Topic~~ | Status |
|---|---|
| File-descriptor leak detection | **ALIVE** — rare, real, ours |
| SQLite query performance | **ALIVE** |
| Thermal throttling | **ALIVE** — and now armed: Google's own ADPF docs concede identical `thermalHeadroom` maps to different `thermalStatus` across devices |
| Touch/input latency (dispatch vs delivery) | **ALIVE** |
| Anything titled "enterprise …" / "on-prem device lab" | **ALIVE as posts.** The only surviving lesson: those are *our* words, not the market's — a titling note, not a kill |
| Appium TV driver how-tos | Still deprioritised — but for a real reason: **HeadSpin authors those drivers**, so it's their ground, and the prior-art audit shows the *mechanisms* are published |
| 53 device-lab-ops topics | Still needs a per-topic call vs devicelab.dev's 105 — overlap permitted but must be deliberate |

**The one genuinely useful signal from that probe** (because these queries *do*
autocomplete, so the reading is valid): `smart tv testing` → **browserstack**,
`app performance testing` → **browserstack**, `device farm` → **aws**,
`flaky tests` → **playwright/selenium**, `real device testing` → **free**.
That tells you who owns the category *nouns* — useful for comparison pages. It tells
you nothing about whether your material is worth writing.

---


Deep research run 2026-07-16 (1 verified workflow round + code inventory of `tv`,
`AppVitals`, `nerve` + targeted agents). Companion to `comparison-pages.md`.

> ## 🛑 THE DEMAND ANALYSIS BELOW IS WRONG — SUPERSEDED 2026-07-16
>
> **Autocomplete cannot answer the question I asked it.** It surfaces only *popular
> prefixes*. Google does not suggest rare queries at all. So every "0 suggestions"
> reading below — `on premise device lab`, `enterprise device lab`, `tizen app
> testing`, `file descriptor leak android` — is **incapable of distinguishing "nobody
> searches this" from "a small number of highly-qualified people search this in
> phrasings too varied to ever autocomplete."** In a niche B2B category, **all** the
> real demand is the second kind. The instrument was structurally wrong, not merely
> limited.
>
> **Ground truth beat it**: robustest.com's own Search Console, one day after launch,
> shows real queries arriving — queries that would **never** have appeared in
> autocomplete. Real data from the actual property > every proxy in this document.
>
> **Consequences:**
> - **Do NOT cut any topic for "zero demand."** The FD-leak and SQLite kills below are
>   void. So are the "enterprise/on-prem has no audience" conclusions.
> - **You cannot reverse-engineer a long tail you cannot see.** The only viable
>   strategy is: write something true and specific; the specific queries find it.
> - Read the §2a tables as **"what's brand-owned/contested"** (still useful — `smart tv
>   testing browserstack` is real) and **ignore every zero.**
>
> Kept below only as a record of the error.
>
> ## ⚠️ (superseded) READ THIS FIRST — demand is largely INFERRED
>
> The workflow round's demand and practitioner-pain angles produced **zero surviving
> claims**; a follow-up agent recovered real practitioner data from primary sources
> (§2d), but **keyword-volume data was never obtained** (Ahrefs/Semrush are paywalled).
>
> **Most topics here are prioritised on competitor-content gaps, regulatory timing and
> code-verified authority — not on evidence that anyone searches for them.** Ahrefs'
> ~14bn-page study found **96.55% of pages get zero Google traffic**
> ([source](https://ahrefs.com/blog/search-traffic-study/)), and their stated first
> cause is *absence of search demand*. A page nobody searches for earns nothing at
> rank #1.
>
> **The TV clusters are a bet, and §2d proves it**: five independent sources show
> **no TV testing community exists anywhere**. That is simultaneously our best
> authority vacuum *and* evidence of no search demand. Fund TV content for citations,
> sales enablement and first-mover position — **not** for traffic. Say so before
> anyone measures it on month-3 sessions.
>
> The 2026 State of Testing Report was never successfully mined (a claim that it lacks
> relevant data was refuted 0–3 — it probably *does* hold usable evidence).

---

## 1. What we can claim authority on (from code, not marketing)

Source: Explore inventory of `tv`, `AppVitals`, `nerve`, 2026-07-16. Full detail in
the ⚠️ block of `comparison-pages.md`. Summary of what's *ours*:

**Tier 1 — nobody else could credibly write these:**
1. **Native Go SDB client for Tizen + the SAKEP ceiling** (`tv/internal/driver/tizenos/`,
   incl. `PROTOCOL.md` — ground truth from sdbd 3.0.50 source, correcting the public
   adb protocol doc, and naming where reverse-engineering hits a proprietary wall).
2. **Native Go SSAP client for webOS** + Developer-Mode SSH install path.
3. **Statistical build-over-build regression detection** (`nerve/auth/appvitals/stats.go`)
   — Mann-Whitney U with tie correction, Cliff's delta, MDE, three-filter verdict,
   cited to Mann & Whitney 1947 / Cliff 1993. ~80 LOC, no gonum. Rare in production.
4. **Mongo-backed SCEP depot / distributed CA for a device fleet** (`nerve/auth/scep/`)
   — written because micromdm/scep only ships filesystem/BoltDB depots and each
   replica behind an LB would otherwise have its own CA.
5. **Forensic per-session video watermarking, pre-encode** (`tv/cmd/tvnode/watermark.go`)
   — timestamp drawn at render time so a leaked frame carries the moment it was shown.

**Tier 2 — strong, ours, adjacent to known art:** CDP 1.3 as a deliberate
compatibility floor across an 8-year TV engine spread; black-box vs white-box as an
architectural axis (DRM renders black under CDP → camera is ground truth); graceful
metric degradation per Chromium version.

**Do NOT write (we'd be faking it):** Android TV / Fire TV / tvOS / Vizio anything;
perfetto/atrace/XCTest-metrics; native iOS protocol internals (that's `go-ios`);
"how we built Roku ECP" (that's Roku Inc.'s Apache-2.0 code); air-gapped deployment;
Kubernetes autoscaling; SAML/LDAP SSO; Jenkins/CircleCI/Azure integrations; video QoE
beyond black/freeze detection; Maestro support in the platform; **"24/7 / continuous /
scheduled" performance monitoring** (no scheduler exists in any repo — but "monitoring
on by default every session, incl. app + device logs" IS true and shippable, §2b).

---

## 2. Six drafts already exist — publish before commissioning

`tv/content/`, written 2026-07-09, ~830–1,690 words each, each with a paired
LinkedIn insight + proof post (10 total). Already in the cited/stats house style.

| Draft | Topic | Maps to authority |
|---|---|---|
| `blog-native-drivers.md` | Why we wrote our own Tizen/webOS drivers instead of using vendor SDKs; has a "What about Appium?" section | Tier 1 #1/#2 |
| `blog-webos26-field-guide.md` | webOS 26 undocumented quirks: `/tmp` execute-only, journalctl blindness, Key Server timing, first-pairing foot-gun | Tier 1 #2 |
| `blog-camera-webrtc.md` | DRM makes screenshots black → watch the panel, not the framebuffer | Tier 2 |
| `blog-ok-not-worked.md` | "It returned OK" ≠ "it worked" — the Magic-Remote second-socket silent-drop bug | Tier 1 #2 |
| `blog-tv-performance.md` | Why TV performance matters differently than mobile; what's worth measuring | Tier 2 |
| `blog-business-testing.md` | Business-level: TV testing bottleneck for streaming teams | positioning |

**Action: these are the fastest path to a live blog.** Review for the ⚠️ claim
corrections (esp. anything implying Android TV support or always-on monitoring),
then publish. Do not commission new TV posts until these ship.

---

## 2b. Cluster 0 — On-premise enterprise device lab (**THE priority cluster**)

> **Correction, 2026-07-16.** An earlier draft of this doc under-served on-prem badly.
> The error: the code inventory found (a) air-gap unevidenced and (b) on-prem
> *licensing* flagged as an open product decision — and that got over-generalised into
> "avoid on-prem." Wrong. On-prem **deployment** is real, it is the business, and it is
> our **strongest verified differentiator** — stronger than TV, where Suitest
> out-covers us on breadth (§3). Air-gap specifically stays off-limits; on-prem does not.

**Why this cluster beats the TV cluster**: the competitor facts are already verified,
already citable, and against-interest (they come from the vendors' own pages).
**Five of eight major device clouds have no customer-premises option at all:**

| Vendor | On-prem? | The citable fact |
|---|---|---|
| BrowserStack | **no** | Custom Device Lab is Enterprise-only and hosted in *their* DCs. Their /private-devices page said "deployed **on-premises** or in the cloud" in Nov 2025 (Wayback 20251113212615); today the word "premise" appears **zero** times. They walked it back. |
| Sauce Labs | **no** | Private Devices = "a dedicated pool of real devices… only accessible to members of your organization", Sauce/AWS-hosted, Enterprise contact-sales |
| AWS Device Farm | **no** | Private devices are "deployed on your behalf **in an Amazon data center**", **us-west-2 only** |
| TestMu AI | **none found** | Private Device Plan exists; no customer-premises deployment documented |
| Firebase Test Lab | **no** | Nothing — CI execution service only |
| Kobiton | Enterprise only | "Dedicated devices & fully on-prem devices" behind "Let's Talk" |
| pCloudy | **yes** | "Lab in a Box" — air-gapped, vendor-stated. The real rival. |
| HeadSpin | claims | "the industry's only enterprise-grade on-premise setup" — contestable (Kobiton/pCloudy/us) |

### Topics (demand inferred; facts verified unless marked)

| # | Topic | Basis | Status |
|---|---|---|---|
| O1 | **"Which device clouds can actually run in your data centre? (2026)"** — the table above, every cell cited to the vendor's own page | all verified | **write now** — highest-value post in the whole program |
| O2 | **"What actually leaves your network when you test on a device cloud"** — BrowserStack documents retention plainly: text logs 60 days, Automate dashboard logs 30 days, App Automate video 30 days, screenshots/logs 60 days, auto-deleted. Your unreleased app's screen recordings sit on a vendor's cloud for a month, by their own docs | verified | **write now**. Framing: *"BrowserStack retains test artifacts for 30–60 days by its own documentation"* — **never** "they hide it" (refuted 0–3; also breaches the India no-disparagement rule) |
| O3 | **"One CA, many replicas: device MDM inside someone else's data centre"** — the Mongo-backed SCEP depot; micromdm/scep ships filesystem/BoltDB depots only, so each replica behind an LB gets its own CA and devices enrolled by A are rejected by B | code-verified (`nerve/auth/scep/depot.go`) | write — this **is** an on-prem post, mis-filed earlier as generic infra |
| O4 | **"Where your test artifacts live"** — per-org pluggable `StorageEngine`: S3 / GCS / **GridFS**. GridFS mode = artifacts never leave your Mongo, no cloud dependency | code-verified (`nerve/auth/storage/fileStorage/{aws,google,mongo}.go`) | write — the honest, concrete residency story |
| O5 | **"Watermarking every frame a tester sees"** — pre-encode forensic watermarking; timestamp drawn at render time so a leaked frame carries the moment it was shown | code-verified (`tv/cmd/tvnode/watermark.go`) | write — enterprise/pre-release streaming angle |
| O6 | **"What RBI/SEBI/IRDAI actually require about test data and device location"** — regulations verified (§5). **Unblocked 2026-07-16**: product says data handling **varies by customer** — we don't control what testers load onto devices. That's the *better* angle: *"you don't know what your QA team put on that device, and neither does your cloud vendor"* — written as the reader's obligation, not our claim | verified | write — hang it on **RBI/SEBI/IRDAI, never DPDP** (§5) |
| O7 | **"Five wire protocols, one device"** — ADB (Android), usbmux/lockdown (iOS), native SDB (Tizen), native SSAP (webOS), ECP (Roku) behind one device model, one API, one reservation queue. **The "almost all devices" flagship** | code-verified across `tv` + `nerve`; **integration confirmed by product 2026-07-16** | **write now** — genuinely novel; nobody drives this spread from one self-hosted control plane |
| O8 | **"iOS is the hard part of an on-prem lab"** — why phones are easy and iPhones aren't: supervision, enrolment, code signing, cert lifetime, macOS hosts. Leads naturally into O3 | code-verified (`nerve/auth/scep/`, nanomdm) | write — the single most-asked on-prem question |
| O9 | **"What 'on-prem' actually means"** — the term hides three different things: devices on your site vs control plane in your DC vs **both**. **RobusTest = full stack — control plane, devices and storage all in the customer's DC** (product-confirmed 2026-07-16). **Suitest = your TVs, their cloud brain**; BrowserStack/Sauce/AWS/TestMu/Firebase = neither; stb-tester = genuinely on-prem but capture-only; pCloudy = full (Lab in a Box). A taxonomy post that makes the buyer ask the right question | ✅ **verified both sides** | **write now** — this is the cluster's sharpest post |
| O10 | **"TVs don't turn themselves on"** — the physical reality of a TV lab vs a phone lab: neither webOS nor Tizen can power on over its control protocol, so you need Wake-on-LAN with a MAC learned while the set was awake. Plus capture (camera vs framebuffer), remotes, HDMI | code-verified (`tv/wol.go`) | write — concrete, funny, true, unmatched |
| O11 | **"The mixed-fleet problem"** — a phone lab and a TV lab are different machines: USB vs network transport, framebuffer vs camera capture, touch vs remote keys, battery vs mains. What it takes to run both under one roof | code-verified (`tv` + `nerve`) | write |
| O12 | **"Watching the traffic without leaving the building"** — mitm network capture **packaged with the node** (product-confirmed): capture, rewrite, mapLocal, breakpoints, protobuf decode, HAR export — **run and controlled entirely by the customer**, traffic never traverses our infra. Contrast: a cloud vendor MITM-ing your app's traffic on their tin | code-verified (`nerve/auth/mitm/` — authz, breakpoint, forward, inline, preset, resolver, session, validate + `mitmproxy/`) | **write now** |
| O16 | **"Your MDM, your CA, your devices"** — iOS supervision/enrolment **runs inside nerve and works air-gapped** (product-confirmed), with our own Mongo-backed SCEP CA. The customer owns the enrolment identity — no external MDM vendor, no Apple Business Manager dependency for it. O3 is the engineering post; **this is the sovereignty post** | code-verified + product-confirmed | **write now** |
| O17 | **"We install it" — white-glove on-prem** (product-confirmed). Reframes the missing Dockerfile/Helm/systemd entirely: on-prem isn't a self-serve download, it's a delivered service. That's an *enterprise selling point*, not a packaging gap — say it deliberately rather than letting a reader infer immaturity | product-confirmed | write |
| O13 | **Build vs buy: what running your own device lab actually costs** | inferred | needs honest gaps (no installer of any kind — see reality check) |
| O14 | **"Your device lab is behind a corporate firewall. Good."** — nodes dial *out* over a WebSocket reverse proxy; nerve routes back through the socket. No inbound ports, no VPN, no DMZ. Exactly the objection every enterprise security team raises | code-verified (`nerve/wsp/`, `tunnel/`, `neuron/wsp/`) | **write now** — new from the inventory, and a real on-prem pain |
| O15 | **"No license server, no callback, no internet"** — validation is a local Mongo lookup; nothing phones home. Contrast: every SaaS device cloud dies without connectivity | code-verified (`middleware/auth/auth.go:409-478`) | write — but **do not** frame as "offline licensing" (unenforced; see reality check) |

### 🔴 On-prem reality check (deployment inventory, 2026-07-16)

A full inventory of `nerve` + `neuron` came back materially harsher than the product
framing. **The runtime story is genuinely good. The packaging, licensing, and TV
device support are not.** Publish nothing from this cluster until the two blocking
questions below are answered.

**What is REAL and worth writing (code-verified):**
- **Runs disconnected — no license server, no runtime callback.** License validation
  is a Mongo lookup (`middleware/auth/auth.go:409-478`). Nothing phones home. Rare
  and true.
- **Artifacts genuinely never leave the network.** GridFS is the *default and
  complete* path (Upload/Get/Delete all work), builds are GridFS-only by design, and
  no cloud storage exists unless an admin creates a storage Integration. "Mongo-only
  mode, zero cloud dependency for artifacts" is fully supported.
- **The Mongo-backed SCEP depot** (`auth/scep/depot.go`) — best-engineered thing in
  the estate; own your CA inside your network, no per-replica CA divergence (O3).
- **NAT-traversing device nodes** (`wsp/`, `tunnel/`) — labs behind corporate
  firewalls dial *out*; no inbound ports. Real on-prem pain, real solution. **New
  topic.**
- **Camera-based TV rig** — `PrimaryScreenSource: "raspberry_pi_camera"` + RTP over
  WebRTC + IR remote (`handler/device/device.go:145-158`). Novel. Describe it as what
  it is: a rig, not device-native capture.
- **Android and iOS support are deep and real** (`neuron/device/android/` ~250KB,
  `device/iOS/iOS.go` 164KB).

**✅ AIR-GAP IS REAL — the inventory was wrong (corrected 2026-07-16).**
Product confirms **air-gapped deployments exist in the field**, and the code backs it:
`util/email/email.go` supports **SMTP as a first-class option** —
`Config.DefaultEmailService == "smtp"` routes to a `gomail.NewDialer(Config.Mail.Host,
Port, Username, Password)` with TLS config and an anonymous-SMTP branch when no
username is set; validation at `util/housekeeping/config.go:38` demands SMTP settings
when that mode is chosen. **Sendgrid is only the default**, overridable from admin.
The inventory agent saw the hardcoded `api.sendgrid.com` inside `sendBySendgrid`
(email.go:105) and wrongly concluded "no SMTP path found".
→ **Air-gap is claimable**: no license callback, no runtime phone-home, artifacts in
GridFS, mail via the customer's own SMTP relay. Remaining cosmetic egress: email
templates reference a `storage.googleapis.com` logo and `mobile.robustest.com`
(email.go:31-33) — broken images in an air-gapped install, worth fixing.

**🛑 Claims still to strike / verify:**
1. ⬜ ~~"License gates nodes and devices"~~ — inventory says `MaxNode`/`MaxDevice` are
   never read and the key is an unsigned `GenerateRandomString(20)`. **Re-verify
   before relying on this** (see reliability note) — but if true, don't market
   "offline licensing"; a competitor reads that as unenforced.
2. ⬜ **Apple TV live screen** — rests on a `model/testSesion.go:22-24` comment saying
   the producer never starts. That comment may be as stale as the Tizen/webOS stubs
   were. Verify on a real Apple TV before publishing either way.

> **⚠️ Reliability note on the deployment inventory.** It produced **two confident
> false negatives** — "no SMTP path" (wrong; SMTP is first-class) and TV support
> (technically true of nerve's legacy stubs, but it missed the live tvnode path and so
> read as "TV is install-only"). Both were *negative* claims drawn from incomplete
> reads. Its file-existence findings (no Dockerfile/Helm/systemd) are more trustworthy
> than its absence-of-capability findings. **Spot-check anything from it before it
> reaches a page.**
**✅ RESOLVED 2026-07-16 (product confirmation) — both earlier blockers are cleared:**
1. **tvnode IS wired to nerve and working for manual + automation.** Tizen/webOS TV
   testing is real and live via the `tv` repo's native SDB/SSAP drivers.
2. **The nerve→devicelab move is a migration, not a withdrawal** — no change to the
   enterprise plan or support. On-prem content is safe to write.

⚠️ **Why the inventory got this wrong — a code-hygiene trap worth fixing.**
`neuron/device/tizen/tizen.go` and `webos.go` remain in the tree as **legacy
install-only stubs** (~38 of ~40 methods return zero; the tizen error string still says
"roku device"), and `tv/cmd/tvnode/appium_mgmt.go` labels several nerve-facing
endpoints "best-effort placeholders". A thorough read of `nerve` alone lands on the
legacy path and wrongly concludes TV is install-only. The tell: `model/testSesion.go`
says Tizen/webOS lack screen capture *"even in the **legacy path**"* — implying a
non-legacy path that a nerve-only inventory never sees. **Stale stubs beside the live
path will mislead the next reader, human or agent.**

**The device matrix (corrected 2026-07-16):**

| Platform | Status |
|---|---|
| Android phone/tablet | ✅ deep (ADB, scrcpy→WebRTC, minitouch, logcat) |
| iOS phone/tablet | ✅ deep (usbmux/go-ios, WDA, DVT→WebRTC) — needs macOS + Xcode |
| **Samsung Tizen** | ✅ **live via tvnode** — native Go SDB driver, CDP, camera/WebRTC |
| **LG webOS** | ✅ **live via tvnode** — native Go SSAP driver, CDP, camera/WebRTC |
| **Roku** | ✅ **full — live screen + automation** (product-confirmed; the "screenshot only, no WebRTC" finding was wrong) |
| "television" (dumb TV) | ✅ camera + IR rig — novel |
| **Apple TV (tvOS)** | ✅ **works — live screen + automation** (product-confirmed; the `model/testSesion.go:22-24` comment is **stale**, like the tizen/webos stubs) |
| Chromecast | ⚠️ thin (CDP websocket, no install/input) |
| **Android TV / Fire TV** | ✅ **live — they are just Android** (product-confirmed 2026-07-16); they ride the generic ADB path, which is why grep finds no dedicated code. The earlier "no Android TV" ⚠️ is **withdrawn** |

**What an on-prem customer does NOT get — know this before writing "enterprise":**
~~no installer~~ (**resolved: white-glove — we install it**; the missing
Docker/Helm/systemd is a *delivery model*, not a gap — but it does mean **no self-serve
trial** of on-prem, and each install is hands-on, which caps how fast the motion
scales); ~~no SMTP~~ (**wrong — SMTP is first-class**);
no SAML/LDAP/AD/SCIM (OAuth2 + password only — the one real enterprise gap);
**hardcoded `admin@robustest.com` and
`robustest.com` in the customer bootstrap** (`util/housekeeping/admin.go:214-345`);
a **19-devices-per-node default ceiling** whose config key isn't in the shipped
`neuron/config.yaml`; no quota enforcement; reservations that never expire; mitmproxy
requiring manual per-node Python installs; and S3/GCS artifacts that can't be fetched
or deleted through the storage API.

**Also do not claim**: Kubernetes autoscaling (`auth/job/kubernative.go` is CRUD over
a config collection, not an orchestrator).

---

## 2c. Cluster 0b — Device lab operations (the seam I nearly missed)

> ## ⚠️ devicelab.dev already covers this ground — write it anyway, for a different reader
>
> **Product decision, 2026-07-16: overlap is fine — the audiences differ.**
> devicelab.dev talks to OSS/self-hosters; robustest.com talks to enterprise buyers.
> Same topics are allowed, with different framing and depth. Accept some keyword
> cannibalisation as the cost. **So this cluster is live for robustest.com** — but
> pitch it at the enterprise reader (fleet scale, compliance, support, procurement),
> not at the hobbyist with five phones on a shelf, or we're just re-writing
> devicelab.dev in a different font.
>
> **Checked 2026-07-16.** `devicelab.dev/blog` contains **105 articles** in categories
> including Comparisons, Engineering, Tutorials, Security, Compliance, Zero Trust,
> **Cloud Exit**, **Enterprise**, **Hardware**, Performance. Among them, verbatim:
> *"How to Build a Mobile Device Lab in…"*, *"Device Lab Hardware Requirements:
> Shopping List"*, *"Build vs Buy Device Lab: Cost Analysis"*, *"Private Device Cloud
> Architecture: Zero-Trust"*, *"Why Fortune 500 Teams Build Private Device Labs"*,
> *"Why Fintech Teams Can't Use BrowserStack"*, *"Cloud Device Labs: The Hidden
> Security Risks"*, *"Managing 50+ Test Devices"*, *"Scaling Mobile Test
> Infrastructure: From 10 to…"*, *"Hybrid Device Lab Strategy: Own + Cloud"*,
> *"Device Lab Cost Calculator: Cloud vs Own"*, *"Test Mobile Apps Without Uploading
> to the Cloud"*, *"HIPAA-Compliant Mobile Testing"*, *"DeviceLab Certified Hardware:
> Hubs & Mounts"*, *"Parallel Slot Tax"*, *"OpenSTF & DeviceFarmer Alternatives"*.
>
> **Check each topic against devicelab.dev before commissioning** — nearly every
> comparison page in `comparison-pages.md` §6 is *also* already there (BrowserStack
> Alternative, Sauce Labs Alternatives, LambdaTest Alternative, TestMu AI
> Alternatives, HeadSpin Alternative, Kobiton vs BrowserStack, per-vendor pricing and
> troubleshooting). Duplicate deliberately, or not at all.
>
> **The uncontested ground is still TV.** devicelab.dev has **zero** TV posts — not
> one of 105 mentions Tizen, webOS, Roku or OTT. Neither does any competitor blog.
> Combined with the demand finding in §2d, that makes robustest.com the natural home
> for **smart TV/OTT**, **deep protocol engineering** (SDB/SSAP/CDP), **production
> statistics** (Mann-Whitney), **MDM/SCEP**, and the **enterprise** story.
>
> Note also devicelab.dev's tone is aggressive ("HeadSpin Review 2026: After the
> Fraud", "BrowserStack Pricing: The Hidden Costs"). An enterprise brand selling to
> BFSI probably should not match it — and the India no-disparagement rule
> (`comparison-pages.md` §1) bites harder on robustest.com than on an OSS blog.
>
> **Method correction that produced this cluster anyway:** Clusters 0/1/6 were built by
> asking "what can we prove in the repo?" — the right bar for comparison pages (legal
> risk), the wrong bar for a blog, where the standard is *"is it true, useful, and did
> we live it."* Keep the list below as the **devicelab.dev backlog / gap-check**, not
> as robustest.com's plan. For scale: BrowserStack's `/guide/` is **186 pages**; their
> `/blog/` is mostly product/community/newsletters.

### A. Why a lab at all (decision-stage)
1. Real devices vs emulators — what emulators **structurally** cannot catch: Widevine
   L1/DRM, thermal throttling, real modem/radio behaviour, camera/NFC/biometrics, GPU
   drivers, OEM skins
2. How many devices do you actually need? Building a coverage matrix from your own
   analytics rather than a vendor's device count
3. The long tail: why the 20th-most-popular device still matters in India
4. Device fragmentation in India — low-RAM Android, storage pressure, entry-level SoCs
5. Cloud vs on-prem: the utilisation crossover point (where owning beats renting)
6. Build vs buy: real TCO of a device lab
7. **When you should NOT run your own lab** (the honest one — earns more trust than
   the other six combined)

### B. Physical/hardware ops — *nobody writes this and everybody hits it*
8. **Battery swelling**: what happens when phones sit on charge for two years
9. Charge/discharge cycling to keep a fleet's batteries alive
10. Thermal: phones throttle in racks, and your benchmarks lie
11. USB at scale — hub topology, powered hubs, cable degradation, port budgets
12. Wi-Fi for 100 devices — AP density, channel planning, 2.4 vs 5 GHz contention
13. Racking and mounting: what actually survives
14. Screen burn-in on always-on displays
15. Device retirement: when a device leaves the fleet

### C. Android at scale
16. ADB at scale — adb server limits, offline devices, reconnect storms
17. USB vs ADB-over-TCP: the tradeoffs
18. Screen streaming: scrcpy vs minicap
19. Input injection with minitouch
20. **OEM quirks**: MIUI permission prompts, Samsung Knox, Vivo/Oppo install dialogs
21. Google Play Protect blocking your own test builds
22. Keeping devices unlocked, awake, and out of doze
23. Disabling OS auto-update — and what breaks the day you forget

### D. iOS at scale — the hard part
24. Why an iOS lab needs Macs (and how many)
25. usbmuxd/lockdown: how iOS device comms actually work
26. WebDriverAgent lifecycle management
27. **Provisioning profiles that expire at the worst possible moment**
28. Supervision + MDM enrolment for a test fleet → leads into O3/O8
29. iOS 17+: `devicectl` vs go-ios
30. "Trust this computer", 200 times
31. Screen capture on iOS via DVT

### E. Smart TV lab (ours alone)
32. **TVs can't turn themselves on** — Wake-on-LAN with a MAC learned while awake
33. Camera vs HDMI capture vs framebuffer — and why DRM forces the choice
34. IR blasters vs network remotes
35. Tizen dev mode + the cert-permit flow (Samsung-account-bound, IP-allowlisted, expiring)
36. webOS dev mode expiry **uninstalls your apps** (~1000h)
37. A TV's Chromium is frozen at manufacture — automation capability never improves

### F. Platform / architecture
38. Devices per host: what actually limits you
39. Video streaming at lab scale (WebRTC)
40. Device reservation and queueing across teams
41. Device health monitoring and self-healing
42. **State hygiene** — cleaning app data, accounts and permissions between sessions
43. Test data and account management on shared devices
44. NAT traversal: nodes dial out (→ O14)
45. Artifact storage: where videos/logs/screenshots actually go (→ O4)

### G. Security / compliance
46. Who can see the screen — pre-release protection on a shared lab (→ O5)
47. Data residency: what leaves your network (→ O2)
48. RBI/SEBI/IRDAI (→ O6, blocked)
49. Network isolation for labs
50. Wiping devices between tenants/teams

### H. Economics
51. Utilisation — the metric that decides cloud vs own
52. Device procurement strategy: what to buy, when
53. Lab ops headcount: the cost nobody budgets

**Note**: several of these (B, C, D) are *not* RobusTest-specific — they're the shared
pain of anyone running a lab. That is a feature: they're the most searched, least
vendor-written content in the space, and we have the scar tissue. They also feed
Cluster 0's product story naturally rather than by force.

---

## 2a. DEMAND — measured at last (autocomplete probe, 2026-07-16)

Three research rounds failed to produce positive demand data. The instruments mostly
don't exist for this niche — **and one that everyone reaches for is actively invalid**
(below). Google autocomplete is the one that works: a **capped 10-item list signals
meaningful volume; a short or empty list signals none.** Probed directly, reproducible.

| Query | Suggestions | Read |
|---|---|---|
| `app startup time` | **10** (→ android, firebase) | ✅ real, open demand |
| `android jank` | **10** (→ jankstats, janky frames, **jank detection**) | ✅ real, open, developer-intent |
| `flaky tests` | **10** (→ **playwright**, **selenium**) | ✅ real — but owned by Playwright/Selenium, *not* Appium |
| `real device testing` | **10** (→ **free**, **online free**, **android free**) | ⚠️ real but **price-seeking** — devicelab.dev's audience, not enterprise |
| `device farm` | **10** (→ **aws**, pricing) | ⚠️ real but **AWS-anchored** |
| `app performance testing` | **10** (→ **browserstack**, tools) | ⚠️ real but **BrowserStack-anchored** |
| `frame drops android` | 7 | 🟡 moderate |
| `android anr` | 7 — but mostly **German noise** (*anruf* = phone call) | 🟡 thin |
| `memory leak android app` | 2 | 🔴 thin |
| `appium flaky` | 2 | 🔴 thin |
| `mobile app performance monitoring` | 2 | 🔴 thin |
| **`on premise device lab`** | **0** | 🔴 zero — **but see the correction below; this phrasing is NOT the concept** |
| **`openstf`** | **10** (→ device farmer, smartphone test farm, **openstf ios**, stf) | ✅ **real — the self-hosting audience searches the TOOL** |
| **`device farmer`** | **10** (→ stf, **github**, **docker**, openstf) | ✅ **real, developer-intent** |
| **`private device cloud`** | **10** (→ **testingbot** private device cloud, private real device cloud) | ✅ real, lightly brand-anchored |
| `self hosted device farm` / `open source device farm` | 4 / 5, mostly noise (recipes, farm equipment) | 🔴 thin |
| `in house device lab` / `device lab setup` / `self hosted mobile testing` / `browserstack alternative self hosted` | **0** | 🔴 zero |
| `on premise testing` | 5, mostly **drug-testing** noise | 🔴 thin |
| `smart tv testing` | 6, top suggestion = **`smart tv testing browserstack`** (relevance 601, highest) | 🔴 tiny + **BrowserStack-owned** |
| `tizen app testing` / `webos app testing` | **0** | 🔴 zero |

### ⚠️ "Enterprise" is dead as a search word (probed 2026-07-16)
`enterprise device lab` → **0**. `enterprise device cloud` → **0**. `enterprise mobile
app testing` → **0**. `corporate device lab` → **0**. `enterprise device farm` → **0**.
`enterprise mobile testing` → 2 (one is *drug testing*). `enterprise test lab` → 2
(Microsoft 365 noise). The only "enterprise" query that caps at 10 is `mobile device
management enterprise` — and that's the **Intune/Jamf IT-admin market, not QA**.
Adjacent-sounding, wrong buyer; our MDM/SCEP capability must not chase it.

**"Enterprise device lab" is our internal language, not the market's.** Never title a
post with it.

### The pattern that actually governs this space
**Buyers don't search by segment adjective. They search by TOOL NAME or by SYMPTOM.**

| They never search | They do search |
|---|---|
| enterprise device lab / device cloud / device farm (all 0) | **`openstf` (10)**, **`device farmer` (10)** |
| on premise device lab / in house device lab / device lab setup (all 0) | **`private device cloud` (10)** |
| self hosted mobile testing / browserstack alternative self hosted (0) | **`android jank` (10)**, **`app startup time` (10)** |
| tizen app testing / webos app testing (0) | `smart tv testing **browserstack**` |

So the tension isn't "our differentiators have no audience" — **it's that the audience
searches in a vocabulary we don't use.** Where we're differentiated, demand exists but
is expressed as *tool names* (openstf/device farmer) and *symptoms* (jank, startup);
where we use category language (enterprise/on-prem/device lab), it's zero. Where the
category terms do have volume, they resolve to *"browserstack"*, *"aws"*, or *"free"*.

**→ Write in the market's words.** Not "our enterprise on-prem device lab" but
"`openstf ios`", "`device farmer docker`", "`android jank detection`".

**Three consequences:**
1. **robustest.com's blog is not primarily a search play.** Judge it on citations,
   sales enablement and being first — not month-3 sessions. Decide that now, in
   writing, or it gets killed at the first traffic review.
2. **The demand-capture play is comparison-framed.** People literally type *"smart tv
   testing browserstack"* and *"app performance testing browserstack"*. Riding that
   brand query is lawful (nominative fair use, `comparison-pages.md` §1) and it is
   where the searchers actually are. That's the bridge between the blog and §6's
   comparison pages.
3. **✅ THE EXCEPTION — and it's the big one: `app startup time` (10) and `android
   jank` (10, → "jank detection") are real, open, developer-intent, and owned by
   NOBODY.** No brand suffix, no "free". And they land exactly on our deepest verified
   authority (§2e: `cold/warm/hot_start_ms`, `jank_rate`, `frozen_rate`, `frame_p95`,
   `slow_frames` via `dumpsys gfxinfo framestats`). **This is the only place where real
   demand, no incumbent, and first-hand authority overlap.** Both prior rounds missed
   it because neither ever measured the perf cluster.

### 🛑 Do NOT use Stack Overflow tag counts as a demand instrument (verified)
Appium's apparent collapse — **1,025 new questions (2019) → 43 (2025) → 22 (trailing
12mo)** — is **~87–97% a platform-abandonment artifact**, not Appium decline:
site-wide SO went **1,752,419 → 111,984** over the same period. Appium retained 4.2%
of its 2019 volume vs 6.4% site-wide. Only **normalised share** (SO's own published
method) is valid, and **SO Tag Trends was retired in March 2026**. Publishing the raw
collapse as a demand signal would be wrong and checkable.

Related corrections to earlier notes in this doc:
- **The State of Testing Report is genuinely empty** — verified by full-text census of
  *both* the 2025 PDF (12th ed.) and 2026 web report (13th ed.): **zero** mentions of
  device, flaky, emulator, simulator, "real device", on-prem, self-host, appium,
  "smart tv", android, ios. The survey never asks. My earlier note that it "probably
  does hold usable evidence" was **wrong** — stop chasing it.
- **"Original data beats how-tos" is UNSOURCED.** The Ahrefs ChatGPT-citation study has
  **no content-format dimension at all** (its variables are retrieval channel, semantic
  relevance, URL slug, content age). Don't justify format choice with it. The KDD 2024
  GEO finding supports **in-page composition** (stats + quotes + citations) regardless
  of format — so enrich the existing drafts rather than restructure them.
- ⚠️ **Citation hygiene**: cite SO figures to the **Stack Exchange API**
  (`/2.3/questions?tagged=…&fromdate=…&filter=total`), never `/tags/…/info` (which
  doesn't emit them). And **never publish `samsung-smart-tv = 574`** as the domain's
  size — it understates ~3–4× (tizen = 1,409); it's cherry-picking and won't survive a
  reader's check. Autocomplete rank is **not** volume.

---

## 2e. Mobile performance — what AppVitals actually captures (verified 2026-07-16)

**19 Android collectors, 14 iOS collectors, 32 scored metrics with good/poor
thresholds** (`AppVitals/pkg/metrics/scoring.go`), weighted into 0–100 category scores
(performance .35 / responsiveness .25 / efficiency .25 / stability .15).

| Group | Metrics |
|---|---|
| Rendering | `fps_avg`, `jank_rate`, `frozen_rate`, `frame_p95`, `slow_frames`, `transition_avg_ms` |
| Startup | `cold_start_ms`, `warm_start_ms`, `hot_start_ms` (separately!) |
| Input latency | `touch_latency_ms`, `input_dispatch_ms`, `input_delivery_ms` (dispatch vs delivery split) |
| CPU / GPU | `cpu_avg`, `cpu_peak`, `gpu_avg` |
| Memory | `memory_avg_mb`, `memory_peak_mb` |
| Battery | `battery_per_min` |
| Stability | `crash_count`, `anr_count`, `exception_count` |
| GC | `gc_pause_total` |
| **Thermal** | `cpu_temp_avg`, `cpu_temp_max`, **`throttle_count`** |
| Network | `packet_loss`, `dns_lookup_ms`, `tcp_connect_ms` |
| **File descriptors** | `open_fds_max`, **`fd_growth_rate`** (leak detection) |
| **SQLite** | `sqlite_slow`, `sqlite_avg_ms` |

Collection: Android via `adb` + `dumpsys gfxinfo framestats` / `/proc/stat` /
`dumpsys meminfo` / `thermalservice` / `am start -W`; iOS via Instruments/DTX
(go-ios) **plus our own raw sysmontap service** exposing per-process data go-ios's
public API doesn't surface (`pkg/ios/services.go`, `sysmontap.go`).

**The rare ones — this is where the story is**: thermal **throttle count**, **FD growth
rate**, **SQLite query latency**, **ANR count**, input **dispatch vs delivery** split,
and **cold/warm/hot startup as three separate metrics**. Competitors surface CPU/memory/
FPS/battery; almost nobody surfaces fd leaks or thermal throttling.

⚠️ **On "more than anyone" — I'd not publish that claim.** HeadSpin markets **"130+
KPIs"** (vendor-stated, and it bundles AV QoE + network which inflates the count).
A metric-count race against that is unwinnable and unverifiable. **The defensible
claims are specificity and the statistics layer**: 32 *scored* metrics with published
thresholds, **plus Mann-Whitney U + Cliff's delta regression detection on top**
(`nerve/auth/appvitals/stats.go`) — nobody else publishes a non-parametric significance
test for mobile perf. *Name the metrics nobody else has; don't count them.*

---

## 2d. Practitioner demand — what the community actually says (2026-07-16)

The workflow round surfaced nothing here; a dedicated agent went to primary sources
(Reddit via old.reddit.com, HN Algolia API, Stack Overflow) and recovered real data.
**Two findings overturn earlier conclusions.**

### 🛑 Drop the "BrowserStack is expensive" angle — for engineer-facing content
A term-sweep of the QA corpus for cost/pricing/`$`/quota/seat found **near-zero
organic hits**. Practitioners attack **reliability, speed and queue latency** — not
invoices: *"The last thing you want is flakiness in… the underlying infrastructure."*

This does **not** contradict the Capterra finding in `comparison-pages.md` (pricing IS
the top Capterra complaint) — it splits by audience:
- **Capterra reviewers = buyers** → cost is the pain → use it on **comparison pages**.
- **Reddit/HN = engineers** → reliability is the pain → use it in the **blog**.
Leading a blog post with "BrowserStack is expensive" talks past the reader.

### ✅ "No TV testing community exists" — the best-evidenced claim in the project
**Three independent sources agree**: zero relevant 2025-26 TV/OTT threads across the QA
subreddits; zero on Hacker News; the `samsung-smart-tv` Stack Overflow tag had **zero
questions in 2025**. Combined with **zero TV posts across 105 devicelab.dev articles**
and **no blog at all from Suitest or Device.io**, the TV authority vacuum is now
established from five directions.

⚠️ **But read it honestly — it cuts both ways.** No community means **no search demand
to capture**. Per Ahrefs, the first cause of a zero-traffic page is absence of demand.
TV content is a bet on **starting a conversation, not joining one** — justified by
citations, sales enablement and being first, **not** by traffic forecasts. Say that out
loud when the program is funded, so nobody judges it on month-3 sessions.

### Verified threads worth writing against
- **[Since when did simulator testing become "good enough"?](https://old.reddit.com/r/iOSProgramming/comments/1szt079/since_when_did_simulator_testing_become_good/)**
  — 2026-04-30, 28 pts, 37 comments. Verbatim OP: *"So did something change? Are
  simulators now good enough for most apps, or are people just accepting this because
  it's easier and cheaper?"* Commenter, verbatim: *"sim is parallelisable, scriptable,
  screenshot-friendly, and roughly free per run. **Real devices are none of those.**"*
  → The best available proof the real-device premise is **live and contested in 2026**.
  Answers there are good, so this is not an unanswered gap — but it's the sharpest
  framing of the objection our whole product must beat. Topic A1 should answer *that
  comment*, by name.
- *"We want an open source alternative to Appium… like Playwright did for Selenium"*
  — [2026-06-01](https://old.reddit.com/r/softwaretesting/comments/1ttmnt3/what_appium_alternative_are_teams_moving_to_in/),
  9 pts. **Mostly vendor spam**, and a skeptic calls the selector-churn complaint
  *"hyperbole… a dev governance issue."* Do not overclaim it.

### Top-ranked unanswered gaps (agent's ranking, TV-side)
1. Samsung's own disclaimers on Tizen behaviour
2. Black TV screenshots (→ our `blog-camera-webrtc.md` draft answers this exactly)
3. Physical lab constraints

---

## 3. Cluster 1 — Smart TV / OTT (strong, but narrower than it looked)

**Verified [3–0]: there IS a content vacuum among general device clouds.**
BrowserStack's smart TV product is an **alpha, request-access only**, covering exactly
four device/OS combos (Fire TV Stick 4K, Nvidia Shield TV Pro, Apple TV 4K, Roku TV) —
**no Tizen, no webOS** on any BrowserStack-controlled page. This is against-interest
evidence: vendors don't label their own product alpha without cause.
Sources: browserstack.com/docs/app-automate/smart-tv/getting-started, /test-on-smart-tv.
**TIME-SENSITIVE — no page timestamp; the whole cluster argument turns on this. Re-verify before publishing.**

**Verified [3–0]: but it is NOT a product vacuum — and we missed the whole specialist
field.** **Suitest** out-covers us (adds Xbox, PlayStation, HbbTV/Freeview Play),
publishes its pricing (€0–€1673/mo), has **103 customer logos** (Sky, DAZN,
Crunchyroll, Rakuten TV, Globo, Viaplay) and an npm API doing 31k downloads/month.
**Device.io** (⚠️ the rebrand of TV Labs) is Appium-first and names HBO Max, ITVX,
CNN, FOX, Peacock. **stb-tester** is genuinely on-prem. **Witbe** does real-device
QoE. Full profile: `comparison-pages.md` §10.

**But neither Suitest nor Device.io publishes a blog at all** (device.io/blog → 404;
Suitest has docs + a changelog only). → **The gap is editorial, not product.** That is
the single best argument for this blog program existing.

⚠️ **This invalidates the comparison research's headline claim** ("only HeadSpin and
RobusTest cover Tizen and webOS"). Retracted — see `comparison-pages.md`.

### ✅ TV performance — much stronger than I first wrote. **Be specific by TV type.**

> **Correction 2026-07-16.** I earlier called our TV perf "black/freeze only". That was
> me reading `tv/capture/qoe.go` — the **camera** layer — as if it were the whole
> story. It isn't. **There are two kinds of TV and they use completely different perf
> stacks.** Never say "TV performance" without saying which.

| TV type | Perf stack | Evidence |
|---|---|---|
| **Android TV / Fire TV** (= Android) | Full Android suite — **17 collectors**: CPU (`/proc/stat`), memory (`dumpsys meminfo`), **FPS/jank** (`dumpsys gfxinfo framestats`), battery, **cold/warm/hot startup** (`am start -W`), network, disk, GPU, thermal, GC, wakelock, FDs, SQLite, logcat streaming | `AppVitals/pkg/android/` |
| **Apple TV** (= iOS) | Full iOS suite over **Instruments/DTX**: CPU, memory (`physFootprint`), FPS (`graphics.opengl`), battery, disk, network, GPU + our **custom raw sysmontap service** | `AppVitals/pkg/ios/` |
| **Tizen / webOS** (= web apps) | **CDP**: `Vitals` — FPS (rAF-sampled), FCP, TTFB, DOMContentLoaded, LoadEvent; `Metrics` — JS heap used/total/limit, DOM node/document/listener counters (leak trends), LayoutCount, RecalcStyleCount, ScriptDuration, TaskDuration, LayoutDuration + raw perf map; `CompositorFPS` via tracing | `tv/internal/driver/cdp/{vitals,metrics,tracing}.go` |
| **Any TV, incl. DRM/native** | **Camera QoE** — black + freeze detection from the H.264 camera feed. Works on apps we can't reach over CDP | `tv/capture/qoe.go` |

**What we still cannot claim**: perceptual **video-quality models** — no VMAF, UVQ,
blockiness, PSNR/SSIM, A/V-sync or rebuffering metrics anywhere. `qoe.go`'s own comment
defers FPS/jank and A/V sync to "follow-ups"; `CompositorFPS` needs per-TV empirical
confirmation of trace category names (a zero means "category absent", not "0 fps").

### The honest competitive position on TV performance
- **vs HeadSpin/Witbe**: they have perceptual video quality (VMAF/UVQ); we don't.
  **Don't fight on video quality.** But we have deep **app** performance on every TV
  type — which is a different axis, not a weaker one. Split the frame: *app
  performance* vs *video quality*.
- **vs Suitest**: they disclaim it in writing — *"we are considering to add more
  video-related features, but **QoS/QoE is not our main focus**"* — and their video
  assertions are property reads only (length/position/state/URL). **This is a clean,
  citable gap.**
- **vs Device.io**: perf is marketing-only; their perf docs page 404s.

### ✅ The architectural point that's ours alone

Suitest reads **the video object's own properties** — "directly from the platform's
video object bypassing all upper layers". That's **white-box**: it reports what the
*player believes*. We watch the panel through a camera because **DRM renders the
framebuffer black** — **black-box**: what the screen *actually shows*.

If playback decodes to black while `position` keeps advancing, a property read says
"playing"; the camera says black. `qoe.go` names this exactly: black and freeze are
*"the two outside-in failure signals **no in-app instrumentation can see**"*. **That is
the "'it returned OK' is not 'it worked'" thesis**, and two drafts already sit on it
(`blog-camera-webrtc.md`, `blog-ok-not-worked.md`). No competitor can write it: Suitest
reads properties by architecture, Device.io proxies Appium.

Frame it as an **architectural difference with their docs quoted verbatim** — state
both approaches, cite, let the reader conclude. Do **not** write "Suitest can't detect
black frames" (unverified competitor claim + disparagement risk, §1).

### ✅ What we can defensibly own instead — and it's better

Suitest's video assertions read **the video object's own properties** — "Video length,
Video position, Video state, Video URL" — "directly from the platform's video object
**bypassing all upper layers**". That is a **white-box** read: it tells you what the
*player believes*. We watch the actual panel through a camera because **DRM renders the
framebuffer black** — a **black-box** read of what the screen *actually shows*.

If playback decodes to black but `position` keeps advancing, a property read says
"playing". The camera says black. **That is exactly the "'it returned OK' is not 'it
worked'" thesis** — and we already have two drafts sitting on it
(`blog-camera-webrtc.md`, `blog-ok-not-worked.md`).

This is ours, it's code-backed, it needs no QoE overclaim, and no competitor can write
it: Suitest architecturally reads properties, Device.io proxies Appium. Frame it as an
**architectural difference with their docs quoted verbatim** — state the two approaches,
cite, let the reader conclude. Do **not** phrase it as "Suitest can't detect black
frames" (that's an unverified competitor claim and disparagement risk under §1).

**Verified [2–1]: HeadSpin owns Appium TV tooling authority.** HeadSpin authors and
maintains `appium-roku-driver`, `appium-tizen-tv-driver` and `appium-lg-webos-driver`,
is **Appium's first named Development Partner**, and employs two Appium core
maintainers (per appium.io's own docs, not HeadSpin marketing). All three repos are
live and unarchived.
→ **We must not position as "the ones who understand the Appium TV drivers."**

**Our actual angle is better than the one the research assumed.** We don't use the
Appium TV drivers — we implemented **SDB and SSAP natively in Go**. So the honest
position isn't "Appium TV authority" (HeadSpin's) or "widest TV matrix" (Suitest's);
it's **"we went a layer below the driver and here's what we found."** That is exactly
what `blog-native-drivers.md` and `blog-webos26-field-guide.md` already do, and no
competitor — including HeadSpin — can write those posts, because they consume the
protocol rather than implement it.

**Also decay-watch**: HeadSpin's driver stewardship shows decay signals (Jonathan
Lipps departed to Sony; Development Partner announcement is 2 years old; HeadSpin
University closed). Re-check in 12 months — if maintenance lapses, the tooling
authority is up for grabs.

---

## 4. Cluster 2 — Appium TV driver how-tos (DEPRIORITISE for us)

**Verified [3–0]**: the smart TV Appium ecosystem is genuinely fragmented — five
drivers, five artifact formats: `appium-roku-driver` (.zip), `appium-lg-webos-driver`
(.ipk), `appium-tizen-tv-driver` (.wgt), `appium-xcuitest-driver` (.ipa),
`appium-uiautomator2-driver` (.apk), plus `appium-smartcast-driver` for Vizio.
Corroborated by Appium's own 3.2 ecosystem docs and TV Labs' platform table.

It's a real, documentable how-to surface — **but it's HeadSpin's home turf (they wrote
the drivers), we don't use these drivers, and BrowserStack/LambdaTest/HeadSpin/Sauce
are all Appium sponsors saturating generic Appium content.** Writing five driver
how-tos would be commodity content on someone else's ground.

**Exception worth one post**: the *artifact-format fragmentation* itself (.zip/.ipk/
.wgt/.ipa/.apk — five packaging formats for one "TV app" concept) is a genuine
practitioner pain we've lived through (Tizen cert-permit flow, webOS ipk staging into
`/media/developer/temp`, dev-mode expiry). That's ours to write.

---

## 5. Cluster 3 — Data residency / regulated industry (high priority, needs a fact fix)

**Verified [3–0] and CRITICAL: do NOT hang this on DPDP.** India's DPDP Act contains
**no general data-localization mandate**. Section 16 is a *blacklist* model — transfer
is permitted to any country *except* those the government notifies, and **no country
has been notified**. The Nov 13 2025 DPDP Rules did not change this (Rule 15 preserves
the permissive default).
**A post claiming "DPDP requires your test devices to stay in India" would be
factually wrong** — and would fail the citation standard the whole content strategy
rests on.

**The real hook is SECTORAL**: RBI (all payment system data "stored only in India",
circular DPSS.CO.OD No. 2785/06.08.005/2017-18, 6 Apr 2018), SEBI, and IRDAI (cite the
**2025** regulations, not the superseded 2015 ones). That's the accurate, and
genuinely under-written, post: *"what RBI/SEBI/IRDAI actually require about test data
and device location."*

**Timing [2–1]** — this is a *rising*, not peaked, topic: DPDP substantive compliance
lands **13 May 2027**; consent-manager provisions **13 Nov 2026**; the Data Protection
Board has been live since 13 Nov 2025. A MeitY proposal may accelerate Significant
Data Fiduciary obligations to Nov 2026 — which likely captures our large streaming and
telecom customers. **Publish now, refresh at the Nov 2026 and May 2027 milestones —
treat as a maintained page, not a one-shot post.**

**Refuted [0–3] — do not use**: the hypothesis that BrowserStack hides where test
artifacts live. They document retention precisely (text logs 60 days; Automate
dashboard logs 30 days; App Automate video 30 days, screenshots/logs 60 days,
auto-deleted) and disclose it openly. Honest framing is *"BrowserStack retains test
artifacts for 30–60 days by its own documentation"* — **never** "BrowserStack hides
that it stores your data," which is unsupported and legally reckless (and would breach
the India no-disparagement rule in `comparison-pages.md` §1).

✅ **UNBLOCKED 2026-07-16.** Product: data handling **varies by customer** — RobusTest
doesn't control or know what testers load onto devices. So don't write *"we handle
regulated production data"*. Write the reader's problem instead: **you don't know what
your QA team put on that device, and neither does your cloud vendor** — which is
precisely why sectoral regulators care where the device sits. Honest, accurate, and a
stronger hook than a capability claim.

---

## 6. Cluster 4 — Untapped Tier 1 engineering posts (no competitor can match)

None of these are drafted. All are code-verified. Demand is unverified but these are
link/citation bait rather than search plays — and per the settled GEO finding
(stats + quotes + cited sources ≈ 30–40% visibility lift), original-data engineering
posts fit the citation strategy better than how-tos.

1. **"Mann-Whitney U in production: how we decide if a build actually got slower"** —
   non-parametric stats because mobile perf data isn't normal; Cliff's delta for effect
   size; the three-filter verdict; comparability blocking. *Rare topic, real code.*
2. **"One CA, many replicas: an SCEP depot on MongoDB"** — why micromdm/scep's
   filesystem/BoltDB depots break behind a load balancer; atomic `$inc` serial minting.
3. **"Watermarking every frame a tester sees"** — pre-encode forensic watermarking for
   pre-release content on a shared device lab.
4. **"Your TV's browser is frozen at manufacture"** — CDP 1.3 as a compatibility floor;
   a TV's automation capability never changes after purchase.

⚠️ Open question from research (unresolved): **do original-data/benchmark posts
outperform how-tos for links and AI citations for dev-tool blogs?** No evidence
survived. Given the GEO finding, benchmark posts are a plausible highest-leverage
format — but that's a hypothesis, not a finding.

---

## 7. Competitive content landscape (who owns what)

| Player | Owns | Don't fight them on |
|---|---|---|
| BrowserStack | Enormous `/guide/` library, head terms, generic Appium/Selenium how-tos | Generic mobile testing how-tos |
| LambdaTest / TestMu AI | Learning hub, generic automation content | Same |
| HeadSpin | Appium TV **tooling** (they author the drivers); Appium Development Partner | Appium TV driver internals |
| Suitest | Widest TV/OTT product matrix (incl. consoles, HbbTV); published pricing; 103 logos | "We support the most TV platforms" |
| Device.io (ex-TV Labs) | Appium TV driver matrix docs; heavy streaming logos | Appium-proxy convenience |
| Witbe | Real-device **QoE** monitoring | Video quality depth |
| stb-tester | HDMI-capture, **genuinely on-prem** | "the only on-prem TV lab" |
| **Nobody** | **Editorial depth on Tizen/webOS at the protocol layer; black-box camera ground truth vs white-box property reads; on-prem device labs under Indian sectoral regulation; production statistics for perf regression** | ← our ground |

Note: **no TV specialist publishes a blog.** Suitest has docs + changelog; Device.io's
blog 404s. That's the editorial vacuum and the best argument for this program. But
their existence means TV breadth is **table stakes**, and QoE is **contested ground we
would lose** (§3).

---

## 7b. ⭐ Cluster P — Mobile performance: the ONLY demand+authority+no-incumbent overlap

**This is the lead cluster.** `app startup time` and `android jank` both cap
autocomplete at 10 with open, developer-intent suffixes ("jank detection", "janky
frames", "jankstats") and **no brand owner** — unlike every other query with demand,
which resolves to browserstack/aws/free (§2a). And they sit on our deepest verified
authority (§2e). Three rounds of research missed this because nobody measured it.

| # | Working title | Demand | Our authority |
|---|---|---|---|
| P1 | **"How to actually measure app startup time: cold vs warm vs hot"** — why one "startup" number is a lie; `am start -W`; what each mode really tests | ✅ `app startup time` = 10 | `cold_start_ms`/`warm_start_ms`/`hot_start_ms` as **three separate scored metrics** |
| P2 | **"Jank detection on Android: what `dumpsys gfxinfo framestats` actually tells you"** | ✅ `android jank` = 10, → "jank detection" | `jank_rate`, `frozen_rate`, `frame_p95`, `slow_frames` |
| P3 | **"Your p95 frame time is lying to you"** — percentiles vs averages in frame data | ✅ (jank cluster) | `frame_p95` + the stats layer |
| P4 | **"Did this build actually get slower? Mann-Whitney U says maybe"** — non-parametric significance for perf, because mobile perf data isn't normal; Cliff's delta for effect size; the three-filter verdict | 🔵 inferred | **Unmatched** — `nerve/auth/appvitals/stats.go` |
| P5 | **"Thermal throttling makes your benchmark a fiction"** — `throttle_count`, `cpu_temp_max`; why rack-mounted phones lie | 🔵 inferred | `dumpsys thermalservice` — **almost nobody surfaces this** |
| P6 | **"Finding file-descriptor leaks in mobile apps"** — `fd_growth_rate`, `open_fds_max` | 🔵 inferred | **Rare — no competitor surfaces FD leak metrics** |
| P7 | **"Your SQLite queries are the reason the UI stutters"** — `sqlite_slow`, `sqlite_avg_ms` | 🔵 inferred | Rare |
| P8 | **"Touch latency: dispatch vs delivery"** — why splitting them finds the real culprit | 🔵 inferred | `input_dispatch_ms` vs `input_delivery_ms` split |
| P9 | **"32 metrics, 4 scores: how we grade an app's performance"** — the threshold table + category weights, published openly | 🔵 inferred | `pkg/metrics/scoring.go` |
| P10 | **"Reading iOS performance without Instruments.app"** — DTX/sysmontap; our raw sysmontap service exposing what go-ios's API doesn't | 🔵 inferred | `pkg/ios/services.go`, `sysmontap.go` — but **credit go-ios**; we didn't write the protocol stack |

**Collision check**: devicelab.dev's 105 posts are device-lab ops, Maestro, and vendor
comparisons — **it has no metric-level performance content**. Clean.
**Competitor check**: `app performance testing` is BrowserStack-anchored, but
`android jank` / `app startup time` are **not** — attack those, not the category term.
**Rule**: name the rare metrics; **never claim "more metrics than anyone"** (§2e).

---

## 8. Recommended program

**First, decide this in writing (§2a):** our differentiators have **zero search
demand** (on-prem = 0 autocomplete, Tizen/webOS = 0, smart TV = tiny + BrowserStack-
owned). This blog is an **authority / AI-citation / sales-enablement** asset, not a
traffic play — with **one exception** (Cluster P). Agree that up front or it dies at
the first traffic review.

1. **⭐ Cluster P — mobile performance (P1, P2 first).** The only place real demand, no
   incumbent, and first-hand authority overlap. `app startup time` and `android jank`
   cap autocomplete at 10 with no brand owner. **This is the traffic play.**
2. **Publish the six existing drafts** (after the ⚠️ claim review — Android TV, TV type
   specificity, no "always-on"). Fastest path live; best engineering material; already
   cited. Enrich with stats/quotes per KDD-2024 GEO — **don't restructure** (the
   "original data beats how-tos" premise is unsourced, §2a).
3. **O1 + O2 (Cluster 0)** — every fact already verified from vendor pages; five of
   eight clouds can't do on-prem at all. **Zero search demand — these are sales
   enablement.** Write them for the buyer conversation, not for Google.
4. **The comparison bridge**: people literally search *"smart tv testing browserstack"*
   and *"app performance testing browserstack"*. That's where the searchers are —
   ride it via `comparison-pages.md` §6, lawfully (§1).
5. **O3/O4/O5/O12/O16** — the on-prem engineering set (SCEP, storage, watermarking,
   customer-run mitm, customer-owned MDM CA). Code-verified; unmatched; citation bait.
6. **O6 (RBI/SEBI/IRDAI)** — never DPDP. Maintained page; refresh Nov 2026 / May 2027.
7. **Do not**: build volume how-tos on BrowserStack's or HeadSpin's ground; position on
   TV breadth (Suitest wins); chase the State of Testing Report (it's empty); cite SO
   tag counts as demand; claim "more metrics than anyone".

## 9. Open gaps
- [ ] Real search-demand data for the TV cluster (paywalled tools blocked the round)
- [ ] Practitioner pain from Reddit/HN/SO/State-of-Testing 2026 — agent dispatched
- [ ] Suitest/TV Labs full profile + completeness of the TV-specialist competitive set
      (Witbe, Testronic, Applause, Eurofins?) — agent dispatched
- [ ] Does RobusTest touch production data? (product question, blocks Cluster 3)
- [ ] Benchmark posts vs how-tos for citations — unresearched
