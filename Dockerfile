# FROM golang:alpine AS build-container
# ADD . /work
# WORKDIR /work
# ENV GO111MODULE=on
# RUN apk --update add --no-cache git mercurial
# RUN go build

# FROM alpine:3.9
# COPY --from=build-container /work/RecommendSystem /usr/local/bin/RecommendSystem
# CMD ["usr/local/bin/RecommendSystem"]

FROM golang:1.13.0-alpine

WORKDIR /go/src

ENV GO111MODULE=on

RUN apk add --no-cache \
  alpine-sdk \
  git \
  && go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]
