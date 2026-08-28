package publish

import (
	"fmt"

	"example.com/casescript/domain"
)

type ReviewGate struct {
	RequireReview bool
	AllowDraft    bool
}

func CheckReviewGate(value domain.LegalCase, reviews []domain.ReviewRequest, gate ReviewGate) error {
	if gate.AllowDraft && value.Status == domain.StatusDraft {
		return nil
	}
	if !gate.RequireReview {
		return nil
	}
	for _, review := range reviews {
		if review.CaseID == value.ID && review.Decision == "approved" {
			return nil
		}
	}
	return fmt.Errorf("case %s needs an approved review", value.ID)
}

func ReviewCount(reviews []domain.ReviewRequest, decision string) int {
	count := 0
	for _, review := range reviews {
		if decision == "" || review.Decision == decision {
			count++
		}
	}
	return count
}

func PublicationReady(value domain.LegalCase, chapters []domain.Chapter, reviews []domain.ReviewRequest) bool {
	if value.Status != domain.StatusDraft || len(chapters) == 0 {
		return false
	}
	return CheckReviewGate(value, reviews, ReviewGate{RequireReview: true}) == nil
}
