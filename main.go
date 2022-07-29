package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/javiermaellasricote/my-prs/repos"
)

func main() {
	rps, err := repos.GetRepos("biblio-tech")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Print(len(rps))

}

// Calls the github cli to retrieve the statuses in a repo
// for all the PRs that belong to the current user
// or that are pending for a review from the current user.
// Returns the standard output from the command and an error
// if the command could not be run successfully.
func ghPRStatus(repo string) (string, error) {
	cmd := exec.Command("gh", "pr", "--repo", repo, "status")
	stdout := bytes.Buffer{}
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}
