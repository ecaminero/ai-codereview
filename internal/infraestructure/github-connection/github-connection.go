package github_connection

import (
	"context"
	"github.com/google/go-github/v61/github"
)

/*
Debe a partir de un PR, obtener todos los commits

obtener los archivos cambiados

con esos commits debe

*/

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
		ctx,
		receiver.RepoOwner,
		receiver.RepositoryName,
		receiver.PullRequestNumber,
		comment_text)

	/*
	* esto es otro metodo CreateReview
	 */

	listOptions := &github.ListOptions{PerPage: 100}
	commits, _, err := receiver.client.PullRequests.ListCommits(
		ctx,
		receiver.RepoOwner,
		receiver.RepositoryName,
		receiver.PullRequestNumber,
		listOptions,
	)
	if err != nil {
		return err
	}

	/*
		receiver.client.Repositories.GetCommit(
			ctx,
			receiver.RepoOwner,
			receiver.RepositoryName,

			listOptions,
		)
	*/

	changes := make([]string, 0)
	for _, commit := range commits {
		for _, file := range commit.Files {
			changes = append(changes, file.GetBlobURL())
		}
	}

	print(changes)
	/*
		fin de metodo CreateReview
	*/

	return err
}

/*

func (receiver *GithubConnection) CreateReview() ([]string, error) {
	ctx := context.Background()

	return changes, nil
}

*/
