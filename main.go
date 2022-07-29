package main

import (
	"fmt"
	"log"

	"github.com/javiermaellasricote/my-prs/status"
)

func main() {
	/*
		rps, err := repos.GetRepos("biblio-tech")
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Print(len(rps))
	*/
	rps, err := status.GetStatus("biblio-tech/content-metadata-dynamodb")
	//rps, err := status.GetStatus("biblio-tech/providers")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%#v", rps)
}
