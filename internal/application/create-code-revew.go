package application

import (
	"log"
)

type Context struct {
	Owner             string
	Repository        string
	PullRequestNumber int
	Token             string
}

func (a *App) CreateCodeReview() {
	comment := a.aiModel.GetComment()
	err := a.codeRepositoryProvider.CreateComment(comment)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
