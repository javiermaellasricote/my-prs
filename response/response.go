package response

import (
	"fmt"

	"github.com/javiermaellasricote/my-prs/status"
)

func PrintResponse(stss []status.RepoStatus) {
	oPRs := []status.PR{}
	rPRs := []status.PR{}
	for _, sts := range stss {
		oPRs = append(oPRs, sts.OpenedPRs...)
		rPRs = append(rPRs, sts.ReviewPRs...)
	}

	fmt.Printf("\n\nOPENED PRs:\n")
	printPRs(oPRs)

	fmt.Printf("\n\nPRs FOR YOU TO REVIEW:\n")
	printPRs(rPRs)
}

func printPRs(prs []status.PR) {
	for _, pr := range prs {
		fmt.Printf("PR: %v\tRepo: %v\n", pr.Name, pr.Repo)
	}
}
