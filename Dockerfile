# For development ====================
FROM golang:1.13.0-alpine as develop
WORKDIR /go/src

ENV GO111MODULE=on
RUN apk --update add --no-cache git \
  alpine-sdk \
  && go get github.com/pilu/fresh

EXPOSE 8080
CMD ["fresh"]
# ====================================

# For production =====================
FROM golang:1.13.0-alpine as build
ADD . /go/src
WORKDIR /go/src

ENV GO111MODULE=on
RUN apk --update add --no-cache git \
  mercurial
RUN go build -o RecommendSystem main.go

FROM alpine:3.9 as release
WORKDIR /apps
COPY --from=build /go/src/.env /apps/
COPY --from=build /go/src/RecommendSystem /usr/local/bin/RecommendSystem
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/RecommendSystem"]
# ====================================
