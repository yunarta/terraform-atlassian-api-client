package bamboo

import "github.com/yunarta/terraform-api-transport/transport"

// Client represents the primary entry point for interacting with the Bamboo services.
// It contains service handlers for the project, deployment, repository and user services.
type Client struct {
	agentAssignmentService *AgentAssignmentService
	projectService         *ProjectService
	planService            *PlanService
	deploymentService      *DeploymentService
	repositoryService      *RepositoryService
	userService            *UserService
}

// NewBambooClient constructs and returns a new instance of the Client type.
// The function takes a transport payload that is used to facilitate the underlying
// communication with the Bamboo services. It initializes each of the service handlers
// with this transport payload.
func NewBambooClient(transport transport.PayloadTransport) *Client {
	return &Client{
		agentAssignmentService: &AgentAssignmentService{transport: transport},
		projectService:         &ProjectService{transport: transport},
		planService:            &PlanService{transport: transport},
		deploymentService:      &DeploymentService{transport: transport},
		repositoryService:      &RepositoryService{transport: transport},
		userService:            &UserService{transport: transport},
	}
}

func (client *Client) AgentAssignmentService() *AgentAssignmentService {
	return client.agentAssignmentService
}

// ProjectService is a getter method that returns the project service handler
// from the Bamboo client. It allows accessibility to the project services like
// creating, updating, or deleting a project.
func (client *Client) ProjectService() *ProjectService {
	return client.projectService
}

func (client *Client) PlanService() *PlanService {
	return client.planService
}

// DeploymentService is a getter method that returns the deployment service handler
// from the Bamboo client. It provides access to deployment services like initiating
// a deployment, fetching deployment info and more.
func (client *Client) DeploymentService() *DeploymentService {
	return client.deploymentService
}

// RepositoryService is a getter method that returns the repository service handler
// from the Bamboo client. It lets you interact with repository services such as
// creating, updating, or deleting a repository.
func (client *Client) RepositoryService() *RepositoryService {
	return client.repositoryService
}

// UserService is a getter method that returns the user service handler
// from the Bamboo client. It opens up access to user related services like
// creating a user, updating user info, or deleting a user.
func (client *Client) UserService() *UserService {
	return client.userService
}
