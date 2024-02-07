package model

import "reflect"

type Asset struct {
	TypeName            string `json:"typeName"`
	Attributes          Attributes
	Guid                string `json:"guid"`
	Status              string
	DisplayText         string
	ClassificationNames []string `json:"classificationNames"`
	Classifications     []string
	MeaningNames        []string `json:"meaningNames"`
	Meanings            []string
	IsIncomplete        bool `json:"isIncomplete"`
	Labels              []string
	CreatedBy           string `json:"createdBy"`
	UpdatedBy           string `json:"updatedBy"`
	CreateTime          int64  `json:"createTime"`
	UpdateTime          int64  `json:"updateTime"`
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

type MutatedEntities struct {
	UPDATE         []*Asset `json:"UPDATE"`
	CREATE         []*Asset `json:"CREATE"`
	DELETE         []*Asset `json:"DELETE"`
	PARTIAL_UPDATE []*Asset `json:"PARTIAL_UPDATE"`
}

type AssetMutationResponse struct {
	GuidAssignments        map[string]string `json:"guidAssignments"`
	MutatedEntities        *MutatedEntities  `json:"mutatedEntities"`
	PartialUpdatedEntities []*Asset          `json:"partialUpdatedEntities"`
}

func (amr *AssetMutationResponse) AssetsUpdated(assetType reflect.Type) []*Asset {
	if amr.MutatedEntities != nil && amr.MutatedEntities.UPDATE != nil {
		var assets []*Asset
		for _, asset := range amr.MutatedEntities.UPDATE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*Asset{}
}
func (amr *AssetMutationResponse) AssetsCreated(assetType reflect.Type) []*Asset {
	if amr.MutatedEntities != nil && amr.MutatedEntities.CREATE != nil {
		var assets []*Asset
		for _, asset := range amr.MutatedEntities.CREATE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*Asset{}
}

func (amr *AssetMutationResponse) AssetsDeleted(assetType reflect.Type) []*Asset {
	if amr.MutatedEntities != nil && amr.MutatedEntities.DELETE != nil {
		var assets []*Asset
		for _, asset := range amr.MutatedEntities.DELETE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*Asset{}
}

func (amr *AssetMutationResponse) AssetsPartiallyUpdated(assetType reflect.Type) []*Asset {
	if amr.MutatedEntities != nil && amr.MutatedEntities.PARTIAL_UPDATE != nil {
		var assets []*Asset
		for _, asset := range amr.MutatedEntities.PARTIAL_UPDATE {
			if reflect.TypeOf(asset).Elem() == assetType {
				assets = append(assets, asset)
			}
		}
		return assets
	}
	return []*Asset{}
}
