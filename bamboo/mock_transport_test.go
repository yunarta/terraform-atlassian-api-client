package bamboo

import (
	"github.com/yunarta/terraform-api-transport/transport"
	"os"
)

func MockPayloadTransporter() *transport.MockPayloadTransport {
	return &transport.MockPayloadTransport{
		Payloads: map[string]transport.PayloadResponse{
			// deployment permissions
			"GET:/rest/api/latest/permissions/deployment/0/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_deployment_0_groups.json"),
			},
			"GET:/rest/api/latest/permissions/deployment/0/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_deployment_0_users.json"),
			},
			"GET:/rest/api/latest/permissions/deployment/0/roles": {
				StatusCode: 200,
				Body:       readFile("test/permissions_deployment_0_roles.json"),
			},
			"GET:/rest/api/latest/permissions/deployment/0/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("test/permissions_deployment_0_users.json"),
			},
			"DELETE:/rest/api/latest/permissions/deployment/0/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/deployment/0/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/deployment/0/groups?limit=1000&name=bamboo-admin": {
				StatusCode: 200,
				Body:       readFile("test/permissions_deployment_0_groups.json"),
			},
			"DELETE:/rest/api/latest/permissions/deployment/0/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/deployment/0/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/deployment/0/roles?name=LOGGED_IN": {
				StatusCode: 200,
				Body:       readFile("test/permissions_deployment_0_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/deployment/0/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/deployment/0/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},

			// repository permissions
			"GET:/rest/api/latest/permissions/repository/0/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_repository_0_groups.json"),
			},
			"GET:/rest/api/latest/permissions/repository/0/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_repository_0_users.json"),
			},
			"GET:/rest/api/latest/permissions/repository/0/roles": {
				StatusCode: 200,
				Body:       readFile("test/permissions_repository_0_roles.json"),
			},
			"GET:/rest/api/latest/permissions/repository/0/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("test/permissions_repository_0_users.json"),
			},
			"DELETE:/rest/api/latest/permissions/repository/0/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/repository/0/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/repository/0/groups?limit=1000&name=bamboo-admin": {
				StatusCode: 200,
				Body:       readFile("test/permissions_repository_0_groups.json"),
			},
			"DELETE:/rest/api/latest/permissions/repository/0/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/repository/0/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/repository/0/roles?name=LOGGED_IN": {
				StatusCode: 200,
				Body:       readFile("test/permissions_repository_0_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/repository/0/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/repository/0/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/deployment/1/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("test/permissions_deployment_1_users.json"),
			},
			"PUT:/rest/api/latest/permissions/deployment/1/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},

			// project permissions
			"GET:/rest/api/latest/permissions/project/PROJECT/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_1_groups.json"),
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_1_users.json"),
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/roles": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_1_roles.json"),
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_1_users.json"),
			},
			"DELETE:/rest/api/latest/permissions/project/PROJECT/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/project/PROJECT/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/groups?limit=1000&name=bamboo-admin": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_1_groups.json"),
			},
			"DELETE:/rest/api/latest/permissions/project/PROJECT/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/project/PROJECT/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/roles?name=LOGGED_IN": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_1_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/project/PROJECT/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/project/PROJECT/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/roles?name=ANONYMOUS": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_1_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/project/PROJECT/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/project/PROJECT/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/project/PROJECT2/roles": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_2_roles.json"),
			},
			"GET:/rest/api/latest/permissions/project/PROJECT2/roles?name=ANONYMOUS": {
				StatusCode: 200,
				Body:       readFile("test/permissions_project_2_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/project/PROJECT2/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/project/PROJECT2/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},

			// project plan permissions
			"GET:/rest/api/latest/permissions/projectplan/PROJECT/groups?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_1_groups.json"),
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT/users?limit=1000": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_1_users.json"),
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT/roles": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_1_roles.json"),
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT/users?limit=1000&name=yunarta": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_1_users.json"),
			},
			"DELETE:/rest/api/latest/permissions/projectplan/PROJECT/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/projectplan/PROJECT/users/yunarta": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT/groups?limit=1000&name=bamboo-admin": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_1_groups.json"),
			},
			"DELETE:/rest/api/latest/permissions/projectplan/PROJECT/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/projectplan/PROJECT/groups/bamboo-admin": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT/roles?name=LOGGED_IN": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_1_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/projectplan/PROJECT/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/projectplan/PROJECT/roles/LOGGED_IN": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT/roles?name=ANONYMOUS": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_1_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/projectplan/PROJECT/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/projectplan/PROJECT/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT2/roles": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_2_roles.json"),
			},
			"GET:/rest/api/latest/permissions/projectplan/PROJECT2/roles?name=ANONYMOUS": {
				StatusCode: 200,
				Body:       readFile("test/permissions_projectplan_2_roles.json"),
			},
			"DELETE:/rest/api/latest/permissions/projectplan/PROJECT2/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},
			"PUT:/rest/api/latest/permissions/projectplan/PROJECT2/roles/ANONYMOUS": {
				StatusCode: 204,
				Body:       "",
			},

			// deployment projects
			"PUT:/rest/api/latest/deploy/project": {
				StatusCode: 200,
				Body:       readFile("test/deploy_project_create.json"),
			},
			"GET:/rest/api/latest/search/deployments?searchTerm=dn": {
				StatusCode: 200,
				Body:       readFile("test/search_deployment_dn.json"),
			},
			"GET:/rest/api/latest/deploy/project/2": {
				StatusCode: 200,
				Body:       readFile("test/deploy_project_2.json"),
			},
			"POST:/rest/api/latest/deploy/project/2": {
				StatusCode: 200,
				Body:       readFile("test/deploy_project_2_update.json"),
			},
			"DELETE:/rest/api/latest/deploy/project/2": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/deploy/project/2/repository": {
				StatusCode: 200,
				Body:       readFile("test/deploy_project_2_repository_get.json"),
			},
			"POST:/rest/api/latest/deploy/project/2/repository": {
				StatusCode: 201,
				Body:       readFile("test/deploy_project_2_repository_post.json"),
			},
			"DELETE:/rest/api/latest/deploy/project/2/repository/1": {
				StatusCode: 204,
				Body:       "",
			},

			// project
			"POST:/rest/api/latest/project": {
				StatusCode: 201,
				Body:       readFile("test/project_create.json"),
			},
			"GET:/rest/api/latest/project/PROJECT": {
				StatusCode: 200,
				Body:       readFile("test/project_create.json"),
			},
			"PUT:/rest/api/latest/project/PROJECT": {
				StatusCode: 200,
				Body:       readFile("test/project_update.json"),
			},
			"DELETE:/rest/api/latest/project/PROJECT": {
				StatusCode: 204,
				Body:       "{}",
			},
			"GET:/rest/api/latest/plan/PROJECT-PLAN": {
				StatusCode: 200,
				Body:       readFile("test/project_plan.json"),
			},
			"GET:/rest/api/latest/project/PROJECT/repository": {
				StatusCode: 200,
				Body:       readFile("test/project_1_repository_get.json"),
			},
			"POST:/rest/api/latest/project/PROJECT/repository": {
				StatusCode: 201,
				Body:       readFile("test/project_1_repository_post.json"),
			},
			"DELETE:/rest/api/latest/project/PROJECT/repository/1": {
				StatusCode: 204,
				Body:       "",
			},

			"POST:/rest/api/latest/import/repository": {
				StatusCode: 200,
				Body:       "https://bamboo.funf?repositoryId=1",
			},
			"GET:/rest/api/latest/repository?max-result=1000&searchTerm=oraclelinux": {
				StatusCode: 200,
				Body:       readFile("test/repository_search.json"),
			},
			"GET:/rest/api/latest/repository?max-result=1000&searchTerm=ubuntu": {
				StatusCode: 200,
				Body:       readFile("test/repository_search.json"),
			},
			"PUT:/rest/api/latest/repository/1/enableCi": {
				StatusCode: 200,
				Body:       "",
			},
			"GET:/rest/api/latest/repository/1/rssrepository": {
				StatusCode: 200,
				Body:       readFile("test/repository_1_rssrepository.json"),
			},
			"POST:/rest/api/latest/repository/1/rssrepository": {
				StatusCode: 201,
				Body:       readFile("test/repository_1_rssrepository_post.json"),
			},
			"DELETE:/rest/api/latest/repository/1/rssrepository/2": {
				StatusCode: 204,
				Body:       readFile("test/repository_1_rssrepository_post.json"),
			},
			"GET:/rest/api/latest/currentUser": {
				StatusCode: 200,
				Body:       readFile("test/currentUser.json"),
			},
			"GET:/rest/api/latest/admin/users?filter=yunarta": {
				StatusCode: 200,
				Body:       readFile("test/admin_users_find_yunarta.json"),
			},
			"GET:/rest/api/latest/admin/groups?filter=bamboo-admin": {
				StatusCode: 200,
				Body:       readFile("test/admin_groups_find_admin.json"),
			},
		},
	}
}

func readFile(name string) string {
	bodyBytes, _ := os.ReadFile(name)
	return string(bodyBytes)
}
