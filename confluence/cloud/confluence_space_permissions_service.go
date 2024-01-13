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

func (service *SpacePermissionsService) Create(spaceKey string, request confluence.AddPermission) (*confluence.Permission, error) {
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

	permission := confluence.Permission{}
	err = reply.Object(&permission)
	if err != nil {
		return nil, err
	}

	return &permission, err
}

func (service *SpacePermissionsService) Read(spaceId int64) (*[]confluence.PermissionV2, error) {
	url := fmt.Sprintf("/wiki/api/v2/spaces/%d/permissions", spaceId)
	var permissions = make([]confluence.PermissionV2, 0)
	for {
		reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
			Method: http.MethodGet,
			Url:    url,
		}, 200)
		if err != nil {
			return nil, err
		}

		response := spacePermissionResponse{}
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

func (service *SpacePermissionsService) Delete(spaceKey string, permissionID string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/wiki/rest/api/space/%s/permission/%s", spaceKey, permissionID),
	}, 204)
	return err
}
