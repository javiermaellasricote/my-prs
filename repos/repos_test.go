package repos

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestGetRepos(t *testing.T) {
	cases := []struct {
		name   string
		newCmd func(string, ...string) *exec.Cmd
		hasErr bool
		res    []string
	}{
		{
			name: "it should return empty result if cmd output empty",
			newCmd: func(name string, args ...string) *exec.Cmd {
				return exec.Command("echo")
			},
			hasErr: false,
			res:    []string{},
		},
		{
			name: "It should error if cmd execution fails",
			newCmd: func(name string, args ...string) *exec.Cmd {
				return exec.Command("exit 1")
			},
			hasErr: true,
			res:    []string{},
		},
		{
			name: "it should return one name when cmd return one non-empty line",
			newCmd: func(name string, args ...string) *exec.Cmd {
				return exec.Command("echo", "hello\tworld")
			},
			hasErr: false,
			res:    []string{"hello"},
		},
		{
			name: "it should return multiple names when cmd return multiple non-empty lines",
			newCmd: func(name string, args ...string) *exec.Cmd {
				return exec.Command("echo", "hello\tworld\n\none\ttwo\nthree\n")
			},
			hasErr: false,
			res:    []string{"hello", "one", "three"},
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			newCommand = tt.newCmd

			res, err := GetRepos("me", 10)
			if (err == nil) == tt.hasErr {
				fmt.Printf("\n%#v\n", err.Error())
				t.Errorf("Unexpected response for error")
			}
			if !reflect.DeepEqual(res, tt.res) {
				t.Errorf("Expected result %v, but received %v", tt.res, res)
			}
		})
	}
}
