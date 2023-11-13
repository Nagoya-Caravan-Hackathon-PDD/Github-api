package gateways

type githubAPIGateway struct {
}

func NewGithubAPIGateway() *githubAPIGateway {
	return &githubAPIGateway{}
}

func (g *githubAPIGateway) GithubAuthentication() error {
	return nil
}
