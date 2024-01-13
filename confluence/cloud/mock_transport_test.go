package cloud

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"os"
)

func MockPayloadTransporter() *transport.MockPayloadTransport {
	return &transport.MockPayloadTransport{
		Payloads: map[string]transport.PayloadResponse{
			"POST:/wiki/rest/api/space": {
				StatusCode: 200,
				Body:       readFile("space_service_create.json"),
			},
			"GET:/wiki/rest/api/space?spaceKey=SK&limit=1&expand=description.plain": {
				StatusCode: 200,
				Body:       readFile("space_service_read.json"),
			},
			"PUT:/wiki/rest/api/space/SK2": {
				StatusCode: 200,
				Body:       readFile("space_service_update2.json"),
			},
			"/wiki/rest/api/space?spaceKey=SK2&limit=1&expand=description.plain": {
				StatusCode: 200,
				Body:       readFile("space_service_read2.json"),
			},
			"DELETE:/wiki/rest/api/space/SK": {
				StatusCode: 202,
				Body:       readFile("space_service_delete.json"),
			},
			"GET:/wiki/rest/api/space?spaceKey=SK3&limit=1&expand=description.plain": {
				StatusCode: 200,
				Body:       readFile("space_service_read3.json"),
			},
			"GET:/wiki/api/v2/spaces/3932162/permissions": {
				StatusCode: 200,
				Body:       readFile("space_service_permissions_read1.json"),
			},
			"GET:/wiki/api/v2/spaces/3932162/permissions?cursor=eyJpZCI6MzkzMjE5Mywic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyIjp7ImZpZWxkIjoiSUQiLCJkaXJlY3Rpb24iOiJBU0NFTkRJTkcifSwic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyVmFsdWUiOjM5MzIxOTN9": {
				StatusCode: 200,
				Body:       readFile("space_service_permissions_read2.json"),
			},
			"GET:/wiki/api/v2/spaces/3932162/permissions?cursor=eyJpZCI6MzkzMjIyMSwic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyIjp7ImZpZWxkIjoiSUQiLCJkaXJlY3Rpb24iOiJBU0NFTkRJTkcifSwic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyVmFsdWUiOjM5MzIyMjF9": {
				StatusCode: 200,
				Body:       readFile("space_service_permissions_read3.json"),
			},
			"GET:/wiki/api/v2/spaces/3932162/permissions?cursor=eyJpZCI6MzkzMjMxNCwic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyIjp7ImZpZWxkIjoiSUQiLCJkaXJlY3Rpb24iOiJBU0NFTkRJTkcifSwic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyVmFsdWUiOjM5MzIzMTR9": {
				StatusCode: 200,
				Body:       readFile("space_service_permissions_read4.json"),
			},
			"POST:/wiki/rest/api/space/SK3/permission": {
				StatusCode: 200,
				Body:       readFile("space_service_permissions_create.json"),
			},
			"DELETE:/wiki/rest/api/space/SK3/permission/3670025": {
				StatusCode: 204,
				Body:       readFile("space_service_permissions_delete.json"),
			},

			"GET:/wiki/rest/api/space?spaceKey=SP&limit=1&expand=description.plain": {
				StatusCode: 200,
				Body:       readFile("space_permission_manager_read.json"),
			},
			"GET:/wiki/api/v2/spaces/3768323/permissions": {
				StatusCode: 200,
				Body:       readFile("space_permission_manager_read_permission1.json"),
			},
			"GET:/wiki/api/v2/spaces/3768323/permissions?cursor=eyJpZCI6MzkzMjUxNCwic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyIjp7ImZpZWxkIjoiSUQiLCJkaXJlY3Rpb24iOiJBU0NFTkRJTkcifSwic3BhY2VQZXJtaXNzaW9uU29ydE9yZGVyVmFsdWUiOjM5MzI1MTR9": {
				StatusCode: 200,
				Body:       readFile("space_permission_manager_read_permission2.json"),
			},
			"GET:/rest/api/latest/user/bulk?accountId=557058:32b276cf-1a9f-45ae-b3f5-f850bc24f1b9": {
				StatusCode: 200,
				Body:       readFile("space_permission_manager_read_user1.json"),
			},
			"GET:/rest/api/latest/group/bulk?groupId=94af0e5e-018a-422d-bffa-e41fc5b71d29": {
				StatusCode: 200,
				Body:       readFile("space_permission_manager_read_group1.json"),
			},
			"POST:/wiki/rest/api/space/SK/permission": {
				StatusCode: 200,
				Body:       readFile("space_permission_manager_write.json"),
			},
			"DELETE:/wiki/rest/api/space/SK/permission": {
				StatusCode: 204,
				Body:       readFile("space_permission_manager_write.json"),
			},
		},
	}
}

func readFile(name string) string {
	bodyBytes, _ := os.ReadFile(fmt.Sprintf("test/%s", name))
	return string(bodyBytes)
}
