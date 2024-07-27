package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v61/github"
)

func HandleCommentReview(owner string, repo string, number int) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	body := `Review Comment ` + time.Now().Format("2006-01-02 15:04:05")
	comment := &github.IssueComment{Body: github.String(body)}
	createdComment, _, err := client.Issues.CreateComment(ctx, owner, repo, number, comment)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("Created comment at %s\n", *createdComment.HTMLURL)

}
