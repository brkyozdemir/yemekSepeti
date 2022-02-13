# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /yemekSepeti

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o yemekSepeti

EXPOSE 9000
