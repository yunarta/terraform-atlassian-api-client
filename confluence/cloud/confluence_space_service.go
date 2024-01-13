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

func (service *SpaceService) Create(request confluence.CreateSpace) (*confluence.Space, error) {
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

func (service *SpaceService) Read(spaceKey string) (*confluence.Space, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/wiki/rest/api/space?spaceKey=%s&limit=1&expand=description.plain", spaceKey),
	}, 200)
	if err != nil {
		return nil, err
	}

	space := getSpacesResponse{}
	err = reply.Object(&space)
	if err != nil {
		return nil, err
	}

	if len(space.Results) == 0 {
		return nil, nil
	}

	return &space.Results[0], err
}

func (service *SpaceService) Update(spaceKey string, request confluence.UpdateSpace) (*confluence.Space, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/wiki/rest/api/space/%s", spaceKey),
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

func (service *SpaceService) Delete(spaceKey string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/wiki/rest/api/space/%s", spaceKey),
	}, 202)
	return err
}
