package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"os"
	"strings"
)

func envVarError(name string) {
	fmt.Errorf("No %v environment variable found", name)
	os.Exit(1)
}

func missingArg(arg string, index int) {
	fmt.Printf("Missing arg %v at position %d\n", arg, index)
	os.Exit(1)
}

func validateArgs(args []string) {
	if len(args) < 2 {
		missingArg("Github Account", 1)
	}
	if len(args) < 3 {
		missingArg("Old Branch", 2)
	}
	if len(args) < 4 {
		missingArg("New Branch", 3)
	}
	if len(args) < 5 {
		missingArg("Repos (e.g. oceans,triton,rhode", 3)
	}
}

func main() {
	uname := os.Getenv("GITHUB_USERNAME")
	pass := os.Getenv("GITHUB_PASSWORD")

	if uname == "" {
		envVarError("GITHUB_USERNAME")
	}
	if pass == "" {
		envVarError("GITHUB_PASSWORD")
	}

	validateArgs(os.Args)

	owner := os.Args[1]
	oldBranch := os.Args[2]
	newBranch := os.Args[3]
	repos := strings.Split(os.Args[4], ",")

	tp := github.BasicAuthTransport{Username: uname, Password: pass}

	client := github.NewClient(tp.Client())

	for _, repo := range repos {
		opt := &github.PullRequestListOptions{"open", "", oldBranch, "created", "desc", github.ListOptions{Page: 1}}
		pulls, _, err := client.PullRequests.List(context.Background(), owner, repo, opt)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		numberOfPulls := len(pulls)
		fmt.Println("number of pulls = ", numberOfPulls)
		for i := 0; i < numberOfPulls; i++ {
			pull := pulls[i]
			pullNumber := *pull.Number
			*pull.Base.Ref = newBranch
			fmt.Printf("Retargeting pull request %v the %v branch\n", pullNumber, *pull.Base.Ref)
			_, _, err := client.PullRequests.Edit(context.Background(), owner, repo, pullNumber, pull)
			if err != nil {
				fmt.Errorf("%d: PullRequests.Edit returned error: %v", i, err)
			} else {
				fmt.Printf("pull request %v retargeted\n", pullNumber)
			}

		}
	}
}
