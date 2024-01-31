package bamboo

type AgentAssignment struct {
	ExecutorType   string `json:"executorType"`
	ExecutorId     int64  `json:"executorId"`
	EntityId       int64  `json:"entityId"`
	AssignmentType string `json:"assignmentType"`
}
