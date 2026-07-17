---
title: "16.67ms is not the number that flags jank"
description: "Everyone quotes the 16.67ms frame budget as the line for jank. Google's own JankStats library doesn't use it — it flags a frame at twice the refresh rate, and that difference is the whole point of measuring jank on real hardware."
category: PERF
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-03
sources:
  - title: "Android — Slow rendering (frame budget by refresh rate)"
    url: "https://developer.android.com/topic/performance/vitals/render"
    note: "16ms @ 60fps, 11ms @ 90fps, 8ms @ 120fps; jank = Choreographer dropping a whole frame."
  - title: "Android — Frozen frames (>700ms)"
    url: "https://developer.android.com/topic/performance/vitals/frozen"
  - title: "Android — JankStats library"
    url: "https://developer.android.com/studio/profile/jankstats"
    note: "Default jank heuristic is 2× the current refresh rate, not a fixed 16.67ms."
---

If you've read anything about Android rendering, you know the number: **16.67
milliseconds** — the time you have to draw a frame at 60fps, and, the story goes, the
line past which a frame is "janky." The budget part is right. The jank-line part is
not, and Google's own tooling disagrees with the popular version.

## The budget depends on the display

16.67ms is the frame *budget* at 60fps — 1000ms divided by 60 frames. But phones
stopped being 60Hz. Google's own docs give the budget per refresh rate:

- **60fps** → render in under **16ms**
- **90fps** → under **11ms**
- **120fps** → under **8ms**

So there is no single "the frame budget." On a 120Hz phone you have half the time you
had on a 60Hz one. A screen that scrolls smoothly on your 60Hz test device can jank on
a user's 120Hz flagship, because the same rendering work now blows a tighter budget.
Which is the first reason the number on a real device matters: the budget is a property
of the hardware, and your test hardware may not be your user's.

## What jank actually is

Jank isn't "a frame that was a bit late." Google's definition is specific: when the
rendering window is overrun, the Choreographer **drops the entire frame** — the display
holds the previous frame for another interval, and the user perceives a stutter. A
dropped frame, not a slow one. That's jank.

At the far end, a **frozen frame** is officially a UI frame that takes longer than
**700ms** to render — long enough that the app looks stuck, not just stuttery.

## The number that actually flags jank

Here's the part that contradicts the folklore. Google's JankStats library — the thing
Google ships specifically to detect jank — does **not** flag frames at 16.67ms. By
default it defines a janky frame as one taking **twice as long as the current refresh
rate**, a deliberate buffer so that normal timing noise doesn't get counted as jank.

That heuristic auto-scales with the display:

- On a 60Hz screen (16.67ms/frame), the jank flag is at ~**33ms**.
- On a 120Hz screen (8.33ms/frame), it's at ~**16.67ms**.

So the famous 16.67ms isn't the jank line at all — on a 60Hz device it's *half* of it,
and it only *becomes* the jank line at 120Hz. Anyone thresholding jank at a fixed
16.67ms is over-counting on 60Hz displays and, worse, applying a number that means
something different on every phone.

## Reading the raw frame data

`dumpsys gfxinfo <package>` gives you the summary — total frames, janky frames,
percentiles. But `dumpsys gfxinfo <package> framestats` gives you the raw ledger: a
per-frame table of nanosecond timestamps for each stage of the pipeline (when the frame
was issued, when the app finished, when it hit the GPU, when it was presented). The
summary tells you *that* frames janked; the framestats table lets you compute it
yourself and see *where* in the pipeline the time went.

[Reading framestats directly](/platform/performance-testing) — rather than trusting a single "jank %" — is what lets you
compare two builds honestly, apply the right per-refresh-rate threshold, and catch the
120Hz regression that a 60Hz summary would have called clean.

## The takeaway

Two things to unlearn: the frame budget is not one number, it's a function of the
display; and 16.67ms is not where jank is flagged — twice the refresh rate is, which is
~33ms at 60Hz. Measure on the refresh rate your users actually have, threshold against
it, and read the frame data rather than a rounded summary. The whole reason to [measure jank on real hardware](/blog/simulators-are-good-enough) is that the budget, the refresh rate, and the GPU are all properties of a device an emulator doesn't have.
