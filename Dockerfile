FROM golang:alpine AS build

RUN apk add --no-cache curl git alpine-sdk

WORKDIR $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator

COPY go.mod go.sum $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator/

RUN go mod tidy

COPY . .

RUN go build -o /terraform-plan-validator

FROM alpine:latest
ENV OPA_GCP_POLICY=opa-gcp-policy.rego
ENV OPA_AZURE_POLICY=opa-azure-policy.rego
ENV OPA_AWS_POLICY=opa-aws-policy.rego
ENV OPA_REGO_QUERY=data.terraform.analysis.authz

RUN apk add --no-cache curl git alpine-sdk

WORKDIR /terraform-plan-validator

COPY app.env ./app.env
COPY opa-azure-policy.rego ./opa-azure-policy.rego
COPY app.env /opt/atlassian/pipelines/agent/build

COPY --from=build /terraform-plan-validator terraform-plan-validator
COPY --from=build terraform-plan-validator /usr/bin/terraform-plan-validator

ENTRYPOINT [ "./terraform-plan-validator" ]
