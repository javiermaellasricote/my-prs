package response

import (
	"fmt"
	"strconv"

	"github.com/javiermaellasricote/my-prs/status"
)

var printF = fmt.Printf

var (
	colorReset = "\033[0m"
	green      = "\033[32m"
	yellow     = "\033[33m"
	red        = "\033[31m"
	blue       = "\033[36m"
	purple     = "\033[35m"
)

// Prints the CLI response to the console, separating between
// opened PRs (opened by the user running the command), and
// PRs for review (they are waiting for the user to add a review).
func PrintResponse(stss []status.RepoStatus) {
	oPRs := []status.PR{}
	rPRs := []status.PR{}
	for _, sts := range stss {
		oPRs = append(oPRs, sts.OpenedPRs...)
		rPRs = append(rPRs, sts.ReviewPRs...)
	}

	printF(string(red) + "\nOPENED PRs:\n")
	printF(string(colorReset))
	printPRs(oPRs)

	printF(string(red) + "\nWAITING FOR REVIEW:\n" + string(colorReset))
	printF(string(colorReset))
	printPRs(rPRs)
}

// Prints PR information with the correct formatting and colors
// to the console.
func printPRs(prs []status.PR) {
	for _, pr := range prs {
		ghLink := "https://github.com/" + pr.Repo + "/pull/" + strconv.Itoa(pr.ID)
		printF(string(yellow)+"  [%v]:", pr.Repo)
		printF(string(colorReset)+" %v\n\t%v", pr.Branch, pr.Status)
		printF(string(blue)+" %v\n\n", ghLink)
		printF(string(colorReset))
	}
}
