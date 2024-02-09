package model

import (
	"encoding/json"
	"reflect"
)

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

// Add Mutated Assets for Respomse in Creation, Updation and Deletion
// Unmarshal on Assets changed the unmarshalling for the whole sdk asset structure
type MutatedAssets struct {
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

// Unmarshalling for Assets from JSON
// Used in RetrieveMinimal Function

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

type MutatedEntities struct {
	//Assets that were assets_updated. The detailed properties of the returned asset will vary based on
	//the type of asset, but listed in the example are the common set of properties across assets.
	UPDATE []*MutatedAssets `json:"UPDATE"`

	// Assets that were created. The detailed properties of the returned asset will vary based on the
	// type of asset, but listed in the example are the common set of properties across assets.
	CREATE []*MutatedAssets `json:"CREATE"`

	// Assets that were deleted. The detailed properties of the returned asset will vary based on the
	// type of asset, but listed in the example are the common set of properties across assets.
	DELETE []*MutatedAssets `json:"DELETE"`

	// Assets that were partially updated. The detailed properties of the returned asset will
	// vary based on the type of asset, but listed in the example are the common set of properties across assets.
	PARTIAL_UPDATE []*MutatedAssets `json:"PARTIAL_UPDATE"`
}

type AssetMutationResponse struct {
	// Map of assigned unique identifiers for the changed assets.
	GuidAssignments map[string]string `json:"guidAssignments,omitempty"`

	// Assets that were changed.
	MutatedEntities *MutatedEntities `json:"mutatedEntities"`

	// Assets that were partially updated.
	PartialUpdatedEntities []*MutatedAssets `json:"partialUpdatedEntities,omitempty"`
}

func (amr *AssetMutationResponse) AssetsUpdated(assetType reflect.Type) []*MutatedAssets {
	if amr.MutatedEntities != nil && amr.MutatedEntities.UPDATE != nil {
		var assets []*MutatedAssets
		for _, asset := range amr.MutatedEntities.UPDATE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*MutatedAssets{}
}
func (amr *AssetMutationResponse) AssetsCreated(assetType reflect.Type) []*MutatedAssets {
	if amr.MutatedEntities != nil && amr.MutatedEntities.CREATE != nil {
		var assets []*MutatedAssets
		for _, asset := range amr.MutatedEntities.CREATE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*MutatedAssets{}
}

func (amr *AssetMutationResponse) AssetsDeleted(assetType reflect.Type) []*MutatedAssets {
	if amr.MutatedEntities != nil && amr.MutatedEntities.DELETE != nil {
		var assets []*MutatedAssets
		for _, asset := range amr.MutatedEntities.DELETE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*MutatedAssets{}
}

func (amr *AssetMutationResponse) AssetsPartiallyUpdated(assetType reflect.Type) []*MutatedAssets {
	if amr.MutatedEntities != nil && amr.MutatedEntities.PARTIAL_UPDATE != nil {
		var assets []*MutatedAssets
		for _, asset := range amr.MutatedEntities.PARTIAL_UPDATE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*MutatedAssets{}
}
