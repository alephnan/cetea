FROM "golang:alpine" AS build-env
# 1. Static
ENV GO111MODULE on
WORKDIR /go/src/app
# 2. Dependencies
RUN apk add git
# 3. Post dependency COPY
# TODO: decouple changes from source vs dependencies.
# cache invalidated when XOR changes, and
# downstream build always redownload dependencies.
COPY *.go go.mod /go/src/app/
# 4. Dynamic
RUN go build -o cetea .

# # Lightweight Linux container
FROM alpine
# 1. Static
WORKDIR /app
# Alpine does not include dictionary of words, needed to
# generating build name.
RUN mkdir -p /usr/share/dict/
# 1.5. *Mostly static*
COPY utility/build_name_words /usr/share/dict/words
COPY startup.sh config ./

#2. Dependencies
RUN apk --no-cache add ca-certificates && apk --no-cache add bash

# Frequently changes
COPY static template ./

# 3. Post dependency COPY
COPY --from=build-env /go/src/app/cetea /app

# 4. Always dynamic
ARG port=8080
ENV PORT=$port
EXPOSE $port

ENTRYPOINT ["sh", "-c", "./startup.sh"]