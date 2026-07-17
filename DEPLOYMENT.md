# Deployment

robustest.com runs on a single GCE instance as a systemd service. Deploys
are one command from a developer machine with `gcloud` access.

## Where it runs

| | |
|---|---|
| GCP project | `devicelab-470109` |
| Instance | `robustest-landing-instance-20251223-062108` (e2-micro) |
| Zone | `us-central1-c` |
| Public IP | `35.225.176.209` (robustest.com A record points here directly — no proxy/CDN) |
| Install dir | `/home/omnarayan/site` (binary, `public/`, `.env`) |
| Service | `robustest-web` (systemd, runs as root, auto-restart, starts on boot) |
| Ports | Binary binds `:443` (TLS terminated in-process) and `:80` (301 → https) |
| TLS certs | `/etc/letsencrypt/live/robustest.com/` (paths set in `.env`) |
| Logs | journald (`make deploy-logs`), plus `contact_form.log` in the install dir |

The unit file lives in this repo (`robustest-web.service`) and is installed
with `make deploy-service`.

## Everyday commands

```bash
make deploy            # build → tarball → gcloud scp → extract → restart → health check
make deploy-status     # systemd state + https://robustest.com/health
make deploy-logs       # live journald tail
make deploy-rollback   # swap the previous binary back and restart (~5s)
make docs-refresh      # publish docs after the team merges markdown (see DOCS.md)
```

A deploy replaces the binary, `public/`, `docs-content/`, and `blog-content/`.
**The tarball deliberately excludes `.env`** — the server's
`/home/omnarayan/site/.env` is the authoritative config and survives every deploy.
Note: `tar` extraction overwrites and adds but does not remove, so a post *deleted*
from `blog-content/` locally must also be removed on the server (the same caveat as
docs). To change env (keys, docs
token), edit that file on the server and `sudo systemctl restart robustest-web`.

Verify a deploy landed: `curl https://robustest.com/health` returns the git
short SHA the binary was built from.

## Server environment (`/home/omnarayan/site/.env`)

| Variable | Purpose |
|---|---|
| `PORT=443` | HTTPS port (binary also opens :80 for redirects when TLS is on) |
| `GIN_MODE=release` | |
| `ASSETS_PATH=./public` | |
| `TLS_CERT` / `TLS_KEY` | Let's Encrypt fullchain/privkey paths |
| `SENDGRID_API_KEY` | Contact-form email |
| `CONTACT_FROM_EMAIL` / `CONTACT_TO_EMAIL` | Sender (SendGrid-verified) / recipient |
| `TURNSTILE_SITE_KEY` / `TURNSTILE_SECRET_KEY` | Anti-spam on the contact form |
| `DOCS_GITHUB_TOKEN` | Read-only PAT for the private docs repo — see [DOCS.md](DOCS.md) |
| `HTTP_REDIRECT_PORT` | Optional override for the :80 redirect listener |
| `DOCS_SYNC_INTERVAL` | Optional (e.g. `10m`) to re-enable automatic docs polling |

## History & legacy fallback

Until 2026-07-14 the site ran manually: a binary at
`/home/omnarayan/robustest-web` started with `sudo ./robustest-web` inside a
screen session (`screen -r server`). That setup is **preserved untouched as an
emergency fallback**:

```bash
# Emergency: fall back to the pre-migration setup
sudo systemctl stop robustest-web
screen -r server            # reattach the old session
cd /home/omnarayan && sudo ./robustest-web
```

Note the legacy binary predates the :80 redirect and the /docs, /partners,
/enterprise pages — it is a last resort, not a rollback (use
`make deploy-rollback` for that).

## Analytics (two systems, deliberately)

- **Plausible (cloud)** — script in both layouts; rich custom events with
  properties (demo/partner lead split). Caveat: silently drops traffic
  from networks it classifies as datacenter/VPN (including our own office
  egress), so treat its absolute counts as a floor.
- **GoatCounter (self-hosted, ground truth)** — `goatcounter` systemd
  service on the same instance, listening on `localhost:8081` only
  (SQLite at `/home/omnarayan/goatcounter/goatcounter.sqlite3`). The
  counting beacon is proxied first-party at `robustest.com/gc/count` and
  `count.js` is served from our own assets, so ad-blockers and
  third-party IP filters can't interfere. Events mirror Plausible's with
  flattened names (e.g. `contact-form-submitted-partner`).
- **Dashboard access** (not publicly exposed): add
  `127.0.0.1 stats.robustest.com` to your local `/etc/hosts`, then
  `gcloud compute ssh <instance> --zone us-central1-c -- -L 8081:localhost:8081`
  and open `http://stats.robustest.com:8081`. Login is
  `hello@robustest.com` (password held by Om, set at install).
- Comparing GoatCounter vs Plausible visitor counts over a few weeks
  measures Plausible's undercount empirically.

## Known gaps / future work

- **Certificate renewal**: check how the Let's Encrypt cert renews
  (`sudo certbot certificates` on the instance). The :80 listener now makes
  webroot HTTP-01 renewal possible if it isn't already automated.
- **Run as non-root**: the service runs as root today (to read LE keys and
  bind :443, matching the legacy setup). Hardening: dedicated user +
  `AmbientCapabilities=CAP_NET_BIND_SERVICE` + cert group readable keys.
- **Secrets in git**: `.env` (with a stale SendGrid key) is still tracked in
  this repo's history — rotate the key and `git rm --cached .env` when
  convenient.
- `contact_form.log` in the install dir has no rotation.
