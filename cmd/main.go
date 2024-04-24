package main

import (
	"ai-codereview/pkg/application"
	"fmt"
	"os"
)

func main() {
	model_retries := os.Getenv("model_retries")
	comment, _ := application.CodeReview()
	fmt.Printf("------------Comment: %s", comment)
	fmt.Printf("------------Comment: %s", model_retries)
}
