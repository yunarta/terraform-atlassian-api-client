package bitbucket

type CreateRepo struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Repository struct {
	ID          int     `json:"id,omitempty"`
	Slug        string  `json:"slug,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Project     Project `json:"project,omitempty"`
}

type RepositoryCommits struct {
	Commits []RepositoryCommit `json:"values,omitempty"`
}

type RepositoryCommit struct {
	Id string `json:"id,omitempty"`
}
