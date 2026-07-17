---
title: "Where we stopped"
description: "We reimplemented Samsung's SDB protocol from the daemon's own source, found the public documentation wrong, and then hit a wall we could not climb. This is the wall."
category: WHY
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-05-12
sources:
  - title: "Samsung — Smart Development Bridge (Tizen Studio common tools)"
    url: "https://github.com/Samsung/tizen-docs/blob/master/docs/application/tizen-studio/common-tools/smart-development-bridge.md"
    note: "The vendor's own description of sdb. Useful for the CLI, silent on the wire format."
---

To [drive a Samsung TV from a test lab](/platform/tv-testing) you need `sdb` — the Smart Development
Bridge, Samsung's fork of Android's `adb`. Samsung ships it as a binary. They do
not ship it for arm64, which is what our lab nodes run on.

So we wrote our own client — part of [why we implemented every TV wire protocol ourselves](/blog/why-we-wrote-our-own-drivers). This is a note about how far that got, and where it
stopped.

## The public documentation is wrong

`sdb` descends from `adb`, and the protocol document everyone reaches for is the
one bundled with `adb` from around 2011. It describes a 24-byte message header, a
magic field that is the command XORed with `0xFFFFFFFF`, and a payload checksum.

Three of those details do not survive contact with a Samsung TV.

The version field is not `adb`'s `0x01000001` — it is `0x02000000`. The maximum
payload is not the documented figure but 256 KiB. And the field the old document
calls a CRC is not a CRC at all: it is an unsigned sum of the payload bytes. If
you implement the published spec faithfully, the TV drops your connection and does
not tell you why.

We only know this because sdbd — the daemon on the TV — has readable source. Every
correction above came from reading it, not from guessing at a handshake until one
worked. We wrote the corrected description down as we went, with each claim
pointing at the file it came from, because the alternative was rediscovering it in
eighteen months.

## Then it stops

Getting a connection is the easy half. The hard half is authentication.

`adb` authenticates with RSA over a message type called `A_AUTH`. That is public,
implemented in a dozen open-source clients, and well understood. Samsung's TVs do
not do this. They negotiate over a different message type — `A_ENCR` — using a
handshake called SAKEP.

SAKEP is not in the sdbd source we can read. It lives in a proprietary plugin,
`_TV_REL`, that ships as a binary. You cannot implement from source what is not in
the source.

So: **on a TV that requires the secure handshake, our native client cannot connect,
and we fall back to Samsung's own `sdb` binary.** That is the wall. We did not get
past it. We are not going to pretend the fallback is a design choice.

## Why publish the wall

Because the useful part of this work is not "we built a thing." Plenty of people
build things. The useful part is a map with the cliff marked on it.

If you are about to reimplement SDB, you now know three things that would each have
cost you a day: the version and payload constants the public doc gets wrong, that
the checksum is a byte sum, and that there is a hard ceiling at the secure
handshake with a proprietary plugin behind it. That last one might save you a
month, because it is the kind of problem you can spend a month believing is your
bug.

The version of this post where we only describe the parts that worked would be a
better advertisement and a worse document. We would rather be the second one.
