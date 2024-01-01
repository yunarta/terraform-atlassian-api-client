package bamboo

import (
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

// PermissionsHelper is a struct that assists in managing permissions.
// It uses a transport layer for API communications and maintains a URL and a list of permissions.
type PermissionsHelper struct {
	Transport   transport.PayloadTransport
	Url         string
	Permissions []string
}

// ReadGroupPermissions fetches group permissions and returns them in a GroupPermissionResponse struct.
// It returns an error if the operation fails.
func (helper PermissionsHelper) ReadGroupPermissions() (*GroupPermissionResponse, error) {
	response := GroupPermissionResponse{} // Initialize an empty response structure.

	// Call readPermissions to fetch permissions. If an error occurs, it returns the error.
	err := helper.readPermissions(&response)
	if err != nil {
		return nil, err
	}

	// Returns the filled response structure.
	return &response, nil
}

// ReadUserPermissions works similarly to ReadGroupPermissions but for user permissions.
func (helper PermissionsHelper) ReadUserPermissions() (*UserPermissionResponse, error) {
	response := UserPermissionResponse{} // Initialize an empty response structure.
	err := helper.readPermissions(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ReadRolePermissions works similarly to ReadGroupPermissions but for role permissions.
func (helper PermissionsHelper) ReadRolePermissions() (*RolePermissionResponse, error) {
	response := RolePermissionResponse{} // Initialize an empty response structure.
	err := helper.readPermissions(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// readPermissions is a helper function that makes an HTTP GET request to fetch permissions.
// It unmarshal the response into the provided response structure.
func (helper PermissionsHelper) readPermissions(response any) error {
	// Send an HTTP GET request and expect a 200 status code in response.
	reply, err := helper.Transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    helper.Url,
	}, 200)

	// If there's an error in sending the request or receiving the response, return the error.
	if err != nil {
		return err
	}

	// Unmarshal the response body into the provided response structure.
	err = reply.Object(response)
	if err != nil {
		return err
	}

	return nil
}

// AddPermissions sends a PUT request to add permissions.
// If there are permissions to add, it returns the result of the transport operation, otherwise nil.
func (helper PermissionsHelper) AddPermissions() error {
	if len(helper.Permissions) > 0 {
		_, err := helper.Transport.SendWithExpectedStatus(&transport.PayloadRequest{
			Method: http.MethodPut,
			Url:    helper.Url,
			Payload: transport.JsonPayloadData{
				Payload: helper.Permissions,
			},
		}, 204, 304)
		return err
	}

	return nil
}

// RemovePermissions sends a DELETE request to remove permissions.
// If there are permissions to remove, it returns the result of the transport operation, otherwise nil.
func (helper PermissionsHelper) RemovePermissions() error {
	if len(helper.Permissions) > 0 {
		_, err := helper.Transport.SendWithExpectedStatus(&transport.PayloadRequest{
			Method: http.MethodDelete,
			Url:    helper.Url,
			Payload: transport.JsonPayloadData{
				Payload: helper.Permissions,
			},
		}, 204, 304)
		return err
	}

	return nil
}

// updateItemPermission is a generic function to update permissions of a specific item.
// It applies new permissions and removes old ones as needed.
func updateItemPermission[K any](
	response PermissionResponse,
	itemId K,
	key string,
	newPermissions []string,
	addPermissions func(itemId K, role string, permissions []string) error,
	removePermissions func(itemId K, role string, permissions []string) error,
) error {
	var err error

	// Find the item in the response using the provided key.
	item := response.Find(key)
	// If the item is found, calculate the permissions to add and remove.
	if item != nil {
		adding, removing := item.DeltaPermissions(newPermissions)
		// Add new permissions if there are any to add.
		if len(adding) > 0 {
			err = addPermissions(itemId, key, adding)
			if err != nil {
				return err
			}
		}

		// Remove old permissions if there are any to remove.
		if len(removing) > 0 {
			err = removePermissions(itemId, key, removing)
			if err != nil {
				return err
			}
		}
	} else if len(newPermissions) > 0 {
		// If the item is not found but there are new permissions, add them.
		err = addPermissions(itemId, key, newPermissions)
		if err != nil {
			return err
		}
	}

	return nil
}
