package structs

type DataDomain struct {
	Asset
	ParentDomain     *DataDomain `json:"parentDomain,omitempty"`
	UniqueAttributes *struct {
		QualifiedName *string `json:"qualifiedName,omitempty"`
	} `json:"uniqueAttributes,omitempty"`
}
