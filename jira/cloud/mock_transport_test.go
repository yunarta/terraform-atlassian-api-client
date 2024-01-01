package cloud

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"os"
)

func MockPayloadTransporter() *transport.MockPayloadTransport {
	return &transport.MockPayloadTransport{
		Payloads: map[string]transport.PayloadResponse{
			"GET:/rest/api/latest/role": {
				StatusCode: 200,
				Body:       readFile("project_role_service_read_all_role.json"),
			},
			"GET:/rest/api/latest/project/P/role": {
				StatusCode: 200,
				Body:       readFile("project_role_service_read_project_role.json"),
			},
			"GET:/rest/api/latest/project/P/role/10002": {
				StatusCode: 200,
				Body:       readFile("project_role_service_read_project_administrators_role.json"),
			},
			"GET:/rest/api/latest/project/P/role/10008": {
				StatusCode: 200,
				Body:       readFile("project_role_service_read_project_developer_role.json"),
			},
			"POST:/rest/api/latest/project/P/role/10008": {
				StatusCode: 200,
				Body:       "",
			},
			"DELETE:/rest/api/latest/project/P/role/10008?&user=557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9": {
				StatusCode: 204,
				Body:       "",
			},
			"DELETE:/rest/api/latest/project/P/role/10008?&user=557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9&&groupId=557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9": {
				StatusCode: 204,
				Body:       "",
			},
			"DELETE:/rest/api/latest/project/P/role/10008?&user=new-uuid": {
				StatusCode: 204,
				Body:       "",
			},
			"POST:/rest/api/latest/project": {
				StatusCode: 201,
				Body:       readFile("project_service_create.json"),
			},
			"GET:/rest/api/latest/project/PROJECT": {
				StatusCode: 200,
				Body:       readFile("project_service_read.json"),
			},
			"PUT:/rest/api/latest/project/PROJECT": {
				StatusCode: 200,
				Body:       readFile("project_service_edit.json"),
			},
			"GET:/rest/api/latest/project/search?startAt=0": {
				StatusCode: 200,
				Body:       readFile("project_service_read_all.json"),
			},
			"DELETE:/rest/api/latest/project/PROJECT?enableUndo=false": {
				StatusCode: 204,
				Body:       "",
			},
			"GET:/rest/api/latest/project/TEMPLATE": {
				StatusCode: 200,
				Body:       readFile("project_service_read_template.json"),
			},
			"POST:/rest/simplified/latest/project/shared": {
				StatusCode: 200,
				Body:       readFile("project_service_clone.json"),
			},
			"GET:/rest/api/latest/user/search?query=yunarta.kartawahyudi@gmail.com": {
				StatusCode: 200,
				Body:       readFile("user_service_find_user.json"),
			},
			"GET:/rest/api/latest/groups/picker?query=jira-admins-mobilesolutionworks": {
				StatusCode: 200,
				Body:       readFile("user_service_find_group.json"),
			},
		},
	}
}

func readFile(name string) string {
	bodyBytes, _ := os.ReadFile(fmt.Sprintf("test/%s", name))
	return string(bodyBytes)
}
