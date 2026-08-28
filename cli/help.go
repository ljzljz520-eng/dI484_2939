package cli

import "io"

func writeUsage(out io.Writer) error {
	_, err := io.WriteString(out, "commands: list, search, summary, student, export, archive\n")
	return err
}
