package bamboo

type AgentQuery struct {
	ExecutorType string `json:"executorType"`
	ExecutorId   int64  `json:"executorId"`
}

type AgentAssignmentRequest struct {
	ExecutorType   string `json:"executorType"`
	ExecutorId     int64  `json:"executorId"`
	EntityId       int64  `json:"entityId"`
	AssignmentType string `json:"assignmentType"`
}

type AgentAssignment struct {
	ExecutorType   string `json:"executorType"`
	ExecutorId     int64  `json:"executorId"`
	ExecutableId   int64  `json:"executableId"`
	ExecutableType string `json:"executableType"`
}
