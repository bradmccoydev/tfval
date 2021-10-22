FROM golang:alpine AS build

RUN apk add --no-cache curl git alpine-sdk

WORKDIR $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator

COPY go.mod go.sum $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator/

RUN go mod tidy

COPY . .

RUN go build -o /terraform-plan-validator

FROM alpine:latest

WORKDIR /terraform-plan-validator

COPY conf.toml ./conf.toml
COPY mock.json ./mock.json

COPY --from=build /terraform-plan-validator terraform-plan-validator
COPY --from=build terraform-plan-validator /usr/bin/terraform-plan-validator

ENTRYPOINT [ "./terraform-plan-validator" ]
