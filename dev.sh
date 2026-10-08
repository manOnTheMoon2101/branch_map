#!/usr/bin/env bash
# dev.sh — starts the Go backend (with live reload) and SvelteKit frontend.
# Usage:  bash dev.sh
#         ./dev.sh       (after: chmod +x dev.sh)

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Make sure Go-installed tools (air) are on PATH
export PATH="${PATH}:$(go env GOPATH)/bin"

BLUE=$'\033[0;36m'
GREEN=$'\033[0;32m'
BOLD=$'\033[1m'
DIM=$'\033[2m'
NC=$'\033[0m'

# Kill the whole process group (both servers + their children) on Ctrl+C or SIGTERM
trap 'printf "\n${BOLD}Stopping servers...${NC}\n"; kill 0' INT TERM

# Auto-install air if missing
if ! command -v air &>/dev/null; then
    printf "${BLUE}[backend]${NC}  air not found — installing...\n"
    go install github.com/air-verse/air@latest
fi

printf "\n${BOLD}Branch Map${NC} — dev mode\n"
printf "  ${BLUE}▸ backend${NC}   http://localhost:8080  (live reload via air)\n"
printf "  ${GREEN}▸ frontend${NC}  http://localhost:5173\n"
printf "  ${DIM}Ctrl+C to stop both${NC}\n\n"

# ── Backend (Go + air live reload) ───────────────────
(
    cd "${ROOT}/backend"
    air 2>&1 | while IFS= read -r line; do
        printf "${BLUE}[backend]${NC}  %s\n" "$line"
    done
) &

sleep 1   # give air a moment before starting the frontend

# ── Frontend (SvelteKit) ─────────────────────────────
(
    cd "${ROOT}/frontend"
    FORCE_COLOR=1 pnpm dev 2>&1 | while IFS= read -r line; do
        printf "${GREEN}[frontend]${NC} %s\n" "$line"
    done
) &

# Wait for both background jobs; exit when both stop
wait
