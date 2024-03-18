package client

import (
	"fmt"
	"net/http"
)

type ErrorInfo struct {
	HTTPErrorCode int
	ErrorID       string
	ErrorMessage  string
	UserAction    string
}

type AtlanError struct {
	ErrorCode     ErrorInfo
	Args          []interface{}
	OriginalError error // Error received from Atlan API

}

func (e AtlanError) Error() string {
	errorMessage := fmt.Sprintf("%s %s", e.ErrorCode.ErrorID, fmt.Sprintf(e.ErrorCode.ErrorMessage, e.Args...))
	if e.ErrorCode.UserAction != "" {
		errorMessage += "\n" + e.ErrorCode.UserAction
	}
	if e.OriginalError != nil {
		errorMessage += "\nError: " + e.OriginalError.Error()
	}
	return errorMessage
}

type ApiConnectionError struct{ AtlanError }
type NotFoundError struct{ AtlanError }
type InvalidRequestError struct{ AtlanError }
type ApiError struct{ AtlanError }
type AuthenticationError struct{ AtlanError }
type PermissionError struct{ AtlanError }
type ConflictError struct{ AtlanError }
type RateLimitError struct{ AtlanError }
type LogicError struct{ AtlanError }

type ErrorCode int

const (
	CONNECTION_ERROR ErrorCode = iota
	INVALID_REQUEST_PASSTHROUGH
	MISSING_GROUP_ID
	MISSING_USER_ID
	MISSING_TERM_GUID
	MISSING_ROLE_NAME
	MISSING_ROLE_ID
	MISSING_ATLAN_TAG_NAME
	MISSING_ATLAN_TAG_ID
	MISSING_CM_NAME
	MISSING_CM_ID
	MISSING_CM_ATTR_NAME
	MISSING_CM_ATTR_ID
	MISSING_ENUM_NAME
	NO_GRAPH_WITH_PROCESS
	UNABLE_TO_TRANSLATE_FILTERS
	UNABLE_TO_CREATE_TYPEDEF_CATEGORY
	UNABLE_TO_UPDATE_TYPEDEF_CATEGORY
	MISSING_GUID_FOR_DELETE
	MISSING_REQUIRED_UPDATE_PARAM
	JSON_ERROR
	NOTHING_TO_ENCODE
	MISSING_REQUIRED_QUERY_PARAM
	NO_CONNECTION_ADMIN
	MISSING_PERSONA_ID
	MISSING_PURPOSE_ID
	NO_ATLAN_TAG_FOR_PURPOSE
	NO_USERS_FOR_POLICY
	MISSING_GROUP_NAME
	MISSING_USER_NAME
	MISSING_USER_EMAIL
	MISSING_GROUP_ALIAS
	NOT_AGGREGATION_METRIC
	MISSING_TOKEN_ID
	MISSING_TOKEN_NAME
	INVALID_LINEAGE_DIRECTION
	INVALID_URL
	INACCESSIBLE_URL
	NO_ATLAN_CLIENT
	MISSING_REQUIRED_RELATIONSHIP_PARAM
	INVALID_QUERY
	MISSING_CREDENTIALS
	FULL_UPDATE_ONLY
	CATEGORIES_CANNOT_BE_ARCHIVED
	AUTHENTICATION_PASSTHROUGH
	NO_API_TOKEN
	EMPTY_API_TOKEN
	INVALID_API_TOKEN
	EXPIRED_API_TOKEN
	PERMISSION_PASSTHROUGH
	UNABLE_TO_IMPERSONATE
	UNABLE_TO_ESCALATE
	NOT_FOUND_PASSTHROUGH
	ASSET_NOT_FOUND_BY_GUID
	ASSET_NOT_TYPE_REQUESTED
	ASSET_NOT_FOUND_BY_QN
	ROLE_NOT_FOUND_BY_NAME
	ROLE_NOT_FOUND_BY_ID
	ATLAN_TAG_NOT_FOUND_BY_NAME
	ATLAN_TAG_NOT_FOUND_BY_ID
	CM_NOT_FOUND_BY_NAME
	CM_NOT_FOUND_BY_ID
	CM_NO_ATTRIBUTES
	CM_ATTR_NOT_FOUND_BY_NAME
	CM_ATTR_NOT_FOUND_BY_ID
	ENUM_NOT_FOUND
	ASSET_NOT_FOUND_BY_NAME
	NO_CATEGORIES
	CONNECTION_NOT_FOUND_BY_NAME
	GROUP_NOT_FOUND_BY_NAME
	GROUP_NOT_FOUND_BY_ID
	USER_NOT_FOUND_BY_NAME
	USER_NOT_FOUND_BY_ID
	USER_NOT_FOUND_BY_EMAIL
	GROUP_NOT_FOUND_BY_ALIAS
	PERSONA_NOT_FOUND_BY_NAME
	PURPOSE_NOT_FOUND_BY_NAME
	COLLECTION_NOT_FOUND_BY_NAME
	QUERY_NOT_FOUND_BY_NAME
	CONFLICT_PASSTHROUGH
	RESERVED_SERVICE_TYPE
	RATE_LIMIT_PASSTHROUGH
	ERROR_PASSTHROUGH
	DUPLICATE_CUSTOM_ATTRIBUTES
	UNABLE_TO_DESERIALIZE
	UNABLE_TO_PARSE_ORIGINAL_QUERY
	FOUND_UNEXPECTED_ASSET_TYPE
	RETRIES_INTERRUPTED
	RETRY_OVERRUN
	TYPEDEF_NOT_FOUND_BY_NAME
)

var errorCodes = map[ErrorCode]ErrorInfo{
	CONNECTION_ERROR: {
		HTTPErrorCode: -1,
		ErrorID:       "ATLAN-GO--1-001",
		ErrorMessage:  "IOException occurred during API request to Atlan.",
		UserAction:    "Please check your internet connection and try again. If this problem persists, you should check Atlan's availability via a browser, or let us know at support@atlan.com.",
	},
	INVALID_REQUEST_PASSTHROUGH: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-000",
		ErrorMessage:  "Server responded with %s: %s. Check the details of the server's message to correct your request.",
		UserAction:    "",
	},
	MISSING_GROUP_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-001",
		ErrorMessage:  "No ID was provided when attempting to retrieve or update the group.",
		UserAction:    "You must provide an ID when attempting to retrieve or update a group.",
	},
	MISSING_USER_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-002",
		ErrorMessage:  "No ID was provided when attempting to retrieve or update the user.",
		UserAction:    "You must provide an ID when attempting to retrieve or update a user.",
	},
	MISSING_TERM_GUID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-003",
		ErrorMessage:  "No GUID was specified for the term to be removed.",
		UserAction:    "You must provide the GUID of the term to be removed.",
	},
	MISSING_ROLE_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-004",
		ErrorMessage:  "No name was provided when attempting to retrieve or update the role.",
		UserAction:    "You must provide a name when attempting to retrieve or update a role.",
	},
	MISSING_ROLE_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-005",
		ErrorMessage:  "No ID was provided when attempting to retrieve a role.",
		UserAction:    "You must provide an ID of a role when attempting to retrieve one.",
	},
	MISSING_ATLAN_TAG_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-006",
		ErrorMessage:  "No name was provided when attempting to retrieve an Atlan tag.",
		UserAction:    "You must provide a name of an Atlan tag when attempting to retrieve one.",
	},
	MISSING_ATLAN_TAG_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-007",
		ErrorMessage:  "No ID was provided when attempting to retrieve an Atlan tag.",
		UserAction:    "You must provide an ID of an Atlan tag when attempting to retrieve one.",
	},
	MISSING_CM_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-008",
		ErrorMessage:  "No name was provided when attempting to retrieve a custom metadata.",
		UserAction:    "You must provide a name of a custom metadata when attempting to retrieve one.",
	},
	MISSING_CM_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-009",
		ErrorMessage:  "No ID was provided when attempting to retrieve a custom metadata.",
		UserAction:    "You must provide an ID of a custom metadata when attempting to retrieve one.",
	},
	MISSING_CM_ATTR_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-010",
		ErrorMessage:  "No name was provided when attempting to retrieve a custom metadata attribute.",
		UserAction:    "You must provide a name of a custom metadata attribute when attempting to retrieve one.",
	},
	MISSING_CM_ATTR_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-011",
		ErrorMessage:  "No ID was provided when attempting to retrieve a custom metadata attribute.",
		UserAction:    "You must provide an ID of a custom metadata attribute when attempting to retrieve one.",
	},
	MISSING_ENUM_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-012",
		ErrorMessage:  "No name was provided when attempting to retrieve an enum.",
		UserAction:    "You must provide a name of an enum when attempting to retrieve one.",
	},
	NO_GRAPH_WITH_PROCESS: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-013",
		ErrorMessage:  "Lineage was retrieved using hideProces=false. We do not provide a graph view in this case.",
		UserAction:    "Retry your request for lineage setting hideProcess=true.",
	},
	UNABLE_TO_TRANSLATE_FILTERS: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-014",
		ErrorMessage:  "Unable to translate the provided include/exclude asset filters into JSON.",
		UserAction:    "Verify the filters you provided. If the problem persists, please raise an issue on the Go SDK GitHub repository providing context in which this error occurred.",
	},
	UNABLE_TO_CREATE_TYPEDEF_CATEGORY: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-015",
		ErrorMessage:  "Unable to create new type definitions of category: %s.",
		UserAction:    "Atlan currently only allows you to create type definitions for new custom metadata, enumerations and Atlan tags.",
	},
	UNABLE_TO_UPDATE_TYPEDEF_CATEGORY: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-016",
		ErrorMessage:  "Unable to update type definitions of category: %s.",
		UserAction:    "Atlan currently only allows you to update type definitions for custom metadata, enumerations and Atlan tags.",
	},
	MISSING_GUID_FOR_DELETE: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-017",
		ErrorMessage:  "Insufficient information provided to delete assets.",
		UserAction:    "You must provide the GUID of the asset(s) to be deleted.",
	},
	MISSING_REQUIRED_UPDATE_PARAM: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-018",
		ErrorMessage:  "One or more required parameters to update %s are missing: %s.",
		UserAction:    "You must provide all of the parameters listed to update assets of this type.",
	},
	JSON_ERROR: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-019",
		ErrorMessage:  "Invalid response object from API: %s. (HTTP response code was %d). Additional details: %s.",
		UserAction:    "Atlan was unable to produce a valid response to your request. Please verify your request is valid.",
	},
	NOTHING_TO_ENCODE: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-020",
		ErrorMessage:  "Invalid null ID found for url path formatting.",
		UserAction:    "Verify the string ID argument to the API method is what you expect. It could be either the string ID itself is null or the relevant field in your Atlan object is null.",
	},
	MISSING_REQUIRED_QUERY_PARAM: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-021",
		ErrorMessage:  "One or more required parameters to query %s are missing: %s.",
		UserAction:    "You must provide all of the parameters listed to query assets of this type.",
	},
	NO_CONNECTION_ADMIN: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-022",
		ErrorMessage:  "No admin provided for the connection.",
		UserAction:    "You must specify at least one connection admin through adminRoles, adminGroups, or adminUsers to create a new connection. Without at least one admin, the connection will be inaccessible to all.",
	},
	MISSING_PERSONA_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-023",
		ErrorMessage:  "No ID was provided when attempting to update the persona.",
		UserAction:    "You must provide an ID when attempting to update a persona.",
	},
	MISSING_PURPOSE_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-024",
		ErrorMessage:  "No ID was provided when attempting to update the purpose.",
		UserAction:    "You must provide an ID when attempting to update a purpose.",
	},
	NO_ATLAN_TAG_FOR_PURPOSE: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-025",
		ErrorMessage:  "No Atlan tags provided for the purpose.",
		UserAction:    "You must specify at least one Atlan tag to create a new purpose.",
	},
	NO_USERS_FOR_POLICY: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-026",
		ErrorMessage:  "No user or group specified for the policy.",
		UserAction:    "You must specify at least one user or group to whom the policy in a purpose will be applied.",
	},
	MISSING_GROUP_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-027",
		ErrorMessage:  "No name was provided when attempting to retrieve a group.",
		UserAction:    "You must provide a name of a group when attempting to retrieve one.",
	},
	MISSING_USER_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-028",
		ErrorMessage:  "No name was provided when attempting to retrieve a user.",
		UserAction:    "You must provide a name of a user when attempting to retrieve one.",
	},
	MISSING_USER_EMAIL: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-029",
		ErrorMessage:  "No email address was provided when attempting to retrieve a user.",
		UserAction:    "You must provide an email address of a user when attempting to retrieve one.",
	},
	MISSING_GROUP_ALIAS: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-030",
		ErrorMessage:  "No alias was provided when attempting to retrieve or update the group.",
		UserAction:    "You must provide an alias when attempting to retrieve or update a group.",
	},
	NOT_AGGREGATION_METRIC: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-031",
		ErrorMessage:  "Requested extracting a metric from a non-metric aggregation result.",
		UserAction:    "You must provide an aggregation result that is a metric aggregation to extract a numeric metric.",
	},
	MISSING_TOKEN_ID: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-032",
		ErrorMessage:  "No ID was provided when attempting to update the API token.",
		UserAction:    "You must provide an ID when attempting to update an API token.",
	},
	MISSING_TOKEN_NAME: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-033",
		ErrorMessage:  "No displayName was provided when attempting to update the API token.",
		UserAction:    "You must provide a displayName for the API token when attempting to update it.",
	},
	INVALID_LINEAGE_DIRECTION: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-034",
		ErrorMessage:  "Can only request upstream or downstream lineage (not both) through the lineage list API.",
		UserAction:    "Change your provided 'direction' parameter to either upstream or downstream.",
	},
	INVALID_URL: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-035",
		ErrorMessage:  "The URL provided for uploading a file was invalid.",
		UserAction:    "Check the provided URL and attempt to upload again.",
	},
	INACCESSIBLE_URL: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-036",
		ErrorMessage:  "The URL provided could not be accessed.",
		UserAction:    "Check the provided URL and attempt to upload again.",
	},
	NO_ATLAN_CLIENT: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-037",
		ErrorMessage:  "No Atlan client has been provided.",
		UserAction:    "You must provide an Atlan client to this operation, or it has no information about which Atlan tenant to run against.",
	},
	MISSING_REQUIRED_RELATIONSHIP_PARAM: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-038",
		ErrorMessage:  "One or more required parameters to create a relationship to %s are missing: %s.",
		UserAction:    "You must provide all of the parameters listed to relate to assets of this type.",
	},
	INVALID_QUERY: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-039",
		ErrorMessage:  "Cannot create a %s query on field: %s.",
		UserAction:    "You can either try a different field, or try a different kind of query.",
	},
	MISSING_CREDENTIALS: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-040",
		ErrorMessage:  "Missing privileged credentials to impersonate users.",
		UserAction:    "You must have both CLIENT_ID and CLIENT_SECRET configured to be able to impersonate users.",
	},
	FULL_UPDATE_ONLY: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-041",
		ErrorMessage:  "Objects of type %s should only be updated in full.",
		UserAction:    "For objects of this type you should not use trimToRequired but instead update the object in full.",
	},
	CATEGORIES_CANNOT_BE_ARCHIVED: {
		HTTPErrorCode: 400,
		ErrorID:       "ATLAN-GO-400-042",
		ErrorMessage:  "Categories cannot be archived (soft-deleted): %s.",
		UserAction:    "Please use the purge operation if you wish to remove a category.",
	},
	AUTHENTICATION_PASSTHROUGH: {
		HTTPErrorCode: 401,
		ErrorID:       "ATLAN-GO-401-000",
		ErrorMessage:  "Server responded with %s: %s",
		UserAction:    "Check the details of the server's message to correct your request.",
	},
	NO_API_TOKEN: {
		HTTPErrorCode: 401,
		ErrorID:       "ATLAN-GO-401-001",
		ErrorMessage:  "No API token provided.",
		UserAction:    "Set your API token using `Atlan.setApiToken(\"<API-TOKEN>\");`. You can generate API tokens from the Atlan Admin Center. See https://ask.atlan.com/hc/en-us/articles/8312649180049 for details or contact support at https://ask.atlan.com/hc/en-us/requests/new if you have any questions.",
	},
	EMPTY_API_TOKEN: {
		HTTPErrorCode: 401,
		ErrorID:       "ATLAN-GO-401-002",
		ErrorMessage:  "Your API token is invalid, as it is an empty string.",
		UserAction:    "You can double-check your API token from the Atlan Admin Center. See https://ask.atlan.com/hc/en-us/articles/8312649180049 for details or contact support at https://ask.atlan.com/hc/en-us/requests/new if you have any questions.",
	},
	INVALID_API_TOKEN: {
		HTTPErrorCode: 401,
		ErrorID:       "ATLAN-GO-401-003",
		ErrorMessage:  "Your API token is invalid, as it contains whitespace.",
		UserAction:    "You can double-check your API token from the Atlan Admin Center. See https://ask.atlan.com/hc/en-us/articles/8312649180049 for details or contact support at https://ask.atlan.com/hc/en-us/requests/new if you have any questions.",
	},
	EXPIRED_API_TOKEN: {
		HTTPErrorCode: 401,
		ErrorID:       "ATLAN-GO-401-004",
		ErrorMessage:  "Your API token is no longer valid, it can no longer lookup base Atlan structures.",
		UserAction:    "You can double-check your API token from the Atlan Admin Center. See https://ask.atlan.com/hc/en-us/articles/8312649180049 for details or contact support at https://ask.atlan.com/hc/en-us/requests/new if you have any questions.",
	},
	PERMISSION_PASSTHROUGH: {
		HTTPErrorCode: 403,
		ErrorID:       "ATLAN-GO-403-000",
		ErrorMessage:  "Server responded with %s: %s",
		UserAction:    "Check the details of the server's message to correct your request.",
	},
	UNABLE_TO_IMPERSONATE: {
		HTTPErrorCode: 403,
		ErrorID:       "ATLAN-GO-403-001",
		ErrorMessage:  "Unable to impersonate requested user.",
		UserAction:    "Check the details of your configured privileged credentials and the user you requested to impersonate.",
	},
	UNABLE_TO_ESCALATE: {
		HTTPErrorCode: 403,
		ErrorID:       "ATLAN-GO-403-002",
		ErrorMessage:  "Unable to escalate to a privileged user.",
		UserAction:    "Check the details of your configured privileged credentials.",
	},
	NOT_FOUND_PASSTHROUGH: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-000",
		ErrorMessage:  "Server responded with %s: %s",
		UserAction:    "Check the details of the server's message to correct your request.",
	},
	ASSET_NOT_FOUND_BY_GUID: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-001",
		ErrorMessage:  "Asset with GUID %s does not exist.",
		UserAction:    "Verify the GUID of the asset you are trying to retrieve.",
	},
	ASSET_NOT_TYPE_REQUESTED: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-002",
		ErrorMessage:  "Asset with GUID %s is not of the type requested: %s.",
		UserAction:    "Verify the GUID and expected type of the asset you are trying to retrieve.",
	},
	ASSET_NOT_FOUND_BY_QN: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-003",
		ErrorMessage:  "Asset with qualifiedName %s of type %s does not exist.",
		UserAction:    "Verify the qualifiedName and expected type of the asset you are trying to retrieve.",
	},
	ROLE_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-004",
		ErrorMessage:  "Role with name %s does not exist.",
		UserAction:    "Verify the role name provided is a valid role name.",
	},
	ROLE_NOT_FOUND_BY_ID: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-005",
		ErrorMessage:  "Role with GUID %s does not exist.",
		UserAction:    "Verify the role GUID provided is a valid role GUID.",
	},
	ATLAN_TAG_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-006",
		ErrorMessage:  "Atlan tag with name %s does not exist.",
		UserAction:    "Verify the Atlan tag name provided is a valid Atlan tag name. This should be the human-readable name of the Atlan tag.",
	},
	ATLAN_TAG_NOT_FOUND_BY_ID: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-007",
		ErrorMessage:  "Atlan tag with ID %s does not exist.",
		UserAction:    "Verify the Atlan tag ID provided is a valid Atlan tag ID. This should be the Atlan-internal hashed string representation.",
	},
	CM_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-008",
		ErrorMessage:  "Custom metadata with name %s does not exist.",
		UserAction:    "Verify the custom metadata name provided is a valid custom metadata name. This should be the human-readable name of the custom metadata.",
	},
	CM_NOT_FOUND_BY_ID: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-009",
		ErrorMessage:  "Custom metadata with ID %s does not exist.",
		UserAction:    "Verify the custom metadata ID provided is a valid custom metadata ID. This should be the Atlan-internal hashed string representation.",
	},
	CM_NO_ATTRIBUTES: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-010",
		ErrorMessage:  "Custom metadata with ID %s does not have any attributes.",
		UserAction:    "Verify the custom metadata ID you are accessing has attributes defined before attempting to retrieve one of them.",
	},
	CM_ATTR_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-011",
		ErrorMessage:  "Custom metadata property with name %s does not exist in custom metadata %s.",
		UserAction:    "Verify the custom metadata ID you are accessing has the attribute you are looking for defined. The name of the attribute should be the human-readable name.",
	},
	CM_ATTR_NOT_FOUND_BY_ID: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-012",
		ErrorMessage:  "Custom metadata property with ID %s does not exist in custom metadata %s.",
		UserAction:    "Verify the custom metadata ID you are accessing has the attribute you are looking for defined. The ID of the attribute should be the Atlan-internal hashed string representation.",
	},
	ENUM_NOT_FOUND: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-013",
		ErrorMessage:  "Enumeration with name %s does not exist.",
		UserAction:    "Verify the enumeration name provided is a valid enumeration name. This should be the human-readable name of the enumeration.",
	},
	ASSET_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-014",
		ErrorMessage:  "The %s asset could not be found by name: %s.",
		UserAction:    "Verify the requested asset type and name exist in your Atlan environment.",
	},
	NO_CATEGORIES: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-015",
		ErrorMessage:  "Unable to find any categories in glossary with GUID %s and qualifiedName %s.",
		UserAction:    "Verify the requested glossary contains categories.",
	},
	CONNECTION_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-016",
		ErrorMessage:  "Unable to find a connection with the name %s of type: %s.",
		UserAction:    "Verify the requested connection exists in your Atlan environment.",
	},
	GROUP_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-017",
		ErrorMessage:  "Group with name %s does not exist.",
		UserAction:    "Verify the group name provided is a valid group name.",
	},
	GROUP_NOT_FOUND_BY_ID: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-018",
		ErrorMessage:  "Group with GUID %s does not exist.",
		UserAction:    "Verify the role GUID provided is a valid group GUID.",
	},
	USER_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-019",
		ErrorMessage:  "User with username %s does not exist.",
		UserAction:    "Verify the username provided is a valid username.",
	},
	USER_NOT_FOUND_BY_ID: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-020",
		ErrorMessage:  "User with GUID %s does not exist.",
		UserAction:    "Verify the user GUID provided is a valid user GUID.",
	},
	USER_NOT_FOUND_BY_EMAIL: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-021",
		ErrorMessage:  "User with email %s does not exist.",
		UserAction:    "Verify the user email provided is a valid user email address.",
	},
	GROUP_NOT_FOUND_BY_ALIAS: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-022",
		ErrorMessage:  "Group with alias %s does not exist.",
		UserAction:    "Verify the group alias provided is a valid group alias.",
	},
	PERSONA_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-023",
		ErrorMessage:  "Unable to find a persona with the name: %s.",
		UserAction:    "Verify the requested persona exists in your Atlan environment.",
	},
	PURPOSE_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-024",
		ErrorMessage:  "Unable to find a purpose with the name: %s.",
		UserAction:    "Verify the requested purpose exists in your Atlan environment.",
	},
	COLLECTION_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-025",
		ErrorMessage:  "Unable to find a query collection with the name: %s.",
		UserAction:    "Verify the requested query collection exists in your Atlan environment.",
	},
	QUERY_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-026",
		ErrorMessage:  "Unable to find a query with the name: %s.",
		UserAction:    "Verify the requested query exists in your Atlan environment.",
	},
	TYPEDEF_NOT_FOUND_BY_NAME: {
		HTTPErrorCode: 404,
		ErrorID:       "ATLAN-GO-404-027",
		ErrorMessage:  "Type definition with name %s does not exist.",
		UserAction:    "Verify the type definition name provided is a valid type definition name. This should be the human-readable name of the type definition.",
	},
	CONFLICT_PASSTHROUGH: {
		HTTPErrorCode: 409,
		ErrorID:       "ATLAN-GO-409-000",
		ErrorMessage:  "Server responded with %s: %s",
		UserAction:    "Check the details of the server's message to correct your request.",
	},
	RESERVED_SERVICE_TYPE: {
		HTTPErrorCode: 409,
		ErrorID:       "ATLAN-GO-409-001",
		ErrorMessage:  "Provided service type is reserved for internal Atlan use only: %s",
		UserAction:    "You cannot create, update or remove any type definitions using this service type, it is reserved for Atlan use only.",
	},
	RATE_LIMIT_PASSTHROUGH: {
		HTTPErrorCode: 429,
		ErrorID:       "ATLAN-GO-429-000",
		ErrorMessage:  "Server responded with %s: %s",
		UserAction:    "Check the details of the server's message to correct your request.",
	},
	ERROR_PASSTHROUGH: {
		HTTPErrorCode: 500,
		ErrorID:       "ATLAN-GO-500-000",
		ErrorMessage:  "Server responded with %s: %s",
		UserAction:    "Check the details of the server's message to correct your request.",
	},
	DUPLICATE_CUSTOM_ATTRIBUTES: {
		HTTPErrorCode: 500,
		ErrorID:       "ATLAN-GO-500-001",
		ErrorMessage:  "Multiple custom attributes with exactly the same name (%s) were found for: %s.",
		UserAction:    "Please raise an issue on the Go SDK GitHub repository providing context in which this error occurred.",
	},
	UNABLE_TO_DESERIALIZE: {
		HTTPErrorCode: 500,
		ErrorID:       "ATLAN-GO-500-002",
		ErrorMessage:  "Unable to deserialize value: %s",
		UserAction:    "Please raise an issue on the Go SDK GitHub repository providing context in which this error occurred.",
	},
	UNABLE_TO_PARSE_ORIGINAL_QUERY: {
		HTTPErrorCode: 500,
		ErrorID:       "ATLAN-GO-500-003",
		ErrorMessage:  "Unable to parse the original query: %s",
		UserAction:    "Please raise an issue on the Go SDK GitHub repository providing context in which this error occurred.",
	},
	FOUND_UNEXPECTED_ASSET_TYPE: {
		HTTPErrorCode: 500,
		ErrorID:       "ATLAN-GO-500-004",
		ErrorMessage:  "Found an asset type that does not match what was requested: %s",
		UserAction:    "Please raise an issue on the Go SDK GitHub repository providing context in which this error occurred.",
	},
	RETRIES_INTERRUPTED: {
		HTTPErrorCode: 500,
		ErrorID:       "ATLAN-GO-500-005",
		ErrorMessage:  "Loop for retrying a failed action was interrupted.",
		UserAction:    "Please raise an issue on the Go SDK GitHub repository providing context in which this error occurred.",
	},
	RETRY_OVERRUN: {
		HTTPErrorCode: 500,
		ErrorID:       "ATLAN-GO-500-006",
		ErrorMessage:  "Loop for retrying a failed action hit the maximum number of retries.",
		UserAction:    "Please raise an issue on the Go SDK GitHub repository providing context in which this error occurred.",
	},
}

func handleApiError(response *http.Response) error {
	if response == nil {
		return ApiConnectionError{AtlanError{ErrorCode: errorCodes[CONNECTION_ERROR]}}
	}
	rc := response.StatusCode

	switch rc {
	case 400:
		return InvalidRequestError{AtlanError{ErrorCode: errorCodes[INVALID_REQUEST_PASSTHROUGH]}}
	case 404:
		return NotFoundError{AtlanError{ErrorCode: errorCodes[NOT_FOUND_PASSTHROUGH]}}
	case 401:
		return AuthenticationError{AtlanError{ErrorCode: errorCodes[AUTHENTICATION_PASSTHROUGH]}}
	case 403:
		return PermissionError{AtlanError{ErrorCode: errorCodes[PERMISSION_PASSTHROUGH]}}
	case 409:
		return ConflictError{AtlanError{ErrorCode: errorCodes[CONFLICT_PASSTHROUGH]}}
	case 429:
		return RateLimitError{AtlanError{ErrorCode: errorCodes[RATE_LIMIT_PASSTHROUGH]}}
	default:
		return ApiError{AtlanError{ErrorCode: errorCodes[ERROR_PASSTHROUGH]}}
	}
	return nil
}
