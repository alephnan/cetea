#!/bin/bash

FROM "golang:alpine" AS build-env
WORKDIR /go/src
RUN apk add --no-cache git
RUN go get github.com/tjarratt/babble
# TODO: evaluate if we can just get google oauth2
RUN go get -v golang.org/x/oauth2/...
ADD . /go/src/app
RUN cd /go/src/app && go build -o cetea

# Lightweight Linux container
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/app /app
# Alpine does not include dictionary of words, needed to
# generating build name.
ADD utility/build_name_words /usr/share/dict/words

EXPOSE 8080
ENTRYPOINT "./cetea"
