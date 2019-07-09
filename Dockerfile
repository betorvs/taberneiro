FROM golang:1.12.6-alpine3.10 AS golang

RUN apk add --no-cache git
RUN go get github.com/golang/dep && go install github.com/golang/dep/cmd/dep
RUN mkdir -p /builds/go/src/github.com/betorvs/taberneiro/
ENV GOPATH /builds/go
COPY . /builds/go/src/github.com/betorvs/taberneiro/
ENV CGO_ENABLED 0
RUN cd /builds/go/src/github.com/betorvs/taberneiro/ && dep ensure -v && go build

FROM alpine:3.10
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
COPY --from=golang /builds/go/src/github.com/betorvs/taberneiro/taberneiro /
RUN update-ca-certificates

EXPOSE 9090
RUN chmod +x /taberneiro
CMD ["/taberneiro"]