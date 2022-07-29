package status

import (
	"strconv"
	"strings"
)

// Extracts the PRs for the user from the information
// string. Returns error if the string passed does not
// match the expected format.
func extractPRs(info, noPRsMessage string) ([]PR, error) {
	if info == noPRsMessage {
		return []PR{}, nil
	}

	cleanInfo := strings.Trim(info, "\n")
	data := strings.Split(cleanInfo, "\n")
	return convertStrsToPRs(data)
}

// Converts a slice of strings into a slice of PR objects,
// extracting all the necessary information and filling the
// PR objects with the appropriate data.
func convertStrsToPRs(data []string) ([]PR, error) {
	prs := make([]PR, len(data)/2)
	for i, item := range data {
		idx := i / 2
		cleanItem := strings.Trim(item, " ")

		switch i % 2 {
		case 0:
			split1 := strings.Split(cleanItem, "#")[1]
			split2 := strings.Split(split1, "  ")
			split3 := strings.Split(split2[1], " [")
			id, err := strconv.Atoi(split2[0])
			if err != nil {
				return []PR{}, err
			}
			prs[idx].ID = id
			prs[idx].Description = split3[0]
			prs[idx].Name = strings.Trim(split3[1], "]")

		case 1:
			prs[idx].Status = cleanItem
		}
	}
	return prs, nil
}
