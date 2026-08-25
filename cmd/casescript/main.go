package main

import (
	"fmt"
	"os"

	"example.com/casescript/cli"
	"example.com/casescript/store"
)

func main() {
	path := os.Getenv("CASESCRIPT_DB")
	if path == "" {
		path = "casescript.db"
	}
	repo, err := store.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer repo.Close()
	if err := cli.New(repo).Run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
