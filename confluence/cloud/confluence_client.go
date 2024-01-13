package cloud

// Importing required packages
import (
	"github.com/yunarta/terraform-api-transport/transport"
	jira "github.com/yunarta/terraform-atlassian-api-client/jira/cloud"
)

// ConfluenceClient is a client for accessing various Confluence services including SpaceService, ActorService,
// SpacePermissionsService, and ActorLookupService.
type ConfluenceClient struct {
	spaceService            *SpaceService
	actorService            *jira.ActorService
	spacePermissionsService *SpacePermissionsService
	actorLookupService      *jira.ActorLookupService
}

// NewConfluenceClient constructs a new instance of the ConfluenceClient. It takes a transport to handle payload
// transport and initializes all the services required to interact with the Confluence services.
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

// SpaceService returns the SpaceService instance of the ConfluenceClient for working with spaces in Confluence.
func (client *ConfluenceClient) SpaceService() *SpaceService {
	return client.spaceService
}

// ActorService returns the ActorService instance of the ConfluenceClient for managing actors (e.g., user accounts).
func (client *ConfluenceClient) ActorService() *jira.ActorService {
	return client.actorService
}

// ActorLookupService returns the ActorLookupService instance of the ConfluenceClient for retrieving
// actor information based on their identifiers.
func (client *ConfluenceClient) ActorLookupService() *jira.ActorLookupService {
	return client.actorLookupService
}

// SpacePermissionsService returns the SpacePermissionsService instance of the ConfluenceClient for managing
// space permissions.
func (client *ConfluenceClient) SpacePermissionsService() *SpacePermissionsService {
	return client.spacePermissionsService
}
