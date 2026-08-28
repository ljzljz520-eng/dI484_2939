package store

import "example.com/casescript/domain"

func (r *Repository) PutReview(value domain.ReviewRequest) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(reviewBucket, value.ID, value)
}

func (r *Repository) GetReview(id string) (domain.ReviewRequest, error) {
	var value domain.ReviewRequest
	return value, r.get(reviewBucket, id, &value)
}

func (r *Repository) ListReviews(caseID string) ([]domain.ReviewRequest, error) {
	items := make([]domain.ReviewRequest, 0)
	err := r.list(reviewBucket, func() any { return &domain.ReviewRequest{} }, func(item any) {
		value := *(item.(*domain.ReviewRequest))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	return items, err
}

func (r *Repository) DeleteReview(id string) error { return r.delete(reviewBucket, id) }
