package status

import (
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

func StoreRepoStatus(stsChan chan RepoStatus) ([]RepoStatus, error) {
	stss := []RepoStatus{}
	for sts := range stsChan {
		if sts.Err != nil {
			return []RepoStatus{}, sts.Err
		}

		if len(sts.OpenedPRs) != 0 || len(sts.ReviewPRs) != 0 {
			stss = append(stss, sts)
		}
	}
	return stss, nil
}

func GetRepoStatus(rps []string, stsChan chan RepoStatus) {
	wg := sync.WaitGroup{}
	for _, rp := range rps {
		go execGetStatus(rp, stsChan, &wg)
	}
	wg.Wait()
	close(stsChan)
}
