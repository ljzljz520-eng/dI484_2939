package cli

import (
	"flag"
	"fmt"
	"io"
	"strings"

	"example.com/casescript/cases"
	"example.com/casescript/publish"
	"example.com/casescript/store"
)

type Runner struct {
	Repo    *store.Repository
	Cases   *cases.Manager
	Publish *publish.Service
}

func New(repo *store.Repository) *Runner {
	return &Runner{Repo: repo, Cases: cases.NewManager(repo), Publish: publish.NewService(repo)}
}

func (r *Runner) Run(args []string, out io.Writer) error {
	if len(args) == 0 {
		return writeUsage(out)
	}
	switch args[0] {
	case "list":
		return r.list(args[1:], out)
	case "export":
		return r.export(args[1:], out)
	case "archive":
		return r.archive(args[1:], out)
	case "summary":
		return r.summary(args[1:], out)
	case "student":
		return r.student(args[1:], out)
	case "search":
		return r.search(args[1:], out)
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func (r *Runner) list(args []string, out io.Writer) error {
	flags := flag.NewFlagSet("list", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	date := flags.String("date", "", "publish date")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*date) == "" {
		return fmt.Errorf("--date is required")
	}
	items, err := r.Cases.ListPublishedByDate(*date)
	if err != nil {
		return err
	}
	for _, value := range items {
		fmt.Fprintf(out, "%s %s\n", value.Number, value.Title)
	}
	return nil
}

func (r *Runner) export(args []string, out io.Writer) error {
	flags := flag.NewFlagSet("export", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	caseID := flags.String("case", "", "case id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if strings.TrimSpace(*caseID) == "" {
		return fmt.Errorf("--case is required")
	}
	handout, err := r.Publish.ExportHandout(*caseID)
	if err != nil {
		return err
	}
	_, err = io.WriteString(out, handout+"\n")
	return err
}

func (r *Runner) archive(args []string, out io.Writer) error {
	flags := flag.NewFlagSet("archive", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	caseID := flags.String("case", "", "case id")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if err := r.Cases.ArchiveCase(*caseID); err != nil {
		return err
	}
	_, err := fmt.Fprintf(out, "archived %s\n", *caseID)
	return err
}
