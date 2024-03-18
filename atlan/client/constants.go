package client

import (
	"atlan-go/atlan/model"
	"fmt"
	"net/http"
	"strings"
	"time"
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
)

// API defines the structure of an API call.
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

// API calls for Atlas
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

	DELETE_ENTITIES_BY_GUIDS = API{
		Path:     ENTITY_BULK_API,
		Method:   http.MethodDelete,
		Status:   http.StatusOK,
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
	CLASSIFICATION_NAMES                               = "__classificationNames"
	CLASSIFICATIONS_TEXT                               = "__classificationsText"
	CREATE_TIME_AS_DATE                                = "__timestamp.date"
	DESCRIPTION                                        = "description"
	MEANINGS_TEXT                                      = "__meaningsText"
	NAME_TEXT                                          = "name"
	QUALIFIED_NAME_TEXT                                = "qualifiedName.text"
	PROPAGATED_CLASSIFICATION_NAMES                    = "__propagatedClassificationNames"
	PROPAGATED_TRAIT_NAMES                             = "__propagatedTraitNames"
	SUPER_TYPE_NAMES_TEXT                              = "__superTypeNames"
	TRAIT_NAMES                                        = "__traitNames"
	UPDATE_TIME_AS_DATE                                = "__modificationTimestamp.date"
	USER_DESCRIPTION                                   = "userDescription"
	ACTIVE                          model.LiteralState = "ACTIVE"
	DELETED                         model.LiteralState = "DELETED"
	PURGED                          model.LiteralState = "PURGED"
	ASCENDING                       model.SortOrder    = "asc"
	Descending                      model.SortOrder    = "desc"
)

const (
	SINGLE model.Cardinality = "SINGLE"
	LIST   model.Cardinality = "LIST"
	SET    model.Cardinality = "SET"
)

// Constants representing tag colors
const (
	AtlanTagColorGreen  model.AtlanTagColor = "Green"
	AtlanTagColorYellow model.AtlanTagColor = "Yellow"
	AtlanTagColorRed    model.AtlanTagColor = "Red"
	AtlanTagColorGray   model.AtlanTagColor = "Gray"
)

// Constants representing type categories
const (
	AtlanTypeCategoryEnum           model.AtlanTypeCategory = "ENUM"
	AtlanTypeCategoryStruct         model.AtlanTypeCategory = "STRUCT"
	AtlanTypeCategoryClassification model.AtlanTypeCategory = "CLASSIFICATION"
	AtlanTypeCategoryEntity         model.AtlanTypeCategory = "ENTITY"
	AtlanTypeCategoryRelationship   model.AtlanTypeCategory = "RELATIONSHIP"
	AtlanTypeCategoryCustomMetadata model.AtlanTypeCategory = "BUSINESS_METADATA"
)

// AdminOperationType - Enum for admin operation types.
type AdminOperationType string

const (
	Create AdminOperationType = "CREATE"
	Update AdminOperationType = "UPDATE"
	Delete AdminOperationType = "DELETE"
	Action AdminOperationType = "ACTION"
)

// AdminResourceType - Enum for admin resource types.
type AdminResourceType string

const (
	Realm                       AdminResourceType = "REALM"
	RealmRole                   AdminResourceType = "REALM_ROLE"
	RealmRoleMapping            AdminResourceType = "REALM_ROLE_MAPPING"
	RealmScopeMapping           AdminResourceType = "REALM_SCOPE_MAPPING"
	AuthFlow                    AdminResourceType = "AUTH_FLOW"
	AuthExecutionFlow           AdminResourceType = "AUTH_EXECUTION_FLOW"
	AuthExecution               AdminResourceType = "AUTH_EXECUTION"
	AuthenticatorConfig         AdminResourceType = "AUTHENTICATOR_CONFIG"
	RequiredAction              AdminResourceType = "REQUIRED_ACTION"
	IdentityProvider            AdminResourceType = "IDENTITY_PROVIDER"
	IdentityProviderMapper      AdminResourceType = "IDENTITY_PROVIDER_MAPPER"
	ProtocolMapper              AdminResourceType = "PROTOCOL_MAPPER"
	User                        AdminResourceType = "USER"
	UserLoginFailure            AdminResourceType = "USER_LOGIN_FAILURE"
	UserSession                 AdminResourceType = "USER_SESSION"
	UserFederationProvider      AdminResourceType = "USER_FEDERATION_PROVIDER"
	UserFederationMapper        AdminResourceType = "USER_FEDERATION_MAPPER"
	Group                       AdminResourceType = "GROUP"
	GroupMembership             AdminResourceType = "GROUP_MEMBERSHIP"
	Client                      AdminResourceType = "CLIENT"
	ClientInitialAccessModel    AdminResourceType = "CLIENT_INITIAL_ACCESS_MODEL"
	ClientRole                  AdminResourceType = "CLIENT_ROLE"
	ClientRoleMapping           AdminResourceType = "CLIENT_ROLE_MAPPING"
	ClientScope                 AdminResourceType = "CLIENT_SCOPE"
	ClientScopeMapping          AdminResourceType = "CLIENT_SCOPE_MAPPING"
	ClientScopeClientMapping    AdminResourceType = "CLIENT_SCOPE_CLIENT_MAPPING"
	ClusterNode                 AdminResourceType = "CLUSTER_NODE"
	Component                   AdminResourceType = "COMPONENT"
	AuthorizationResourceServer AdminResourceType = "AUTHORIZATION_RESOURCE_SERVER"
	AuthorizationResource       AdminResourceType = "AUTHORIZATION_RESOURCE"
	AuthorizationScope          AdminResourceType = "AUTHORIZATION_SCOPE"
	AuthorizationPolicy         AdminResourceType = "AUTHORIZATION_POLICY"
	Custom                      AdminResourceType = "CUSTOM"
)

// AnnouncementType represents the type of an announcement.
type AnnouncementType string

const (
	Information AnnouncementType = "information"
	Warning     AnnouncementType = "warning"
	Issue       AnnouncementType = "issue"
)

// AssetSidebarTab represents the tabs available in the asset sidebar.
type AssetSidebarTab string

const (
	Overview         AssetSidebarTab = "overview"
	Columns          AssetSidebarTab = "Columns"
	Runs             AssetSidebarTab = "Runs"
	Tasks            AssetSidebarTab = "Tasks"
	Components       AssetSidebarTab = "Components"
	Projects         AssetSidebarTab = "Projects"
	Collections      AssetSidebarTab = "Collections"
	Usage            AssetSidebarTab = "Usage"
	Objects          AssetSidebarTab = "Objects"
	Lineage          AssetSidebarTab = "Lineage"
	Incidents        AssetSidebarTab = "Incidents"
	Fields           AssetSidebarTab = "Fields"
	Visuals          AssetSidebarTab = "Visuals"
	Visualizations   AssetSidebarTab = "Visualizations"
	SchemaObjects    AssetSidebarTab = "Schema Objects"
	Relations        AssetSidebarTab = "Relations"
	FactDimRelations AssetSidebarTab = "Fact-Dim Relations"
	Profile          AssetSidebarTab = "Profile"
	Assets           AssetSidebarTab = "Assets"
	Activity         AssetSidebarTab = "Activity"
	Schedules        AssetSidebarTab = "Schedules"
	Resources        AssetSidebarTab = "Resources"
	Queries          AssetSidebarTab = "Queries"
	Requests         AssetSidebarTab = "Requests"
	Properties       AssetSidebarTab = "Properties"
	MonteCarlo       AssetSidebarTab = "Monte Carlo"
	DbtTest          AssetSidebarTab = "dbt Test"
	Soda             AssetSidebarTab = "Soda"
)

// AtlanComparisonOperator represents comparison operators in Atlan.
type AtlanComparisonOperator string

const (
	LT          AtlanComparisonOperator = "<"
	GT          AtlanComparisonOperator = ">"
	LTE         AtlanComparisonOperator = "<="
	GTE         AtlanComparisonOperator = ">="
	EQ          AtlanComparisonOperator = "="
	NEQ         AtlanComparisonOperator = "!="
	In          AtlanComparisonOperator = "in"
	Like        AtlanComparisonOperator = "like"
	StartsWith  AtlanComparisonOperator = "startsWith"
	EndsWith    AtlanComparisonOperator = "endsWith"
	Contains    AtlanComparisonOperator = "contains"
	NotContains AtlanComparisonOperator = "not_contains"
	ContainsAny AtlanComparisonOperator = "containsAny"
	ContainsAll AtlanComparisonOperator = "containsAll"
	IsNull      AtlanComparisonOperator = "isNull"
	NotNull     AtlanComparisonOperator = "notNull"
	TimeRange   AtlanComparisonOperator = "timerange"
	NotEmpty    AtlanComparisonOperator = "notEmpty"
)

// AtlanConnectionCategory represents the category of a connection in Atlan.
type AtlanConnectionCategory string

const (
	Warehouse      AtlanConnectionCategory = "warehouse"
	BI             AtlanConnectionCategory = "bi"
	ObjectStore    AtlanConnectionCategory = "ObjectStore"
	SAAS           AtlanConnectionCategory = "SaaS"
	Lake           AtlanConnectionCategory = "lake"
	QueryEngine    AtlanConnectionCategory = "queryengine"
	ELT            AtlanConnectionCategory = "elt"
	Database       AtlanConnectionCategory = "database"
	api            AtlanConnectionCategory = "API"
	EventBus       AtlanConnectionCategory = "eventbus"
	DataQuality    AtlanConnectionCategory = "data-quality"
	SchemaRegistry AtlanConnectionCategory = "schema-registry"
)

// AtlanConnectorType represents connector types with their categories.
type AtlanConnectorType struct {
	Value    string
	Category AtlanConnectionCategory
}

// ConnectorTypes is a map of all connector types for easy lookup.
var ConnectorTypes = map[string]AtlanConnectorType{
	"snowflake": {Value: "snowflake", Category: Warehouse},
	"tableau":   {Value: "tableau", Category: BI},
	// Add other connectors here
}

// NewAtlanConnectorType creates a new AtlanConnectorType with the given value and category.
func NewAtlanConnectorType(value string, category AtlanConnectionCategory) AtlanConnectorType {
	return AtlanConnectorType{Value: value, Category: category}
}

// ToQualifiedName generates a qualified name for the AtlanConnectorType.
func (a AtlanConnectorType) ToQualifiedName() string {
	return fmt.Sprintf("default/%s/%d", a.Value, time.Now().Unix())
}

// GetConnectorTypeFromQualifiedName attempts to extract an AtlanConnectorType from a qualified name.
func GetConnectorTypeFromQualifiedName(qualifiedName string) (AtlanConnectorType, error) {
	tokens := strings.Split(qualifiedName, "/")
	if len(tokens) > 1 {
		if ct, exists := ConnectorTypes[tokens[1]]; exists {
			return ct, nil
		}
	}
	return AtlanConnectorType{}, fmt.Errorf("could not determine AtlanConnectorType from %s", qualifiedName)
}

// AtlanCustomAttributePrimitiveType simulates an enum for custom attribute primitive types.
type AtlanCustomAttributePrimitiveType string

const (
	String  AtlanCustomAttributePrimitiveType = "string"
	Integer AtlanCustomAttributePrimitiveType = "int"
	Decimal AtlanCustomAttributePrimitiveType = "float"
	Boolean AtlanCustomAttributePrimitiveType = "boolean"
	Date    AtlanCustomAttributePrimitiveType = "date"
	Options AtlanCustomAttributePrimitiveType = "enum"
	Users   AtlanCustomAttributePrimitiveType = "users"
	Groups  AtlanCustomAttributePrimitiveType = "groups"
	URL     AtlanCustomAttributePrimitiveType = "url"
	SQL     AtlanCustomAttributePrimitiveType = "SQL"
)

// AtlanDeleteType simulates an enum for delete types.
type AtlanDeleteType string

const (
	Hard  AtlanDeleteType = "HARD"
	Soft  AtlanDeleteType = "SOFT"
	Purge AtlanDeleteType = "PURGE"
)
