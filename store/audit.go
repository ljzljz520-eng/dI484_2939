package store

import (
	"fmt"
	"sort"

	"example.com/casescript/domain"
)

type AuditEntry struct {
	ID      string `json:"id"`
	CaseID  string `json:"case_id"`
	Actor   string `json:"actor"`
	Action  string `json:"action"`
	Date    string `json:"date"`
	Details string `json:"details"`
}

func (e AuditEntry) Validate() error {
	if e.ID == "" || e.CaseID == "" || e.Actor == "" || e.Action == "" || e.Details == "" {
		return fmt.Errorf("audit fields are required")
	}
	if !domain.ValidDate(e.Date) {
		return fmt.Errorf("audit date must be YYYY-MM-DD")
	}
	return nil
}

func (r *Repository) PutAudit(value AuditEntry) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(auditBucket, value.ID, value)
}

func (r *Repository) ListAudit(caseID string) ([]AuditEntry, error) {
	items := make([]AuditEntry, 0)
	err := r.list(auditBucket, func() any { return &AuditEntry{} }, func(item any) {
		value := *(item.(*AuditEntry))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	if err != nil {
		return nil, err
	}
	sort.SliceStable(items, func(i, j int) bool {
		if items[i].Date != items[j].Date {
			return items[i].Date < items[j].Date
		}
		return items[i].ID < items[j].ID
	})
	return items, nil
}

func AuditActions(items []AuditEntry) map[string]int {
	result := make(map[string]int)
	for _, item := range items {
		result[item.Action]++
	}
	return result
}
