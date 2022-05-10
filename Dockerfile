FROM golang:alpine AS build

RUN apk add --no-cache curl git alpine-sdk

WORKDIR $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator

COPY go.mod go.sum $GOPATH/src/github.com/bradmccoydev/terraform-plan-validator/

RUN go mod tidy

COPY . .

RUN go build -o /terraform-plan-validator

FROM alpine:latest
ENV OPA_REGO_QUERY=data.terraform.analysis.authz

RUN apk add --no-cache curl git alpine-sdk

RUN curl -SL "https://github.com/aquasecurity/tfsec/releases/download/v1.20.0/tfsec-linux-amd64" --output tfsec && \    
    chmod +x tfsec && \
    mv tfsec /usr/local/bin

RUN curl -SL "https://releases.hashicorp.com/terraform/1.1.9/terraform_1.1.9_linux_amd64.zip" --output terraform.zip && \
    unzip "terraform.zip" && \
    mv terraform /usr/local/bin && \
    rm terraform.zip

RUN curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | sh

WORKDIR /terraform-plan-validator

COPY app.env ./app.env
COPY app.env /opt/atlassian/pipelines/agent/build

COPY --from=build /terraform-plan-validator terraform-plan-validator
COPY --from=build terraform-plan-validator /usr/bin/terraform-plan-validator

ENTRYPOINT [ "./terraform-plan-validator" ]
