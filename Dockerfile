FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/leadphone-validator .

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/leadphone-validator /app/leadphone-validator

USER appuser

EXPOSE 3007

ENV PORT=3007

ENTRYPOINT ["/app/leadphone-validator"]
