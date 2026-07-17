---
title: "Why every device node is a Mac mini"
description: "iOS testing can only run on a Mac — Apple's license requires it. So the node driving your iPhones and Androids is a Mac mini. No compromise."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-17
sources:
  - title: "Appium — XCUITest driver system requirements"
    url: "https://appium.github.io/appium-xcuitest-driver/latest/getting-started/system-requirements/"
    note: "Verbatim: 'There are three primary requirements to use the XCUITest driver: macOS host machine, Xcode, Appium.' Windows/Linux hosts have only limited support."
  - title: "Apple — macOS Software License Agreement (Tahoe 26)"
    url: "https://www.apple.com/legal/sla/docs/macOSTahoe.pdf"
    note: "§2.A licenses one copy 'on a single Apple-branded computer'; §2.J forbids non-Apple hardware; VMs are permitted only on Apple hardware. macOS legally cannot run on a generic server."
---

Every device node in a RobusTest lab is a Mac mini — the same small silver box for the
rack that drives your iPhones, and the one that drives your Androids. That isn't an
Apple preference. It's the cleanest answer to a constraint you cannot design your way
around.

## You cannot test iOS without a Mac. Two reasons, both hard.

The **tooling** reason: Apple ships Xcode — and with it the signing, building, and
instruments that iOS automation depends on — for macOS only. Appium's own XCUITest
driver documentation puts it plainly: *"There are three primary requirements to use the
XCUITest driver: macOS host machine, Xcode, Appium."* Windows and Linux hosts have, in
Appium's words, only limited support. There is no supported path to driving a real
iPhone from a Linux box.

The **legal** reason: even if you found a way to run macOS elsewhere, Apple's license
forbids it. The macOS Software License Agreement grants you one copy *"on a single
Apple-branded computer,"* explicitly bars running it on non-Apple hardware, and permits
virtual machines only on Apple hardware. macOS on a generic server isn't a technical
hack around a rule — it *is* the rule you'd be breaking.

Put together: a lab that tests iOS needs real Macs. That part isn't a decision. The
only decision is *which* Mac.

## The mini turns the requirement into an advantage

Once you accept you need Macs anyway, the Mac mini stops being a tax and starts being
the obvious node — because it does more than satisfy the iOS requirement.

- **One node type, both platforms.** A Mac runs the Android toolchain — ADB, Appium —
  perfectly well. So the same box that drives iPhones drives Androids too. Instead of a
  split fleet (Linux hosts for Android, Macs for iOS), the lab has one kind of node to
  provision, image, and support.
- **Built for the rack already.** Small, low-power, quiet, and mountable — it slots
  into lab infrastructure without a custom chassis or a workstation's footprint.
- **First-class support for the thing you're testing against.** The node lives in the
  same ecosystem as the devices and the OS releases it has to keep up with, with the
  longevity and update cadence Apple hardware gets.

It's a genuinely good node on its own merits: simple, small, reliable. The fact that it
also happens to be the only legal way to test iOS is what makes it the whole answer
instead of half of one.

## The general lesson

There's a pattern here we keep applying. A constraint you can't remove is best turned
into a standard you lean on. We could have fought the Mac requirement — chased macOS
VMs on foreign hardware, run two separate node fleets, treated iOS as a bolt-on. Each
of those trades a clean rule for a permanent source of complexity. Standardizing on the
one box that satisfies the constraint *and* pulls double duty is why [every node in the
lab is identical](/blog/we-ship-the-lab-you-keep-the-phones) — and why "which server
should this be?" is a question we only had to answer once. It's a Mac mini, [and it's
driving your real devices, not a simulator](/platform/device-lab).
