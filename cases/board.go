package cases

import (
	"fmt"
	"sort"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type TeacherBoard struct {
	Items       []BoardItem
	Published   int
	Drafts      int
	NeedsReview int
}

type BoardItem struct {
	Case         domain.LegalCase
	ChapterCount int
	NoteCount    int
	ReviewState  string
	NextAction   string
}

func BuildTeacherBoard(repo *store.Repository) (TeacherBoard, error) {
	items, err := repo.ListCases()
	if err != nil {
		return TeacherBoard{}, err
	}
	board := TeacherBoard{Items: make([]BoardItem, 0, len(items))}
	for _, value := range items {
		chapters, chapterErr := repo.ListChapters(value.ID)
		if chapterErr != nil {
			return TeacherBoard{}, chapterErr
		}
		notes, noteErr := repo.ListNotes(value.ID)
		if noteErr != nil {
			return TeacherBoard{}, noteErr
		}
		reviews, reviewErr := repo.ListReviews(value.ID)
		if reviewErr != nil {
			return TeacherBoard{}, reviewErr
		}
		reviewState := reviewLabel(reviews)
		item := BoardItem{Case: value, ChapterCount: len(chapters), NoteCount: len(notes), ReviewState: reviewState}
		item.NextAction = NextAction(item)
		board.Items = append(board.Items, item)
		switch value.Status {
		case domain.StatusPublished:
			board.Published++
		case domain.StatusDraft:
			board.Drafts++
		}
		if reviewState == "pending" {
			board.NeedsReview++
		}
	}
	sort.Slice(board.Items, func(i, j int) bool {
		if board.Items[i].Case.PublishDate != board.Items[j].Case.PublishDate {
			return board.Items[i].Case.PublishDate < board.Items[j].Case.PublishDate
		}
		return domain.CaseNumberKey(board.Items[i].Case.Number) < domain.CaseNumberKey(board.Items[j].Case.Number)
	})
	return board, nil
}

func reviewLabel(reviews []domain.ReviewRequest) string {
	if len(reviews) == 0 {
		return "not requested"
	}
	for _, review := range reviews {
		if review.Decision == "approved" {
			return "approved"
		}
		if review.Decision == "changes_requested" {
			return "changes requested"
		}
	}
	return "pending"
}

func NextAction(item BoardItem) string {
	if item.Case.Status == domain.StatusArchived {
		return "hidden"
	}
	if item.ChapterCount == 0 {
		return "add chapter"
	}
	if item.ReviewState == "not requested" && item.Case.Status == domain.StatusDraft {
		return "request review"
	}
	if item.ReviewState == "changes requested" {
		return "revise draft"
	}
	if item.Case.Status == domain.StatusDraft {
		return "publish"
	}
	return "review classroom link"
}

func (b TeacherBoard) Find(term string) []BoardItem {
	term = strings.ToLower(strings.TrimSpace(term))
	if term == "" {
		return append([]BoardItem(nil), b.Items...)
	}
	result := make([]BoardItem, 0)
	for _, item := range b.Items {
		if strings.Contains(strings.ToLower(item.Case.Title), term) || strings.Contains(strings.ToLower(item.Case.Number), term) {
			result = append(result, item)
		}
	}
	return result
}

func (b TeacherBoard) Summary() string {
	return fmt.Sprintf("%d cases · %d published · %d drafts · %d review items", len(b.Items), b.Published, b.Drafts, b.NeedsReview)
}
