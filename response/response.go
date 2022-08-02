package response

import (
	"fmt"
	"strconv"

	"github.com/javiermaellasricote/my-prs/status"
)

var (
	colorReset = "\033[0m"
	green      = "\033[32m"
	yellow     = "\033[33m"
	red        = "\033[31m"
	blue       = "\033[36m"
	purple     = "\033[35m"
)

func PrintResponse(stss []status.RepoStatus) {
	oPRs := []status.PR{}
	rPRs := []status.PR{}
	for _, sts := range stss {
		oPRs = append(oPRs, sts.OpenedPRs...)
		rPRs = append(rPRs, sts.ReviewPRs...)
	}

	fmt.Printf(string(red) + "\nOPENED PRs:\n")
	fmt.Printf(string(colorReset))
	printPRs(oPRs)

	fmt.Printf(string(red) + "\nWAITING FOR REVIEW:\n" + string(colorReset))
	fmt.Printf(string(colorReset))
	printPRs(rPRs)
}

func printPRs(prs []status.PR) {
	for _, pr := range prs {
		ghLink := "https://github.com/" + pr.Repo + "/pull/" + strconv.Itoa(pr.ID)
		fmt.Printf(string(yellow)+"  [%v]:", pr.Repo)
		fmt.Printf(string(colorReset)+" %v\n\t%v", pr.Branch, pr.Status)
		fmt.Printf(string(blue)+" %v\n\n", ghLink)
		fmt.Printf(string(colorReset))
	}
}
