package response

import (
	"fmt"
	"strconv"

	"github.com/javiermaellasricote/my-prs/status"
)

func PrintResponse(stss []status.RepoStatus) {
	oPRs := []status.PR{}
	rPRs := []status.PR{}
	for _, sts := range stss {
		oPRs = append(oPRs, sts.OpenedPRs...)
		rPRs = append(rPRs, sts.ReviewPRs...)
	}

	fmt.Printf("\nOPENED PRs:\n")
	printPRs(oPRs)

	fmt.Printf("\nWAITING FOR REVIEW:\n")
	printPRs(rPRs)
}

func printPRs(prs []status.PR) {
	for _, pr := range prs {
		ghLink := "https://github.com/" + pr.Repo + "/pull/" + strconv.Itoa(pr.ID)
		fmt.Printf("  [%v]: %v\n\t%v %v\n\n", pr.Repo, pr.Branch, pr.Status, ghLink)
	}
}
