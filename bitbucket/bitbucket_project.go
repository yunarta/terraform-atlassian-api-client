package bitbucket

type Project struct {
	ID          int64  `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type CreateProject struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ProjectUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
