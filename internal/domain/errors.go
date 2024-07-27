package domain

import "fmt"

var ErrAIBotNotFound = fmt.Errorf("AI Bot not exist")
var ErrPullRequestFormat = fmt.Errorf("error converting PR number to int")
