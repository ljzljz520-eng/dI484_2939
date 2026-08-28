package store

import (
	"example.com/casescript/domain"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"sort"
)

func (r *Repository) PutTimelineEvent(value domain.TimelineEvent) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(timelineBucket, value.ID, value)
}

func (r *Repository) GetTimelineEvent(id string) (domain.TimelineEvent, error) {
	var value domain.TimelineEvent
	return value, r.get(timelineBucket, id, &value)
}

func (r *Repository) ListTimeline(caseID string) ([]domain.TimelineEvent, error) {
	items := make([]domain.TimelineEvent, 0)
	err := r.list(timelineBucket, func() any { return &domain.TimelineEvent{} }, func(item any) {
		value := *(item.(*domain.TimelineEvent))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	if err != nil {
		return nil, err
	}
	return domain.SortTimeline(items), nil
}

func (r *Repository) DeleteTimelineEvent(id string) error { return r.delete(timelineBucket, id) }

func (r *Repository) CountBucket(bucket []byte) (int, error) {
	count := 0
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return fmt.Errorf("bucket %q not found", string(bucket))
		}
		return b.ForEach(func(key, value []byte) error {
			if value != nil {
				count++
			}
			return nil
		})
	})
	return count, err
}

func sortStrings(values []string) []string {
	result := append([]string(nil), values...)
	sort.Strings(result)
	return result
}
