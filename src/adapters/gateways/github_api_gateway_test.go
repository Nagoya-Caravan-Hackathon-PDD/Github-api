package gateways

import (
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/driver/github_api"
)

func TestGetStatus(t *testing.T) {
	gateway := NewGithubAPIGateway(github_api.GithubAuthentication())

	testCases := []struct {
		name  string
		input string
	}{
		{
			name:  "not nil check",
			input: "murasame29",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := gateway.GetStatus(tc.input)
			if err != nil {
				t.Errorf("request error: %v", err)
			}

			if status == nil {
				t.Errorf("status is nil")
			}

			t.Log(status)
		})
	}

}
