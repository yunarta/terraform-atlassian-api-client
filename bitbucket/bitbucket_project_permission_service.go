package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/util"
	"net/http"
)

const (
	groupProjectPermissionEndPoint     = "/rest/api/latest/projects/%s/permissions/groups?name=%s"
	readGroupProjectPermissionEndPoint = "/rest/api/latest/projects/%s/permissions/groups?limit=1000%s"
	addGroupProjectPermissionEndPoint  = "/rest/api/latest/projects/%s/permissions/groups?name=%s&permission=%s"
	userProjectPermissionEndPoint      = "/rest/api/latest/projects/%s/permissions/users?name=%s"
	readUserProjectPermissionEndPoint  = "/rest/api/latest/projects/%s/permissions/users?limit=1000%s"
	addUserProjectPermissionEndPoint   = "/rest/api/latest/projects/%s/permissions/users?name=%s&permission=%s"
)

func (service *ProjectService) ReadPermissions(projectKey string) (*ObjectPermission, error) {
	groupPermissions, err := service.readGroupsPermission(projectKey, "")
	if err != nil {
		return nil, err
	}

	userPermissions, err := service.readUsersPermission(projectKey, "")
	if err != nil {
		return nil, err
	}

	objectPermission := ObjectPermission{
		Groups: groupPermissions.Values,
		Users:  userPermissions.Values,
	}

	return &objectPermission, nil
}

func (service *ProjectService) readGroupsPermission(projectKey string, groupName string) (*GroupPermissionResponse, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(readGroupProjectPermissionEndPoint, projectKey, util.QueryParam("name", groupName)),
	}, 200)
	if err != nil {
		return nil, err
	}

	response := GroupPermissionResponse{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *ProjectService) addGroupPermission(projectKey string, username string, permission string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(addGroupProjectPermissionEndPoint, projectKey, username, permission),
	}, 204)
	return err
}

func (service *ProjectService) removeGroupPermission(projectKey string, username string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(groupProjectPermissionEndPoint, projectKey, username),
	}, 204)
	return err
}

func (service *ProjectService) readUsersPermission(projectKey string, user string) (*UserPermissionResponse, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(readUserProjectPermissionEndPoint, projectKey, util.QueryParam("name", user)),
	}, 200)
	if err != nil {
		return nil, err
	}

	response := UserPermissionResponse{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (service *ProjectService) addUserPermission(projectKey string, username string, permission string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(addUserProjectPermissionEndPoint, projectKey, username, permission),
	}, 204)
	return err
}

func (service *ProjectService) removeUserPermission(projectKey string, username string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(userProjectPermissionEndPoint, projectKey, username),
	}, 204)
	return err
}

func (service *ProjectService) UpdateGroupPermission(projectKey string, groupName string, newPermission string) error {
	groupPermissions, err := service.readGroupsPermission(projectKey, groupName)
	if err != nil {
		return err
	}

	group := groupPermissions.Find(groupName)
	if group != nil {
		if len(newPermission) > 0 {
			err = service.addGroupPermission(projectKey, groupName, newPermission)
		} else {
			err = service.removeGroupPermission(projectKey, groupName)
		}
	} else if len(newPermission) > 0 {
		err = service.addGroupPermission(projectKey, groupName, newPermission)
	}

	return err
}

func (service *ProjectService) UpdateUserPermission(projectKey string, username string, newPermission string) error {
	userPermissions, err := service.readUsersPermission(projectKey, username)
	if err != nil {
		return err
	}

	user := userPermissions.Find(username)
	if user != nil {
		if len(newPermission) > 0 {
			err = service.addUserPermission(projectKey, username, newPermission)
		} else {
			err = service.removeUserPermission(projectKey, username)
		}
	} else if len(newPermission) > 0 {
		err = service.addUserPermission(projectKey, username, newPermission)
	}

	return err
}
