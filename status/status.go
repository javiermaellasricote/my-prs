package status

import (
	"strings"
	"sync"
)

var (
	yourPRsMsg     = "Created by you"
	codeReviewMsg  = "\nRequesting a code review from you\n"
	noReviewPRsMsg = "You have no pull requests to review"
	noOpenedPRsMsg = "You have no open pull requests"
)

type PR struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Repo        string `json:"repository"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type RepoStatus struct {
	OpenedPRs []PR  `json:"opened_prs"`
	ReviewPRs []PR  `json:"review_prs"`
	Err       error `json:"-"`
}

// Retrieves the status for all the PRs in the repo
// related to the user making the request. Returns
// error if the statuses cannot be retrieved or parsed
// successfully.
func GetRepoStatus(repo string, stsChan chan RepoStatus, wg *sync.WaitGroup) {
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
