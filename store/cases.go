package store

import (
	"example.com/casescript/domain"
	bolt "go.etcd.io/bbolt"
)

func (r *Repository) PutCase(value domain.LegalCase) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(caseBucket, value.ID, value)
}

func (r *Repository) GetCase(id string) (domain.LegalCase, error) {
	var value domain.LegalCase
	return value, r.get(caseBucket, id, &value)
}

func (r *Repository) ListCases() ([]domain.LegalCase, error) {
	items := make([]domain.LegalCase, 0)
	err := r.list(caseBucket, func() any { return &domain.LegalCase{} }, func(item any) {
		items = append(items, *(item.(*domain.LegalCase)))
	})
	return items, err
}

func (r *Repository) DeleteCase(id string) error {
	return r.db.Update(func(tx *bolt.Tx) error { return tx.Bucket(caseBucket).Delete([]byte(id)) })
}
