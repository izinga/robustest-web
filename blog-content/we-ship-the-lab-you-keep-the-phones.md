---
title: "We ship the lab. You keep the phones."
description: "A device lab has two kinds of hardware: the devices under test, and the infrastructure around them. We standardize and supply the second — here's why."
category: WHY
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-17
sources:
  - title: "DeviceFarmer/stf — the maintained fork of OpenSTF"
    url: "https://github.com/DeviceFarmer/stf"
    note: "Its own README: 'development is still largely funded by individual team members and their unpaid free time, leading to slow progress.'"
---

A device lab has two kinds of hardware, and the whole operating model turns on
telling them apart.

There are the **devices under test** — the phones and TVs your users actually hold.
And there is the **infrastructure around them** — the nodes that drive each device,
the servers that run the platform, the capture hardware, the cables. When we set up a
lab, we supply the second and never the first. The devices stay the customer's:
their real fleet, their exact model list, their kit, [on their own
network](/platform/device-lab). We don't ship you phones and tell you what to test.

The infrastructure is the opposite. We supply it, and we make it identical
everywhere. Both halves of that rule were learned the expensive way.

## Fragmentation is a bug you didn't write and can't fix

Back in 2016 our device nodes were Intel-based Linux boxes, and we hit a
platform-level bug in the Intel USB host controller on them. After a handful of
connect/disconnect cycles, USB connections would drop at random: a phone that was
plugged in and working would simply fall offline mid-test, and only a controller reset
or a full reboot brought it back. It wasn't our code, our cables, or the phones — it
was the host silicon and its firmware, and there was nothing we could fix from our
side.

From where the customer sat, none of that mattered. The lab we installed dropped
devices at random, so the lab we installed was unreliable. We spent days proving the
fault lived in the host hardware itself — hardware in a configuration we couldn't
control — and that episode nearly cost us an enterprise relationship, over a bug that
was never ours.

The lesson wasn't "Intel is risky." It was that **every uncontrolled variable in the
infrastructure is a latent failure you will eventually own the blame for.** The only
way to not debug N different hardware configurations is to not have N. We supply one
known-good node configuration so there is exactly one to trust — and so the next
host-level bug, when it comes, is our problem to absorb, not a mystery in your rack.
It's a large part of why [the node is a standardized box we
choose](/blog/why-every-device-node-is-a-mac-mini), not whatever hardware happened to
clear procurement.

## Procurement is where deployments go to die

The second reason is less dramatic and, honestly, kills more projects. Enterprise
hardware purchasing is slow by design. When the node has to move through the
customer's procurement — a purchase order, a spec review, a vendor onboarding — the
lab stops. And the questions that stall it are maddeningly small: *why do you need
this much RAM, why is there no GPU, would a cheaper CPU do?* Each one is reasonable in
isolation and each one adds a week. We have watched deployments sit blocked for months
on sign-off for a machine that costs less than the delay.

When we supply the standardized node as part of the lab, that entire path disappears.
The deployment runs on our timeline, not on a purchasing queue's.

## Devices are exactly where variation *should* live

None of this means uniformity is the goal everywhere. The devices under test should be
gloriously varied — that variation is your users' reality, and pretending it away is
how bugs ship. The customer picks the exact models; we never substitute a "representative"
device for the real one.

The discipline is simply putting variation where it belongs and removing it where it's
pure downside. On the devices, variation is the whole point. On the infrastructure, it
is only ever a source of failures nobody budgeted for — which is the same reason the
DIY route is heavier than it looks. The main open-source lab stack (OpenSTF, now
maintained as DeviceFarmer) is real and still developed, but by its maintainers' own
words, "development is still largely funded by individual team members and their unpaid
free time, leading to slow progress." Running the infrastructure yourself means
[owning that maintenance tax forever](/blog/why-device-labs-get-abandoned).

So: we ship the lab, standardized down to the node. You keep the phones, in all their
messy, accurate variety. Each side owns the hardware it should — and neither of us
spends a week debugging the other's.
