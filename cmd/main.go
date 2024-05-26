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

func main() {
	fmt.Println("---- Start ----")
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

	// debug data
	fmt.Printf("Repository: %s\n", repository)
	fmt.Printf("Event Name: %s\n", eventName)
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
