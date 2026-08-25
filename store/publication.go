package store

import "example.com/casescript/domain"

func (r *Repository) PutPublication(value domain.Publication) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(publicationBucket, value.CaseID, value)
}

func (r *Repository) GetPublication(caseID string) (domain.Publication, error) {
	var value domain.Publication
	return value, r.get(publicationBucket, caseID, &value)
}

func (r *Repository) ListPublications() ([]domain.Publication, error) {
	items := make([]domain.Publication, 0)
	err := r.list(publicationBucket, func() any { return &domain.Publication{} }, func(item any) {
		items = append(items, *(item.(*domain.Publication)))
	})
	return items, err
}
