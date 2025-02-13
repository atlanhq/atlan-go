package assets

import (
	"encoding/json"
	"fmt"
	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
	"time"
)

const (
	workflowRunSchedule = "orchestration.atlan.com/schedule"
	workflowRunTimezone = "orchestration.atlan.com/timezone"
)

type WorkflowClient struct {
	*AtlanClient
}

// FindByType searches for workflows by their type prefix.
func (w *WorkflowClient) FindByType(prefix atlan.WorkflowPackage, maxResults int) ([]structs.WorkflowSearchResult, error) {

	var query model.Query = &model.BoolQuery{
		Filter: []model.Query{
			&model.NestedQuery{
				Path: "metadata",
				Query: &model.PrefixQuery{
					Field: "metadata.name.keyword",
					Value: prefix.Name,
				},
			},
		},
	}

	nestedPath := "metadata"

	// Add sorting based on creation timestamp in descending order.
	sortItems := []model.SortItem{
		{
			Order:      atlan.SortOrderDescending,
			Field:      "metadata.creationTimestamp",
			NestedPath: &nestedPath,
		},
	}

	request := model.WorkflowSearchRequest{
		Query: query,
		Size:  maxResults,
		Sort:  sortItems,
	}

	rawJSON, err := DefaultAtlanClient.CallAPI(&WORKFLOW_INDEX_SEARCH, nil, &request)
	if err != nil {
		return nil, err
	}
	var response structs.WorkflowSearchResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return response.Hits.Hits, nil
}

// FindByID searches for a workflow by its ID.
func (w *WorkflowClient) FindByID(id string) (*structs.WorkflowSearchResult, error) {
	var query model.Query = &model.BoolQuery{
		Filter: []model.Query{
			&model.NestedQuery{
				Path: "metadata",
				Query: &model.TermQuery{
					Field: "metadata.name.keyword",
					Value: id,
				},
			},
		},
	}

	request := model.WorkflowSearchRequest{
		Query: query,
		Size:  1,
	}

	rawJSON, err := DefaultAtlanClient.CallAPI(&WORKFLOW_INDEX_SEARCH, nil, &request)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowSearchResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	if len(response.Hits.Hits) > 0 {
		return &response.Hits.Hits[0], nil
	}
	return nil, nil
}

// FindRunByID searches for a workflow run by its ID.
func (w *WorkflowClient) FindRunByID(id string) (*structs.WorkflowSearchResult, error) {
	var query model.Query = &model.BoolQuery{
		Filter: []model.Query{
			&model.TermQuery{
				Field: "_id",
				Value: id,
			},
		},
	}

	response, err := w.findRuns(query, 0, 1)
	if err != nil {
		return nil, err
	}

	if len(response.Hits.Hits) > 0 {
		return &response.Hits.Hits[0], nil
	}
	return nil, nil
}

// findRuns retrieves existing workflow runs.
func (w *WorkflowClient) findRuns(query model.Query, from, size int) (*structs.WorkflowSearchResponse, error) {
	request := model.WorkflowSearchRequest{
		Query: query,
		From:  from,
		Size:  size,
	}

	rawJSON, err := DefaultAtlanClient.CallAPI(&WORKFLOW_INDEX_RUN_SEARCH, nil, &request)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowSearchResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// FindLatestRun retrieves the latest run of a given workflow.
func (w *WorkflowClient) FindLatestRun(workflowName string) (*structs.WorkflowSearchResult, error) {
	var query model.Query = &model.BoolQuery{
		Filter: []model.Query{
			&model.NestedQuery{
				Path: "spec",
				Query: &model.TermQuery{
					Field: "spec.workflowTemplateRef.name.keyword",
					Value: workflowName,
				},
			},
		},
	}

	response, err := w.findRuns(query, 0, 1)
	if err != nil {
		return nil, err
	}

	if len(response.Hits.Hits) > 0 {
		return &response.Hits.Hits[0], nil
	}
	return nil, nil
}

// FindCurrentRun retrieves the most current, still-running workflow.
func (w *WorkflowClient) FindCurrentRun(workflowName string) (*structs.WorkflowSearchResult, error) {
	var query model.Query = &model.BoolQuery{
		Filter: []model.Query{
			&model.NestedQuery{
				Path: "spec",
				Query: &model.TermQuery{
					Field: "spec.workflowTemplateRef.name.keyword",
					Value: workflowName,
				},
			},
		},
	}

	response, err := w.findRuns(query, 0, 50)
	if err != nil {
		return nil, err
	}

	for _, result := range response.Hits.Hits {
		if *result.Source.Status.Phase == atlan.AtlanWorkflowPhasePending.Name || *result.Source.Status.Phase == atlan.AtlanWorkflowPhaseRunning.Name {
			return &result, nil
		}
	}
	return nil, nil
}

// GetRuns retrieves all workflow runs filtered by workflow name and phase.
func (w *WorkflowClient) GetRuns(workflowName string, workflowPhase atlan.AtlanWorkflowPhase, from, size int) ([]structs.WorkflowSearchResult, error) {
	var query model.Query = &model.BoolQuery{
		Must: []model.Query{
			&model.NestedQuery{
				Path: "spec",
				Query: &model.TermQuery{
					Field: "spec.workflowTemplateRef.name.keyword",
					Value: workflowName,
				},
			},
		},
		Filter: []model.Query{
			&model.TermQuery{
				Field: "status.phase.keyword",
				Value: workflowPhase.Name,
			},
		},
	}

	response, err := w.findRuns(query, from, size)
	if err != nil {
		return nil, err
	}

	return response.Hits.Hits, nil
}

// Stop stops a running workflow.
func (w *WorkflowClient) Stop(workflowRunID string) (*structs.WorkflowRunResponse, error) {

	api := &STOP_WORKFLOW_RUN
	api.Path = fmt.Sprintf("runs/%s/stop", workflowRunID)

	rawJSON, err := DefaultAtlanClient.CallAPI(api, nil, "")
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowRunResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Delete archives (deletes) the provided workflow.
func (w *WorkflowClient) Delete(workflowName string) error {
	api := &WORKFLOW_ARCHIVE
	api.Path = fmt.Sprintf("workflows/%s/archive", workflowName)
	_, err := DefaultAtlanClient.CallAPI(api, nil, "")
	return err
}

// handleWorkflowTypes determines the workflow details based on its type.
func (w *WorkflowClient) handleWorkflowTypes(workflow interface{}) (*structs.WorkflowSearchResultDetail, error) {
	switch wf := workflow.(type) {
	case atlan.WorkflowPackage:
		results, err := w.FindByType(wf, 1) // Fetching at most 1 result
		if err != nil {
			return nil, err
		}
		if len(results) == 0 {
			return nil, fmt.Errorf("no prior run available for workflow: %v", wf)
		}
		return &results[0].Source, nil

	case structs.WorkflowSearchResult:
		return &wf.Source, nil

	case *structs.WorkflowSearchResultDetail:
		return wf, nil

	default:
		return nil, fmt.Errorf("invalid workflow type provided")
	}
}

// Rerun executes the workflow immediately if it has been run before.
// If idempotent is true, it only reruns if not already running.
func (w *WorkflowClient) Rerun(workflow interface{}, idempotent bool) (*structs.WorkflowRunResponse, error) {
	detail, err := w.handleWorkflowTypes(workflow)
	if err != nil {
		return nil, err
	}

	if idempotent && *detail.Metadata.Name != "" {
		// Wait before checking the current workflow run status
		time.Sleep(10 * time.Second)

		currentRun, err := w.FindCurrentRun(*detail.Metadata.Name)
		if err == nil && currentRun != nil && currentRun.Source.Status != nil {
			return &structs.WorkflowRunResponse{
				WorkflowResponse: structs.WorkflowResponse{
					Metadata: &currentRun.Source.Metadata,
					Spec:     &currentRun.Source.Spec,
				},
				Status: *currentRun.Source.Status,
			}, nil
		}
	}

	request := structs.ReRunRequest{
		Namespace:    *detail.Metadata.Namespace,
		ResourceName: *detail.Metadata.Name,
	}

	rawJSON, err := DefaultAtlanClient.CallAPI(&WORKFLOW_RERUN, nil, &request)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowRunResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// Update modifies the configuration of an existing workflow.
func (w *WorkflowClient) Update(workflow *structs.Workflow) (*structs.WorkflowResponse, error) {

	api := &WORKFLOW_UPDATE
	api.Path = fmt.Sprintf("workflows/%s", *workflow.Metadata.Name)

	rawJSON, err := DefaultAtlanClient.CallAPI(api, nil, workflow)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateOwner assigns a new owner to the workflow.
func (w *WorkflowClient) UpdateOwner(workflowName, username string) (*structs.WorkflowResponse, error) {

	api := &WORKFLOW_CHANGE_OWNER
	api.Path = fmt.Sprintf("workflows/%s/changeownership", workflowName)

	queryParams := map[string]string{"username": username}

	rawJSON, err := DefaultAtlanClient.CallAPI(api, queryParams, nil)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
