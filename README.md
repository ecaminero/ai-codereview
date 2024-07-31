<a name="readme-top"></a> 

<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]

<a href="https://github.com/ecaminero/ai-codereview">
  <img width="300px" src="https://cdn.githubraw.com/ecaminero/ai-codereview/main/docs/images/AI-bot-codereview.png" alt="Logo" width="800" />
</a>
  
## AI CodeReview
AI CodeReview 'ai-review' is an AI-based code reviewer and summarizer for GitHub pull requests. You can use your own AI model, its designed to be used as a GitHub Action and can be configured to run on every pull request and review comments

</div>

### Install:
```bash
go mod tidy

# Build Image
GOOS=linux GOARCH=amd64 go build -o dist/main_linux cmd/main.go
```

#### Environment variables

- `GITHUB_TOKEN`: This should already be available to the GitHub Action
  environment. This is used to add comments to the pull request.

### Models:
* [codellama](https://ollama.com/library/codellama) - A large language model that can use text prompts to generate and discuss code.
    Code Llama is a model for generating and discussing code, built on top of Llama 2. It’s designed to make workflows faster and efficient for developers and make it easier for people to learn how to code. It can generate both code and natural language about code. Code Llama supports many of the most popular programming languages used today, including Python, C++, Java, PHP, Typescript (Javascript), C#, Bash and more.

## Contributors
**¡Thanks to all the collaborators who have made this project possible!**

[![Contributors](https://contrib.rocks/image?repo=ecaminero/ai-codereview)](https://github.com/ecaminero/ai-codereview/graphs/contributors)

<p align="right">(<a href="#readme-top">Go back up</a>)</p>
## 🛠️ Stack
- [![Go][go-badge]][go-url] - An open-source programming language supported by Google.

<p align="right">(<a href="#readme-top">Go back up</a>)</p>

[go-url]: https://go.dev/
[githubaction-url]: https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions
[go-badge]: https://img.shields.io/badge/go-fff?style=for-the-badge&logo=go&logoColor=bd303a&color=35256


[contributors-url]: https://github.com/ecaminero/ai-codereview/graphs/contributors
[contributors-shield]: https://img.shields.io/github/contributors/ecaminero/ai-codereview.svg?style=for-the-badge

[forks-url]: https://github.com/ecaminero/ai-codereview/network/members
[forks-shield]: https://img.shields.io/github/forks/ecaminero/ai-codereview.svg?style=for-the-badge

[stars-url]: https://github.com/ecaminero/ai-codereview/stargazers
[stars-shield]: https://img.shields.io/github/stars/ecaminero/ai-codereview.svg?style=for-the-badge

[issues-shield]: https://img.shields.io/github/issues/ecaminero/ai-codereview.svg?style=for-the-badge
[issues-url]: https://github.com/ecaminero/ai-codereview/issues
