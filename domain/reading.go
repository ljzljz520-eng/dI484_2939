package domain

import (
	"fmt"
	"sort"
	"strings"
)

type ReadingPacket struct {
	ID           string   `json:"id"`
	CaseID       string   `json:"case_id"`
	Heading      string   `json:"heading"`
	Objectives   []string `json:"objectives"`
	KeyTerms     []string `json:"key_terms"`
	Questions    []string `json:"questions"`
	Instructions string   `json:"instructions"`
}

type VocabularyEntry struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Example    string `json:"example"`
	ChapterID  string `json:"chapter_id"`
}

func (p ReadingPacket) Validate() error {
	if strings.TrimSpace(p.ID) == "" || strings.TrimSpace(p.CaseID) == "" {
		return fmt.Errorf("reading packet ids are required")
	}
	if strings.TrimSpace(p.Heading) == "" || strings.TrimSpace(p.Instructions) == "" {
		return fmt.Errorf("reading packet heading and instructions are required")
	}
	if len(p.Objectives) == 0 || len(p.Questions) == 0 {
		return fmt.Errorf("reading packet requires objectives and questions")
	}
	for _, objective := range p.Objectives {
		if strings.TrimSpace(objective) == "" {
			return fmt.Errorf("reading objective cannot be empty")
		}
	}
	for _, question := range p.Questions {
		if strings.TrimSpace(question) == "" {
			return fmt.Errorf("reading question cannot be empty")
		}
	}
	return nil
}

func (p ReadingPacket) QuestionCount() int { return len(p.Questions) }

func (p ReadingPacket) PrimaryObjective() string {
	if len(p.Objectives) == 0 {
		return ""
	}
	return p.Objectives[0]
}

func (p ReadingPacket) NormalizedTerms() []string {
	terms := make([]string, 0, len(p.KeyTerms))
	seen := make(map[string]bool)
	for _, term := range p.KeyTerms {
		term = strings.TrimSpace(term)
		if term != "" && !seen[term] {
			seen[term] = true
			terms = append(terms, term)
		}
	}
	sort.Strings(terms)
	return terms
}

func (v VocabularyEntry) Validate() error {
	if strings.TrimSpace(v.Term) == "" || strings.TrimSpace(v.Definition) == "" {
		return fmt.Errorf("vocabulary term and definition are required")
	}
	if strings.TrimSpace(v.ChapterID) == "" {
		return fmt.Errorf("vocabulary chapter is required")
	}
	return nil
}

func ExtractTerms(text string, candidates []string) []string {
	lower := strings.ToLower(text)
	result := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		candidate = strings.TrimSpace(candidate)
		if candidate != "" && strings.Contains(lower, strings.ToLower(candidate)) {
			result = append(result, candidate)
		}
	}
	sort.Strings(result)
	return result
}

func JoinObjectives(objectives []string) string {
	clean := make([]string, 0, len(objectives))
	for _, objective := range objectives {
		if value := strings.TrimSpace(objective); value != "" {
			clean = append(clean, value)
		}
	}
	return strings.Join(clean, "；")
}
