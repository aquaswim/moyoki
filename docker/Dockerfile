FROM alpine

LABEL org.opencontainers.image.title="Moyoki" \
      org.opencontainers.image.description="Mock api"

RUN addgroup -S app && adduser -S app -G app

USER app

WORKDIR /app
ENV PANEL_LISTEN_ADDR=:3000 \
    MOCK_LISTEN_ADDR=:3001 \
    DB_DSN=/config/moyoki.db \
    TZ=Asia/Jakarta

EXPOSE 3000 3001

VOLUME ["/config"]

COPY docker/moyoki /usr/bin/moyoki

ENTRYPOINT ["moyoki"]