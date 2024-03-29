// Contains the search model for the Atlas search DSL.

package model

import (
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model/assets"
)

// Query is an interface that represents the base query behavior.
type Query interface {
	ToJSON() map[string]interface{}
}

// TermQuery represents a term query in the Atlas search DSL.
type TermQuery struct {
	Field string
	Value interface{}
}

// BoolQuery represents a boolean query in the Atlas search DSL.
type BoolQuery struct {
	Must               []Query
	Should             []Query
	MustNot            []Query
	Filter             []Query
	TypeName           string
	Boost              *float64
	MinimumShouldMatch *int
}

// MatchAll represents a match_all query in the Atlas search DSL.
type MatchAll struct {
	Boost *float64
}

type MatchNone struct{}

type Exists struct {
	Field string
}

// NestedQuery represents a nested query in the Atlas search DSL.
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
	Gt       *float64
	Gte      *float64
	Lt       *float64
	Lte      *float64
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
	Order      atlan.SortOrder
	NestedPath *string
}

// ToJSON returns the JSON representation of the TermQuery.
func (t *TermQuery) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"term": map[string]interface{}{
			t.Field: map[string]interface{}{
				"value": t.Value,
			},
		},
	}
}

// ToJSON returns the JSON representation of the BoolQuery.
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

// ToJSON returns the JSON representation of the MatchAll query.
func (m *MatchAll) ToJSON() map[string]interface{} {
	query := make(map[string]interface{})
	if m.Boost != nil {
		query["boost"] = *m.Boost
	}
	return map[string]interface{}{
		"match_all": query,
	}
}

// ToJSON returns the JSON representation of the MatchNone query.
func (m *MatchNone) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"match_none": map[string]interface{}{},
	}
}

// ToJSON returns the JSON representation of the Exists query.
func (e *Exists) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"exists": map[string]interface{}{
			"field": e.Field,
		},
	}
}

// ToJSON returns the JSON representation of the NestedQuery.
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

// ToJSON returns the JSON representation of the Terms query.
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

// ToJSON returns the JSON representation of the PrefixQuery.
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

// ToJSON returns the JSON representation of the RangeQuery.
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

// ToJSON returns the JSON representation of the WildcardQuery.
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

// ToJSON returns the JSON representation of the RegexpQuery.
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

// ToJSON returns the JSON representation of the FuzzyQuery.
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

// ToJSON returns the JSON representation of the MatchQuery.
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

// ToJSON returns the JSON representation of the SortItem.
func (s *SortItem) ToJSON() map[string]interface{} {
	sortField := map[string]interface{}{"order": s.Order}
	if s.NestedPath != nil {
		sortField["nested"] = map[string]interface{}{"path": *s.NestedPath}
	}
	return map[string]interface{}{s.Field: sortField}
}

// SearchRequest represents a search request in the Atlas search DSL.
type SearchRequest struct {
	Attributes          []string `json:"attributes,omitempty"`
	Offset              int      `json:"from,omitempty"`
	Size                int      `json:"size,omitempty"`
	RelationsAttributes []string `json:"relationsAttributes,omitempty"`
}

// IndexSearchRequest represents a search request in the Atlas search DSL.
type IndexSearchRequest struct {
	SearchRequest
	Dsl                    Dsl      `json:"dsl"`
	RelationAttributes     []string `json:"relationAttributes,omitempty"`
	SuppressLogs           bool     `json:"suppressLogs"`
	ShowSearchScore        bool     `json:"showSearchScore"`
	ExcludeMeanings        bool     `json:"excludeMeanings"`
	ExcludeClassifications bool     `json:"excludeClassifications"`
}

// Dsl represents the DSL for the Atlas search request.
type Dsl struct {
	From                int                      `json:"from"`
	Size                int                      `json:"size"`
	Aggregation         map[string]interface{}   `json:"aggregation,omitempty"`
	Query               map[string]interface{}   `json:"query"`
	TrackTotalHits      bool                     `json:"track_total_hits"`
	PostFilter          *Query                   `json:"post_filter,omitempty"`
	Sort                []map[string]interface{} `json:"sort,omitempty"`
	IncludesOnResults   []string                 `json:"includesOnResults,omitempty"`
	IncludesOnRelations []string                 `json:"includesOnRelations,omitempty"`
}

// IndexSearchResponse represents a search response in the Atlas search DSL.
type IndexSearchResponse struct {
	QueryType        string           `json:"queryType"`
	SearchParameters SearchParameters `json:"searchParameters"`
	Entities         []SearchAssets   `json:"entities"`
	ApproximateCount int64            `json:"approximateCount"`
}

// SearchParameters represents the search parameters in the Atlas search response.
type SearchParameters struct {
	ShowSearchScore       bool                   `json:"showSearchScore"`
	SuppressLogs          bool                   `json:"suppressLogs"`
	ExcludeMeanings       bool                   `json:"excludeMeanings"`
	ExcludeAtlanTags      bool                   `json:"excludeClassifications"`
	AllowDeletedRelations bool                   `json:"allowDeletedRelations"`
	SaveSearchLog         bool                   `json:"saveSearchLog"`
	RequestMetadata       map[string]interface{} `json:"requestMetadata"`
	Dsl                   Dsl                    `json:"dsl"`
	Query                 string                 `json:"query"`
}

// Entity represents an entity in the Atlas search response.
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

type SearchAssets struct {
	assets.Asset
	assets.Table
	assets.Column
	SearchAttributes *SearchAttributes `json:"Attributes,omitempty"`
}

type SearchAttributes struct {
	QualifiedName *string `json:"qualifiedName,omitempty"`
	Name          *string `json:"name,omitempty"`
}
