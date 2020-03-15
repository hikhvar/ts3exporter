FROM golang:1.14 as build-env

WORKDIR /go/src/github.com/hikhvar/ts3exporter
ADD . /go/src/github.com/hikhvar/ts3exporter

RUN go get -d -v ./...

RUN go build -o /go/bin/ts3exporter

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/ts3exporter /
ENTRYPOINT ["/ts3exporter"]