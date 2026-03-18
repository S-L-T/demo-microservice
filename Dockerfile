FROM golang:1.16.3 AS build-env

ADD . /dockerdev
WORKDIR /dockerdev
RUN GOOS=linux go build -gcflags="all=-N -l" -o /main

FROM debian:buster
EXPOSE 8080 8081 40000
WORKDIR /
COPY --from=build-env /main /
ENTRYPOINT ./main