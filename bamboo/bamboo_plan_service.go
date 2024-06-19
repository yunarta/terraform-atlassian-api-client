package bamboo

import (
	"fmt"
	"github.com/yunarta/terraform-api-transport/transport"
	"net/http"
)

const (
	createPlanEndPoint = "/rest/api/latest/import/plan"
	planEndPoint       = "/rest/api/latest/plan/%s"
)

type PlanService struct {
	// transport member is a client responsible for sending HTTP requests and receiving HTTP responses.
	transport transport.PayloadTransport
}

func (service *PlanService) Create(request CreatePlan) (*Plan, error) {
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodPost,
		Url:    createPlanEndPoint,
		// The payload is formatted in YAML and specifies various repository properties.
		Payload: &XYamlPayload{
			Data: fmt.Sprintf(`
--- !!com.atlassian.bamboo.specs.util.BambooSpecProperties
rootEntity: !!com.atlassian.bamboo.specs.api.model.plan.PlanProperties
  key: %s
  name: %s
  project:
    key: %s
`,
				request.PlanKey,
				request.Name,
				request.ProjectKey,
			),
		},
	}, 200)

	if err != nil {
		return nil, err
	}

	return service.Read(fmt.Sprintf("%s-%s", request.ProjectKey, request.PlanKey))
}

// The ReadPlan function fetches the details of a plan given its key.
// A plan in Bamboo defines a single build, including source code repository, optional triggers, commands to execute, and test results to collect.
func (service *PlanService) Read(planKey string) (*Plan, error) {
	// We send a GET request to fetch the plan. A 200 status code signifies success.
	reply, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodGet,
		Url:    fmt.Sprintf(planEndPoint, planKey),
	}, 200)
	// If there's a communication error, we return it immediately.
	if err != nil {
		return nil, err
	}

	response := Plan{}
	// Parse the return data into a Plan struct.
	err = reply.Object(&response)
	if err != nil {
		// If parsing fails, we return the error.
		return nil, err
	}

	// If everything works as it should, we return the fetched plan.
	return &response, nil
}

func (service *PlanService) Delete(planKey string) error {
	// We send a GET request to fetch the plan. A 200 status code signifies success.
	_, err := service.transport.SendWithExpectedStatus(&transport.PayloadRequest{
		Method: http.MethodDelete,
		Url:    fmt.Sprintf(planEndPoint, planKey),
	}, 204)

	// If everything works as it should, we return the fetched plan.
	return err
}
