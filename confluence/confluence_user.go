package confluence

type User struct {
	Type        string `json:"type,omitempty"`
	AccountId   string `json:"accountId,omitempty"`
	Email       string `json:"email,omitempty"`
	PublicName  string `json:"publicName,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
