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

type RepoStatus struct {
	Name      string `json:"name"`
	OpenedPRs []PR   `json:"opened_prs"`
	ReviewPRs []PR   `json:"review_prs"`
}

var (
	noReviewPRs = "You have no pull requests to review\n\n"
	noOpenedPRs = "\n  You have no open pull requests\n"
)

func GetStatus(repo string) (RepoStatus, error) {
	stdout, err := ghPRStatus(repo)
	if err != nil {
		return RepoStatus{}, err
	}

	info := strings.Split(stdout, "Created by you")[1]
	infos := strings.Split(info, "\nRequesting a code review from you\n")

	oPRs, err := extractPRs(infos[0], noOpenedPRs)
	if err != nil {
		return RepoStatus{}, err
	}

	rPRs, err := extractPRs(infos[1], noReviewPRs)
	if err != nil {
		return RepoStatus{}, err
	}

	return RepoStatus{
		Name:      repo,
		OpenedPRs: oPRs,
		ReviewPRs: rPRs,
	}, nil
}
