package repos

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

// Retrieves a slice of strings with all the GitHub repo names
//belonging to an owner. The owner can be an individual or a project.
func GetRepos(owner string, repoLimit int) ([]string, error) {
	stdout, err := ghRepoList(owner, repoLimit)
	if err != nil {
		return []string{}, err
	}

	infos := strings.Split(stdout, "\n")

	names := []string{}
	for _, info := range infos {
		name := strings.Split(info, "\t")[0]
		if name != "" {
			names = append(names, name)
		}
	}
	return names, nil
}

// Calls the github cli to retrieve all the repos belonging
// to a specific owner (it can be an individual or a project).
// Returns the standard output from the command and an error
// if the command could not be run successfully.
func ghRepoList(owner string, repoLimit int) (string, error) {
	lmtStr := strconv.Itoa(repoLimit)
	cmd := exec.Command("gh", "repo", "list", "--limit", lmtStr, owner)
	stdout := bytes.Buffer{}
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}
