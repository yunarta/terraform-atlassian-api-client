package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"net/url"
)

func (service *RepositoryService) CreateBranchRestrictions(project, repo string, restriction []BranchRestriction) ([]BranchRestrictionReply, error) {
	// Sending a POST request to create a project
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/branch-permissions/latest/projects/%s/repos/%s/restrictions", url.QueryEscape(project), url.QueryEscape(repo)),
		Payload: &XBulkJsonPayload{
			Payload: restriction,
		},
	}, 200)

	if err != nil {
		return nil, err
	}

	var response []BranchRestrictionReply

	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *RepositoryService) ReadBranchRestriction(project, repo string, restrictionId int64) (*BranchRestrictionReply, error) {
	// Sending a POST request to create a project
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/branch-permissions/latest/projects/%s/repos/%s/restrictions/%d", url.QueryEscape(project), url.QueryEscape(repo), restrictionId),
	}, 200)

	if err != nil {
		return nil, err
	}

	var response = &BranchRestrictionReply{}

	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *RepositoryService) DeleteBranchRestriction(project, repo string, restrictionId int64) error {
	// Sending a POST request to create a project
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/branch-permissions/latest/projects/%s/repos/%s/restrictions/%d", url.QueryEscape(project), url.QueryEscape(repo), restrictionId),
	}, 204)
	return err
}
