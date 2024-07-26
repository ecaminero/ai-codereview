package main

import (
	"ai-codereview/internal/application"
	"ai-codereview/internal/domain"
	stub_persistence "ai-codereview/internal/infraestructure/stub"

	"github.com/google/go-github/v61/github"

	"context"
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
	repository := strings.Join(strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1:], "")
	repoOwner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ModelRepository := stub_persistence.NewStubModelRepository()
	ctx := context.Background()
	GithubClient := github.NewClient(nil).WithAuthToken(token)
	repo, _, err := GithubClient.Repositories.Get(ctx, repoOwner, repository)
	if err != nil {
		fmt.Println(err)
		return
	}
	prNumber, err := strconv.Atoi(os.Getenv("GITHUB_PR_NUMBER"))
	if err != nil {
		fmt.Println(domain.ErrPullRequestFormat.Error())
	}

	context := application.Context{
		Repository:        *repo.Name,
		Owner:             repoOwner,
		PullRequestNumber: prNumber,
	}

	switch eventName {
	case "pull_request_target", "pull_request":
		application.CreateCodeReview(ModelRepository, context, GithubClient)
	case "pull_request_review_comment":
		application.HandleCommentReview(repoOwner, repository, prNumber)
	default:
		log.Fatalf("Skipped: current event is %s, only support pull_request event", eventName)
	}

	if err != nil {
		log.Fatalf("Error handling : %s", err)
	}
	fmt.Println("---- END ----")
}
