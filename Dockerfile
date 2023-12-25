FROM golang:1.20-alpine

WORKDIR /app

# This will download all certificates (ca-certificates) and builds it in a
# single file under /etc/ssl/certs/ca-certificates.crt (update-ca-certificates)
# I also add git so that we can download with `go mod download` and
# tzdata to configure timezone in final image
RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

COPY go.mod go.sum ./

RUN go mod download

ENV GO111MODULE=on

COPY ./app/port_domain_service .

RUN chmod +x /app/port_domain_service

CMD /app/port_domain_service