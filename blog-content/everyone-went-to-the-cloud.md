---
title: "Everyone went to the cloud. We didn't."
description: "For most of a decade, building an on-premise device lab looked like a mistake. What the bet cost — and why the argument has quietly changed underneath it."
category: BUSINESS
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-06-24
sources:
  - title: "37signals — Our cloud exit savings will now top ten million over five years"
    url: "https://world.hey.com/dhh/our-cloud-exit-savings-will-now-top-ten-million-over-five-years-c7d9b5bd"
    note: "Self-reported, gross of colo and power. Their 2024 figures, not an analyst's."
  - title: "Wayback Machine — BrowserStack /private-devices, 13 November 2025"
    url: "http://web.archive.org/web/20251113212615/https://www.browserstack.com/private-devices"
    note: "Snapshot showing the on-premise wording that is no longer on the live page."
---

Every year for about ten years, someone has explained to us that on-premise is
over. Usually kindly. Occasionally with a slide.

They were describing a real trend. The device-testing market went to the public
cloud almost completely, and the reasons were good ones: no hardware to buy, no
racks to cool, a credit card instead of a procurement cycle. We watched
competitors raise against that story and we did not have a better story. We had a
different one.

## What the bet actually was

Not "the cloud is bad." The cloud is extremely good at most things.

The bet was narrower: **that a meaningful number of companies would never be
allowed to send their unreleased app to someone else's data centre, and that
serving them properly meant putting the entire lab — control plane, devices,
storage — inside their building.**

Not devices on their premises with the brain in our cloud. The whole thing. It is
a harder product to build, a harder product to install, and it makes you a
services business whether you like it or not.

## What it cost

Growth, mostly.

You cannot run [an on-premise lab](/platform/device-lab) as a self-serve trial. Nobody swipes a card and
gets a device lab in their data centre by Tuesday. Every install is a
conversation, then a room, then hardware, then people. That is a slower curve than
a signup form, and for a decade the market rewarded the signup form.

It also cost us the easy comparison. When a buyer lines vendors up in a
spreadsheet, "deployed in your data centre" does not fit in the column next to a
per-minute price.

## What changed

Two things, neither of which we predicted.

The first is that running your own infrastructure stopped being embarrassing.
37signals published their numbers: a cloud bill taken from a $3.2 million a year
run rate down to $1.3 million, with about $700,000 of hardware paid off out of half
a year's savings. Those are their figures, self-reported, and gross of the colo and
power they do not itemise — take them as one company's honest accounting rather
than a law of nature. But the taboo broke. It became sayable that owning the
machine can be cheaper than renting it.

We want to be careful here, because this is where the argument usually gets
oversold. You will see a statistic quoted that 83% of enterprises are repatriating
from the public cloud. That number comes from a CIO *sentiment* survey, and what it
actually counts is CIOs who **plan** to move **at least some** workloads. It is not
a mass exodus, and anyone selling you on-premise by waving it is hoping you do not
check. The honest version is smaller and still enough: repatriation is real,
case by case, where the economics or the rules make it real.

The second thing is quieter. Sometime between November 2025 and now, BrowserStack's
private-devices page stopped saying its devices could be "deployed on-premises or
in the cloud." The Wayback snapshot still has the old wording. The live page does
not. We are not going to guess at their reasoning — plans change, and a marketing
page is not a roadmap. But we noticed, because we have spent twelve years being the
only people in the room arguing for that sentence.

## Where this leaves us

Still holding it. The difference is that we now get asked about it by people who
already know why they are asking — [a bank whose security review said no](/blog/which-device-clouds-run-in-your-data-centre), a broadcaster who [cannot post an unreleased build to anyone's cloud](/blog/what-actually-leaves-your-network), a services firm
that just read an RFP requiring a lab it does not have.

That is not vindication. Vindication would be the market agreeing with us, and it
mostly still doesn't. It is just that the people who need this now know they need
it, and for a long stretch they didn't know it was a thing you could buy.
