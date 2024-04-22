# Docs

## CDN Icons
Use this URL for development
https://githubraw.com/ecaminero/ai-codereview/main/docs/images/AI-bot-codereview.png

Use this URL in production
https://cdn.githubraw.com/ecaminero/ai-codereview/main/docs/images/AI-bot-codereview.png


## Testing Actions
### Executions
```bash
# Basic use
act 

# Use on OSX
act pull_request --container-architecture linux/amd64

# List all the actions in your YAML file
act -l 

# Get Act to run the workflow as if a specific push to master event occured
act push

# List actions for a specific event (here the event is push)
act push -l

# Get Act to run a specific job
act -j review 

# The graph option draws the available workflow jobs structure in your terminal as a graph.
act --graph
 ╭─────────────╮ ╭──────────────────────────────────╮
 │ AI Reviewer │ │ Analyze (${{ matrix.language }}) │
 ╰─────────────╯ ╰──────────────────────────────────╯

# pass secrets into a job so that the GitHub action can consume them
act -s MY_TOKEN_SECRET=<token> -s MY_NETLIFY_SITE_ID=<site_id> 

# run a GitHub action that uses artifacts between jobs
act --artifact-server-path /tmp/artifacts push
```









