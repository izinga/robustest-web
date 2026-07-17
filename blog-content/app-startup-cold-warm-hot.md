---
title: "The startup number everyone quotes is the failure line, not the target"
description: "'Keep cold start under 5 seconds' is everywhere — but 5s is Google's excessive-failure threshold, not a target. What the tools actually measure."
category: PERF
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-06-30
sources:
  - title: "Android — App startup time (Android vitals launch-time thresholds)"
    url: "https://developer.android.com/topic/performance/vitals/launch-time"
    note: "Startup is flagged excessive at cold ≥5s, warm ≥2s, hot ≥1.5s. Also the am start -W example output."
---

There is a number that gets repeated in every app-performance discussion: "keep your
cold start under five seconds." It is wrong, and it is wrong in a way that matters,
because it tells you to aim for the exact point at which Google considers your app
*broken*.

## Five seconds is the failure line

Google's own Android vitals documentation defines startup times that it flags as
**excessive**: a cold startup is excessive at **5 seconds or longer**, a warm startup
at **2 seconds or longer**, a hot startup at **1.5 seconds or longer**. Those are not
recommendations. There is no official "cold start should be under X ms" target
anywhere in Google's docs — the five-second figure that circulates as a goal is
actually the line past which your startup is bad enough to be reported against you.

Treating the failure threshold as your target is like treating the speed limit as your
cruising speed on a road where everyone else is going half that. If your cold start is
4.8 seconds you are not "within spec." You are one small regression away from the line,
and you are already slower than a user will tolerate.

## Cold, warm, hot — three different events

The three aren't severities of the same thing; they're different scenarios:

- **Cold start** — the app starts from scratch. The system has to create the process,
  because this is the first launch since boot, or since the app was killed. The most
  expensive case, and the one users hit after not opening your app for a while.
- **Warm start** — a subset of the cold-start work. The process may already exist, but
  an activity has to be recreated (its `onCreate` runs again, possibly restoring saved
  state). Cheaper than cold, more expensive than hot.
- **Hot start** — the activity is still resident in memory and just brought to the
  foreground. The cheapest case.

Reporting one "startup time" number hides which of these you measured — and a build
that looks fine because someone benchmarked a hot start can ship a cold start that's
twice the excessive line. This is why [we record cold, warm, and hot as three separate metrics](/platform/performance-testing), not one.

## What `am start -W` actually reports

The standard way to measure this from `adb` is `am start -W`, and it prints three
numbers that are constantly conflated. Google's own example output:

```
ThisTime:  2044
TotalTime: 2044
WaitTime:  2054
```

Three fields, and they are not the same measurement:

- **ThisTime** — the launch time of the *last* activity in the start sequence.
- **TotalTime** — the launch time of *all* the activities the launch brought up. This
  is the number that reflects your app's own startup cost, and the one to track over
  builds.
- **WaitTime** — `TotalTime` plus the system's own overhead, as observed by the
  ActivityManager. Slightly larger (2054 vs 2044 above), because it includes work
  outside your app.

The common mistake is to grab `WaitTime` (it's the biggest, so it looks like "the real
number") or to quote whichever field a script happened to parse. If you want to know
whether *your code* got slower between two builds, `TotalTime` is the field — and you
want it for a cold start specifically, on a real device, because [an emulator's process creation and I/O don't match the phone your user is holding](/blog/simulators-are-good-enough).

## The point

Two corrections, both small, both consequential:

1. Don't aim for five seconds. Five seconds cold is the line where Google calls your
   app excessive; your target should be a fraction of it. Aim to be *fast*, and measure
   against your own previous build, not against the failure threshold.
2. Say which start you measured, and which field you read. "Startup: 2044 ms" is
   almost meaningless. "Cold TotalTime, 2044 ms, median of 20 launches on a mid-range
   device" is a number you can act on.
