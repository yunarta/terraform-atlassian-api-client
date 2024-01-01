package cloud

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/jira"
	"net/http"
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
	//
	//actors, err := service.ReadProjectRoleActors(projectIdOrKey, roleID)
	//if err != nil {
	//	return false, err
	//}
	//
	//existingUsers, existingGroups := service.computeRoleActors(actors)
	//
	//matchingUsers, _ := collections.Delta(existingUsers, userAccountIds)
	//matchingGroups, _ := collections.Delta(existingGroups, groupIds)

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
	//
	//actors, err := service.ReadProjectRoleActors(projectIdOrKey, roleID)
	//if err != nil {
	//	return false, err
	//}
	//
	//existingUsers, existingGroups := service.computeRoleActors(actors)
	//
	//matchingUsers := collections.Intersect(existingUsers, userAccountIds)
	//matchingGroups := collections.Intersect(existingGroups, groupIds)

	var queries = make([]string, 0)
	queries = append(queries, collections.Map(userAccountIds, func(k string) string {
		return fmt.Sprintf("&user=%s", k)
	})...)
	queries = append(queries, collections.Map(groupIds, func(k string) string {
		return fmt.Sprintf("&groupId=%s", k)
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

func (service *ProjectRoleService) computeRoleActors(actors []jira.Actor) (users []string, groups []string) {
	var existingUsers = make([]string, 0)
	var existingGroups = make([]string, 0)

	for _, actor := range actors {
		switch actor.Type {
		case "atlassian-user-role-actor":
			existingUsers = append(existingUsers, actor.ActorUser.AccountID)
			break
		case "atlassian-group-role-actor":
			existingGroups = append(existingGroups, actor.ActorGroup.GroupId)
			break
		}
	}
	return existingUsers, existingGroups
}

//func (service *ProjectRoleService) UpdateProjectGroupRole(projectIdOrKey string, roleID string, groupId []string) error {
//	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
//		Method: http.MethodPost,
//		Url:    fmt.Sprintf("/rest/api/latest/project/%s/role/%s", projectIdOrKey, roleID),
//		Payload: transport.JsonPayloadData{
//			Payload: map[string][]string{
//				"groupId": groupId,
//			},
//		},
//	}, 200, 400)
//
//	return err
//}

//func (service *ProjectRoleService) UpdateProjectUserRole(projectIdOrKey string, roleID string, userAccountId []string) error {
//	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
//		Method: http.MethodPost,
//		Url:    fmt.Sprintf("/rest/api/latest/project/%s/role/%s", projectIdOrKey, roleID),
//		Payload: transport.JsonPayloadData{
//			Payload: map[string][]string{
//				"user": userAccountId,
//			},
//		},
//	}, 200, 400)
//
//	return err
//}
