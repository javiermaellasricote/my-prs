package status

import (
	"bytes"
	"os/exec"
	"strings"
	"sync"
)

// Retrieves the status for all the PRs in the repo
// related to the user making the request. Returns
// error if the statuses cannot be retrieved or parsed
// successfully.
func getStatusRoutine(repo string, stsChan chan RepoStatus, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	stdout, err := ghPRStatus(repo)
	if err != nil {
		stsChan <- RepoStatus{Err: err}
	}

	info := strings.Split(stdout, yourPRsMsg)[1]
	infos := strings.Split(info, codeReviewMsg)

	oPRs, err := extractPRs(infos[0], noOpenedPRsMsg, repo)
	if err != nil {
		stsChan <- RepoStatus{Err: err}
	}

	rPRs, err := extractPRs(infos[1], noReviewPRsMsg, repo)
	if err != nil {
		stsChan <- RepoStatus{Err: err}
	}

	stsChan <- RepoStatus{
		OpenedPRs: oPRs,
		ReviewPRs: rPRs,
	}
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
