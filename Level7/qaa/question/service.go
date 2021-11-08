package question

import (
	"errors"
)

var ErrInvalidArgument = errors.New("invalid argument")

type Service interface {
	//GetQuestionByID(id question.QuestionID) (Question, error)
	Questions() []Question
	//GetQuestionsByUser(user string) []Question
	NewQuestion(statement string, answer string, user string) (QuestionID, error)
	//UpdateQuestion(id question.QuestionID, newQuestion Question) Question
	//DeleteQuestion(id question.QuestionID) error
}

type service struct {
	questions Repository
}

func (s *service) NewQuestion(statement, answer, user string) (QuestionID, error) {
	if statement == "" || user == "" {
		return "", ErrInvalidArgument
	}

	id := NextQuestionID()

	c := New(id, statement, answer, user)

	if err := s.questions.Store(c); err != nil {
		return "", err
	}

	return c.QuestionID, nil
}

func NewService(questions Repository) Service {
	return &service{
		questions: questions,
	}
}

func (s *service) Questions() []Question {
	var result []Question
	for _, v := range s.questions.FindAll() {
		result = append(result, Question{
			QuestionID:    v.QuestionID,
			Statement:     v.Statement,
			Answer:        v.Answer,
			CreatedByUser: v.CreatedByUser,
		})
	}
	return result
}
