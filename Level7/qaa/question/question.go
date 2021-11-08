package question

import (
	"errors"
	"github.com/pborman/uuid"
	"strings"
)

type QuestionID string

type Question struct {
	QuestionID    QuestionID `json:"id,omitempty"`
	Statement     string     `json:"statement"`
	Answer        string     `json:"answer"`
	CreatedByUser string     `json:"createdByUser"`
}

type Repository interface {
	Store(question *Question) error
	Find(id QuestionID) (*Question, error)
	FindAll() []*Question
}

var ErrUnknown = errors.New("unknown question")

func NextQuestionID() QuestionID {
	return QuestionID(strings.Split(strings.ToUpper(uuid.New()), "-")[0])
}

func New(id QuestionID, stm string, ans string, usr string) *Question {
	return &Question{
		QuestionID:    id,
		Statement:     stm,
		Answer:        ans,
		CreatedByUser: usr,
	}
}
