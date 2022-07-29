package status

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type PR struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func GetStatus(repo string) ([]PR, error) {
	stdout, err := ghPRStatus(repo)
	if err != nil {
		return []PR{}, err
	}

	data := strings.Split(stdout, "Created by you")[1]
	spltData := strings.Split(data, "\nRequesting a code review from you\n")

	return getOpenPRs(spltData[0])
}

func getOpenPRs(prInfo string) ([]PR, error) {
	noPRsMsg := "You have no open pull requests\n"
	if prInfo == noPRsMsg {
		return []PR{}, nil
	}
	data := strings.Split(prInfo, "\n")

	prs := make([]PR, len(data)/4)
	for i, item := range data {
		idx := i / 4
		switch i % 4 {
		case 0:
			continue
		case 1:
			//TODO: Get the name, description and id
		case 2:
			prs[idx].Status = item
		case 3:
			continue
		default:
			err := errors.New("More items than expected")
			return []PR{}, err
		}
	}
	return prs, nil
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
