package bitbucket

type BranchRestrictionMatcherType struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type BranchRestrictionMatcher struct {
	Id        string                       `json:"id,omitempty"`
	DisplayId string                       `json:"displayId,omitempty"`
	Type      BranchRestrictionMatcherType `json:"type,omitempty"`
}

type BranchRestrictionScope struct {
	Type       string `json:"type,omitempty"`
	ResourceId int    `json:"resourceId,omitempty"`
}

type BranchRestriction struct {
	Id      int                      `json:"id,omitempty"`
	Matcher BranchRestrictionMatcher `json:"matcher,omitempty"`
	Scope   BranchRestrictionScope   `json:"scope,omitempty"`
	Type    string                   `json:"type,omitempty"`
	Users   []string                 `json:"users,omitempty"`
	Groups  []string                 `json:"groups,omitempty"`
}

type BranchRestrictionReply struct {
	Id      int64                    `json:"id,omitempty"`
	Matcher BranchRestrictionMatcher `json:"matcher,omitempty"`
	Scope   BranchRestrictionScope   `json:"scope,omitempty"`
	Type    string                   `json:"type,omitempty"`
	Users   []User                   `json:"users,omitempty"`
	Groups  []string                 `json:"groups,omitempty"`
}
