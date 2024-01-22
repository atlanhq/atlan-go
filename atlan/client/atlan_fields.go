package client

type AtlanField interface {
	GetAtlanFieldName() string
}

type RelationField struct {
	AtlanFieldName string
}

func NewRelationField(atlanFieldName string) *RelationField {
	return &RelationField{AtlanFieldName: atlanFieldName}
}

func (rf *RelationField) GetAtlanFieldName() string {
	return rf.AtlanFieldName
}

type SearchableField struct {
	AtlanFieldName    string
	ElasticFieldName  string
	InternalFieldName string
}

func NewSearchableField(atlanFieldName, elasticFieldName string) *SearchableField {
	return &SearchableField{
		AtlanFieldName:    atlanFieldName,
		ElasticFieldName:  elasticFieldName,
		InternalFieldName: atlanFieldName,
	}
}

func (sf *SearchableField) GetAtlanFieldName() string {
	return sf.AtlanFieldName
}

func (sf *SearchableField) GetElasticFieldName() string {
	return sf.ElasticFieldName
}

func (sf *SearchableField) GetInternalFieldName() string {
	return sf.InternalFieldName
}

func (sf *SearchableField) HasAnyValue() Query {
	return &Exists{sf.ElasticFieldName}
}

func (sf *SearchableField) Order(order SortOrder) SortItem {
	return SortItem{
		Field:      sf.ElasticFieldName,
		Order:      order,
		NestedPath: nil,
	}
}

type BooleanField struct {
	*SearchableField
	BooleanFieldName string
}

func NewBooleanField(atlanFieldName, booleanFieldName string) *BooleanField {
	searchableField := NewSearchableField(atlanFieldName, booleanFieldName)
	return &BooleanField{
		SearchableField:  searchableField,
		BooleanFieldName: booleanFieldName,
	}
}

func (bf *BooleanField) GetBooleanFieldName() string {
	return bf.BooleanFieldName
}

func (bf *BooleanField) Eq(value bool) Query {
	return &TermQuery{
		Field: bf.BooleanFieldName,
		Value: value,
	}
}

type KeywordField struct {
	*SearchableField
	KeywordFieldName string
}

func NewKeywordField(atlanFieldName, keywordFieldName string) *KeywordField {
	searchableField := NewSearchableField(atlanFieldName, keywordFieldName)
	return &KeywordField{
		SearchableField:  searchableField,
		KeywordFieldName: keywordFieldName,
	}
}

func (kf *KeywordField) GetKeywordFieldName() string {
	return kf.KeywordFieldName
}

func (kf *KeywordTextField) StartsWith(value string, caseInsensitive *bool) Query {
	return &PrefixQuery{
		Field:           kf.KeywordFieldName,
		Value:           value,
		CaseInsensitive: caseInsensitive,
	}
}

func (kf *KeywordTextField) Eq(value string, caseInsensitive *bool) Query {
	return &TermQuery{
		Field: kf.KeywordFieldName,
		Value: value,
	}
}

func (kf *KeywordField) Within(values []string) Query {
	return &Terms{
		Field:  kf.KeywordFieldName,
		Values: values,
	}
}

type KeywordTextField struct {
	*SearchableField
	KeywordFieldName string
	TextFieldName    string
}

func NewKeywordTextField(atlanFieldName, keywordFieldName string, textfieldname string) *KeywordTextField {
	searchableField := NewSearchableField(atlanFieldName, keywordFieldName)
	return &KeywordTextField{
		SearchableField:  searchableField,
		KeywordFieldName: keywordFieldName,
		TextFieldName:    textfieldname,
	}
}

type TextField struct {
	*SearchableField
	TextFieldName string
}

func NewTextField(atlanFieldName, textFieldName string) *TextField {
	searchableField := NewSearchableField(atlanFieldName, textFieldName)
	return &TextField{
		SearchableField: searchableField,
		TextFieldName:   textFieldName,
	}
}

func (tf *TextField) GetTextFieldName() string {
	return tf.TextFieldName
}

func (tf *TextField) Match(value string) Query {
	return &MatchQuery{
		Field: tf.TextFieldName,
		Query: value,
	}
}

type NumericField struct {
	*SearchableField
	NumericFieldName string
}

func NewNumericField(atlanFieldName, numericFieldName string) *NumericField {
	searchableField := NewSearchableField(atlanFieldName, numericFieldName)
	return &NumericField{
		SearchableField:  searchableField,
		NumericFieldName: numericFieldName,
	}
}

func (nf *NumericField) GetNumericFieldName() string {
	return nf.NumericFieldName
}

func (nf *NumericField) Eq(value interface{}) Query {
	return &TermQuery{
		Field: nf.NumericFieldName,
		Value: value,
	}
}

func (nf *NumericField) Gt(value *interface{}) Query {
	return &RangeQuery{
		Field: nf.NumericFieldName,
		Gt:    value,
	}
}

func (nf *NumericField) Gte(value *interface{}) Query {
	return &RangeQuery{
		Field: nf.NumericFieldName,
		Gte:   value,
	}
}

func (nf *NumericField) Lt(value *interface{}) Query {
	return &RangeQuery{
		Field: nf.NumericFieldName,
		Lt:    value,
	}
}

func (nf *NumericField) Lte(value *interface{}) Query {
	return &RangeQuery{
		Field: nf.NumericFieldName,
		Lte:   value,
	}
}

func (nf *NumericField) Between(minimum, maximum *interface{}) Query {
	return &RangeQuery{
		Field: nf.NumericFieldName,
		Gte:   minimum,
		Lte:   maximum,
	}
}

/*
func (nf *NumericField) Avg() Aggregation {
	return Aggregation{
		Root: map[string]interface{}{
			"avg": map[string]interface{}{
				"field": nf.NumericFieldName,
			},
		},
	}
}

func (nf *NumericField) Sum() Aggregation {
	return Aggregation{
		Root: map[string]interface{}{
			"sum": map[string]interface{}{
				"field": nf.NumericFieldName,
			},
		},
	}
}

func (nf *NumericField) Min() Aggregation {
	return Aggregation{
		Root: map[string]interface{}{
			"min": map[string]interface{}{
				"field": nf.NumericFieldName,
			},
		},
	}
}

func (nf *NumericField) Max() Aggregation {
	return Aggregation{
		Root: map[string]interface{}{
			"max": map[string]interface{}{
				"field": nf.NumericFieldName,
			},
		},
	}
}
*/
