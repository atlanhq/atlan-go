package assets

import (
	"encoding/json"
	"github.com/atlanhq/atlan-go/atlan"
)

type AtlasGlossaryAttributes AtlasGlossary

type AtlasGlossary struct {
	Relation
	Asset
	ShortDescription     *string                  `json:"shortDescription,omitempty"`
	LongDescription      *string                  `json:"longDescription,omitempty"`
	Language             *string                  `json:"language,omitempty"`
	Usage                *string                  `json:"usage,omitempty"`
	AdditionalAttributes *map[string]string       `json:"additionalAttributes,omitempty"`
	Terms                *[]AtlasGlossaryTerm     `json:"terms,omitempty"`
	Categories           *[]AtlasGlossaryCategory `json:"categories,omitempty"`
}

func (ag *AtlasGlossary) UnmarshalJSON(data []byte) error {
	// Define a temporary structure with the expected JSON structure.
	var temp struct {
		ReferredEntities map[string]interface{} `json:"referredEntities"`
		Entity           struct {
			TypeName               string            `json:"typeName"`
			AttributesJSON         json.RawMessage   `json:"attributes"`
			Guid                   string            `json:"guid"`
			IsIncomplete           bool              `json:"isIncomplete"`
			Status                 atlan.AtlanStatus `json:"status"`
			CreatedBy              string            `json:"createdBy"`
			UpdatedBy              string            `json:"updatedBy"`
			CreateTime             int64             `json:"createTime"`
			UpdateTime             int64             `json:"updateTime"`
			Version                int               `json:"version"`
			RelationshipAttributes struct {
				SchemaRegistrySubjects []SchemaRegistrySubject `json:"schemaRegistrySubjects"`
				McMonitors             []MCMonitor             `json:"mcMonitors"`
				Terms                  []AtlasGlossaryTerm     `json:"terms"`
				OutputPortDataProducts []string                `json:"outputPortDataProducts"`
				Files                  []File                  `json:"files"`
				McIncidents            []MCIncident            `json:"mcIncidents"`
				Links                  []Link                  `json:"links"`
				Categories             []AtlasGlossaryCategory `json:"categories"`
				Metrics                []Metric                `json:"metrics"`
				Readme                 []Readme                `json:"readme"`
				Meanings               []Meaning               `json:"meanings"`
				SodaChecks             []SodaCheck             `json:"sodaChecks"`
			} `json:"relationshipAttributes"`
			Labels []interface{} `json:"labels"`
		} `json:"entity"`
	}

	// Unmarshal the JSON into the temporary structure
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Map the fields from the temporary structure to your AtlasGlossary struct
	ag.TypeName = &temp.Entity.TypeName
	ag.Guid = &temp.Entity.Guid
	ag.IsIncomplete = &temp.Entity.IsIncomplete
	ag.Status = &temp.Entity.Status
	ag.CreatedBy = &temp.Entity.CreatedBy
	ag.UpdatedBy = &temp.Entity.UpdatedBy
	ag.CreateTime = &temp.Entity.CreateTime
	ag.UpdateTime = &temp.Entity.UpdateTime

	var asset AtlasGlossaryAttributes
	if err := json.Unmarshal(temp.Entity.AttributesJSON, &asset); err != nil {
		return err
	}

	// Map Asset fields
	ag.Name = asset.Name
	ag.AssetIcon = asset.AssetIcon
	ag.QualifiedName = asset.QualifiedName
	ag.ShortDescription = asset.ShortDescription
	ag.LongDescription = asset.LongDescription
	ag.Language = asset.Language
	ag.Usage = asset.Usage
	ag.AssetIcon = asset.AssetIcon

	return nil
}

func (ag *AtlasGlossary) ToJSON() ([]byte, error) {
	return json.MarshalIndent(ag, "", "  ")
}

func FromJSON(data []byte) (*AtlasGlossary, error) {
	var glossaryResponse AtlasGlossary
	//fmt.Println("data")
	//fmt.Println(string(data))
	err := json.Unmarshal(data, &glossaryResponse)

	return &glossaryResponse, err
}

// Glossary represents a glossary in Atlan.
type Glossary struct {
	TypeName     string             `json:"typeName,omitempty"`
	Attributes   GlossaryAttributes `json:"attributes,omitempty"`
	Guid         string             `json:"guid,omitempty"`
	IsIncomplete bool               `json:"isIncomplete,omitempty"`
	Status       string             `json:"status,omitempty"`
	CreatedBy    string             `json:"createdBy,omitempty"`
	UpdatedBy    string             `json:"updatedBy,omitempty"`
	CreateTime   int64              `json:"createTime,omitempty"`
	UpdateTime   int64              `json:"updateTime,omitempty"`
	Version      int                `json:"version,omitempty"`
	Terms        []struct {
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
	} `json:"terms,omitempty"`
	RelationshipAttributes struct {
		SchemaRegistrySubjects []interface{} `json:"schemaRegistrySubjects,omitempty"`
		McMonitors             []interface{} `json:"mcMonitors,omitempty"`
		OutputPortDataProducts []interface{} `json:"outputPortDataProducts,omitempty"`
		Files                  []interface{} `json:"files,omitempty"`
		McIncidents            []interface{} `json:"mcIncidents,omitempty"`
		Links                  []interface{} `json:"links,omitempty"`
		Categories             []interface{} `json:"categories,omitempty"`
		Metrics                []interface{} `json:"metrics,omitempty"`
		Readme                 interface{}   `json:"readme,omitempty"`
		Meanings               []interface{} `json:"meanings,omitempty"`
		SodaChecks             []interface{} `json:"sodaChecks,omitempty"`
	} `json:"relationshipAttributes,omitempty"`
	Labels []interface{} `json:"labels,omitempty"`
}

// GlossaryAttributes represents the attributes of a glossary in Atlan.
type GlossaryAttributes struct {
	PopularityScore                           float64       `json:"popularityScore,omitempty"`
	AssetDbtJobLastRunQueuedDuration          interface{}   `json:"assetDbtJobLastRunQueuedDuration,omitempty"`
	AssetMcMonitorNames                       []string      `json:"assetMcMonitorNames,omitempty"`
	Usage                                     interface{}   `json:"usage,omitempty"`
	HasLineage                                bool          `json:"__hasLineage,omitempty"`
	AssetDbtTestStatus                        interface{}   `json:"assetDbtTestStatus,omitempty"`
	LastSyncRun                               interface{}   `json:"lastSyncRun,omitempty"`
	AssetSodaLastSyncRunAt                    int           `json:"assetSodaLastSyncRunAt,omitempty"`
	StarredCount                              int           `json:"starredCount,omitempty"`
	LastRowChangedAt                          int           `json:"lastRowChangedAt,omitempty"`
	SourceReadRecentUserList                  []interface{} `json:"sourceReadRecentUserList,omitempty"`
	AssetMcIncidentQualifiedNames             []interface{} `json:"assetMcIncidentQualifiedNames,omitempty"`
	AssetMcIncidentTypes                      []interface{} `json:"assetMcIncidentTypes,omitempty"`
	AssetSodaLastScanAt                       int           `json:"assetSodaLastScanAt,omitempty"`
	SourceUpdatedAt                           int           `json:"sourceUpdatedAt,omitempty"`
	AssetDbtJobLastRunArtifactsSaved          bool          `json:"assetDbtJobLastRunArtifactsSaved,omitempty"`
	StarredDetailsList                        []interface{} `json:"starredDetailsList,omitempty"`
	AssetDbtJobLastRunQueuedDurationHumanized interface{}   `json:"assetDbtJobLastRunQueuedDurationHumanized,omitempty"`
	AssetDbtJobStatus                         interface{}   `json:"assetDbtJobStatus,omitempty"`
	AssetDbtJobLastRunArtifactS3Path          interface{}   `json:"assetDbtJobLastRunArtifactS3Path,omitempty"`
	CertificateStatusMessage                  interface{}   `json:"certificateStatusMessage,omitempty"`
	SourceCreatedAt                           int           `json:"sourceCreatedAt,omitempty"`
	AssetDbtJobLastRunDequedAt                int           `json:"assetDbtJobLastRunDequedAt,omitempty"`
	AssetDbtTags                              []interface{} `json:"assetDbtTags,omitempty"`
	SourceReadSlowQueryRecordList             []interface{} `json:"sourceReadSlowQueryRecordList,omitempty"`
	AssetDbtAccountName                       interface{}   `json:"assetDbtAccountName,omitempty"`
	SourceQueryComputeCostList                []interface{} `json:"sourceQueryComputeCostList,omitempty"`
	AssetDbtJobLastRunOwnerThreadId           interface{}   `json:"assetDbtJobLastRunOwnerThreadId,omitempty"`
	AssetDbtJobLastRunNotificationsSent       bool          `json:"assetDbtJobLastRunNotificationsSent,omitempty"`
	AssetDbtEnvironmentDbtVersion             interface{}   `json:"assetDbtEnvironmentDbtVersion,omitempty"`
	AssetDbtMeta                              interface{}   `json:"assetDbtMeta,omitempty"`
	AssetMcMonitorTypes                       []interface{} `json:"assetMcMonitorTypes,omitempty"`
	GlossaryType                              interface{}   `json:"glossaryType,omitempty"`
	AssetDbtJobLastRunTotalDuration           interface{}   `json:"assetDbtJobLastRunTotalDuration,omitempty"`
	AssetSodaCheckCount                       int           `json:"assetSodaCheckCount,omitempty"`
	Examples                                  []interface{} `json:"examples,omitempty"`
	SourceLastReadAt                          int           `json:"sourceLastReadAt,omitempty"`
	AssetDbtJobLastRunTotalDurationHumanized  interface{}   `json:"assetDbtJobLastRunTotalDurationHumanized,omitempty"`
	SubType                                   interface{}   `json:"subType,omitempty"`
	AssetMcIncidentSeverities                 []interface{} `json:"assetMcIncidentSeverities,omitempty"`
	ConnectionName                            interface{}   `json:"connectionName,omitempty"`
	AssetDbtSourceFreshnessCriteria           interface{}   `json:"assetDbtSourceFreshnessCriteria,omitempty"`
	Metrics                                   []interface{} `json:"metrics,omitempty"`
	AdditionalAttributes                      interface{}   `json:"additionalAttributes,omitempty"`
	AssetSodaCheckStatuses                    interface{}   `json:"assetSodaCheckStatuses,omitempty"`
	CertificateStatus                         string        `json:"certificateStatus,omitempty"`
	AssetDbtJobLastRunExecutedByThreadId      interface{}   `json:"assetDbtJobLastRunExecutedByThreadId,omitempty"`
	ReplicatedFrom                            interface{}   `json:"replicatedFrom,omitempty"`
	AssetDbtJobLastRunHasSourcesGenerated     bool          `json:"assetDbtJobLastRunHasSourcesGenerated,omitempty"`
	DisplayName                               interface{}   `json:"displayName,omitempty"`
	SourceCostUnit                            interface{}   `json:"sourceCostUnit,omitempty"`
	AssetDbtUniqueId                          interface{}   `json:"assetDbtUniqueId,omitempty"`
	AssetSodaDQStatus                         interface{}   `json:"assetSodaDQStatus,omitempty"`
	TermType                                  interface{}   `json:"termType,omitempty"`
	AssetDbtJobLastRunHasDocsGenerated        bool          `json:"assetDbtJobLastRunHasDocsGenerated,omitempty"`
	AssetTags                                 []interface{} `json:"assetTags,omitempty"`
	AssetDbtSemanticLayerProxyUrl             interface{}   `json:"assetDbtSemanticLayerProxyUrl,omitempty"`
	CertificateUpdatedBy                      string        `json:"certificateUpdatedBy,omitempty"`
	AssetMcMonitorQualifiedNames              []interface{} `json:"assetMcMonitorQualifiedNames,omitempty"`
	AssetDbtJobLastRunStartedAt               int           `json:"assetDbtJobLastRunStartedAt,omitempty"`
	AnnouncementType                          interface{}   `json:"announcementType,omitempty"`
	ViewerUsers                               []interface{} `json:"viewerUsers,omitempty"`
	ViewScore                                 float64       `json:"viewScore,omitempty"`
	SourceOwners                              interface{}   `json:"sourceOwners,omitempty"`
	UserDescription                           string        `json:"userDescription,omitempty"`
	AdminGroups                               []interface{} `json:"adminGroups,omitempty"`
	AssetSodaSourceURL                        interface{}   `json:"assetSodaSourceURL,omitempty"`
	AssetDbtJobLastRunCreatedAt               int           `json:"assetDbtJobLastRunCreatedAt,omitempty"`
	AssetDbtJobNextRun                        int           `json:"assetDbtJobNextRun,omitempty"`
	AssetCoverImage                           interface{}   `json:"assetCoverImage,omitempty"`
	Abbreviation                              interface{}   `json:"abbreviation,omitempty"`
	SourceReadPopularQueryRecordList          []interface{} `json:"sourceReadPopularQueryRecordList,omitempty"`
	SourceTotalCost                           float64       `json:"sourceTotalCost,omitempty"`
	TenantId                                  interface{}   `json:"tenantId,omitempty"`
	AnnouncementMessage                       interface{}   `json:"announcementMessage,omitempty"`
	SourceEmbedURL                            interface{}   `json:"sourceEmbedURL,omitempty"`
	AssetDbtJobLastRunUrl                     interface{}   `json:"assetDbtJobLastRunUrl,omitempty"`
	Name                                      string        `json:"name,omitempty"`
	QualifiedName                             string        `json:"qualifiedName,omitempty"`
	AssetIcon                                 string        `json:"assetIcon,omitempty"`
}

// GlossaryTerm represents a term in a glossary in Atlan.
type GlossaryTerm struct {
	TypeName               string             `json:"typeName"`
	Attributes             GlossaryAttributes `json:"attributes"`
	Guid                   string             `json:"guid"`
	IsIncomplete           bool               `json:"isIncomplete"`
	Status                 string             `json:"status"`
	CreatedBy              string             `json:"createdBy"`
	UpdatedBy              string             `json:"updatedBy"`
	CreateTime             int64              `json:"createTime"`
	UpdateTime             int64              `json:"updateTime"`
	Version                int                `json:"version"`
	ValidValuesFor         []interface{}      `json:"validValuesFor"`
	SchemaRegistrySubjects []interface{}      `json:"schemaRegistrySubjects"`
	ValidValues            []interface{}      `json:"validValues"`
	SeeAlso                []interface{}      `json:"seeAlso"`
	IsA                    []interface{}      `json:"isA"`
	Antonyms               []interface{}      `json:"antonyms"`
	AssignedEntities       []interface{}      `json:"assignedEntities"`
	McIncidents            []interface{}      `json:"mcIncidents"`
	Links                  []interface{}      `json:"links"`
	Classifies             []interface{}      `json:"classifies"`
	Categories             []interface{}      `json:"categories"`
	PreferredToTerms       []interface{}      `json:"preferredToTerms"`
	PreferredTerms         []interface{}      `json:"preferredTerms"`
	TranslationTerms       []interface{}      `json:"translationTerms"`
	Synonyms               []interface{}      `json:"synonyms"`
	ReplacedBy             []interface{}      `json:"replacedBy"`
	OutputPortDataProducts []interface{}      `json:"outputPortDataProducts"`
	Readme                 interface{}        `json:"readme"`
	ReplacementTerms       []interface{}      `json:"replacementTerms"`
	Meanings               []interface{}      `json:"meanings"`
	McMonitors             []interface{}      `json:"mcMonitors"`
	TranslatedTerms        []interface{}      `json:"translatedTerms"`
	Files                  []interface{}      `json:"files"`
	Metrics                []interface{}      `json:"metrics"`
	SodaChecks             []interface{}      `json:"sodaChecks"`
	Tags                   []AtlanTag         `json:"classifications"`
	Anchor                 struct {
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
	} `json:"anchor"`
	Labels []interface{} `json:"labels"`
}
