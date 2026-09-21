FROM node:22-alpine AS frontend-builder

WORKDIR /web
COPY web/package*.json ./
RUN npm install
COPY web/ ./
RUN npm run build

FROM golang:1.24-alpine AS go-builder

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

COPY --from=go-builder /app/leadphone-validator /app/leadphone-validator
COPY --from=frontend-builder /web/dist /app/web/dist
COPY openapi.yaml /app/openapi.yaml
COPY documentation.yaml /app/documentation.yaml

USER appuser

EXPOSE 3007

ENV PORT=3007

ENTRYPOINT ["/app/leadphone-validator"]
