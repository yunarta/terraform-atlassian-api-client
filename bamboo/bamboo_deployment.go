package bamboo

// CreateDeployment struct defines the blueprint for creating a new deployment.
// Name: the deployment's human-readable display name.
// PlanKey: a unique ID (key) that identifies the associated build plan.
// Description: a brief overview of what the deployment is all about.
// PublicAccess: a flag indicating whether the deployment should be accessible to the public.
type CreateDeployment struct {
	Name         string `json:"name,omitempty"`
	PlanKey      Key    `json:"planKey,omitempty"`
	Description  string `json:"description,omitempty"`
	PublicAccess bool   `json:"publicAccess,omitempty"`
}

// UpdateDeployment struct shares similar fields with CreateDeployment,
// but it's for updating an existing deployment rather than creating a new one.
type UpdateDeployment struct {
	Name        string `json:"name,omitempty"`
	PlanKey     Key    `json:"planKey,omitempty"`
	Description string `json:"description,omitempty"`
}

// Key struct defines our key data type, which seems to represent some sort of unique ID.
type Key struct {
	Key string `json:"key"`
}

// DeploymentList struct represents a paginated response of deployments from an API.
// Start: the starting index of the list.
// MaxResult: the maximum number of results that can be returned.
// Results: slice of deployment data (DeploymentItem) returned by the search query.
type DeploymentList struct {
	Start     int              `json:"start-index,omitempty"`
	MaxResult int              `json:"max-result,omitempty"`
	Results   []DeploymentItem `json:"searchResults,omitempty"`
}

// DeploymentItem struct represents a single deployment item in a paginated list response.
// Id: the unique ID of this deployment.
// Type: the type of this entity (we don't know the specific types from this code).
// SearchEntity: detailed information about the deployment.
type DeploymentItem struct {
	Id           string           `json:"id,omitempty"`
	Type         string           `json:"type,omitempty"`
	SearchEntity DeploymentEntity `json:"searchEntity,omitempty"`
}

// DeploymentEntity struct is the search entity from DeploymentItem with some extra information.
// Id: the unique ID of this entity.
// Key: the unique key of this entity.
// ProjectName: the name of the project this deployment entity is attached to.
// Description: A brief description of this entity.
type DeploymentEntity struct {
	Id          string `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	Description string `json:"description,omitempty"`
}

// Deployment struct is an overall view of a single deployment with all necessary details.
// ID: the unique ID of this deployment.
// Name: the readable name of this deployment.
// PlanKey: the unique key that linked to the execution plan of this deployment.
// Description: Detail information about this deployment.
// Environments: a list of environment settings associated with this deployment.
// RepositorySpecsManaged: a flag that shows whether repository settings are managed automatically.
type Deployment struct {
	ID                     int           `json:"id,omitempty"`
	Name                   string        `json:"name,omitempty"`
	PlanKey                Key           `json:"planKey,omitempty"`
	Description            string        `json:"description,omitempty"`
	Environments           []Environment `json:"environments,omitempty"`
	RepositorySpecsManaged bool          `json:"repositorySpecsManaged,omitempty"`
}

// Environment struct represents each environment within a deployment setting.
// ID: the unique identifier of this environment within the deployment.
// Name: the Display name of this environment.
// Description: More detail about what this environment is and why it exists.
type Environment struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
