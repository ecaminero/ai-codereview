package main

import (
	"ai-codereview/pkg/application"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPRNumber(ref string) (int, error) {
	parts := strings.Split(ref, "/")
	prNumberStr := parts[len(parts)-2]
	return strconv.Atoi(prNumberStr)
}

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
	workflow := os.Getenv("GITHUB_WORKFLOW")
	runNumber := os.Getenv("GITHUB_RUN_NUMBER")
	actor := os.Getenv("GITHUB_ACTOR")
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	head_ref := os.Getenv("GITHUB_HEAD_REF")
	base_ref := os.Getenv("GITHUB_BASE_REF")
	ref := os.Getenv("GITHUB_REF")
	ref_name := os.Getenv("GITHUB_REF_NAME")

	switch eventName {
	case "pull_request_target", "pull_request":
		fmt.Println("Code review for pull request")
	case "pull_request_review_comment":
		fmt.Println("A pull request review comment event occurred")
	default:
		fmt.Println("This event is not supported")
	}

	// debug data
	fmt.Printf("Repository: %s\n", repository)
	fmt.Printf("Workflow: %s\n", workflow)
	fmt.Printf("Run Number: %s\n", runNumber)
	fmt.Printf("Actor: %s\n", actor)
	fmt.Printf("Event Path: %s\n", eventPath)
	fmt.Printf("Repo owner: %s\n", repo_owner)
	fmt.Printf("Head Ref: %s\n", head_ref)
	fmt.Printf("Ref: %s\n", ref)
	fmt.Printf("RefName: %s\n", ref_name)

	fmt.Printf("Github Base ref: %s\n", base_ref)

	// Convert the pull request number to an integer
	prNumber, err := getPRNumber(ref)
	if err != nil {
		fmt.Println("Error converting pull request number to integer:", err)
		return
	}

	fmt.Printf("Pull Request Number: %d\n", prNumber)

	comment, _ := application.CodeReview(repo_owner, repository, prNumber)
	fmt.Println("------------Comment:", comment)
	fmt.Println("---- END ----")
}
