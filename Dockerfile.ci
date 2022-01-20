FROM golang:alpine AS build

RUN apk add --no-cache curl git alpine-sdk

WORKDIR $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator

COPY go.mod go.sum $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator/

RUN go mod tidy

COPY . .

RUN go build -o /terraform-plan-validator

FROM node:17-alpine
ENV OPA_GCP_POLICY=opa-gcp-policy.rego
ENV OPA_AZURE_POLICY=opa-azure-policy.rego
ENV OPA_AWS_POLICY=opa-aws-policy.rego
ENV OPA_REGO_QUERY=data.terraform.analysis.authz

RUN apk --no-cache add curl git alpine-sdk openssh-client

RUN curl -SL "https://github.com/aquasecurity/tfsec/releases/download/v0.63.1/tfsec-linux-amd64" --output tfsec && \
  chmod +x tfsec && \
  mv tfsec /usr/local/bin

# amit needs: 0.13.4
# RUN curl -SL "https://releases.hashicorp.com/terraform/0.13.4/terraform_0.13.4_linux_amd64.zip" --output terraform.zip && \
RUN curl -SL "https://releases.hashicorp.com/terraform/1.0.11/terraform_1.0.11_linux_amd64.zip" --output terraform.zip && \
  unzip "terraform.zip" && \
  mv terraform /usr/local/bin && \
  rm terraform.zip

WORKDIR /terraform-plan-validator

COPY app.env ./app.env
COPY policies/opa-azure-policy.rego ./opa-azure-policy.rego
COPY policies/opa-gcp-policy.rego ./opa-gcp-policy.rego
COPY app.env /opt/atlassian/pipelines/agent/build

COPY --from=build /terraform-plan-validator terraform-plan-validator
COPY --from=build terraform-plan-validator /usr/bin/terraform-plan-validator

LABEL "com.azure.dev.pipelines.agent.handler.node.path"="/usr/local/bin/node"

RUN apk add --no-cache --virtual .pipeline-deps readline linux-pam \
  && apk add bash sudo shadow \
  && apk del .pipeline-deps

CMD [ "node" ]
