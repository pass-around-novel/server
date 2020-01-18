package graph

// Predicate from schema
type Predicate struct {
	Name       string   `json:"predicate,omitempty"`
	Type       string   `json:"type,omitempty"`
	IsIndex    bool     `json:"index,omitempty"`
	IsList     bool     `json:"list,omitempty"`
	HasReverse bool     `json:"reverse,omitempty"`
	Tokenizer  []string `json:"tokenizer,omitempty"`
}

// Equals checks if two predicates are equal
func (a Predicate) Equals(b Predicate) bool {
	if a.Name == b.Name &&
		a.Type == b.Type &&
		a.IsIndex == b.IsIndex &&
		a.IsList == b.IsList &&
		a.HasReverse == b.HasReverse &&
		len(a.Tokenizer) == len(b.Tokenizer) {
		for i, v := range a.Tokenizer {
			if v != b.Tokenizer[i] {
				return false
			}
		}
		return true
	}
	return false
}
