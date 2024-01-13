package confluence

type ContentValue struct {
	Value string `json:"value,omitempty"`
}

type Description struct {
	Plain ContentValue `json:"plain,omitempty"`
}

type Space struct {
	Id          int64       `json:"id,omitempty"`
	Key         string      `json:"key,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description Description `json:"description,omitempty"`
}

type CreateSpace struct {
	Key         string      `json:"key,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description Description `json:"description,omitempty"`
}
type UpdateSpace struct {
	Name        string      `json:"name,omitempty"`
	Description Description `json:"description,omitempty"`
}
