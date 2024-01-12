package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type Attributes string
type LiteralState string
type SortOrder string

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
	Ascending                     SortOrder    = "asc"
	Descending                    SortOrder    = "desc"
)

// Query is an interface that represents the base query behavior.
type Query interface {
	ToJSON() map[string]interface{}
}

type TermQuery struct {
	Field string
	Value interface{}
}

type BoolQuery struct {
	Must               []Query
	Should             []Query
	MustNot            []Query
	Filter             []Query
	TypeName           string
	Boost              *float64
	MinimumShouldMatch *int
}

type MatchAll struct {
	Boost *float64
}

type MatchNone struct{}

type Exists struct {
	Field string
}

type NestedQuery struct {
	Path           string
	Query          Query
	ScoreMode      string
	IgnoreUnmapped bool
}

type Terms struct {
	Field  string
	Values []string
	Boost  *float64
}

type PrefixQuery struct {
	Field           string
	Value           interface{}
	Boost           *float64
	CaseInsensitive *bool
	TypeName        string
}

type RangeQuery struct {
	Field    string
	Gt       *interface{}
	Gte      *interface{}
	Lt       *interface{}
	Lte      *interface{}
	Boost    *float64
	Format   *string
	Relation *string
	TimeZone *string
	TypeName string
}

type WildcardQuery struct {
	Field           string
	Value           string
	Boost           *float64
	CaseInsensitive *bool
}

type RegexpQuery struct {
	Field                 string
	Value                 string
	Boost                 *float64
	CaseInsensitive       *bool
	MaxDeterminizedStates *int
}

type FuzzyQuery struct {
	Field          string
	Value          string
	Fuzziness      *string
	MaxExpansions  *int
	PrefixLength   *int
	Transpositions *bool
	Rewrite        *string
}

type MatchQuery struct {
	Field                           string
	Query                           string
	Analyzer                        *string
	AutoGenerateSynonymsPhraseQuery *bool
	Fuzziness                       *string
	FuzzyTranspositions             *bool
	FuzzyRewrite                    *string
	Lenient                         *bool
	Operator                        *string
	MinimumShouldMatch              *int
	ZeroTermsQuery                  *string
	MaxExpansions                   *int
	PrefixLength                    *int
}

type SortItem struct {
	Field      string
	Order      *SortOrder
	NestedPath *string
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
	boolQuery := make(map[string]interface{})
	if len(b.Must) > 0 {
		mustClauses := make([]map[string]interface{}, len(b.Must))
		for i, q := range b.Must {
			mustClauses[i] = q.ToJSON()
		}
		boolQuery["must"] = mustClauses
	}
	if len(b.Should) > 0 {
		shouldClauses := make([]map[string]interface{}, len(b.Should))
		for i, q := range b.Should {
			shouldClauses[i] = q.ToJSON()
		}
		boolQuery["should"] = shouldClauses
	}
	if len(b.MustNot) > 0 {
		mustNotClauses := make([]map[string]interface{}, len(b.MustNot))
		for i, q := range b.MustNot {
			mustNotClauses[i] = q.ToJSON()
		}
		boolQuery["must_not"] = mustNotClauses
	}
	if len(b.Filter) > 0 {
		filterClauses := make([]map[string]interface{}, len(b.Filter))
		for i, q := range b.Filter {
			filterClauses[i] = q.ToJSON()
		}
		boolQuery["filter"] = filterClauses
	}
	if b.Boost != nil {
		boolQuery["boost"] = *b.Boost
	}
	if b.MinimumShouldMatch != nil {
		boolQuery["minimum_should_match"] = *b.MinimumShouldMatch
	}
	return map[string]interface{}{"bool": boolQuery}
}

func (m *MatchAll) ToJSON() map[string]interface{} {
	query := make(map[string]interface{})
	if m.Boost != nil {
		query["boost"] = *m.Boost
	}
	return map[string]interface{}{
		"match_all": query,
	}
}

func (m *MatchNone) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"match_none": map[string]interface{}{},
	}
}

func (e *Exists) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"exists": map[string]interface{}{
			"field": e.Field,
		},
	}
}

func (n *NestedQuery) ToJSON() map[string]interface{} {
	query := map[string]interface{}{
		"path":  n.Path,
		"query": n.Query.ToJSON(),
	}
	if n.ScoreMode != "" {
		query["score_mode"] = n.ScoreMode
	}
	if n.IgnoreUnmapped {
		query["ignore_unmapped"] = n.IgnoreUnmapped
	}
	return map[string]interface{}{
		"nested": query,
	}
}

func (t *Terms) ToJSON() map[string]interface{} {
	query := map[string]interface{}{
		"terms": map[string]interface{}{
			t.Field: t.Values,
		},
	}
	if t.Boost != nil {
		query["terms"].(map[string]interface{})["boost"] = *t.Boost
	}
	return query
}

func (p *PrefixQuery) ToJSON() map[string]interface{} {
	prefixQuery := map[string]interface{}{
		"value": p.Value,
	}
	if p.Boost != nil {
		prefixQuery["boost"] = *p.Boost
	}
	if p.CaseInsensitive != nil {
		prefixQuery["case_insensitive"] = *p.CaseInsensitive
	}
	return map[string]interface{}{
		"prefix": map[string]interface{}{
			p.Field: prefixQuery,
		},
	}
}

func (r *RangeQuery) ToJSON() map[string]interface{} {
	rangeQuery := make(map[string]interface{})
	if r.Gt != nil {
		rangeQuery["gt"] = *r.Gt
	}
	if r.Gte != nil {
		rangeQuery["gte"] = *r.Gte
	}
	if r.Lt != nil {
		rangeQuery["lt"] = *r.Lt
	}
	if r.Lte != nil {
		rangeQuery["lte"] = *r.Lte
	}
	if r.Boost != nil {
		rangeQuery["boost"] = *r.Boost
	}
	if r.Format != nil {
		rangeQuery["format"] = *r.Format
	}
	if r.Relation != nil {
		rangeQuery["relation"] = *r.Relation
	}
	if r.TimeZone != nil {
		rangeQuery["time_zone"] = *r.TimeZone
	}
	return map[string]interface{}{
		"range": map[string]interface{}{
			r.Field: rangeQuery,
		},
	}
}

func (w *WildcardQuery) ToJSON() map[string]interface{} {
	wildcardQuery := map[string]interface{}{
		"value": w.Value,
	}
	if w.Boost != nil {
		wildcardQuery["boost"] = *w.Boost
	}
	if w.CaseInsensitive != nil {
		wildcardQuery["case_insensitive"] = *w.CaseInsensitive
	}
	return map[string]interface{}{
		"wildcard": map[string]interface{}{
			w.Field: wildcardQuery,
		},
	}
}

func (r *RegexpQuery) ToJSON() map[string]interface{} {
	regexpQuery := map[string]interface{}{
		"value": r.Value,
	}
	if r.Boost != nil {
		regexpQuery["boost"] = *r.Boost
	}
	if r.CaseInsensitive != nil {
		regexpQuery["case_insensitive"] = *r.CaseInsensitive
	}
	if r.MaxDeterminizedStates != nil {
		regexpQuery["max_determinized_states"] = *r.MaxDeterminizedStates
	}
	return map[string]interface{}{
		"regexp": map[string]interface{}{
			r.Field: regexpQuery,
		},
	}
}

func (f *FuzzyQuery) ToJSON() map[string]interface{} {
	fuzzyQuery := map[string]interface{}{
		"value": f.Value,
	}
	if f.Fuzziness != nil {
		fuzzyQuery["fuzziness"] = *f.Fuzziness
	}
	if f.MaxExpansions != nil {
		fuzzyQuery["max_expansions"] = *f.MaxExpansions
	}
	if f.PrefixLength != nil {
		fuzzyQuery["prefix_length"] = *f.PrefixLength
	}
	if f.Transpositions != nil {
		fuzzyQuery["transpositions"] = *f.Transpositions
	}
	if f.Rewrite != nil {
		fuzzyQuery["rewrite"] = *f.Rewrite
	}
	return map[string]interface{}{
		"fuzzy": map[string]interface{}{
			f.Field: fuzzyQuery,
		},
	}
}

func (m *MatchQuery) ToJSON() map[string]interface{} {
	matchQuery := map[string]interface{}{
		"query": m.Query,
	}
	if m.Analyzer != nil {
		matchQuery["analyzer"] = *m.Analyzer
	}
	if m.AutoGenerateSynonymsPhraseQuery != nil {
		matchQuery["auto_generate_synonyms_phrase_query"] = *m.AutoGenerateSynonymsPhraseQuery
	}
	if m.Fuzziness != nil {
		matchQuery["fuzziness"] = *m.Fuzziness
	}
	if m.FuzzyTranspositions != nil {
		matchQuery["fuzzy_transpositions"] = *m.FuzzyTranspositions
	}
	if m.FuzzyRewrite != nil {
		matchQuery["fuzzy_rewrite"] = *m.FuzzyRewrite
	}
	if m.Lenient != nil {
		matchQuery["lenient"] = *m.Lenient
	}
	if m.Operator != nil {
		matchQuery["operator"] = *m.Operator
	}
	if m.MinimumShouldMatch != nil {
		matchQuery["minimum_should_match"] = *m.MinimumShouldMatch
	}
	if m.ZeroTermsQuery != nil {
		matchQuery["zero_terms_query"] = *m.ZeroTermsQuery
	}
	if m.MaxExpansions != nil {
		matchQuery["max_expansions"] = *m.MaxExpansions
	}
	if m.PrefixLength != nil {
		matchQuery["prefix_length"] = *m.PrefixLength
	}
	return map[string]interface{}{
		"match": map[string]interface{}{
			m.Field: matchQuery,
		},
	}
}

func (s *SortItem) ToJSON() map[string]interface{} {
	parameters := map[string]interface{}{
		"order": s.Order,
	}
	if s.NestedPath != nil {
		parameters["nested"] = map[string]interface{}{
			"path": s.NestedPath,
		}
	}
	return map[string]interface{}{
		s.Field: parameters,
	}
}

type IndexSearchIterator struct {
	request        IndexSearchRequest
	currentPage    int
	pageSize       int
	totalResults   int64
	hasMoreResults bool
}

type SearchRequest struct {
	Attributes []string `json:"attributes,omitempty"`
	Offset     int      `json:"from,omitempty"`
	Size       int      `json:"size,omitempty"`
}

func NewIndexSearchIterator(pageSize int, initialRequest IndexSearchRequest) *IndexSearchIterator {
	return &IndexSearchIterator{
		request:        initialRequest,
		currentPage:    0,
		pageSize:       pageSize,
		totalResults:   0,
		hasMoreResults: true,
	}
}

func (it *IndexSearchIterator) NextPage() (*IndexSearchResponse, error) {
	if !it.hasMoreResults {
		return nil, fmt.Errorf("no more results available")
	}

	it.request.Dsl.From = it.currentPage * it.pageSize
	it.request.Dsl.Size = it.pageSize

	response, err := search(it.request)
	if err != nil {
		return nil, err
	}

	it.totalResults = response.ApproximateCount
	it.hasMoreResults = int64(it.request.Dsl.From+it.pageSize) < it.totalResults
	it.currentPage++

	return response, nil
}

func (it *IndexSearchIterator) CurrentPage() int {
	return it.currentPage
}

func (it *IndexSearchIterator) IteratePages() ([]*IndexSearchResponse, error) {
	if !it.hasMoreResults {
		return nil, fmt.Errorf("no more results available")
	}

	// Perform an initial search to get the approximateCount
	it.request.Dsl.From = 0
	it.request.Dsl.Size = it.pageSize
	response, err := search(it.request)
	if err != nil {
		return nil, err
	}
	it.totalResults = response.ApproximateCount
	it.hasMoreResults = it.totalResults > 0
	if !it.hasMoreResults {
		return nil, fmt.Errorf("no more results available")
	}

	// If approximateCount is 1, return the response immediately
	if it.totalResults == 1 {
		return []*IndexSearchResponse{response}, nil
	}

	// Num of pages to fetch
	numPageGroups := int((it.totalResults + int64(it.pageSize) - 1) / int64(it.pageSize))
	var wg sync.WaitGroup
	responses := make([]*IndexSearchResponse, numPageGroups)
	errors := make([]error, numPageGroups)

	// Fetch all pages in parallel
	for i := 0; i < numPageGroups; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			it.request.Dsl.From = i * it.pageSize
			it.request.Dsl.Size = it.pageSize
			response, err := search(it.request)
			if err != nil {
				errors[i] = err
				return
			}
			responses[i] = response
		}(i)
	}

	wg.Wait()

	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}

	// Update the iterator state
	it.hasMoreResults = int64(it.request.Dsl.From+it.pageSize) < it.totalResults
	it.currentPage++

	return responses, nil
}

func (it *IndexSearchIterator) HasMoreResults() bool {
	return it.hasMoreResults
}

type IndexSearchRequest struct {
	SearchRequest
	Dsl                    dsl      `json:"dsl"`
	RelationAttributes     []string `json:"relationAttributes,omitempty"`
	SuppressLogs           bool     `json:"suppressLogs"`
	ShowSearchScore        bool     `json:"showSearchScore"`
	ExcludeMeanings        bool     `json:"excludeMeanings"`
	ExcludeClassifications bool     `json:"excludeClassifications"`
}

type dsl struct {
	From                int                    `json:"from"`
	Size                int                    `json:"size"`
	aggregation         map[string]interface{} `json:"aggregation,omitempty"`
	Query               map[string]interface{} `json:"query"`
	TrackTotalHits      bool                   `json:"track_total_hits"`
	PostFilter          *Query                 `json:"post_filter,omitempty"`
	Sort                []SortItem             `json:"sort,omitempty"`
	IncludesOnResults   []string               `json:"includesOnResults,omitempty"`
	IncludesOnRelations []string               `json:"includesOnRelations,omitempty"`
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
	if err != nil {
		return nil, err
	}
	pageSize := 1

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

	iterator := NewIndexSearchIterator(pageSize, request)

	for iterator.HasMoreResults() {
		responses, err := iterator.IteratePages()
		if err != nil {
			return nil, fmt.Errorf("error executing search: %v", err)
		}
		for _, response := range responses {
			for _, entity := range response.Entities {
				if entity.TypeName == "AtlasGlossary" {
					return response, nil
				}
			}
		}
	}
	// Call the search function
	//response, err := search(request)
	//if err != nil {
	//	return nil, fmt.Errorf("error executing search: %v", err)
	//}

	// return response, nil
	return nil, nil
}

func FindCategoryByName(categoryName string, glossaryQualifiedName string) (*IndexSearchResponse, error) {
	boolQuery, err := WithActiveCategory(categoryName, glossaryQualifiedName)
	if err != nil {
		return nil, err
	}
	pageSize := 1

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

	iterator := NewIndexSearchIterator(pageSize, request)

	for iterator.HasMoreResults() {
		response, err := iterator.NextPage()
		if err != nil {
			return nil, fmt.Errorf("error executing search: %v", err)
		}
		fmt.Println("Current Page: ", iterator.CurrentPage())
		for _, entity := range response.Entities {
			if entity.TypeName == "AtlasGlossaryCategory" {
				return response, err
			}
		}
	}
	// Call the search function
	//response, err := search(request)
	//if err != nil {
	//	return nil, fmt.Errorf("error executing search: %v", err)
	//}

	// return response, nil
	return nil, nil
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

func WithActiveCategory(name string, glossaryqualifiedname string) (*BoolQuery, error) {
	q1, err := WithState("ACTIVE")
	if err != nil {
		return nil, err
	}
	q2 := WithTypeName("AtlasGlossaryCategory")
	q3 := WithName(name)
	q4 := WithGlossary(glossaryqualifiedname)
	return &BoolQuery{
		Filter: []Query{q1, q2, q3, q4},
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

func WithGlossary(value string) *TermQuery {
	return &TermQuery{
		Field: string(Glossary),
		Value: value,
	}
}
