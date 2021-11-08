package question

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type questionRequest struct {
	Statement string
	Answer    string
	User      string
}

type questionResponse struct {
	ID  QuestionID `json:"question_id,omitempty"`
	Err error      `json:"error,omitempty"`
}

type listQuestionsRequest struct {
}

type listQuestionsResponse struct {
	Questions []Question `json:"questions,omitempty"`
	Err       error      `json:"error,omitempty"`
}

func makeQuestionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(questionRequest)
		id, err := s.NewQuestion(req.Statement, req.Answer, req.User)
		return questionResponse{ID: id, Err: err}, nil
	}
}

func makeListQuestionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listQuestionsRequest)
		return listQuestionsResponse{Questions: s.Questions(), Err: nil}, nil
	}
}
