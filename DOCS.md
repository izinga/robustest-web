# Documentation publishing (/docs)

The website serves product documentation at `https://robustest.com/docs`,
rendered from the private repo
[`izinga/robustest_documentation_md`](https://github.com/izinga/robustest_documentation_md).
The docs repo is the single source of truth — its content is **never
committed into robustest-web**. The server fetches it at runtime as a
GitHub tarball and renders the markdown (goldmark, docsify-style
`_sidebar.md` navigation, per-page table of contents).

Publishing is **manual by design** (the repo is private; there is no
background polling).

## One-time setup

1. Create a GitHub **fine-grained personal access token**:
   GitHub → Settings → Developer settings → Fine-grained tokens →
   Generate new token
   - Repository access: **only** `izinga/robustest_documentation_md`
   - Permissions: **Contents: Read** (nothing else)
2. Add it to the server's env file (`/home/omnarayan/.env` on the
   production instance):

   ```bash
   DOCS_GITHUB_TOKEN=github_pat_XXXX...
   ```

3. Restart the site binary. It performs one initial sync at startup.

Without the token, `/docs` shows a "docs are syncing" page and
`/docs/refresh` returns the GitHub API error.

## Publishing flow (every docs change)

1. **Write** — edit or add markdown in the docs repo. New pages must
   also be added to `_sidebar.md` to appear in the navigation. Merge to
   `main`. Nothing changes on the site yet.
2. **Publish** — trigger a sync, any of:
   - `make docs-refresh` (from this repo; hits production)
   - `curl https://robustest.com/docs/refresh`
   - open `https://robustest.com/docs/refresh` in a browser
3. **Verify** — the response is:

   ```json
   {"status":"ok","sha":"<commit>","synced_at":"..."}
   ```

   If `sha` matches the docs repo's latest commit on `main`, the site is
   serving exactly what was merged.

The sync downloads the tarball, extracts it next to the current copy
under `data/docs/`, and swaps atomically — readers never see a
half-updated state.

## Failure behavior

- If a refresh fails (expired token, GitHub unreachable), the endpoint
  returns the error and the site **keeps serving the last good copy**.
- A server restart re-syncs once at boot, so docs survive redeploys.
- `data/` is gitignored; synced content never enters this repo.

## Environment variables

| Variable | Default | Purpose |
|---|---|---|
| `DOCS_GITHUB_TOKEN` | *(required in prod)* | Fine-grained read-only PAT for the private docs repo |
| `DOCS_REPO` | `izinga/robustest_documentation_md` | Repo to sync |
| `DOCS_BRANCH` | `main` | Branch to sync |
| `DOCS_DIR` | `./data/docs` | Where synced trees are extracted |
| `DOCS_SYNC_INTERVAL` | *(unset — disabled)* | Set e.g. `10m` to re-enable automatic polling |

## Rendering conventions (for docs authors)

- `_sidebar.md` drives the left navigation: bold lines are section
  headings, list links are pages.
- Links between pages use repo-root-relative `.md` paths
  (`guides/foo.md`) — the site rewrites them to `/docs/...` URLs.
- Images live in `assets/images/` and are referenced relatively
  (`assets/images/x.png` or `../assets/images/x.png`).
- Files and folders starting with `_` (e.g. `_archive/`) are not served.
- The page title is the first `#` heading; `##`/`###` headings build the
  "On this page" table of contents.
