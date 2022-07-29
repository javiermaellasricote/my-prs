package status

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Converts a slice of strings into a slice of PR objects,
// extracting all the necessary information and filling the
// PR objects with the appropriate data.
func convertStrsToPRs(data []string) ([]PR, error) {
	prs := make([]PR, len(data)/4)
	for i, item := range data {
		idx := i / 4
		cleanItem := strings.Trim(item, " ")

		switch i % 4 {
		case 0:
			continue

		case 1:
			fmt.Println(cleanItem)
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

		case 2:
			prs[idx].Status = cleanItem

		case 3:
			continue

		default:
			err := errors.New("More items than expected")
			return []PR{}, err
		}
	}
	return prs, nil
}
