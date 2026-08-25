package content

import (
	"fmt"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type ReviewManager struct{ repo *store.Repository }

func NewReviewManager(repo *store.Repository) *ReviewManager { return &ReviewManager{repo: repo} }

func (m *ReviewManager) Request(caseID, reviewer, requestedAt string) (domain.ReviewRequest, error) {
	request := domain.ReviewRequest{ID: caseID + ":review:" + reviewer, CaseID: caseID, Reviewer: strings.TrimSpace(reviewer), RequestedAt: requestedAt}
	if err := request.Validate(); err != nil {
		return domain.ReviewRequest{}, err
	}
	if err := m.repo.PutReview(request); err != nil {
		return domain.ReviewRequest{}, err
	}
	return request, nil
}

func (m *ReviewManager) Decide(id, decision, feedback string) (domain.ReviewRequest, error) {
	request, err := m.repo.GetReview(id)
	if err != nil {
		return domain.ReviewRequest{}, err
	}
	decision = strings.TrimSpace(decision)
	if decision != "approved" && decision != "changes_requested" {
		return domain.ReviewRequest{}, fmt.Errorf("review decision must approve or request changes")
	}
	request.Decision = decision
	request.Feedback = strings.TrimSpace(feedback)
	if err := m.repo.PutReview(request); err != nil {
		return domain.ReviewRequest{}, err
	}
	return request, nil
}

func (m *ReviewManager) Pending(caseID string) ([]domain.ReviewRequest, error) {
	items, err := m.repo.ListReviews(caseID)
	if err != nil {
		return nil, err
	}
	result := make([]domain.ReviewRequest, 0, len(items))
	for _, item := range items {
		if !item.IsResolved() {
			result = append(result, item)
		}
	}
	return result, nil
}

func ReviewStatus(request domain.ReviewRequest) string {
	if request.Decision == "approved" {
		return "reviewed"
	}
	if request.Decision == "changes_requested" {
		return "changes requested"
	}
	return "pending"
}
