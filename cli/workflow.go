package cli

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"example.com/casescript/cases"
	"example.com/casescript/publish"
)

func (r *Runner) summary(args []string, out io.Writer) error {
	flags := flag.NewFlagSet("summary", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	caseID := flags.String("case", "", "case id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*caseID) == "" {
		return fmt.Errorf("--case is required")
	}
	workspace, err := cases.NewManager(r.Repo).LoadWorkspace(*caseID)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(out, cases.WorkspaceSummary(workspace).Label())
	return err
}

func (r *Runner) student(args []string, out io.Writer) error {
	flags := flag.NewFlagSet("student", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	caseID := flags.String("case", "", "case id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*caseID) == "" {
		return fmt.Errorf("--case is required")
	}
	view, err := publish.BuildStudentView(r.Repo, *caseID)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(out, "%s\n", view.Summary); err != nil {
		return err
	}
	for _, chapter := range view.Chapters {
		if _, err := fmt.Fprintf(out, "%d. %s\n", chapter.Position, chapter.Title); err != nil {
			return err
		}
	}
	return nil
}

func (r *Runner) search(args []string, out io.Writer) error {
	flags := flag.NewFlagSet("search", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	term := flags.String("text", "", "search text")
	date := flags.String("date", "", "publish date")
	if err := flags.Parse(args); err != nil {
		return err
	}
	items, err := r.Cases.ListAll()
	if err != nil {
		return err
	}
	filtered := cases.ApplyFilter(items, cases.CaseFilter{Text: *term, Date: *date})
	for _, value := range filtered {
		if _, err := fmt.Fprintf(out, "%s %s %s\n", value.PublishDate, value.Number, value.Title); err != nil {
			return err
		}
	}
	return nil
}
