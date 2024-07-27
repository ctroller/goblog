FROM golang:1.22.5-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o blog cmd/server.go

# Build stage for TailwindCSS
FROM alpine:3.20.2 AS tailwind
WORKDIR /app
RUN apk add --no-cache curl
RUN curl --proto '=https' --tlsv1.2 -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.6/tailwindcss-linux-x64 && \
    chmod +x tailwindcss-linux-x64 && \
    mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

COPY ui /app/ui
COPY tailwind.config.js /app/tailwind.config.js
RUN tailwindcss -i /app/ui/src/css/main.css -o /app/ui/static/css/tailwind.css --minify

# combine
FROM alpine:3.20.2
WORKDIR /app
COPY --from=builder /app/blog /app/
COPY --from=tailwind /app/ui/static/css/tailwind.css /app/ui/static/css/tailwind.css

USER goblog

EXPOSE 8080

CMD ["./blog"]
