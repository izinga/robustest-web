---
title: "A smart TV's browser is frozen the day it ships"
description: "A phone updates its web engine; a TV never does. That one fact decides how deeply you can automate each panel — and why a TV fleet is never uniform."
category: TV
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-17
sources:
  - title: "LG — Web API and web engine specifications"
    url: "https://webostv.developer.lge.com/develop/specifications/web-api-and-web-engine"
    note: "LG's own table pinning each webOS release to a specific Chromium version — the engine is set per model year."
  - title: "Chrome DevTools Protocol"
    url: "https://chromedevtools.github.io/devtools-protocol/"
    note: "The protocol behind element-level web automation; its available domains and commands track the underlying Chromium version."
---

A phone's web engine is a moving target in the good way: Chrome and the system WebView
update themselves for years after you buy the device, so the automation surface you
test against keeps pace. A smart TV is the opposite. Its browser engine is set at the
factory and, for practical purposes, never moves again. The panel you buy in 2024 will
still be running its 2024 engine in 2030.

That single fact governs how deeply you can automate any given TV — and it is why a
shelf of TVs is never one uniform target.

## The engine is pinned to the model year

This is not a rumour; the vendors publish it. LG's own developer documentation maps
each webOS release to the specific Chromium version it ships with. Samsung's Tizen
line works the same way. Pick a model year and you have pinned the browser engine, and
with it the exact web-platform features and APIs that app can rely on.

For app developers this is a familiar constraint — you code to the oldest engine you
support. For a **test lab** it has a sharper consequence, because the tool you drive
the app with is subject to the same limit.

## Why that sets an automation floor

Element-level automation of a TV web app — click this button, read that text, assert
on real page state rather than a screenshot — runs over the Chrome DevTools Protocol.
CDP is not one fixed thing; its available domains and commands are a function of the
Chromium version underneath it. A newer engine exposes a richer, more reliable
protocol surface. An older one exposes less, and behaves differently where it does
overlap.

So an older panel and a newer panel are not "the same automation, slower hardware."
They present **different control surfaces**. On a recent engine you get full
element-level automation. Below a certain vintage, the protocol simply isn't there to
support it, and the honest thing to do is fall back to what always works on any TV:
[driving the remote and watching the real
screen](/blog/the-camera-sees-what-the-framebuffer-cant). That floor is set by the
model year, not by how much effort you throw at it.

## What it means for a fleet

The practical upshot: automation depth across a TV fleet is not uniform, and you
cannot write one script and assume it runs everywhere. The same test that reads
elements on a 2022 set may only be able to send keys and verify by capture on a 2016
one. A team that discovers this per-device, mid-suite, loses a lot of afternoons to
"why does this pass here and not there."

The fix isn't clever code — it's knowing the boundary in advance. That is why our
[smart-TV support is a model-year matrix you can look up](/platform/tv-testing) rather
than a promise you find the edges of at runtime: for each Samsung and LG vintage, what
you get is full element automation, core automation on an older engine, or remote-plus-video
only. It reflects a constraint the TVs impose — [which is also why we drive them at the
protocol layer in the first place](/blog/why-we-wrote-our-own-drivers) —
not one we chose. Knowing where the floor is turns a surprise into a spec.
