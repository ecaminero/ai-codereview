package github_connection

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v61/github"
)

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

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	if eventName == "" {
		return nil, fmt.Errorf("GITHUB_EVENT_NAME environment variable is not set")
	}
	GithubClient := github.NewClient(nil).WithAuthToken(token)

	return &GithubConnection{
		client:            GithubClient,
		RepositoryName:    githubRepositoryName,
		RepoOwner:         repoOwner,
		PullRequestNumber: pullRequestNumber,
		eventName:         eventName,
	}, nil
}

type GithubConnection struct {
	client            *github.Client
	RepositoryName    string
	RepoOwner         string
	PullRequestNumber int
	eventName         string
}

func (receiver *GithubConnection) GetEventName() string {
	return receiver.eventName
}

func (receiver *GithubConnection) GetRepository() string {
	return receiver.RepositoryName
}

func (receiver GithubConnection) CreateComment(comment string) error {
	ctx := context.Background()

	comment_text := &github.IssueComment{
		Body: github.String(comment),
	}

	_, _, err := receiver.client.Issues.CreateComment(
		ctx, receiver.RepoOwner,
		receiver.RepositoryName,
		receiver.PullRequestNumber,
		comment_text)
	return err
}
