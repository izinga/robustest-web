---
title: "Simulators win on everything except being real"
description: "The case for simulators is genuinely strong — cheaper, faster, parallel, scriptable. We'll concede all of it. Then there's the one class of bug a stock image structurally cannot contain, and it only ships to real users."
category: BUSINESS
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-05-06
sources:
  - title: "Demystifying Device-specific Compatibility Issues in Android Apps (ICSME 2024)"
    url: "https://arxiv.org/abs/2408.01810"
    note: "197 device-specific issues across 94 open-source repos; the large majority trace to OEM customization of the Android system."
  - title: "Android — Thermal (ADPF) documentation"
    url: "https://developer.android.com/games/optimize/adpf/thermal"
    note: "Google's own docs: the same thermalHeadroom maps to a different thermalStatus on different device models."
  - title: "Don't kill my app! — OEM background-killing benchmark"
    url: "https://dontkillmyapp.com/"
    note: "A whole project (urbandroid-team) exists to rank how aggressively OEMs kill background apps — a class of behaviour absent from any stock image."
---

There is a recurring argument in every mobile QA team, and a version of it turned
up on Reddit not long ago, put about as well as it can be put:

> sim is parallelisable, scriptable, screenshot-friendly, and roughly free per run.
> Real devices are none of those.

That is correct. All of it. We are not going to pretend otherwise, because pretending
otherwise is how real-device vendors lose the argument — by overclaiming, so the
reader stops trusting them at exactly the point they should start.

So let's concede the whole case first, properly.

## Everything the simulator wins

A simulator boots in seconds, on a machine you already have. You can run a hundred in
parallel on a CI box for the price of the CI box. It never has a swollen battery, a
cracked screen, or a cable that came loose. It is scriptable, resettable, and
disposable. For the vast bulk of testing — does the button do the thing, does the
layout hold, does the API call fire — a simulator is not just adequate, it is the
*correct* tool, and reaching for a real device instead is a waste of money and time.

If your test would pass or fail identically on an emulator and a real phone, run it on
the emulator. We build a real-device lab and we will still tell you that.

## The bug that isn't in the image

Here is the part the cost argument can't reach.

A simulator runs a **stock** operating system image. A real device in a user's hand
runs that OS *after the manufacturer has changed it* — the skin, the modified
framework, the OEM services, the power management, the modem stack, the GPU driver.
And a large, well-documented class of bugs lives specifically in that gap.

A peer-reviewed 2024 study catalogued 197 device-specific compatibility issues across
94 open-source Android apps and found the large majority trace directly to
manufacturer customization of the Android system — behaviours that exist on the real
device and simply are not present in the stock image a simulator runs. These are not
exotic edge cases. They are functionality that works on the emulator, works on your
Pixel, and breaks on a specific vendor's phone in the hands of a real user, because
that vendor changed something underneath the app.

You cannot test for the absence of something. A simulator can't reproduce an OEM's
modified permission dialog, its aggressive background-process killer, or its custom
camera pipeline, because *those things are not in it*. Not "hard to reproduce" —
absent by definition.

There is a famous illustration of exactly this. An entire community project —
[Don't kill my app!](https://dontkillmyapp.com/) — exists solely to benchmark and
publicly rank how aggressively different manufacturers kill background apps to save
battery, with names like Nokia, OnePlus, Xiaomi and Huawei near the top of the
offenders list. An app that runs fine in the background on the emulator and on a Pixel
gets silently killed on those OEMs' phones, breaking alarms, sync, notifications, and
long-running work — for real users, in the field. There is a *website* dedicated to a
class of bug that, by construction, cannot occur on a stock image. That's the gap.

## The two it can't even attempt

Two more are worth naming because they are total, not partial:

**Thermal.** Sustained load makes a real phone hot, the SoC throttles, and everything
slows down in a feedback loop your users feel. A simulator has no SoC and no heat, so
it has no throttling. And this isn't uniform across hardware — Google's own thermal
documentation states plainly that the same thermal-headroom value "may be mapped to a
certain thermalStatus on one device model but a different thermalStatus on another,"
because devices are physically built differently. There is no stock image of physics.

**Protected video.** High-end DRM — Widevine L1, the level that streams HD and 4K —
decrypts and decodes inside hardware that an emulator does not have. On a simulator you
get L3 and a downgraded stream, or nothing. If your app plays premium content, the most
important thing it does cannot be verified on a simulator at all.

## The hardware it simply doesn't have

Beyond the OS layer, a real device is a box full of physical parts an emulator only
pretends at. A fingerprint sensor and a face-unlock camera behave differently across
vendors and generations, and their failure modes — a wet finger, a bad-angle face, a
locked-out retry — don't exist on a simulated one. NFC taps, real camera pipelines
(HDR, night mode, the OEM's own post-processing), true GPS drift, actual accelerometer
and gyroscope noise, and a real modem negotiating a weak signal on a real network: none
of these have a faithful emulator equivalent, because the whole point of them is the
messy analogue world the emulator was built to abstract away.

This is why our own [performance measurement](/platform/performance-testing) reads the
counters the operating system keeps on real hardware rather than anything a simulator
can report — the numbers that matter (thermal state, real frame timing, battery drain)
only exist on the physical device.

## Where that leaves it

The honest position is not "real devices, always." It is a division of labour: run the
bulk on simulators, because they win on cost, speed, and parallelism — and run the
things that only exist on real hardware on real hardware, because a simulator cannot
fail a test for a bug it does not contain.

The trap is a test suite that is green on emulators and quietly blind to an entire
category of failure — the category that, by construction, only ever reaches a real
user. (It's also worth knowing, if you do reach for a real-device cloud, [what that
means for where your app ends up](/blog/what-actually-leaves-your-network) — and
[which of those clouds can run inside your own network](/blog/which-device-clouds-run-in-your-data-centre).)
Simulators win on everything except being real. Unfortunately, being real is where a
meaningful share of your production bugs live.
