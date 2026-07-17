---
title: "A field guide to webOS 26"
description: "Bringing a new LG webOS 26 set into an automated lab, we hit a run of quirks no document warns you about. Every one of these cost us real time."
category: FIELD
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
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

The obvious place to stage an install package is `/tmp`. On webOS 26 you cannot.
The directory is mounted `d--x--x--x`: you can traverse it, you cannot write to it.
The failure is not a permission error you can read at a glance — you get a copy
that appears to go nowhere.

Stage into `/media/developer/temp` instead. Live-validated on webOS 26; older sets
do not care either way, so it is safe to do everywhere.

## 2. `journalctl` returns nothing — for *you*

Running `journalctl` as `prisoner` produces empty output. Not an error, not a
permission denial. Empty.

This is the worst possible failure mode when you are debugging an install, because
it looks exactly like "nothing was logged," which sends you off to investigate a
system that is in fact logging perfectly well to a journal your user cannot read.

## 3. Developer Mode expires — and takes your apps with it

The developer session lasts roughly 1,000 hours. When it ends, the TV does not just
stop accepting new installs. **It uninstalls the developer apps already on it.**

If you [run a lab](/platform/tv-testing), this arrives as a device that was working last week and now has
no app on it and no obvious reason why. The session token lives in the TV's
preferences, and there is a service call to renew it before it lapses. Read it, and
renew on a schedule, or plan to reinstall your fleet every six weeks by surprise.

## 4. Compression will silently eat your frames

Both major TV vendors mishandle the WebSocket `permessage-deflate` extension. If
you let your client negotiate compression — and most clients will, by default — the
TV will accept the connection, complete the handshake, and then silently drop
frames or close the socket after registration.

There is no error. Disable compression explicitly.

## 5. Try port 3001 first, and only 3001

The control service listens on 3000 and 3001. Newer firmware refuses 3000 with an
immediate reset, which is at least honest. But a failure on 3000 against some sets
is not a reset — it is a **silent 120-second hang**, which no fallback logic can
catch, because from the client's side it is indistinguishable from a slow TV.

Try 3001 first. Do not be clever about falling back.

## The meta-lesson

Every one of these is a case of the system succeeding at the API level while
failing at the level you care about. A copy that writes nowhere. A log that returns
empty. A socket that accepts frames it will not act on. A port that hangs instead
of refusing.

Which is really the same lesson we keep relearning: on a TV, [the only thing that knows whether it worked is the panel](/blog/the-camera-sees-what-the-framebuffer-cant).
