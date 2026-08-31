---
title: "Smart TV app testing: what changes when there is no touchscreen"
description: "Smart TV and OTT app testing is not mobile testing on a bigger screen: focus instead of touch, apps you do not own, expiring developer modes, frozen engines."
category: TV
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-08-20
sources:
  - title: "Android — TV navigation"
    url: "https://developer.android.com/training/tv/start/navigation"
    note: "\"On a TV device, users navigate using either a D-pad or arrow keys. This type of control limits movement to up, down, left, and right.\""
  - title: "Samsung — Remote Control (Tizen TV web apps)"
    url: "https://developer.samsung.com/smarttv/develop/guides/user-interaction/remote-control.html"
    note: "Arrow, Enter and Back key events arrive automatically; every other remote key has to be registered by the app."
  - title: "LG — Magic Remote"
    url: "https://webostv.developer.lge.com/develop/guides/magic-remote"
    note: "webOS's remote is a pointer (\"a wand\") as well as a 5-way key set — two input modes, one app."
  - title: "Apple — Focus-based navigation"
    url: "https://developer.apple.com/documentation/uikit/focus-based-navigation"
    note: "tvOS moves a focus ring between views; there is no touch coordinate to tap."
  - title: "Roku — External Control Protocol (ECP)"
    url: "https://developer.roku.com/docs/developer-program/dev-tools/external-control-api.md"
    note: "\"The External Control Protocol (ECP) enables a Roku device to be controlled over a local area network\" — a REST API on port 8060."
  - title: "LG — Developer Mode app"
    url: "https://webostv.developer.lge.com/develop/getting-started/developer-mode-app"
    note: "When the session runs out, \"the installed apps that you were using on Developer Mode are uninstalled\"."
  - title: "Samsung — TV Device (Developer Mode)"
    url: "https://developer.samsung.com/smarttv/develop/getting-started/using-sdk/tv-device.html"
    note: "Developer Mode is unlocked with a code and bound to the IP address of the host that is allowed to connect."
---

Teams that ship a streaming app to phones and then to living rooms usually assume the
second job is the first one with a bigger screen. Then they try to test it. Smart TV
app testing — and OTT app testing on the boxes under the TV — changes four things at
once, and every one of them breaks a habit that mobile testing taught you.

## 1. There is no touch. There is focus.

A phone test taps a coordinate. A TV app has no coordinate to tap: it has a *focused*
element, and a remote that moves the focus. Android's own guidance is blunt about
it — "on a TV device, users navigate using either a D-pad or arrow keys. This type of
control limits movement to up, down, left, and right." tvOS does the same with a focus
ring. Samsung's Tizen web apps get arrow, Enter and Back automatically and must
register every other key. LG's webOS adds a twist: the Magic Remote is a pointer
("a wand", in LG's words) *and* a 5-way key set, so the same app has two input modes.

For testing this means your test case is a *path*, not a *point*. "Tap Play" becomes
"Right, Right, Down, Enter" — and whether that path lands on Play depends on the focus
order the app declares, which is exactly the thing that breaks when a designer moves a
tile. Focus-order regressions are the single most common smart TV app bug we see, and
no phone test suite has a concept for them.

## 2. Half the apps you test are not yours

On mobile, you test your own build. On a TV, the thing you need to verify is often
something you cannot modify: the competitor app you are benchmarking against, the
platform's launcher that decides whether your tile is visible, a retail channel, or
your own app *as installed from the store* with production DRM keys.

That splits smart TV app testing into two levels. Apps you own and can sign can be
driven white-box — element-level control through the TV's debugging engine. Everything
else is driven black-box: remote keys go in, the real screen and audio come out, and
[the panel is the only oracle](/blog/the-camera-sees-what-the-framebuffer-cant). A
serious OTT test lab needs both, because [a test that passed on the framebuffer can
still have had no sound](/blog/your-tv-test-passed-there-was-no-sound).

## 3. Developer mode is a lease, not a setting

A phone stays in developer mode until you turn it off. A TV doesn't.

LG's Developer Mode runs on a session clock; when it runs out, in LG's own words,
"the installed apps that you were using on Developer Mode are uninstalled." Samsung's
is unlocked with a code and bound to the IP address of the one host allowed to
connect — move the set to a different node and it stops accepting installs. We wrote
up what this looks like on a real panel in the
[webOS 26 field guide](/blog/webos-26-field-guide): a TV that worked last week and now
has no app on it, for no reason it will tell you.

Across a wall of thirty sets, this is a fleet-management problem before it is a testing
problem. Someone has to track every expiry, extend it before it lapses, and re-pair
every set that was moved. If nobody does, the lab quietly wipes itself one TV at a
time.

## 4. The engine never updates

Every phone test matrix is organised by OS version, because the OS updates. A smart
TV's browser engine is fixed at manufacture and stays fixed for the life of the
set — [the TV browser is frozen the day it ships](/blog/the-tv-browser-is-frozen-the-day-it-ships).
So a TV test matrix is organised by *model year*, and "we support Tizen" means nothing
until you say which years. Automation depth follows the engine: a 2022 panel gives you
full element-level control, a 2017 panel gives you keys and video, and no firmware
update will ever move a set from one row to the other.

## What "automated smart TV testing" actually means

Because of all four, there is no single framework for automated smart TV testing the
way Appium is for phones. There are five paths, one per platform family:

| Platform | Automation path | What you can drive |
|---|---|---|
| Samsung Tizen | native `sdb` + Chrome DevTools Protocol | remote keys, install, element-level control on your own web app |
| LG webOS | native SSAP + CDP | same, plus pointer gestures |
| Roku | External Control Protocol (REST on port 8060) | keys, launch and deep-link, UI tree, media-player state |
| Apple TV | XCUITest, same pipeline as iOS | native element automation |
| Android TV / Fire TV | it's Android — ADB and your existing frameworks | everything you already do on phones |

Plus a sixth for everything with only an HDMI port — cable boxes, consoles — where the
"automation" is a remote and a capture card, and the assertion is made on the picture.

The honest summary: smart TV app testing is a *lab* problem more than a *framework*
problem. The framework is the easy part. Keeping forty panels [awake](/blog/tvs-dont-turn-themselves-on), paired,
in-session and at known model years — on [a TV wall inside your own
building](/platform/tv-testing) — is the work.
