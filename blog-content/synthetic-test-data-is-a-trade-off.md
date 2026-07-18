---
title: "Synthetic test data is a trade-off, not an escape hatch"
description: "Synthetic or masked data can move test data out of GDPR scope — but only if it's truly anonymous, and the more you sanitize it, the fewer bugs it catches."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-14
sources:
  - title: "GDPR — Recital 26 (not applicable to anonymous data)"
    url: "https://gdpr-info.eu/recitals/no-26/"
    note: "Only data rendered anonymous 'in such a manner that the data subject is not or no longer identifiable' is out of scope, judged by 'all the means reasonably likely to be used.' Pseudonymised data remains personal data."
  - title: "DataMasque — Risks of using production data in test environments"
    url: "https://www.datamasque.com/blog/risks-of-using-production-data-in-test-environments"
    note: "Non-production environments are a leading site of PII exposure — test data is widely copied, shared with external teams, and less controlled than production."
---

In each of the three [data-residency](/blog/data-residency-reaches-the-test-device)
posts I conceded the same escape hatch: *if you test only with synthetic data, the
calculus softens.* That's true. It is also worth being honest about how much — because
it softens less than it sounds, for two reasons that pull in opposite directions.

## What actually takes data out of scope

Under the GDPR, the bar is specific and high. Recital 26 says the rules don't apply to
data "rendered anonymous in such a manner that the data subject is not or no longer
identifiable" — and identifiability is judged against "all the means reasonably likely
to be used" to single someone out, accounting for cost, time, and available technology.
Clear that bar and the data is genuinely out of scope.

Here is the catch most teams miss: **pseudonymised data is still personal data.** If a
mapping exists anywhere that could re-link the records — a key, a lookup table, a
reversible masking scheme — you have pseudonymised, not anonymised, and it remains
fully in scope. A great many "we masked the test data" setups are exactly that, which
means they never actually left the regulation. And the sectoral localization rules —
[India's RBI and SEBI](/blog/data-residency-reaches-the-test-device), for instance —
don't offer the GDPR's clean "anonymous data is out" carve-out at all. Synthetic data
helps most under a privacy law that has that door, and least under a hard localization
mandate that doesn't.

## What you give up to clear it

The reason to test on real devices with realistic data is to catch what only appears at
production's data shapes: the 300-character company name, the account created in a
timezone nobody tested, the payment that partially succeeded, the referential-integrity
tangle three tables deep. Synthetic data, by nature, sticks to the routes and scenarios
someone thought to generate. It rarely reproduces the sheer mess of how real people use
software.

So there's a direct tension. The more thoroughly you sanitise the data to move it out of
scope, the more you strip out the exact distributions that expose the bug. Sanitise it
enough to be safely anonymous, and you have often sanitised away the failure you were
hoping to find. (In fairness, synthetic data cuts the other way too — it can be
*configured* to inject the systematic edge cases real data lacks: maximum-length strings,
Unicode, nulls, boundary dates. It's not strictly worse; it's different, and it misses a
different set of things.)

## The part that's backwards

"It's only test data" is the phrase that gets this exactly wrong. Non-production
environments are one of the *highest*-risk places for sensitive data, not the lowest.
Production access is locked down and audited; test schemas are frequently readable by
whoever has database access, copied freely, shared with external and offshore teams, and
increasingly scraped to fine-tune internal AI models. Test data sitting on a third-party
device cloud is not a lesser exposure than production — routinely, it is a greater one.

## Why on-premise dissolves the trade-off

Notice that the whole tension only exists because you're trying to keep data *off
someone else's infrastructure.* Run the lab [inside your own
network](/platform/device-lab) and it simply disappears. You can test with realistic,
production-shaped data *and* keep it within your walls, because there is no third party
and no border for it to cross. You are no longer choosing between a faithful test and a
compliant one — the same way you don't choose between [a real device and a
convenient](/blog/simulators-are-good-enough) one when the real device is right there in
the rack.

Synthetic data is a genuine tool; use it, especially to manufacture edge cases real data
won't. But treat it as a trade-off between fidelity and compliance, not a free pass out
of residency. The setup that never makes you trade is the one where the realistic data
never leaves in the first place.
