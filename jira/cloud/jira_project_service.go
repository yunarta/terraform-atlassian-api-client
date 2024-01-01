package cloud

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
	"net/http"
)

type ProjectService struct {
	transport transport.PayloadTransport
}

func (service *ProjectService) Create(request jira.CreateProject) (*jira.Project, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/rest/api/latest/project",
		Payload: transport.JsonPayloadData{
			Payload: request,
		},
	}, 201)
	if err != nil {
		return nil, err
	}

	project := jira.CreateProjectResponse{}
	err = reply.Object(&project)
	if err != nil {
		return nil, err
	}

	return service.Read(project.Key)
}

type cloneProjectRequest struct {
	Key               string `json:"key,omitempty"`
	Name              string `json:"name,omitempty"`
	ExistingProjectId string `json:"existingProjectId,omitempty"`
}

func (service *ProjectService) Clone(request jira.CloneProject) (*jira.Project, error) {
	project, err := service.Read(request.ExistingProjectKey)
	if err != nil {
		return nil, err
	}

	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/rest/simplified/latest/project/shared",
		Payload: transport.JsonPayloadData{
			Payload: cloneProjectRequest{
				Key:               request.Key,
				Name:              request.Name,
				ExistingProjectId: project.ID,
			},
		},
	}, 200)
	if err != nil {
		return nil, err
	}

	response := jira.CloneProjectResponse{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return service.Read(response.ProjectKey)
}

func (service *ProjectService) Update(projectIdOrKey string, project jira.UpdateProject) (*jira.Project, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Payload: transport.JsonPayloadData{
			Payload: project,
		},
		Url: fmt.Sprintf("/rest/api/latest/project/%s", projectIdOrKey),
	}, 200)
	if err != nil {
		return nil, err
	}

	response := jira.Project{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *ProjectService) Read(projectIdOrKey string) (*jira.Project, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s", projectIdOrKey),
	}, 200)
	if err != nil {
		return nil, err
	}

	response := jira.Project{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *ProjectService) ReadAll() ([]jira.Project, error) {

	var projects []jira.Project
	start := 0
	for {
		reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/rest/api/latest/project/search?startAt=%d", start),
		}, 200)
		if err != nil {
			return nil, err
		}

		response := jira.SearchProjectResponse{}
		err = reply.Object(&response)
		if err != nil {
			return nil, err
		}

		projects = append(projects, response.Projects...)

		if response.IsLast {
			break
		}

		start += response.Total
	}

	return projects, nil
}

func (service *ProjectService) Delete(projectIdOrKey string, enableUndo bool) (bool, error) {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s?enableUndo=%v", projectIdOrKey, enableUndo),
	}, 204)
	if err != nil {
		return false, err
	}

	return true, nil
}
