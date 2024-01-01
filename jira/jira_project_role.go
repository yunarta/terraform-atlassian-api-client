package jira

import "github.com/yunarta/golang-quality-of-life-pack/collections"

type ObjectRoles struct {
	RoleMap map[string]string
	Groups  []GroupRoles
	Users   []UserRoles
}

func (r ObjectRoles) FindUser(accountId string) *UserRoles {
	for _, user := range r.Users {
		if user.AccountId == accountId {
			return &user
		}
	}

	return nil
}

func (r ObjectRoles) FindGroup(groupId string) *GroupRoles {
	for _, group := range r.Groups {
		if group.AccountId == groupId {
			return &group
		}
	}

	return nil
}

type GroupRoles struct {
	Name      string
	AccountId string
	Roles     []string
}

func (r GroupRoles) DeltaRoles(newRoles []string) (adding []string, removing []string) {
	return collections.Delta(r.Roles, newRoles)
}

type UserRoles struct {
	Name      string
	AccountId string
	Roles     []string
}

func (r UserRoles) DeltaRoles(newRoles []string) (adding []string, removing []string) {
	return collections.Delta(r.Roles, newRoles)
}

type Role struct {
	Self        string  `json:"self" structs:"self"`
	Name        string  `json:"name" structs:"name"`
	ID          int     `json:"id" structs:"id"`
	Description string  `json:"description" structs:"description"`
	Actors      []Actor `json:"actors" structs:"actors"`
}

type RoleType struct {
	Name     string
	RollLink string
	ID       string
}

// Actor represents a Jira actor
type Actor struct {
	ID          int        `json:"id" structs:"id"`
	DisplayName string     `json:"displayName" structs:"displayName"`
	Type        string     `json:"type" structs:"type"`
	AvatarURL   string     `json:"avatarUrl" structs:"avatarUrl"`
	ActorUser   ActorUser  `json:"actorUser" structs:"actorUser"`
	ActorGroup  ActorGroup `json:"actorGroup" structs:"actorGroup"`
}

// ActorUser contains the account id of the actor/user
type ActorUser struct {
	AccountID string `json:"accountId" structs:"accountId"`
}

// ActorGroup contains the account id of the actor/user
type ActorGroup struct {
	GroupId string `json:"groupId" structs:"groupId"`
}
