package assets

type AtlanAnnouncementType string

type AtlanIcon string

type AtlanConnectorType string

/*
 * Base class for all entities
 */

type AtlanStatus string

type Referenceable struct {
	// Type of the asset. For example Table, Column, and so on.
	TypeName string `json:"typeName"`
	// Globally unique identifier (GUID) of any object in Atlan.
	Guid string `json:"guid"`
	// Atlan user who created this asset.
	CreatedBy string `json:"createdBy"`
	// Atlan user who last updated this asset.
	UpdatedBy string `json:"updatedBy"`
	// Asset status in Atlan (active vs deleted)
	Status AtlanStatus `json:"status"`
	// All directly-assigned Atlan tags that exist on an asset, searchable by internal hashed-string ID of the Atlan tag.
	AtlanTags []AtlanTag `json:"classifications"`
	// All propagated Atlan tags that exist on an asset, searchable by internal hashed-string ID of the Atlan tag.
	PropagatedAtlanTags string `json:"classifications"`
	// All terms attached to an asset, searchable by the term's qualifiedName.
	AssignedTerms []AtlasGlossaryTerm `json:"meanings"`
	// All super types of an asset.
	SuperTypeNames string `json:"typeName"`
	// Time (in milliseconds) when the asset was created.
	CreateTime int `json:"createTime"`
	// Time (in milliseconds) when the asset was last updated.
	UpdateTime string `json:"updateTime"`
	// Unique fully-qualified name of the asset in Atlan.
	QualifiedName string `json:"qualifiedName"`
}

/*
 * Base class for all assets.
 */

type CertificateStatus string

type PopularityInsights int

type Link string

type MCIncident string

type MCMonitor string

type File string

type CustomMetadataAttributes string

type SchemaRegistrySubject string

type Metric string

type Readme string

type SodaCheck string

type SourceCostUnitType string

type StarredDetails string

type Meaning string

type Asset struct {
	Referenceable
	// List of groups who administer this asset.
	AdminGroups []string `json:"adminGroups"`
	// List of roles who administer this asset.
	AdminRoles []string `json:"adminRoles"`
	// List of users who administer this asset.
	AdminUsers []string `json:"adminUsers"`
	// Detailed message to include in the announcement on this asset.
	AnnouncementMessage string `json:"announcementMessage"`
	// Brief title for the announcement on this asset.
	AnnouncementTitle string `json:"announcementTitle"`
	// Type of announcement on this asset.
	AnnouncementType AtlanAnnouncementType `json:"announcementType"`
	// Time (epoch) at which the announcement was last updated, in milliseconds.
	AnnouncementUpdatedAt int64 `json:"announcementUpdatedAt"`
	// Name of the user who last updated the announcement.
	AnnouncementUpdatedBy string `json:"announcementUpdatedBy"`
	// Name of the account in which this asset exists in dbt.
	AssetDbtAccountName string `json:"assetDbtAccountName"`
	// Alias of this asset in dbt.
	AssetDbtAlias string `json:"assetDbtAlias"`
	// Version of the environment in which this asset is materialized in dbt.
	AssetDbtEnvironmentDbtVersion string `json:"assetDbtEnvironmentDbtVersion"`
	// Name of the environment in which this asset is materialized in dbt.
	AssetDbtEnvironmentName string `json:"assetDbtEnvironmentName"`
	// Time (epoch) at which the job that materialized this asset in dbt last ran, in milliseconds.
	AssetDbtJobLastRun int64 `json:"assetDbtJobLastRun"`
	// Path in S3 to the artifacts saved from the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunArtifactS3Path string `json:"assetDbtJobLastRunArtifactS3Path"`
	// Whether artifacts were saved from the last run of the job that materialized this asset in dbt (true) or not (false).
	AssetDbtJobLastRunArtifactsSaved bool `json:"assetDbtJobLastRunArtifactsSaved"`
	// Time (epoch) at which the job that materialized this asset in dbt was last created, in milliseconds.
	AssetDbtJobLastRunCreatedAt int64 `json:"assetDbtJobLastRunCreatedAt"`
	// Time (epoch) at which the job that materialized this asset in dbt was dequeued, in milliseconds.
	AssetDbtJobLastRunDequedAt int64 `json:"assetDbtJobLastRunDequedAt"`
	// Thread ID of the user who executed the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunExecutedByThreadId string `json:"assetDbtJobLastRunExecutedByThreadId"`
	// Branch in git from which the last run of the job that materialized this asset in dbt ran.
	AssetDbtJobLastRunGitBranch string `json:"assetDbtJobLastRunGitBranch"`
	// SHA hash in git for the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunGitSha string `json:"assetDbtJobLastRunGitSha"`
	// Whether docs were generated from the last run of the job that materialized this asset in dbt (true) or not (false).
	AssetDbtJobLastRunHasDocsGenerated bool `json:"assetDbtJobLastRunHasDocsGenerated"`
	// Whether sources were generated from the last run of the job that materialized this asset in dbt (true) or not (false).
	AssetDbtJobLastRunHasSourcesGenerated bool `json:"assetDbtJobLastRunHasSourcesGenerated"`
	// Whether notifications were sent from the last run of the job that materialized this asset in dbt (true) or not (false).
	AssetDbtJobLastRunNotificationsSent bool `json:"assetDbtJobLastRunNotificationsSent"`
	// Thread ID of the owner of the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunOwnerThreadId string `json:"assetDbtJobLastRunOwnerThreadId"`
	// Total duration the job that materialized this asset in dbt spent being queued.
	AssetDbtJobLastRunQueuedDuration string `json:"assetDbtJobLastRunQueuedDuration"`
	// Human-readable total duration of the last run of the job that materialized this asset in dbt spend being queued.
	AssetDbtJobLastRunQueuedDurationHumanized string `json:"assetDbtJobLastRunQueuedDurationHumanized"`
	// Run duration of the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunRunDuration string `json:"assetDbtJobLastRunRunDuration"`
	// Human-readable run duration of the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunRunDurationHumanized string `json:"assetDbtJobLastRunRunDurationHumanized"`
	// Time (epoch) at which the job that materialized this asset in dbt was started running, in milliseconds.
	AssetDbtJobLastRunStartedAt int64 `json:"assetDbtJobLastRunStartedAt"`
	// Status message of the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunStatusMessage string `json:"assetDbtJobLastRunStatusMessage"`
	// Total duration of the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunTotalDuration string `json:"assetDbtJobLastRunTotalDuration"`
	// Human-readable total duration of the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunTotalDurationHumanized string `json:"assetDbtJobLastRunTotalDurationHumanized"`
	// Time (epoch) at which the job that materialized this asset in dbt was last updated, in milliseconds.
	AssetDbtJobLastRunUpdatedAt int64 `json:"assetDbtJobLastRunUpdatedAt"`
	// URL of the last run of the job that materialized this asset in dbt.
	AssetDbtJobLastRunUrl string `json:"assetDbtJobLastRunUrl"`
	// Name of the job that materialized this asset in dbt.
	AssetDbtJobName string `json:"assetDbtJobName"`
	// Time (epoch) when the next run of the job that materializes this asset in dbt is scheduled.
	AssetDbtJobNextRun int64 `json:"assetDbtJobNextRun"`
	// Human-readable time when the next run of the job that materializes this asset in dbt is scheduled.
	AssetDbtJobNextRunHumanized string `json:"assetDbtJobNextRunHumanized"`
	// Schedule of the job that materialized this asset in dbt.
	AssetDbtJobSchedule string `json:"assetDbtJobSchedule"`
	// Human-readable cron schedule of the job that materialized this asset in dbt.
	AssetDbtJobScheduleCronHumanized string `json:"assetDbtJobScheduleCronHumanized"`
	// Status of the job that materialized this asset in dbt.
	AssetDbtJobStatus string `json:"assetDbtJobStatus"`
	// Metadata for this asset in dbt, specifically everything under the 'meta' key in the dbt object.
	AssetDbtMeta string `json:"assetDbtMeta"`
	// Name of the package in which this asset exists in dbt.
	AssetDbtPackageName string `json:"assetDbtPackageName"`
	// Name of the project in which this asset exists in dbt.
	AssetDbtProjectName string `json:"assetDbtProjectName"`
	// URL of the semantic layer proxy for this asset in dbt.
	AssetDbtSemanticLayerProxyUrl string `json:"assetDbtSemanticLayerProxyUrl"`
	// Freshness criteria for the source of this asset in dbt.
	AssetDbtSourceFreshnessCriteria string `json:"assetDbtSourceFreshnessCriteria"`
	// List of tags attached to this asset in dbt.
	AssetDbtTags []string `json:"assetDbtTags"`
	// All associated dbt test statuses.
	AssetDbtTestStatus string `json:"assetDbtTestStatus"`
	// Unique identifier of this asset in dbt.
	AssetDbtUniqueId string `json:"assetDbtUniqueId"`
	// Name of the icon to use for this asset.
	AssetIcon AtlanIcon `json:"assetIcon"`
	// List of Monte Carlo incident names attached to this asset.
	AssetMcIncidentNames []string `json:"assetMcIncidentNames"`
	// List of unique Monte Carlo incident names attached to this asset.
	AssetMcIncidentQualifiedNames []string `json:"assetMcIncidentQualifiedNames"`
	// List of Monte Carlo incident severities associated with this asset.
	AssetMcIncidentSeverities []string `json:"assetMcIncidentSeverities"`
	// List of Monte Carlo incident states associated with this asset.
	AssetMcIncidentStates []string `json:"assetMcIncidentStates"`
	// List of Monte Carlo incident sub-types associated with this asset.
	AssetMcIncidentSubTypes []string `json:"assetMcIncidentSubTypes"`
	// List of Monte Carlo incident types associated with this asset.
	AssetMcIncidentTypes []string `json:"assetMcIncidentTypes"`
	// Time (epoch) at which this asset was last synced from Monte Carlo.
	AssetMcLastSyncRunAt int64 `json:"assetMcLastSyncRunAt"`
	// List of Monte Carlo monitor names attached to this asset.
	AssetMcMonitorNames []string `json:"assetMcMonitorNames"`
	// List of unique Monte Carlo monitor names attached to this asset.
	AssetMcMonitorQualifiedNames []string `json:"assetMcMonitorQualifiedNames"`
	// Schedules of all associated Monte Carlo monitors.
	AssetMcMonitorScheduleTypes []string `json:"assetMcMonitorScheduleTypes"`
	// Statuses of all associated Monte Carlo monitors.
	AssetMcMonitorStatuses []string `json:"assetMcMonitorStatuses"`
	// Types of all associated Monte Carlo monitors.
	AssetMcMonitorTypes []string `json:"assetMcMonitorTypes"`
	// Number of checks done via Soda.
	AssetSodaCheckCount int64 `json:"assetSodaCheckCount"`
	// All associated Soda check statuses.
	AssetSodaCheckStatuses string `json:"assetSodaCheckStatuses"`
	// Status of data quality from Soda.
	AssetSodaDQStatus string `json:"assetSodaDQStatus"`
	// Time (epoch) at which the last scan via Soda occurred, in milliseconds.
	AssetSodaLastScanAt int64 `json:"assetSodaLastScanAt"`
	// Time (epoch) at which this asset was last synced via Soda, in milliseconds.
	AssetSodaLastSyncRunAt int64 `json:"assetSodaLastSyncRunAt"`
	// URL of the source for Soda.
	AssetSodaSourceURL string `json:"assetSodaSourceURL"`
	// List of tags attached to this asset.
	AssetTags []string `json:"assetTags"`
	// Glossary terms that are linked to this asset.
	AssignedTerms []AtlasGlossaryTerm `json:"assignedTerms"`
	// Status of this asset's certification.
	CertificateStatus CertificateStatus `json:"certificateStatus"`
	// Human-readable descriptive message used to provide further detail to certificateStatus.
	CertificateStatusMessage string `json:"certificateStatusMessage"`
	// Time (epoch) at which the certification was last updated, in milliseconds.
	CertificateUpdatedAt int64 `json:"certificateUpdatedAt"`
	// Name of the user who last updated the certification of this asset.
	CertificateUpdatedBy string `json:"certificateUpdatedBy"`
	// Simple name of the connection through which this asset is accessible.
	ConnectionName string `json:"connectionName"`
	// Unique name of the connection through which this asset is accessible.
	ConnectionQualifiedName string `json:"connectionQualifiedName"`
	// Type of the connector through which this asset is accessible.
	ConnectorType AtlanConnectorType `json:"connectorType"`
	// Unique name of this asset in dbt.
	DbtQualifiedName string `json:"dbtQualifiedName"`
	// Description of this asset, for example as crawled from a source.
	Description string `json:"description"`
	// Human-readable name of this asset used for display purposes (in user interface).
	DisplayName string `json:"displayName"`
	// List of files associated with this asset.
	Files []File `json:"files"`
	// Whether this asset has lineage (true) or not (false).
	HasLineage bool `json:"hasLineage"`
	// Whether this asset is AI-generated (true) or not (false).
	IsAIGenerated bool `json:"isAIGenerated"`
	// Whether this asset is discoverable through the UI (true) or not (false).
	IsDiscoverable bool `json:"isDiscoverable"`
	// Whether this asset can be edited in the UI (true) or not (false).
	IsEditable bool `json:"isEditable"`
	// Time (epoch) of the last operation that inserted, updated, or deleted rows, in milliseconds.
	LastRowChangedAt int64 `json:"lastRowChangedAt"`
	// Name of the last run of the crawler that last synchronized this asset.
	LastSyncRun string `json:"lastSyncRun"`
	// Time (epoch) at which this asset was last crawled, in milliseconds.
	LastSyncRunAt int64 `json:"lastSyncRunAt"`
	// Name of the crawler that last synchronized this asset.
	LastSyncWorkflowName string `json:"lastSyncWorkflowName"`
	// Links that are attached to this asset.
	Links []Link `json:"links"`
	// Monte Carlo incidents associated with this asset.
	McIncidents []MCIncident `json:"mcIncidents"`
	// Monte Carlo monitors that observe this asset.
	McMonitors []MCMonitor `json:"mcMonitors"`
	// Metrics associated with this asset.
	Metrics []Metric `json:"metrics"`
	// Name of this asset.
	Name string `json:"name"`
	// List of groups who own this asset.
	OwnerGroups []string `json:"ownerGroups"`
	// List of users who own this asset.
	OwnerUsers []string `json:"ownerUsers"`
	// Popularity score for this asset.
	PopularityScore float64 `json:"popularityScore"`
	// Qualified name of this asset.
	QualifiedName string `json:"qualifiedName"`
	// README that is linked to this asset.
	Readme Readme `json:"readme"`
	// URL for sample data for this asset.
	SampleDataUrl string `json:"sampleDataUrl"`
	// Subjects in the schema registry for this asset.
	SchemaRegistrySubjects []SchemaRegistrySubject `json:"schemaRegistrySubjects"`
	// Soda checks associated with this asset.
	SodaChecks []SodaCheck `json:"sodaChecks"`
	// Unit of measure for sourceTotalCost.
	SourceCostUnit SourceCostUnitType `json:"sourceCostUnit"`
	// Time (epoch) at which this asset was created in the source system, in milliseconds.
	SourceCreatedAt int64 `json:"sourceCreatedAt"`
	// User who created this asset in the source system.
	SourceCreatedBy string `json:"sourceCreatedBy"`
	// URL to create an embed for a resource (for example, an image of a dashboard) within Atlan.
	SourceEmbedURL string `json:"sourceEmbedURL"`
	// Timestamp of most recent read operation.
	SourceLastReadAt int64 `json:"sourceLastReadAt"`
	// Owners of this asset in the source system.
	SourceOwners string `json:"sourceOwners"`
	// Records of most expensive warehouse with extra insights.
	SourceQueryComputeCostRecords []PopularityInsights `json:"sourceQueryComputeCostRecords"`
	// Names of most expensive warehouses.
	SourceQueryComputeCosts []string `json:"sourceQueryComputeCosts"`
	// Total count of all read operations at source.
	SourceReadCount int64 `json:"sourceReadCount"`
	// Records of most expensive queries that accessed this asset.
	SourceReadExpensiveQueryRecords []PopularityInsights `json:"sourceReadExpensiveQueryRecords"`
	// Records of most popular queries that accessed this asset.
	SourceReadPopularQueryRecords []PopularityInsights `json:"sourceReadPopularQueryRecords"`
	// Total cost of read queries at source.
	SourceReadQueryCost float64 `json:"sourceReadQueryCost"`
	// Records of most recent users who read this asset.
	SourceReadRecentUserRecords []PopularityInsights `json:"sourceReadRecentUserRecords"`
	// Names of most recent users who read this asset.
	SourceReadRecentUsers []string `json:"sourceReadRecentUsers"`
	// Records of slowest queries that accessed this asset.
	SourceReadSlowQueryRecords []PopularityInsights `json:"sourceReadSlowQueryRecords"`
	// Records of users who read this asset the most.
	SourceReadTopUserRecords []PopularityInsights `json:"sourceReadTopUserRecords"`
	// Names of users who read this asset the most.
	SourceReadTopUsers []string `json:"sourceReadTopUsers"`
	// Total number of unique users that read data from asset.
	SourceReadUserCount int64 `json:"sourceReadUserCount"`
	// Total cost of all operations at source.
	SourceTotalCost float64 `json:"sourceTotalCost"`
	// URL to the resource within the source application.
	SourceURL string `json:"sourceURL"`
	// Time (epoch) at which this asset was last updated in the source system, in milliseconds.
	SourceUpdatedAt int64 `json:"sourceUpdatedAt"`
	// User who last updated this asset in the source system.
	SourceUpdatedBy string `json:"sourceUpdatedBy"`
	// Users who have starred this asset.
	StarredBy []string `json:"starredBy"`
	// Number of users who have starred this asset.
	StarredCount int `json:"starredCount"`
	// Details of users who have starred this asset.
	StarredDetails []StarredDetails `json:"starredDetails"`
	// Subtype of this asset.
	SubType string `json:"subType"`
	// Name of the Atlan workspace in which this asset exists.
	TenantId string `json:"tenantId"`
	// Description of this asset as provided by a user.
	UserDescription string `json:"userDescription"`
	// View score for this asset.
	ViewScore float64 `json:"viewScore"`
	// List of groups who can view assets contained in a collection.
	ViewerGroups []string `json:"viewerGroups"`
	// List of users who can view assets contained in a collection.
	ViewerUsers []string `json:"viewerUsers"`
	// Internal tracking of fields that should be serialized with null values.
	NullFields []string `json:"nullFields"`
	// Atlan tags assigned to the asset.
	AtlanTags []AtlanTag `json:"atlanTags"`
	// Map of custom metadata attributes and values defined on the asset.
	CustomMetadataSets map[string]CustomMetadataAttributes `json:"customMetadataSets"`
	// Status of the asset.
	Status AtlanStatus `json:"status"`
	// User or account that created the asset.
	CreatedBy string `json:"createdBy"`
	// User or account that last updated the asset.
	UpdatedBy string `json:"updatedBy"`
	// Time (epoch) at which the asset was created, in milliseconds.
	CreateTime int64 `json:"createTime"`
	// Time (epoch) at which the asset was last updated, in milliseconds.
	UpdateTime int64 `json:"updateTime"`
	// Details on the handler used for deletion of the asset.
	DeleteHandler string `json:"deleteHandler"`
	// Names of the Atlan tags that exist on the asset.
	AtlanTagNames []string `json:"atlanTagNames"`
	// Unused.
	IsIncomplete bool `json:"isIncomplete"`
	// Names of terms that have been linked to this asset.
	MeaningNames []string `json:"meaningNames"`
	// Details of terms that have been linked to this asset.
	Meanings []Meaning `json:"meanings"`
	// Unique identifiers (GUIDs) for any background tasks that are yet to operate on this asset.
	PendingTasks []string `json:"pendingTasks"`
}

/*
	type Asset struct {
		TypeName            string     `json:"typeName"`
		Attributes          Attributes `json:"attributes"`
		Guid                string     `json:"guid"`
		Status              string     `json:"status"`
		DisplayText         string     `json:"displayText"`
		ClassificationNames []string   `json:"classificationNames"`
		MeaningNames        []string   `json:"meaningNames"`
		Meanings            []string   `json:"meanings"`
		IsIncomplete        bool       `json:"isIncomplete"`
		Labels              []string   `json:"labels"`
		CreatedBy           string     `json:"createdBy"`
		UpdatedBy           string     `json:"updatedBy"`
		CreateTime          int64      `json:"createTime"`
		UpdateTime          int64      `json:"updateTime"`
	}
*/
type Attributes struct {
	PopularityScore                       float64 `json:"popularityScore"`
	AssetMcMonitorNames                   []string
	LastSyncRunAt                         int64 `json:"lastSyncRunAt"`
	__hasLineage                          bool
	SourceQueryComputeCostRecordList      []interface{}
	AssetSodaLastSyncRunAt                int64 `json:"assetSodaLastSyncRunAt"`
	StarredCount                          int
	AdminUsers                            []interface{}
	LastRowChangedAt                      int64 `json:"lastRowChangedAt"`
	SourceReadRecentUserList              []interface{}
	AssetMcIncidentQualifiedNames         []interface{}
	AssetMcIncidentTypes                  []interface{}
	AssetSodaLastScanAt                   int64 `json:"assetSodaLastScanAt"`
	SourceUpdatedAt                       int64 `json:"sourceUpdatedAt"`
	AssetDbtJobLastRunArtifactsSaved      bool
	StarredDetailsList                    []interface{}
	IsEditable                            bool `json:"isEditable"`
	SourceReadCount                       int
	AnnouncementUpdatedAt                 int64 `json:"announcementUpdatedAt"`
	SourceCreatedAt                       int64 `json:"sourceCreatedAt"`
	AssetDbtJobLastRunDequedAt            int64 `json:"assetDbtJobLastRunDequedAt"`
	AssetDbtTags                          []interface{}
	SourceReadSlowQueryRecordList         []interface{}
	QualifiedName                         string `json:"qualifiedName"`
	SourceQueryComputeCostList            []interface{}
	AssetDbtJobLastRunNotificationsSent   bool
	AssetMcMonitorTypes                   []interface{}
	AssetSodaCheckCount                   int
	AssetMcMonitorStatuses                []interface{}
	StarredBy                             []interface{}
	SourceLastReadAt                      int64 `json:"sourceLastReadAt"`
	Name                                  string
	CertificateUpdatedAt                  int64 `json:"certificateUpdatedAt"`
	AssetMcIncidentSeverities             []interface{}
	SourceReadQueryCost                   float64 `json:"sourceReadQueryCost"`
	OwnerUsers                            []interface{}
	AssetDbtJobLastRunHasSourcesGenerated bool
	AssetMcIncidentSubTypes               []interface{}
	IsAIGenerated                         bool `json:"isAIGenerated"`
	AssetDbtJobLastRunHasDocsGenerated    bool `json:"assetDbtJobLastRunHasDocsGenerated"`
	AssetTags                             []interface{}
	AssetMcIncidentStates                 []interface{}
	AssetDbtJobLastRunUpdatedAt           int64 `json:"assetDbtJobLastRunUpdatedAt"`
	OwnerGroups                           []interface{}
	AssetMcMonitorQualifiedNames          []interface{}
	SourceReadExpensiveQueryRecordList    []interface{}
	AssetDbtJobLastRunStartedAt           int64 `json:"assetDbtJobLastRunStartedAt"`
	IsDiscoverable                        bool  `json:"isDiscoverable"`
	IsPartial                             bool  `json:"isPartial"`
	SourceReadTopUserRecordList           []interface{}
	AssetMcMonitorScheduleTypes           []interface{}
	ViewerUsers                           []interface{}
	ViewScore                             float64 `json:"viewScore"`
	SourceReadTopUserList                 []interface{}
	AssetMcIncidentNames                  []interface{}
	AdminRoles                            []interface{}
	AdminGroups                           []interface{}
	AssetDbtJobLastRunCreatedAt           int64 `json:"assetDbtJobLastRunCreatedAt"`
	AssetDbtJobNextRun                    int64 `json:"assetDbtJobNextRun"`
	SourceReadRecentUserRecordList        []interface{}
	AssetIcon                             string `json:"assetIcon"`
	SourceReadPopularQueryRecordList      []interface{}
	SourceTotalCost                       float64 `json:"sourceTotalCost"`
	AssetMcLastSyncRunAt                  int64   `json:"assetMcLastSyncRunAt"`
	SourceReadUserCount                   int
	ViewerGroups                          []interface{}
	AssetDbtJobLastRun                    int64 `json:"assetDbtJobLastRun"`
}

// AtlanTag represents a tag in Atlan.
type AtlanTag struct {
	TypeName                          string `json:"typeName"`
	EntityGuid                        string `json:"entityGuid"`
	EntityStatus                      string `json:"entityStatus"`
	Propagate                         bool   `json:"propagate"`
	RemovePropagationsOnEntityDelete  bool   `json:"removePropagationsOnEntityDelete"`
	RestrictPropagationThroughLineage bool   `json:"restrictPropagationThroughLineage"`
}

/*
func (a *Asset) UnmarshalJSON(data []byte) error {
	var temp struct {
		Entity struct {
			TypeName               string     `json:"typeName"`
			Attributes             Attributes `json:"attributes"`
			Guid                   string     `json:"guid"`
			IsIncomplete           bool       `json:"isIncomplete"`
			Status                 string     `json:"status"`
			CreatedBy              string     `json:"createdBy"`
			UpdatedBy              string     `json:"updatedBy"`
			CreateTime             int64      `json:"createTime"`
			UpdateTime             int64      `json:"updateTime"`
			Version                int        `json:"version"`
			RelationshipAttributes struct {
				SchemaRegistrySubjects []interface{} `json:"schemaRegistrySubjects"`
				McMonitors             []interface{} `json:"mcMonitors"`
				Terms                  []struct {
					Guid                   string `json:"guid"`
					TypeName               string `json:"typeName"`
					EntityStatus           string `json:"entityStatus"`
					DisplayText            string `json:"displayText"`
					RelationshipType       string `json:"relationshipType"`
					RelationshipGuid       string `json:"relationshipGuid"`
					RelationshipStatus     string `json:"relationshipStatus"`
					RelationshipAttributes struct {
						TypeName string `json:"typeName"`
					} `json:"relationshipAttributes"`
				} `json:"terms"`
				OutputPortDataProducts []interface{} `json:"outputPortDataProducts"`
				Files                  []interface{} `json:"files"`
				McIncidents            []interface{} `json:"mcIncidents"`
				Links                  []interface{} `json:"links"`
				Categories             []interface{} `json:"categories"`
				Metrics                []interface{} `json:"metrics"`
				Readme                 interface{}   `json:"readme"`
				Meanings               []interface{} `json:"meanings"`
				SodaChecks             []interface{} `json:"sodaChecks"`
			} `json:"relationshipAttributes"`
			Labels []interface{} `json:"labels"`
		} `json:"entity"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Copy fields
	a.TypeName = temp.Entity.TypeName
	//a.Attributes = temp.Entity.Attributes
	a.Guid = temp.Entity.Guid
	a.IsIncomplete = temp.Entity.IsIncomplete
	//a.Status = temp.Entity.Status
	a.CreatedBy = temp.Entity.CreatedBy
	a.UpdatedBy = temp.Entity.UpdatedBy
	a.CreateTime = temp.Entity.CreateTime
	a.UpdateTime = temp.Entity.UpdateTime

	return nil
}
*/
