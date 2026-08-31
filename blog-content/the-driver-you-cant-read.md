---
title: "The driver you ~can't read~"
description: "Vega's official Appium driver is built on open-source Appium and shipped under a proprietary license. The gate on your release is a box you can't open."
category: FIELD
draft: true
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-09-02
sources:
  - title: "npm — @amazon-devices/appium-kepler-driver"
    url: "https://www.npmjs.com/package/@amazon-devices/appium-kepler-driver"
    note: "v3.33.1 at the time of writing. license field: \"See license in LICENSE\"; no repository or homepage in the package metadata; description \"Appium driver for Kepler.\"; dependencies include @appium/base-driver and appium-uiautomator2-server."
  - title: "Amazon — Program Materials License Agreement"
    url: "https://developer.amazon.com/support/legal/pml"
    note: "The proprietary license the driver's LICENSE file points to: no independent redistribution, no derivative works beyond incorporation, and \"You will not reverse engineer, disassemble or decompile the Program Materials.\""
  - title: "npm — @appium/base-driver"
    url: "https://www.npmjs.com/package/@appium/base-driver"
    note: "license: Apache-2.0. The open foundation the closed driver is built on."
---

[Part one](/blog/the-one-line-fork) was about `vda` — adb, renamed. This is the same
move, one layer up, where it matters more: not the debug bridge, but the *test driver*.
The thing your CI leans on to decide whether a build ships.

Vega's official UI-automation path is an Appium driver:
`@amazon-devices/appium-kepler-driver`, still carrying the platform's pre-launch
"Kepler" codename. Appium itself is a gift to the industry — Apache-2.0, forkable,
inspectable, a large part of why mobile [test
automation](/platform/test-automation) is a solved problem. Amazon's driver is built
directly on Appium's open `@appium/base-driver`. And then, at the last step, it stops
being open.

## The manifest of closedness

Everything below is quoted from the package as published on npm:

| field | value |
| --- | --- |
| license | "See license in LICENSE" — the Program Materials License Agreement, Amazon's proprietary license |
| repository | none — no public source |
| homepage | none |
| what ships | compiled JavaScript — build output, not source |
| built on | `@appium/base-driver` — Apache-2.0, open |
| you may | use it, under Amazon's terms |
| you may not | redistribute it, make derivative works, or reverse engineer it |

That last row is not a paraphrase of vibes; it is the Program Materials License
Agreement, which says plainly: "You will not reverse engineer, disassemble or decompile
the Program Materials." The dependency this driver stands on says *take this, change
it, keep your changes.* The thing built from it says *you may run this and nothing
else.* Open source flowed in; nothing flows back out.

> Your test infrastructure decides what ships. On Vega, it's a black box you're
> licensed to run and forbidden to read.

## It's a thin client over a protocol you can watch

Here is the part that makes the closedness pointless rather than merely unfriendly. You
don't need to open the box to see what it does — you can run it, as licensed, and watch
the traffic your own test session produces. It isn't magic. The driver forwards a port
to the device and speaks plain JSON-RPC to an automation server already running there:

```
# 1. reach the on-device automation server
adb forward tcp:PORT tcp:8383

# 2. speak JSON-RPC to it — that's the whole "driver"
POST /jsonrpc  {"method":"getPageSource"}       # the UI tree
POST /jsonrpc  {"method":"findObject", ...}     # locate an element
POST /jsonrpc  {"method":"click", ...}          # act on it
POST /jsonrpc  {"method":"injectInputKeyEvent"} # keys
POST /jsonrpc  {"method":"takeScreenshot"}      # base64 PNG
```

We reproduced the full surface — page tree, element find, click, key inject,
screenshot — talking to `:8383` directly, no driver involved, in about an afternoon of
watching the wire. Which means the closed license isn't protecting a moat. The protocol
*is* the interface; the driver is a courtesy wrapper around it. Closing it guards
nothing and blocks everyone.

The package's own dependency list tells you the ancestry, too: alongside Appium's
base driver it pulls in `appium-uiautomator2-server` — the automation server from the
Android Appium stack, on the platform whose whole pitch is *not Android*.

## The failure mode is the whole problem

A test tool you can't read is fine right up until it misbehaves — and test tools always
misbehave. When the driver returns `socket hang up` because the on-device server wasn't
enabled, or a selector silently stops matching after an OS update, a closed binary
leaves you exactly two options:

**Wait for Amazon.** File a ticket, describe the symptom, and hope a fix lands before
your release does. Your test pipeline's uptime is now someone else's roadmap.

**Rebuild the understanding from the outside.** Trace the traffic, probe the on-device
server, reconstruct what the tool must be doing from what you can observe. You end up
owning the knowledge anyway — just the hard way, and without ever being allowed to
confirm it against the source.

Both roads end at the same place: you needed to understand the tool, and it was
designed so you couldn't. For a debugger, that's annoying. For [the gate on your
release](/blog/it-returned-ok), it's a liability you didn't choose. It is also why,
on the TV platforms that shipped no automation story at all, we chose to [write our
own drivers](/blog/why-we-wrote-our-own-drivers) — a stack we can read is a stack we
can fix at 2 a.m.

## The open path was right there

None of the reasons to ship this closed hold up, because the alternatives cost the same
or less.

**Open the driver.** It is built on Apache-2.0 Appium; releasing it as Apache-2.0 too
would be the natural, frictionless choice — and it would let the community fix
selectors, add capabilities, and keep it alive across OS versions *for* Amazon.

**Or just document `:8383`.** Publish the JSON-RPC methods and anyone can build a
client in any language. The protocol is stable enough that we mapped it from the
outside; a spec would take an afternoon to write and would outlast any single driver.

## The verdict

Amazon took open-source Appium, wrapped a simple on-device protocol in a thin layer of
request-shaping, and shipped the result under a license that forbids opening it. It's
the fork pattern from part one, in the place it hurts most: a closed test driver adds
no capability the underlying protocol doesn't already provide — and it puts the gate on
your release behind glass you can look through but never open. The kindest fix is also
the easiest: open it, or spec the protocol it hides.

*Companion to [The one-line fork](/blog/the-one-line-fork). Same pattern — open in,
closed out — one layer up. License and package metadata are quoted from the package as
published on npm; the on-device automation server and its JSON-RPC surface were
observed on the wire against a live Fire TV Stick running Vega OS 1.2. Appium is
Apache-2.0 (OpenJS Foundation); Vega and Fire TV are Amazon's. This is independent
commentary, affiliated with neither.*
