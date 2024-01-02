package client

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Attributes string
type LiteralState string

const (
	ConnectorName                 Attributes   = "connectorName"
	Categories                    Attributes   = "__categories"
	CreateTimeAsTimestamp         Attributes   = "__timestamp"
	CreatedBy                     Attributes   = "__createdBy"
	Glossary                      Attributes   = "__glossary"
	GUID                          Attributes   = "__guid"
	HasLineage                    Attributes   = "__hasLineage"
	Meanings                      Attributes   = "__meanings"
	ModifiedBy                    Attributes   = "__modifiedBy"
	Name                          Attributes   = "name.keyword"
	OwnerGroups                   Attributes   = "ownerGroups"
	OwnerUsers                    Attributes   = "ownerUsers"
	ParentCategory                Attributes   = "__parentCategory"
	PopularityScore               Attributes   = "popularityScore"
	QualifiedName                 Attributes   = "qualifiedName"
	State                         Attributes   = "__state"
	SuperTypeNames                Attributes   = "__superTypeNames.keyword"
	TypeName                      Attributes   = "__typeName.keyword"
	UpdateTimeAsTimestamp         Attributes   = "__modificationTimestamp"
	CertificateStatus             Attributes   = "certificateStatus"
	ClassificationNames           Attributes   = "__classificationNames"
	ClassificationsText           Attributes   = "__classificationsText"
	CreateTimeAsDate              Attributes   = "__timestamp.date"
	Description                   Attributes   = "description"
	MeaningsText                  Attributes   = "__meaningsText"
	PropagatedClassificationNames Attributes   = "__propagatedClassificationNames"
	PropagatedTraitNames          Attributes   = "__propagatedTraitNames"
	SuperTypeNamesText            Attributes   = "__superTypeNames"
	TraitNames                    Attributes   = "__traitNames"
	UpdateTimeAsDate              Attributes   = "__modificationTimestamp.date"
	UserDescription               Attributes   = "userDescription"
	Active                        LiteralState = "ACTIVE"
	Deleted                       LiteralState = "DELETED"
	Purged                        LiteralState = "PURGED"
)

type TermQuery struct {
	Field string
	Value interface{}
}

// Query is an interface that represents the base query behavior.
type Query interface {
	ToJSON() map[string]interface{}
}

type BoolQuery struct {
	Filter []Query
}

func (t *TermQuery) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"term": map[string]interface{}{
			t.Field: map[string]interface{}{
				"value": t.Value,
			},
		},
	}
}

func (b *BoolQuery) ToJSON() map[string]interface{} {
	filter := make([]interface{}, len(b.Filter))
	for i, f := range b.Filter {
		filter[i] = f.ToJSON()
	}
	return map[string]interface{}{
		"bool": map[string]interface{}{
			"filter": filter,
		},
	}
}

type IndexSearchRequest struct {
	Dsl                    dsl  `json:"dsl"`
	SuppressLogs           bool `json:"suppressLogs"`
	ShowSearchScore        bool `json:"showSearchScore"`
	ExcludeMeanings        bool `json:"excludeMeanings"`
	ExcludeClassifications bool `json:"excludeClassifications"`
}

type dsl struct {
	From           int                    `json:"from"`
	Size           int                    `json:"size"`
	Query          map[string]interface{} `json:"query"`
	TrackTotalHits bool                   `json:"track_total_hits"`
}

type IndexSearchResponse struct {
	QueryType        string           `json:"queryType"`
	SearchParameters SearchParameters `json:"searchParameters"`
	Entities         []Entity         `json:"entities"`
	ApproximateCount int64            `json:"approximateCount"`
}

type SearchParameters struct {
	ShowSearchScore       bool                   `json:"showSearchScore"`
	SuppressLogs          bool                   `json:"suppressLogs"`
	ExcludeMeanings       bool                   `json:"excludeMeanings"`
	ExcludeAtlanTags      bool                   `json:"excludeClassifications"`
	AllowDeletedRelations bool                   `json:"allowDeletedRelations"`
	SaveSearchLog         bool                   `json:"saveSearchLog"`
	RequestMetadata       map[string]interface{} `json:"requestMetadata"`
	Dsl                   dsl                    `json:"dsl"`
	Query                 string                 `json:"query"`
}

type Entity struct {
	TypeName            string                 `json:"typeName"`
	Attributes          map[string]interface{} `json:"attributes"`
	Guid                string                 `json:"guid"`
	Status              string                 `json:"status"`
	DisplayText         string                 `json:"displayText"`
	ClassificationNames []string               `json:"classificationNames"`
	Tags                []interface{}          `json:"classifications"`
	MeaningNames        []interface{}          `json:"meaningNames"`
	Meanings            []interface{}          `json:"meanings"`
	IsIncomplete        bool                   `json:"isIncomplete"`
	Labels              []interface{}          `json:"labels"`
	CreatedBy           string                 `json:"createdBy"`
	UpdatedBy           string                 `json:"updatedBy"`
	CreateTime          int64                  `json:"createTime"`
	UpdateTime          int64                  `json:"updateTime"`
}

func search(request IndexSearchRequest) (*IndexSearchResponse, error) {
	// Define the API endpoint and method
	api := &INDEX_SEARCH

	// Call the API
	responseBytes, err := DefaultAtlanClient.CallAPI(api, nil, &request)
	if err != nil {
		return nil, fmt.Errorf("error calling API: %v", err)
	}

	// Unmarshal the response
	var response IndexSearchResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &response, nil
}

func FindGlossaryByName(glossaryName string) (*IndexSearchResponse, error) {
	boolQuery, err := WithActiveGlossary(glossaryName)
	request := IndexSearchRequest{
		Dsl: dsl{
			From:           0,
			Size:           2,
			Query:          boolQuery.ToJSON(),
			TrackTotalHits: true,
		},
		SuppressLogs:           true,
		ShowSearchScore:        false,
		ExcludeMeanings:        false,
		ExcludeClassifications: false,
	}
	// Call the search function
	response, err := search(request)
	if err != nil {
		return nil, fmt.Errorf("error executing search: %v", err)
	}

	return response, nil
}

// Methods

func WithActiveGlossary(name string) (*BoolQuery, error) {
	q1, err := WithState("ACTIVE")
	if err != nil {
		return nil, err
	}
	q2 := WithTypeName("AtlasGlossary")
	q3 := WithName(name)

	return &BoolQuery{
		Filter: []Query{q1, q2, q3},
	}, nil
}

// Helper Functions

func WithState(value string) (*TermQuery, error) {
	if value != string(Active) && value != string(Deleted) && value != string(Purged) {
		return nil, errors.New("invalid state")
	}
	return &TermQuery{
		Field: string(State),
		Value: value,
	}, nil
}

func WithTypeName(value string) *TermQuery {
	return &TermQuery{
		Field: string(TypeName),
		Value: value,
	}
}

func WithName(value string) *TermQuery {
	return &TermQuery{
		Field: string(Name),
		Value: value,
	}
}
