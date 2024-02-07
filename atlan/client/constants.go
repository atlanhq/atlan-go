package client

import "net/http"

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
)

type API struct {
	Path     string
	Method   string
	Status   int
	Endpoint Endpoint
}

type Endpoint struct {
	Atlas string
}

var AtlasEndpoint = Endpoint{
	Atlas: "/api/meta/",
}

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
)
