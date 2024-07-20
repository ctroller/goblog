FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o blog cmd/blog/server.go

# Build stage for TailwindCSS
FROM alpine as tailwind
WORKDIR /app
RUN apk add --no-cache curl
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.6/tailwindcss-linux-x64 && \
    chmod +x tailwindcss-linux-x64 && \
    mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

COPY ui/src/css /app/ui/src/css
COPY tailwind.config.js /app/tailwind.config.js
RUN tailwindcss -i /app/ui/src/css/main.css -o /app/ui/static/css/tailwind.css --minify

# combine
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/blog /app/
COPY --from=tailwind /app/ui/static /app/ui/static

EXPOSE 8080

CMD ["./blog"]