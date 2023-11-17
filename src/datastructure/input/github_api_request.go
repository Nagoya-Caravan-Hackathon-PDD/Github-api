package input

type GithubAPIRequest struct {
	GithubID    string `json:"github_id"`
	GitmonName  string `json:"gitmon_name"`
	GitmonImage int    `json:"gitmon_image"`
}
