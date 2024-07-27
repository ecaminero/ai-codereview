package domain

import "time"

type AIModel struct {
	Name string
}

func (m *AIModel) GetComment() string {
	return `Code Review ` + time.Now().Format("2006-01-02 15:04:05")
}
