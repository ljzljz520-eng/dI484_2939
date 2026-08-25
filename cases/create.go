package cases

import (
	"fmt"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type Manager struct {
	repo *store.Repository
}

type CreateInput struct {
	ID          string
	Number      string
	Title       string
	Summary     string
	PublishDate string
	CreatedAt   string
}

func NewManager(repo *store.Repository) *Manager { return &Manager{repo: repo} }

func (m *Manager) CreateCase(input CreateInput) (domain.LegalCase, error) {
	value := domain.LegalCase{
		ID: input.ID, Number: strings.TrimSpace(input.Number), Title: strings.TrimSpace(input.Title),
		Summary: strings.TrimSpace(input.Summary), PublishDate: input.PublishDate,
		Status: domain.StatusDraft, CreatedAt: input.CreatedAt,
	}
	if err := value.Validate(); err != nil {
		return domain.LegalCase{}, err
	}
	if err := m.repo.PutCase(value); err != nil {
		return domain.LegalCase{}, fmt.Errorf("save case: %w", err)
	}
	return value, nil
}

func (m *Manager) SaveCase(value domain.LegalCase) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return m.repo.PutCase(value)
}

func (m *Manager) GetCase(id string) (domain.LegalCase, error) { return m.repo.GetCase(id) }

func (m *Manager) ArchiveCase(id string) error {
	value, err := m.repo.GetCase(id)
	if err != nil {
		return err
	}
	if value.Status == domain.StatusArchived {
		return nil
	}
	value.Status = domain.StatusArchived
	return m.repo.PutCase(value)
}
