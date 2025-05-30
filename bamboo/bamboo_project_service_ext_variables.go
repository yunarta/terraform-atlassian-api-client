package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"net/url"
)

func (service *ProjectService) GetVariables(projectKey string, key string) (string, error) {
	// We send a DELETE request to remove a repository from a project. A 204 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/variable/%s", url.QueryEscape(projectKey), url.QueryEscape(key)),
	}, 200)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return "", err
	}

	var variable Variable
	err = reply.Object(&variable)
	if err != nil {
		return "", err
	}

	return variable.Value, nil
}

func (service *ProjectService) PutVariables(projectKey string, key string, value string) error {
	// We send a DELETE request to remove a repository from a project. A 204 status code signifies success.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/variable", url.QueryEscape(projectKey)),
		Payload: transport.JsonPayloadData{
			Payload: map[string]string{
				"name":  key,
				"value": value,
			},
		},
	}, 200, 201)
	// If there's a communication error, we return it immediately.
	return err
}

func (service *ProjectService) DeleteVariables(projectKey string, key string) error {
	// We send a DELETE request to remove a repository from a project. A 204 status code signifies success.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/variable/%s", url.QueryEscape(projectKey), url.QueryEscape(key)),
	}, 204)
	// If there's a communication error, we return it immediately.
	return err
}
