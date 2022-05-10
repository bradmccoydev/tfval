#!/bin/bash

function terragruntInit {
  # Gather the output of `terragrunt init`.
  echo "init: info: initializing Terragrunt configuration in ${tfWorkingDir}"
  initOutput=$(${tfBinary} init -input=false ${*} 2>&1)
  initExitCode=${?}

  # Exit code of 0 indicates success. Print the output and exit.
  if [ ${initExitCode} -eq 0 ]; then
    echo "init: info: successfully initialized Terragrunt configuration in ${tfWorkingDir}"
    echo "${initOutput}"
    echo
    exit ${initExitCode}
  fi

  # Exit code of !0 indicates failure.
  echo "init: error: failed to initialize Terragrunt configuration in ${tfWorkingDir}"
  echo "${initOutput}"
  echo

  # Comment on the pull request if necessary.
  if [ "$GITHUB_EVENT_NAME" == "pull_request" ] && [ "${tfComment}" == "1" ]; then
    initCommentWrapper="#### \`${tfBinary} init\` Failed

\`\`\`
${initOutput}
\`\`\`

*Workflow: \`${GITHUB_WORKFLOW}\`, Action: \`${GITHUB_ACTION}\`, Working Directory: \`${tfWorkingDir}\`, Workspace: \`${tfWorkspace}\`*"

    initCommentWrapper=$(stripColors "${initCommentWrapper}")
    echo "init: info: creating JSON"
    initPayload=$(echo "${initCommentWrapper}" | jq -R --slurp '{body: .}')
    initCommentsURL=$(cat ${GITHUB_EVENT_PATH} | jq -r .pull_request.comments_url)
    echo "init: info: commenting on the pull request"
    echo "${initPayload}" | curl -s -S -H "Authorization: token ${GITHUB_TOKEN}" --header "Content-Type: application/json" --data @- "${initCommentsURL}" > /dev/null
  fi

  exit ${initExitCode}
}
