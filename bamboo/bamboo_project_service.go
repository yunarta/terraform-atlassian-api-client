package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

const (
	projectsEndPoint         = "/rest/api/latest/project"
	projectEndpoint          = "/rest/api/latest/project/%s"
	planEndPoint             = "/rest/api/latest/plan/%s"
	specRepositoriesEndPoint = "/rest/api/latest/project/%s/repository"
	specRepositoryEndPoint   = "/rest/api/latest/project/%s/repository/%d"
)

// ProjectService is a key struct in this package.
// This struct holds a client that sends HTTP requests and receives HTTP responses.
// We will use instances of ProjectService to perform operations on Bamboo servers.
type ProjectService struct {
	// transport member is a client responsible for sending HTTP requests and receiving HTTP responses.
	transport transport.PayloadTransport
}

// The Create function allows us to create a new project in Bamboo.
// It receives project data as input and returns the newly created project if everything goes okay.
func (service *ProjectService) Create(request CreateProject) (*Project, error) {
	// We send a POST request to create a new project. A 201 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method:  http.MethodPost,
		Url:     projectsEndPoint,
		Payload: transport.JsonPayloadData{Payload: request},
	}, 201)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return nil, err
	}

	response := Project{}
	// We parse the return data into a Create struct.
	err = reply.Object(&response)
	if err != nil {
		// If parsing fails, we return the error.
		return nil, err
	}

	// If everything works as it should, we return the newly created project.
	return &response, nil
}

// The Read function fetches the details of a project given its key.
func (service *ProjectService) Read(projectKey string) (*Project, error) {
	// We send a GET request to fetch the project. A 200 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(projectEndpoint, projectKey),
	}, 200)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return nil, err
	}

	response := Project{}
	// We parse the returned data into a Project struct.
	err = reply.Object(&response)
	if err != nil {
		// If parsing fails, we return the error.
		return nil, err
	}

	// If everything works as it should, we return the fetched project.
	return &response, nil
}

// ProjectService struct should already be defined somewhere. It handles the operations related to the Project.

// Update is a method on the ProjectService struct. It sends a PUT request to update the project.
// projectKey is a unique identifier for a project and update is the data that needs to be updated.
// The function will return the updated project or any error occurred during the process.
func (service *ProjectService) Update(projectKey string, update UpdateProject) (*Project, error) {
	// The SendWithExpectedStatus method of the transport field on the service struct is called.
	// This method is used to send requests with a specific expected HTTP status code (200 in this case, which means Success).
	// The method, URL, and payload for the request are provided.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(projectEndpoint, projectKey),
		Payload: transport.JsonPayloadData{
			Payload: update,
		},
	}, 200)

	// If there's a communication error (network issues, server not responding, etc.), it gets returned immediately.
	if err != nil {
		return nil, err
	}

	// Create an empty Project struct to hold the response.
	response := Project{}

	// The Object method is called on reply to parse the response body into the Project struct.
	err = reply.Object(&response)
	if err != nil {
		// If there's a parsing error (e.g., the server responded with unexpected data), it gets returned immediately.
		return nil, err
	}

	// Return the updated project along with nil as the error, indicating success.
	return &response, nil
}

// Delete is another method on ProjectService. It sends a DELETE request to delete the project specified by the provided projectKey.
// It returns error (if any) that occurred during the deletion.
func (service *ProjectService) Delete(projectKey string) error {
	// The SendWithExpectedStatus function on the service's transport struct sends a DELETE request. It expects a 204 status code response.
	// 204 indicates that the server successfully processed the request and there's no additional content to send in the response payload.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(projectEndpoint, projectKey),
	}, 204)

	// Return the error (if any). If the function call was successful the error will be nil.
	return err
}

// The ReadPlan function fetches the details of a plan given its key.
// A plan in Bamboo defines a single build, including source code repository, optional triggers, commands to execute, and test results to collect.
func (service *ProjectService) ReadPlan(planKey string) (*Plan, error) {
	// We send a GET request to fetch the plan. A 200 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(planEndPoint, planKey),
	}, 200)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return nil, err
	}

	response := Plan{}
	// Parse the return data into a Plan struct.
	err = reply.Object(&response)
	if err != nil {
		// If parsing fails, we return the error.
		return nil, err
	}

	// If everything works as it should, we return the fetched plan.
	return &response, nil
}

// The GetSpecRepositories function fetches the repositories of a given project.
// It sends a GET request to fetch the repositories and returns them if everything works out.
func (service *ProjectService) GetSpecRepositories(projectKey string) ([]Repository, error) {
	// We send a GET request to fetch the repositories for a project. A 200 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(specRepositoriesEndPoint, projectKey),
	}, 200)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return nil, err
	}

	var repositories []Repository
	// We parse the return data into a slice of Repository struct.
	err = reply.Object(&repositories)
	if err != nil {
		// If parsing fails, we return the error.
		return nil, err
	}

	// If everything works as it should, we return the fetched repositories.
	return repositories, nil
}

// The AddSpecRepositories function adds a new repository to a given project.
// It sends a POST request to add a new repository to a project and returns the repository if addition is successful.
func (service *ProjectService) AddSpecRepositories(projectKey string, repositoryId int) (*Repository, error) {
	// We send a POST request to add a new repository to a project. A 201 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf(specRepositoriesEndPoint, projectKey),
		Payload: transport.JsonPayloadData{
			Payload: map[string]int{
				"id": repositoryId,
			},
		},
	}, 201)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return nil, err
	}

	var repository Repository
	// We parse the return data into a Repository struct.
	err = reply.Object(&repository)
	if err != nil {
		// If parsing fails, we return the error.
		return nil, err
	}

	// If everything works as it should, we return the added repository.
	return &repository, nil
}

// The RemoveSpecRepositories function removes a repository from a given project.
// It sends a DELETE request to remove a repository from a project and returns an error if the operation is unsuccessful.
func (service *ProjectService) RemoveSpecRepositories(projectKey string, repositoryId int) error {
	// We send a DELETE request to remove a repository from a project. A 204 status code signifies success.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(specRepositoryEndPoint, projectKey, repositoryId),
	}, 204)
	// If there's a communication error, we return it immediately.
	return err
}
