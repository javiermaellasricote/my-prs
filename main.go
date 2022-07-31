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

	wg := sync.WaitGroup{}
	stsChan := make(chan status.RepoStatus)
	for _, rp := range rps {
		go status.GetRepoStatus(rp, stsChan, wg)
		wg.Wait()
		close(stsChan)
	}

	stss := []status.RepoStatus{}
	for sts := range stsChan {
		if sts.Err != nil {
			log.Fatalf(sts.Err.Error())
		}

		if len(sts.OpenedPRs) != 0 || len(sts.ReviewPRs) != 0 {
			stss = append(stss, sts)
		}
	}

	jsonStss, err := json.MarshalIndent(stss, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Print(string(jsonStss))
}
