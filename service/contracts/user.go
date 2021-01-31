package contracts

import (
	"strings"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// CreateUser defines a contract for creating a user
type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate function to validate the input contract
func (cu *CreateUser) Validate() bool {
	email := strings.TrimSpace(cu.Email)
	password := strings.TrimSpace(cu.Password)

	return email != "" && password != ""
}

// CreateUserResponse defines a response from a create user request
type CreateUserResponse struct {
	Response Response `json:"response"`
	ID       string   `json:"id"`
}

// UpdateUserDisplayName defines a contract for updating a users' display name
type UpdateUserDisplayName struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

// Validate function to validate the input contract
func (udn *UpdateUserDisplayName) Validate() bool {
	name := strings.TrimSpace(udn.Name)
	token := strings.TrimSpace(udn.Token)

	return name != "" && token != ""
}

// UpdateUserDisplayNameResponse defines a response from an update user display name request
type UpdateUserDisplayNameResponse struct {
	Response Response `json:"response"`
}

// GetUserResponse defines a response from a get user by "filter" request
type GetUserResponse struct {
	Response  Response          `json:"response"`
	FoundUser *cognito.UserType `json:"found_user"`
}

// GetUsersResponse defines a response from a get users by "filter" request
type GetUsersResponse struct {
	Response   Response
	FoundUsers []*cognito.UserType `json:"found_users"`
}

// GetUserByID defines a contract for getting a user by id
type GetUserByID struct {
	ID string `json:"id"`
}

// GetUserByEmail defines a contract for getting a user by email
type GetUserByEmail struct {
	Email string `json:"email"`
}

// GetUserByDisplayName defines a contract for getting a user by display name
type GetUserByDisplayName struct {
	Name string `json:"name"`
}

// GetUsersByStartDisplayName defines a contract for getting users by a starting string of a user's display name
type GetUsersByStartDisplayName struct {
	Substring string `json:"substring"`
	Limit     int64  `json:"limit"`
}

// GetUsersByStartEmail defines a contract for getting users by a starting string of a user's email
type GetUsersByStartEmail struct {
	Substring string `json:"substring"`
	Limit     int64  `json:"limit"`
}
