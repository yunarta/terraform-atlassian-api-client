package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"net/url"
)

func (service *ProjectService) AddDefaultReviewers(project string, reviewers DefaultReviewers) (*DefaultReviewers, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/condition", url.QueryEscape(project)),
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

func (service *ProjectService) ReadDefaultReviewers(project string, id int64) (*ReadDefaultReviewers, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/conditions", url.QueryEscape(project)),
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

func (service *ProjectService) UpdateDefaultReviewers(project string, id int64, reviewers DefaultReviewers) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/condition/%d", url.QueryEscape(project), id),
		Payload: &transport.JsonPayloadData{
			Payload: reviewers,
		},
	}, 200)
	return err
}

func (service *ProjectService) DeleteDefaultReviewers(project string, id int64) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/default-reviewers/1.0/projects/%s/condition/%d", url.QueryEscape(project), id),
	}, 200)
	return err
}
