package bamboo

// CurrentUser represents the current authenticated user in the system.
// It includes basic details like name, full name, and email.
type CurrentUser struct {
	Name     string `json:"name,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email,omitempty"`
}

// User represents a user in the system.
// It is a simpler structure compared to CurrentUser, containing only the name.
type User struct {
	Name string `json:"name,omitempty"`
}

// UserResponse is used to encapsulate a response containing multiple users.
// It is typically used in API responses where a list of users is returned.
type UserResponse struct {
	Results []User `json:"results,omitempty"`
}

// Group represents a user group in the system.
// Groups are typically used for managing a collection of users with common characteristics or permissions.
type Group struct {
	Name string `json:"name,omitempty"`
}

// GroupResponse is used to encapsulate a response containing multiple groups.
// It is often used in API responses where a list of groups is returned.
type GroupResponse struct {
	Results []Group `json:"results,omitempty"`
}
