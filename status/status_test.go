package status

import (
	"reflect"
	"testing"
)

func TestGetStatus(t *testing.T) {
	cases := []struct {
		name   string
		rps    []string
		res    []RepoStatus
		hasErr bool
	}{}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetStatus(tt.rps)

			if (err == nil) == tt.hasErr {
				t.Errorf("Unexpected response for error")
			}

			if !reflect.DeepEqual(res, tt.res) {
				t.Errorf("Expected result %v, but got %v", tt.res, res)
			}
		})
	}
}
