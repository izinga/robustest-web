# Documentation publishing (/docs)

The website serves product documentation at `https://robustest.com/docs`,
rendered from the private repo
[`izinga/robustest_documentation_md`](https://github.com/izinga/robustest_documentation_md).
The docs repo is the single source of truth — its content is **never
committed into robustest-web** (`docs-content/` is gitignored).

Docs are **fetched locally at build time and shipped with the deploy** —
the production server needs no GitHub credentials and never calls GitHub.

## How it works

- `make docs-fetch` pulls the docs repo tarball (using your local `gh`
  login) into `./docs-content/`.
- `make release-linux` (and therefore `make deploy`) runs `docs-fetch`
  automatically and bundles `docs-content/` into the release tarball.
- The server detects the bundled directory and serves it: markdown is
  rendered with goldmark, `_sidebar.md` drives the navigation, `##`/`###`
  headings build the per-page table of contents.

## Publishing flow (every docs change)

1. **Write** — edit or add markdown in the docs repo. New pages must also
   be added to `_sidebar.md` to appear in the navigation. Merge to `main`.
2. **Publish** — from this repo, one command:

   ```bash
   make docs-publish     # fetch latest docs, ship ONLY docs, refresh — no binary deploy
   ```

   (A full `make deploy` also carries the latest docs, since the release
   bundles them.)
3. **Verify** — the command ends by calling `/docs/refresh`, which
   returns `{"status":"ok","sha":"<content fingerprint>",...}` after
   re-scanning the shipped content. Spot-check a changed page on
   `https://robustest.com/docs`.

Requirements for whoever publishes: `gh` CLI authenticated with access to
the docs repo, and `gcloud` access to the instance — the same access
needed to deploy the site.

## Failure behavior

- The docs swap on the server is atomic (`docs-content.new` → rename), so
  readers never see a half-updated tree; the previous copy is kept as
  `docs-content.old`.
- If `/docs` shows "Docs are syncing", the bundled directory is missing —
  run `make docs-publish` (or a full `make deploy`).

## Optional: server-side GitHub sync

The server can instead fetch from GitHub itself (the pre-bundling mode):
set `DOCS_GITHUB_TOKEN` (fine-grained PAT, Contents: Read, docs repo only)
in the server env, and optionally `DOCS_SYNC_INTERVAL=10m` for polling.
Bundled content takes precedence when `docs-content/` exists.

| Variable | Default | Purpose |
|---|---|---|
| `DOCS_LOCAL_DIR` | `./docs-content` | Bundled docs directory (preferred mode) |
| `DOCS_GITHUB_TOKEN` | *(unset)* | Enables server-side GitHub fetch mode |
| `DOCS_REPO` | `izinga/robustest_documentation_md` | Repo (both modes) |
| `DOCS_BRANCH` | `main` | Branch (both modes) |
| `DOCS_DIR` | `./data/docs` | Extraction dir for GitHub mode |
| `DOCS_SYNC_INTERVAL` | *(unset — disabled)* | Polling interval for GitHub mode |

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
