package cloud

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/confluence"
	"net/http"
)

type SpaceService struct {
	transport transport.PayloadTransport
}

func (service *SpaceService) Create(request confluence.CreateSpaceRequest) (*confluence.Space, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/wiki/rest/api/space",
		Payload: transport.JsonPayloadData{
			Payload: request,
		},
	}, 200)
	if err != nil {
		return nil, err
	}

	space := confluence.Space{}
	err = reply.Object(&space)
	if err != nil {
		return nil, err
	}

	return &space, err
}

func (service *SpaceService) Get(spaceKey string) (*confluence.Space, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/wiki/rest/api/space?spaceKey=%s&limit=1", spaceKey),
	}, 200)
	if err != nil {
		return nil, err
	}

	space := confluence.GetSpacesResponse{}
	err = reply.Object(&space)
	if err != nil {
		return nil, err
	}

	if len(space.Results) == 0 {
		return nil, nil
	}

	return &space.Results[0], err
}

func (service *SpaceService) Delete(spaceKey string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/wiki/rest/api/space/%s", spaceKey),
	}, 202)
	return err
}
