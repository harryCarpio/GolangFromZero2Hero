package question

import (
	"context"
	"encoding/json"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type errorer interface {
	error() error
}

func MakeHandler(qs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	questionHandler := kithttp.NewServer(
		makeQuestionEndpoint(qs),
		decodeQuestionRequest,
		encodeResponse,
		opts...,
	)

	listQuestionHandler := kithttp.NewServer(
		makeListQuestionsEndpoint(qs),
		decodeListQuestionsRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/questions/v1/question", questionHandler).Methods("POST")
	r.Handle("/questions/v1/question", listQuestionHandler).Methods("GET")

	return r
}

func decodeListQuestionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listQuestionsRequest{}, nil
}

func decodeQuestionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Statement string `json:"statement"`
		Answer    string `json:"answer"`
		User      string `json:"user"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return questionRequest{
		Statement: body.Statement,
		Answer:    body.Answer,
		User:      body.User,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
