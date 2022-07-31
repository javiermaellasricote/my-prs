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
	fmt.Println("Obtaining repos...")
	rps, err := repos.GetRepos(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Repos obtained!")

	fmt.Println("Obtaining PRs...")
	stsChan := make(chan status.RepoStatus)
	go status.GetRepoStatus(rps, stsChan)

	stss, err := status.StoreRepoStatus(stsChan)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("PRs obtained!")

	jsonStss, err := json.MarshalIndent(stss, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Print(string(jsonStss))
}
