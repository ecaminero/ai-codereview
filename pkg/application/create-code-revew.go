package application

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v61/github"
	"golang.org/x/oauth2"
)

func CodeReview() (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// owner and repo correspond to the Github repository you want to interact with
	owner := "ecaminero"
	repo := os.Getenv("GITHUB_REPOSITORY")
	ref := os.Getenv("GITHUB_REF")
	splitRef := strings.Split(ref, "/")
	number, err := strconv.Atoi(splitRef[2])

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}
	body := `This is a multi-line string.
	It can contain multiple lines.
	Each new line is represented by a new line in the string.`
	comment := &github.IssueComment{Body: github.String(body)}
	comment, _, err = client.Issues.CreateComment(ctx, owner, repo, number, comment)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "", err
	}

	fmt.Printf("Created comment at %s\n", *comment.HTMLURL)

	return body, nil
}
