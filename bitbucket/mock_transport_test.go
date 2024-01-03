package bitbucket

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"os"
)

func MockPayloadTransporter() *transport.MockPayloadTransport {
	return &transport.MockPayloadTransport{
		Payloads: map[string]transport.PayloadResponse{
			"POST:/rest/api/latest/projects": {
				StatusCode: 201,
				Body:       readFile("create_project.json"),
			},
			"GET:/rest/api/latest/projects/KEY": {
				StatusCode: 200,
				Body:       readFile("create_project.json"),
			},
			"PUT:/rest/api/latest/projects/KEY": {
				StatusCode: 200,
				Body:       readFile("update_project.json"),
			},
			// project permissions
			"GET:/rest/api/latest/projects/A/permissions/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("project_a_permissions_users_read.json"),
			},
			"GET:/rest/api/latest/projects/A/permissions/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("project_a_permissions_groups_read.json"),
			},
			"GET:/rest/api/latest/projects/B/permissions/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("project_b_permissions_users_read.json"),
			},
			"GET:/rest/api/latest/projects/B/permissions/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("project_b_permissions_groups_read.json"),
			},

			"GET:/rest/api/latest/projects/A/permissions/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("project_a_permissions_users_read.json"),
			},
			"DELETE:/rest/api/latest/projects/A/permissions/users?name=yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/projects/A/permissions/users?name=yunarta&permission=PROJECT_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			"GET:/rest/api/latest/projects/A/permissions/groups?limit=1000&name=bitbucket-admin": {
				StatusCode: 200,
				Body:       readFile("project_a_permissions_groups_read.json"),
			},
			"DELETE:/rest/api/latest/projects/A/permissions/groups?name=bitbucket-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/projects/A/permissions/groups?name=bitbucket-admin&permission=PROJECT_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			"GET:/rest/api/latest/projects/B/permissions/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("project_b_permissions_users_read.json"),
			},
			"PUT:/rest/api/latest/projects/B/permissions/users?name=yunarta&permission=PROJECT_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			"GET:/rest/api/latest/projects/B/permissions/groups?limit=1000&name=bitbucket-admin": {
				StatusCode: 200,
				Body:       readFile("project_b_permissions_groups_read.json"),
			},
			"PUT:/rest/api/latest/projects/B/permissions/groups?name=bitbucket-admin&permission=PROJECT_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			// repository permissions
			"GET:/rest/api/latest/projects/A/repos/r/permissions/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("repository_a_r_permissions_users_read.json"),
			},
			"GET:/rest/api/latest/projects/A/repos/r/permissions/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("repository_a_r_permissions_groups_read.json"),
			},
			"GET:/rest/api/latest/projects/B/repos/r/permissions/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("repository_b_r_permissions_users_read.json"),
			},
			"GET:/rest/api/latest/projects/B/repos/r/permissions/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("repository_b_r_permissions_groups_read.json"),
			},

			"GET:/rest/api/latest/projects/A/repos/r/permissions/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("repository_a_r_permissions_users_read.json"),
			},
			"DELETE:/rest/api/latest/projects/A/repos/r/permissions/users?name=yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/projects/A/repos/r/permissions/users?name=yunarta&permission=REPO_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			"GET:/rest/api/latest/projects/A/repos/r/permissions/groups?limit=1000&name=bitbucket-admin": {
				StatusCode: 200,
				Body:       readFile("repository_a_r_permissions_groups_read.json"),
			},
			"DELETE:/rest/api/latest/projects/A/repos/r/permissions/groups?name=bitbucket-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/projects/A/repos/r/permissions/groups?name=bitbucket-admin&permission=REPO_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			"GET:/rest/api/latest/projects/B/repos/r/permissions/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("repository_b_r_permissions_users_read.json"),
			},
			"PUT:/rest/api/latest/projects/B/repos/r/permissions/users?name=yunarta&permission=REPO_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			"GET:/rest/api/latest/projects/B/repos/r/permissions/groups?limit=1000&name=bitbucket-admin": {
				StatusCode: 200,
				Body:       readFile("repository_b_r_permissions_groups_read.json"),
			},
			"PUT:/rest/api/latest/projects/B/repos/r/permissions/groups?name=bitbucket-admin&permission=REPO_ADMIN": {
				StatusCode: 204,
				Body:       "",
			},

			// repository
			"POST:/rest/api/latest/projects/A/repos": {
				StatusCode: 201,
				Body:       readFile("project_repo_create.json"),
			},
			"GET:/rest/api/latest/projects/A/repos/name": {
				StatusCode: 200,
				Body:       readFile("project_repo_create.json"),
			},
			"PUT:/rest/api/latest/projects/A/repos/name": {
				StatusCode: 200,
				Body:       readFile("project_repo_update.json"),
			},
			"DELETE:/rest/api/latest/projects/A/repos/name": {
				StatusCode: 202,
				Body:       "",
			},
			"GET:/rest/api/latest/projects/A/repos/new/commits?limit=1": {
				StatusCode: 200,
				Body:       readFile("project_repo_new_commits.json"),
			},
			"PUT:/rest/api/latest/projects/A/repos/new/browse/README.md": {
				StatusCode: 200,
				Body:       "",
			},
			"GET:/rest/api/latest/projects/A/repos/exists/commits?limit=1": {
				StatusCode: 200,
				Body:       readFile("project_repo_existings_commits.json"),
			},
			"PUT:/rest/api/latest/projects/A/repos/exists/browse/README.md": {
				StatusCode: 200,
				Body:       "",
			},

			"GET:/rest/api/latest/users?filter=yunarta": {
				StatusCode: 200,
				Body:       readFile("find_users.json"),
			},
			"GET:/rest/api/latest/admin/groups?filter=bitbucket-admin": {
				StatusCode: 200,
				Body:       readFile("find_groups.json"),
			},
		},
	}
}

func readFile(name string) string {
	bodyBytes, _ := os.ReadFile(fmt.Sprintf("test/%s", name))
	return string(bodyBytes)
}
