package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

func (service *RepositoryService) AddDefaultReviewers(project, repository string, reviewers DefaultReviewers) (*DefaultReviewers, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/repos/%s/condition", project, repository),
		Payload: &transport.JsonPayloadData{
			Payload: reviewers,
		},
	}, 200)

	if err != nil {
		return nil, err
	}

	var response *DefaultReviewers

	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *RepositoryService) ReadDefaultReviewers(project, repository string, id int64) (*ReadDefaultReviewers, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/repos/%s/conditions", project, repository),
	}, 200)

	if err != nil {
		return nil, err
	}

	var response []ReadDefaultReviewers

	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	for _, item := range response {
		if item.Id == id {
			return &item, nil
		}
	}

	return nil, nil
}

func (service *RepositoryService) UpdateDefaultReviewers(project, repository string, id int64, reviewers DefaultReviewers) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/repos/%s/condition/%d", project, repository, id),
		Payload: &transport.JsonPayloadData{
			Payload: reviewers,
		},
	}, 200)
	return err
}

func (service *RepositoryService) DeleteDefaultReviewers(project, repository string, id int64) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/repos/%s/condition/%d", project, repository, id),
	}, 200)
	return err
}
