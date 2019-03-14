#/cetea/Dockerfile

FROM "golang:alpine" AS build-env
WORKDIR /go/src
ADD . /go/src/app
RUN cd /go/src/app && go build -o cetea

# Lightweight Linux container
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
WORKDIR /app
COPY --from=build-env /go/src/app /app

EXPOSE 8080
ENTRYPOINT "./cetea"
