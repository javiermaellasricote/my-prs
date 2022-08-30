package repos

import (
	"reflect"
	"testing"
)

func TestGetRepos(t *testing.T) {
	cases := []struct {
		name      string
		owner     string
		repoLimit int
		err       error
		res       []string
	}{}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetRepos(tt.owner, tt.repoLimit)
			if err != tt.err {
				t.Errorf("Expected error %v, but received %v", tt.err, err)
			}
			if !reflect.DeepEqual(res, tt.res) {
				t.Errorf("Expected result %v, but received %v", tt.res, res)
			}
		})
	}
}
