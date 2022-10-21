FROM golang:1.19-alpine3.16 AS builder
WORKDIR /go/src/app

COPY . .

RUN go build -o reckoner-go

FROM alpine:3.14
COPY --chown=65534:65534 --from=builder /go/src/app/reckoner-go ./reckoner
USER 65534

ENTRYPOINT [ "./reckoner" ]
