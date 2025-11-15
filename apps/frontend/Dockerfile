# ---------- Runtime: Nginx ----------
FROM nginx:1.27-alpine

# Без лишнего шума в таймзоне/логах
ENV TZ=UTC

# Опционально: свой конфиг с кэшированием, gzip и SPA-фолбеком
# Если не нужен — закомментируй COPY, Nginx возьмёт дефолтный.
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# Копируем статические файлы сайта внутрь образа
# Помести свои файлы рядом с Dockerfile в папке ./web (index.html и т.д.)
COPY web/src/ /usr/share/nginx/html/

# На всякий случай: правильные права и владелец
RUN chmod -R 755 /usr/share/nginx/html && \
    chown -R nginx:nginx /usr/share/nginx/html

EXPOSE 80
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget -qO- http://localhost/ || exit 1

CMD ["nginx", "-g", "daemon off;"]
