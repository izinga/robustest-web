---
title: "What a $30 stick does *under load*"
description: "No SDK, no source — adb and 45 seconds of Perfetto on a Fire TV Stick. A native player vs a webview one, and the tax that shows up under load."
category: PERF
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-09-04
cta_heading: "Vega OS is already in the rack."
cta_text: "RobusTest runs Amazon Fire TV — sticks included, Fire OS and the new Vega OS — alongside your phones, tablets and other TVs: manual sessions, test automation and performance runs, on real devices inside your own network."
sources:
  - title: "Perfetto — System profiling and app tracing"
    url: "https://perfetto.dev/docs/"
    note: "The open-source tracer used for the capture. The stick runs the tracer on-device; adb is the only tool on the host."
---

A thing the Vega docs bury: you don't need Amazon's SDK, a developer account, or an
app's source to profile it. The Fire TV Stick runs the tracer on-device; `adb` is the
only thing on your machine. Start a trace, do the thing, pull the file, read it in
Perfetto. That works on *any* app on the box — including ones you'll never hold the
source to, which after [part one of this series](/blog/the-one-line-fork) should not
surprise you: the bridge is adb, so the whole adb workflow comes with it.

So we took two commercial streaming apps and gave each the same job: play video for 45
seconds on the same device — a stick with 4 cores and 974 MB of RAM — captured the same
way. The only structural difference between them is architectural: one renders through
a **native** player, the other paints its UI and video surface through a **webview**.
That single choice turned out to be the whole story.

## Native vs webview, both playing video

The intuition most people start with is that a video app is a video app — the heavy
lifting is decode, the OS does that either way, so the app layer shouldn't matter
much. The trace disagrees. The rendering architecture leaves a clear, measurable tax
on the two things a memory-starved stick cares about most.

| 45 s of playback | Native player | WebView player |
| --- | --- | --- |
| CPU, average (% of 4-core SoC) | 11.9% | 15.9% |
| CPU, peak | 16.3% | **28.0%** |
| Memory, average (RSS) | 192 MB | 196 MB |
| Memory, peak | 197 MB | **412 MB** |

Two things jump out. The native player runs lighter *and* steadier — the webview
player carries its rendering through an extra engine, costing more CPU on average and
peaking nearly 2× higher. And the memory story is sneakier: **nearly identical on
average**, 192 vs 196 MB. Averages would call this a tie. But the webview player
spiked to 412 MB while the native one sat flat at 197 MB — and on a 974 MB device,
that spike is the difference between calm playback and the OS hunting for memory to
reclaim.

> Same movie, same decode, same silicon. The webview tax lands entirely in the app
> layer — and it's the volatile one.

## The trace draws the whole video pipeline

What makes this more than a pair of bar charts: the trace shows the entire hardware
playback chain lighting up, identical on both apps, because it's the OS doing this
work, not the app — DRM decrypt (`drmserver`), a vendor hardware-decode daemon
(`vpud`), media transform and buffering, the Wayland compositor (`weston`), the audio
mixer. That shared pipeline costs about 24–25% of the SoC under either app.

The *app's* CPU sits on top of that floor. So when the webview player burns four
points more than the native one, that's pure app-layer overhead — and on this class of
device, four points of SoC is a fan you can hear.

## Idle browsing lied to us first

Our first capture of the webview app looked fantastic — 6.5% CPU, 83 MB. Then we
noticed what it was doing: *menu browsing*, not playback. Comparing a scrolling menu
to a playing movie is how you "prove" whatever you already believed. The moment both
apps did the same thing, the ranking flipped.

The lesson isn't about either app. It's that a performance number without a matched
workload is a vibe, not a measurement — the same reason
[startup time needs cold, warm and hot kept apart](/blog/app-startup-cold-warm-hot),
and the reason our own [performance testing](/platform/performance-testing) is built
around measuring apps from the outside, no SDK in the build, so the workload is the
real app doing the real thing.

## What we actually learned

You can profile any Fire TV app — source or no source — with adb and a trace. And when
you match the workload, architecture shows up in the numbers: the native player beat
the webview player on CPU stability (16.3% vs 28.0% peak) and on memory calm (a flat
197 MB against a 412 MB spike), over an identical 24–25% shared pipeline. On a 974 MB
stick, that difference is the product.

*Method note: single 45 s sample per app, different content and scenes, no per-frame
FPS — Vega's frame-trace category emits no slices, so frame timing needs the SDK's own
tooling. CPU is scheduler on-CPU time bucketed per second; memory is per-process RSS
polled at 1 Hz. Directional, not certification-grade — but repeatable by anyone with a
cable. App names withheld deliberately: this is a native-vs-webview architecture
comparison, not a brand scorecard. Captured on a live Fire TV Stick (Vega OS 1.2);
part of the Vega teardown series with [The one-line fork](/blog/the-one-line-fork),
[The driver you can't read](/blog/the-driver-you-cant-read) and [Fire TV killed Chrome
DevTools. Good.](/blog/fire-tv-killed-chrome-devtools-good) Independent commentary,
not affiliated with Amazon or Google.*
