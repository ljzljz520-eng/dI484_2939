package domain

import "testing"

func TestDomainValidationAndFormatting(t *testing.T) {
	value := LegalCase{ID: "case-1", Number: "2024-001", Title: "标题", Summary: "摘要", PublishDate: "2024-06-01", Status: StatusDraft}
	if err := value.Validate(); err != nil {
		t.Fatal(err)
	}
	if !ValidDate(value.PublishDate) || CaseNumberKey("2024/001") != "2024:001" {
		t.Fatal("date or number normalization failed")
	}
	if NormalizeText("  一  二 ") != "一 二" || StatusLabel(StatusPublished) != "已发布" {
		t.Fatal("formatting failed")
	}
	if y, m, d, ok := DateParts("2024-06-01"); !ok || y != 2024 || m != 6 || d != 1 {
		t.Fatalf("date parts %d-%d-%d %v", y, m, d, ok)
	}
}
