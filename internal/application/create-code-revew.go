package application

import (
	"context"
	"fmt"
	"log"

	"ai-codereview/internal/domain"

	"github.com/google/go-github/v61/github"
)

type Context struct {
	Owner             string
	Repository        string
	PullRequestNumber int
	Token             string
}

func CreateCodeReview(
	ModelRepository domain.ModelRepository,
	gitContext Context,
	githubClient *github.Client) {
	ctx := context.Background()
	comment_text := &github.IssueComment{
		Body: github.String(ModelRepository.GetComment()),
	}

	createdComment, _, err := githubClient.Issues.CreateComment(
		ctx, gitContext.Owner,
		gitContext.Repository,
		gitContext.PullRequestNumber,
		comment_text)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Created comment at %s\n", *createdComment.HTMLURL)
}
