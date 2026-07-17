---
title: "One less connection is one less thing to debug"
description: "Our node barely touches CPU or RAM, yet we cap phones at the Mac mini's native USB ports — no hubs. Reliability, not capacity, sets the limit."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-05
sources:
  - title: "Hackaday — USB and the myth of 500 milliamps"
    url: "https://hackaday.com/2024/07/03/usb-and-the-myth-of-500-milliamps/"
    note: "USB ports carry a hard per-port power budget — 500 mA on the original spec (100 mA before negotiation), raised to 900 mA on USB 3."
  - title: "Red Hat Bugzilla #1411604 — USB enumeration failure, and the fix"
    url: "https://bugzilla.redhat.com/show_bug.cgi?id=1411604"
    note: "A real kernel-level USB failure where the maintainer's guidance is to 'plug into a controller port directly (no hub).' When the debugging advice is 'remove the hub,' that's the tell."
---

Our device node is almost embarrassingly light. It sits at 10–20 MB of RAM and under
1% CPU while it drives a phone. On paper, one Mac mini could babysit a large stack of
devices without breaking a sweat.

We don't do that. We connect as many devices as the machine has **native USB ports**,
and not one more. No hubs, no port expanders, no oversubscription. The obvious question
is why a node with that much headroom would leave capacity on the table — and the
answer is that in a device lab, the bottleneck was never CPU or RAM.

## The limit is the bus, not the box

The principle behind it is one sentence: **one less connection is one less thing to
debug or fix.** Every hub you add is another device on the bus, another power domain,
another piece of firmware, another set of connectors — another thing that can go wrong
at 3 a.m. and take a real device down with it. Capacity is easy to add. Reliability is
what you're actually selling. Those two pull in opposite directions the moment a hub
enters the picture, and we've been burned enough to know which one wins.

Two concrete reasons the trade isn't close.

**Power doesn't multiply.** A USB port carries a hard power budget — 500 mA on the
original spec (just 100 mA before the device negotiates), raised to 900 mA on USB 3. A
phone under test is not sipping: its screen is on, radios are up, the CPU is working,
and it's charging. Hang several of those off a bus-powered hub and they share one
upstream port's budget — a hub can't hand out power it doesn't have. Under simultaneous
load, something browns out: charging stalls, a device resets, the connection drops. And
a phone that browns out in the middle of a run looks exactly like [a failed
test](/blog/it-returned-ok), sending you off to debug software that was never the
problem.

**Hubs are a documented source of flakiness in their own right.** When USB behaves
strangely, the standard diagnostic — including in vendor and kernel bug reports — is to
remove the hub and plug straight into a controller port. That's the guidance attached
to real enumeration failures: *connect directly, no hub.* When the accepted fix for a
whole class of bugs is "take the hub out," adding one to save a few dollars per device
is a bad trade. Hubs also differ from each other in how they handle edge-case
signalling, so "it works on this hub but not that one" becomes a genuine variable — the
opposite of what [a standardized lab is for](/blog/we-ship-the-lab-you-keep-the-phones).

## The 127-device number is a fantasy

USB's spec famously allows 127 devices on a bus, and that number is worth naming only
to dismiss it. It's a theoretical addressing limit, not an operating target. Long
before you reach it, shared power and per-controller reliability have capped you at a
far smaller, far more honest number: the devices you can hang off real ports and trust.

So we treat the native port count as the actual device limit — not because the node
can't *logically* handle more, but because each connection past the port is a
subtraction from the one property the lab exists to provide. A green run should mean the
app worked, [not that no phone happened to fall off a hub that night](/platform/device-lab).
Trading paper density for a connection you can trust is the right trade every single
time — and the whole reason to make it is so there's one less thing to debug when
something else, inevitably, breaks.
