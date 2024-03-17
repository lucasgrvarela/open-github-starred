package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	username := os.Getenv("GITHUB_USERNAME")
	token := os.Getenv("GITHUB_TOKEN")

	defaultOpenPerPage := flag.Int("open-per-page", 2, "Number of repositories to open per page")
	flag.Parse()

	if *defaultOpenPerPage > 100 {
		fmt.Println("open-per-page cannot be greater than 100. Exiting...")
		os.Exit(1)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	page := 1

	for {
		opts := &github.ActivityListStarredOptions{ListOptions: github.ListOptions{Page: page, PerPage: *defaultOpenPerPage}}
		repos, _, err := client.Activity.ListStarred(ctx, username, opts)
		if err != nil {
			fmt.Printf("Error fetching starred repositories: %v\n", err)
			return
		}

		for _, repo := range repos {
			fmt.Println(*repo.Repository.Name)
			url := *repo.Repository.HTMLURL
			if err := openURL(url); err != nil {
				fmt.Printf("Error opening URL %s: %v\n", url, err)
			}
		}

		if len(repos) == 0 {
			fmt.Println("No more repositories found. Exiting...")
			break
		}

		page++

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Continue to open more %d repositories? (Press Enter or Y to confirm): ", opts.PerPage)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" || strings.EqualFold(text, "Y") {
			fmt.Println("Confirmation received. Proceeding...")
		} else {
			fmt.Println("Confirmation not received. Exiting...")
			break
		}
	}
}

func openURL(url string) error {
	cmd := exec.Command("xdg-open", url)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
