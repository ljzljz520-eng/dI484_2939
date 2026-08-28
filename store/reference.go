package store

import "example.com/casescript/domain"

func (r *Repository) PutReference(value domain.CaseReference) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(referenceBucket, value.ID, value)
}

func (r *Repository) GetReference(id string) (domain.CaseReference, error) {
	var value domain.CaseReference
	return value, r.get(referenceBucket, id, &value)
}

func (r *Repository) ListReferences(caseID string) ([]domain.CaseReference, error) {
	items := make([]domain.CaseReference, 0)
	err := r.list(referenceBucket, func() any { return &domain.CaseReference{} }, func(item any) {
		value := *(item.(*domain.CaseReference))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	return items, err
}

func (r *Repository) DeleteReference(id string) error { return r.delete(referenceBucket, id) }
