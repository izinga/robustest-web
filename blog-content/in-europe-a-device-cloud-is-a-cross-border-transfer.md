---
title: "In Europe, a device cloud is a cross-border transfer"
description: "GDPR conditions every move of EU personal data outside the EEA, and the EU–US basis has been struck down twice. A device cloud inherits all of it."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-10
sources:
  - title: "GDPR — Chapter V (Articles 44–50), transfers to third countries"
    url: "https://gdpr-info.eu/chapter-5/"
    note: "Art. 44: personal data may leave the EEA only under Chapter V conditions, and the level of protection must not be undermined. Art. 45: free transfer only to countries with an EU 'adequacy' decision."
  - title: "CJEU — Schrems II (Case C-311/18, 16 July 2020)"
    url: "https://www.europarl.europa.eu/RegData/etudes/ATAG/2020/652073/EPRS_ATA(2020)652073_EN.pdf"
    note: "Invalidated the EU–US Privacy Shield: US surveillance law (e.g. PRISM/UPSTREAM) was not 'essentially equivalent' to EU protection. SCCs survive but need case-by-case 'supplementary measures.'"
  - title: "IAPP — General Court dismisses Latombe challenge, upholds the Data Privacy Framework (Sept 2025)"
    url: "https://iapp.org/news/a/european-general-court-dismisses-latombe-challenge-upholds-eu-us-data-privacy-framework"
    note: "The 2023 EU–US Data Privacy Framework survived its first challenge, but is now under appeal at the CJEU (Case C-703/25 P, Oct 2025)."
---

*This is the Europe edition of a short series on data residency and the test lab. The
[India edition is here](/blog/data-residency-reaches-the-test-device), and the
[US edition here](/blog/the-us-has-no-data-residency-law).*

European data-protection law does not forbid personal data from leaving the EEA. It does
something more demanding: it **conditions** every such move, and puts the burden on you
to justify it. A device cloud that runs outside the EEA — or is operated by a company
that can be compelled to hand data to a foreign government — inherits that whole burden,
whether your test plan meant to take it on or not.

## What GDPR actually requires

Chapter V of the GDPR (Articles 44–50) governs transfers of personal data to "third
countries." The general principle in Article 44 is that data leaving the EEA must keep
an essentially equivalent level of protection — the transfer "shall not undermine" what
the GDPR guarantees. In practice that means one of three things: the destination has an
EU **adequacy decision** (Article 45); or you put **appropriate safeguards** in place,
typically Standard Contractual Clauses (Article 46); or you fall back to a narrow
**derogation** (Article 49). Ship EU users' personal data to a device cloud in a
non-adequate country with none of these, and the transfer is simply unlawful.

## The US basis has collapsed twice, and is being challenged a third time

The hardest case is the most common one: a US-headquartered device cloud. Transfers of
EU personal data to the US have rested on three successive frameworks, and the pattern
is not reassuring:

- **Safe Harbor** — struck down by the CJEU in 2015 (Schrems I).
- **Privacy Shield** — struck down on 16 July 2020 (Schrems II, Case C-311/18), because
  US surveillance law was found not "essentially equivalent" to EU protection.
- **Data Privacy Framework** — the current basis, adopted July 2023. It survived its
  first challenge in September 2025, but is now on appeal at the CJEU (Case C-703/25 P).

Building your test pipeline on the framework of the day means building on ground the
same court has given way under twice — and is being asked to rule on a third time. When
a US cloud tells you "we're certified, transfers are fine," it is describing a status
that its two predecessors also held right up until they didn't.

## Why this lands on the test lab

None of this is about production alone. A test run generates an app binary, screen
recordings, logs, and network captures — and for an app used by EU residents, those
artifacts routinely contain personal data: names, emails, locations, account details on
screen. Run that test on a device cloud outside the EEA and you have made a Chapter V
transfer of that personal data, and [it sits on that provider's infrastructure for as
long as they retain it](/blog/what-actually-leaves-your-network). Your DPO now owns a
transfer to assess, safeguards to document, and a legal basis that may not survive the
next ruling.

## Why on-premise answers it cleanly

If the device lab runs inside the EEA and inside your own network — the devices, the builds,
and the captured data never crossing a border — there is no Chapter V transfer to
justify. No adequacy question, no SCCs to paper, no supplementary measures to argue, no
framework whose validity you have to track in the legal press. The data stayed where the
regulation wanted it. That is why teams under real GDPR exposure tend to want [the lab,
and everything it captures, inside their own perimeter](/enterprise) — and why "can this
device cloud actually run in our own data centre" is [the question that narrows the field
fast](/blog/which-device-clouds-run-in-your-data-centre).

## The honest boundary

Transfers can be made lawful — SCCs plus genuine supplementary measures do work, and
fully anonymised or [synthetic test data](/blog/synthetic-test-data-is-a-trade-off) sidesteps GDPR's scope altogether. But real test
environments touch real personal data, and "we have the SCCs on file" is an ongoing
burden of assessment, audit, and legal-risk tracking that on-premise testing simply does
not create. The cleanest transfer to defend is the one you never made.
