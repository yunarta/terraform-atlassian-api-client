package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

// RepositoryService is a struct that provides methods to interact with repositories.
// It contains a transport mechanism for API calls.
type RepositoryService struct {
	transport transport.PayloadTransport
}

// CreateRepository holds the data necessary to create a new repository.
type CreateRepository struct {
	Name             string
	ProjectKey       string
	RepositorySlug   string
	RepositoryBranch string
	ServerId         string
	ServerName       string
	CloneUrl         string
}

// CreateProjectRepository holds the data necessary to create a new repository.
type CreateProjectRepository struct {
	Project          string
	Name             string
	ProjectKey       string
	RepositorySlug   string
	RepositoryBranch string
	ServerId         string
	ServerName       string
	CloneUrl         string
}

// Create initializes a new repository with the specified parameters.
// It returns the created repository's ID or an error if the creation fails.
func (service *RepositoryService) Create(create CreateRepository) (int, error) {
	// Send a POST request to create a repository with the specified details.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/rest/api/latest/import/repository",
		// The payload is formatted in YAML and specifies various repository properties.
		Payload: &XYamlPayload{
			Data: fmt.Sprintf(`
!!com.atlassian.bamboo.specs.util.BambooSpecProperties
rootEntity: !!com.atlassian.bamboo.specs.model.repository.bitbucket.server.BitbucketServerRepositoryProperties
  name: %s
  projectKey: %s
  repositorySlug: %s
  branch: master
  server:
    id: %s
    name: %s
  sshCloneUrl: %s
  useLfs: false
specModelVersion: 9.3.0
`,
				create.Name,
				create.ProjectKey,
				create.RepositorySlug,
				create.ServerId,
				create.ServerName,
				create.CloneUrl,
			),
		},
	}, 200)

	// Handle any errors that occur during the API call.
	if err != nil {
		return 0, err
	}

	// Parse the URL from the response to extract the repository ID.
	parsedURL, err := url.Parse(reply.Body)
	if err != nil {
		return 0, err
	}

	// Extract the repository ID from the query parameters of the URL.
	queryParams := parsedURL.Query()
	repositoryId := queryParams.Get("repositoryId")
	if repositoryId == "" {
		return 0, fmt.Errorf("repositoryId not found")
	} else {
		// Convert the repository ID from string to int.
		value, err := strconv.Atoi(repositoryId)
		if err != nil {
			return 0, err
		} else {
			return value, nil
		}
	}
}

func (service *RepositoryService) CreateProject(create CreateProjectRepository) (int, error) {
	// Send a POST request to create a repository with the specified details.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/rest/api/latest/import/repository",
		// The payload is formatted in YAML and specifies various repository properties.
		Payload: &XYamlPayload{
			Data: fmt.Sprintf(`
!!com.atlassian.bamboo.specs.util.BambooSpecProperties
rootEntity: !!com.atlassian.bamboo.specs.model.repository.bitbucket.server.BitbucketServerRepositoryProperties
  project:
    key: %s
  name: %s
  projectKey: %s
  repositorySlug: %s
  branch: %s
  server:
    id: %s
    name: %s
  sshCloneUrl: %s
  useLfs: false
specModelVersion: 9.3.0
`,
				create.Project,
				create.Name,
				create.ProjectKey,
				create.RepositorySlug,
				create.RepositoryBranch,
				create.ServerId,
				create.ServerName,
				create.CloneUrl,
			),
		},
	}, 200)

	// Handle any errors that occur during the API call.
	if err != nil {
		return 0, err
	}

	// Parse the URL from the response to extract the repository ID.
	parsedURL, err := url.Parse(reply.Body)
	if err != nil {
		return 0, err
	}

	// Extract the repository ID from the query parameters of the URL.
	queryParams := parsedURL.Query()
	repositoryId := queryParams.Get("repositoryId")
	if repositoryId == "" {
		return 0, fmt.Errorf("repositoryId not found")
	} else {
		// Convert the repository ID from string to int.
		value, err := strconv.Atoi(repositoryId)
		if err != nil {
			return 0, err
		} else {
			return value, nil
		}
	}
}

// Read searches for and retrieves a repository by its name.
// It returns the found repository or nil if no repository is found.
func (service *RepositoryService) Read(name string) (*Repository, error) {
	// Send a GET request to search for a repository by name.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/repository?max-result=1000&searchTerm=%s", name),
	}, 200)
	if err != nil {
		return nil, err
	}

	// Parse the response into a RepositoryList struct.
	response := RepositoryList{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	// Iterate over the results to find the repository with the matching name.
	for _, repository := range response.Results {
		if repository.Name == name {
			return &repository, nil
		}
	}

	// Return nil if no matching repository is found.
	return nil, nil
}

func (service *RepositoryService) ReadProject(project string, name string) (*Repository, error) {
	// Send a GET request to search for a repository by name.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/project/%s/repositories?max-result=1000&filter=%s", project, name),
	}, 200)
	if err != nil {
		return nil, err
	}

	// Parse the response into a RepositoryList struct.
	response := ProjectRepositoryList{}
	err = reply.Object(&response)
	if err != nil {
		return nil, err
	}

	// Iterate over the results to find the repository with the matching name.
	for _, repository := range response.Results {
		if repository.Name == name {
			return &repository, nil
		}
	}

	// Return nil if no matching repository is found.
	return nil, nil
}

// EnableCI enables or disables Continuous Integration (CI) for a specified repository.
func (service *RepositoryService) EnableCI(repositoryId int, enableCi bool) error {
	// Send a PUT request to update the repository's CI settings.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPut,
		Url:    fmt.Sprintf("/rest/api/latest/repository/%d/enableCi", repositoryId),
		// Payload specifies whether to enable or disable CI.
		Payload: transport.JsonPayloadData{
			Payload: map[string]bool{
				"enable": enableCi,
			},
		},
	}, 200)
	return err
}

func (service *RepositoryService) ScanCI(repositoryId int) error {
	// Send a PUT request to update the repository's CI settings.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/api/latest/repository/%d/scanNow", repositoryId),
	}, 204)
	return err
}

func (service *RepositoryService) Update(repositoryId int, request CreateRepository) error {
	// Send a PUT request to update the repository's CI settings.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/api/latest/export/repository/id/%d", repositoryId),
	}, 200)
	if err != nil {
		return err
	}

	var export []string
	err = reply.Object(&export)
	if err != nil {
		return err
	}

	if len(export) == 0 {
		return fmt.Errorf("unable to find repository with id %d", repositoryId)
	}
	filename := filepath.Base(export[0])
	oid := strings.TrimSuffix(filename, filepath.Ext(filename))

	_, err = service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/rest/api/latest/import/repository",
		// The payload is formatted in YAML and specifies various repository properties.
		Payload: &XYamlPayload{
			Data: fmt.Sprintf(`
!!com.atlassian.bamboo.specs.util.BambooSpecProperties
rootEntity: !!com.atlassian.bamboo.specs.model.repository.bitbucket.server.BitbucketServerRepositoryProperties
  oid:
    oid: %s
  name: %s
  projectKey: %s
  repositorySlug: %s
  branch: %s
  server:
    id: %s
    name: %s
  sshCloneUrl: %s
  useLfs: false
specModelVersion: 9.3.0
`,
				oid,
				request.Name,
				request.ProjectKey,
				request.RepositorySlug,
				request.RepositoryBranch,
				request.ServerId,
				request.ServerName,
				request.CloneUrl,
			),
		},
	}, 200)
	if err != nil {
		return err
	}

	return nil
}

func (service *RepositoryService) UpdateProject(repositoryId int, request CreateProjectRepository) error {
	// Send a PUT request to update the repository's CI settings.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/api/latest/export/repository/id/%d", repositoryId),
	}, 200)
	if err != nil {
		return err
	}

	var export []string
	err = reply.Object(&export)
	if err != nil {
		return err
	}

	if len(export) == 0 {
		return fmt.Errorf("unable to find repository with id %d", repositoryId)
	}
	filename := filepath.Base(export[0])
	oid := strings.TrimSuffix(filename, filepath.Ext(filename))

	_, err = service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/rest/api/latest/import/repository",
		// The payload is formatted in YAML and specifies various repository properties.
		Payload: &XYamlPayload{
			Data: fmt.Sprintf(`
!!com.atlassian.bamboo.specs.util.BambooSpecProperties
rootEntity: !!com.atlassian.bamboo.specs.model.repository.bitbucket.server.BitbucketServerRepositoryProperties
  oid:
    oid: %s  
  project:
    key: %s
  name: %s
  projectKey: %s
  repositorySlug: %s
  branch: %s
  server:
    id: %s
    name: %s
  sshCloneUrl: %s
  useLfs: false
specModelVersion: 9.3.0
`,
				oid,
				request.Project,
				request.Name,
				request.ProjectKey,
				request.RepositorySlug,
				request.RepositoryBranch,
				request.ServerId,
				request.ServerName,
				request.CloneUrl,
			),
		},
	}, 200)
	if err != nil {
		return err
	}

	return nil
}

func (service *RepositoryService) Delete(repositoryId int) error {
	// Send a DELETE request to remove an accessor from the repository.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/admin/deleteLinkedRepository.action",
		Payload: &XFormPayload{
			Data: fmt.Sprintf("repositoryId=%d", repositoryId),
		},
		Headers: map[string]string{
			"X-Atlassian-Token": "no-check",
		},
	}, 200, 302)

	return err
}

func (service *RepositoryService) DeleteProject(project string, repositoryId int) error {
	// Send a DELETE request to remove an accessor from the repository.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    "/project/deleteProjectRepository.action",
		Payload: &XFormPayload{
			Data: fmt.Sprintf("repositoryId=%d&projectKey=%s", repositoryId, project),
		},
		Headers: map[string]string{
			"X-Atlassian-Token": "no-check",
		},
	}, 200, 302)

	return err
}

// ReadAccessor retrieves a list of repositories that a specific repository has access to.
func (service *RepositoryService) ReadAccessor(repositoryId int) ([]Repository, error) {
	// Send a GET request to retrieve accessible repositories.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf("/rest/api/latest/repository/%d/rssrepository", repositoryId),
	}, 200)
	if err != nil {
		return nil, err
	}

	var repositories []Repository
	// Parse the response into a slice of Repository.
	err = reply.Object(&repositories)
	if err != nil {
		return nil, err
	}

	return repositories, nil
}

// AddAccessor adds an accessor (another repository) to a specific repository.
// It returns the updated repository or nil if the operation is not successful.
func (service *RepositoryService) AddAccessor(repositoryId int, accessorId int) (*Repository, error) {
	// Send a POST request to add an accessor to the repository.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    fmt.Sprintf("/rest/api/latest/repository/%d/rssrepository", repositoryId),
		// The payload contains the ID of the accessor to be added.
		Payload: transport.JsonPayloadData{
			Payload: map[string]int{
				"id": accessorId,
			},
		},
	}, 201, 500)
	if err != nil {
		return nil, err
	}

	// If the response status is 201 (Created), parse the response into a Repository.
	if reply.StatusCode == 201 {
		var repository Repository
		err = reply.Object(&repository)
		if err != nil {
			return nil, err
		}
		return &repository, nil
	} else {
		// Return nil if the status code is not 201.
		return nil, nil
	}
}

// RemoveAccessor removes an accessor (another repository) from a specific repository.
func (service *RepositoryService) RemoveAccessor(repositoryId int, accessorId int) error {
	// Send a DELETE request to remove an accessor from the repository.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf("/rest/api/latest/repository/%d/rssrepository/%d", repositoryId, accessorId),
	}, 204, 500)

	return err
}
