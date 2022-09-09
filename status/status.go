package status

import (
	"sync"
)

type PR struct {
	ID          int    `json:"id"`
	Branch      string `json:"branch"`
	Repo        string `json:"repository"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type RepoStatus struct {
	OpenedPRs []PR `json:"opened_prs"`
	ReviewPRs []PR `json:"review_prs"`
}

// Gets the status for each repo name passed as an argument.
// Returns error if the status of any of the repos could not
// be retrieved.
func GetStatus(rps []string) ([]RepoStatus, error) {
	stsChan := make(chan statusChan)
	go executeRoutines(rps, stsChan)
	return retrieveResponse(stsChan)
}

// Executes the go routines that get the status for the different
// repos. This function should be executed as a go routine so it
// doesn't block the main thread.
func executeRoutines(rps []string, stsChan chan statusChan) {
	defer close(stsChan)
	wg := sync.WaitGroup{}
	for _, rp := range rps {
		wg.Add(1)
		go getStatusRoutine(rp, stsChan, &wg)
	}
	wg.Wait()
}

// Retrieves the response from the statusChan channel and returns the
// status obtained as a RepoStatus slice. It returns an error if the
// channel response contains one.
func retrieveResponse(stsChan chan statusChan) ([]RepoStatus, error) {
	stss := []RepoStatus{}
	for sts := range stsChan {
		if sts.err != nil {
			return []RepoStatus{}, sts.err
		}

		if len(sts.openedPRs) != 0 || len(sts.reviewPRs) != 0 {
			stss = append(stss, RepoStatus{
				OpenedPRs: sts.openedPRs,
				ReviewPRs: sts.reviewPRs,
			})
		}
	}
	return stss, nil
}
