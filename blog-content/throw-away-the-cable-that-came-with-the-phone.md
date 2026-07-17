---
title: "Throw away the cable that came with the phone"
description: "The in-box phone cable is the most common invisible cause of a flaky device lab. A known-good, tested cable removes a whole class of debugging."
category: FIELD
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-09
sources:
  - title: "9to5Google — A bad USB-C cable fried a Google engineer's Chromebook (Feb 2016)"
    url: "https://9to5google.com/2016/02/03/bad-usb-c-cable-fries-chromebook/"
    note: "Benson Leung's USB-C cable-review campaign; one miswired SurjTech cable destroyed his Chromebook Pixel and two USB analyzers — proof that cables are not interchangeable."
  - title: "USB-IF — Cables and Connectors (certification program)"
    url: "https://www.usb.org/cable_connector"
    note: "The industry runs a cable-certification program precisely because, in its words, 'all cables do not have the same capabilities.'"
---

Early on, customers wired up their devices with the cable that came in the phone's box.
Reasonable — it's right there, it's the manufacturer's own, and it works. Then the lab
starts doing the thing that costs the most time in this business: failing
intermittently, for no reason anyone can see.

A device drops off its node mid-test. A build transfer that takes twenty seconds
suddenly takes four minutes. A phone that is plugged in, charging, and visibly on
simply isn't there as far as the node is concerned — until you unplug and replug it,
and then it is. Every one of these looks like a flaky test, or a flaky device, or a
flaky platform. Almost none of them is. The cause is thirty centimetres of copper
nobody thought to suspect.

## Why the cable is the worst kind of failure

It's invisible. When a test flakes, you look at the test. When a device misbehaves, you
look at the device, the OS, the automation. The cable is the last thing on the list,
because it "obviously" works — it charged the phone, didn't it? So you spend a day
ruling out everything expensive before you get to the cheapest part in the rack. It's
[the same trap as "it returned OK"](/blog/it-returned-ok): the signal you're reading
says success while the thing you care about quietly failed.

## Cables are not interchangeable, and the industry knows it

The idea that a cable is a cable is simply wrong, and there's a famous demonstration of
how wrong. In 2015-2016 a Google engineer, Benson Leung, ran a public campaign
reviewing USB-C cables for spec compliance and found a startling number that weren't —
including one miswired cable that physically **destroyed** the laptop he was testing
with, plus two of his analyzers, because its ground and power pins were crossed. That
episode is a big reason the USB-IF runs a cable-certification program at all; its own
page exists because, in plain terms, *"all cables do not have the same capabilities."*

That story is the dramatic extreme — a dangerously miswired cable is rare. The failure
a lab actually lives with is the quieter cousin: a perfectly ordinary cable that is
marginally out of spec, or was fine on day one and degrades. Which brings us to why
in-box cables in particular don't survive a lab.

## A consumer cable is not built for a lab's life

The cable in the box is engineered for a person who plugs in once a night. A lab does
something completely different to it: connect and disconnect thousands of times, hold a
stable data link and charge current for hours, run 24/7, for months. Connectors wear.
Marginal conductors that passed at room temperature on day one start dropping the data
line under sustained use. And because in-box cables aren't marked with which spec they
meet, you have a rack full of subtly different unknowns — the exact fragmentation
problem, [shrunk to the smallest part in the system](/blog/we-ship-the-lab-you-keep-the-phones).

So we don't use them. We ship known-good, tested cables as part of the lab and treat
them as a controlled, consumable part — the same make, to a known spec, replaced on a
schedule rather than when they fail. It sounds almost too small to matter. It removes
an entire category of "the lab is flaky" tickets, because the question *"is it the
cable?"* is answered before anyone has to ask it. The cheapest component in a device
lab causes a wildly disproportionate share of its failures — [precisely because it's
the last one you suspect](/platform/device-lab). Control it, and a whole class of
ghosts goes away.
