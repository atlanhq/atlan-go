package assets

type DataContract struct {
	Catalog
	DataContractJson *string `json:"dataContractJson,omitempty"`
	Version          *string `json:"dataContractVersion,omitempty"`
	AssetGuid        *string `json:"dataContractAssetGuid,omitempty"`
	// Relationships
	LatestCertified         *[]Asset        `json:"dataContractLatestCertified,omitempty"`
	ContractAssetCertified  *[]DataContract `json:"dataContractAssetCertified,omitempty"`
	ContractLatest          *[]Asset        `json:"dataContractLatest,omitempty"`
	ContractAssetLatest     *[]DataContract `json:"dataContractAssetLatest,omitempty"`
	ContractPreviousVersion *[]DataContract `json:"dataContractPreviousVersion,omitempty"`
	ContractNextVersion     *[]DataContract `json:"dataContractNextVersion,omitempty"`
}

/*
type DataContract struct {
	Asset
	Name              *string            `json:"name,omitempty"`
	QualifiedName     *string            `json:"qualifiedName,omitempty"`
	Description       *string            `json:"description,omitempty"`
	CertificateStatus *CertificateStatus `json:"certificateStatus,omitempty"`
	Contract          *string            `json:"contract,omitempty"`
	RelAsset          *RelAttribute      `json:"asset,omitempty"`
}

type UniqueAttributes struct {
	QualifiedName *string `json:"qualifiedName,omitempty"`
}

type RelAttribute struct {
	TypeName         *string           `json:"typeName,omitempty"`
	UniqueAttributes *UniqueAttributes `json:"uniqueAttributes,omitempty"`
}
*/
