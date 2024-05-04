package main

import (
	"ai-codereview/pkg/application"
	"fmt"
	"os"
)

func main() {
	fmt.Println("-- Start --:")
	model_retries := os.Getenv("model_retries")
	comment, _ := application.CodeReview()
	fmt.Println("------------Comment:", comment)
	fmt.Println("------------Comment:", model_retries)
}
