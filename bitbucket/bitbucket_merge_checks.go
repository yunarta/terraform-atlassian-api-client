package bitbucket

type MergeCheckDetail struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}
type MergeCheck struct {
	Details    MergeCheckDetail `json:"details"`
	Enabled    bool             `json:"enabled"`
	Configured bool             `json:"configured"`
}

type MergeChecksReply struct {
	Values []MergeCheck `json:"values"`
}

type MergeCheckSetting struct {
	RequiredCount string `json:"requiredCount"`
}
