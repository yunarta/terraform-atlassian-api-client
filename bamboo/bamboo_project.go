package bamboo

// CreateProject is a structure that holds the necessary information need to create a new Bamboo project.
// This includes properties like Name, Key, Description, and PublicAccess.
type CreateProject struct {
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	Description  string `json:"description,omitempty"`
	PublicAccess bool   `json:"publicAccess,omitempty"`
}

// Project is a structure representing a Bamboo project as returned by the API.
// This includes properties like Key, Name, and Description.
type Project struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Plan is a structure that represents a Bamboo build plan -- a collection of related build jobs.
// This includes properties like ProjectKey, ProjectName, Description, Key, ShortName, and Name.
type Plan struct {
	ProjectKey  string `json:"projectKey,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key,omitempty"`
	ShortName   string `json:"shortName,omitempty"`
	Name        string `json:"name,omitempty"`
}
