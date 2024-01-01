package cloud

import (
	"github.com/yunarta/terraform-api-transport/transport"
)

type JiraClient struct {
	projectService     *ProjectService
	actorService       *ActorService
	projectRoleService *ProjectRoleService
	actorLookupService *ActorLookupService
}

func NewJiraClient(transport transport.PayloadTransport) *JiraClient {
	actorService := &ActorService{transport: transport}

	return &JiraClient{
		projectService:     &ProjectService{transport: transport},
		actorService:       actorService,
		projectRoleService: &ProjectRoleService{transport: transport},
		actorLookupService: NewActorLookupService(actorService),
	}
}

func (client *JiraClient) ProjectService() *ProjectService {
	return client.projectService
}

func (client *JiraClient) ActorService() *ActorService {
	return client.actorService
}

func (client *JiraClient) ProjectRoleService() *ProjectRoleService {
	return client.projectRoleService
}

func (client *JiraClient) ActorLookupService() *ActorLookupService {
	return client.actorLookupService
}
