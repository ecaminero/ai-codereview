package github_connection

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v61/github"
)

type GithubConnection struct {
	Client            *github.Client `json:"client"`
	RepositoryName    string         `json:"repository_name"`
	RepoOwner         string         `json:"repo_owner,omitempty"`
	PullRequestNumber int            `json:"pull_request_number"`
	EventName         string         `json:"event_name"`
}

type FileChange struct {
	Filename    string
	LineNumbers []int
	Content     string
}

func NewGithubConnection() (*GithubConnection, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable is not set")
	}

	repoFullName := os.Getenv("GITHUB_REPOSITORY")
	if repoFullName == "" {
		return nil, fmt.Errorf("GITHUB_REPOSITORY environment variable is not set")
	}
	repoParts := strings.Split(repoFullName, "/")
	if len(repoParts) != 2 {
		return nil, fmt.Errorf("invalid GITHUB_REPOSITORY format: %s", repoFullName)
	}
	repoOwner := repoParts[0]
	githubRepositoryName := repoParts[1]

	pullRequestNumberStr := os.Getenv("GITHUB_PR_NUMBER")
	if pullRequestNumberStr == "" {
		return nil, fmt.Errorf("GITHUB_PR_NUMBER environment variable is not set")
	}
	pullRequestNumber, err := strconv.Atoi(pullRequestNumberStr)
	if err != nil {
		return nil, fmt.Errorf("invalid GITHUB_PR_NUMBER: %w", err)
	}

	eventName := os.Getenv("GITHUB_EVENT_NAME")
	if eventName == "" {
		return nil, fmt.Errorf("GITHUB_EVENT_NAME environment variable is not set")
	}
	GithubClient := github.NewClient(nil).WithAuthToken(token)

	return &GithubConnection{
		Client:            GithubClient,
		RepositoryName:    githubRepositoryName,
		RepoOwner:         repoOwner,
		PullRequestNumber: pullRequestNumber,
		EventName:         eventName,
	}, nil
}

func (receiver *GithubConnection) GetEventName() string {
	return receiver.EventName
}

func (receiver *GithubConnection) GetRepository() string {
	return receiver.RepositoryName
}

func (receiver *GithubConnection) GetPullRequestChanges() ([]FileChange, error) {
	var ctx = context.Background()
	var opt = &github.ListOptions{PerPage: 100}
	var allAdditions []FileChange

	for {
		files, resp, err := receiver.Client.PullRequests.ListFiles(
			ctx, receiver.RepoOwner, receiver.RepositoryName, receiver.PullRequestNumber, opt)
		if err != nil {
			return nil, fmt.Errorf("error listing pull request files: %w", err)
		}
		for _, file := range files {
			if file.Patch == nil {
				continue
			}
			additions := parseAdditions(file)
			if len(additions.LineNumbers) > 0 {
				allAdditions = append(allAdditions, additions)
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allAdditions, nil
}

func (receiver *GithubConnection) CreateComment(comment string) error {
	// ctx := context.Background()

	allFiles, err := receiver.GetPullRequestChanges()
	if err != nil {
		return fmt.Errorf("error getting commits changes: %w", err)
	}

	for _, file := range allFiles {
		fmt.Printf("File: %s, RawURL: %s\n", file.Filename, file.Content)
		// commentData := &github.PullRequestComment{
		// 	CommitID: github.String(os.Getenv("GITHUB_SHA")),
		// 	Path:     github.String(*file.BlobURL),
		// 	Body:     github.String(comment),
		// }
		// var prComment, _, err = receiver.Client.PullRequests.CreateComment(
		// 	ctx,
		// 	receiver.RepoOwner,
		// 	receiver.RepositoryName,
		// 	receiver.PullRequestNumber,
		// 	commentData)
		// if err != nil {
		// 	return fmt.Errorf("error PullRequests: %w", err)
		// }
	}

	return nil
}

func parseAdditions(file *github.CommitFile) FileChange {
	lines := strings.Split(*file.Patch, "\n")
	var lineNumbers []int
	var content strings.Builder

	lineNumber := 0

	for _, line := range lines {
		if strings.HasPrefix(line, "@@") {
			parts := strings.Split(line, " ")
			if len(parts) >= 3 {
				lineNumberStr := strings.TrimPrefix(parts[2], "+")
				lineNumber, _ = strconv.Atoi(strings.Split(lineNumberStr, ",")[0])
			}
			continue
		}

		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			lineNumbers = append(lineNumbers, lineNumber)
		}

		if !strings.HasPrefix(line, "-") {
			lineNumber++
			content.WriteString(line + "\n")

		}
	}
	fmt.Println(lineNumbers)

	return FileChange{
		Filename:    *file.Filename,
		LineNumbers: lineNumbers,
		Content:     content.String(),
	}
}
