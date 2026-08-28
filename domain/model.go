package domain

import (
	"fmt"
	"strings"
)

type CaseStatus string

const (
	StatusDraft     CaseStatus = "draft"
	StatusPublished CaseStatus = "published"
	StatusArchived  CaseStatus = "archived"
)

type LegalCase struct {
	ID          string     `json:"id"`
	Number      string     `json:"number"`
	Title       string     `json:"title"`
	Summary     string     `json:"summary"`
	PublishDate string     `json:"publish_date"`
	Status      CaseStatus `json:"status"`
	CreatedAt   string     `json:"created_at"`
}

type Chapter struct {
	ID       string `json:"id"`
	CaseID   string `json:"case_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Position int    `json:"position"`
}

type ClassroomNote struct {
	ID        string `json:"id"`
	CaseID    string `json:"case_id"`
	ChapterID string `json:"chapter_id"`
	Text      string `json:"text"`
	Author    string `json:"author"`
	UpdatedAt string `json:"updated_at"`
}

type Publication struct {
	CaseID      string     `json:"case_id"`
	Status      CaseStatus `json:"status"`
	PublishedAt string     `json:"published_at"`
	Version     int        `json:"version"`
}

func (c LegalCase) Validate() error {
	if strings.TrimSpace(c.ID) == "" {
		return fmt.Errorf("case id is required")
	}
	if strings.TrimSpace(c.Number) == "" || strings.TrimSpace(c.Title) == "" {
		return fmt.Errorf("case number and title are required")
	}
	if strings.TrimSpace(c.Summary) == "" {
		return fmt.Errorf("case summary is required")
	}
	if !ValidDate(c.PublishDate) {
		return fmt.Errorf("publish date must be YYYY-MM-DD")
	}
	if c.Status != StatusDraft && c.Status != StatusPublished && c.Status != StatusArchived {
		return fmt.Errorf("unsupported case status %q", c.Status)
	}
	return nil
}

func (c LegalCase) IsPublished() bool { return c.Status == StatusPublished }

func (c Chapter) Validate() error {
	if strings.TrimSpace(c.ID) == "" || strings.TrimSpace(c.CaseID) == "" {
		return fmt.Errorf("chapter ids are required")
	}
	if strings.TrimSpace(c.Title) == "" || strings.TrimSpace(c.Body) == "" {
		return fmt.Errorf("chapter title and body are required")
	}
	if c.Position < 1 {
		return fmt.Errorf("chapter position must be positive")
	}
	return nil
}

func (n ClassroomNote) Validate() error {
	if strings.TrimSpace(n.ID) == "" || strings.TrimSpace(n.CaseID) == "" {
		return fmt.Errorf("note ids are required")
	}
	if strings.TrimSpace(n.Text) == "" || strings.TrimSpace(n.Author) == "" {
		return fmt.Errorf("note text and author are required")
	}
	if !ValidDate(n.UpdatedAt) {
		return fmt.Errorf("note updated date must be YYYY-MM-DD")
	}
	return nil
}

func (p Publication) Validate() error {
	if strings.TrimSpace(p.CaseID) == "" {
		return fmt.Errorf("publication case id is required")
	}
	if p.Status != StatusPublished {
		return fmt.Errorf("publication must be published")
	}
	if !ValidDate(p.PublishedAt) || p.Version < 1 {
		return fmt.Errorf("publication date and version are required")
	}
	return nil
}
