package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"github.com/yunarta/terraform-atlassian-api-client/util"
	"net/http"
)

const (
	readGroupRepositoryPermissionEndPoint = "/rest/api/latest/projects/%s/repos/%s/permissions/groups?limit=1000%s"
	addGroupRepositoryPermissionEndPoint  = "/rest/api/latest/projects/%s/repos/%s/permissions/groups?name=%s&permission=%s"
	groupRepositoryPermissionEndPoint     = "/rest/api/latest/projects/%s/repos/%s/permissions/groups?name=%s"
	readUserRepositoryPermissionEndPoint  = "/rest/api/latest/projects/%s/repos/%s/permissions/users?limit=1000%s"
	addUserRepositoryPermissionEndPoint   = "/rest/api/latest/projects/%s/repos/%s/permissions/users?name=%s&permission=%s"
	userRepositoryPermissionEndPoint      = "/rest/api/latest/projects/%s/repos/%s/permissions/users?name=%s"
)

func (service *RepositoryService) ReadPermissions(projectKey string, slug string) (*ObjectPermission, error) {
	groupPermissions, err := service.readGroupsPermission(projectKey, slug, "")
	if err != nil {
		return nil, err
	}

	userPermissions, err := service.readUsersPermission(projectKey, slug, "")
	if err != nil {
		return nil, err
	}

	objectPermission := ObjectPermission{
		Groups: groupPermissions.Values,
		Users:  userPermissions.Values,
	}

	return &objectPermission, nil
}

func (service *RepositoryService) readGroupsPermission(projectKey string, slug string, group string) (*GroupPermissionResponse, error) {
	var err error

	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(readGroupRepositoryPermissionEndPoint, projectKey, slug, util.QueryParam("name", group)),
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

func (service *RepositoryService) addGroupPermission(projectKey string, slug string, name string, permission string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(addGroupRepositoryPermissionEndPoint, projectKey, slug, name, permission),
	}, 204)
	return err
}

func (service *RepositoryService) removeGroupPermission(projectKey string, slug string, name string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(groupRepositoryPermissionEndPoint, projectKey, slug, name),
	}, 204)
	return err
}

func (service *RepositoryService) readUsersPermission(projectKey string, slug string, name string) (*UserPermissionResponse, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(readUserRepositoryPermissionEndPoint, projectKey, slug, util.QueryParam("name", name)),
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

func (service *RepositoryService) addUserPermission(projectKey string, slug string, username string, permission string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf(addUserRepositoryPermissionEndPoint, projectKey, slug, username, permission),
	}, 204)
	return err
}

func (service *RepositoryService) removeUserPermission(projectKey string, slug string, username string) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(userRepositoryPermissionEndPoint, projectKey, slug, username),
	}, 204)
	return err
}

func (service *RepositoryService) UpdateGroupPermission(projectKey string, slug string, name string, newPermission string) error {
	groupPermissions, err := service.readGroupsPermission(projectKey, slug, name)
	if err != nil {
		return err
	}

	group := groupPermissions.Find(name)
	if group != nil {
		if len(newPermission) > 0 {
			err = service.addGroupPermission(projectKey, slug, name, newPermission)
			if err != nil {
				return err
			}
		} else {
			err = service.removeGroupPermission(projectKey, slug, name)
			if err != nil {
				return err
			}
		}
	} else if len(newPermission) > 0 {
		err = service.addGroupPermission(projectKey, slug, name, newPermission)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *RepositoryService) UpdateUserPermission(projectKey string, slug string, name string, newPermission string) error {
	userPermissions, err := service.readUsersPermission(projectKey, slug, name)
	if err != nil {
		return err
	}

	user := userPermissions.Find(name)
	if user != nil {
		if len(newPermission) > 0 {
			err = service.addUserPermission(projectKey, slug, name, newPermission)
			if err != nil {
				return err
			}
		} else {
			err = service.removeUserPermission(projectKey, slug, name)
			if err != nil {
				return err
			}
		}
	} else if len(newPermission) > 0 {
		err = service.addUserPermission(projectKey, slug, name, newPermission)
		if err != nil {
			return err
		}
	}

	return nil
}
