# Etapa 1: Build
FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache git curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Asegura que la variable DATABASE_URL esté definida para el `generate`
ENV DATABASE_URL="mysql://root:HkQwtPxmrifezyMHwxTXBRBwdcsDNPAA@nozomi.proxy.rlwy.net:26283/railway"

RUN go run github.com/steebchen/prisma-client-go generate && \
    go build -o main .

# Etapa 2: Imagen final
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .env
COPY --from=builder /app/prisma ./prisma

EXPOSE 8080

CMD ["./main"]
