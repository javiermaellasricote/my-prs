package status

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestGetStatus(t *testing.T) {
	cases := []struct {
		name   string
		rps    []string
		newCmd func(string, ...string) *exec.Cmd
		res    []RepoStatus
		hasErr bool
	}{
		{
			name: "it should return empty",
			rps:  []string{"pr1", "pr2", "pr3"},
			newCmd: func(name string, args ...string) *exec.Cmd {
				return exec.Command("echo")
			},
			res: []RepoStatus{},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			newCommand = tt.newCmd
			res, err := GetStatus(tt.rps)

			if (err == nil) == tt.hasErr {
				t.Errorf("Unexpected response for error: %v", err)
			}

			if !reflect.DeepEqual(res, tt.res) {
				t.Errorf("Expected result %v, but got %v", tt.res, res)
			}
		})
	}
}
