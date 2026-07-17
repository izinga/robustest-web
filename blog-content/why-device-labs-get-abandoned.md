---
title: "The device lab you build yourself runs on someone's unpaid weekends"
description: "The open-source stack most in-house device labs run on tells the buy-vs-build story: its maintainers call it underfunded and — their word — a money sink."
category: BUSINESS
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-06-17
sources:
  - title: "DeviceFarmer / stf — project README and releases"
    url: "https://github.com/DeviceFarmer/stf"
    note: "Successor to OpenSTF (openstf/stf), previously backed by CyberAgent and HeadSpin. Self-described funding and pace."
---

When a team decides to build its own device lab instead of [buying one](/blog/which-device-clouds-run-in-your-data-centre), it almost
always builds on the same foundation: **STF** — Smartphone Test Farm — the open-source
project originally released as **OpenSTF**. It's genuinely good software, it's been
around for years, and it's the reason a lot of in-house labs exist at all.

It also, in its own maintainers' words, tells you exactly why those labs so often go
quiet.

## Read the project's own description

STF was developed as OpenSTF and backed by CyberAgent and HeadSpin. Active development
of the original stopped around **July 2020**. What exists today is the community
successor, DeviceFarmer/stf, and its README does not oversell itself. It says the
project is:

> in continued, active development, but development is still largely funded by
> individual team members and their unpaid free time, leading to slow progress…
> somewhat of a money sink.

That is not us characterising a competitor. That is the maintainers of the most widely
used open-source device-lab stack describing, honestly, what it takes to keep it alive:
unpaid weekends, and hardware costs that make it a money sink.

To be fair and accurate: it is **not** abandoned. The latest release, v3.7.9, shipped
on **8 July 2026**. But look at the cadence — the release before it was nearly a full
year earlier — and the nature of the work: OS-image bumps and dependency updates,
maintenance rather than new capability. That is precisely what a project on unpaid time
looks like: kept breathing, not moving forward.

## What "build it yourself" actually signs you up for

The software is the easy part, and it's free. The lab is the hard part, and it isn't.
When you run your own, you inherit the work the open-source project can't do for you:

- **The hardware is a money sink** — the maintainers said it, and they only mean the
  server side. Add a shelf of phones that depreciate, throttle, and swell. Batteries on
  devices kept permanently on charge distend within a couple of years; that's a
  recurring hardware-replacement line nobody budgets on day one.
- **OS updates break your automation** — every Android and iOS release can move the
  ground under your tooling, exactly as it does for the STF maintainers shipping
  image-bump releases. Except now it's your team doing the bumping, on your schedule,
  under your deadlines.
- **The knowledge lives in one person** — the colleague who learned why a device drops
  offline, why enrolment fails intermittently, which cable batch went bad. When they
  leave, the lab starts its slow decline, because the work was never a product, it was
  a person.

None of this is hypothetical. It's the same set of forces that keeps the flagship
open-source project on a once-a-year release cadence funded by free time — applied to a
team whose actual job is shipping an app, not maintaining infrastructure.

## The honest version of the pitch

We're not going to tell you an in-house lab can't work. It can, and for some teams it's
the right call. What we'll tell you is what the open-source project already tells you in
its own README: keeping a device lab alive is continuous, underfunded, unglamorous work
that competes with everything else your engineers are supposed to be doing — and the
first time it competes and loses, the lab starts to rot.

[Buying a managed lab](/platform/device-lab) isn't buying software. The software is free. It's buying the
unpaid weekends — someone else's problem to have the swollen battery, the broken
automation, and the person who left.
