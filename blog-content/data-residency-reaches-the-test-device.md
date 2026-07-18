---
title: "Data residency reaches all the way to the test device"
description: "India's RBI, IRDAI and SEBI rules keep regulated data in India — and a third-party device cloud is a way it leaves. Why regulated apps get tested on-prem."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-07-18
sources:
  - title: "RBI — Storage of Payment System Data (FAQ + April 2018 directive)"
    url: "https://www.rbi.org.in/commonman/english/scripts/FAQs.aspx?Id=2995"
    note: "Circular DPSS.CO.OD.No 2785/06.08.005/2017-18, 6 Apr 2018: 'the entire data relating to payment systems… is stored in a system only in India.'"
  - title: "SEBI — Framework for Adoption of Cloud Services by Regulated Entities"
    url: "https://www.sebi.gov.in/legal/circulars/mar-2023/framework-for-adoption-of-cloud-services-by-sebi-regulated-entities-res-_68740.html"
    note: "Circular SEBI/HO/ITD/ITD_VAPT/P/CIR/2023/033, 6 Mar 2023: RE data stored within India; the regulated entity remains accountable for its data and logs on a provider's cloud."
  - title: "IRDAI (Maintenance of Insurance Records) Regulations, 2015 — data-localisation requirement"
    url: "https://digitalpolicyalert.org/event/8796-adopted-irdai-maintenance-of-insurance-records-regulations-2015-including-data-localisation-requirement"
    note: "Records of insurance policies and claims made in India must be held in data centres located and maintained in India."
---

*This is the India edition of a short series on data residency and the test lab. See also
the [Europe edition](/blog/in-europe-a-device-cloud-is-a-cross-border-transfer) and the
[US edition](/blog/the-us-has-no-data-residency-law).*

For most engineering teams the security bar on a test lab is "keep the data safe." For
an Indian bank, insurer, or broker, the bar is stricter and far more specific: keep
certain data **in India**, and stay **accountable** for it — and that requirement does
not stop at your production servers. It reaches the device you run your tests on.

## What the rules actually say

Three regulators, three overlapping mandates, all of them about where regulated data is
allowed to live.

**Banking and payments — RBI.** The Reserve Bank's April 2018 directive on Storage of
Payment System Data (circular DPSS.CO.OD.No 2785/06.08.005/2017-18) is blunt: *"the
entire data relating to payment systems… is stored in a system only in India."* There
is a narrow allowance for the foreign leg of a cross-border transaction, but even then
the data must be *"deleted from the systems abroad and brought back to India not later
than… 24 hours from payment processing."* Payment data does not get to sit overseas.

**Insurance — IRDAI.** The IRDAI (Maintenance of Insurance Records) Regulations, 2015
require records of insurance policies and claims made in India to be held in data
centres located and maintained in India. The insurer's records stay on Indian soil.

**Securities — SEBI.** SEBI's 2023 Framework for Adoption of Cloud Services (circular
SEBI/HO/ITD/ITD_VAPT/P/CIR/2023/033) requires regulated entities' data to be stored
within India — and it is pointed about responsibility: adopting a cloud provider does
not move the accountability. The regulated entity remains answerable for the
confidentiality, integrity, and security of its data and logs, provider or no provider.

One thing these are *not*: they are not India's general privacy law. The DPDP Act, 2023
is comparatively permissive about moving personal data abroad. The hard localization
lives in these **sectoral** rules — which is why "our privacy team signed off on the
vendor" is not the same statement as "our sector regulator is satisfied."

## Why this lands on the test lab

None of those circulars mentions device testing. They govern the data. But a test
*produces* data — and for a regulated app, that data is exactly the kind the rules pin
to India.

Run a test and you generate an app binary, screen recordings, logs, and network
captures. On a payments app, a lending flow, or a policy-servicing screen, those
artifacts routinely contain account numbers, transaction details, and personal
information — the payment and policyholder data the rules localize. Send that test to a
third-party device cloud and every one of those artifacts lands on infrastructure you
don't own, frequently outside India, and [stays there for as long as the vendor's
retention window](/blog/what-actually-leaves-your-network).

SEBI's framing is the one to sit with: you can move the *workload* to someone else's
cloud, but the *accountability* stays with you. You can outsource the infrastructure;
you cannot outsource the responsibility. A regulator asking "where did this data go, and
who could read it" is not answered by "our testing vendor handles that."

## Why on-premise answers it cleanly

If the lab runs inside your own network — the devices, the builds, the captured video,
and the logs never leaving your walls, air-gapped if the review demands it — the
residency question answers itself. There is no foreign leg to document, no vendor
retention window to prove you closed, no data you have to certify you repatriated
within 24 hours. The test data was never anywhere it wasn't allowed to be. That is the
whole reason a bank or a broadcaster tends to want [the lab, and everything it captures,
inside its own perimeter](/enterprise) — and it is the practical difference between the
device clouds that [can run in your data centre and the ones that
can't](/blog/which-device-clouds-run-in-your-data-centre).

## The honest boundary

If you test strictly with synthetic data, the calculus softens — genuinely. But real
test environments touch real builds, real integrations, and real captures, and none of
these rules carries a "but it was only a test" exemption. The accountable-entity
standard does not pause for a QA run. For a regulated entity, "where does the test data
live?" is a question with a regulator attached — and on-premise is how you keep the
answer inside your own walls.
