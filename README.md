# ReleaseGate

Pre-deployment readiness, SLA risk, and staged rollout gatekeeper built with Go standard-library serverless handlers and a zero-build Vanilla JavaScript frontend.

## Run locally

```bash
go run ./cmd/server
```

Then open http://localhost:8080. The local server serves `public/` and exposes `/api/releases` and `/api/evaluate`.

## Deploy to Vercel

Import this repository, keep the project root as the root directory, and deploy. Vercel serves the static app from `public/` and routes `/api/*` to the Go API entrypoint. No Node build step is required.

## API

- `GET /api/releases` — list seeded and in-memory release candidates.
- `POST /api/releases` — create or update a release candidate.
- `POST /api/evaluate` — run concurrent readiness checks and return a verdict, score, rollout plan, and action items.

The in-memory registry is process-local and resets when a serverless instance is recycled. Use a durable database for shared production state.
