---
title: "Your TV test passed. There was no sound."
description: "A video-only check on a TV app can pass while the audio is silent, wrong, or out of sync — the failures users notice first. Audio is its own test signal."
category: TV
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-08
sources:
  - title: "EBU Recommendation R37 — The relative timing of sound and vision"
    url: "https://tech.ebu.ch/publications/r037"
    note: "Broadcast tolerance for audio/video sync: sound should not lead vision by more than ~40 ms nor lag by more than ~60 ms — beyond that, viewers notice."
  - title: "Xiph.Org — Opus audio codec"
    url: "https://opus-codec.org/"
    note: "A common codec for streamed audio; captured TV audio arrives encoded and must be decoded correctly before it can be verified."
---

Almost every app on a television exists to play something you can hear. Yet the
default way teams test TV apps checks only the thing you can see: did the screen
change, is the picture there, did the video start. A test built entirely on the
picture will report a clean pass while the app is playing to an empty room.

Audio is not a detail of the video test. On a TV it is a separate signal, and it fails
in ways the picture never reveals.

## The failures a picture can't show

Three of them, all real, all invisible to a screenshot:

- **Silence.** The video renders perfectly and there is simply no sound — a muted
  audio track, a failed audio-focus handoff, a decoder that never started. The frame
  looks correct. The experience is broken.
- **The wrong audio.** Playback is fine but on the wrong track: the default language
  instead of the user's, commentary instead of the main mix, an ad's audio bleeding
  over the content. The picture is right; the sound is wrong.
- **Drift.** Audio and video start together and slowly separate. Broadcasters treat
  this as a hard limit — EBU R37 puts the tolerance at roughly *sound no more than
  40 ms ahead of vision, no more than 60 ms behind* — because past that, every viewer
  notices the lips not matching. A picture-only test cannot see a millisecond of it.

Each of these is the kind of failure a user reports in the first ten seconds and a
green test suite never mentions. It is the [same trap as "it returned
OK"](/blog/it-returned-ok): the check you ran succeeded, so you believe the thing you
cared about did too.

## Why audio is harder to check than video

You cannot assert on TV audio the way you eyeball a screenshot. The sound has to be
captured as its own stream, decoded from whatever codec it arrives in — Opus, AAC, and
others are all in play — and then actually evaluated: is there signal or silence, is it
on the expected channels, does it line up in time with the picture. Skip the decode
step, or assume the wrong container, and you get silence that reads as success — which
is worse than no test at all, because it is a test that lies.

None of that is exotic once you treat audio as a first-class signal. It is only a trap
if you bolt it on as an afterthought to a video test.

## The point

Verifying a TV app means verifying what comes out of the speakers, not just what lands
on the panel. That is why a real TV session in our lab [captures the actual audio
alongside the live screen](/platform/tv-testing), so a tester hears what the device is
playing and a check can assert on it — and why "the video played" is never, on its own,
allowed to stand in for "the app worked." On a device built to be heard, a silent pass
is not a pass.
