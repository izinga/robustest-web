---
title: "Your unreleased app sits on someone else's cloud for weeks"
description: "When you test on a device cloud, your app binary, your screen recordings, and your logs are stored on the vendor's infrastructure — for 30 to 400 days, by their own documentation. Here's exactly how long, per vendor."
category: LAB
author: Om Narayan
author_role: "Co-founder"
author_url: "https://www.linkedin.com/company/robustest/"
author_company: "RobusTest"
author_company_url: "https://www.linkedin.com/company/robustest/"
date: 2026-06-24
sources:
  - title: "AWS Device Farm — data protection & retention"
    url: "https://docs.aws.amazon.com/devicefarm/latest/developerguide/data-protection.html"
    note: "Uploaded applications 30 days; logs 400 days; video & other artifacts 400 days; then deleted — archiving is the customer's responsibility."
  - title: "Sauce Labs — viewing test results (asset retention)"
    url: "https://docs.saucelabs.com/test-results/viewing-test-results/"
    note: "Videos, screenshots, and logs retained for 30 days."
  - title: "Firebase Test Lab — analyzing results"
    url: "https://firebase.google.com/docs/test-lab/android/analyzing-results"
    note: "Detailed results (logs, screenshots, videos) retained 90 days in a Cloud Storage bucket."
  - title: "BrowserStack — App Automate (artifact retention)"
    url: "https://www.browserstack.com/docs/app-automate"
    note: "Video 30 days; screenshots and logs 60 days."
---

To test your app on a public device cloud, you upload it. That sentence is the whole
security review, and most teams skate past it. The app you uploaded — often an
unreleased build — does not vanish when the test ends. It sits on the vendor's
infrastructure, along with the video of it running and the logs it produced, for a
documented period. Here's how long, according to each vendor's own docs.

## The retention periods, by vendor

| Vendor | Uploaded app binary | Video / screenshots / logs |
|---|---|---|
| AWS Device Farm | 30 days | logs & video: **400 days** |
| BrowserStack (App Automate) | (per plan) | video 30 days; screenshots & logs 60 days |
| Sauce Labs | (with assets) | videos, screenshots, logs: 30 days |
| Firebase Test Lab | (in results) | detailed results: 90 days |

None of these is a scandal. They're all documented, all defensible, and mostly there so
you can actually go back and look at a failed run. The point isn't that any vendor is
hiding something — it's the arithmetic most buyers never do: **your unreleased binary
and a screen recording of it running live on a third party's servers for somewhere
between a month and well over a year.**

AWS Device Farm is the most explicit, and worth reading closely because it's the most
generous with time: uploaded **applications** are retained 30 days, but **logs and
video recordings are retained 400 days** — and after that, the data is deleted, with
AWS stating plainly that archiving anything you want to keep is *your* responsibility.
So the recording of your pre-release app runs for over a year by default, and the
cleanup is on you to think about.

## Why this is the question a security review actually asks

For most teams none of this matters. For some it's the whole decision:

- A **fintech** app under data-localisation rules can't have its build and [network traffic](/platform/network-capture) sitting in a vendor's region for 400 days.
- A **broadcaster** testing an unreleased streaming app — the thing competitors would
  most like to see early — has to account for a screen recording of it living on
  someone else's cloud for weeks.
- Any team whose **security questionnaire** has a "where does our data reside and for
  how long" line now has to fill it in with a number they don't control.

The honest framing is not "these vendors are careless." It's that **using a public
device cloud means your app and its artifacts leave your network and persist on
someone else's, for a period they set** — and that's a fact to weigh, not a footnote to
skip.

## The alternative that removes the question

There's a category of team for whom the cleanest answer is to not have the data leave at
all. That's [the bet RobusTest is built on](/blog/everyone-went-to-the-cloud): the devices, the control plane, and the
storage all run inside your own network, so the app binary, the video, and the logs
never sit on anyone else's infrastructure — there is no retention period to reason
about, because nothing was uploaded anywhere.

If your security review has never asked how long your unreleased build lives on a
vendor's cloud, it's worth asking once. The numbers above are the ones you'll get back,
and for a surprising number of teams, the right number is zero.
