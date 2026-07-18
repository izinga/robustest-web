---
title: "Which device clouds can actually run in your data centre?"
description: "If your security review said no to the public cloud, most device clouds can't run inside your walls at all — by their own docs. The mid-2026 survey."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_image: "/assets/images/authors/author-om.png"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-06-10
sources:
  - title: "BrowserStack — Custom Device Lab (Enterprise; hosted in BrowserStack data centres)"
    url: "https://www.browserstack.com/docs/custom-device-lab/overview"
  - title: "Sauce Labs — Private Devices (a dedicated pool on the Real Device Cloud)"
    url: "https://saucelabs.com/products/private-devices-real-device-cloud"
  - title: "AWS Device Farm — private devices (deployed in an Amazon data centre, us-west-2)"
    url: "https://docs.aws.amazon.com/devicefarm/latest/developerguide/working-with-private-devices.html"
  - title: "Kobiton — pricing (on-prem / dedicated devices under Enterprise)"
    url: "https://kobiton.com/pricing/"
  - title: "pCloudy — Lab in a Box (on-premise appliance)"
    url: "https://www.pcloudy.com/lab-in-a-box-on-premise-testing/"
---

Some teams can [send an unreleased build to a third party's cloud](/blog/what-actually-leaves-your-network), and most do. Some
teams cannot — a bank whose regulator localises payment data, a broadcaster whose
unreleased app is the product, a health app under a data-protection regime. For those
teams the useful question isn't "which device cloud is best," it's the much narrower
"which one can run **inside our own network** at all."

That field is far smaller than the field of vendors. Here is where the major device
clouds stand, as of mid-2026, according to each vendor's own documentation. We've
linked every source; check them, because these pages change.

## What "on-prem" turns out to mean

The word does a lot of quiet work, so it helps to separate three different things
vendors call by similar names:

- **Public cloud** — shared devices in the vendor's data centre.
- **Private / dedicated cloud** — devices reserved to you, still in the *vendor's*
  data centre.
- **On-premises** — the devices, and ideally the control plane and storage, inside
  *your* network.

Most "private" offerings are the middle one. The third is the one a security review
actually asks for, and it is rare.

## The survey

| Vendor | Runs inside your network? | What they offer |
|---|---|---|
| BrowserStack | No | "Custom Device Lab" (Enterprise) — dedicated devices, hosted in BrowserStack's own data centres. Its page advertised "on-premises" wording in Nov 2025 that is no longer present. |
| Sauce Labs | No | "Private Devices" — a dedicated pool "only accessible to members of your organization," on the Sauce-hosted Real Device Cloud. |
| AWS Device Farm | No | Private devices are "deployed on your behalf in an Amazon data center… exclusive to your AWS account," available in us-west-2 only. |
| TestMu AI (ex-LambdaTest) | No | No customer-premises deployment documented. |
| Firebase Test Lab | No | A CI test-execution service; no private or on-prem option. |
| Kobiton | Contact sales | On-prem / dedicated devices appear only under the Enterprise tier ("Let's Talk"); self-serve tiers are public cloud. |
| pCloudy | Yes | "Lab in a Box" — the real-device cloud running inside your data centre, air-gapped setups included (vendor-stated). |
| HeadSpin | Claims to | Markets "the industry's only enterprise-grade on-premise setup." |

Read the pattern rather than any single row: **five of the eight cannot run inside
your network at all.** For most of the market, "private" means a reserved shelf in
someone else's building — which is a fine thing, and not the thing a data-residency
requirement is asking for.

## Two honest notes

**HeadSpin's superlative is contestable.** They market "the industry's only
enterprise-grade on-premise setup." By this same survey, pCloudy offers an on-premise
appliance, Kobiton offers on-prem under Enterprise, and [RobusTest runs the full stack on your premises](/enterprise) — so "only" is a strong word. We'd quote the claim and let you weigh
it, rather than characterise it.

**"On-prem" and "air-gapped" are not the same.** On-prem means the gear is in your
building. [Air-gapped](/blog/air-gapped-is-not-just-on-prem) means it also has no path to the internet — which breaks anything
that needs to phone home (a licence check, an email relay, a push service). If your
requirement is the stricter one, ask each vendor specifically, because a page that says
"on-premise" rarely says "and it runs disconnected."

## Where we sit

We'll be direct about our own position, since this is our blog: RobusTest runs the
**entire** lab — control plane, devices, and storage — inside the customer's network,
air-gap included. That's [the bet the rest of the site is about](/blog/everyone-went-to-the-cloud), and the honest reason
this survey exists is that we kept meeting teams who assumed it wasn't an option,
because the loudest vendors in the category can't offer it.

If you're weighing this, the useful exercise is not a feature grid. It's one question
put to each vendor in writing: *does the device, the control plane, and the test data
stay inside our network — and does it work with no path out?* The answers narrow the
field fast.
