package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

// UserService struct represents a service to interact with user-related functionalities.
// It uses a PayloadTransport from the transport package for API communication.
type UserService struct {
	transport transport.PayloadTransport
	users     []string
	groups    []string
}

func NewUserService(transport transport.PayloadTransport) *UserService {
	return &UserService{
		transport: transport,
		users:     make([]string, 0),
		groups:    make([]string, 0),
	}
}

// CurrentUser retrieves the current user's information.
// It returns a pointer to a CurrentUser object or an error.
func (service *UserService) CurrentUser() (*CurrentUser, error) {
	// Sending a GET request to the "/rest/api/latest/currentUser" endpoint.
	// Expecting an HTTP status code of 200 for a successful request.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    "/rest/api/latest/currentUser",
	}, 200)

	// Error handling - if an error occurs during the request, return the error.
	if err != nil {
		return nil, err
	}

	// Initializing an empty CurrentUser struct to store the response.
	response := CurrentUser{}
	// Parsing the reply into the response struct.
	err = reply.Object(&response)

	// Error handling - if an error occurs during parsing, return the error.
	if err != nil {
		return nil, err
	}

	// Returning the response object if no errors occurred.
	return &response, nil
}

// FindUser searches for a user by their username.
// It returns a pointer to a User object or an error.
func (service *UserService) FindUser(user string) (*User, error) {
	// Sending a GET request to search for the user.
	// The URL includes a query parameter for filtering by the username.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/admin/users?filter=%s", user),
	}, 200)

	// Error handling - if an error occurs during the request, return the error.
	if err != nil {
		return nil, err
	}

	// Initializing a UserResponse struct to store the response.
	response := UserResponse{}
	// Parsing the reply into the response struct.
	err = reply.Object(&response)

	// Error handling - if an error occurs during parsing, return the error.
	if err != nil {
		return nil, err
	}

	// Iterating through the results to find a matching user.
	for _, item := range response.Results {
		// If a match is found, return the user.
		if item.Name == user {
			return &item, nil
		}
	}

	// If no match is found, return nil.
	return nil, nil
}

// FindGroup searches for a group by its name.
// It returns a pointer to a Group object or an error.
func (service *UserService) FindGroup(group string) (*Group, error) {
	// Sending a GET request to search for the group.
	// The URL includes a query parameter for filtering by the group name.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/admin/groups?filter=%s", group),
	}, 200)

	// Error handling - if an error occurs during the request, return the error.
	if err != nil {
		return nil, err
	}

	// Initializing a GroupResponse struct to store the response.
	response := GroupResponse{}
	// Parsing the reply into the response struct.
	err = reply.Object(&response)

	// Error handling - if an error occurs during parsing, return the error.
	if err != nil {
		return nil, err
	}

	// Iterating through the results to find a matching group.
	for _, item := range response.Results {
		// If a match is found, return the group.
		if item.Name == group {
			return &item, nil
		}
	}

	// If no match is found, return nil.
	return nil, nil
}

func (service *UserService) LookupUser(user string) bool {
	return contains(service.users, user)
}

func (service *UserService) ValidateUser(user string) {
	service.users = append(service.users, user)
}

func (service *UserService) LookupGroup(group string) bool {
	return contains(service.groups, group)
}

func (service *UserService) ValidateGroup(group string) {
	service.groups = append(service.groups, group)
}
