package main

import (
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

	stss := make([]status.RepoStatus, len(rps))
	for idx, rp := range rps {
		stss[idx], err = status.GetRepoStatus(rp)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
}
