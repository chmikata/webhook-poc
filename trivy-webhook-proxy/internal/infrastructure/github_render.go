package infrastructure

type GithubRender struct {
	organization string
	repository   string
	token        string
}

func NewGithubRender(organization, repository, token string) *GithubRender {
	return &GithubRender{
		organization: organization,
		repository:   repository,
		token:        token,
	}
}
