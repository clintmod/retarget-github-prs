package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"os"
)

func main() {
	tp := github.BasicAuthTransport{
		Username: os.Getenv("GITHUB_USERNAME"),
		Password: os.Getenv("GITHUB_PASSWORD"),
	}
	owner := "clintmod"
	repo := "Oceans"
	oldBranch := "master"
	newBranch := "development"

	client := github.NewClient(tp.Client())

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
