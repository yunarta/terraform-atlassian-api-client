package cloud

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
	"net/http"
	"net/url"
	"strings"
)

type ActorService struct {
	transport transport.PayloadTransport
}

func NewActorService(transport transport.PayloadTransport) *ActorService {
	return &ActorService{transport: transport}
}

func (service *ActorService) ReadUser(emailAddress string) (*jira.User, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/user/search?query=%s", url.QueryEscape(emailAddress)),
	}, 200)
	if err != nil {
		return nil, err
	}

	var users []jira.User
	err = reply.Object(&users)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.EmailAddress == emailAddress {
			return &user, nil
		}
	}

	return nil, nil
}

func (service *ActorService) ReadGroup(name string) (*jira.Group, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/groups/picker?query=%s", url.QueryEscape(name)),
	}, 200)
	if err != nil {
		return nil, err
	}

	var response jira.FindGroupsResponse
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	for _, group := range response.Groups {
		if group.Name == name {
			return &group, nil
		}
	}

	return nil, nil
}

func (service *ActorService) BulkGetUsers(accountIds []string) ([]jira.User, error) {
	accountIdQuery := strings.Join(collections.Map(accountIds, func(e string) string {
		return fmt.Sprintf("accountId=%s", url.QueryEscape(e))
	}), "&")

	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/user/bulk?%s", accountIdQuery),
	}, 200)
	if err != nil {
		return nil, err
	}

	var response jira.BulkUsersResponse
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response.Users, nil
}

func (service *ActorService) BulkGetGroupsById(groupId []string) ([]jira.Group, error) {
	groupIdQuery := strings.Join(collections.Map(groupId, func(e string) string {
		return fmt.Sprintf("groupId=%s", url.QueryEscape(e))
	}), "&")

	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/group/bulk?%s", groupIdQuery),
	}, 200)
	if err != nil {
		return nil, err
	}

	var response jira.BulkGroupsResponse
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response.Groups, nil
}

func (service *ActorService) BulkGetGroupsByName(groupName []string) ([]jira.Group, error) {
	groupNameQuery := strings.Join(collections.Map(groupName, func(e string) string {
		return fmt.Sprintf("groupName=%s", url.QueryEscape(e))
	}), "&")

	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/group/bulk?%s", groupNameQuery),
	}, 200)
	if err != nil {
		return nil, err
	}

	var response jira.BulkGroupsResponse
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return response.Groups, nil
}
