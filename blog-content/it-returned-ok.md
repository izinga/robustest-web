---
title: "\"It returned OK\" is not \"it worked\""
description: "Our TV remote reported every keypress delivered. The TV received nothing. A debugging story, and why \"the write succeeded\" is one of the most dangerous phrases in distributed systems."
category: FIELD
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2025-11-18
---

We were [driving an LG TV over a WebSocket to its own input service](/platform/tv-testing) — arrow keys,
OK, back. On our side every keypress came back clean: sent, no error, logged as
delivered. On the TV: nothing moved.

Worse, it was inconsistent in a maddeningly specific way. The remote worked in the
first session after a restart, then went dead in the next. Within a session, if you
stopped pressing keys for a minute and came back, it had quietly stopped
responding — until you started a fresh session, at which point it worked again.

## Why "success" lies

A write to a socket succeeds when the bytes leave. That is all it means.

It does not mean the other end read them. It does not mean the other end is the
right other end. On a TCP connection that has been silently abandoned by the peer,
writes will keep succeeding into a buffer for a long time before anything
complains. Your logs will show a clean run of delivered keypresses to a socket
nobody is listening on.

This is not a TV problem. It is the shape of every "fire and check the return code"
bug there is. But TVs make it vivid, because the feedback channel — did the
highlight move? — is a physical panel two metres away, not an assertion.

## What was actually happening

LG's Magic Remote pointer runs on its own socket, separate from the main control
channel. We were opening a second one.

The TV accepts the second socket. It does not error. It does not close the first.
It simply stops acting on the frames from the older one — and the older one is the
one we were still writing to. Every write succeeded. Every keypress went nowhere.

The session-boundary pattern fell out of that immediately. A fresh session opened a
fresh socket and became the newest, so it worked. Any session that had been alive
long enough to have a second socket opened behind it was writing into the void.

## The fix

Two rules, both boring:

- **One pointer socket, ever.** Own it, reuse it, and make opening a second one
  impossible rather than merely discouraged.
- **A 15-second zero-delta keepalive.** A stale socket still accepts writes without
  error, so silence is not evidence of health. You have to keep saying something to
  find out you have been talking to yourself.

## The lesson worth keeping

We now treat any "delivered" log line on a channel we did not read back from as
approximately worthless. If the only proof that a command worked is that we
successfully sent it, we do not have proof — we have a hope with a timestamp on it.

[The camera watching the panel](/blog/the-camera-sees-what-the-framebuffer-cant) is not a nicety in that world. It is the only witness
that is not also the defendant.
