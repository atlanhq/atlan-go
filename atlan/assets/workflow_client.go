package assets

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/atlanhq/atlan-go/atlan"
	"github.com/atlanhq/atlan-go/atlan/model"
	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

const (
	workflowRunSchedule = "orchestration.atlan.com/schedule"
	workflowRunTimezone = "orchestration.atlan.com/timezone"
	MonitorSleepSeconds = 5
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
		if *result.Source.Status.Phase == atlan.AtlanWorkflowPhasePending || *result.Source.Status.Phase == atlan.AtlanWorkflowPhaseRunning {
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

// Methods related to workflow schedules

// Monitor the status of the workflow's run.
func (w *WorkflowClient) Monitor(workflowResponse *structs.WorkflowResponse, logger *log.Logger) (*atlan.AtlanWorkflowPhase, error) {
	if workflowResponse.Metadata == nil || *workflowResponse.Metadata.Name == "" {
		if logger != nil {
			logger.Println("Skipping workflow monitoring — nothing to monitor.")
		}
		return nil, nil
	}

	name := workflowResponse.Metadata.Name
	var status *atlan.AtlanWorkflowPhase

	for status == nil || (*status != atlan.AtlanWorkflowPhaseSuccess && *status != atlan.AtlanWorkflowPhaseError && *status != atlan.AtlanWorkflowPhaseFailed) {
		time.Sleep(MonitorSleepSeconds * time.Second)
		runDetails, _ := w.FindLatestRun(*name)
		if runDetails != nil {
			status = runDetails.Status()
		}
		if logger != nil {
			logger.Printf("Workflow status: %s\n", status)
		}
	}
	if logger != nil {
		logger.Printf("Workflow completion status: %s\n", status)
	}
	return status, nil
}

func (w *WorkflowClient) addSchedule(workflow *structs.WorkflowSearchResultDetail, schedule *structs.WorkflowSchedule) {
	if workflow.Metadata.Annotations == nil {
		workflow.Metadata.Annotations = make(map[string]string)
	}
	workflow.Metadata.Annotations[workflowRunSchedule] = schedule.CronSchedule
	workflow.Metadata.Annotations[workflowRunTimezone] = schedule.Timezone
}

func (w *WorkflowClient) AddSchedule(workflow interface{}, schedule *structs.WorkflowSchedule) (*structs.WorkflowResponse, error) {
	workflowToUpdate, err := w.handleWorkflowTypes(workflow)
	if err != nil {
		return nil, err
	}

	w.addSchedule(workflowToUpdate, schedule)

	api := &WORKFLOW_UPDATE
	api.Path = fmt.Sprintf("workflows/%s", *workflowToUpdate.Metadata.Name)

	rawJSON, err := DefaultAtlanClient.CallAPI(api, nil, workflowToUpdate)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (w *WorkflowClient) RemoveSchedule(workflow interface{}) (*structs.WorkflowResponse, error) {
	workflowToUpdate, err := w.handleWorkflowTypes(workflow)
	if err != nil {
		return nil, err
	}

	if workflowToUpdate.Metadata.Annotations != nil {
		delete(workflowToUpdate.Metadata.Annotations, workflowRunSchedule)
	}

	api := &WORKFLOW_UPDATE
	api.Path = fmt.Sprintf("workflows/%s", *workflowToUpdate.Metadata.Name)

	rawJSON, err := DefaultAtlanClient.CallAPI(api, nil, workflowToUpdate)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (w *WorkflowClient) GetAllScheduledRuns() ([]structs.WorkflowScheduleResponse, error) {
	rawJSON, err := DefaultAtlanClient.CallAPI(&GET_ALL_SCHEDULE_RUNS, nil, nil)
	if err != nil {
		return nil, err
	}

	var response []structs.WorkflowScheduleResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (w *WorkflowClient) GetScheduledRun(workflowName string) (*structs.WorkflowScheduleResponse, error) {
	api := &GET_SCHEDULE_RUN
	api.Path = fmt.Sprintf("schedules/%s/cron", workflowName)

	rawJSON, err := DefaultAtlanClient.CallAPI(api, nil, nil)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowScheduleResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (w *WorkflowClient) FindScheduleQuery(savedQueryID string, maxResults int) ([]structs.WorkflowSearchResult, error) {
	if maxResults <= 0 {
		maxResults = 10
	}

	var query model.Query = &model.BoolQuery{
		Filter: []model.Query{
			&model.NestedQuery{
				Path: "metadata",
				Query: &model.PrefixQuery{
					Field: "metadata.name.keyword",
					Value: fmt.Sprintf("asq-%s", savedQueryID),
				},
			},
			&model.NestedQuery{
				Path: "metadata",
				Query: &model.TermQuery{
					Field: "metadata.annotations.package.argoproj.io/name.keyword",
					Value: "@atlan/schedule-query",
				},
			},
		},
	}

	request := model.WorkflowSearchRequest{
		Query: query,
		Size:  maxResults,
	}

	rawJSON, err := DefaultAtlanClient.CallAPI(&WORKFLOW_INDEX_SEARCH, nil, request)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowSearchResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}

	if response.Hits.Hits != nil {
		return response.Hits.Hits, nil
	}
	return nil, nil
}

func (w *WorkflowClient) ReRunScheduleQuery(scheduleQueryID string) (*structs.WorkflowRunResponse, error) {
	request := structs.ReRunRequest{
		Namespace:    "default",
		ResourceName: scheduleQueryID,
	}

	rawJSON, err := DefaultAtlanClient.CallAPI(&WORKFLOW_OWNER_RERUN, nil, &request)
	if err != nil {
		return nil, err
	}

	var response structs.WorkflowRunResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (w *WorkflowClient) FindScheduleQueryBetween(request structs.ScheduleQueriesSearchRequest, missed bool) ([]structs.WorkflowRunResponse, error) {
	queryParams := map[string]string{
		"startDate": request.StartDate,
		"endDate":   request.EndDate,
	}

	searchAPI := SCHEDULE_QUERY_WORKFLOWS_SEARCH
	if missed {
		searchAPI = SCHEDULE_QUERY_WORKFLOWS_MISSED
	}

	rawJSON, err := DefaultAtlanClient.CallAPI(&searchAPI, queryParams, nil)
	if err != nil {
		return nil, err
	}

	var response []structs.WorkflowRunResponse
	if err := json.Unmarshal(rawJSON, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// Run executes an Atlan workflow with an optional schedule.
func (w *WorkflowClient) Run(workflow interface{}, schedule *structs.WorkflowSchedule) (*structs.WorkflowResponse, error) {
	if workflow == nil {
		return nil, errors.New("workflow cannot be nil")
	}

	if schedule != nil {
		_, err := w.AddSchedule(workflow, schedule)
		if err != nil {
			return nil, err
		}
	}

	responseData, err := DefaultAtlanClient.CallAPI(&WORKFLOW_RUN, nil, workflow)
	if err != nil {
		return nil, fmt.Errorf("error executing workflow: %w", err)
	}

	var workflowResponse structs.WorkflowResponse
	if err := json.Unmarshal(responseData, &workflowResponse); err != nil {
		return nil, fmt.Errorf("failed to parse workflow response: %w", err)
	}

	return &workflowResponse, nil
}
