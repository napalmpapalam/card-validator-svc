FROM golang:1.23-alpine as buildbase

WORKDIR /go/src/github.com/napalmpapalam/card-validator-svc
COPY vendor .
COPY config.yaml .
COPY . .

ENV GO111MODULE="on"
ENV CGO_ENABLED=0
ENV GOOS="linux"

RUN go build -o /usr/local/bin/card-validator-svc github.com/napalmpapalam/card-validator-svc

###

FROM alpine:3.9 as calpine

RUN apk add --no-cache ca-certificates

FROM calpine

COPY --from=buildbase /usr/local/bin/card-validator-svc /usr/local/bin/card-validator-svc
COPY --from=buildbase /go/src/github.com/napalmpapalam/card-validator-svc/config.yaml /etc/card-validator-svc/config.yaml

EXPOSE 8000
ENV KV_VIPER_FILE="/etc/card-validator-svc/config.yaml"

ENTRYPOINT ["card-validator-svc"]
