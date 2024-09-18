package bitbucket

type DefaultReviewerId struct {
	Id string `json:"id,omitempty"`
}

type DefaultReviewerUserId struct {
	Id int64 `json:"id,omitempty"`
}

type SourceMatcher struct {
	Id   string            `json:"id,omitempty"`
	Type DefaultReviewerId `json:"type,omitempty"`
}

type TargetMatcher struct {
	Id   string            `json:"id,omitempty"`
	Type DefaultReviewerId `json:"type,omitempty"`
}

type DefaultReviewers struct {
	Id                int64         `json:"id,omitempty"`
	SourceMatcher     SourceMatcher `json:"sourceMatcher,omitempty"`
	TargetMatcher     TargetMatcher `json:"targetMatcher,omitempty"`
	Reviewers         []User        `json:"reviewers,omitempty"`
	RequiredApprovals int64         `json:"requiredApprovals,omitempty"`
}

type ReadDefaultReviewers struct {
	Id                int64         `json:"id,omitempty"`
	SourceMatcher     SourceMatcher `json:"sourceRefMatcher,omitempty"`
	TargetMatcher     TargetMatcher `json:"targetRefMatcher,omitempty"`
	Reviewers         []User        `json:"reviewers,omitempty"`
	RequiredApprovals int64         `json:"requiredApprovals,omitempty"`
}
