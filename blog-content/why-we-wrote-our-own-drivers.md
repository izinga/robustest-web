---
title: "The SDK stops where the test lab begins"
description: "We wanted to drive real Samsung and LG TVs at scale — install, launch, press buttons, read performance, headless. The vendor tooling isn't built for that. So we implemented the wire protocols ourselves."
category: WHY
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-02-18
sources:
  - title: "Samsung — Smart Development Bridge (sdb)"
    url: "https://github.com/Samsung/tizen-docs/blob/master/docs/application/tizen-studio/common-tools/smart-development-bridge.md"
    note: "The vendor CLI: install, launch, shell. Nothing for driving the remote."
  - title: "LG — CLI (ares) introduction"
    url: "https://webostv.developer.lge.com/develop/tools/cli-introduction"
    note: "ares-install / ares-launch / ares-inspect. Also no remote control."
  - title: "Appium — driver ecosystem"
    url: "https://appium.io/docs/en/latest/ecosystem/drivers/"
    note: "The TV drivers are third-party, authored by HeadSpin, and proxy CDP over the same SDK limits."
  - title: "appium-tizen-tv-driver (HeadSpin)"
    url: "https://github.com/headspinio/appium-tizen-tv-driver"
  - title: "appium-lg-webos-driver (HeadSpin)"
    url: "https://github.com/headspinio/appium-lg-webos-driver"
---

Both Samsung and LG ship perfectly good developer tooling. It is built for
*developing an app*. Running a test lab is a different job, and the tools stop
before it.

## What the SDKs give you

**Tizen** connects over `sdb`, the Smart Development Bridge — a client/server/daemon
CLI that talks to the TV over TCP and can package, install, launch and uninstall
apps. **webOS** has `ares-cli`: `ares-install` for an `.ipk`, `ares-launch`,
`ares-inspect` to attach the web inspector. Both give you a solid inner loop for
one developer, one TV, one app.

They do not give you a lab.

## Where they stop

You hit three walls the moment you try to automate a shelf of TVs.

**There is no programmatic remote.** This is the big one. The official CLIs cover
install, launch and debug — and nothing for *driving the remote*. There is no
`sdb press-key RIGHT`. To simulate a physical remote you have to speak protocols
the SDK never exposes: Samsung's WebSocket remote channel, and on webOS a separate
Magic-Remote pointer socket that only the community has ever documented. A TV test
that cannot press buttons is not a TV test.

**Everything is IDE-first and single-device.** The CLIs act on one target at a
time — `sdb -s <serial>`, `ares -d <device>` — with no fleet abstraction; fan-out
is your problem. Onboarding is manual and hostile to scale: webOS Developer Mode is
time-limited and, on expiry, *uninstalls the apps you side-loaded*; Tizen bakes each
device's hardware ID into the signing certificate, which cannot be changed after the
cert is made, so adding a TV means re-issuing certs.

**There is no unified interface.** Tizen (`sdb`, `.wgt`, Samsung's remote channel)
and webOS (`ares`, `.ipk`, the WebSocket API plus a pointer socket) are entirely
separate stacks — different transport, packaging, signing and key schemes. Nothing
lets you say "install this build and press OK" across both brands.

## What about Appium?

You would hope the automation ecosystem papered over this. It mostly hasn't — and
credit where it is due: the TV drivers that exist are **authored and maintained by
HeadSpin**, who are Appium's first development partner and employ Appium core
maintainers. They did real work that we lean on the existence of.

But the `appium-tizen-tv-driver` and `appium-lg-webos-driver` automate
*web-framework apps only*, essentially as CDP proxies that require Developer Mode
and a debuggable app. TVs ship very old Chromium, so you hand-match an old
Chromedriver, and some older sets cannot be automated that way at all. It is a proxy
over the same SDK limits, not a way around them.

## So we implemented the protocols

Instead of wrapping tools that couldn't do what we needed, we wrote native drivers —
one per TV OS, the wire protocol owned end to end.

**Tizen** rides `sdb`, which is a fork of Android's `adb` — not "similar to," a fork,
with AOSP copyright headers in its source. So we implemented the transport framing
directly and run install, shell and the debugger over one persistent connection, no
SDK process in the loop. That reimplementation is also where we found the public
protocol documentation is wrong in three specific ways, and eventually hit a wall we
couldn't climb — [a separate post](/blog/where-we-stopped) is entirely about that.

**webOS** rides the SSAP WebSocket for app lifecycle, the Magic-Remote pointer socket
for input — D-pad *and* a real cursor — and Developer-Mode SSH for installs.

On top of both sits **one interface** the rest of our system talks to: install this
build, launch it, press OK, read performance. The brand differences stop at the
driver boundary.

## Why owning the protocol was worth it

More work up front, plainly. You inherit the reverse-engineering, and you track
undocumented changes yourself — a new OS version can move the ground under you. But
for a *control plane*, that ownership is the whole point:

- **You can make it resilient.** Reconnect logic, keepalives, isolating the remote
  channel from control-connection churn — the difference between a demo and a shelf
  of TVs running for months. We learned that the hard way, with [a remote that reported every keypress "delivered" while the TV received nothing](/blog/it-returned-ok).
- **You get the capability the SDK doesn't have** — real remote and pointer input —
  because you speak the same protocol the physical remote does.
- **You can go headless and fan out**, because you are not driving a single-target
  CLI.
- **One surface, two brands.** New automation never asks whether it's Samsung or LG.

The vendor SDK is the right tool for writing an app. For [turning a rack of real TVs into a dependable, automatable fleet](/platform/tv-testing), the SDK is where the work *starts* — and owning
the protocol is how you finish it.
