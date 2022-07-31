package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/javiermaellasricote/my-prs/repos"
	"github.com/javiermaellasricote/my-prs/status"
)

func main() {
	rps, err := repos.GetRepos(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}

	stss := []status.RepoStatus{}
	for _, rp := range rps {
		sts, err := status.GetRepoStatus(rp)
		if err != nil {
			log.Fatalf(err.Error())
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
