package main

import (
	"ai-codereview/pkg/application"
	"fmt"
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
	repository := strings.Join(strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/")[1:], "")
	repo_owner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	eventName := os.Getenv("GITHUB_EVENT_NAME")

	prNumber, err := strconv.Atoi(os.Getenv("GITHUB_PR_NUMBER"))
	if err != nil {
		fmt.Println("Error converting PR number to int")

	}
	switch eventName {
	case "pull_request_target", "pull_request":
		fmt.Println("Code review for pull request")
	case "pull_request_review_comment":
		fmt.Println("A pull request review comment event occurred")
	default:
		fmt.Println("This event is not supported")
	}

	comment, _ := application.CodeReview(repo_owner, repository, prNumber)
	fmt.Println("------------Comment:", comment)
	fmt.Println("---- END ----")
}
