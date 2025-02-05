package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

const repositoryEndPoint = "/rest/api/latest/projects/%s/repos"
const repositoryEndPointWithId = "/rest/api/latest/projects/%s/repos/%s"
const getRepositoryLastCommit = "/rest/api/latest/projects/%s/repos/%s/commits?limit=1"
const browserReadMeFile = "/rest/api/latest/projects/%s/repos/%s/browse/README.md"

type RepositoryService struct {
	transport transport.PayloadTransport
}

func (service *RepositoryService) Create(project string, request CreateRepo) (*Repository, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf(repositoryEndPoint, project),
		Payload: transport.JsonPayloadData{
			Payload: request,
		},
	}, 201)
	if err != nil {
		return nil, err
	}

	response := Repository{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *RepositoryService) Read(project string, repo string) (*Repository, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(repositoryEndPointWithId, project, repo),
	}, 200)
	if err != nil {
		return nil, err
	}

	response := Repository{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *RepositoryService) Update(project string, repo string, description string) (*Repository, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(repositoryEndPointWithId, project, repo),
		Payload: transport.JsonPayloadData{
			Payload: map[string]string{
				"description": description,
			},
		},
	}, 200)
	if err != nil {
		return nil, err
	}

	response := Repository{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *RepositoryService) Delete(project string, repo string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(repositoryEndPointWithId, project, repo),
	}, 202)
	return err
}

func (service *RepositoryService) Rename(project string, repo, newRepo string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(repositoryEndPointWithId, project, repo),
		Payload: transport.JsonPayloadData{
			Payload: map[string]string{
				"name": newRepo,
			},
		},
	}, 201)
	return err
}

func (service *RepositoryService) Initialize(project string, repo string, content string) (bool, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(getRepositoryLastCommit, project, repo),
	}, 200)
	if err != nil {
		return false, err
	}

	var response RepositoryCommits
	err = reply.Object(&response)
	if err != nil {
		return false, err
	}

	if len(response.Commits) == 0 {
		_, err = service.transport.Send(&transport.PayloadRequest{
			Method: http.MethodPut,
			Url:    fmt.Sprintf(browserReadMeFile, project, repo),
			Payload: &transport.MultipartPayload{
				Form: map[string]string{
					"message": "Initial commit",
				},
				File: &transport.MultipartFile{
					Key:     "content",
					Name:    "README.md",
					Content: content,
				},
			},
		})

		return true, nil
	} else {
		return false, nil
	}
}
