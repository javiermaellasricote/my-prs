package status

import (
	"strings"
)

type PR struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func GetStatus(repo string) (string, error) {
	stdout, err := ghPRStatus(repo)
	if err != nil {
		return "", err
	}

	data := strings.Split(stdout, "Created by you")[1]
	spltData := strings.Split(data, "\nRequesting a code review from you\n")

	return spltData[1], nil
}

// Extracts the PRs that the user has opened from the
// information string. Returns error if the string passed
// does not match the expected format.
func getOpenPRs(info string) ([]PR, error) {
	noPRsMsg := "\n  You have no open pull requests\n"
	if info == noPRsMsg {
		return []PR{}, nil
	}

	data := strings.Split(info, "\n")
	return convertStrsToPRs(data)
}
