package main

import (
	"ai-codereview/internal/application"
	"fmt"
)

func main() {

	comment := application.GetRandomComments()
	fmt.Printf("Comment: %s", comment)
}
