#!/bin/bash

FROM "golang:alpine" AS build-env
WORKDIR /go/src
RUN apk add --no-cache git
RUN go get -v google.golang.org/api/cloudresourcemanager/v1
RUN go get -v github.com/tjarratt/babble
RUN go get -v github.com/alephnan/google-auth-id-token-verifier
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

# Docker is so dumb. This ARG 'port' resolves to empty string when declared
# above...
ARG port=8080
ENV PORT=$port
EXPOSE $port
# https://github.com/moby/moby/issues/13877
# TODO: Problem ENTRYPOINT /CMD pattern is NGINX does not receive kill signal
# and container left dangling.
ENTRYPOINT ["sh", "-c"]
CMD ["./cetea --port $PORT"]
