package store

import (
	"example.com/casescript/domain"
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func (r *Repository) delete(bucket []byte, key string) error {
	if key == "" {
		return fmt.Errorf("record key is required")
	}
	return r.db.Update(func(tx *bolt.Tx) error {
		value := tx.Bucket(bucket)
		if value == nil {
			return fmt.Errorf("bucket %q not found", string(bucket))
		}
		return value.Delete([]byte(key))
	})
}

func (r *Repository) Exists(bucket []byte, key string) (bool, error) {
	if key == "" {
		return false, fmt.Errorf("record key is required")
	}
	var exists bool
	err := r.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(bucket)
		if value == nil {
			return fmt.Errorf("bucket %q not found", string(bucket))
		}
		exists = value.Get([]byte(key)) != nil
		return nil
	})
	return exists, err
}

func (r *Repository) UpdateCaseAndPublication(value domain.LegalCase, publication domain.Publication) error {
	if err := value.Validate(); err != nil {
		return err
	}
	if err := publication.Validate(); err != nil {
		return err
	}
	caseData, err := encode(value)
	if err != nil {
		return err
	}
	publicationData, err := encode(publication)
	if err != nil {
		return err
	}
	return r.db.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(caseBucket).Put([]byte(value.ID), caseData); err != nil {
			return err
		}
		return tx.Bucket(publicationBucket).Put([]byte(value.ID), publicationData)
	})
}
