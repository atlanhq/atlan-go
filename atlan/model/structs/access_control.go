package structs

import "github.com/atlanhq/atlan-go/atlan"

// AccessControl represents the attributes of the AccessControl asset.
type AccessControl struct {
	Asset
	IsAccessControlEnabled  *bool         `json:"isAccessControlEnabled,omitempty"`
	DenyCustomMetadataGuids *[]string     `json:"denyCustomMetadataGuids,omitempty"`
	DenyAssetTabs           *[]string     `json:"denyAssetTabs,omitempty"`
	DenyAssetFilters        *[]string     `json:"denyAssetFilters,omitempty"`
	ChannelLink             *string       `json:"channelLink,omitempty"`
	DenyAssetTypes          *[]string     `json:"denyAssetTypes,omitempty"`
	DenyNavigationPages     *[]string     `json:"denyNavigationPages,omitempty"`
	DefaultNavigation       *string       `json:"defaultNavigation,omitempty"`
	DisplayPreferences      *[]string     `json:"displayPreferences,omitempty"`
	Policies                *[]AuthPolicy `json:"policies,omitempty"` // Relationship
}

// AuthPolicy represents a policy with various attributes.
type AuthPolicy struct {
	Asset
	PolicyType              *atlan.AuthPolicyType               `json:"policyType,omitempty"`
	PolicyServiceName       *string                             `json:"policyServiceName,omitempty"`
	PolicyCategory          *string                             `json:"policyCategory,omitempty"`
	PolicySubCategory       *string                             `json:"policySubCategory,omitempty"`
	PolicyUsers             *[]string                           `json:"policyUsers,omitempty"`
	PolicyGroups            *[]string                           `json:"policyGroups,omitempty"`
	PolicyRoles             *[]string                           `json:"policyRoles,omitempty"`
	PolicyActions           *[]string                           `json:"policyActions,omitempty"`
	PolicyResources         *[]string                           `json:"policyResources,omitempty"`
	PolicyResourceCategory  *string                             `json:"policyResourceCategory,omitempty"`
	PolicyPriority          *int                                `json:"policyPriority,omitempty"`
	IsPolicyEnabled         *bool                               `json:"isPolicyEnabled,omitempty"`
	PolicyMaskType          *string                             `json:"policyMaskType,omitempty"`
	PolicyValiditySchedule  *[]atlan.AuthPolicyValiditySchedule `json:"policyValiditySchedule,omitempty"`
	PolicyResourceSignature *string                             `json:"policyResourceSignature,omitempty"`
	PolicyDelegateAdmin     *bool                               `json:"policyDelegateAdmin,omitempty"`
	PolicyConditions        *[]atlan.AuthPolicyCondition        `json:"policyConditions,omitempty"`
	AccessControl           *AccessControl                      `json:"accessControl,omitempty"` // Relationship
}

// Persona represents the attributes of the Persona asset.
type Persona struct {
	AccessControl
	PersonaGroups *[]string `json:"personaGroups,omitempty"`
	PersonaUsers  *[]string `json:"personaUsers,omitempty"`
	RoleId        *string   `json:"roleId,omitempty"`
}
