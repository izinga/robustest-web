---
title: "Why there's an MDM inside our test lab"
description: "Nobody expects a certificate authority in a device farm. We built one because enrolling iOS behind a load balancer breaks in a way you miss until production."
category: WHY
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-03-24
sources:
  - title: "micromdm/scep — SCEP server and depot implementations"
    url: "https://github.com/micromdm/scep"
    note: "Ships filesystem and BoltDB depots; both are local-disk."
  - title: "RFC 8894 — Simple Certificate Enrolment Protocol (SCEP)"
    url: "https://www.rfc-editor.org/rfc/rfc8894.html"
---

To install an enterprise build on a supervised iPhone without a human tapping
through prompts, the device has to be enrolled in an MDM. To be enrolled, it needs a
certificate. To get a certificate, something has to sign it. That something is a
certificate authority — and in a test lab, that is a stranger thing to own than it
sounds.

This is a note about why [a device farm](/platform/device-lab) ended up with a CA in it, and the specific,
non-obvious way the obvious approach fails.

## The obvious approach

There is a good open-source SCEP server — `micromdm/scep` — and the standard way to
run MDM enrolment. SCEP is the protocol a device speaks to request its certificate;
the server hands back a signed one. Point your devices at it, and enrolment works.

On one machine, it works perfectly.

## Where it breaks

`micromdm/scep` ships two ways to store the CA and the certificates it has issued:
a filesystem depot and a BoltDB depot. **Both are local disk.**

That is completely fine until you run more than one instance of the server, which any
real lab does — you put the enrolment endpoint behind a load balancer for
availability. Now you have a problem that does not show up in testing and does not
show up on the first day in production:

**Each replica has its own CA.**

Replica A generates a CA on its disk. Replica B generates a different CA on its disk.
A device that happens to enrol through A gets a certificate signed by A's CA. Later,
when that device checks in and the load balancer routes it to B, B looks at the
certificate, does not recognise the authority that signed it, and rejects it. The
device was enrolled correctly. It is now un-enrollable, intermittently, depending on
which replica answers.

It is the worst class of bug: correct on one node, correct in a demo, and broken only
under the exact conditions you built the second node for.

## What we did instead

We moved the depot into the database the rest of the fleet already uses. The CA
certificate, the issued-certificate records, and the serial counter all live in
MongoDB, which every replica shares.

Now there is one CA. A device enrolled through any replica gets a certificate every
replica recognises, because they are all reading from the same place. The load
balancer becomes what it was supposed to be — a detail — instead of a source of
intermittent enrolment failures.

One piece of it is worth calling out, because it is the part that has to be right
under concurrency: serial numbers. Every certificate a CA issues needs a unique
serial. With replicas issuing certificates at the same time, "read the highest serial,
add one" is a race — two replicas read the same number and mint two certificates with
the same serial. We mint serials with an atomic increment in the database instead, so
uniqueness holds across every replica without any of them coordinating or locking.

## Why this counts as lab infrastructure, not a side quest

Because the whole reason the lab exists is to run unattended. A tester should book a
supervised device, push an enterprise build, and have it install — no trust prompts,
no babysitting. That only works if enrolment is boring and reliable, and enrolment is
only boring and reliable if the CA is singular and shared.

The visible feature is "iOS devices take enterprise builds without anyone tapping
through dialogs." The unglamorous thing that makes it true is a certificate authority
that survives having more than one copy of the server in front of it. [Most of running a lab is like that](/blog/why-device-labs-get-abandoned): the feature is what you see, and the engineering is the failure
mode you made sure nobody ever hits.
