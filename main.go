package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/javiermaellasricote/my-prs/repos"
	"github.com/javiermaellasricote/my-prs/status"
)

func main() {
	rps, err := repos.GetRepos(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}

	stsChan := make(chan status.RepoStatus)
	go execGetRepoStatus(rps, stsChan)

	stss, err := storeRepoStatus(stsChan)
	if err != nil {
		log.Fatalf(err.Error())
	}

	jsonStss, err := json.MarshalIndent(stss, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Print(string(jsonStss))
}

func storeRepoStatus(stsChan chan status.RepoStatus) ([]status.RepoStatus, error) {
	stss := []status.RepoStatus{}
	for sts := range stsChan {
		if sts.Err != nil {
			return []status.RepoStatus{}, sts.Err
		}

		if len(sts.OpenedPRs) != 0 || len(sts.ReviewPRs) != 0 {
			stss = append(stss, sts)
		}
	}
	return stss, nil
}

func execGetRepoStatus(rps []string, stsChan chan status.RepoStatus) {
	wg := sync.WaitGroup{}
	for _, rp := range rps {
		go status.GetRepoStatus(rp, stsChan, &wg)
	}
	wg.Wait()
	close(stsChan)
}
