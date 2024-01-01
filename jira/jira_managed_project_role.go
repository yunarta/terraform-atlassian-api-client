package jira

type ObjectPermission struct {
	Groups []GroupRoles
	Users  []UserRoles
}
