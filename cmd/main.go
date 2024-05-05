package main

import (
	"ai-codereview/pkg/application"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("-- Start --:")
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
	fmt.Printf("Github Base ref: %s\n", base_ref)

	comment, _ := application.CodeReview(repo_owner, repository, 9)
	fmt.Println("------------Comment:", comment)
}
