package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
	"net/url"
)

type AgentAssignmentService struct {
	// Dependency for handling communication via Payload Transport (HTTP requests).
	transport transport.PayloadTransport
}

func (service *AgentAssignmentService) Read(request AgentQuery) (*[]AgentAssignment, error) {
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url: fmt.Sprintf("/rest/api/latest/agent/assignment?executorType=%s&executorId=%d",
			url.QueryEscape(request.ExecutorType),
			request.ExecutorId,
		),
		// Check for any errors. If there's an error, it's returned immediately.
	}, 200)

	if err != nil {
		return nil, err
	}

	deployment := make([]AgentAssignment, 0)
	err = reply.Object(&deployment)
	if err != nil {
		return nil, err
	}

	return &deployment, nil
}

func (service *AgentAssignmentService) Create(request AgentAssignmentRequest) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url: fmt.Sprintf("/rest/api/latest/agent/assignment?executorType=%s&executorId=%d&entityId=%d&assignmentType=%s",
			url.QueryEscape(request.ExecutorType),
			request.ExecutorId,
			request.EntityId,
			url.QueryEscape(request.AssignmentType),
		),
		// Check for any errors. If there's an error, it's returned immediately.
	}, 200)
	return err
}

func (service *AgentAssignmentService) Delete(request AgentAssignmentRequest) error {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url: fmt.Sprintf("/rest/api/latest/agent/assignment?executorType=%s&executorId=%d&entityId=%d&assignmentType=%s",
			url.QueryEscape(request.ExecutorType),
			request.ExecutorId,
			request.EntityId,
			url.QueryEscape(request.AssignmentType),
		),
	}, 204)
	return err
}
