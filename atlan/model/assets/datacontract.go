package assets

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
