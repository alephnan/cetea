FROM "golang:alpine" AS build-env
RUN apk add git
COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -d -v
RUN go build -o cetea

# Lightweight Linux container
FROM alpine
RUN apk update && apk add ca-certificates && apk add bash && rm -rf /var/cache/apk*
COPY --from=build-env /go/src/app /app
WORKDIR /app
# Alpine does not include dictionary of words, needed to
# generating build name.
RUN mkdir -p /usr/share/dict/ && mv utility/build_name_words /usr/share/dict/words

# Docker is so dumb. This ARG 'port' resolves to empty string when declared
# above...
ARG port=8080
ENV PORT=$port
EXPOSE $port
ENTRYPOINT ["sh", "-c", "./startup.sh"]