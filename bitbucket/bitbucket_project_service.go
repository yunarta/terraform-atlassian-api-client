package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

const (
	projectsEndpoint = "/rest/api/latest/projects"
	projectEndpoint  = "/rest/api/latest/projects/%s"
)

// ProjectService struct represents a project service
type ProjectService struct {
	transport transport.PayloadTransport
}

// Create function creates a new project and returns the project or an error
func (service *ProjectService) Create(create CreateProject) (*Project, error) {
	// Sending a POST request to create a project
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    projectsEndpoint,
		Payload: &transport.JsonPayloadData{
			Payload: create,
		},
	}, 201)

	if err != nil {
		return nil, err
	}

	response := &Project{}
	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Read function reads an existing project and returns the project or an error
func (service *ProjectService) Read(key string) (*Project, error) {
	// Sending a GET request to read a project
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(projectEndpoint, key),
	}, 200)

	if err != nil {
		return nil, err
	}

	response := &Project{}
	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Update function updates an existing project and returns the updated project or an error
func (service *ProjectService) Update(key string, update ProjectUpdate) (*Project, error) {
	// Sending a GET request to update a project
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(projectEndpoint, key),
		Payload: transport.JsonPayloadData{
			Payload: update,
		},
	}, 200)

	if err != nil {
		return nil, err
	}

	response := &Project{}
	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Delete sends an HTTP DELETE request to delete an existing project. It takes in the project's key as an argument and returns any error encountered during the operation.
func (service *ProjectService) Delete(key string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(projectEndpoint, key),
	}, 204)

	// We just return the result of the delete command.
	return err
}
