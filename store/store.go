package store

import (
	"encoding/json"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

var (
	caseBucket        = []byte("legal_cases")
	chapterBucket     = []byte("chapters")
	noteBucket        = []byte("classroom_notes")
	publicationBucket = []byte("publications")
	lessonBucket      = []byte("lesson_plans")
	referenceBucket   = []byte("case_references")
	reviewBucket      = []byte("review_requests")
	timelineBucket    = []byte("timeline_events")
	readingBucket     = []byte("reading_packets")
	vocabularyBucket  = []byte("vocabulary_entries")
	auditBucket       = []byte("audit_entries")
	feedbackBucket    = []byte("student_feedback")
)

type Repository struct {
	db *bolt.DB
}

func Open(path string) (*Repository, error) {
	if path == "" {
		return nil, fmt.Errorf("database path is required")
	}
	if err := os.MkdirAll(filepathDir(path), 0755); err != nil {
		return nil, err
	}
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	r := &Repository{db: db}
	if err := db.Update(func(tx *bolt.Tx) error {
		for _, name := range [][]byte{caseBucket, chapterBucket, noteBucket, publicationBucket, lessonBucket, referenceBucket, reviewBucket, timelineBucket, readingBucket, vocabularyBucket, auditBucket, feedbackBucket} {
			if _, err := tx.CreateBucketIfNotExists(name); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		db.Close()
		return nil, err
	}
	return r, nil
}

func filepathDir(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			if i == 0 {
				return "/"
			}
			return path[:i]
		}
	}
	return "."
}

func (r *Repository) Close() error {
	if r == nil || r.db == nil {
		return nil
	}
	return r.db.Close()
}

func encode(value any) ([]byte, error) { return json.Marshal(value) }

func decode(data []byte, target any) error {
	if len(data) == 0 {
		return fmt.Errorf("record not found")
	}
	return json.Unmarshal(data, target)
}

func (r *Repository) put(bucket []byte, key string, value any) error {
	if key == "" {
		return fmt.Errorf("record key is required")
	}
	payload, err := encode(value)
	if err != nil {
		return err
	}
	return r.db.Update(func(tx *bolt.Tx) error { return tx.Bucket(bucket).Put([]byte(key), payload) })
}

func (r *Repository) get(bucket []byte, key string, target any) error {
	return r.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(bucket).Get([]byte(key))
		if value == nil {
			return fmt.Errorf("record %q not found", key)
		}
		copyValue := append([]byte(nil), value...)
		return decode(copyValue, target)
	})
}

func (r *Repository) list(bucket []byte, factory func() any, appendValue func(any)) error {
	return r.db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(bucket).ForEach(func(_, value []byte) error {
			if value == nil {
				return nil
			}
			item := factory()
			if err := decode(value, item); err != nil {
				return err
			}
			appendValue(item)
			return nil
		})
	})
}
