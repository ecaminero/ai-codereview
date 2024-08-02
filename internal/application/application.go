package application

import "ai-codereview/internal/domain"

type ICodeRepositoryProvider interface {
	GetEventName() string
	GetRepository() string
	CreateComment(comment string) error
}

type App struct {
	codeRepositoryProvider ICodeRepositoryProvider
	aiModel                domain.IAIModel
}

func NewApp(
	codeRepositoryProvider ICodeRepositoryProvider,
	model domain.IAIModel) *App {
	return &App{
		codeRepositoryProvider: codeRepositoryProvider,
		aiModel:                model,
	}
}
