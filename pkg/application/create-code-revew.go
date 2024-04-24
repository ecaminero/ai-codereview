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
	now := time.Now()
	dateTime := now.Format("2006-01-02 15:04:05")

	// owner and repo correspond to the Github repository you want to interact with
	owner := "ecaminero"
	repo := os.Getenv("GITHUB_REPOSITORY")
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	data, err := os.ReadFile(eventPath)
	if err != nil {
		return "", err
	}
	var event Event
	err = json.Unmarshal(data, &event)
	if err != nil {
		return "", err
	}

	number := event.Issue.GetNumber()

	body := `This is an example comment` + dateTime
	comment := &github.IssueComment{Body: github.String(body)}
	comment, _, err = client.Issues.CreateComment(ctx, owner, repo, number, comment)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	fmt.Printf("Created comment at %s\n", *comment.HTMLURL)

	return body, nil
}
