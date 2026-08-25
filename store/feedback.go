package store

import "example.com/casescript/domain"

func (r *Repository) PutFeedback(value domain.StudentFeedback) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(feedbackBucket, value.ID, value)
}

func (r *Repository) GetFeedback(id string) (domain.StudentFeedback, error) {
	var value domain.StudentFeedback
	return value, r.get(feedbackBucket, id, &value)
}

func (r *Repository) ListFeedback(caseID string) ([]domain.StudentFeedback, error) {
	items := make([]domain.StudentFeedback, 0)
	err := r.list(feedbackBucket, func() any { return &domain.StudentFeedback{} }, func(item any) {
		value := *(item.(*domain.StudentFeedback))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	return items, err
}
