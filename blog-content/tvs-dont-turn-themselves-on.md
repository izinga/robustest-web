---
title: "TVs don't turn themselves on"
description: "A phone on USB is never off. A TV in a rack goes to standby, and neither Tizen nor webOS can power on over its control protocol. What a lab does instead."
category: TV
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-08-24
sources:
  - title: "Wikipedia — Wake-on-LAN"
    url: "https://en.wikipedia.org/wiki/Wake-on-LAN"
    note: "The magic packet: \"6 bytes of all 255 ... followed by sixteen repetitions of the target computer's 48-bit MAC address, for a total of 102 bytes\"; sent to UDP port 0, 7 or 9; \"can generally only be used within a subnet.\" Origin: AMD with HP, proposed in a November 1995 white paper."
  - title: "LG — Developer Mode app"
    url: "https://webostv.developer.lge.com/develop/getting-started/developer-mode-app"
    note: "Why a sleeping webOS set is a liability beyond the session that can't start: the Developer Mode session clock keeps running whether or not you are using the TV."
---

A phone in a device lab is never off. It hangs on a USB cable, the cable feeds it, and
the control channel is up the whole time. If a session starts at 3 a.m. the phone is
there to take it. Every assumption in a phone lab — reservation queues, health checks,
"is this device online" — rests on that.

A TV is not like that. A TV in a rack goes to standby the way it does in a living room:
someone presses power, or it times out, or the firmware decides. And once it is in
standby, the control channel you drive it through is gone.

## You can turn it off. You cannot turn it on.

This is the part nobody expects. LG's webOS control protocol has a *turn off* call.
Samsung's Tizen has the equivalent. Neither has a *turn on* — not a hidden one, not an
undocumented one. The socket you would send the command over lives in the part of the
TV that just went to sleep. The comment in our own driver says it plainly: "there's no
SSAP power-on — use Wake-on-LAN."

So a TV lab needs a second channel that works when the first one is dead. There is
exactly one candidate, and it is thirty years old.

## A 102-byte packet from 1995

Wake-on-LAN was proposed by AMD and HP in 1995 and has not changed since. The "magic
packet" is six bytes of `0xFF` followed by the target's MAC address repeated sixteen
times — 102 bytes, sent as a UDP broadcast, conventionally to port 9. The network
interface stays powered in standby, sees its own MAC sixteen times in a row, and wakes
the machine.

Three things about this shape a TV lab more than you would guess.

**You need the MAC, and only an awake TV will tell you.** The packet is addressed by
hardware address, not IP. We learn each set's MAC from the TV's own device information
while it is online and keep it with the device record. A set that was racked but never
brought online has no MAC on file, and there is nothing to wake it with — which is why
onboarding a TV in the lab is "connect it, read it, *then* let it sleep", not the other
way round.

**It is a broadcast, so it does not cross subnets.** Wake-on-LAN "can generally only be
used within a subnet." That is a network-design constraint: the node that wakes a TV
must sit on the same broadcast domain as the TV. It is one of the reasons a TV wall is
[one small node per wall](/platform/tv-testing), on the same switch as its panels,
rather than a controller somewhere else in the building. (Linux also refuses to send to
`255.255.255.255` unless you set `SO_BROADCAST` on the socket — an hour we will not get
back.)

**The TV has to agree.** WoL is a setting on the set, off by default on some models,
under a different menu name on every brand. A TV that was never configured for it
ignores the packet, silently. That goes on the onboarding checklist too.

## What the lab does with it

Every session starts by sending the magic packet whether or not the TV looks awake —
our cached status can be stale, and the packet is harmless to a set that is already on.
Then the node re-probes until the TV answers, about thirty seconds worst case, so the
install step does not fire at a panel that is still booting. The same path is wired to
the *Power* key in the live remote, because the remote socket is dead while the panel
sleeps and that is the one key that cannot ride it.

And where the platform allows it, we avoid the whole problem: webOS can blank the
screen while keeping the TV powered and reachable, so a wall can run dark overnight
without ever crossing into standby. That is the difference between a lab that is ready
at 3 a.m. and one that has to be woken up first.

## Why it matters beyond the packet

A sleeping TV is not just unavailable; it is a liability. LG's Developer Mode session
clock does not stop because the set is asleep — the panel can lapse out of developer
mode and [uninstall your apps](/blog/webos-26-field-guide) while nobody is watching. A
lab that cannot reliably wake its sets cannot reliably keep them in session either.

It is the same lesson every TV in the rack keeps teaching: the platform reports success
at the API level while the thing you care about — [is the panel actually
there](/blog/it-returned-ok) — is a separate question. Phones let you skip that
question. Televisions never do.
