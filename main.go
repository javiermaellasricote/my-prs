package main

import (
	"log"
	"os"

	"github.com/javiermaellasricote/my-prs/repos"
	"github.com/javiermaellasricote/my-prs/response"
	"github.com/javiermaellasricote/my-prs/status"
)

func main() {
	rps, err := repos.GetRepos(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}

	stss, err := status.GetStatus(rps)
	if err != nil {
		log.Fatalf(err.Error())
	}

	response.PrintResponse(stss)
}
