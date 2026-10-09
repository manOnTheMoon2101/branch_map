# Capitec Branch Map

An interactive full-screen map that displays Capitec Bank branch locations across South Africa. Click any marker to view branch details and get directions.

---

## Stack

### Frontend

| Layer | Choice |
|---|---|
| Framework | [SvelteKit 3](https://svelte.dev/docs/kit) + [Svelte 5](https://svelte.dev) (runes mode) |
| Language | TypeScript |
| Styling | [Tailwind CSS v4](https://tailwindcss.com) |
| Components | [shadcn-svelte](https://www.shadcn-svelte.com) (vega style, neutral palette) |
| Map | [MapLibre GL v6](https://maplibre.org) |
| Map tiles | [Carto Voyager](https://carto.com/basemaps) (light) / Carto Dark Matter (dark) |
| Icons | [Lucide Svelte](https://lucide.dev) |
| Font | [Roboto](https://fonts.google.com/specimen/Roboto) 300 · 400 · 500 · 700 |
| Package manager | [pnpm](https://pnpm.io) |
| Adapter | [@sveltejs/adapter-node](https://svelte.dev/docs/kit/adapter-node) |

### Backend

| Layer | Choice |
|---|---|
| Language | Go 1.21 |
| Server | `net/http` (standard library only) |
| Data | Dummy JSON (in-memory) |
| Port | `8080` |

---

## Design Tokens

### Colors

| Role | Tailwind | Hex |
|---|---|---|
| Branch marker / header | `sky-600` | `#0284c7` |
| Head Office marker / header | `red-500` | `#ef4444` |

### Font

**Roboto** — loaded via [`@fontsource/roboto`](https://fontsource.org/fonts/roboto) (self-hosted, no CDN).

```css
@import "@fontsource/roboto/300.css";  /* Light    */
@import "@fontsource/roboto/400.css";  /* Regular  */
@import "@fontsource/roboto/500.css";  /* Medium   */
@import "@fontsource/roboto/700.css";  /* Bold     */
```

---

## Project Structure

```
branch_map/
├── backend/
│   ├── main.go            # HTTP API server
│   ├── go.mod
│   ├── Dockerfile
│   └── .air.toml          # Live-reload config (dev only)
├── frontend/
│   ├── src/
│   │   ├── assets/
│   │   │   └── capitec.png
│   │   ├── lib/
│   │   │   └── components/ui/map/   # MapLibre GL component library
│   │   └── routes/
│   │       ├── +layout.svelte
│   │       ├── +page.svelte         # Full-screen map page
│   │       └── layout.css           # Tailwind + shadcn theme
│   ├── Dockerfile
│   └── vite.config.ts
├── docker-compose.yml
├── dev.sh                 # Starts both servers locally
└── README.md
```

---

## Running Locally

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [Node.js 22+](https://nodejs.org) + [pnpm](https://pnpm.io/installation)
- [air](https://github.com/air-verse/air) — installed automatically by `dev.sh` if missing

### One command

```bash
bash dev.sh
```

This starts both servers with live reload and colour-coded output:

| Service | URL |
|---|---|
| Frontend (SvelteKit) | http://localhost:5173 |
| Backend (Go API) | http://localhost:8080 |

Press **Ctrl+C** to stop both.

### Manual start

```bash
# Terminal 1 — backend (with live reload)
cd backend && air

# Terminal 2 — frontend
cd frontend && pnpm dev
```

---

## Running with Docker

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) with the Compose plugin

### Start

```bash
docker compose up --build
```

| Service | URL |
|---|---|
| Frontend | http://localhost:3000 |
| Backend API | http://localhost:8080 |

### Stop

```bash
docker compose down
```

### Rebuild after code changes

```bash
docker compose up --build --force-recreate
```

---

## API

### `GET /api/branches`

Returns all branch locations.

**Response**

```json
[
  {
    "id": 1,
    "name": "Capitec Bank Wellington",
    "type": "branch",
    "address": "28 Church St",
    "city": "Wellington",
    "province": "Western Cape",
    "phone": "+27 860 102 043",
    "hours": "Mon–Fri 08:00–16:30 | Sat 08:00–13:00",
    "latitude": -33.639735,
    "longitude": 19.008627
  }
]
```

**Branch types**

| `type` | Description |
|---|---|
| `branch` | Full-service branch |
| `head` | Head Office |
