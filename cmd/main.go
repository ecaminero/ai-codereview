package main

import (
	"ai-codereview/internal/application"
	github_connection "ai-codereview/internal/infraestructure/github-connection"
	stub_persistence "ai-codereview/internal/infraestructure/stub"

	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func print_all_variables() {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if strings.HasPrefix(pair[0], "GITHUB_") {
			fmt.Println(pair[0], ":", pair[1])
		}
	}
}

func main() {
	print_all_variables()
	token := os.Getenv("GITHUB_TOKEN")
	githubRepositoryName := strings.Join(strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1:], "")
	repoOwner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	pullRequestNumber, err := strconv.Atoi(os.Getenv("GITHUB_PR_NUMBER"))

	if err != nil {
		log.Fatal(err)
	}

	params := github_connection.GithubConnectionParams{
		RepositoryName:    githubRepositoryName,
		Token:             token,
		RepoOwner:         repoOwner,
		PullRequestNumber: pullRequestNumber,
	}
	githubCodeRepositoryProvider, err := github_connection.NewGithubConnection(params)
	if err != nil {
		log.Fatal(err)
	}

	aiModel := stub_persistence.NewStubModelRepository()

	app := application.NewApp(githubCodeRepositoryProvider, aiModel)

	switch eventName {
	case "pull_request_target", "pull_request":
		app.CreateCodeReview()
	case "pull_request_review_comment":
		// application.HandleCommentReview(repoOwner, githubRepositoryName, prNumber)
		print("HandleCommentReview")
	default:
		log.Fatalf("Skipped: current event is %s, only support pull_request event", eventName)
	}
}
