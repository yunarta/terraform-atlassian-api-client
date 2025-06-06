package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"net/url"
	"strings"
)

const findUser = "/rest/api/latest/users?filter=%s"
const findGroup = "/rest/api/latest/groups?filter=%s"

type UserService struct {
	transport transport.PayloadTransport
	userCache map[string]User
}

func NewUserService(transport transport.PayloadTransport) *UserService {
	return &UserService{
		transport: transport,
		userCache: make(map[string]User),
	}
}

func (service *UserService) FindUser(user string) (*User, error) {
	foundUser, ok := service.userCache[user]
	if ok {
		return &foundUser, nil
	}

	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(findUser, url.QueryEscape(user)),
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
		if strings.EqualFold(test.Name, user) {
			service.userCache[test.EmailAddress] = test
			return &test, nil
		}
	}

	return nil, nil
}

func (service *UserService) FindGroup(group string) (*Group, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(findGroup, url.QueryEscape(group)),
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
		if strings.EqualFold(item, group) {
			return &Group{
				Name: item,
			}, nil
		}
	}

	return nil, nil
}
