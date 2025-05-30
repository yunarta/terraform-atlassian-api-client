package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"strconv"
	"strings"
)

// Define constants representing the endpoints of our application's API.
// These will be used to send requests to execute specific operations.
const (
	deploymentSearchEndpoint    = "/rest/api/latest/search/deployments?searchTerm=%s"
	deploymentEndPoint          = "/rest/api/latest/deploy/project"
	deploymentWithIdEndPoint    = "/rest/api/latest/deploy/project/%d"
	deploymentRssEndPoint       = "/rest/api/latest/deploy/project/%d/repository"
	deploymentRssWithIdEndPoint = "/rest/api/latest/deploy/project/%d/repository/%d"
)

// DeploymentService contains our transport method (transport.PayloadTransport).
// DeploymentService is used to send requests related to the various operations on Deployment instances.
type DeploymentService struct {
	// Dependency for handling communication via Payload Transport (HTTP requests).
	transport transport.PayloadTransport
}

// The Create function allows creating a new Deployment instance using provided request data.
// Specifically, it sends an HTTP PUT request to the Deployment endpoint.
func (service *DeploymentService) Create(request CreateDeployment) (*Deployment, error) {
	// Send a payload to our endpoint.
	// Here we use the HTTP PUT method and expect a 200 OK status response from the server.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method:  http.MethodPut,
		Url:     deploymentEndPoint,
		Payload: transport.JsonPayloadData{Payload: request},
	}, 200)
	// Check for any errors. If there's an error, it's returned immediately.
	if err != nil {
		return nil, err
	}

	// Initialize an empty Deployment struct to hold our server response data.
	response := Deployment{}
	// Here, we transfer the data from the received payload into our Deployment struct.
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	// Successful operation - return the created Deployment.
	return &response, nil
}

// The Read function retrieves a Deployment instance's data using the provided Deployment name.
func (service *DeploymentService) Read(deploymentName string) (*Deployment, error) {
	// Send a GET request to the search endpoint and expect a 200 status response.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(deploymentSearchEndpoint, deploymentName),
	}, 200)

	// If there's an error, return it immediately.
	if err != nil {
		return nil, err
	}

	// Initialize a DeploymentList struct to hold our server response data.
	deploymentList := DeploymentList{}
	err = reply.Object(&deploymentList)
	if err != nil {
		return nil, err
	}

	// Iterate over the Deployments to find a matching Deployment name.
	for _, deployment := range deploymentList.Results {
		if strings.EqualFold(deployment.SearchEntity.ProjectName, deploymentName) {
			// If found, go ahead and get the full Deployment data using its ID.
			id, err := strconv.Atoi(deployment.Id)
			if err != nil {
				return nil, err
			}

			// Return the full Deployment data.
			return service.ReadWithId(id)
		}
	}

	// Return nil if no matching Deployment was found.
	return nil, nil
}

// The ReadWithId function retrieves a Deployment instance's data using its ID.
func (service *DeploymentService) ReadWithId(deploymentId int) (*Deployment, error) {
	// Send a GET request to the search endpoint with the Deployment ID in URL.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(deploymentWithIdEndPoint, deploymentId),
	}, 200)
	// If there's an error, return it immediately.
	if err != nil {
		return nil, err
	}

	// Initialize a Deployment struct to hold our server response data.
	deployment := Deployment{}
	// Transfer the response data to our Deployment struct.
	err = reply.Object(&deployment)
	if err != nil {
		return nil, err
	}

	// Return the found Deployment instance.
	return &deployment, nil
}

// The UpdateWithId function updates an existing Deployment instance's data using its id and provided request data.
func (service *DeploymentService) UpdateWithId(deploymentId int, request UpdateDeployment) (*Deployment, error) {
	// Send a POST request containing updated data to the endpoint with the Deployment ID in URL.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf(deploymentWithIdEndPoint, deploymentId),
		Payload: transport.JsonPayloadData{
			Payload: request,
		},
	}, 200)
	// If there's an error, return it immediately.
	if err != nil {
		return nil, err
	}

	// Initialize a Deployment struct to hold the updated Deployment data.
	deployment := Deployment{}
	// Transfer the received data into our Deployment struct.
	err = reply.Object(&deployment)
	if err != nil {
		return nil, err
	}

	// Return the updated Deployment instance.
	return &deployment, nil
}

// The Delete function deletes a Deployment instance using its id.
func (service *DeploymentService) Delete(deploymentId int) error {
	// Send a DELETE request to the server.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(deploymentWithIdEndPoint, deploymentId),
	}, 204)

	// Return any error that occurred during the request.
	return err
}

// The GetSpecRepositories function retrieves the spec repository for a given Deployment id.
func (service *DeploymentService) GetSpecRepositories(deploymentId int) ([]Repository, error) {
	// Send a GET request to the spec repository endpoint with Deployment ID in URL.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(deploymentRssEndPoint, deploymentId),
	}, 200)
	// If there's an error during sending request, it's returned immediately.
	if err != nil {
		return nil, err
	}

	// Initialize a slice of Repository structs to hold our server response data.
	var repositories []Repository
	// Transfer the received data to our repositories slice.
	err = reply.Object(&repositories)
	if err != nil {
		return nil, err
	}

	// Return the slice of found repositories.
	return repositories, nil
}

// The AddSpecRepositories function adds a spec Repository to a Deployment using provided Deployment and Repository IDs.
func (service *DeploymentService) AddSpecRepositories(deploymentId int, repositoryId int) (*Repository, error) {
	// Send a POST request to add a spec repository with the corresponding repository and deployment IDs in the payload.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf(deploymentRssEndPoint, deploymentId),
		Payload: transport.JsonPayloadData{
			Payload: map[string]int{
				"id": repositoryId,
			},
		},
	}, 201)
	// If there's an error, return it immediately.
	if err != nil {
		return nil, err
	}

	// Initialize a Repository struct to hold our server response data.
	var repository Repository
	// Transfer the response data to our Repository struct.
	err = reply.Object(&repository)
	if err != nil {
		return nil, err
	}

	// Return the added Repository.
	return &repository, nil
}

// The RemoveSpecRepositories function removes a spec Repository from the Deployment using the provided Deployment and Repository IDs.
func (service *DeploymentService) RemoveSpecRepositories(deploymentId int, repositoryId int) error {
	// Send a DELETE request with the corresponding repository and deployment IDs in the payload.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(deploymentRssWithIdEndPoint, deploymentId, repositoryId),
		Payload: transport.JsonPayloadData{
			Payload: map[string]int{
				"id": repositoryId,
			},
		},
	}, 204)

	// Return any error that occurred during the request.
	return err
}
