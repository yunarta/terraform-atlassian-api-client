package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

// ProjectService is a key struct in this package.
// This struct holds a client that sends HTTP requests and receives HTTP responses.
// We will use instances of ProjectService to perform operations on Bamboo servers.
type ProjectService struct {
	// transport member is a client responsible for sending HTTP requests and receiving HTTP responses.
	transport transport.PayloadTransport
}

// The CreateProject function allows us to create a new project in Bamboo.
// It receives project data as input and returns the newly created project if everything goes okay.
func (service *ProjectService) CreateProject(request CreateProject) (*CreateProject, error) {
	// We send a POST request to create a new project. A 201 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method:  http.MethodPost,
		Url:     "/rest/api/latest/project",
		Payload: transport.JsonPayloadData{Payload: request},
	}, 201)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return nil, err
	}

	response := CreateProject{}
	// We parse the return data into a CreateProject struct.
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
		Url:    fmt.Sprintf("/rest/api/latest/project/%s", projectKey),
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

// The ReadPlan function fetches the details of a plan given its key.
// A plan in Bamboo defines a single build, including source code repository, optional triggers, commands to execute, and test results to collect.
func (service *ProjectService) ReadPlan(planKey string) (*Plan, error) {
	// We send a GET request to fetch the plan. A 200 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/plan/%s", planKey),
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
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/repository", projectKey),
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
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/repository", projectKey),
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
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/repository/%d", projectKey, repositoryId),
	}, 204)
	// If there's a communication error, we return it immediately.
	return err
}
