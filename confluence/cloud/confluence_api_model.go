package cloud

import "github.com/yunarta/terraform-atlassian-api-client/confluence"

type getSpacesResponse struct {
	Results []confluence.Space `json:"results,omitempty"`
}

type Links struct {
	Next string `json:"next,omitempty"`
}

type spacePermissionResponse struct {
	Results []confluence.PermissionV2 `json:"results,omitempty"`
	Links   *Links                    `json:"_links,omitempty"`
}

type searchUserResponse struct {
	Results []confluence.User `json:"results,omitempty"`
}
