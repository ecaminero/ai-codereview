package application

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v61/github"
)

type Event struct {
	Issue *github.Issue `json:"issue"`
}

func CodeReview() (string, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)
	dateTime := time.Now().Format("2006-01-02 15:04:05")

	// owner and repo correspond to the Github repository you want to interact with
	// Debug vars
	owner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	repo := os.Getenv("GITHUB_REPOSITORY")

	// Get the GITHUB_EVENT_PATH environment variable
	eventPath := os.Getenv("GITHUB_EVENT_PATH")

	// Read the event payload file
	data, err := os.ReadFile(eventPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}
	// Unmarshal the event payload into an Event struct
	var event Event
	fmt.Println("Data: ", data)
	err = json.Unmarshal(data, &event)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	// Check if event.Issue and event.Issue.Title are not nil
	if event.Issue != nil && event.Issue.Title != nil {
		// Get the pull request title
		title := *event.Issue.Title
		fmt.Printf("Pull request title: %s\n", title)
	} else {
		fmt.Println("Issue or Issue Title is nil")
		return "", fmt.Errorf("issue or Issue Title is nil")
	}

	number := 9

	body := `This is an example comment` + dateTime
	comment := &github.IssueComment{Body: github.String(body)}
	createdComment, _, err := client.Issues.CreateComment(ctx, owner, repo, number, comment)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
		return "", err
	}
	fmt.Printf("Created comment at %s\n", *createdComment.HTMLURL)

	return body, nil
}
