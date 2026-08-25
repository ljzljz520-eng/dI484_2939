package cases

import (
	"testing"

	"example.com/casescript/domain"
)

func TestApplyFilterAndMetrics(t *testing.T) {
	items := []domain.LegalCase{{ID: "a", Number: "2024-002", Title: "证据", Summary: "公开", PublishDate: "2024-06-01", Status: domain.StatusPublished}, {ID: "b", Number: "2024-001", Title: "合同", Summary: "草稿", PublishDate: "2024-06-01", Status: domain.StatusDraft}}
	filtered := ApplyFilter(items, CaseFilter{Date: "2024-06-01", Status: domain.StatusPublished})
	if len(filtered) != 1 || filtered[0].Number != "2024-002" {
		t.Fatalf("filtered=%#v", filtered)
	}
	metrics := BuildDateMetrics(items)
	if len(metrics) != 1 || metrics[0].Published != 1 || metrics[0].Drafts != 1 {
		t.Fatalf("metrics=%#v", metrics)
	}
}
