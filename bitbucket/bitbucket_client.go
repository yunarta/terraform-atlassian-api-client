package bitbucket

import "github.com/yunarta/terraform-api-transport/transport"

type Client struct {
	projectService    *ProjectService
	repositoryService *RepositoryService
	userService       *UserService
}

func NewBitbucketClient(transport transport.PayloadTransport) *Client {
	return &Client{
		projectService:    &ProjectService{transport: transport},
		repositoryService: &RepositoryService{transport: transport},
		userService:       &UserService{transport: transport},
	}
}

func (client *Client) ProjectService() *ProjectService {
	return client.projectService
}

func (client *Client) RepositoryService() *RepositoryService {
	return client.repositoryService
}
func (client *Client) UserService() *UserService {
	return client.userService
}
