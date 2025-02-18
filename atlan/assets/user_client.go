package assets

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/getsynq/atlan-go/atlan/model/structs"
)

type (
	AtlanUser  structs.AtlanUser
	UserClient AtlanClient
)

type CreateUser struct {
	Email    string `json:"email"`
	RoleName string `json:"roleName"`
	RoleID   string `json:"roleId"`
}

// CreateUserRequest represents a request to create users.
type CreateUserRequest struct {
	Users []CreateUser `json:"users"`
}

// Create initializes a new AtlanUser with email and role name.
func (u *AtlanUser) Create(email, roleName string) (*AtlanUser, error) {
	return &AtlanUser{
		Email:         email,
		WorkspaceRole: roleName,
	}, nil
}

// Updater initializes an AtlanUser for modification with a GUID.
func (u *AtlanUser) Updater(guid string) (*AtlanUser, error) {
	return &AtlanUser{
		ID: guid,
	}, nil
}

func (uc *UserClient) CreateUsers(users []AtlanUser, returnInfo bool) ([]AtlanUser, error) {
	if len(users) == 0 {
		return nil, fmt.Errorf("no users provided for creation")
	}

	var cur CreateUserRequest
	for _, user := range users {
		if user.WorkspaceRole == "" || user.Email == "" {
			return nil, fmt.Errorf("email and workspace role must not be nil")
		}

		// Fetch the role ID from roleCache
		roleID, err := GetRoleIDForRoleName(user.WorkspaceRole)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch role ID for role '%s': %w", user.WorkspaceRole, err)
		}

		cur.Users = append(cur.Users, CreateUser{
			Email:    user.Email,
			RoleName: user.WorkspaceRole,
			RoleID:   roleID,
		})
	}

	_, err := DefaultAtlanClient.CallAPI(&CREATE_USERS, nil, cur)
	if err != nil {
		return nil, fmt.Errorf("failed to create users: %w", err)
	}

	// If returnInfo is true, fetch details of created users
	if returnInfo {
		emails := []string{}
		for _, user := range cur.Users {
			emails = append(emails, user.Email)
		}
		return uc.GetByEmails(emails, 20, 0)
	}

	return nil, nil
}

// Get retrieves a UserResponse which contains a list of users defined in Atlan.
func (uc *UserClient) Get(limit int, postFilter string, sort string, count bool, offset int) (*UserResponse, error) {
	if limit == 0 {
		limit = 20
	}

	request := &structs.UserRequest{
		PostFilter: &postFilter,
		Limit:      limit,
		Sort:       &sort,
		Count:      count,
		Offset:     offset,
		Columns: []string{
			"firstName",
			"lastName",
			"username",
			"id",
			"email",
			"emailVerified",
			"enabled",
			"roles",
			"defaultRoles",
			"groupCount",
			"attributes",
			"personas",
			"createdTimestamp",
			"lastLoginTime",
			"loginEvents",
			"isLocked",
			"workspaceRole",
		},
	}

	queryParams := request.QueryParams()

	rawJson, err := DefaultAtlanClient.CallAPI(&GET_USERS, queryParams, nil)
	if err != nil {
		return nil, err
	}

	var userResponse UserResponse
	err = json.Unmarshal(rawJson, &userResponse)
	if err != nil {
		return nil, err
	}

	userResponse.Client = DefaultAtlanClient
	userResponse.Endpoint = &GET_USERS
	userResponse.Criteria = request
	userResponse.Start = request.Offset
	userResponse.Size = request.Limit

	return &userResponse, nil
}

// GetAll retrieves all users defined in Atlan.
func (uc *UserClient) GetAll(limit int, offset int, sort string) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 20
	}

	if sort == "" {
		sort = "username"
	}

	userResponse, err := uc.Get(limit, "", sort, true, offset)
	if err != nil {
		return nil, err
	}

	var users []AtlanUser
	users = append(users, userResponse.Records...)

	return users, nil
}

// GetByEmail retrieves all users with email addresses that contain the provided email.
func (uc *UserClient) GetByEmail(email string, limit int, offset int) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 20
	}

	postFilter := `{"email":{"$ilike":"%` + email + `%"}}`
	userResponse, err := uc.Get(limit, postFilter, "", true, offset)
	if err != nil {
		return nil, err
	}

	return userResponse.Records, nil
}

// GetByEmails retrieves all users with email addresses that match the provided list of emails.
func (uc *UserClient) GetByEmails(emails []string, limit int, offset int) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 20
	}

	emailJSON, err := json.Marshal(emails)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal usernames: %w", err)
	}

	emailFilter := fmt.Sprintf(`{"email":{"$in":%s}}`, string(emailJSON))
	userResponse, err := uc.Get(limit, emailFilter, "", true, offset)
	if err != nil {
		return nil, err
	}

	return userResponse.Records, nil
}

// GetByUsername retrieves a user based on the username.
func (uc *UserClient) GetByUsername(username string) (*AtlanUser, error) {
	postFilter := fmt.Sprintf(`{"$and":[{"username":{"$eq":"%s"}}]}`, username)
	userResponse, err := uc.Get(5, postFilter, "", true, 0)
	if err != nil {
		return nil, err
	}

	if len(userResponse.Records) > 0 {
		return &userResponse.Records[0], nil
	}

	return nil, nil
}

// GetByUsernames retrieves users based on their usernames.
func (uc *UserClient) GetByUsernames(usernames []string, limit int, offset int) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 5
	}

	usernamesJSON, err := json.Marshal(usernames)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal usernames: %w", err)
	}

	usernameFilter := fmt.Sprintf(`{"username":{"$in":%s}}`, string(usernamesJSON))
	userResponse, err := uc.Get(limit, usernameFilter, "", true, offset)
	if err != nil {
		return nil, err
	}

	return userResponse.Records, nil
}

func (uc *UserClient) GetGroups(guid string, request *structs.GroupRequest) ([]*AtlanGroup, error) {
	// If no request is provided, initialize a default one
	if request == nil {
		request = &structs.GroupRequest{}
	}

	api := &GET_USER_GROUPS
	api.Path = fmt.Sprintf("users/%s/groups", guid)

	queryParams := request.QueryParams()

	responseData, err := DefaultAtlanClient.CallAPI(api, queryParams, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve groups for user %s: %w", guid, err)
	}

	var response GroupResponse
	if err := json.Unmarshal(responseData, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	return response.Records, nil
}

// AddToGroupsRequest represents the request payload for adding a user to groups.
type AddToGroupsRequest struct {
	Groups []string `json:"groups"`
}

// AddUserToGroups adds a user to one or more groups.
//
// Parameters:
//
// - guid: unique identifier (GUID) of the user to add into groups.
//
// - groupIDs: unique identifiers (GUIDs) of the groups to add the user into.
//
// Errors:
// - Returns an AtlanError if any API communication occurs
func (uc *UserClient) AddUserToGroups(guid string, groupIDs []string) error {
	if guid == "" {
		return fmt.Errorf("user GUID cannot be empty")
	}
	if len(groupIDs) == 0 {
		return fmt.Errorf("group IDs cannot be empty")
	}

	requestPayload := AddToGroupsRequest{
		Groups: groupIDs,
	}

	api := &ADD_USER_TO_GROUPS
	api.Path = fmt.Sprintf("users/%s/groups", guid)

	_, err := DefaultAtlanClient.CallAPI(api, nil, requestPayload)
	if err != nil {
		return fmt.Errorf("failed to add user to groups: %w", err)
	}

	return nil
}

// ChangeRoleRequest represents the request payload for changing the role of a user.
type ChangeRoleRequest struct {
	RoleID string `json:"roleId"`
}

/*
ChangeUserRole changes the role of a user.

Parameters:

- guid: Unique identifier (GUID) of the user whose role should be changed.

- roleID: Unique identifier (GUID) of the role to assign to the user.

Errors:

- Returns an error if any API communication issue occurs.
*/
func (uc *UserClient) ChangeUserRole(guid string, roleID string) error {
	if guid == "" {
		return fmt.Errorf("user GUID cannot be empty")
	}
	if roleID == "" {
		return fmt.Errorf("role ID cannot be empty")
	}

	requestPayload := ChangeRoleRequest{
		RoleID: roleID,
	}

	api := &CHANGE_USER_ROLE
	api.Path = fmt.Sprintf("users/%s/roles/update", guid)

	_, err := DefaultAtlanClient.CallAPI(api, nil, requestPayload)
	if err != nil {
		return fmt.Errorf("failed to change user role: %w", err)
	}

	return nil
}

// Client for searching

// UserResponse represents the response containing a list of users.
type UserResponse struct {
	Size         int                  // Number of users in the current response
	Start        int                  // Offset for pagination
	Endpoint     *API                 // API endpoint used for the request
	Client       *AtlanClient         // API client used for making requests
	Criteria     *structs.UserRequest // Criteria used for the request
	TotalRecord  *int                 `json:"totalRecord,omitempty"`  // Total number of users
	FilterRecord *int                 `json:"filterRecord,omitempty"` // Number of users matching the filter
	Records      []AtlanUser          `json:"records,omitempty"`      // List of user records
}

// NewUserResponse initializes a UserResponse.
func NewUserResponse(
	client *AtlanClient,
	endpoint *API,
	criteria *structs.UserRequest,
	start int,
	size int,
	totalRecord *int,
	filterRecord *int,
	records []AtlanUser,
) *UserResponse {
	return &UserResponse{
		Size:         size,
		Start:        start,
		Endpoint:     endpoint,
		Client:       client,
		Criteria:     criteria,
		TotalRecord:  totalRecord,
		FilterRecord: filterRecord,
		Records:      records,
	}
}

// CurrentPage returns the current page of user records.
func (r *UserResponse) CurrentPage() []AtlanUser {
	return r.Records
}

// NextPage retrieves the next page of user records.
func (r *UserResponse) NextPage(start *int, size *int) (bool, error) {
	if start != nil {
		r.Start = *start
	} else {
		r.Start += r.Size
	}

	if size != nil {
		r.Size = *size
	}

	return r.getNextPage()
}

// getNextPage fetches the next page of records using the API client.
func (r *UserResponse) getNextPage() (bool, error) {
	r.Criteria.Offset = r.Start
	r.Criteria.Limit = r.Size

	queryParams := r.Criteria.QueryParams()
	responseBytes, err := DefaultAtlanClient.CallAPI(r.Endpoint, queryParams, nil)
	if err != nil {
		return false, err
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(responseBytes, &responseMap); err != nil {
		return false, errors.New("failed to unmarshal API response")
	}

	records, ok := responseMap["records"].([]interface{})
	if !ok || len(records) == 0 {
		r.Records = []AtlanUser{}
		return false, nil
	}

	var users []AtlanUser
	for _, record := range records {
		user, err := ParseAtlanUser(record)
		if err != nil {
			return false, errors.New("failed to parse user record")
		}
		users = append(users, user)
	}

	r.Records = users
	return true, nil
}

// Iterator returns an iterator over the UserResponse records.
func (r *UserResponse) Iterator() <-chan AtlanUser {
	ch := make(chan AtlanUser)

	go func() {
		defer close(ch)
		for {
			for _, user := range r.CurrentPage() {
				ch <- user
			}
			ok, err := r.NextPage(nil, nil)
			if err != nil || !ok {
				break
			}
		}
	}()

	return ch
}

// ParseAtlanUser parses a user record from a generic map into an AtlanUser object.
func ParseAtlanUser(data interface{}) (AtlanUser, error) {
	var user AtlanUser

	// Convert data to map if it's not already a map
	record, ok := data.(map[string]interface{})
	if !ok {
		return user, errors.New("invalid data format, expected map[string]interface{}")
	}

	// Marshal the map back to JSON to use standard unmarshalling
	recordBytes, err := json.Marshal(record)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(recordBytes, &user); err != nil {
		return user, err
	}

	return user, nil
}
