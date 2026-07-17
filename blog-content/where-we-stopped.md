---
title: "Where we stopped"
description: "We reimplemented Samsung's SDB protocol from the daemon's source, found the public docs wrong, then hit a wall we couldn't climb. This is the wall."
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
one bundled with `adb` from around 2011. Follow it faithfully against a Samsung TV
and the connection silently drops — no error, no explanation, no clue that the
problem is the spec and not your code.

The reason is that several of the fields the old document describes — the version
it advertises, the maximum payload it allows, the way it checksums a message — are
simply not what a Samsung TV expects. The public map is wrong in ways it gives you
no way to detect from the outside. You will spend a day certain the bug is yours.

We know where it diverges because sdbd — the daemon on the TV — has readable
source, and we read it rather than guessing at a handshake until one happened to
work. (We're keeping the exact corrections to ourselves — the point here is that
the published map has unmarked cliffs on it, not to hand over ours.)

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

## Why this is worth telling

Not to hand over the fixes — those we're keeping. It's the *shape* of the problem
that's worth seeing. Reimplementing a vendor's own bridge protocol, from a daemon's
source, only to dead-end at a proprietary encryption plugin you cannot implement
from source, is the kind of multi-month detour that "we'll just build our own TV
lab" quietly signs you up for. We went down it so that driving a Samsung TV is a
solved problem on our side of the line — and a wall you never have to meet on yours.

That's the honest version of this post. Not a recipe. A marker on a map: here is how
deep the water gets, and here is the one spot even we had to turn back.
