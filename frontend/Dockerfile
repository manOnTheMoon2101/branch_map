# ── Build stage ───────────────────────────────────────────────────────────────
FROM node:22-alpine AS builder

WORKDIR /app

# Copy package manifest and install all deps with npm
# (avoids pnpm lockfile/workspace issues in a clean Docker context)
COPY package.json ./
RUN npm install

# Copy source and build
COPY . .
RUN npm run build

# ── Runtime stage ─────────────────────────────────────────────────────────────
FROM node:22-alpine

WORKDIR /app

# adapter-node bundles everything — no node_modules needed at runtime
COPY --from=builder /app/build ./build

ENV PORT=3000
ENV HOST=0.0.0.0

EXPOSE 3000

CMD ["node", "build"]
