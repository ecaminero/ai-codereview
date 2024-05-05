package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v61/github"
)

type Event struct {
	Issue *github.Issue `json:"issue"`
}

func CodeReview(owner string, repo string, number int) (string, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	body := `This is an example comment ` + time.Now().Format("2006-01-02 15:04:05")
	comment := &github.IssueComment{Body: github.String(body)}
	createdComment, _, err := client.Issues.CreateComment(ctx, owner, repo, number, comment)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
		return "", err
	}
	fmt.Printf("Created comment at %s\n", *createdComment.HTMLURL)

	return body, nil
}
