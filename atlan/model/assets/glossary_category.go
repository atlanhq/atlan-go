package assets

type AtlasGlossaryCategory struct {
	Asset
	AdditionAttributes map[string]string       `json:"additionalAttributes"`
	Anchor             []AtlasGlossary         `json:"anchor"`
	ChildrenCategories []AtlasGlossaryCategory `json:"childrenCategories"`
	LongDescription    string                  `json:"longDescription"`
	ParentCategory     []AtlasGlossaryCategory `json:"parentCategory"`
	ShortDescription   string                  `json:"shortDescription"`
	Terms              []AtlasGlossaryTerm     `json:"terms"`
}
