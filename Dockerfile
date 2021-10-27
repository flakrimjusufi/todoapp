# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go test ./...

RUN go build -o /main

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY templates /templates

COPY --from=build /main /main

EXPOSE 8088

ENTRYPOINT ["/main"]