package confluence

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
)

const (
	PrincipalUser  = "user"
	PrincipalGroup = "group"
	PrincipalRole  = "role"
)

type Subject struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"identifier,omitempty"`
}

type Principal struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`
}

const (
	OpUse             = "use"
	OpCreate          = "create"
	OpRead            = "read"
	OpUpdate          = "update"
	OpDelete          = "delete"
	OpCopy            = "copy"
	OpMove            = "move"
	OpExport          = "export"
	OpPurge           = "purge"
	OpPurgeVersion    = "purge_version"
	OpAdminister      = "administer"
	OpRestore         = "restore"
	OpCreateSpace     = "create_space"
	OpRestrictContent = "restrict_content"
	OpArchive         = "archive"
)

const (
	TargetPage       = "page"
	TargetBlogpost   = "blogpost"
	TargetComment    = "comment"
	TargetAttachment = "attachment"
	TargetSpace      = "space"
)

var operationPriorityOrder = map[string]int{
	"read_space":             1,
	"delete_space":           2, // Lets users delete their own pages, blogs, and attachments. (For comments, the "Add" permission includes "Delete Own")
	"create_page":            10,
	"delete_page":            11,
	"archive_page":           12,
	"create_blogpost":        20,
	"delete_blogpost":        21,
	"delete_comment":         30,
	"create_comment":         31,
	"create_attachment":      40,
	"delete_attachment":      41,
	"administer_space":       100,
	"restrict_content_space": 50, // Lets users add and remove restrictions directly on content. Does not affect inherited restrictions.
	"export_space":           90,
}

func SortOperation(a, b string) int {
	// compare operationPriorityOrder[a] < 	operationPriorityOrder[b]
	if operationPriorityOrder[a] < operationPriorityOrder[b] {
		return -1
	} else if operationPriorityOrder[a] > operationPriorityOrder[b] {
		return 1
	}
	return 0
}

type Operation struct {
	Key    string `json:"key,omitempty"`
	Target string `json:"target,omitempty"`
}

func (operation Operation) GetSlug() string {
	return fmt.Sprintf("%s_%s", operation.Key, operation.Target)
}

type OperationV2 struct {
	Key    string `json:"key,omitempty"`
	Target string `json:"targetType,omitempty"`
}

func (operation OperationV2) GetSlug() string {
	return fmt.Sprintf("%s_%s", operation.Key, operation.Target)
}

type AddOperation struct {
	Key    string `json:"key,omitempty"`
	Target string `json:"target,omitempty"`
}

type AddPermission struct {
	Subject   Subject      `json:"subject,omitempty"`
	Operation AddOperation `json:"operation,omitempty"`
}

type Permission struct {
	Id        int64     `json:"id,omitempty"`
	Subject   Subject   `json:"subject,omitempty"`
	Operation Operation `json:"operation,omitempty"`
}

type PermissionV2 struct {
	Id        string      `json:"id,omitempty"`
	Principal Principal   `json:"principal,omitempty"`
	Operation OperationV2 `json:"operation,omitempty"`
}

type ObjectPermissions struct {
	Groups []GroupPermissions
	Users  []UserPermissions
}

func (op ObjectPermissions) FindUser(accountId string) *UserPermissions {
	for _, user := range op.Users {
		if user.AccountId == accountId {
			return &user
		}
	}

	return nil
}

func (op ObjectPermissions) FindGroup(groupId string) *GroupPermissions {
	for _, group := range op.Groups {
		if group.AccountId == groupId {
			return &group
		}
	}

	return nil
}

type GroupPermissions struct {
	Name        string
	AccountId   string
	Permissions []string
}

func (g GroupPermissions) DeltaPermissions(newPermissions []string) (adding []string, removing []string) {
	return collections.Delta(g.Permissions, newPermissions)
}

type UserPermissions struct {
	Name        string
	AccountId   string
	Permissions []string
}

func (u UserPermissions) DeltaPermissions(newPermissions []string) (adding []string, removing []string) {
	return collections.Delta(u.Permissions, newPermissions)
}
