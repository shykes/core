package main

import (
)

type Core struct{}

func (m *Core) LoadContainerFromID(identifier string) *Container {
	return dag.LoadContainerFromID(ContainerID(identifier))
}

func (m *Core) Container() *Container {
	return dag.Container()
}

func (m *Core) LoadGitRepositoryFromID(identifier string) *GitRepository {
	return dag.LoadGitRepositoryFromID(GitRepositoryID(identifier))
}

func (m *Core) Git(url string) *GitRepository {
	return dag.Git(url)
}

func (m *Core) LoadDirectoryFromID(identifier string) *Directory {
	return dag.LoadDirectoryFromID(DirectoryID(identifier))
}

func (m *Core) Directory() *Directory {
	return dag.Directory()
}
