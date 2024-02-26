package assets

type AtlasGlossaryTerm struct {
	Asset
	ShortDescription     string                  `json:"shortDescription"`
	LongDescription      string                  `json:"longDescription"`
	Example              string                  `json:"example"`
	Abbreviation         string                  `json:"abbreviation"`
	Usage                string                  `json:"usage"`
	AdditionalAttributes map[string]string       `json:"additionalAttributes"`
	Anchor               []AtlasGlossary         `json:"anchor"`
	Antonyms             []AtlasGlossaryTerm     `json:"antonyms"`
	AssignedEntities     []Asset                 `json:"assignedEntities"`
	Categories           []AtlasGlossaryCategory `json:"categories"`
	ValidValuesFor       []AtlasGlossaryTerm     `json:"validValuesFor"`
	ValidValues          []AtlasGlossaryTerm     `json:"validValues"`
	SeeAlso              []AtlasGlossaryTerm     `json:"seeAlso"`
	IsA                  []AtlasGlossaryTerm     `json:"isA"`
	CLASSIFIES           []AtlasGlossaryTerm     `json:"classifies"`
	PreferredToTerms     []AtlasGlossaryTerm     `json:"preferredToTerms"`
	PreferredTerms       []AtlasGlossaryTerm     `json:"preferredTerms"`
	TranslationTerms     []AtlasGlossaryTerm     `json:"translationTerms"`
	Synonyms             []AtlasGlossaryTerm     `json:"synonyms"`
	ReplacedBy           []AtlasGlossaryTerm     `json:"replacedBy"`
	ReplacementTerms     []AtlasGlossaryTerm     `json:"replacementTerms"`
	TranslatedTerms      []AtlasGlossaryTerm     `json:"translatedTerms"`
}
