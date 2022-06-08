![passing](https://github.com/bradmccoydev/tfval/actions/workflows/ci.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/bradmccoydev/tfval)](https://goreportcard.com/report/github.com/bradmccoydev/tfval) ![GitHub](https://img.shields.io/github/license/bradmccoydev/tfval) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/bradmccoydev/tfval)

# TFVAL
This tool validates Terraform Plans it has been developed in golang as a wrapper around TFSEC and OPA to provide guardrails when deploying in CI/CD pipelines. You can find the latest release at the release page

### Command Description

| Command | Parameters |
| --- | --- |
| check | Check if the plan passes OPA and TFSEC Policy |
| checkopa | Check if the plan passes OPA Policy |
| opascore | Gets the OPA score report |
| tfsec | Outputs TfSec vulnerability report |
| sendreport | Sends Terraform validation Report to slack |
| costs | Matches Infracost and Budget |

### Commands Parameters

| Command | Parameters |
| --- | --- |
| tfsec | --tfsecReport "delete-rg-test.json" --tfsecMaxSeverity "CRITICAL" |
| check | --planFileName "policies/delete-rg-test.json" --tfsecReportLocation "pkg/tfsec/mock.json" --tfsecMaxSeverity "CRITICAL" --opaConfig "[{\"location\":\"policies/opa-azure-policy.rego\",\"query\":\"data.terraform.analysis.deny[x]\"}]" |
| checkopa | --planFileName "policies/delete-rg-test.json" --opaConfig "[{\"location\":\"policies/opa-azure-policy.rego\",\"query\":\"data.terraform.analysis.authz\"}]" |
| opascore | --planFileName "delete-rg-test.json" --policyLocation "opa-aws-policy.rego" |
| sendreport | --fileName "delete-rg-test.json" --slackWebhook "*" --prNumber "1" --repoFullUrl "x" --tfsecMaxSeverity "MEDIUM"  |
| cost | --fileName "policies/delete-rg-test.json" --opaConfig "[{\"location\":\"policies/opa-azure-policy.rego\",\"query\":\"data.terraform.analysis.authz\"}]"  |

 - /usr/bin/tfsec-analysis-terraform tfsec "$BITBUCKET_PR_ID" "$BITBUCKET_GIT_HTTP_ORIGIN" "tfsec-report.json" "$SLACK_WEBHOOK"
 
### Docker
```bash
docker pull bradmccoydev/tfval:latest
docker run -p 80:80 bradmccoydev/tfval:latest check --planFileName "delete-rg-test.json" --policyLocation "opa-aws-policy.rego" --tfsecMaxSeverity "CRITICAL" --opaRegoQuery "data.terraform.analysis.authz"
```

### Maintainers:
* Brad McCoy ([@bradmccoydev](https://github.com/bradmccoydev)), Basiq
* Ben Poh ([@benhpoh](https://github.com/benhpoh)), Moula

## Thanks to all the contributors ❤️
<a href = "https://github.com/bradmccoydev/tfval/graphs/contributors">
  <img src = "https://contrib.rocks/image?repo=bradmccoydev/tfval"/>
</a>

### License

Terraform Plan Validator is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/bradmccoydev/tfval/blob/main/LICENSE)

opa eval --fail-defined --format raw --input policies/delete-rg-test.json --data policies/opa-azure-policy.rego 'data.terraform.analysis.authz'

opa eval --fail-defined --format raw --input policies/delete-rg-test.json --data policies/tags-policy.rego 'data.terraform.common.deny[x]'
