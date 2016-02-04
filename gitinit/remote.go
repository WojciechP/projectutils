package gitinit

// RemoteRepository represents a usually existing remote repo
type RemoteRepository interface {
	URL() *string
}

// RemoteRepositoryBuilder creates remote repositories
type RemoteRepositoryBuilder interface {
	Build(namespace, name, fullName, description string) (RemoteRepository, error)
}
