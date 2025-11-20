############################################################
# Билдим и запускаем Next.js приложение (apps/casino/web/src)
############################################################

# ---------- Builder ----------
FROM node:20-alpine AS builder
WORKDIR /app

# Без телеметрии и dev-зависимостей в финальном образе
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1

# Депсы отдельно для лучшего кеша
COPY src/package*.json ./
RUN npm ci

# Исходники + сборка
COPY src ./
RUN npm run build && npm prune --omit=dev

# ---------- Runtime ----------
FROM node:20-alpine AS runner
WORKDIR /app

ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1 \
    HOST=0.0.0.0 \
    PORT=3000

# Берём только то, что нужно для запуска
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/next.config.mjs ./next.config.mjs

EXPOSE 3000
HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD wget -qO- http://localhost:${PORT}/ || exit 1

CMD ["npm", "run", "start"]
