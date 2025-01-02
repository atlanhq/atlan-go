package assets

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

type AtlanUser structs.AtlanUser

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
func Create(email, roleName string) (*AtlanUser, error) {
	return &AtlanUser{
		Email:         email,
		WorkspaceRole: roleName,
	}, nil
}

// Updater initializes an AtlanUser for modification with a GUID.
func Updater(guid string) (*AtlanUser, error) {
	return &AtlanUser{
		ID: guid,
	}, nil
}

func (u *AtlanUser) CreateUsers(users []AtlanUser, returnInfo bool) ([]AtlanUser, error) {
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
		return u.GetByEmails(emails, 20, 0)
	}

	return nil, nil
}

// Get retrieves a UserResponse which contains a list of users defined in Atlan.
func (u *AtlanUser) Get(limit int, postFilter string, sort string, count bool, offset int) (*UserResponse, error) {
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
func (u *AtlanUser) GetAll(limit int, offset int, sort string) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 20
	}

	if sort == "" {
		sort = "username"
	}

	userResponse, err := u.Get(limit, "", sort, true, offset)
	if err != nil {
		return nil, err
	}

	var users []AtlanUser
	for _, user := range userResponse.Records {
		users = append(users, user)
	}

	return users, nil
}

// GetByEmail retrieves all users with email addresses that contain the provided email.
func (u *AtlanUser) GetByEmail(email string, limit int, offset int) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 20
	}

	postFilter := `{"email":{"$ilike":"%` + email + `%"}}`
	userResponse, err := u.Get(limit, postFilter, "", true, offset)
	if err != nil {
		return nil, err
	}

	return userResponse.Records, nil
}

// GetByEmails retrieves all users with email addresses that match the provided list of emails.
func (u *AtlanUser) GetByEmails(emails []string, limit int, offset int) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 20
	}

	emailFilter := `{"email":{"$in":` + fmt.Sprintf("%v", emails) + "}}"
	userResponse, err := u.Get(limit, emailFilter, "", true, offset)
	if err != nil {
		return nil, err
	}

	return userResponse.Records, nil
}

// GetByUsername retrieves a user based on the username.
func (u *AtlanUser) GetByUsername(username string) (*AtlanUser, error) {
	userResponse, err := u.Get(5, `{"username":"`+username+`"}`, "", true, 0)
	if err != nil {
		return nil, err
	}

	if len(userResponse.Records) > 0 {
		return &userResponse.Records[0], nil
	}

	return nil, nil
}

// GetByUsernames retrieves users based on their usernames.
func (u *AtlanUser) GetByUsernames(usernames []string, limit int, offset int) ([]AtlanUser, error) {
	if limit == 0 {
		limit = 5
	}

	usernameFilter := `{"username":{"$in":` + fmt.Sprintf("%v", usernames) + "}}"
	userResponse, err := u.Get(limit, usernameFilter, "", true, offset)
	if err != nil {
		return nil, err
	}

	return userResponse.Records, nil
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
