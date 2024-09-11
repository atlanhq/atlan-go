package assets

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const (
	// Types API
	TYPES_API            = "types/"
	TYPEDEFS_API         = TYPES_API + "typedefs/"
	TYPEDEF_BY_NAME      = TYPES_API + "typedef/name/"
	TYPEDEF_BY_GUID      = TYPES_API + "typedef/guid/"
	GET_BY_NAME_TEMPLATE = TYPES_API + "{path_type}/name/{name}"
	GET_BY_GUID_TEMPLATE = TYPES_API + "{path_type}/guid/{guid}"

	// Entities API
	ENTITY_API      = "entity/"
	ENTITY_BULK_API = "entity/bulk/"

	// Files API
	FILES_API = "files/"
)

// API defines the structure of an API call.
type API struct {
	Path     string
	Method   string
	Status   int
	Endpoint Endpoint
	Consumes string
	Produces string
}

type Endpoint struct {
	Atlas string
}

var AtlasEndpoint = Endpoint{
	Atlas: "/api/meta/",
}

var HeraclesEndpoint = Endpoint{
	Atlas: "/api/service/",
}

// API calls to various services (Atlas, Heracles etc)
var (
	GET_TYPEDEF_BY_NAME = API{
		Path:     TYPEDEF_BY_NAME,
		Method:   http.MethodGet,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	GET_TYPEDEF_BY_GUID = API{
		Path:     TYPEDEF_BY_GUID,
		Method:   http.MethodGet,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	GET_ALL_TYPE_DEFS = API{
		Path:     TYPEDEFS_API,
		Method:   http.MethodGet,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	GET_ALL_TYPE_DEF_HEADERS = API{
		Path:     TYPEDEFS_API + "headers",
		Method:   http.MethodGet,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	UPDATE_TYPE_DEFS = API{
		Path:     TYPEDEFS_API,
		Method:   http.MethodPut,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	CREATE_TYPE_DEFS = API{
		Path:     TYPEDEFS_API,
		Method:   http.MethodPut,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	DELETE_TYPE_DEFS = API{
		Path:     TYPEDEFS_API,
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		Endpoint: AtlasEndpoint,
	}

	DELETE_TYPE_DEF_BY_NAME = API{
		Path:     TYPEDEF_BY_NAME,
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		Endpoint: AtlasEndpoint,
	}

	GET_ENTITY_BY_GUID = API{
		Path:     ENTITY_API + "guid/",
		Method:   http.MethodGet,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	GET_ENTITY_BY_UNIQUE_ATTRIBUTE = API{
		Path:     ENTITY_API + "uniqueAttribute/type/",
		Method:   http.MethodGet,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	INDEX_SEARCH = API{
		Path:     "search/indexsearch/",
		Method:   http.MethodPost,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	CREATE_ENTITY = API{
		Path:     ENTITY_API,
		Method:   http.MethodPost,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	CREATE_ENTITIES = API{
		Path:     ENTITY_BULK_API,
		Method:   http.MethodPost,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	DELETE_ENTITIES_BY_GUIDS = API{
		Path:     ENTITY_BULK_API,
		Method:   http.MethodDelete,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	PRESIGNED_URL = API{
		Path:     FILES_API + "presignedUrl",
		Method:   http.MethodPost,
		Status:   http.StatusOK,
		Endpoint: HeraclesEndpoint,
	}

	UPDATE_ENTITY_BY_ATTRIBUTE = API{
		Path:     ENTITY_API + "uniqueAttribute/type/",
		Method:   http.MethodPost,
		Status:   http.StatusNoContent,
		Endpoint: AtlasEndpoint,
	}

	PARTIAL_UPDATE_ENTITY_BY_ATTRIBUTE = API{
		Path:     ENTITY_API + "uniqueAttribute/type/",
		Method:   http.MethodPut,
		Status:   http.StatusOK,
		Endpoint: AtlasEndpoint,
	}

	DELETE_ENTITY_BY_ATTRIBUTE = API{
		Path:     ENTITY_API + "uniqueAttribute/type/",
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		Endpoint: AtlasEndpoint,
	}
)

// Constants for the Atlas search DSL
const (
	// TermAttributes Constants
	CONNECTOR_NAME           = "connectorName"
	CATEGORIES               = "__categories"
	CREATE_TIME_AS_TIMESTAMP = "__timestamp"
	CREATED_BY               = "__createdBy"
	GLOSSARY                 = "__glossary"
	GUID                     = "__guid"
	HAS_LINEAGE              = "__hasLineage"
	MEANINGS                 = "__meanings"
	MODIFIED_BY              = "__modifiedBy"
	NAME                     = "name.keyword"
	OWNER_GROUPS             = "ownerGroups"
	OWNER_USERS              = "ownerUsers"
	PARENT_CATEGORY          = "__parentCategory"
	POPULARITY_SCORE         = "popularityScore"
	QUALIFIED_NAME           = "qualifiedName"
	STATE                    = "__state"
	SUPER_TYPE_NAMES         = "__superTypeNames.keyword"
	TYPE_NAME                = "__typeName.keyword"
	UPDATE_TIME_AS_TIMESTAMP = "__modificationTimestamp"
	CERTIFICATE_STATUS       = "certificateStatus"

	// TextAttributes Constants
	CLASSIFICATION_NAMES            = "__classificationNames"
	CLASSIFICATIONS_TEXT            = "__classificationsText"
	CREATE_TIME_AS_DATE             = "__timestamp.date"
	DESCRIPTION                     = "description"
	MEANINGS_TEXT                   = "__meaningsText"
	NAME_TEXT                       = "name"
	QUALIFIED_NAME_TEXT             = "qualifiedName.text"
	PROPAGATED_CLASSIFICATION_NAMES = "__propagatedClassificationNames"
	PROPAGATED_TRAIT_NAMES          = "__propagatedTraitNames"
	SUPER_TYPE_NAMES_TEXT           = "__superTypeNames"
	TRAIT_NAMES                     = "__traitNames"
	UPDATE_TIME_AS_DATE             = "__modificationTimestamp.date"
	USER_DESCRIPTION                = "userDescription"
)

// FormatPathWithParams returns a new API object with the path formatted by joining the provided parameters.
func (api *API) FormatPathWithParams(params ...string) (*API, error) {
	// Join the base path with the additional params
	requestPath, err := MultipartURLJoin(api.Path, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to join URL parts: %w", err)
	}

	// Return a new API object with the formatted path
	return &API{
		Path:     requestPath,
		Method:   api.Method,
		Status:   api.Status,
		Endpoint: api.Endpoint,
		Consumes: api.Consumes,
		Produces: api.Produces,
	}, nil
}

// MultipartURLJoin joins the base path with the provided segments.
func MultipartURLJoin(basePath string, params ...string) (string, error) {
	// Parse the base path as a URL
	u, err := url.Parse(basePath)
	if err != nil {
		return "", fmt.Errorf("invalid base path: %w", err)
	}

	// Join additional path segments
	u.Path = path.Join(u.Path, path.Join(params...))

	// Return the final formatted URL
	return u.String(), nil
}
