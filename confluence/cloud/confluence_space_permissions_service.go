package cloud

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"net/http"
)

type SpacePermissionsService struct {
	transport transport.PayloadTransport
}

func (service *SpacePermissionsService) GetPermissions(spaceId int64) (*[]confluence.SpacePermission, error) {
	url := fmt.Sprintf("/wiki/api/v2/spaces/%d/permissions", spaceId)
	var permissions = make([]confluence.SpacePermission, 0)
	for {
		reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
			Method: http.MethodGet,
			Url:    url,
		}, 200)
		if err != nil {
			return nil, err
		}

		response := confluence.SpacePermissionResponse{}
		err = reply.Object(&response)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, response.Results...)

		if response.Links.Next == "" {
			break
		} else {
			url = response.Links.Next
		}
	}

	return &permissions, nil
}

func (service *SpacePermissionsService) AddPermission(spaceKey string, request confluence.AddPermissionRequest) (*confluence.SpacePermission, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/wiki/rest/api/space/%s/permission", spaceKey),
		Payload: transport.JsonPayloadData{
			Payload: request,
		},
	}, 200)
	if err != nil {
		return nil, err
	}

	space := confluence.SpacePermission{}
	err = reply.Object(&space)
	if err != nil {
		return nil, err
	}

	return &space, err
}

func (service *SpacePermissionsService) RemovePermission(spaceKey string, request string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/wiki/rest/api/space/%s/permission/%s", spaceKey, request),
	}, 204)
	return err
}
