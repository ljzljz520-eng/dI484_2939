package publish

import (
	"fmt"

	"example.com/casescript/domain"
)

func CanTransition(from, to domain.CaseStatus) bool {
	if from == domain.StatusDraft && to == domain.StatusPublished {
		return true
	}
	if from == domain.StatusPublished && to == domain.StatusArchived {
		return true
	}
	return false
}

func Transition(value domain.LegalCase, target domain.CaseStatus) (domain.LegalCase, error) {
	if !CanTransition(value.Status, target) {
		return value, fmt.Errorf("cannot transition %s to %s", value.Status, target)
	}
	value.Status = target
	return value, nil
}

func IsVisibleToStudent(value domain.LegalCase) bool { return value.Status == domain.StatusPublished }
