FROM golang:1.9.1 AS build-env

# TODO:
#   * Set ENV (enviornment variable) for Firebase config


ADD . /go/src/github.com/ecclesia-dev/campaign-service/
RUN cd /go/src/github.com/ecclesia-dev/campaign-service/app && go get && go build -o campaign-service

WORKDIR /go/src/github.com/ecclesia-dev/campaign-service/app

ENTRYPOINT ./account-service | tee ../logs/campaign-service.log