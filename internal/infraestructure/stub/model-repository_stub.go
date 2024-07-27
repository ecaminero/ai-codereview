package stub_persistence

import (
	"time"
)

func NewStubModelRepository() *StubModelRepository {
	return &StubModelRepository{}
}

type StubModelRepository struct{}

func (r *StubModelRepository) GetComment() (comment string) {
	comment = `Code Review ` + time.Now().Format("2006-01-02 15:04:05")
	return comment
}
