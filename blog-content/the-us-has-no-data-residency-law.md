---
title: "The US has no data-residency law. Its rules are stricter."
description: "No US law says keep data in-country. But ITAR, FedRAMP and CJIS demand US soil and US-persons-only access — which a shared device cloud can't give."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-12
sources:
  - title: "Braumiller Law — Storing export-controlled data in the cloud"
    url: "https://www.braumillerlaw.com/storing-export-controlled-data-in-the-cloud-whats-the-latest/"
    note: "ITAR/EAR technical data must stay on US-located systems accessible only to US persons; foreign-person access — even by IT admins — is a 'deemed export.' Cloud is allowed only with strong encryption where no foreign person holds the keys."
  - title: "AWS GovCloud (US) — Compliance (FedRAMP, CJIS, ITAR, EAR)"
    url: "https://docs.aws.amazon.com/govcloud-us/latest/UserGuide/govcloud-compliance.html"
    note: "Why a separate US region exists at all: government/regulated workloads require US-person administration and in-country data residency that a normal commercial cloud doesn't guarantee."
---

*This is the US edition of a short series on data residency and the test lab. See also
the [India edition](/blog/data-residency-reaches-the-test-device) and the
[Europe edition](/blog/in-europe-a-device-cloud-is-a-cross-border-transfer).*

The United States is the country everyone assumes has no data-residency rules, and in
one narrow sense that's true: there is no single federal law that says "keep your data
in the country." It is also the most misleading true statement in this whole topic —
because several US regimes are *stricter* than a plain localization law. They don't just
want the data in the country. They want it on US soil, touchable only by US persons, and
they attach criminal penalties when it isn't.

## The rules that actually bite

**Export control — ITAR and the EAR.** This is the sharp one. Controlled technical data
must be stored on systems physically located in the United States and accessible only by
"US persons." A foreign national touching it — including an IT administrator or a
contractor — is a **"deemed export,"** treated as an export to their home country, even
if they never leave the building. Allowing access from outside the US, including through
a cloud reachable abroad, is an export outright. Cloud storage is permitted only with
strong (FIPS-validated) encryption where no foreign person can hold the keys. Get it
wrong and the exposure is civil and criminal, not a fine you write off.

**Federal and criminal-justice data — FedRAMP and CJIS.** Cloud services handling
federal-agency data must be FedRAMP-authorized; criminal-justice information falls under
the FBI's CJIS Security Policy. The reason cloud vendors run entirely separate
"GovCloud" regions is precisely this: those workloads require US-person administration
and in-country data residency that an ordinary commercial cloud doesn't promise.

**Health data — HIPAA.** HIPAA doesn't mandate US storage, but it does demand strict
safeguards and makes anyone processing protected health information a "business
associate," contractually and legally on the hook. A device cloud holding PHI-laden
screen recordings is a business associate you now have to paper and trust.

## Why this lands on the test lab

Look at the shape those rules require — US soil, US-person access only, no offshore
administration — and then look at the shape of a public device cloud: data centres
placed wherever is cheapest, and staff administering them from around the world. Those
two shapes are opposites.

Test a defense-adjacent app, a federal-contractor app, or a health app on a shared
global device cloud, and the app binary, the screen captures, and the logs land on
infrastructure with foreign administrators and offshore data centres — the exact
condition export control calls a deemed export and FedRAMP/CJIS simply forbid. A test run
is [a lot of data leaving your control](/blog/what-actually-leaves-your-network), and for
these regimes "leaving your control" is the whole violation.

## Why on-premise answers it cleanly

If the device lab runs inside your own US facility, on your own network, administered by your
own cleared staff — the devices, the builds, and the captures never leaving the
building — export-controlled data was never exported, no foreign person could access it,
and there is no third-party business associate to vet. The US-soil and US-persons-only
requirements aren't safeguards you bolt on; they're just true, because nothing moved.
That is why regulated and defense-facing teams want [the whole lab inside their own
perimeter](/enterprise), and why the practical filter is [which device clouds can even
run in your data centre at all](/blog/which-device-clouds-run-in-your-data-centre).

## The honest boundary

Compliant cloud paths do exist: FedRAMP-authorized GovCloud regions, and ITAR storage
under FIPS-validated encryption with US-held keys. If your device-testing vendor offers
one of those, use it. But most commercial device clouds are none of those things — not
FedRAMP-authorized, not ITAR-scoped, administered globally — and for export-controlled,
federal, or criminal-justice data, that's disqualifying, not a detail. On-premise skips
the entire question: the strictest rule in the set is satisfied by the data simply never
leaving your walls.
