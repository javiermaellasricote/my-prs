package response

import (
	"testing"

	"github.com/javiermaellasricote/my-prs/status"
)

func TestPrintResponse(t *testing.T) {
	cases := []struct {
		name       string
		stss       []status.RepoStatus
		printCount int
	}{
		{
			name:       "It should not print any PRs when RepoStatus slice empty",
			stss:       []status.RepoStatus{},
			printCount: 4,
		},
		{
			name: "It should print PRs that are opened",
			stss: []status.RepoStatus{
				{
					OpenedPRs: []status.PR{
						{
							ID:          1,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
					},
				},
			},
			printCount: 8,
		},
		{
			name: "It should print PRs that are in review",
			stss: []status.RepoStatus{
				{
					ReviewPRs: []status.PR{
						{
							ID:          2,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
					},
				},
			},
			printCount: 8,
		},
		{
			name: "It should print all PRs if there are more than one in the array",
			stss: []status.RepoStatus{
				{
					OpenedPRs: []status.PR{
						{
							ID:          3,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
						{
							ID:          1,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
					},
					ReviewPRs: []status.PR{
						{
							ID:          4,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
						{
							ID:          2,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
					},
				},
			},
			printCount: 20,
		},
		{
			name: "It should print PRs that are opened and in review",
			stss: []status.RepoStatus{
				{
					OpenedPRs: []status.PR{
						{
							ID:          1,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
					},
					ReviewPRs: []status.PR{
						{
							ID:          2,
							Branch:      "branch",
							Repo:        "repo",
							Description: "description",
							Status:      "status",
						},
					},
				},
			},
			printCount: 12,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			printF = func(format string, args ...any) (int, error) {
				count++
				return 0, nil
			}

			PrintResponse(tt.stss)
			if count != tt.printCount {
				t.Errorf("Expected %v print calls, got %v calls", count, tt.printCount)
			}
		})
	}
}
