package structs

type DataProduct struct {
	Asset
	DataProductDataDomain *DataDomain `json:"dataDomain,omitempty"`
	DataProductAssetsDSL  *string     `json:"dataProductAssetsDSL,omitempty"`
}
