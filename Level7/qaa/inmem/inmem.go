package inmem

import (
	"sync"

	"qaa/question"
)

type questionRepository struct {
	mtx       sync.RWMutex
	questions map[question.QuestionID]*question.Question
}

func (r *questionRepository) Store(q *question.Question) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.questions[q.QuestionID] = q
	return nil
}

func (r *questionRepository) Find(id question.QuestionID) (*question.Question, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.questions[id]; ok {
		return val, nil
	}
	return nil, question.ErrUnknown
}

func (r *questionRepository) FindAll() []*question.Question {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	q := make([]*question.Question, 0, len(r.questions))
	for _, val := range r.questions {
		q = append(q, val)
	}
	return q
}

func NewQuestionRepository() question.Repository {
	return &questionRepository{
		questions: make(map[question.QuestionID]*question.Question),
	}
}
