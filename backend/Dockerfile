FROM "golang:alpine" AS build-env
# 1. Static
ENV GO111MODULE on
WORKDIR /go/src/app
# 2. OS Dependencies should not change much between runs
RUN apk add git
# 3. Network fetch for go deps is expensive, cache as much as possible.
COPY go.mod /go/src/app/
RUN go mod download
# 4. Not too bad to run every time, once deps downloaded.
COPY *.go /go/src/app/
RUN go build -o cetea .

# # Lightweight Linux container
FROM alpine
# 1. Static
WORKDIR /app
# Alpine does not include dictionary of words, needed to
# generating build name.
RUN mkdir -p /usr/share/dict/ ./config ./static ./template
# 1.5. *Mostly static*
COPY utility/build_name_words /usr/share/dict/words
COPY startup.sh ./
COPY config config

#2. Dependencies
RUN apk --no-cache add ca-certificates && apk --no-cache add bash

# 3. Post dependency COPY
COPY --from=build-env /go/src/app/cetea /app

# 4. Always dynamic
ARG port=8080
ENV PORT=$port
EXPOSE $port

ENTRYPOINT ["sh", "-c", "./startup.sh"]