package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

func (service *ProjectService) CreateBranchRestrictions(project string, restriction []BranchRestriction) ([]BranchRestrictionReply, error) {
	// Sending a POST request to create a project
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/branch-permissions/latest/projects/%s/restrictions", project),
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

func (service *ProjectService) ReadBranchRestriction(project string, restrictionId int64) (*BranchRestrictionReply, error) {
	// Sending a POST request to create a project
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/branch-permissions/latest/projects/%s/restrictions/%d", project, restrictionId),
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

func (service *ProjectService) DeleteBranchRestriction(project string, restrictionId int64) error {
	// Sending a POST request to create a project
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/branch-permissions/latest/projects/%s/restrictions/%d", project, restrictionId),
	}, 204)
	return err
}
