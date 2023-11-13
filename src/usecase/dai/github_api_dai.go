package dai

type GithubAPIDai interface {
	GetGithubStatus() (string, error)
}
