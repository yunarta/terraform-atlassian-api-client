package confluence

type ContentValue struct {
	Value string `json:"value,omitempty"`
}

type SpaceDescription struct {
	Plain ContentValue `json:"plain,omitempty"`
}
type CreateSpaceRequest struct {
	Key         string           `json:"key,omitempty"`
	Name        string           `json:"name,omitempty"`
	Description SpaceDescription `json:"description,omitempty"`
}

type GetSpacesResponse struct {
	Results []Space `json:"results,omitempty"`
}

type Space struct {
	Id          int64            `json:"id,omitempty"`
	Key         string           `json:"key,omitempty"`
	Name        string           `json:"name,omitempty"`
	Description SpaceDescription `json:"description,omitempty"`
}
