package cases

import (
	"fmt"
	"strings"

	"example.com/casescript/domain"
)

type IntakeChecklist struct {
	HasSourceMaterial bool
	HasCourtNumber    bool
	HasLearningGoal   bool
	HasPublishDate    bool
}

func (c IntakeChecklist) Complete() bool {
	return c.HasSourceMaterial && c.HasCourtNumber && c.HasLearningGoal && c.HasPublishDate
}

func (m *Manager) ValidateIntake(value domain.LegalCase, checklist IntakeChecklist) error {
	if err := value.Validate(); err != nil {
		return err
	}
	if err := domain.EnsureCaseNumber(value.Number); err != nil {
		return err
	}
	if !checklist.HasSourceMaterial {
		return fmt.Errorf("source material is required")
	}
	if !checklist.HasCourtNumber {
		return fmt.Errorf("court number is required")
	}
	if !checklist.HasLearningGoal {
		return fmt.Errorf("learning goal is required")
	}
	if !checklist.HasPublishDate {
		return fmt.Errorf("publish date is required")
	}
	return nil
}

func (m *Manager) CreateFromDraft(input CreateInput, checklist IntakeChecklist) (domain.LegalCase, error) {
	if strings.TrimSpace(input.ID) == "" {
		return domain.LegalCase{}, fmt.Errorf("case id is required")
	}
	value, err := m.CreateCase(input)
	if err != nil {
		return domain.LegalCase{}, err
	}
	if err := m.ValidateIntake(value, checklist); err != nil {
		return domain.LegalCase{}, err
	}
	return value, nil
}

func DraftLabel(value domain.LegalCase) string {
	return fmt.Sprintf("%s %s [%s]", value.Number, value.Title, domain.StatusLabel(value.Status))
}
