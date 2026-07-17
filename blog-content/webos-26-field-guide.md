---
title: "A field guide to webOS 26"
description: "Bringing a new LG webOS 26 set into an automated lab, we hit a run of quirks no document warns you about. Every one of these cost us real time."
category: FIELD
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2025-09-30
updated: 2026-04-02
sources:
  - title: "LG — Web API and web engine specifications"
    url: "https://webostv.developer.lge.com/develop/specifications/web-api-and-web-engine"
    note: "LG's own table pinning each webOS release to its Chromium version."
  - title: "LG — Developer Mode app"
    url: "https://webostv.developer.lge.com/develop/getting-started/developer-mode-app"
    note: "The documented happy path: developer account, passphrase, SSH as prisoner on 9922."
---

The documented path onto a webOS TV is short: install the Developer Mode app, sign
in with an LG developer account, note the passphrase, and you get SSH as the
`prisoner` user on port 9922. Apps are `.ipk` packages, installed through a
developer service call.

That is the happy path. Here is where webOS 26 leaves it.

## 1. `/tmp` is execute-only

The obvious place to stage an install package is `/tmp`. On webOS 26 you can't —
the directory is mounted execute-only. You can traverse it; you cannot write to it.
And the failure isn't a permission error you can read at a glance: you get a copy
that appears to succeed and goes nowhere. There is a directory that works. Finding
it is the afternoon this quirk costs you.

## 2. `journalctl` returns nothing — for *you*

Running `journalctl` as `prisoner` produces empty output. Not an error, not a
permission denial. Empty.

This is the worst possible failure mode when you are debugging an install, because
it looks exactly like "nothing was logged," which sends you off to investigate a
system that is in fact logging perfectly well to a journal your user cannot read.

## 3. Developer Mode expires — and takes your apps with it

The developer session lasts roughly 1,000 hours. When it ends, the TV does not just
stop accepting new installs. **It uninstalls the developer apps already on it.**

If you [run a lab](/platform/tv-testing), this arrives as a device that was working
last week and now has no app on it and no obvious reason why. You either learn to
track and pre-empt the expiry across a whole fleet, or the fleet quietly wipes
itself every six weeks — one TV at a time, on no schedule you set.

## 4. Compression will silently eat your frames

Both major TV vendors mishandle a standard WebSocket compression extension that most
clients negotiate automatically. Let yours turn it on — as most will, by default —
and the TV accepts the connection, completes the handshake, and then silently drops
frames or closes the socket after registration.

No error. Nothing to grep for. And the cause is nowhere near where you'd look, so
the days go to not knowing rather than to fixing.

## 5. One port hangs instead of refusing

The control service listens on two ports. One, on newer firmware, refuses cleanly —
honest, easy to handle. The other, on some sets, doesn't refuse: it hangs silently
for two minutes, which from the client's side is indistinguishable from a slow TV,
so no naive fallback can catch it. Knowing which port to trust, and in which order,
is the kind of thing you only learn by losing an afternoon to it first.

## The meta-lesson — and the point

Every one of these is the system succeeding at the API level while failing at the
level you care about. A copy that writes nowhere. A log that returns empty. A socket
that accepts frames it will not act on. A port that hangs instead of refusing. It's
the same lesson we keep relearning: on a TV, [the only thing that knows whether it
worked is the panel](/blog/the-camera-sees-what-the-framebuffer-cant).

We're not going to hand over the exact fix for each — those are ours. The point is
the tax. This is *one* TV, *one* firmware, and it fought us on five silent things no
document warns about. Multiply that by every brand, every OS version, every year.
That is what running your own TV lab actually is — and it is precisely the work we
do so that you never have to meet any of it.
