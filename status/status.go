package status

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type PR struct {
	ID          int    `json:"id"`
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

// Extracts the PRs that the user has opened from the
// information string. Returns error if the string passed
// does not match the expected format.
func getOpenPRs(prInfo string) ([]PR, error) {
	noPRsMsg := "\n  You have no open pull requests\n"
	if prInfo == noPRsMsg {
		return []PR{}, nil
	}
	data := strings.Split(prInfo, "\n")

	prs := make([]PR, len(data)/4)
	for i, item := range data {
		idx := i / 4
		cleanItem := strings.Trim(item, " ")

		switch i % 4 {
		case 0:
			continue
		case 1:
			fmt.Println(cleanItem)
			split1 := strings.Split(cleanItem, "#")[1]
			split2 := strings.Split(split1, "  ")
			split3 := strings.Split(split2[1], " [")
			id, err := strconv.Atoi(split2[0])
			if err != nil {
				return []PR{}, err
			}
			prs[idx].ID = id
			prs[idx].Description = split3[0]
			prs[idx].Name = strings.Trim(split3[1], "]")
		case 2:
			prs[idx].Status = cleanItem
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
