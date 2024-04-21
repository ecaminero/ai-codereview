package main

import (
	"ai-codereview/pkg/application"
	"fmt"
)

func main() {

	comment := application.GetRandomComments()
	fmt.Printf("Comment: %s", comment)
}
