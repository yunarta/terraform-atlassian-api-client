package confluence

import (
	"fmt"
	"github.com/yunarta/golang-quality-of-life-pack/collections"
)

const (
	PrincipalTypeUser  = "user"
	PrincipalTypeGroup = "group"
	PrincipalTypeRole  = "role"
)

type SpacePermissionSubject struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"identifier,omitempty"`
}

type SpacePermissionPrincipal struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`
}

const (
	OperationUse             = "use"
	OperationCreate          = "create"
	OperationRead            = "read"
	OperationUpdate          = "update"
	OperationDelete          = "delete"
	OperationCopy            = "copy"
	OperationMove            = "move"
	OperationExport          = "export"
	OperationPurge           = "purge"
	OperationPurgeVersion    = "purge_version"
	OperationAdminister      = "administer"
	OperationRestore         = "restore"
	OperationCreateSpace     = "create_space"
	OperationRestrictContent = "restrict_content"
	OperationArchive         = "archive"
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

type SpacePermissionOperation struct {
	Key    string `json:"key,omitempty"`
	Target string `json:"targetType,omitempty"`
}

type SpacePermissionOperation2 struct {
	Key    string `json:"key,omitempty"`
	Target string `json:"target,omitempty"`
}

func (o SpacePermissionOperation) ToSlug() string {
	return fmt.Sprintf("%s_%s", o.Key, o.Target)
}

type SpacePermission struct {
	Id        string                   `json:"id,omitempty"`
	Principal SpacePermissionPrincipal `json:"principal,omitempty"`
	Operation SpacePermissionOperation `json:"operation,omitempty"`
}

type Links struct {
	Next string `json:"next,omitempty"`
}

type SpacePermissionResponse struct {
	Results []SpacePermission `json:"results,omitempty"`
	Links   *Links            `json:"_links,omitempty"`
}

type AddPermissionRequest struct {
	Subject   SpacePermissionSubject    `json:"subject,omitempty"`
	Operation SpacePermissionOperation2 `json:"operation,omitempty"`
}

type ObjectPermissions struct {
	Groups []GroupPermissions
	Users  []UserPermissions
}

func (r ObjectPermissions) FindUser(accountId string) *UserPermissions {
	for _, user := range r.Users {
		if user.AccountId == accountId {
			return &user
		}
	}

	return nil
}

func (r ObjectPermissions) FindGroup(groupId string) *GroupPermissions {
	for _, group := range r.Groups {
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

func (r GroupPermissions) DeltaPermissions(newPermissions []string) (adding []string, removing []string) {
	return collections.Delta(r.Permissions, newPermissions)
}

type UserPermissions struct {
	Name        string
	AccountId   string
	Permissions []string
}

func (r UserPermissions) DeltaPermissions(newPermissions []string) (adding []string, removing []string) {
	return collections.Delta(r.Permissions, newPermissions)
}
