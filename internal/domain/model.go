package domain

import "time"

type Model struct {
	Name string
}

func (m *Model) GetComment() string {
	return `Code Review ` + time.Now().Format("2006-01-02 15:04:05")
}