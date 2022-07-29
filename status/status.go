package status

import (
	"strings"
)

var (
	yourPRsMsg     = "Created by you"
	codeReviewMsg  = "\nRequesting a code review from you\n"
	noReviewPRsMsg = "You have no pull requests to review\n\n"
	noOpenedPRsMsg = "\n  You have no open pull requests\n"
)

type PR struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type RepoStatus struct {
	Name      string `json:"name"`
	OpenedPRs []PR   `json:"opened_prs"`
	ReviewPRs []PR   `json:"review_prs"`
}

// Retrieves the status for all the PRs in the repo
// related to the user making the request. Returns
// error if the statuses cannot be retrieved or parsed
// successfully.
func GetRepoStatus(repo string) (RepoStatus, error) {
	stdout, err := ghPRStatus(repo)
	if err != nil {
		return RepoStatus{}, err
	}

	info := strings.Split(stdout, yourPRsMsg)[1]
	infos := strings.Split(info, codeReviewMsg)

	oPRs, err := extractPRs(infos[0], noOpenedPRsMsg)
	if err != nil {
		return RepoStatus{}, err
	}

	rPRs, err := extractPRs(infos[1], noReviewPRsMsg)
	if err != nil {
		return RepoStatus{}, err
	}

	return RepoStatus{
		Name:      repo,
		OpenedPRs: oPRs,
		ReviewPRs: rPRs,
	}, nil
}
