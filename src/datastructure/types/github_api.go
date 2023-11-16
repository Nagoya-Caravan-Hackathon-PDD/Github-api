package types

type GitHubStatusQuery struct {
	User struct {
		Name         string
		Login        string
		Url          string
		AvatarUrl    string
		Bio          string
		Repositories struct {
			Nodes []struct {
				Name      string
				Languages struct {
					Edges []struct {
						Size int
						Node struct {
							Color string
							Name  string
						}
					}
				} `graphql:"languages(first: 10, orderBy: {field: SIZE, direction: DESC)"`
			}
		} `graphql:"repositories(first: 100, ownerAffiliations: OWNER, isFork: false)"`
		ContributionsCollection struct {
			ContributionCalendar struct {
				TotalContributions int
			}
		}
		Followers struct {
			TotalCount int
		}
	} `graphql:"user(login: $githubID)"`
}
