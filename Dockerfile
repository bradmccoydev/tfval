FROM golang:alpine AS build

RUN apk add --no-cache curl git alpine-sdk

WORKDIR $GOPATH/src/github.com/bradmccoydev/tfval

COPY go.mod go.sum $GOPATH/src/github.com/bradmccoydev/tfval/

RUN go mod tidy

COPY . .

RUN go build -o /tfval

FROM alpine:latest

RUN apk add --no-cache curl git alpine-sdk

RUN curl -SL "https://github.com/aquasecurity/tfsec/releases/download/v1.23.3/tfsec-linux-amd64" --output tfsec && \    
    chmod +x tfsec && \
    mv tfsec /usr/local/bin

RUN curl -SL "https://releases.hashicorp.com/terraform/1.2.2/terraform_1.2.2_linux_amd64.zip" --output terraform.zip && \
    unzip "terraform.zip" && \
    mv terraform /usr/local/bin && \
    rm terraform.zip

RUN curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | sh

WORKDIR /tfval

COPY --from=build /tfval tfval
COPY --from=build tfval /usr/bin/tfval

ENTRYPOINT [ "./tfval" ]
