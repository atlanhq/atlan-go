package model

import "encoding/json"

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
	a.Attributes = temp.Entity.Attributes
	a.Guid = temp.Entity.Guid
	a.IsIncomplete = temp.Entity.IsIncomplete
	a.Status = temp.Entity.Status
	a.CreatedBy = temp.Entity.CreatedBy
	a.UpdatedBy = temp.Entity.UpdatedBy
	a.CreateTime = temp.Entity.CreateTime
	a.UpdateTime = temp.Entity.UpdateTime

	return nil
}
