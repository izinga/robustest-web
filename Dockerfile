# Build stage
FROM node:20-alpine AS assets
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci
COPY src/ ./src/
COPY tailwind.config.js* ./
RUN mkdir -p public/assets/css && \
    npx @tailwindcss/cli -i ./src/css/input.css -o ./public/assets/css/app.css --minify

FROM golang:1.23-alpine AS builder
RUN go install github.com/a-h/templ/cmd/templ@latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN templ generate
COPY --from=assets /app/public/assets/css/app.css ./public/assets/css/app.css
ARG VERSION=dev
ARG BUILD_TIME=unknown
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -s -w" \
    -o robustest-web ./cmd/server

# Runtime stage
FROM alpine:3.19
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /opt/robustest-web
COPY --from=builder /app/robustest-web .
COPY --from=builder /app/public ./public
EXPOSE 3000
ENTRYPOINT ["./robustest-web"]
