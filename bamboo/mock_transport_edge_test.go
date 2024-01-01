package bamboo

import (
	"github.com/yunarta/terraform-api-transport/transport"
)

// EdgeTestMockPayloadTransporter help test situation where the requested resource not exists.
// Bamboo usually return 400 or 404 when the resource in question cannot be found
func EdgeTestMockPayloadTransporter() *transport.MockPayloadTransport {
	return &transport.MockPayloadTransport{
		Payloads: map[string]transport.PayloadResponse{
			// deployment permissions
			"GET:/rest/api/latest/permissions/deployment/0/groups?limit=1000": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/permissions/deployment/0/users?limit=1000&name=yunarta": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/permissions/deployment/0/groups?limit=1000&name=bamboo-admin": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/permissions/deployment/0/roles": {
				StatusCode: 400,
			},
			"PUT:/rest/api/latest/deploy/project": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/search/deployments?searchTerm=dn": {
				StatusCode: 200,
				Body:       "{\"size\": 1, \"searchResults\": [], \"start-index\": 0, \"max-result\": 1}",
			},
			"GET:/rest/api/latest/deploy/project/2": {
				StatusCode: 400,
			},
			"POST:/rest/api/latest/deploy/project/2": {
				StatusCode: 400,
			},
			"DELETE:/rest/api/latest/deploy/project/2": {
				StatusCode: 404,
			},
			"GET:/rest/api/latest/deploy/project/2/repository": {
				StatusCode: 404,
			},
			"POST:/rest/api/latest/deploy/project/2/repository": {
				StatusCode: 400,
			},
			"DELETE:/rest/api/latest/deploy/project/2/repository/1": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/groups?limit=1000": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/users?limit=1000&name=yunarta": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/groups?limit=1000&name=bamboo-admin": {
				StatusCode: 400,
			},
			"GET:/rest/api/latest/permissions/project/PROJECT/roles": {
				StatusCode: 400,
			},
		},
	}
}
