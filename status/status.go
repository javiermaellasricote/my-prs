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
	OpenedPRs []PR  `json:"opened_prs"`
	ReviewPRs []PR  `json:"review_prs"`
	Err       error `json:"-"`
}

// Gets the status for each repo name passed as an argument.
// Returns error if the status of any of the repos could not
// be retrieved.
func GetStatus(rps []string) ([]RepoStatus, error) {
	stsChan := make(chan RepoStatus)
	go executeRoutines(rps, stsChan)
	return retrieveResponse(stsChan)
}

// Executes the go routines that get the status for the different
// repos. This function should be executed as a go routine so it
// doesn't block the main thread.
func executeRoutines(rps []string, stsChan chan RepoStatus) {
	defer close(stsChan)
	wg := sync.WaitGroup{}
	for _, rp := range rps {
		go getStatusRoutine(rp, stsChan, &wg)
	}
	wg.Wait()
}

// Retrieves the response from the RepoStatus channel and returns the
// status obtained as a RepoStatus slice. It returns an error if the
// channel response contains one.
func retrieveResponse(stsChan chan RepoStatus) ([]RepoStatus, error) {
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
