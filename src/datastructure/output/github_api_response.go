package output

type GithubAPIResponse struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	Contribution int    `json:"contribution"`
	Stars        int    `json:"stars"`
	Followers    int    `json:"followers"`
}
