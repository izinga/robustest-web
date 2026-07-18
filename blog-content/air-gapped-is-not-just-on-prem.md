---
title: "Air-gapped is not just on-prem with the cable pulled"
description: "'On-premise' and 'air-gapped' aren't the same. Cutting a lab's last path to the internet breaks licensing, push, updates and time. Here's what it takes."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-18
sources:
  - title: "Air gap (networking) — definition"
    url: "https://en.wikipedia.org/wiki/Air_gap_(networking)"
    note: "An air-gapped network has 'no network interfaces, either wired or wireless, connected to outside networks' — a permanent architectural isolation, not a firewall rule."
---

"On-premise" and "air-gapped" get used as if they mean the same thing. They don't, and
the gap between them is where most self-hosted setups quietly fail a real security
review. On-premise means the hardware sits in your building. Air-gapped means it also
has **no path to the internet at all** — no wired or wireless interface to any outside
network, as a permanent architectural choice rather than a firewall rule you could
change on a Friday.

That second condition sounds like a small addition. It is not. A device lab is a web of
quiet assumptions that the internet is there, and air-gapping surfaces every one of
them at once.

## What breaks the moment you pull the cable

None of these is exotic. Each is a thing that "just works" with a connection and stops
dead without one:

- **Licensing and activation.** Software that phones home to activate or validate a
  licence simply can't. If the platform running your lab needs to check in with a vendor
  server to stay alive, an air gap kills it.
- **Push and cloud services.** Anything that rides a cloud push service or calls a cloud
  API — for control, for notifications, for a hosted dashboard — has nothing to reach.
- **Updates, over the air.** Device OS updates, platform updates, package pulls: all of
  them expect a repository on the internet. Behind an air gap, nothing arrives unless you
  brought it in deliberately.
- **Time.** With no public NTP source, clocks drift. That is not cosmetic — TLS
  handshakes, certificate validity, and any time-sensitive test start failing in ways
  that look like unrelated bugs.
- **Certificate checks.** Revocation lookups (CRL/OCSP) that reach out to a certificate
  authority hang or fail closed, so a perfectly valid connection times out for a reason
  nobody thinks to check.
- **Runtime dependencies.** Every font, map tile, analytics endpoint, or SDK call your
  app makes at runtime to a server that no longer answers.

Individually, each is a shrug. Together, they are why "we support on-premise" and "we
run air-gapped" are very different claims.

## Why most "on-premise" can't actually air-gap

Here is the uncomfortable part for a lot of the market: plenty of "on-premise" offerings
still call home. They phone out for licensing, they ship telemetry, they keep a control
plane in the vendor's cloud that the on-prem piece connects back to, or they pull updates
from the internet. That is on-premise *and connected* — which is fine right up until a
reviewer asks you to prove the lab works with the cable physically out, and it doesn't.
It is the same reason [most device clouds can't run in your data centre at
all](/blog/which-device-clouds-run-in-your-data-centre): the architecture assumed a
connection that a real air gap forbids.

## What it actually takes

Running a device lab truly air-gapped is an architecture decision made up front, not a
flag you flip. The whole stack has to survive with no path out: offline licensing that
doesn't call a vendor; an internal registry or mirror so builds and updates come from
inside; an internal time source; a control plane that runs entirely on your side with no
mandatory cloud dependency; and no phone-home telemetry that fails the review by simply
existing. Every place the system would reach the internet has to have an inside answer.

The test for it is refreshingly simple: **unplug the internet and see what still works.**
For the environments where this matters — defence, central banks, critical
infrastructure, anyone whose [data legally cannot
leave](/blog/data-residency-reaches-the-test-device) — that is the actual requirement,
and "on-premise" that still reaches out does not meet it. That is why, when we build a
lab for those teams, [everything it needs lives inside the perimeter with
it](/enterprise): pull the cable, and the lab keeps testing.
