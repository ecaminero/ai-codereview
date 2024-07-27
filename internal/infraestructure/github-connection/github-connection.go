package github_connection

import (
	"context"
	"github.com/google/go-github/v61/github"
)

type GithubConnectionParams struct {
	RepositoryName    string
	Token             string
	RepoOwner         string
	PullRequestNumber int
}

func NewGithubConnection(params GithubConnectionParams) (*GithubConnection, error) {

	GithubClient := github.NewClient(nil).WithAuthToken(params.Token)

	return &GithubConnection{
		client:            GithubClient,
		RepositoryName:    params.RepositoryName,
		RepoOwner:         params.RepoOwner,
		PullRequestNumber: params.PullRequestNumber,
	}, nil
}

type GithubConnection struct {
	client            *github.Client
	RepositoryName    string
	RepoOwner         string
	PullRequestNumber int
}

func (receiver *GithubConnection) GetRepositoryName() string {
	return receiver.RepositoryName
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
