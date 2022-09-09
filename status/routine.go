package status

import (
	"bytes"
	"os/exec"
	"strings"
	"sync"
)

var newCommand = exec.Command

var (
	yourPRsMsg     = "Created by you"
	codeReviewMsg  = "\nRequesting a code review from you\n"
	noReviewPRsMsg = "You have no pull requests to review"
	noOpenedPRsMsg = "You have no open pull requests"
)

type statusChan struct {
	openedPRs []PR
	reviewPRs []PR
	err       error
}

// Retrieves the status for all the PRs related to the user making the request
// for the specified repo. Returns error if the statuses cannot be retrieved
// or parsed successfully.
func getStatusRoutine(repo string, stsChan chan statusChan, wg *sync.WaitGroup) {
	defer wg.Done()

	stdout, err := ghPRStatus(repo)
	if err != nil {
		stsChan <- statusChan{err: err}
	}

	info := strings.Split(stdout, yourPRsMsg)[1]
	infos := strings.Split(info, codeReviewMsg)

	oPRs, err := extractPRs(infos[0], noOpenedPRsMsg, repo)
	if err != nil {
		stsChan <- statusChan{err: err}
	}

	rPRs, err := extractPRs(infos[1], noReviewPRsMsg, repo)
	if err != nil {
		stsChan <- statusChan{err: err}
	}

	stsChan <- statusChan{
		openedPRs: oPRs,
		reviewPRs: rPRs,
	}
}

// Calls the github cli to retrieve the statuses in a repo
// for all the PRs that belong to the current user
// or that are pending for a review from the current user.
// Returns the standard output from the command and an error
// if the command could not be run successfully.
func ghPRStatus(repo string) (string, error) {
	cmd := newCommand("gh", "pr", "--repo", repo, "status")
	stdout := bytes.Buffer{}
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}
