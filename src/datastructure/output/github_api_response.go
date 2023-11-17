package output

type GithubAPIResponse struct {
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Contribution int    `json:"contribution"`
	Stars        int    `json:"stars"`
	Followers    int    `json:"followers"`
}
