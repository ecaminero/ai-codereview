# Documentation for Llama

## Model Documentation
https://llama.meta.com/docs/model-cards-and-prompt-formats/llama3_1/

https://cdn.githubraw.com/ecaminero/ai-codereview/main/docs/images/AI-bot-codereview.png


### Executions
'<|begin_of_text|><|start_header_id|>system<|end_header_id|>

  You are `@codereview` (aka `github-actions[bot]`), a language model 
      trained by Meta. Your purpose is to act as a highly experienced 
      software engineer and provide a thorough review of the code hunks
      and suggest code snippets to improve key areas such as:
        - Logic
        - Security
        - Performance
        - Data races
        - Consistency
        - Error handling
        - Maintainability
        - Modularity
        - Complexity
        - Optimization
        - Best practices: DRY, SOLID, KISS

      Do not comment on minor code style issues, missing 
      comments/documentation. Identify and resolve significant 
      concerns to improve overall code quality while deliberately 
      disregarding minor issues.<|eot_id|><|start_header_id|>user<|end_header_id|>


import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v61/github"
)

type GithubConnectionParams struct {
	RepositoryName    string
	Token             string
	RepoOwner         string
	PullRequestNumber int
type GithubConnection struct {
	Client            *github.Client `json:"client"`
	RepositoryName    string         `json:"repository_name"`
	RepoOwner         string         `json:"repo_owner,omitempty"`
	PullRequestNumber int            `json:"pull_request_number"`
	EventName         string         `json:"event_name"`
}

type FileChange struct {
	Filename    string
	LineNumbers []int
	Content     string
}

func NewGithubConnection(params GithubConnectionParams) (*GithubConnection, error) {
func NewGithubConnection() (*GithubConnection, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable is not set")
	}

	repoFullName := os.Getenv("GITHUB_REPOSITORY")
	if repoFullName == "" {
		return nil, fmt.Errorf("GITHUB_REPOSITORY environment variable is not set")
	}
	repoParts := strings.Split(repoFullName, "/")
	if len(repoParts) != 2 {
		return nil, fmt.Errorf("invalid GITHUB_REPOSITORY format: %s", repoFullName)
	}
	repoOwner := repoParts[0]
	githubRepositoryName := repoParts[1]

	pullRequestNumberStr := os.Getenv("GITHUB_PR_NUMBER")
	if pullRequestNumberStr == "" {
		return nil, fmt.Errorf("GITHUB_PR_NUMBER environment variable is not set")
	}
	pullRequestNumber, err := strconv.Atoi(pullRequestNumberStr)
	if err != nil {
		return nil, fmt.Errorf("invalid GITHUB_PR_NUMBER: %w", err)
	}

	GithubClient := github.NewClient(nil).WithAuthToken(params.Token)
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	if eventName == "" {
		return nil, fmt.Errorf("GITHUB_EVENT_NAME environment variable is not set")
	}
	GithubClient := github.NewClient(nil).WithAuthToken(token)

	return &GithubConnection{
		client:            GithubClient,
		RepositoryName:    params.RepositoryName,
		RepoOwner:         params.RepoOwner,
		PullRequestNumber: params.PullRequestNumber,
		Client:            GithubClient,
		RepositoryName:    githubRepositoryName,
		RepoOwner:         repoOwner,
		PullRequestNumber: pullRequestNumber,
		EventName:         eventName,
	}, nil
}

type GithubConnection struct {
	client            *github.Client
	RepositoryName    string
	RepoOwner         string
	PullRequestNumber int
func (receiver *GithubConnection) GetEventName() string {
	return receiver.EventName
}

func (receiver *GithubConnection) GetRepositoryName() string {
func (receiver *GithubConnection) GetRepository() string {
	return receiver.RepositoryName
}

func (receiver *GithubConnection) GetRepository() string {
	return receiver.RepositoryName
func (receiver *GithubConnection) GetPullRequestChanges() ([]FileChange, error) {
	var ctx = context.Background()
	var opt = &github.ListOptions{PerPage: 100}
	var allAdditions []FileChange

	for {
		files, resp, err := receiver.Client.PullRequests.ListFiles(
			ctx, receiver.RepoOwner, receiver.RepositoryName, receiver.PullRequestNumber, opt)
		if err != nil {
			return nil, fmt.Errorf("error listing pull request files: %w", err)
		}
		for _, file := range files {
			if file.Patch == nil {
				continue
			}
			additions := parseAdditions(file)
			if len(additions.LineNumbers) > 0 {
				allAdditions = append(allAdditions, additions)
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allAdditions, nil
}


<|eot_id|><|start_header_id|>assistant<|end_header_id|>
'