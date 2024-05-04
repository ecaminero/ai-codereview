package application

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v61/github"
	"golang.org/x/oauth2"
)

type Event struct {
	Issue *github.Issue `json:"issue"`
}

func CodeReview() (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	dateTime := time.Now().Format("2006-01-02 15:04:05")

	// owner and repo correspond to the Github repository you want to interact with
	owner := "ecaminero"
	repo := os.Getenv("GITHUB_REPOSITORY")
	// Get the event from the GITHUB_EVENT_PATH
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	data, err := os.ReadFile(eventPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	var event Event
	err = json.Unmarshal(data, &event)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}
	number := event.Issue.GetNumber()
	body := `This is an example comment` + dateTime
	comment := &github.IssueComment{Body: github.String(body)}
	createdComment, _, err := client.Issues.CreateComment(ctx, owner, repo, number, comment)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}
	fmt.Printf("Created comment at %s\n", *createdComment.HTMLURL)

	return body, nil
}
