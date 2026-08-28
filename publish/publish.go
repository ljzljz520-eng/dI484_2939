package publish

import (
	"fmt"
	"strings"

	"example.com/casescript/cases"
	"example.com/casescript/content"
	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type Service struct {
	repo    *store.Repository
	caseMgr *cases.Manager
	content *content.Manager
}

func NewService(repo *store.Repository) *Service {
	return &Service{repo: repo, caseMgr: cases.NewManager(repo), content: content.NewManager(repo)}
}

func (s *Service) PublishCase(caseID, publishedAt string) (domain.Publication, error) {
	value, err := s.caseMgr.GetCase(caseID)
	if err != nil {
		return domain.Publication{}, err
	}
	if value.Status != domain.StatusDraft {
		return domain.Publication{}, fmt.Errorf("case %s is not a draft", caseID)
	}
	if err := s.content.ValidateDraft(caseID); err != nil {
		return domain.Publication{}, err
	}
	previous, previousErr := s.repo.GetPublication(caseID)
	version := 1
	if previousErr == nil {
		version = previous.Version + 1
	}
	publication := domain.Publication{CaseID: caseID, Status: domain.StatusPublished, PublishedAt: publishedAt, Version: version}
	if err := publication.Validate(); err != nil {
		return domain.Publication{}, err
	}
	value.Status = domain.StatusPublished
	if err := s.caseMgr.SaveCase(value); err != nil {
		return domain.Publication{}, err
	}
	if err := s.repo.PutPublication(publication); err != nil {
		return domain.Publication{}, err
	}
	return publication, nil
}

func (s *Service) ExportHandout(caseID string) (string, error) {
	value, err := s.caseMgr.GetCase(caseID)
	if err != nil {
		return "", err
	}
	if !value.IsPublished() {
		return "", fmt.Errorf("case %s is not published", caseID)
	}
	chapters, err := s.content.Chapters(caseID)
	if err != nil {
		return "", err
	}
	notes, err := s.content.Notes(caseID)
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	builder.WriteString(value.Number + " " + value.Title + "\n")
	builder.WriteString(value.Summary + "\n\n")
	builder.WriteString(content.RenderOutline(chapters))
	if noteText := content.RenderNotes(notes); noteText != "" {
		builder.WriteString("\n\n课堂备注\n")
		builder.WriteString(noteText)
	}
	return builder.String(), nil
}

func (s *Service) Publication(caseID string) (domain.Publication, error) {
	return s.repo.GetPublication(caseID)
}
