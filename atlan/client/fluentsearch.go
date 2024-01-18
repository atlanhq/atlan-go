package client

import (
	"fmt"
)

type FluentSearch struct {
	Wheres              []Query
	WhereNots           []Query
	WhereSomes          []Query
	MinSomes            *int
	Sorts               []SortItem
	PageSize            int
	Aggregations        map[string]interface{}
	IncludesOnResults   []string
	IncludesOnRelations []string
}

type Aggregation struct {
}

func (fs *FluentSearch) ActiveAssets() *FluentSearch {
	activeAssetsCondition := &TermQuery{
		Field: string(State),
		Value: "ACTIVE",
	}
	fs.Wheres = append(fs.Wheres, activeAssetsCondition)
	return fs
}
func (fs *FluentSearch) ArchivedAssets() *FluentSearch {
	archivedAssetsCondition := &TermQuery{
		Field: string(State),
		Value: "Deleted",
	}
	fs.Wheres = append(fs.Wheres, archivedAssetsCondition)
	return fs
}

func (fs *FluentSearch) AssetType(of string) *FluentSearch {
	assetTypeCondition := &TermQuery{
		Field: string(TypeName),
		Value: of,
	}
	fs.Wheres = append(fs.Wheres, assetTypeCondition)
	return fs
}

func (fs *FluentSearch) AssetTypes(oneOf []string) *FluentSearch {
	assetTypesCondition := &TermQuery{
		Field: string(TypeName),
		Value: oneOf,
	}
	fs.Wheres = append(fs.Wheres, assetTypesCondition)
	return fs
}

func NewFluentSearch() *FluentSearch {
	return &FluentSearch{}
}

// Where adds a TermQuery to the Wheres slice.
func (fs *FluentSearch) Where(field string, value string) *FluentSearch {
	fs.Wheres = append(fs.Wheres, &TermQuery{Field: field, Value: value})
	return fs
}

// WhereNot adds a BoolQuery with MustNot clause to the WhereNots slice.
func (fs *FluentSearch) WhereNot(queries ...Query) *FluentSearch {
	boolQuery := &BoolQuery{MustNot: queries}
	fs.WhereNots = append(fs.WhereNots, boolQuery)
	return fs
}

// WhereSome adds a BoolQuery with Should clause to the WhereSomes slice.
func (fs *FluentSearch) WhereSome(queries ...Query) *FluentSearch {
	boolQuery := &BoolQuery{Should: queries}
	fs.WhereSomes = append(fs.WhereSomes, boolQuery)
	return fs
}

// MinSome sets the MinSomes field.
func (fs *FluentSearch) MinSome(minSomes int) *FluentSearch {
	fs.MinSomes = &minSomes
	return fs
}

// Sort adds a SortItem to the Sorts slice.
func (fs *FluentSearch) Sort(field string, order SortOrder) *FluentSearch {
	fs.Sorts = append(fs.Sorts, SortItem{Field: field, Order: order})
	return fs
}

// Aggregate adds an aggregation to the Aggregations map.
func (fs *FluentSearch) Aggregate(name string, aggregation map[string]interface{}) *FluentSearch {
	if fs.Aggregations == nil {
		fs.Aggregations = make(map[string]interface{})
	}
	fs.Aggregations[name] = aggregation
	return fs
}

// PageSize sets the PageSize field.
func (fs *FluentSearch) PageSizes(size int) *FluentSearch {
	fs.PageSize = size
	return fs
}

// IncludeOnResults adds fields to the IncludesOnResults slice.
func (fs *FluentSearch) IncludeOnResults(fields ...string) *FluentSearch {
	fs.IncludesOnResults = append(fs.IncludesOnResults, fields...)
	return fs
}

// IncludeOnRelations adds fields to the IncludesOnRelations slice.
func (fs *FluentSearch) IncludeOnRelations(fields ...string) *FluentSearch {
	fs.IncludesOnRelations = append(fs.IncludesOnRelations, fields...)
	return fs
}

// Execute performs the search and returns the results.
func (fs *FluentSearch) Execute() ([]*IndexSearchResponse, error) {
	pageSize := fs.PageSize
	request := fs.ToRequest()

	iterator := NewIndexSearchIterator(pageSize, *request)
	responses := make([]*IndexSearchResponse, 0)

	for iterator.HasMoreResults() {
		{
			response, err := iterator.NextPage()
			if err != nil {
				fmt.Printf("Error executing search: %v\n", err)
				return nil, err
			}

			responses = append(responses, response)

		}
	}

	return responses, nil
}

// Sort by GUID by default only if not already specified by the developer
func (fs *FluentSearch) SortByGuidDefault() *FluentSearch {
	// Check if "guid" is already in the list of sort criteria
	guidAlreadySorted := false
	for _, item := range fs.Sorts {
		if item.Field == string(GUID) {
			guidAlreadySorted = true
			break
		}
	}

	// If "guid" is not already in the list, add it as the final sort criteria
	if !guidAlreadySorted {
		fs.Sort(string(GUID), Ascending)
	}

	return fs
}

// ToRequest converts FluentSearch to IndexSearchRequest.
func (fs *FluentSearch) ToRequest() *IndexSearchRequest {
	// Create a new IndexSearchRequest and set its properties based on FluentSearch
	request := &IndexSearchRequest{
		SearchRequest: SearchRequest{
			Attributes:          fs.IncludesOnResults,
			RelationsAttributes: fs.IncludesOnRelations,
		},
		Dsl: dsl{
			From:           0,
			Size:           fs.PageSize,
			aggregation:    fs.Aggregations,
			TrackTotalHits: true,
		},
	}

	// Add Wheres to Query
	if len(fs.Wheres) > 0 {
		boolQuery := &BoolQuery{Filter: fs.Wheres}
		request.Dsl.Query = boolQuery.ToJSON()
	}

	// Add WhereNots to Query
	if len(fs.WhereNots) > 0 {
		boolQuery := &BoolQuery{MustNot: fs.WhereNots}
		if request.Dsl.Query == nil {
			request.Dsl.Query = boolQuery.ToJSON()
		} else {
			request.Dsl.Query = map[string]interface{}{
				"bool": map[string]interface{}{
					"filter":   request.Dsl.Query,
					"must":     request.Dsl.Query,
					"must_not": boolQuery.ToJSON(),
				},
			}
		}
	}

	// Add WhereSomes to Query
	if len(fs.WhereSomes) > 0 {
		boolQuery := &BoolQuery{Should: fs.WhereSomes, MinimumShouldMatch: fs.MinSomes}
		if request.Dsl.Query == nil {
			request.Dsl.Query = boolQuery.ToJSON()
		} else {
			request.Dsl.Query = map[string]interface{}{
				"bool": map[string]interface{}{
					"filter": request.Dsl.Query,
					"must":   request.Dsl.Query,
					"should": boolQuery.ToJSON(),
				},
			}
		}
	}

	// Add Sorts to Dsl.Sort
	if len(fs.Sorts) > 0 {
		fs.SortByGuidDefault()
		sortItems := fs.Sorts
		sortItemsJSON := make([]map[string]interface{}, len(sortItems))
		for i, item := range sortItems {
			sortItemsJSON[i] = item.ToJSON()
		}
		request.Dsl.Sort = sortItemsJSON
	}

	return request
}
