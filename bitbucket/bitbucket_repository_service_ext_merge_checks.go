package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"strconv"
)

func (service *RepositoryService) EnableMergeCheck(project, repo, check string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/rest/api/latest/projects/%s/repos/%s/settings/hooks/%s/enabled?enrich=true", project, repo, check),
	}, 200)
	return err
}

func (service *RepositoryService) DisableMergeCheck(project, repo, check string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/api/latest/projects/%s/repos/%s/settings/hooks/%s/enabled?enrich=true", project, repo, check),
	}, 200)
	return err
}

func (service *RepositoryService) ConfigureMergeCheck(project, repo, check string, value int) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/rest/api/latest/projects/%s/repos/%s/settings/hooks/%s/enabled?enrich=true", project, repo, check),
		Payload: &transport.JsonPayloadData{
			Payload: map[string]string{
				"requiredCount": strconv.Itoa(value),
			},
		},
	}, 200)
	return err
}

func (service *RepositoryService) GetMergeChecks(project, repo string) ([]MergeCheck, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/projects/%s/repos/%s/settings/hooks", project, repo),
	}, 200)

	if err != nil {
		return nil, err
	}

	response := &MergeChecksReply{}
	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response.Values, nil
}

func (service *RepositoryService) GetMergeCheckSetting(project, check string) (*MergeCheckSetting, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/projects/%s/settings/hooks/%s/settings", project, check),
	}, 200)

	if err != nil {
		return nil, err
	}

	response := &MergeCheckSetting{}
	// Parsing the response
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
