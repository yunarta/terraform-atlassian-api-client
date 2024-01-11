package cloud

import (
	"github.com/yunarta/terraform-api-transport/transport"
	jira "github.com/yunarta/terraform-atlassian-api-client/jira/cloud"
)

type ConfluenceClient struct {
	spaceService            *SpaceService
	actorService            *jira.ActorService
	spacePermissionsService *SpacePermissionsService
	actorLookupService      *jira.ActorLookupService
}

func NewConfluenceClient(transport transport.PayloadTransport) *ConfluenceClient {
	spaceService := &SpaceService{transport: transport}
	actorService := jira.NewActorService(transport)
	return &ConfluenceClient{
		spaceService:            spaceService,
		actorService:            actorService,
		spacePermissionsService: &SpacePermissionsService{transport: transport},
		actorLookupService:      jira.NewActorLookupService(actorService),
	}
}

func (client *ConfluenceClient) SpaceService() *SpaceService {
	return client.spaceService
}

func (client *ConfluenceClient) ActorService() *jira.ActorService {
	return client.actorService
}

func (client *ConfluenceClient) ActorLookupService() *jira.ActorLookupService {
	return client.actorLookupService
}

func (client *ConfluenceClient) SpacePermissionsService() *SpacePermissionsService {
	return client.spacePermissionsService
}
