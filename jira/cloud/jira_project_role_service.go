package cloud

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

type ProjectRoleService struct {
	transport transport.PayloadTransport
}

func (service *ProjectRoleService) ReadAllRole() ([]jira.Role, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    "/rest/api/latest/role",
	}, 200)
	if err != nil {
		return nil, err
	}

	var rl []jira.Role
	err = reply.Object(&rl)
	if err != nil {
		return nil, err
	}

	return rl, nil
}

func (service *ProjectRoleService) ReadProjectRoles(projectIdOrKey string) ([]jira.RoleType, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/role", projectIdOrKey),
	}, 200)
	if err != nil {
		return nil, err
	}

	var doc interface{}
	err = reply.Object(&doc)
	if err != nil {
		return nil, err
	}

	var rl []jira.RoleType
	for k, v := range doc.(map[string]interface{}) {
		var r jira.RoleType
		r.Name = k
		r.RollLink = v.(string)

		pos := strings.LastIndex(r.RollLink, "/role/")
		adjustedPos := pos + len("/role/")
		r.ID = r.RollLink[adjustedPos:len(r.RollLink)]
		rl = append(rl, r)
	}

	slices.SortFunc(rl, func(a, b jira.RoleType) int {
		return strings.Compare(a.Name, b.Name)
	})
	return rl, nil
}

func (service *ProjectRoleService) ReadProjectRoleActors(projectIdOrKey string, roleID string) ([]jira.Actor, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/role/%s", projectIdOrKey, roleID),
	}, 200)
	if err != nil {
		return nil, err
	}

	var doc jira.Role
	err = reply.Object(&doc)
	if err != nil {
		return nil, err
	}

	return doc.Actors, nil
}

func (service *ProjectRoleService) AddProjectRole(projectIdOrKey string, roleID string, userAccountIds []string, groupIds []string) error {
	var err error

	if len(userAccountIds) > 0 || len(groupIds) > 0 {
		_, err = service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
			Method: http.MethodPost,
			Url:    fmt.Sprintf("/rest/api/latest/project/%s/role/%s", projectIdOrKey, roleID),
			Payload: transport.JsonPayloadData{
				Payload: map[string][]string{
					"user":    userAccountIds,
					"groupId": groupIds,
				},
			},
		}, 200)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *ProjectRoleService) RemoveProjectRole(projectIdOrKey string, roleID string, userAccountIds []string, groupIds []string) error {
	var err error

	var queries = make([]string, 0)
	queries = append(queries, collections.Map(userAccountIds, func(k string) string {
		return fmt.Sprintf("&user=%s", url.QueryEscape(k))
	})...)
	queries = append(queries, collections.Map(groupIds, func(k string) string {
		return fmt.Sprintf("&groupId=%s", url.QueryEscape(k))
	})...)

	if len(queries) > 0 {
		queryString := strings.Join(queries, "&")

		_, err = service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
			Method: http.MethodDelete,
			Url:    fmt.Sprintf("/rest/api/latest/project/%s/role/%s?%s", projectIdOrKey, roleID, queryString),
		}, 204)
		if err != nil {
			return err
		}
	}

	return nil
}
