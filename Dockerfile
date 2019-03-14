#/cetea/Dockerfile

FROM "golang:alpine"

WORKDIR /go/src

COPY . /go/src
# TODO: point to github

RUN cd /go/src && go build -o cetea

EXPOSE 8080

ENTRYPOINT "./main"