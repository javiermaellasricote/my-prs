package main

import (
	"log"
	"os"
	"strconv"

	"github.com/javiermaellasricote/my-prs/repos"
	"github.com/javiermaellasricote/my-prs/response"
	"github.com/javiermaellasricote/my-prs/status"
)

func main() {
	rpLmt := 10
	if len(os.Args) >= 3 {
		lmt, err := strconv.Atoi(os.Args[2])
		if err == nil {
			rpLmt = lmt
		}
	}

	rps, err := repos.GetRepos(os.Args[1], rpLmt)
	if err != nil {
		log.Fatalf(err.Error())
	}

	stss, err := status.GetStatus(rps)
	if err != nil {
		log.Fatalf(err.Error())
	}

	response.PrintResponse(stss)
}
