package gateways

import (
	"testing"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/driver/github_api"
)

func TestGetStatus(t *testing.T) {
	config.LoadEnv()
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

			if status != nil {
				t.Errorf("status is nil")
			}

		})
	}
}

// この形でよさそう？　求めてる形になってる？　おｋ
// 書いた通りに帰ってきてるからいいと思う bioいらなかな?
// query: {
// User:{
// Name: Login:murasame29
// Url:https://github.com/murasame29
// AvatarUrl:https://avatars.githubusercontent.com/u/77433104?v=4
// Bio:初心者えんじにゃー(=^・^=)
// ContributionsCollection:{
// ContributionCalendar:{
// TotalContributions:1386
// }
// }
// Followers:{
// TotalCount:4
// }
// }
// }
