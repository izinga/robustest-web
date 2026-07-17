---
title: "This market eats companies"
description: "A billion-dollar fraud, a $4B valuation nobody has matched since, and a decade of vendors quietly renaming themselves. A short history of device testing, from someone still in it."
category: BUSINESS
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-16
sources:
  - title: "DOJ — Silicon Valley start-up founder sentenced to 18 months in prison for wire fraud and securities fraud"
    url: "https://www.justice.gov/usao-ndca/pr/silicon-valley-start-founder-sentenced-18-months-prison-wire-fraud-and-securities"
    note: "US Attorney's Office, Northern District of California. Describes HeadSpin's business and the sentence."
  - title: "BrowserStack closes $200 million in Series B funding at a $4 billion valuation"
    url: "https://www.browserstack.com/press/browserstack-closes-200-million-in-series-b-funding-at-a-4-billion-valuation"
    note: "16 June 2021, BOND-led. Still the category's high-water mark."
  - title: "Perforce completes Perfecto acquisition"
    url: "https://www.perforce.com/press-releases/perforce-completes-perfecto-acquisition"
    note: "Closed 3 December 2018."
---

We started building device labs in 2014. Most of the companies that were in this
market when we started are not in it now — not because they failed loudly, but
because this industry has a habit of absorbing people quietly.

Here is the record, as far as it is public.

## The peak is dated, and nobody has been back

On 16 June 2021, BrowserStack announced a $200 million Series B at a $4 billion
valuation, led by BOND. That is the largest number this category has ever printed.
Five years on, it still is. There has been no higher round, no up-round, no IPO.

That is not a criticism of BrowserStack — they built the thing everyone else gets
compared to, and "we're like BrowserStack but…" is still how half this market
introduces itself. It is an observation about the shape of the category. The peak
has a date on it, and the date is in the past.

## A billion-dollar fraud, inside this exact category

In 2023, the co-founder and former CEO of HeadSpin pleaded guilty to securities
fraud. In April 2024, he was sentenced to 18 months in prison. The Department of
Justice's own sentencing release, drawing on the plea agreement, describes the
company like this:

> HeadSpin provided clients with software tools and access to remote devices to
> test mobile applications.

That is our category, described by a federal prosecutor. The mechanism was the
oldest one there is: the revenue was not what the revenue was said to be.

We are not going to editorialise about a man who has been sentenced. The point
that matters for anyone buying testing infrastructure is narrower and less
satisfying: **a company can raise over a hundred million dollars, be valued near a
billion, and have its numbers turn out to be fiction — in a market where the
product is genuinely hard to evaluate from the outside.** If you cannot easily tell
whether a device cloud is doing what it says, neither can an investor.

## The quiet part: everyone else just changed their name

The fraud is the loud story. The pattern underneath it is duller and more telling.

In December 2018, Perfecto was absorbed into Perforce. Mobile Labs went into
Kobiton. S3 Group's StormTest went to Accenture. Eurofins Digital Testing was
divested and became Resillion. Testronic's digital TV arm left the family
altogether. OpenSTF — the open-source device farm that a generation of teams built
their labs on — has been effectively unmaintained for years, and the community fork
carries the load. In January 2026, LambdaTest stopped being LambdaTest and became
TestMu AI. Around the same time, TV Labs became Device.io. Rokuality and TestTrakt
still rank in Google; their domains no longer resolve.

None of that is scandal. It is just what a decade looks like in a market that is
harder than it appears from a pitch deck.

## Why it eats people

Our theory, after twelve years of it, is that device testing looks like software
and behaves like hardware.

Software businesses scale by adding instances. This one scales by buying physical
objects, mounting them somewhere, keeping them charged, keeping them cool, keeping
them on the right OS version, and replacing them when their batteries swell. Every
phone is a small, depreciating, slowly-failing asset with a USB cable attached to
it. The gross margin of a software company and the operational reality of a
warehouse do not sit comfortably together, and the gap between them is where
companies get into trouble — either by pretending the hardware isn't there, or by
raising against a multiple the physics won't support.

The vendors who survive tend to be the ones who made peace with being partly a
hardware operation.

## What we did with that

We stayed small, we stayed profitable, and we bet on the least fashionable thing
available: [putting the whole lab inside the customer's building](/enterprise) instead of ours.

For most of the last decade that looked like a mistake. Everyone was going the
other way, and the money agreed with them. We've since written about [that bet, and what it cost us to hold it](/blog/everyone-went-to-the-cloud).

But it is worth saying plainly, in a post about a market that keeps eating people:
we are still here, and the reason is not that we were cleverer. It is that we never
had to grow into a number someone else had already announced.
