FROM golang:1.12-alpine as builder
RUN apk add -U --no-cache ca-certificates git

RUN mkdir -p /app
COPY . /app

RUN \
  cd /app && \
  CGO_ENABLED=0 go build -o /app/release/drone-amazon-secrets github.com/izikaj/drone-amazon-secrets/cmd/drone-amazon-secrets

FROM alpine:3.6
EXPOSE 3000

ENV GODEBUG netdns=go

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/release/drone-amazon-secrets /bin/

ENTRYPOINT ["/bin/drone-amazon-secrets"]
