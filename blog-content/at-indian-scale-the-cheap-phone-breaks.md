---
title: "At Indian scale, the cheap phone is the one that breaks"
description: "India buys ~152M phones a year, ~95% Android, most under $200 on 3–4GB of RAM. Test on the flagship and you'll miss where your app actually fails."
category: BUSINESS
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-15
sources:
  - title: "StatCounter — Mobile Operating System Market Share India"
    url: "https://gs.statcounter.com/os-market-share/mobile/india"
    note: "Android sits around 95% of India's mobile OS share; iOS in low single digits."
  - title: "IDC — India's smartphone market grew 4% in 2024 to 151 million units"
    url: "https://my.idc.com/getdoc.jsp?containerId=prAP53185725"
    note: "Annual shipment scale; 2025 closed near 152M by Counterpoint's count."
  - title: "Counterpoint / AppleMagazine — India smartphone shipments 2025"
    url: "https://applemagazine.com/india-smartphone-shipments-2025-growth-counterpoint/"
    note: "2025 shipments ~155M; growth led by premium, but entry and budget still the bulk of volume."
  - title: "DataReportal — Digital 2025: India"
    url: "https://datareportal.com/reports/digital-2025-india"
    note: "Median mobile download speed ~100.78 Mbps (Ookla, Jan 2025) — a national median that hides enormous local variance."
---

If you build a mobile app for India, the device that decides whether it succeeds is
not the one your team carries. It is a two-year-old Android phone that cost under
₹15,000, has 3 or 4 GB of RAM, and is three OS versions behind. That phone is the
market. The flagship in your pocket is the exception.

The numbers are not subtle about this.

## The shape of the market

India buys on the order of **152 million smartphones a year** — 151 million in 2024
by IDC's count, around 155 million in 2025 by Counterpoint's. Of the phones running,
**roughly 95% are Android** (StatCounter); iOS is in the low single digits. Whatever
your global instinct about an iOS-first or a fifty-fifty split, it does not apply here.

And the Android that ships is overwhelmingly the cheap Android. In 2025 the
entry-level tier (under ~$100) was about **16%** of shipments and the mass-budget
tier (~$100–200) about **41%** — together nearly **three in five new phones sold**.
Premium is growing and grabs the headlines, but volume still lives at the bottom of
the price ladder, on 3–4 GB of RAM.

That is the device your median user is holding. It is memory-constrained, its OEM
skin is heavier than stock Android, and it will be on that OS version for years
because its browser and system components are effectively frozen after launch.

## Why the cheap phone is where bugs live

A flagship hides your mistakes. It has the RAM to keep your process warm, the CPU to
absorb a sloppy main thread, and the thermal headroom to never throttle. Ship a
memory leak or a heavy cold start and a flagship shrugs it off.

The budget phone does not. On 3 GB of RAM, the OS kills your backgrounded process
far sooner, so users hit a [cold start where you assumed a warm
one](/blog/app-startup-cold-warm-hot) — and cold start is exactly the case that
budget hardware makes slow. The same code that renders at a smooth frame rate on a
flagship [drops frames on a cheaper GPU](/blog/what-a-janky-frame-actually-is). None
of this shows up unless you run on the actual hardware, because [an emulator inherits
your workstation's memory and speed](/blog/simulators-are-good-enough), not the
phone's.

Then there is fragmentation. Hundreds of distinct smartphone models are on sale in
India at any time, across a handful of OEMs who each layer their own launcher,
permission behaviour, and power management on top of Android. A permission dialog
that appears on stock Android may be auto-denied by one vendor's battery optimiser
and silently deferred by another's. There is no "representative device" that stands
in for that spread — there is only the device list your users actually own.

## The network is a median, not a promise

It is tempting to read "India's median mobile download speed is ~100 Mbps" (Ookla,
early 2025) and design for it. Don't. A national median is the middle of an enormous
distribution: the same app runs on new 5G in a metro and on congested, packet-losing
mobile data in a tier-3 town on the same afternoon. The behaviour that matters is not
what your app does at 100 Mbps — it is what it does when the connection drops mid-request
and comes back thirty seconds later. That is a real-device, real-network test, not a
spec-sheet number.

## What this actually teaches

Twelve years of running device labs in India comes down to one discipline: **test the
low end on purpose, on real hardware, on the exact models your users carry.** Not a
flagship as a proxy. Not an emulator with your laptop's resources. Not the median as
a stand-in for the tail.

It is unglamorous. It is also the difference between an app that works for the top 5%
of the market and one that works for the market. When we [spec a lab for a
customer](/platform/device-lab), the budget-phone shelf is not the afterthought — it
is the point.
