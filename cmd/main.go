package main

import (
	"ai-codereview/pkg/application"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-- Start --:")
	repo := os.Getenv("GITHUB_REPOSITORY")
	repo_owner := os.Getenv("GITHUB_REPOSITORY_OWNER")
	eventName := os.Getenv("GITHUB_EVENT_NAME")
	workflow := os.Getenv("GITHUB_WORKFLOW")
	runNumber := os.Getenv("GITHUB_RUN_NUMBER")
	actor := os.Getenv("GITHUB_ACTOR")
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	head_ref := os.Getenv("GITHUB_HEAD_REF")
	ref := os.Getenv("GITHUB_REF")

	// debug data
	fmt.Printf("Repository: %s\n", repo)
	fmt.Printf("Event Name: %s\n", eventName)
	fmt.Printf("Workflow: %s\n", workflow)
	fmt.Printf("Run Number: %s\n", runNumber)
	fmt.Printf("Actor: %s\n", actor)
	fmt.Printf("Event Path: %s\n", eventPath)
	fmt.Printf("Repo owner: %s\n", repo_owner)
	fmt.Printf("Head Ref: %s\n", head_ref)
	fmt.Printf("Ref: %s\n", ref)

	comment, _ := application.CodeReview()
	fmt.Println("------------Comment:", comment)
}
