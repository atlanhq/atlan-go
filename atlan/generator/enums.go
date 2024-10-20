// # **************************************
// # CODE BELOW IS GENERATED NOT MODIFY  **
// # **************************************

package generator

import (
	"encoding/json"
)

type WorkflowRunType struct {
	Name string
}

func (a WorkflowRunType) String() string {
	return a.Name
}

var (
	WorkflowRunTypeData_access            = WorkflowRunType{"DATA_ACCESS"}
	WorkflowRunTypePolicy                 = WorkflowRunType{"POLICY"}
	WorkflowRunTypeChange_management      = WorkflowRunType{"CHANGE_MANAGEMENT"}
	WorkflowRunTypePublication_management = WorkflowRunType{"PUBLICATION_MANAGEMENT"}
	WorkflowRunTypeImpact_analysis        = WorkflowRunType{"IMPACT_ANALYSIS"}
)

// UnmarshalJSON customizes the unmarshalling of a WorkflowRunType from JSON.
func (c *WorkflowRunType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DATA_ACCESS":
		*c = WorkflowRunTypeData_access

	case "POLICY":
		*c = WorkflowRunTypePolicy

	case "CHANGE_MANAGEMENT":
		*c = WorkflowRunTypeChange_management

	case "PUBLICATION_MANAGEMENT":
		*c = WorkflowRunTypePublication_management

	case "IMPACT_ANALYSIS":
		*c = WorkflowRunTypeImpact_analysis
	default:
		*c = WorkflowRunType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a WorkflowRunType to JSON.
func (c WorkflowRunType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSLeaseStatus struct {
	Name string
}

func (a ADLSLeaseStatus) String() string {
	return a.Name
}

var (
	ADLSLeaseStatusLocked   = ADLSLeaseStatus{"Locked"}
	ADLSLeaseStatusUnlocked = ADLSLeaseStatus{"Unlocked"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSLeaseStatus from JSON.
func (c *ADLSLeaseStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Locked":
		*c = ADLSLeaseStatusLocked

	case "Unlocked":
		*c = ADLSLeaseStatusUnlocked
	default:
		*c = ADLSLeaseStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSLeaseStatus to JSON.
func (c ADLSLeaseStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type APIQueryParamTypeEnum struct {
	Name string
}

func (a APIQueryParamTypeEnum) String() string {
	return a.Name
}

var ( // This field is an input parameter to a Query.
	APIQueryParamTypeEnumInput  = APIQueryParamTypeEnum{"Input"} // This field is an output parameter to a Query.
	APIQueryParamTypeEnumOutput = APIQueryParamTypeEnum{"Output"}
)

// UnmarshalJSON customizes the unmarshalling of a APIQueryParamTypeEnum from JSON.
func (c *APIQueryParamTypeEnum) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Input":
		*c = APIQueryParamTypeEnumInput

	case "Output":
		*c = APIQueryParamTypeEnumOutput
	default:
		*c = APIQueryParamTypeEnum{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a APIQueryParamTypeEnum to JSON.
func (c APIQueryParamTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type DataProductStatus struct {
	Name string
}

func (a DataProductStatus) String() string {
	return a.Name
}

var (
	DataProductStatusActive   = DataProductStatus{"Active"}
	DataProductStatusSunset   = DataProductStatus{"Sunset"}
	DataProductStatusArchived = DataProductStatus{"Archived"}
)

// UnmarshalJSON customizes the unmarshalling of a DataProductStatus from JSON.
func (c *DataProductStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Active":
		*c = DataProductStatusActive

	case "Sunset":
		*c = DataProductStatusSunset

	case "Archived":
		*c = DataProductStatusArchived
	default:
		*c = DataProductStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a DataProductStatus to JSON.
func (c DataProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type table_type struct {
	Name string
}

func (a table_type) String() string {
	return a.Name
}

var (
	table_typeTemporary = table_type{"TEMPORARY"}
	table_typeIceberg   = table_type{"ICEBERG"}
)

// UnmarshalJSON customizes the unmarshalling of a table_type from JSON.
func (c *table_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "TEMPORARY":
		*c = table_typeTemporary

	case "ICEBERG":
		*c = table_typeIceberg
	default:
		*c = table_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a table_type to JSON.
func (c table_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type DynamoDBSecondaryIndexProjectionType struct {
	Name string
}

func (a DynamoDBSecondaryIndexProjectionType) String() string {
	return a.Name
}

var ( // Only the index and primary keys are projected into the index.
	DynamoDBSecondaryIndexProjectionTypeKeys_only = DynamoDBSecondaryIndexProjectionType{"KEYS_ONLY"} // In addition to the attributes described in KEYS_ONLY, the secondary index will include other non-key attributes that you specify.
	DynamoDBSecondaryIndexProjectionTypeInclude   = DynamoDBSecondaryIndexProjectionType{"INCLUDE"}   // All of the table attributes are projected into the index.
	DynamoDBSecondaryIndexProjectionTypeAll       = DynamoDBSecondaryIndexProjectionType{"ALL"}
)

// UnmarshalJSON customizes the unmarshalling of a DynamoDBSecondaryIndexProjectionType from JSON.
func (c *DynamoDBSecondaryIndexProjectionType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "KEYS_ONLY":
		*c = DynamoDBSecondaryIndexProjectionTypeKeys_only

	case "INCLUDE":
		*c = DynamoDBSecondaryIndexProjectionTypeInclude

	case "ALL":
		*c = DynamoDBSecondaryIndexProjectionTypeAll
	default:
		*c = DynamoDBSecondaryIndexProjectionType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a DynamoDBSecondaryIndexProjectionType to JSON.
func (c DynamoDBSecondaryIndexProjectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSObjectType struct {
	Name string
}

func (a ADLSObjectType) String() string {
	return a.Name
}

var (
	ADLSObjectTypeBlockblob  = ADLSObjectType{"BlockBlob"}
	ADLSObjectTypePageblob   = ADLSObjectType{"PageBlob"}
	ADLSObjectTypeAppendblob = ADLSObjectType{"AppendBlob"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSObjectType from JSON.
func (c *ADLSObjectType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "BlockBlob":
		*c = ADLSObjectTypeBlockblob

	case "PageBlob":
		*c = ADLSObjectTypePageblob

	case "AppendBlob":
		*c = ADLSObjectTypeAppendblob
	default:
		*c = ADLSObjectType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSObjectType to JSON.
func (c ADLSObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type FivetranConnectorStatus struct {
	Name string
}

func (a FivetranConnectorStatus) String() string {
	return a.Name
}

var (
	FivetranConnectorStatusSuccessful        = FivetranConnectorStatus{"SUCCESSFUL"}
	FivetranConnectorStatusFailure           = FivetranConnectorStatus{"FAILURE"}
	FivetranConnectorStatusFailure_with_task = FivetranConnectorStatus{"FAILURE_WITH_TASK"}
	FivetranConnectorStatusRescheduled       = FivetranConnectorStatus{"RESCHEDULED"}
	FivetranConnectorStatusNo_status         = FivetranConnectorStatus{"NO_STATUS"}
)

// UnmarshalJSON customizes the unmarshalling of a FivetranConnectorStatus from JSON.
func (c *FivetranConnectorStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "SUCCESSFUL":
		*c = FivetranConnectorStatusSuccessful

	case "FAILURE":
		*c = FivetranConnectorStatusFailure

	case "FAILURE_WITH_TASK":
		*c = FivetranConnectorStatusFailure_with_task

	case "RESCHEDULED":
		*c = FivetranConnectorStatusRescheduled

	case "NO_STATUS":
		*c = FivetranConnectorStatusNo_status
	default:
		*c = FivetranConnectorStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a FivetranConnectorStatus to JSON.
func (c FivetranConnectorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSObjectArchiveStatus struct {
	Name string
}

func (a ADLSObjectArchiveStatus) String() string {
	return a.Name
}

var (
	ADLSObjectArchiveStatusRehydrate_pending_to_hot  = ADLSObjectArchiveStatus{"rehydrate-pending-to-hot"}
	ADLSObjectArchiveStatusRehydrate_pending_to_cool = ADLSObjectArchiveStatus{"rehydrate-pending-to-cool"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSObjectArchiveStatus from JSON.
func (c *ADLSObjectArchiveStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "rehydrate-pending-to-hot":
		*c = ADLSObjectArchiveStatusRehydrate_pending_to_hot

	case "rehydrate-pending-to-cool":
		*c = ADLSObjectArchiveStatusRehydrate_pending_to_cool
	default:
		*c = ADLSObjectArchiveStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSObjectArchiveStatus to JSON.
func (c ADLSObjectArchiveStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type powerbi_endorsement struct {
	Name string
}

func (a powerbi_endorsement) String() string {
	return a.Name
}

var (
	powerbi_endorsementPromoted  = powerbi_endorsement{"Promoted"}
	powerbi_endorsementCertified = powerbi_endorsement{"Certified"}
)

// UnmarshalJSON customizes the unmarshalling of a powerbi_endorsement from JSON.
func (c *powerbi_endorsement) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Promoted":
		*c = powerbi_endorsementPromoted

	case "Certified":
		*c = powerbi_endorsementCertified
	default:
		*c = powerbi_endorsement{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a powerbi_endorsement to JSON.
func (c powerbi_endorsement) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AtlasGlossaryType struct {
	Name string
}

func (a AtlasGlossaryType) String() string {
	return a.Name
}

var ( // Glossary will be used to store knowledge as documents
	AtlasGlossaryTypeKnowledge_hub = AtlasGlossaryType{"KNOWLEDGE_HUB"}
)

// UnmarshalJSON customizes the unmarshalling of a AtlasGlossaryType from JSON.
func (c *AtlasGlossaryType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "KNOWLEDGE_HUB":
		*c = AtlasGlossaryTypeKnowledge_hub
	default:
		*c = AtlasGlossaryType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AtlasGlossaryType to JSON.
func (c AtlasGlossaryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type kafka_topic_cleanup_policy struct {
	Name string
}

func (a kafka_topic_cleanup_policy) String() string {
	return a.Name
}

var (
	kafka_topic_cleanup_policyCompact = kafka_topic_cleanup_policy{"compact"}
	kafka_topic_cleanup_policyDelete  = kafka_topic_cleanup_policy{"delete"}
)

// UnmarshalJSON customizes the unmarshalling of a kafka_topic_cleanup_policy from JSON.
func (c *kafka_topic_cleanup_policy) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "compact":
		*c = kafka_topic_cleanup_policyCompact

	case "delete":
		*c = kafka_topic_cleanup_policyDelete
	default:
		*c = kafka_topic_cleanup_policy{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a kafka_topic_cleanup_policy to JSON.
func (c kafka_topic_cleanup_policy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type query_username_strategy struct {
	Name string
}

func (a query_username_strategy) String() string {
	return a.Name
}

var (
	query_username_strategyConnectionusername = query_username_strategy{"connectionUsername"}
	query_username_strategyAtlanusername      = query_username_strategy{"atlanUsername"}
)

// UnmarshalJSON customizes the unmarshalling of a query_username_strategy from JSON.
func (c *query_username_strategy) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "connectionUsername":
		*c = query_username_strategyConnectionusername

	case "atlanUsername":
		*c = query_username_strategyAtlanusername
	default:
		*c = query_username_strategy{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a query_username_strategy to JSON.
func (c query_username_strategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type DataProductCriticality struct {
	Name string
}

func (a DataProductCriticality) String() string {
	return a.Name
}

var (
	DataProductCriticalityLow    = DataProductCriticality{"Low"}
	DataProductCriticalityMedium = DataProductCriticality{"Medium"}
	DataProductCriticalityHigh   = DataProductCriticality{"High"}
)

// UnmarshalJSON customizes the unmarshalling of a DataProductCriticality from JSON.
func (c *DataProductCriticality) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Low":
		*c = DataProductCriticalityLow

	case "Medium":
		*c = DataProductCriticalityMedium

	case "High":
		*c = DataProductCriticalityHigh
	default:
		*c = DataProductCriticality{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a DataProductCriticality to JSON.
func (c DataProductCriticality) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type quick_sight_dataset_field_type struct {
	Name string
}

func (a quick_sight_dataset_field_type) String() string {
	return a.Name
}

var (
	quick_sight_dataset_field_typeString   = quick_sight_dataset_field_type{"STRING"}
	quick_sight_dataset_field_typeInteger  = quick_sight_dataset_field_type{"INTEGER"}
	quick_sight_dataset_field_typeDecimal  = quick_sight_dataset_field_type{"DECIMAL"}
	quick_sight_dataset_field_typeDatetime = quick_sight_dataset_field_type{"DATETIME"}
)

// UnmarshalJSON customizes the unmarshalling of a quick_sight_dataset_field_type from JSON.
func (c *quick_sight_dataset_field_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "STRING":
		*c = quick_sight_dataset_field_typeString

	case "INTEGER":
		*c = quick_sight_dataset_field_typeInteger

	case "DECIMAL":
		*c = quick_sight_dataset_field_typeDecimal

	case "DATETIME":
		*c = quick_sight_dataset_field_typeDatetime
	default:
		*c = quick_sight_dataset_field_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a quick_sight_dataset_field_type to JSON.
func (c quick_sight_dataset_field_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type matillion_job_type struct {
	Name string
}

func (a matillion_job_type) String() string {
	return a.Name
}

var (
	matillion_job_typeOrchestration  = matillion_job_type{"ORCHESTRATION"}
	matillion_job_typeTransformation = matillion_job_type{"TRANSFORMATION"}
)

// UnmarshalJSON customizes the unmarshalling of a matillion_job_type from JSON.
func (c *matillion_job_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "ORCHESTRATION":
		*c = matillion_job_typeOrchestration

	case "TRANSFORMATION":
		*c = matillion_job_typeTransformation
	default:
		*c = matillion_job_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a matillion_job_type to JSON.
func (c matillion_job_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSAccountStatus struct {
	Name string
}

func (a ADLSAccountStatus) String() string {
	return a.Name
}

var (
	ADLSAccountStatusAvailable   = ADLSAccountStatus{"Available"}
	ADLSAccountStatusUnavailable = ADLSAccountStatus{"Unavailable"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSAccountStatus from JSON.
func (c *ADLSAccountStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Available":
		*c = ADLSAccountStatusAvailable

	case "Unavailable":
		*c = ADLSAccountStatusUnavailable
	default:
		*c = ADLSAccountStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSAccountStatus to JSON.
func (c ADLSAccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type atlas_operation struct {
	Name string
}

func (a atlas_operation) String() string {
	return a.Name
}

var (
	atlas_operationOthers              = atlas_operation{"OTHERS"}
	atlas_operationPurge               = atlas_operation{"PURGE"}
	atlas_operationExport              = atlas_operation{"EXPORT"}
	atlas_operationImport              = atlas_operation{"IMPORT"}
	atlas_operationImport_delete_repl  = atlas_operation{"IMPORT_DELETE_REPL"}
	atlas_operationType_def_create     = atlas_operation{"TYPE_DEF_CREATE"}
	atlas_operationType_def_update     = atlas_operation{"TYPE_DEF_UPDATE"}
	atlas_operationType_def_delete     = atlas_operation{"TYPE_DEF_DELETE"}
	atlas_operationServer_start        = atlas_operation{"SERVER_START"}
	atlas_operationServer_state_active = atlas_operation{"SERVER_STATE_ACTIVE"}
)

// UnmarshalJSON customizes the unmarshalling of a atlas_operation from JSON.
func (c *atlas_operation) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "OTHERS":
		*c = atlas_operationOthers

	case "PURGE":
		*c = atlas_operationPurge

	case "EXPORT":
		*c = atlas_operationExport

	case "IMPORT":
		*c = atlas_operationImport

	case "IMPORT_DELETE_REPL":
		*c = atlas_operationImport_delete_repl

	case "TYPE_DEF_CREATE":
		*c = atlas_operationType_def_create

	case "TYPE_DEF_UPDATE":
		*c = atlas_operationType_def_update

	case "TYPE_DEF_DELETE":
		*c = atlas_operationType_def_delete

	case "SERVER_START":
		*c = atlas_operationServer_start

	case "SERVER_STATE_ACTIVE":
		*c = atlas_operationServer_state_active
	default:
		*c = atlas_operation{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a atlas_operation to JSON.
func (c atlas_operation) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AdfActivityState struct {
	Name string
}

func (a AdfActivityState) String() string {
	return a.Name
}

var (
	AdfActivityStateActive   = AdfActivityState{"ACTIVE"}
	AdfActivityStateInactive = AdfActivityState{"INACTIVE"}
)

// UnmarshalJSON customizes the unmarshalling of a AdfActivityState from JSON.
func (c *AdfActivityState) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "ACTIVE":
		*c = AdfActivityStateActive

	case "INACTIVE":
		*c = AdfActivityStateInactive
	default:
		*c = AdfActivityState{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AdfActivityState to JSON.
func (c AdfActivityState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSProvisionState struct {
	Name string
}

func (a ADLSProvisionState) String() string {
	return a.Name
}

var (
	ADLSProvisionStateCreating     = ADLSProvisionState{"Creating"}
	ADLSProvisionStateResolvingdns = ADLSProvisionState{"ResolvingDNS"}
	ADLSProvisionStateSucceeded    = ADLSProvisionState{"Succeeded"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSProvisionState from JSON.
func (c *ADLSProvisionState) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Creating":
		*c = ADLSProvisionStateCreating

	case "ResolvingDNS":
		*c = ADLSProvisionStateResolvingdns

	case "Succeeded":
		*c = ADLSProvisionStateSucceeded
	default:
		*c = ADLSProvisionState{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSProvisionState to JSON.
func (c ADLSProvisionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type SourceCostUnitType struct {
	Name string
}

func (a SourceCostUnitType) String() string {
	return a.Name
}

var (
	SourceCostUnitTypeCredits = SourceCostUnitType{"Credits"}
	SourceCostUnitTypeBytes   = SourceCostUnitType{"bytes"}
	SourceCostUnitTypeSlot_ms = SourceCostUnitType{"slot-ms"}
)

// UnmarshalJSON customizes the unmarshalling of a SourceCostUnitType from JSON.
func (c *SourceCostUnitType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Credits":
		*c = SourceCostUnitTypeCredits

	case "bytes":
		*c = SourceCostUnitTypeBytes

	case "slot-ms":
		*c = SourceCostUnitTypeSlot_ms
	default:
		*c = SourceCostUnitType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a SourceCostUnitType to JSON.
func (c SourceCostUnitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type DomoCardType struct {
	Name string
}

func (a DomoCardType) String() string {
	return a.Name
}

var (
	DomoCardTypeDoc        = DomoCardType{"DOC"}
	DomoCardTypeDoc_card   = DomoCardType{"DOC CARD"}
	DomoCardTypeChart      = DomoCardType{"CHART"}
	DomoCardTypeDrill_view = DomoCardType{"DRILL VIEW"}
	DomoCardTypeNotebook   = DomoCardType{"NOTEBOOK"}
)

// UnmarshalJSON customizes the unmarshalling of a DomoCardType from JSON.
func (c *DomoCardType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DOC":
		*c = DomoCardTypeDoc

	case "DOC CARD":
		*c = DomoCardTypeDoc_card

	case "CHART":
		*c = DomoCardTypeChart

	case "DRILL VIEW":
		*c = DomoCardTypeDrill_view

	case "NOTEBOOK":
		*c = DomoCardTypeNotebook
	default:
		*c = DomoCardType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a DomoCardType to JSON.
func (c DomoCardType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type SchemaRegistrySchemaCompatibility struct {
	Name string
}

func (a SchemaRegistrySchemaCompatibility) String() string {
	return a.Name
}

var (
	SchemaRegistrySchemaCompatibilityBackward            = SchemaRegistrySchemaCompatibility{"BACKWARD"}
	SchemaRegistrySchemaCompatibilityBackward_transitive = SchemaRegistrySchemaCompatibility{"BACKWARD_TRANSITIVE"}
	SchemaRegistrySchemaCompatibilityForward             = SchemaRegistrySchemaCompatibility{"FORWARD"}
	SchemaRegistrySchemaCompatibilityForward_transitive  = SchemaRegistrySchemaCompatibility{"FORWARD_TRANSITIVE"}
	SchemaRegistrySchemaCompatibilityFull                = SchemaRegistrySchemaCompatibility{"FULL"}
	SchemaRegistrySchemaCompatibilityFull_transitive     = SchemaRegistrySchemaCompatibility{"FULL_TRANSITIVE"}
	SchemaRegistrySchemaCompatibilityNone                = SchemaRegistrySchemaCompatibility{"NONE"}
)

// UnmarshalJSON customizes the unmarshalling of a SchemaRegistrySchemaCompatibility from JSON.
func (c *SchemaRegistrySchemaCompatibility) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "BACKWARD":
		*c = SchemaRegistrySchemaCompatibilityBackward

	case "BACKWARD_TRANSITIVE":
		*c = SchemaRegistrySchemaCompatibilityBackward_transitive

	case "FORWARD":
		*c = SchemaRegistrySchemaCompatibilityForward

	case "FORWARD_TRANSITIVE":
		*c = SchemaRegistrySchemaCompatibilityForward_transitive

	case "FULL":
		*c = SchemaRegistrySchemaCompatibilityFull

	case "FULL_TRANSITIVE":
		*c = SchemaRegistrySchemaCompatibilityFull_transitive

	case "NONE":
		*c = SchemaRegistrySchemaCompatibilityNone
	default:
		*c = SchemaRegistrySchemaCompatibility{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a SchemaRegistrySchemaCompatibility to JSON.
func (c SchemaRegistrySchemaCompatibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSLeaseState struct {
	Name string
}

func (a ADLSLeaseState) String() string {
	return a.Name
}

var (
	ADLSLeaseStateAvailable = ADLSLeaseState{"Available"}
	ADLSLeaseStateLeased    = ADLSLeaseState{"Leased"}
	ADLSLeaseStateExpired   = ADLSLeaseState{"Expired"}
	ADLSLeaseStateBreaking  = ADLSLeaseState{"Breaking"}
	ADLSLeaseStateBroken    = ADLSLeaseState{"Broken"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSLeaseState from JSON.
func (c *ADLSLeaseState) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Available":
		*c = ADLSLeaseStateAvailable

	case "Leased":
		*c = ADLSLeaseStateLeased

	case "Expired":
		*c = ADLSLeaseStateExpired

	case "Breaking":
		*c = ADLSLeaseStateBreaking

	case "Broken":
		*c = ADLSLeaseStateBroken
	default:
		*c = ADLSLeaseState{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSLeaseState to JSON.
func (c ADLSLeaseState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type WorkflowStatus struct {
	Name string
}

func (a WorkflowStatus) String() string {
	return a.Name
}

var (
	WorkflowStatusPublished = WorkflowStatus{"PUBLISHED"}
	WorkflowStatusDraft     = WorkflowStatus{"DRAFT"}
	WorkflowStatusDisabled  = WorkflowStatus{"DISABLED"}
)

// UnmarshalJSON customizes the unmarshalling of a WorkflowStatus from JSON.
func (c *WorkflowStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "PUBLISHED":
		*c = WorkflowStatusPublished

	case "DRAFT":
		*c = WorkflowStatusDraft

	case "DISABLED":
		*c = WorkflowStatusDisabled
	default:
		*c = WorkflowStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a WorkflowStatus to JSON.
func (c WorkflowStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type google_datastudio_asset_type struct {
	Name string
}

func (a google_datastudio_asset_type) String() string {
	return a.Name
}

var (
	google_datastudio_asset_typeReport      = google_datastudio_asset_type{"REPORT"}
	google_datastudio_asset_typeData_source = google_datastudio_asset_type{"DATA_SOURCE"}
)

// UnmarshalJSON customizes the unmarshalling of a google_datastudio_asset_type from JSON.
func (c *google_datastudio_asset_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "REPORT":
		*c = google_datastudio_asset_typeReport

	case "DATA_SOURCE":
		*c = google_datastudio_asset_typeData_source
	default:
		*c = google_datastudio_asset_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a google_datastudio_asset_type to JSON.
func (c google_datastudio_asset_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type quick_sight_dataset_import_mode struct {
	Name string
}

func (a quick_sight_dataset_import_mode) String() string {
	return a.Name
}

var (
	quick_sight_dataset_import_modeSpice        = quick_sight_dataset_import_mode{"SPICE"}
	quick_sight_dataset_import_modeDirect_query = quick_sight_dataset_import_mode{"DIRECT_QUERY"}
)

// UnmarshalJSON customizes the unmarshalling of a quick_sight_dataset_import_mode from JSON.
func (c *quick_sight_dataset_import_mode) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "SPICE":
		*c = quick_sight_dataset_import_modeSpice

	case "DIRECT_QUERY":
		*c = quick_sight_dataset_import_modeDirect_query
	default:
		*c = quick_sight_dataset_import_mode{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a quick_sight_dataset_import_mode to JSON.
func (c quick_sight_dataset_import_mode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type MongoDBCollectionValidationLevel struct {
	Name string
}

func (a MongoDBCollectionValidationLevel) String() string {
	return a.Name
}

var ( // OFF means no validation for inserts or updates
	MongoDBCollectionValidationLevelOff      = MongoDBCollectionValidationLevel{"OFF"}    // STRICT means apply validation rules to all inserts and all updates
	MongoDBCollectionValidationLevelStrict   = MongoDBCollectionValidationLevel{"STRICT"} // MODERATE means apply validation rules to inserts and to updates on existing valid documents. Do not apply rules to updates on existing invalid documents.
	MongoDBCollectionValidationLevelModerate = MongoDBCollectionValidationLevel{"MODERATE"}
)

// UnmarshalJSON customizes the unmarshalling of a MongoDBCollectionValidationLevel from JSON.
func (c *MongoDBCollectionValidationLevel) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "OFF":
		*c = MongoDBCollectionValidationLevelOff

	case "STRICT":
		*c = MongoDBCollectionValidationLevelStrict

	case "MODERATE":
		*c = MongoDBCollectionValidationLevelModerate
	default:
		*c = MongoDBCollectionValidationLevel{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a MongoDBCollectionValidationLevel to JSON.
func (c MongoDBCollectionValidationLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type certificate_status struct {
	Name string
}

func (a certificate_status) String() string {
	return a.Name
}

var (
	certificate_statusDeprecated = certificate_status{"DEPRECATED"}
	certificate_statusDraft      = certificate_status{"DRAFT"}
	certificate_statusVerified   = certificate_status{"VERIFIED"}
)

// UnmarshalJSON customizes the unmarshalling of a certificate_status from JSON.
func (c *certificate_status) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DEPRECATED":
		*c = certificate_statusDeprecated

	case "DRAFT":
		*c = certificate_statusDraft

	case "VERIFIED":
		*c = certificate_statusVerified
	default:
		*c = certificate_status{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a certificate_status to JSON.
func (c certificate_status) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type icon_type struct {
	Name string
}

func (a icon_type) String() string {
	return a.Name
}

var (
	icon_typeImage = icon_type{"image"}
	icon_typeEmoji = icon_type{"emoji"}
)

// UnmarshalJSON customizes the unmarshalling of a icon_type from JSON.
func (c *icon_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "image":
		*c = icon_typeImage

	case "emoji":
		*c = icon_typeEmoji
	default:
		*c = icon_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a icon_type to JSON.
func (c icon_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type DataProductSensitivity struct {
	Name string
}

func (a DataProductSensitivity) String() string {
	return a.Name
}

var (
	DataProductSensitivityPublic       = DataProductSensitivity{"Public"}
	DataProductSensitivityInternal     = DataProductSensitivity{"Internal"}
	DataProductSensitivityConfidential = DataProductSensitivity{"Confidential"}
)

// UnmarshalJSON customizes the unmarshalling of a DataProductSensitivity from JSON.
func (c *DataProductSensitivity) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Public":
		*c = DataProductSensitivityPublic

	case "Internal":
		*c = DataProductSensitivityInternal

	case "Confidential":
		*c = DataProductSensitivityConfidential
	default:
		*c = DataProductSensitivity{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a DataProductSensitivity to JSON.
func (c DataProductSensitivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSEncryptionTypes struct {
	Name string
}

func (a ADLSEncryptionTypes) String() string {
	return a.Name
}

var (
	ADLSEncryptionTypesMicrosoftstorage  = ADLSEncryptionTypes{"Microsoft.Storage"}
	ADLSEncryptionTypesMicrosoftkeyvault = ADLSEncryptionTypes{"Microsoft.Keyvault"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSEncryptionTypes from JSON.
func (c *ADLSEncryptionTypes) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Microsoft.Storage":
		*c = ADLSEncryptionTypesMicrosoftstorage

	case "Microsoft.Keyvault":
		*c = ADLSEncryptionTypesMicrosoftkeyvault
	default:
		*c = ADLSEncryptionTypes{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSEncryptionTypes to JSON.
func (c ADLSEncryptionTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type SchemaRegistrySchemaType struct {
	Name string
}

func (a SchemaRegistrySchemaType) String() string {
	return a.Name
}

var (
	SchemaRegistrySchemaTypeAvro     = SchemaRegistrySchemaType{"AVRO"}
	SchemaRegistrySchemaTypeJson     = SchemaRegistrySchemaType{"JSON"}
	SchemaRegistrySchemaTypeProtobuf = SchemaRegistrySchemaType{"PROTOBUF"}
)

// UnmarshalJSON customizes the unmarshalling of a SchemaRegistrySchemaType from JSON.
func (c *SchemaRegistrySchemaType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "AVRO":
		*c = SchemaRegistrySchemaTypeAvro

	case "JSON":
		*c = SchemaRegistrySchemaTypeJson

	case "PROTOBUF":
		*c = SchemaRegistrySchemaTypeProtobuf
	default:
		*c = SchemaRegistrySchemaType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a SchemaRegistrySchemaType to JSON.
func (c SchemaRegistrySchemaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type incident_severity struct {
	Name string
}

func (a incident_severity) String() string {
	return a.Name
}

var (
	incident_severityLow    = incident_severity{"LOW"}
	incident_severityMedium = incident_severity{"MEDIUM"}
	incident_severityHigh   = incident_severity{"HIGH"}
)

// UnmarshalJSON customizes the unmarshalling of a incident_severity from JSON.
func (c *incident_severity) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "LOW":
		*c = incident_severityLow

	case "MEDIUM":
		*c = incident_severityMedium

	case "HIGH":
		*c = incident_severityHigh
	default:
		*c = incident_severity{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a incident_severity to JSON.
func (c incident_severity) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type quick_sight_folder_type struct {
	Name string
}

func (a quick_sight_folder_type) String() string {
	return a.Name
}

var (
	quick_sight_folder_typeShared = quick_sight_folder_type{"SHARED"}
)

// UnmarshalJSON customizes the unmarshalling of a quick_sight_folder_type from JSON.
func (c *quick_sight_folder_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "SHARED":
		*c = quick_sight_folder_typeShared
	default:
		*c = quick_sight_folder_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a quick_sight_folder_type to JSON.
func (c quick_sight_folder_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSStorageKind struct {
	Name string
}

func (a ADLSStorageKind) String() string {
	return a.Name
}

var (
	ADLSStorageKindBlobstorage      = ADLSStorageKind{"BlobStorage"}
	ADLSStorageKindBlockblobstorage = ADLSStorageKind{"BlockBlobStorage"}
	ADLSStorageKindFilestorage      = ADLSStorageKind{"FileStorage"}
	ADLSStorageKindStorage          = ADLSStorageKind{"Storage"}
	ADLSStorageKindStoragev2        = ADLSStorageKind{"StorageV2"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSStorageKind from JSON.
func (c *ADLSStorageKind) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "BlobStorage":
		*c = ADLSStorageKindBlobstorage

	case "BlockBlobStorage":
		*c = ADLSStorageKindBlockblobstorage

	case "FileStorage":
		*c = ADLSStorageKindFilestorage

	case "Storage":
		*c = ADLSStorageKindStorage

	case "StorageV2":
		*c = ADLSStorageKindStoragev2
	default:
		*c = ADLSStorageKind{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSStorageKind to JSON.
func (c ADLSStorageKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type DataProductVisibility struct {
	Name string
}

func (a DataProductVisibility) String() string {
	return a.Name
}

var (
	DataProductVisibilityPrivate   = DataProductVisibility{"Private"}
	DataProductVisibilityProtected = DataProductVisibility{"Protected"}
	DataProductVisibilityPublic    = DataProductVisibility{"Public"}
)

// UnmarshalJSON customizes the unmarshalling of a DataProductVisibility from JSON.
func (c *DataProductVisibility) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Private":
		*c = DataProductVisibilityPrivate

	case "Protected":
		*c = DataProductVisibilityProtected

	case "Public":
		*c = DataProductVisibilityPublic
	default:
		*c = DataProductVisibility{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a DataProductVisibility to JSON.
func (c DataProductVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ModelCardinalityType struct {
	Name string
}

func (a ModelCardinalityType) String() string {
	return a.Name
}

var ( // An entity (E) is connected to at most one other entity (F), and vice versa.
	ModelCardinalityTypeOne_to_one   = ModelCardinalityType{"ONE-TO-ONE"}  // An entity (E) can be associated with multiple entities (F), but each entity (F) is associated with at most one entity (E).
	ModelCardinalityTypeOne_to_many  = ModelCardinalityType{"ONE-TO-MANY"} // Multiple entities (E) can be connected to the same entity (F), but each entity (F) is associated with at most one entity (E).
	ModelCardinalityTypeMany_to_one  = ModelCardinalityType{"MANY-TO-ONE"} // Entities (E) can be associated with multiple other entities (F), and entities (F) can be associated with multiple entities (E).
	ModelCardinalityTypeMany_to_many = ModelCardinalityType{"MANY-TO-MANY"}
)

// UnmarshalJSON customizes the unmarshalling of a ModelCardinalityType from JSON.
func (c *ModelCardinalityType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "ONE-TO-ONE":
		*c = ModelCardinalityTypeOne_to_one

	case "ONE-TO-MANY":
		*c = ModelCardinalityTypeOne_to_many

	case "MANY-TO-ONE":
		*c = ModelCardinalityTypeMany_to_one

	case "MANY-TO-MANY":
		*c = ModelCardinalityTypeMany_to_many
	default:
		*c = ModelCardinalityType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ModelCardinalityType to JSON.
func (c ModelCardinalityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSReplicationType struct {
	Name string
}

func (a ADLSReplicationType) String() string {
	return a.Name
}

var ( // Locally redundant storage
	ADLSReplicationTypeLrs    = ADLSReplicationType{"LRS"}  // Zone-redundant storage
	ADLSReplicationTypeZrs    = ADLSReplicationType{"ZRS"}  // Geo-redundant storage
	ADLSReplicationTypeGrs    = ADLSReplicationType{"GRS"}  // Geo-zone-redundant storage
	ADLSReplicationTypeGzrs   = ADLSReplicationType{"GZRS"} // Read-access geo-redundant storage
	ADLSReplicationTypeRa_grs = ADLSReplicationType{"RA-GRS"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSReplicationType from JSON.
func (c *ADLSReplicationType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "LRS":
		*c = ADLSReplicationTypeLrs

	case "ZRS":
		*c = ADLSReplicationTypeZrs

	case "GRS":
		*c = ADLSReplicationTypeGrs

	case "GZRS":
		*c = ADLSReplicationTypeGzrs

	case "RA-GRS":
		*c = ADLSReplicationTypeRa_grs
	default:
		*c = ADLSReplicationType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSReplicationType to JSON.
func (c ADLSReplicationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type WorkflowType struct {
	Name string
}

func (a WorkflowType) String() string {
	return a.Name
}

var (
	WorkflowTypeData_access            = WorkflowType{"DATA_ACCESS"}
	WorkflowTypePolicy                 = WorkflowType{"POLICY"}
	WorkflowTypeChange_management      = WorkflowType{"CHANGE_MANAGEMENT"}
	WorkflowTypePublication_management = WorkflowType{"PUBLICATION_MANAGEMENT"}
	WorkflowTypeImpact_analysis        = WorkflowType{"IMPACT_ANALYSIS"}
)

// UnmarshalJSON customizes the unmarshalling of a WorkflowType from JSON.
func (c *WorkflowType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DATA_ACCESS":
		*c = WorkflowTypeData_access

	case "POLICY":
		*c = WorkflowTypePolicy

	case "CHANGE_MANAGEMENT":
		*c = WorkflowTypeChange_management

	case "PUBLICATION_MANAGEMENT":
		*c = WorkflowTypePublication_management

	case "IMPACT_ANALYSIS":
		*c = WorkflowTypeImpact_analysis
	default:
		*c = WorkflowType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a WorkflowType to JSON.
func (c WorkflowType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AtlasGlossaryTermAssignmentStatus struct {
	Name string
}

func (a AtlasGlossaryTermAssignmentStatus) String() string {
	return a.Name
}

var ( // DISCOVERED means that the semantic assignment was added by a discovery engine.
	AtlasGlossaryTermAssignmentStatusDiscovered = AtlasGlossaryTermAssignmentStatus{"DISCOVERED"} // PROPOSED means that the semantic assignment was proposed by person - they may be a subject matter expert, or consumer of the Referenceable asset
	AtlasGlossaryTermAssignmentStatusProposed   = AtlasGlossaryTermAssignmentStatus{"PROPOSED"}   // IMPORTED means that the semantic assignment has been imported from outside of the open metadata cluster
	AtlasGlossaryTermAssignmentStatusImported   = AtlasGlossaryTermAssignmentStatus{"IMPORTED"}   // VALIDATED means that the semantic assignment has been reviewed and is highly trusted.
	AtlasGlossaryTermAssignmentStatusValidated  = AtlasGlossaryTermAssignmentStatus{"VALIDATED"}  // DEPRECATED means that the semantic assignment is being phased out. There may be another semantic assignment to the Referenceable that will ultimately replace this one.
	AtlasGlossaryTermAssignmentStatusDeprecated = AtlasGlossaryTermAssignmentStatus{"DEPRECATED"} // OBSOLETE means that the semantic assignment is no longer in use,
	AtlasGlossaryTermAssignmentStatusObsolete   = AtlasGlossaryTermAssignmentStatus{"OBSOLETE"}   // OTHER means that the semantic assignment value does not match any of the other Term Assignment Status values
	AtlasGlossaryTermAssignmentStatusOther      = AtlasGlossaryTermAssignmentStatus{"OTHER"}
)

// UnmarshalJSON customizes the unmarshalling of a AtlasGlossaryTermAssignmentStatus from JSON.
func (c *AtlasGlossaryTermAssignmentStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DISCOVERED":
		*c = AtlasGlossaryTermAssignmentStatusDiscovered

	case "PROPOSED":
		*c = AtlasGlossaryTermAssignmentStatusProposed

	case "IMPORTED":
		*c = AtlasGlossaryTermAssignmentStatusImported

	case "VALIDATED":
		*c = AtlasGlossaryTermAssignmentStatusValidated

	case "DEPRECATED":
		*c = AtlasGlossaryTermAssignmentStatusDeprecated

	case "OBSOLETE":
		*c = AtlasGlossaryTermAssignmentStatusObsolete

	case "OTHER":
		*c = AtlasGlossaryTermAssignmentStatusOther
	default:
		*c = AtlasGlossaryTermAssignmentStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AtlasGlossaryTermAssignmentStatus to JSON.
func (c AtlasGlossaryTermAssignmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type MongoDBCollectionValidationAction struct {
	Name string
}

func (a MongoDBCollectionValidationAction) String() string {
	return a.Name
}

var ( // ERROR means validator will throw an error in case the validation fails
	MongoDBCollectionValidationActionError = MongoDBCollectionValidationAction{"ERROR"} // WARN means validator will throw an warning in case the validation fails
	MongoDBCollectionValidationActionWarn  = MongoDBCollectionValidationAction{"WARN"}
)

// UnmarshalJSON customizes the unmarshalling of a MongoDBCollectionValidationAction from JSON.
func (c *MongoDBCollectionValidationAction) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "ERROR":
		*c = MongoDBCollectionValidationActionError

	case "WARN":
		*c = MongoDBCollectionValidationActionWarn
	default:
		*c = MongoDBCollectionValidationAction{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a MongoDBCollectionValidationAction to JSON.
func (c MongoDBCollectionValidationAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type kafka_topic_compression_type struct {
	Name string
}

func (a kafka_topic_compression_type) String() string {
	return a.Name
}

var (
	kafka_topic_compression_typeUncompressed = kafka_topic_compression_type{"uncompressed"}
	kafka_topic_compression_typeZstd         = kafka_topic_compression_type{"zstd"}
	kafka_topic_compression_typeLz4          = kafka_topic_compression_type{"lz4"}
	kafka_topic_compression_typeSnappy       = kafka_topic_compression_type{"snappy"}
	kafka_topic_compression_typeGzip         = kafka_topic_compression_type{"gzip"}
	kafka_topic_compression_typeProducer     = kafka_topic_compression_type{"producer"}
)

// UnmarshalJSON customizes the unmarshalling of a kafka_topic_compression_type from JSON.
func (c *kafka_topic_compression_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "uncompressed":
		*c = kafka_topic_compression_typeUncompressed

	case "zstd":
		*c = kafka_topic_compression_typeZstd

	case "lz4":
		*c = kafka_topic_compression_typeLz4

	case "snappy":
		*c = kafka_topic_compression_typeSnappy

	case "gzip":
		*c = kafka_topic_compression_typeGzip

	case "producer":
		*c = kafka_topic_compression_typeProducer
	default:
		*c = kafka_topic_compression_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a kafka_topic_compression_type to JSON.
func (c kafka_topic_compression_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AuthPolicyCategory struct {
	Name string
}

func (a AuthPolicyCategory) String() string {
	return a.Name
}

var (
	AuthPolicyCategoryBootstrap = AuthPolicyCategory{"bootstrap"}
	AuthPolicyCategoryPersona   = AuthPolicyCategory{"persona"}
	AuthPolicyCategoryPurpose   = AuthPolicyCategory{"purpose"}
)

// UnmarshalJSON customizes the unmarshalling of a AuthPolicyCategory from JSON.
func (c *AuthPolicyCategory) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "bootstrap":
		*c = AuthPolicyCategoryBootstrap

	case "persona":
		*c = AuthPolicyCategoryPersona

	case "purpose":
		*c = AuthPolicyCategoryPurpose
	default:
		*c = AuthPolicyCategory{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AuthPolicyCategory to JSON.
func (c AuthPolicyCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AuthPolicyResourceCategory struct {
	Name string
}

func (a AuthPolicyResourceCategory) String() string {
	return a.Name
}

var (
	AuthPolicyResourceCategoryEntity       = AuthPolicyResourceCategory{"ENTITY"}
	AuthPolicyResourceCategoryRelationship = AuthPolicyResourceCategory{"RELATIONSHIP"}
	AuthPolicyResourceCategoryTag          = AuthPolicyResourceCategory{"TAG"}
	AuthPolicyResourceCategoryCustom       = AuthPolicyResourceCategory{"CUSTOM"}
	AuthPolicyResourceCategoryTypedefs     = AuthPolicyResourceCategory{"TYPEDEFS"}
	AuthPolicyResourceCategoryAdmin        = AuthPolicyResourceCategory{"ADMIN"}
)

// UnmarshalJSON customizes the unmarshalling of a AuthPolicyResourceCategory from JSON.
func (c *AuthPolicyResourceCategory) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "ENTITY":
		*c = AuthPolicyResourceCategoryEntity

	case "RELATIONSHIP":
		*c = AuthPolicyResourceCategoryRelationship

	case "TAG":
		*c = AuthPolicyResourceCategoryTag

	case "CUSTOM":
		*c = AuthPolicyResourceCategoryCustom

	case "TYPEDEFS":
		*c = AuthPolicyResourceCategoryTypedefs

	case "ADMIN":
		*c = AuthPolicyResourceCategoryAdmin
	default:
		*c = AuthPolicyResourceCategory{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AuthPolicyResourceCategory to JSON.
func (c AuthPolicyResourceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSAccessTier struct {
	Name string
}

func (a ADLSAccessTier) String() string {
	return a.Name
}

var (
	ADLSAccessTierCool    = ADLSAccessTier{"Cool"}
	ADLSAccessTierHot     = ADLSAccessTier{"Hot"}
	ADLSAccessTierArchive = ADLSAccessTier{"Archive"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSAccessTier from JSON.
func (c *ADLSAccessTier) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Cool":
		*c = ADLSAccessTierCool

	case "Hot":
		*c = ADLSAccessTierHot

	case "Archive":
		*c = ADLSAccessTierArchive
	default:
		*c = ADLSAccessTier{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSAccessTier to JSON.
func (c ADLSAccessTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type WorkflowRunStatus struct {
	Name string
}

func (a WorkflowRunStatus) String() string {
	return a.Name
}

var (
	WorkflowRunStatusPending    = WorkflowRunStatus{"PENDING"}
	WorkflowRunStatusApproved   = WorkflowRunStatus{"APPROVED"}
	WorkflowRunStatusRejected   = WorkflowRunStatus{"REJECTED"}
	WorkflowRunStatusSuccess    = WorkflowRunStatus{"SUCCESS"}
	WorkflowRunStatusFailure    = WorkflowRunStatus{"FAILURE"}
	WorkflowRunStatusWithdrawn  = WorkflowRunStatus{"WITHDRAWN"}
	WorkflowRunStatusExpired    = WorkflowRunStatus{"EXPIRED"}
	WorkflowRunStatusTerminated = WorkflowRunStatus{"TERMINATED"}
)

// UnmarshalJSON customizes the unmarshalling of a WorkflowRunStatus from JSON.
func (c *WorkflowRunStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "PENDING":
		*c = WorkflowRunStatusPending

	case "APPROVED":
		*c = WorkflowRunStatusApproved

	case "REJECTED":
		*c = WorkflowRunStatusRejected

	case "SUCCESS":
		*c = WorkflowRunStatusSuccess

	case "FAILURE":
		*c = WorkflowRunStatusFailure

	case "WITHDRAWN":
		*c = WorkflowRunStatusWithdrawn

	case "EXPIRED":
		*c = WorkflowRunStatusExpired

	case "TERMINATED":
		*c = WorkflowRunStatusTerminated
	default:
		*c = WorkflowRunStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a WorkflowRunStatus to JSON.
func (c WorkflowRunStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AtlasGlossaryCategoryType struct {
	Name string
}

func (a AtlasGlossaryCategoryType) String() string {
	return a.Name
}

var ( // Document Folder will contain Documents
	AtlasGlossaryCategoryTypeDocument_folder = AtlasGlossaryCategoryType{"DOCUMENT_FOLDER"}
)

// UnmarshalJSON customizes the unmarshalling of a AtlasGlossaryCategoryType from JSON.
func (c *AtlasGlossaryCategoryType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DOCUMENT_FOLDER":
		*c = AtlasGlossaryCategoryTypeDocument_folder
	default:
		*c = AtlasGlossaryCategoryType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AtlasGlossaryCategoryType to JSON.
func (c AtlasGlossaryCategoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AuthPolicyType struct {
	Name string
}

func (a AuthPolicyType) String() string {
	return a.Name
}

var (
	AuthPolicyTypeAllow           = AuthPolicyType{"allow"}
	AuthPolicyTypeDeny            = AuthPolicyType{"deny"}
	AuthPolicyTypeAllowexceptions = AuthPolicyType{"allowExceptions"}
	AuthPolicyTypeDenyexceptions  = AuthPolicyType{"denyExceptions"}
	AuthPolicyTypeDatamask        = AuthPolicyType{"dataMask"}
	AuthPolicyTypeRowfilter       = AuthPolicyType{"rowFilter"}
)

// UnmarshalJSON customizes the unmarshalling of a AuthPolicyType from JSON.
func (c *AuthPolicyType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "allow":
		*c = AuthPolicyTypeAllow

	case "deny":
		*c = AuthPolicyTypeDeny

	case "allowExceptions":
		*c = AuthPolicyTypeAllowexceptions

	case "denyExceptions":
		*c = AuthPolicyTypeDenyexceptions

	case "dataMask":
		*c = AuthPolicyTypeDatamask

	case "rowFilter":
		*c = AuthPolicyTypeRowfilter
	default:
		*c = AuthPolicyType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AuthPolicyType to JSON.
func (c AuthPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type file_type struct {
	Name string
}

func (a file_type) String() string {
	return a.Name
}

var (
	file_typePdf  = file_type{"pdf"}
	file_typeDoc  = file_type{"doc"}
	file_typeXls  = file_type{"xls"}
	file_typePpt  = file_type{"ppt"}
	file_typeCsv  = file_type{"csv"}
	file_typeTxt  = file_type{"txt"}
	file_typeJson = file_type{"json"}
	file_typeXml  = file_type{"xml"}
	file_typeZip  = file_type{"zip"}
)

// UnmarshalJSON customizes the unmarshalling of a file_type from JSON.
func (c *file_type) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "pdf":
		*c = file_typePdf

	case "doc":
		*c = file_typeDoc

	case "xls":
		*c = file_typeXls

	case "ppt":
		*c = file_typePpt

	case "csv":
		*c = file_typeCsv

	case "txt":
		*c = file_typeTxt

	case "json":
		*c = file_typeJson

	case "xml":
		*c = file_typeXml

	case "zip":
		*c = file_typeZip
	default:
		*c = file_type{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a file_type to JSON.
func (c file_type) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AtlasGlossaryTermType struct {
	Name string
}

func (a AtlasGlossaryTermType) String() string {
	return a.Name
}

var ( // Will represent Document as knowledge hub
	AtlasGlossaryTermTypeDocument = AtlasGlossaryTermType{"DOCUMENT"}
)

// UnmarshalJSON customizes the unmarshalling of a AtlasGlossaryTermType from JSON.
func (c *AtlasGlossaryTermType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DOCUMENT":
		*c = AtlasGlossaryTermTypeDocument
	default:
		*c = AtlasGlossaryTermType{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AtlasGlossaryTermType to JSON.
func (c AtlasGlossaryTermType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type DynamoDBStatus struct {
	Name string
}

func (a DynamoDBStatus) String() string {
	return a.Name
}

var ( // The table/index is being created
	DynamoDBStatusCreating                            = DynamoDBStatus{"CREATING"}                            // The table/index is being updated
	DynamoDBStatusUpdating                            = DynamoDBStatus{"UPDATING"}                            // The table/index is being deleted
	DynamoDBStatusDeleting                            = DynamoDBStatus{"DELETING"}                            // The table/index is ready for use
	DynamoDBStatusActive                              = DynamoDBStatus{"ACTIVE"}                              // The AWS KMS key used to encrypt the table in inaccessible. Table operations may fail due to failure to use the AWS KMS key. DynamoDB will initiate the table archival process when a table's AWS KMS key remains inaccessible for more than seven days.
	DynamoDBStatusInaccessible_encryption_credentials = DynamoDBStatus{"INACCESSIBLE_ENCRYPTION_CREDENTIALS"} // The table is being archived. Operations are not allowed until archival is complete.
	DynamoDBStatusArchiving                           = DynamoDBStatus{"ARCHIVING"}                           // The table has been archived.
	DynamoDBStatusArchived                            = DynamoDBStatus{"ARCHIVED"}
)

// UnmarshalJSON customizes the unmarshalling of a DynamoDBStatus from JSON.
func (c *DynamoDBStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "CREATING":
		*c = DynamoDBStatusCreating

	case "UPDATING":
		*c = DynamoDBStatusUpdating

	case "DELETING":
		*c = DynamoDBStatusDeleting

	case "ACTIVE":
		*c = DynamoDBStatusActive

	case "INACCESSIBLE_ENCRYPTION_CREDENTIALS":
		*c = DynamoDBStatusInaccessible_encryption_credentials

	case "ARCHIVING":
		*c = DynamoDBStatusArchiving

	case "ARCHIVED":
		*c = DynamoDBStatusArchived
	default:
		*c = DynamoDBStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a DynamoDBStatus to JSON.
func (c DynamoDBStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type AtlasGlossaryTermRelationshipStatus struct {
	Name string
}

func (a AtlasGlossaryTermRelationshipStatus) String() string {
	return a.Name
}

var ( // DRAFT means the relationship is under development.
	AtlasGlossaryTermRelationshipStatusDraft      = AtlasGlossaryTermRelationshipStatus{"DRAFT"}      // ACTIVE means the relationship is validated and in use.
	AtlasGlossaryTermRelationshipStatusActive     = AtlasGlossaryTermRelationshipStatus{"ACTIVE"}     // DEPRECATED means the the relationship is being phased out.
	AtlasGlossaryTermRelationshipStatusDeprecated = AtlasGlossaryTermRelationshipStatus{"DEPRECATED"} // OBSOLETE means that the relationship should not be used anymore.
	AtlasGlossaryTermRelationshipStatusObsolete   = AtlasGlossaryTermRelationshipStatus{"OBSOLETE"}   // OTHER means that there is another status.
	AtlasGlossaryTermRelationshipStatusOther      = AtlasGlossaryTermRelationshipStatus{"OTHER"}
)

// UnmarshalJSON customizes the unmarshalling of a AtlasGlossaryTermRelationshipStatus from JSON.
func (c *AtlasGlossaryTermRelationshipStatus) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "DRAFT":
		*c = AtlasGlossaryTermRelationshipStatusDraft

	case "ACTIVE":
		*c = AtlasGlossaryTermRelationshipStatusActive

	case "DEPRECATED":
		*c = AtlasGlossaryTermRelationshipStatusDeprecated

	case "OBSOLETE":
		*c = AtlasGlossaryTermRelationshipStatusObsolete

	case "OTHER":
		*c = AtlasGlossaryTermRelationshipStatusOther
	default:
		*c = AtlasGlossaryTermRelationshipStatus{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a AtlasGlossaryTermRelationshipStatus to JSON.
func (c AtlasGlossaryTermRelationshipStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type OpenLineageRunState struct {
	Name string
}

func (a OpenLineageRunState) String() string {
	return a.Name
}

var (
	OpenLineageRunStateStart    = OpenLineageRunState{"START"}
	OpenLineageRunStateRunning  = OpenLineageRunState{"RUNNING"}
	OpenLineageRunStateComplete = OpenLineageRunState{"COMPLETE"}
	OpenLineageRunStateAbort    = OpenLineageRunState{"ABORT"}
	OpenLineageRunStateFail     = OpenLineageRunState{"FAIL"}
	OpenLineageRunStateOther    = OpenLineageRunState{"OTHER"}
)

// UnmarshalJSON customizes the unmarshalling of a OpenLineageRunState from JSON.
func (c *OpenLineageRunState) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "START":
		*c = OpenLineageRunStateStart

	case "RUNNING":
		*c = OpenLineageRunStateRunning

	case "COMPLETE":
		*c = OpenLineageRunStateComplete

	case "ABORT":
		*c = OpenLineageRunStateAbort

	case "FAIL":
		*c = OpenLineageRunStateFail

	case "OTHER":
		*c = OpenLineageRunStateOther
	default:
		*c = OpenLineageRunState{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a OpenLineageRunState to JSON.
func (c OpenLineageRunState) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type ADLSPerformance struct {
	Name string
}

func (a ADLSPerformance) String() string {
	return a.Name
}

var (
	ADLSPerformanceStandard = ADLSPerformance{"Standard"}
	ADLSPerformancePremium  = ADLSPerformance{"Premium"}
)

// UnmarshalJSON customizes the unmarshalling of a ADLSPerformance from JSON.
func (c *ADLSPerformance) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "Standard":
		*c = ADLSPerformanceStandard

	case "Premium":
		*c = ADLSPerformancePremium
	default:
		*c = ADLSPerformance{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a ADLSPerformance to JSON.
func (c ADLSPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}

type quick_sight_analysis_status struct {
	Name string
}

func (a quick_sight_analysis_status) String() string {
	return a.Name
}

var (
	quick_sight_analysis_statusCreation_in_progress = quick_sight_analysis_status{"CREATION_IN_PROGRESS"}
	quick_sight_analysis_statusCreation_successful  = quick_sight_analysis_status{"CREATION_SUCCESSFUL"}
	quick_sight_analysis_statusCreation_failed      = quick_sight_analysis_status{"CREATION_FAILED"}
	quick_sight_analysis_statusUpdate_in_progress   = quick_sight_analysis_status{"UPDATE_IN_PROGRESS"}
	quick_sight_analysis_statusUpdate_successful    = quick_sight_analysis_status{"UPDATE_SUCCESSFUL"}
	quick_sight_analysis_statusUpdate_failed        = quick_sight_analysis_status{"UPDATE_FAILED"}
	quick_sight_analysis_statusDeleted              = quick_sight_analysis_status{"DELETED"}
)

// UnmarshalJSON customizes the unmarshalling of a quick_sight_analysis_status from JSON.
func (c *quick_sight_analysis_status) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}

	switch name {

	case "CREATION_IN_PROGRESS":
		*c = quick_sight_analysis_statusCreation_in_progress

	case "CREATION_SUCCESSFUL":
		*c = quick_sight_analysis_statusCreation_successful

	case "CREATION_FAILED":
		*c = quick_sight_analysis_statusCreation_failed

	case "UPDATE_IN_PROGRESS":
		*c = quick_sight_analysis_statusUpdate_in_progress

	case "UPDATE_SUCCESSFUL":
		*c = quick_sight_analysis_statusUpdate_successful

	case "UPDATE_FAILED":
		*c = quick_sight_analysis_statusUpdate_failed

	case "DELETED":
		*c = quick_sight_analysis_statusDeleted
	default:
		*c = quick_sight_analysis_status{Name: name}
	}

	return nil
}

// MarshalJSON customizes the marshalling of a quick_sight_analysis_status to JSON.
func (c quick_sight_analysis_status) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Name)
}
