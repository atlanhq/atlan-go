// # **************************************
// # CODE BELOW IS GENERATED NOT MODIFY  **
// # **************************************

package generator

import "time"

// BadgeCondition: Detailed information about a condition used in coloring a custom metadata badge in Atlan.

// BadgeCondition represents the struct for BadgeCondition.
type BadgeCondition struct {
	Badgeconditionoperator *string `json:"badgeconditionoperator,omitempty"`
	Badgeconditionvalue    *string `json:"badgeconditionvalue,omitempty"`
	Badgeconditioncolorhex *string `json:"badgeconditioncolorhex,omitempty"`
}

// StarredDetails: Detailed information about the users who have starred an asset.

// StarredDetails represents the struct for StarredDetails.
type StarredDetails struct {
	Assetstarredby *string    `json:"assetstarredby,omitempty"`
	Assetstarredat *time.Time `json:"assetstarredat,omitempty"`
}

// AzureTag: Detailed information about an Azure tag.

// AzureTag represents the struct for AzureTag.
type AzureTag struct {
	Azuretagkey   *string `json:"azuretagkey,omitempty"`
	Azuretagvalue *string `json:"azuretagvalue,omitempty"`
}

// GoogleLabel: Detailed information about a Google label.

// GoogleLabel represents the struct for GoogleLabel.
type GoogleLabel struct {
	Googlelabelkey   *string `json:"googlelabelkey,omitempty"`
	Googlelabelvalue *string `json:"googlelabelvalue,omitempty"`
}

// SourceTagAttribute: Detailed information about a source tag's attributes.

// SourceTagAttribute represents the struct for SourceTagAttribute.
type SourceTagAttribute struct {
	Tagattributekey        *string            `json:"tagattributekey,omitempty"`
	Tagattributevalue      *string            `json:"tagattributevalue,omitempty"`
	Tagattributeproperties map[string]*string `json:"tagattributeproperties,omitempty"`
}

// MCRuleSchedule: Detailed information about the schedule for a Monte Carlo rule.

// MCRuleSchedule represents the struct for MCRuleSchedule.
type MCRuleSchedule struct {
	Mcrulescheduletype              *string    `json:"mcrulescheduletype,omitempty"`
	Mcrulescheduleintervalinminutes *int       `json:"mcrulescheduleintervalinminutes,omitempty"`
	Mcruleschedulestarttime         *time.Time `json:"mcruleschedulestarttime,omitempty"`
	Mcruleschedulecrontab           *string    `json:"mcruleschedulecrontab,omitempty"`
}

// KafkaTopicConsumption: Detailed information about the consumption of a Kafka topic.

// KafkaTopicConsumption represents the struct for KafkaTopicConsumption.
type KafkaTopicConsumption struct {
	Topicname          *string `json:"topicname,omitempty"`
	Topicpartition     *string `json:"topicpartition,omitempty"`
	Topiclag           *int64  `json:"topiclag,omitempty"`
	Topiccurrentoffset *int64  `json:"topiccurrentoffset,omitempty"`
}

// SourceTagAttachment: Detailed information about the attachment of a tag to an Atlan asset, synced from source.

// SourceTagAttachment represents the struct for SourceTagAttachment.
type SourceTagAttachment struct {
	Sourcetagname          *string                     `json:"sourcetagname,omitempty"`
	Sourcetagqualifiedname *string                     `json:"sourcetagqualifiedname,omitempty"`
	Sourcetagguid          *string                     `json:"sourcetagguid,omitempty"`
	Sourcetagconnectorname *string                     `json:"sourcetagconnectorname,omitempty"`
	Sourcetagvalue         []*SourceTagAttachmentValue `json:"sourcetagvalue,omitempty"`
	Issourcetagsynced      *bool                       `json:"issourcetagsynced,omitempty"`
	Sourcetagsynctimestamp *time.Time                  `json:"sourcetagsynctimestamp,omitempty"`
	Sourcetagsyncerror     *string                     `json:"sourcetagsyncerror,omitempty"`
}

// AuthPolicyCondition: Policy condition schedule struct

// AuthPolicyCondition represents the struct for AuthPolicyCondition.
type AuthPolicyCondition struct {
	Policyconditiontype   *string   `json:"policyconditiontype,omitempty"`
	Policyconditionvalues []*string `json:"policyconditionvalues,omitempty"`
}

// AwsTag: Detailed information about an AWS tag.

// AwsTag represents the struct for AwsTag.
type AwsTag struct {
	Awstagkey   *string `json:"awstagkey,omitempty"`
	Awstagvalue *string `json:"awstagvalue,omitempty"`
}

// DbtJobRun: Detailed information about a dbt job run.

// DbtJobRun represents the struct for DbtJobRun.
type DbtJobRun struct {
	Dbtjobid             *string    `json:"dbtjobid,omitempty"`
	Dbtjobname           *string    `json:"dbtjobname,omitempty"`
	Dbtenvironmentid     *string    `json:"dbtenvironmentid,omitempty"`
	Dbtenvironmentname   *string    `json:"dbtenvironmentname,omitempty"`
	Dbtjobrunid          *string    `json:"dbtjobrunid,omitempty"`
	Dbtjobruncompletedat *time.Time `json:"dbtjobruncompletedat,omitempty"`
	Dbtjobrunstatus      *string    `json:"dbtjobrunstatus,omitempty"`
	Dbttestrunstatus     *string    `json:"dbttestrunstatus,omitempty"`
	Dbtmodelrunstatus    *string    `json:"dbtmodelrunstatus,omitempty"`
	Dbtcompiledsql       *string    `json:"dbtcompiledsql,omitempty"`
	Dbtcompiledcode      *string    `json:"dbtcompiledcode,omitempty"`
}

// AwsCloudWatchMetric: Detailed information about an AWS CloudWatch metric.

// AwsCloudWatchMetric represents the struct for AwsCloudWatchMetric.
type AwsCloudWatchMetric struct {
	Awscloudwatchmetricname  *string `json:"awscloudwatchmetricname,omitempty"`
	Awscloudwatchmetricscope *string `json:"awscloudwatchmetricscope,omitempty"`
}

// Action: Action for the task

// Action represents the struct for Action.
type Action struct {
	Taskactionfulfillmenturl     *string `json:"taskactionfulfillmenturl,omitempty"`
	Taskactionfulfillmentmethod  *string `json:"taskactionfulfillmentmethod,omitempty"`
	Taskactionfulfillmentpayload *string `json:"taskactionfulfillmentpayload,omitempty"`
	Taskactiondisplaytext        *string `json:"taskactiondisplaytext,omitempty"`
}

// ColumnValueFrequencyMap: Detailed information representing a column value and it's frequency.

// ColumnValueFrequencyMap represents the struct for ColumnValueFrequencyMap.
type ColumnValueFrequencyMap struct {
	Columnvalue          *string `json:"columnvalue,omitempty"`
	Columnvaluefrequency *int64  `json:"columnvaluefrequency,omitempty"`
}

// DbtMetricFilter: Detailed information about a dbt Metric Filter.

// DbtMetricFilter represents the struct for DbtMetricFilter.
type DbtMetricFilter struct {
	Dbtmetricfiltercolumnqualifiedname *string `json:"dbtmetricfiltercolumnqualifiedname,omitempty"`
	Dbtmetricfilterfield               *string `json:"dbtmetricfilterfield,omitempty"`
	Dbtmetricfilteroperator            *string `json:"dbtmetricfilteroperator,omitempty"`
	Dbtmetricfiltervalue               *string `json:"dbtmetricfiltervalue,omitempty"`
}

// MCRuleComparison: Detailed information about the comparison logic of a Monte Carlo rule.

// MCRuleComparison represents the struct for MCRuleComparison.
type MCRuleComparison struct {
	Mcrulecomparisontype                *string  `json:"mcrulecomparisontype,omitempty"`
	Mcrulecomparisonfield               *string  `json:"mcrulecomparisonfield,omitempty"`
	Mcrulecomparisonmetric              *string  `json:"mcrulecomparisonmetric,omitempty"`
	Mcrulecomparisonoperator            *string  `json:"mcrulecomparisonoperator,omitempty"`
	Mcrulecomparisonthreshold           *float64 `json:"mcrulecomparisonthreshold,omitempty"`
	Mcrulecomparisonisthresholdrelative *bool    `json:"mcrulecomparisonisthresholdrelative,omitempty"`
}

// Histogram: Detailed information representing a histogram of values.

// Histogram represents the struct for Histogram.
type Histogram struct {
	Boundaries  []*float64 `json:"boundaries,omitempty"`
	Frequencies []*float64 `json:"frequencies,omitempty"`
}

// SourceTagAttachmentValue: Detailed information about the value of a source tag's attachment to an asset.

// SourceTagAttachmentValue represents the struct for SourceTagAttachmentValue.
type SourceTagAttachmentValue struct {
	Tagattachmentkey   *string `json:"tagattachmentkey,omitempty"`
	Tagattachmentvalue *string `json:"tagattachmentvalue,omitempty"`
}

// GoogleTag: Detailed information about a Google tag.

// GoogleTag represents the struct for GoogleTag.
type GoogleTag struct {
	Googletagkey   *string `json:"googletagkey,omitempty"`
	Googletagvalue *string `json:"googletagvalue,omitempty"`
}

// AuthPolicyValiditySchedule: Validity schedule struct for policy

// AuthPolicyValiditySchedule represents the struct for AuthPolicyValiditySchedule.
type AuthPolicyValiditySchedule struct {
	Policyvalidityschedulestarttime *string `json:"policyvalidityschedulestarttime,omitempty"`
	Policyvalidityscheduleendtime   *string `json:"policyvalidityscheduleendtime,omitempty"`
	Policyvalidityscheduletimezone  *string `json:"policyvalidityscheduletimezone,omitempty"`
}

// PopularityInsights: Detailed information about an asset's usage or popularity based on aggregated queries.

// PopularityInsights represents the struct for PopularityInsights.
type PopularityInsights struct {
	Recorduser            *string             `json:"recorduser,omitempty"`
	Recordquery           *string             `json:"recordquery,omitempty"`
	Recordqueryduration   *int64              `json:"recordqueryduration,omitempty"`
	Recordquerycount      *int64              `json:"recordquerycount,omitempty"`
	Recordtotalusercount  *int64              `json:"recordtotalusercount,omitempty"`
	Recordcomputecost     *float64            `json:"recordcomputecost,omitempty"`
	Recordmaxcomputecost  *float64            `json:"recordmaxcomputecost,omitempty"`
	Recordcomputecostunit *SourceCostUnitType `json:"recordcomputecostunit,omitempty"`
	Recordlasttimestamp   *time.Time          `json:"recordlasttimestamp,omitempty"`
	Recordwarehouse       *string             `json:"recordwarehouse,omitempty"`
}
