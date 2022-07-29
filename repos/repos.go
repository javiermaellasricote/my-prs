package repos

import (
	"bytes"
	"os/exec"
	"strings"
)

// Retrieves a slice of strings with all the GitHub repo
// names belonging to an owner. The owner can be an individual
// or a project.
func GetRepos(owner string) ([]string, error) {
	stdout, err := ghSearchRepos(owner)
	if err != nil {
		return []string{}, err
	}
	infos, err := strings.Split(stdout, "\n"), nil
	if err != nil {
		return []string{}, err
	}

	names := []string{}
	for _, info := range infos {
		name := strings.Split(info, "\t")[0]
		names = append(names, name)
	}
	return names, nil
}

// Calls the github cli to retrieve all the repos belonging
// to a specific owner (it can be an individual or a project).
// Returns the standard output from the command and an error
// if the command could not be run successfully.
func ghSearchRepos(owner string) (string, error) {
	cmd := exec.Command("gh", "search", "repos", "--owner", owner)
	stdout := bytes.Buffer{}
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}
