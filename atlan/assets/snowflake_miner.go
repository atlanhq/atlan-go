package assets

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/atlanhq/atlan-go/atlan"

	"github.com/atlanhq/atlan-go/atlan/model/structs"
)

// SnowflakeMiner represents a Snowflake miner package
type SnowflakeMiner struct {
	*AbstractMiner
	AdvancedConfig bool
	Name           string
	PackageName    string
	PackagePrefix  string
	ConnectorType  string
	PackageIcon    string
	PackageLogo    string
}

// NewSnowflakeMiner initializes a Snowflake miner
func NewSnowflakeMiner(connectionQualifiedName string) *SnowflakeMiner {
	return &SnowflakeMiner{
		AbstractMiner: NewAbstractMiner(
			connectionQualifiedName,
			"@atlan/snowflake-miner",
			"atlan-snowflake-miner",
		),
		AdvancedConfig: false,
		Name:           "snowflake",
		PackageName:    "@atlan/snowflake-miner",
		PackagePrefix:  atlan.WorkflowPackageSnowflakeMiner.Name,
		ConnectorType:  "snowflake",
		PackageIcon:    "https://docs.snowflake.com/en/_images/logo-snowflake-sans-text.png",
		PackageLogo:    "https://1amiydhcmj36tz3733v94f15-wpengine.netdna-ssl.com/wp-content/themes/snowflake/assets/img/logo-blue.svg",
	}
}

// Direct sets up the miner to extract directly from Snowflake
func (s *SnowflakeMiner) Direct(startEpoch int64, database, schema string) *SnowflakeMiner {
	if database == "" && schema == "" {
		s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "snowflake-database", Value: "default"})
	} else {
		s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "database-name", Value: database})
		s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "schema-name", Value: schema})
	}

	s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "extraction-method", Value: "query_history"})
	s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "miner-start-time-epoch", Value: strconv.FormatInt(startEpoch, 10)})

	return s
}

// ExcludeUsers excludes certain users from usage metrics
func (s *SnowflakeMiner) ExcludeUsers(users []string) *SnowflakeMiner {
	userJSON, _ := json.Marshal(users)
	s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "popularity-exclude-user-config", Value: string(userJSON)})
	return s
}

// PopularityWindow sets the number of days for popularity metrics
func (s *SnowflakeMiner) PopularityWindow(days int) *SnowflakeMiner {
	s.AdvancedConfig = true
	s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "popularity-window-days", Value: strconv.Itoa(days)})
	return s
}

// NativeLineage enables or disables native lineage from Snowflake
func (s *SnowflakeMiner) NativeLineage(enabled bool) *SnowflakeMiner {
	s.AdvancedConfig = true
	s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "native-lineage-active", Value: fmt.Sprintf("%t", enabled)})
	return s
}

// CustomConfig sets a custom configuration for the miner
func (s *SnowflakeMiner) CustomConfig(config map[string]interface{}) *SnowflakeMiner {
	if len(config) > 0 {
		configJSON, _ := json.Marshal(config)
		s.Parameters = append(s.Parameters, structs.NameValuePair{Name: "control-config", Value: string(configJSON)})
	}
	s.AdvancedConfig = true
	return s
}

// GetMetadata generates workflow metadata for Snowflake Miner
func (s *SnowflakeMiner) GetMetadata() *structs.WorkflowMetadata {
	return &structs.WorkflowMetadata{
		Name:      structs.StringPtr(fmt.Sprintf("%s-%d", s.PackagePrefix, s.Epoch)),
		Namespace: structs.StringPtr("default"),
		Labels: map[string]string{
			"orchestration.atlan.com/certified":      "true",
			"orchestration.atlan.com/source":         s.Name,
			"orchestration.atlan.com/sourceCategory": "warehouse",
			"orchestration.atlan.com/type":           "miner",
			"orchestration.atlan.com/verified":       "true",
			"package.argoproj.io/installer":          "argopm",
			"package.argoproj.io/name":               fmt.Sprintf("%s-miner", s.Name),
			"package.argoproj.io/registry":           "httpsc-o-l-o-ns-l-a-s-hs-l-a-s-hpackages.atlan.com",
			"orchestration.atlan.com/atlan-ui":       "true",
		},
		Annotations: map[string]string{
			"orchestration.atlan.com/allowSchedule":   "true",
			"orchestration.atlan.com/categories":      "warehouse,miner",
			"orchestration.atlan.com/docsUrl":         "https://ask.atlan.com/hc/en-us/articles/6482067592337",
			"orchestration.atlan.com/emoji":           "\\uD83D\\uDE80",
			"orchestration.atlan.com/icon":            s.PackageIcon,
			"orchestration.atlan.com/logo":            s.PackageLogo,
			"orchestration.atlan.com/marketplaceLink": fmt.Sprintf("https://packages.atlan.com/-/web/detail/%s", s.PackageName),
			"orchestration.atlan.com/name":            "Snowflake Miner",
			"package.argoproj.io/author":              "Atlan",
			"package.argoproj.io/description":         "Package to mine query history data from Snowflake and store it for further processing. The data mined will be used for generating lineage and usage metrics.", //nolint
			"package.argoproj.io/homepage":            fmt.Sprintf("https://packages.atlan.com/-/web/detail/%s", s.PackageName),
			"package.argoproj.io/keywords":            `["snowflake","warehouse","connector","miner"]`,
			"package.argoproj.io/name":                s.PackageName,
			"package.argoproj.io/registry":            "https://packages.atlan.com",
			"package.argoproj.io/repository":          "https://github.com/atlanhq/marketplace-packages.git",
			"package.argoproj.io/support":             "support@atlan.com",
			"orchestration.atlan.com/atlanName":       fmt.Sprintf("%s-%d", s.PackagePrefix, s.Epoch),
		},
	}
}

// ToWorkflow generates a workflow from the miner configuration
func (s *SnowflakeMiner) ToWorkflow() *structs.Workflow {
	workflow := s.AbstractPackage.ToWorkflow()
	workflow.Metadata = s.GetMetadata() // Override metadata
	return workflow
}
