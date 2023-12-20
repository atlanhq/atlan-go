package model

import "encoding/json"

type Glossary struct {
	TypeName     string             `json:"typeName"`
	Attributes   GlossaryAttributes `json:"attributes"`
	Guid         string             `json:"guid"`
	IsIncomplete bool               `json:"isIncomplete"`
	Status       string             `json:"status"`
	CreatedBy    string             `json:"createdBy"`
	UpdatedBy    string             `json:"updatedBy"`
	CreateTime   int64              `json:"createTime"`
	UpdateTime   int64              `json:"updateTime"`
	Version      int                `json:"version"`
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
	} `json:"terms"`
	RelationshipAttributes struct {
		SchemaRegistrySubjects []interface{} `json:"schemaRegistrySubjects"`
		McMonitors             []interface{} `json:"mcMonitors"`
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
}

type GlossaryAttributes struct {
	PopularityScore                           float64       `json:"popularityScore"`
	AssetDbtJobLastRunQueuedDuration          interface{}   `json:"assetDbtJobLastRunQueuedDuration"`
	AssetMcMonitorNames                       []string      `json:"assetMcMonitorNames"`
	Usage                                     interface{}   `json:"usage"`
	HasLineage                                bool          `json:"__hasLineage"`
	AssetDbtTestStatus                        interface{}   `json:"assetDbtTestStatus"`
	LastSyncRun                               interface{}   `json:"lastSyncRun"`
	AssetSodaLastSyncRunAt                    int           `json:"assetSodaLastSyncRunAt"`
	StarredCount                              int           `json:"starredCount"`
	LastRowChangedAt                          int           `json:"lastRowChangedAt"`
	SourceReadRecentUserList                  []interface{} `json:"sourceReadRecentUserList"`
	AssetMcIncidentQualifiedNames             []interface{} `json:"assetMcIncidentQualifiedNames"`
	AssetMcIncidentTypes                      []interface{} `json:"assetMcIncidentTypes"`
	AssetSodaLastScanAt                       int           `json:"assetSodaLastScanAt"`
	SourceUpdatedAt                           int           `json:"sourceUpdatedAt"`
	AssetDbtJobLastRunArtifactsSaved          bool          `json:"assetDbtJobLastRunArtifactsSaved"`
	StarredDetailsList                        []interface{} `json:"starredDetailsList"`
	AssetDbtJobLastRunQueuedDurationHumanized interface{}   `json:"assetDbtJobLastRunQueuedDurationHumanized"`
	AssetDbtJobStatus                         interface{}   `json:"assetDbtJobStatus"`
	AssetDbtJobLastRunArtifactS3Path          interface{}   `json:"assetDbtJobLastRunArtifactS3Path"`
	CertificateStatusMessage                  interface{}   `json:"certificateStatusMessage"`
	SourceCreatedAt                           int           `json:"sourceCreatedAt"`
	AssetDbtJobLastRunDequedAt                int           `json:"assetDbtJobLastRunDequedAt"`
	AssetDbtTags                              []interface{} `json:"assetDbtTags"`
	SourceReadSlowQueryRecordList             []interface{} `json:"sourceReadSlowQueryRecordList"`
	AssetDbtAccountName                       interface{}   `json:"assetDbtAccountName"`
	SourceQueryComputeCostList                []interface{} `json:"sourceQueryComputeCostList"`
	AssetDbtJobLastRunOwnerThreadId           interface{}   `json:"assetDbtJobLastRunOwnerThreadId"`
	AssetDbtJobLastRunNotificationsSent       bool          `json:"assetDbtJobLastRunNotificationsSent"`
	AssetDbtEnvironmentDbtVersion             interface{}   `json:"assetDbtEnvironmentDbtVersion"`
	AssetDbtMeta                              interface{}   `json:"assetDbtMeta"`
	AssetMcMonitorTypes                       []interface{} `json:"assetMcMonitorTypes"`
	GlossaryType                              interface{}   `json:"glossaryType"`
	AssetDbtJobLastRunTotalDuration           interface{}   `json:"assetDbtJobLastRunTotalDuration"`
	AssetSodaCheckCount                       int           `json:"assetSodaCheckCount"`
	Examples                                  []interface{} `json:"examples"`
	SourceLastReadAt                          int           `json:"sourceLastReadAt"`
	AssetDbtJobLastRunTotalDurationHumanized  interface{}   `json:"assetDbtJobLastRunTotalDurationHumanized"`
	SubType                                   interface{}   `json:"subType"`
	AssetMcIncidentSeverities                 []interface{} `json:"assetMcIncidentSeverities"`
	ConnectionName                            interface{}   `json:"connectionName"`
	AssetDbtSourceFreshnessCriteria           interface{}   `json:"assetDbtSourceFreshnessCriteria"`
	Metrics                                   []interface{} `json:"metrics"`
	AdditionalAttributes                      interface{}   `json:"additionalAttributes"`
	AssetSodaCheckStatuses                    interface{}   `json:"assetSodaCheckStatuses"`
	CertificateStatus                         string        `json:"certificateStatus"`
	AssetDbtJobLastRunExecutedByThreadId      interface{}   `json:"assetDbtJobLastRunExecutedByThreadId"`
	ReplicatedFrom                            interface{}   `json:"replicatedFrom"`
	AssetDbtJobLastRunHasSourcesGenerated     bool          `json:"assetDbtJobLastRunHasSourcesGenerated"`
	DisplayName                               interface{}   `json:"displayName"`
	SourceCostUnit                            interface{}   `json:"sourceCostUnit"`
	AssetDbtUniqueId                          interface{}   `json:"assetDbtUniqueId"`
	AssetSodaDQStatus                         interface{}   `json:"assetSodaDQStatus"`
	TermType                                  interface{}   `json:"termType"`
	AssetDbtJobLastRunHasDocsGenerated        bool          `json:"assetDbtJobLastRunHasDocsGenerated"`
	AssetTags                                 []interface{} `json:"assetTags"`
	AssetDbtSemanticLayerProxyUrl             interface{}   `json:"assetDbtSemanticLayerProxyUrl"`
	CertificateUpdatedBy                      string        `json:"certificateUpdatedBy"`
	AssetMcMonitorQualifiedNames              []interface{} `json:"assetMcMonitorQualifiedNames"`
	AssetDbtJobLastRunStartedAt               int           `json:"assetDbtJobLastRunStartedAt"`
	AnnouncementType                          interface{}   `json:"announcementType"`
	ViewerUsers                               []interface{} `json:"viewerUsers"`
	ViewScore                                 float64       `json:"viewScore"`
	SourceOwners                              interface{}   `json:"sourceOwners"`
	UserDescription                           string        `json:"userDescription"`
	AdminGroups                               []interface{} `json:"adminGroups"`
	AssetSodaSourceURL                        interface{}   `json:"assetSodaSourceURL"`
	AssetDbtJobLastRunCreatedAt               int           `json:"assetDbtJobLastRunCreatedAt"`
	AssetDbtJobNextRun                        int           `json:"assetDbtJobNextRun"`
	AssetCoverImage                           interface{}   `json:"assetCoverImage"`
	Abbreviation                              interface{}   `json:"abbreviation"`
	SourceReadPopularQueryRecordList          []interface{} `json:"sourceReadPopularQueryRecordList"`
	SourceTotalCost                           float64       `json:"sourceTotalCost"`
	TenantId                                  interface{}   `json:"tenantId"`
	AnnouncementMessage                       interface{}   `json:"announcementMessage"`
	SourceEmbedURL                            interface{}   `json:"sourceEmbedURL"`
	AssetDbtJobLastRunUrl                     interface{}   `json:"assetDbtJobLastRunUrl"`
}

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
	Tags                   []struct {
		TypeName                          string `json:"typeName"`
		EntityGuid                        string `json:"entityGuid"`
		EntityStatus                      string `json:"entityStatus"`
		Propagate                         bool   `json:"propagate"`
		RemovePropagationsOnEntityDelete  bool   `json:"removePropagationsOnEntityDelete"`
		RestrictPropagationThroughLineage bool   `json:"restrictPropagationThroughLineage"`
	} `json:"classifications"`
	Anchor struct {
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

func (g *Glossary) UnmarshalJSON(data []byte) error {
	var temp struct {
		Entity struct {
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
	g.TypeName = temp.Entity.TypeName
	g.Attributes = temp.Entity.Attributes
	g.Guid = temp.Entity.Guid
	g.IsIncomplete = temp.Entity.IsIncomplete
	g.Status = temp.Entity.Status
	g.CreatedBy = temp.Entity.CreatedBy
	g.UpdatedBy = temp.Entity.UpdatedBy
	g.CreateTime = temp.Entity.CreateTime
	g.UpdateTime = temp.Entity.UpdateTime
	g.Version = temp.Entity.Version
	g.Terms = temp.Entity.RelationshipAttributes.Terms

	return nil
}

func (g *Glossary) ToJSON() ([]byte, error) {
	return json.MarshalIndent(g, "", "  ")
}

func FromJSON(data []byte) (*Glossary, error) {
	var glossaryResponse Glossary
	err := json.Unmarshal(data, &glossaryResponse)

	return &glossaryResponse, err
}

func (gt *GlossaryTerm) UnmarshalJSON(data []byte) error {
	var temp struct {
		Entity struct {
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
			RelationshipAttributes struct {
				ValidValuesFor         []interface{} `json:"validValuesFor"`
				SchemaRegistrySubjects []interface{} `json:"schemaRegistrySubjects"`
				ValidValues            []interface{} `json:"validValues"`
				SeeAlso                []interface{} `json:"seeAlso"`
				IsA                    []interface{} `json:"isA"`
				Antonyms               []interface{} `json:"antonyms"`
				AssignedEntities       []interface{} `json:"assignedEntities"`
				McIncidents            []interface{} `json:"mcIncidents"`
				Links                  []interface{} `json:"links"`
				Classifies             []interface{} `json:"classifies"`
				Categories             []interface{} `json:"categories"`
				PreferredToTerms       []interface{} `json:"preferredToTerms"`
				PreferredTerms         []interface{} `json:"preferredTerms"`
				TranslationTerms       []interface{} `json:"translationTerms"`
				Synonyms               []interface{} `json:"synonyms"`
				ReplacedBy             []interface{} `json:"replacedBy"`
				OutputPortDataProducts []interface{} `json:"outputPortDataProducts"`
				Readme                 interface{}   `json:"readme"`
				ReplacementTerms       []interface{} `json:"replacementTerms"`
				Meanings               []interface{} `json:"meanings"`
				McMonitors             []interface{} `json:"mcMonitors"`
				TranslatedTerms        []interface{} `json:"translatedTerms"`
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
				Files      []interface{} `json:"files"`
				Metrics    []interface{} `json:"metrics"`
				SodaChecks []interface{} `json:"sodaChecks"`
			} `json:"relationshipAttributes"`
			Tags []struct {
				TypeName                          string `json:"typeName"`
				EntityGuid                        string `json:"entityGuid"`
				EntityStatus                      string `json:"entityStatus"`
				Propagate                         bool   `json:"propagate"`
				RemovePropagationsOnEntityDelete  bool   `json:"removePropagationsOnEntityDelete"`
				RestrictPropagationThroughLineage bool   `json:"restrictPropagationThroughLineage"`
			} `json:"classifications"`
			Anchor struct {
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
		} `json:"entity"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Copy fields
	gt.TypeName = temp.Entity.TypeName
	gt.Attributes = temp.Entity.Attributes
	gt.Guid = temp.Entity.Guid
	if temp.Entity.Anchor.Guid != "" {
		gt.Anchor = temp.Entity.Anchor
	}
	if temp.Entity.RelationshipAttributes.Anchor.Guid != "" {
		gt.Anchor = temp.Entity.RelationshipAttributes.Anchor
	}
	gt.McMonitors = temp.Entity.RelationshipAttributes.McMonitors
	gt.Tags = temp.Entity.Tags

	return nil
}

func FromJSONTerm(data []byte) (*GlossaryTerm, error) {
	var glossaryResponse GlossaryTerm
	err := json.Unmarshal(data, &glossaryResponse)

	return &glossaryResponse, err
}
