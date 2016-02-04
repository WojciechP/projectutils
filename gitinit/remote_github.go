package gitinit

import "github.com/google/go-github/github"

// GitHubRemote represents a GitHub repository
type GitHubRemote struct {
	client *github.Client
	repo   *github.Repository
	org    string
}

// URL returns the human-viewable url of the repo
func (gh *GitHubRemote) URL() *string {
	return gh.repo.HTMLURL
}

// GitHubRemoteBuilder creates remote GitHub repos
type GitHubRemoteBuilder struct {
	github.Client
}

// Build creates a GitHub repository specification
func (client *GitHubRemoteBuilder) Build(namespace, name, fullName, description string) (RemoteRepository, error) {
	repo := github.Repository{
		Name:        &name,
		Description: &description,
		FullName:    &fullName,
	}
	newrepo, _, err := client.Repositories.Create(namespace, &repo)
	if err != nil {
		return nil, err
	}
	return &GitHubRemote{
		client: &client.Client,
		repo:   newrepo,
		org:    namespace,
	}, nil
}
