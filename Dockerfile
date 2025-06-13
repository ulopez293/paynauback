# Etapa 2: Imagen final
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
# Eliminamos esta línea:
# COPY --from=builder /app/.env .env
COPY --from=builder /app/prisma ./prisma

EXPOSE 8080

CMD ["./main"]
