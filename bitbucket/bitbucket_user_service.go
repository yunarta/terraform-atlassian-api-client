package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

const findUser = "/rest/api/latest/users?filter=%s"
const findGroup = "/rest/api/latest/admin/groups?filter=%s"

type UserService struct {
	transport transport.PayloadTransport
}

func (service *UserService) FindUser(user string) (*User, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(findUser, user),
	}, 200)

	if err != nil {
		return nil, err
	}

	response := UserResponse{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	for _, test := range response.Values {
		if test.Name == user {
			return &test, nil
		}
	}

	return nil, nil
}

func (service *UserService) FindGroup(group string) (*Group, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(findGroup, group),
	}, 200)

	if err != nil {
		return nil, err
	}

	response := GroupResponse{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	for _, item := range response.Values {
		if item.Name == group {
			return &Group{
				Name: item.Name,
			}, nil
		}
	}

	return nil, nil
}
