package bitbucket

type Group struct {
	Name string `json:"name,omitempty"`
}

type GroupResponse struct {
	Values []Group `json:"values,omitempty"`
}

type User struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	Active       bool   `json:"active,omitempty"`
}

type UserResponse struct {
	Values []User `json:"values,omitempty"`
}
