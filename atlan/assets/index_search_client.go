package assets

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model"
)

// Call the search API
func Search(request model.IndexSearchRequest) (*model.IndexSearchResponse, error) {
	// Define the API endpoint and method
	api := &INDEX_SEARCH

	// Call the API
	responseBytes, err := DefaultAtlanClient.CallAPI(api, nil, &request)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response
	var response model.IndexSearchResponse
	err = json.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// FindGlossaryByName searches for a glossary by name.
func FindGlossaryByName(glossaryName string) (*model.IndexSearchResponse, error) {
	boolQuery, err := WithActiveGlossary(glossaryName)
	if err != nil {
		return nil, err
	}
	pageSize := 1

	sortItems := []model.SortItem{{Field: string(MODIFIED_BY), Order: atlan.SortOrderAscending}}
	sortItemsJSON := make([]map[string]interface{}, len(sortItems))
	for i, item := range sortItems {
		sortItemsJSON[i] = item.ToJSON()
	}

	request := model.IndexSearchRequest{
		Dsl: model.Dsl{
			From:           0,
			Size:           2,
			Query:          boolQuery.ToJSON(),
			TrackTotalHits: true,
			Sort:           sortItemsJSON,
		},
		SuppressLogs:     true,
		ShowSearchScore:  false,
		ExcludeMeanings:  false,
		ExcludeAtlanTags: false,
		Metadata: model.Metadata{
			SaveSearchLog: true,
			UtmTags:       []string{atlan.PROJECT_SDK_GO.String()},
		},
	}

	iterator := NewIndexSearchIterator(pageSize, request)

	for iterator.HasMoreResults() {
		responses, err := iterator.IteratePages()
		if err != nil {
			return nil, fmt.Errorf("error executing search: %v", err)
		}
		for _, response := range responses {
			for _, entity := range response.Entities {
				if *entity.TypeName == "AtlasGlossary" {
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

// FindCategoryByName searches for a category by name.
func FindCategoryByName(categoryName string, glossaryQualifiedName string) (*model.IndexSearchResponse, error) {
	boolQuery, err := WithActiveCategory(categoryName, glossaryQualifiedName)
	if err != nil {
		return nil, err
	}
	pageSize := 1

	request := model.IndexSearchRequest{
		Dsl: model.Dsl{
			From:           0,
			Size:           2,
			Query:          boolQuery.ToJSON(),
			TrackTotalHits: true,
		},
		SuppressLogs:     true,
		ShowSearchScore:  false,
		ExcludeMeanings:  false,
		ExcludeAtlanTags: false,
	}

	iterator := NewIndexSearchIterator(pageSize, request)

	for iterator.HasMoreResults() {
		response, err := iterator.NextPage()
		if err != nil {
			return nil, fmt.Errorf("error executing search: %v", err)
		}
		fmt.Println("Current Page: ", iterator.CurrentPage())
		for _, entity := range response.Entities {
			if *entity.TypeName == "AtlasGlossaryCategory" {
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

// WithActiveGlossary returns a query for an active glossary by name.
func WithActiveGlossary(name string) (*model.BoolQuery, error) {
	q1, err := WithState("ACTIVE")
	if err != nil {
		return nil, err
	}
	q2 := WithTypeName("AtlasGlossary")
	q3 := WithName(name)

	return &model.BoolQuery{
		Filter: []model.Query{q1, q2, q3},
	}, nil
}

// WithActiveCategory returns a query for an active category by name.
func WithActiveCategory(name string, glossaryqualifiedname string) (*model.BoolQuery, error) {
	q1, err := WithState("ACTIVE")
	if err != nil {
		return nil, err
	}
	q2 := WithTypeName("AtlasGlossaryCategory")
	q3 := WithName(name)
	q4 := WithGlossary(glossaryqualifiedname)
	return &model.BoolQuery{
		Filter: []model.Query{q1, q2, q3, q4},
	}, nil
}

// WithActivePersona returns a query for an active persona by name.
func WithActivePersona(name string) (*model.BoolQuery, error) {
	q1, err := WithState("ACTIVE")
	if err != nil {
		return nil, err
	}
	q2 := WithTypeName("Persona")
	q3 := WithName(name)

	return &model.BoolQuery{
		Filter: []model.Query{q1, q2, q3},
	}, nil
}

// WithActivePurpose returns a query for an active purpose by name.
func WithActivePurpose(name string) (*model.BoolQuery, error) {
	q1, err := WithState("ACTIVE")
	if err != nil {
		return nil, err
	}
	q2 := WithTypeName("Purpose")
	q3 := WithName(name)

	return &model.BoolQuery{
		Filter: []model.Query{q1, q2, q3},
	}, nil
}

// Helper Functions

// WithState returns a query for an entity with a specific state.
func WithState(value string) (*model.TermQuery, error) {
	if value != atlan.LiteralStateActive.String() && value != atlan.LiteralStateDeleted.String() && value != atlan.LiteralStatePurged.String() {
		return nil, errors.New("invalid state")
	}
	return &model.TermQuery{
		Field: string(STATE),
		Value: value,
	}, nil
}

// WithTypeName returns a query for an entity with a specific type name.
func WithTypeName(value string) *model.TermQuery {
	return &model.TermQuery{
		Field: string(TYPE_NAME),
		Value: value,
	}
}

// WithName returns a query for an entity with a specific name.
func WithName(value string) *model.TermQuery {
	return &model.TermQuery{
		Field: string(NAME),
		Value: value,
	}
}

// WithGlossary returns a query for an entity with a specific glossary.
func WithGlossary(value string) *model.TermQuery {
	return &model.TermQuery{
		Field: string(GLOSSARY),
		Value: value,
	}
}

// Pagination Implemented here:

type IndexSearchIterator struct {
	request        model.IndexSearchRequest
	currentPage    int
	pageSize       int
	totalResults   int64
	hasMoreResults bool
}

func NewIndexSearchIterator(pageSize int, initialRequest model.IndexSearchRequest) *IndexSearchIterator {
	return &IndexSearchIterator{
		request:        initialRequest,
		currentPage:    0,
		pageSize:       pageSize,
		totalResults:   0,
		hasMoreResults: true,
	}
}

// NextPage returns the next page of search results.
func (it *IndexSearchIterator) NextPage() (*model.IndexSearchResponse, error) {
	if !it.hasMoreResults {
		return nil, fmt.Errorf("no more results available")
	}

	it.request.Dsl.From = it.currentPage * it.pageSize
	it.request.Dsl.Size = it.pageSize

	response, err := Search(it.request)
	if err != nil {
		return nil, err
	}

	it.totalResults = response.ApproximateCount
	it.hasMoreResults = int64(it.request.Dsl.From+it.pageSize) < it.totalResults
	it.currentPage++

	return response, nil
}

// CurrentPage returns the current page number.
func (it *IndexSearchIterator) CurrentPage() int {
	return it.currentPage
}

// IteratePages returns all pages of search results.
func (it *IndexSearchIterator) IteratePages() ([]*model.IndexSearchResponse, error) {
	if !it.hasMoreResults {
		return nil, fmt.Errorf("no more results available")
	}

	if it.pageSize == 0 {
		it.pageSize = 300 // Set Default Page Size
	}
	// Perform an initial search to get the approximateCount
	it.request.Dsl.From = 0
	it.request.Dsl.Size = it.pageSize
	response, err := Search(it.request)
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
		return []*model.IndexSearchResponse{response}, nil
	}

	// Num of pages to fetch
	numPageGroups := int((it.totalResults + int64(it.pageSize) - 1) / int64(it.pageSize))
	var wg sync.WaitGroup
	responses := make([]*model.IndexSearchResponse, numPageGroups)
	errors := make([]error, numPageGroups)

	// Fetch all pages in parallel
	for i := 0; i < numPageGroups; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			it.request.Dsl.From = i * it.pageSize
			it.request.Dsl.Size = it.pageSize
			response, err := Search(it.request)
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

// HasMoreResults returns whether there are more results available.
func (it *IndexSearchIterator) HasMoreResults() bool {
	return it.hasMoreResults
}
