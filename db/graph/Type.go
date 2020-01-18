package graph

// Type from schema
type Type struct {
	Name   string `json:"name,omitempty"`
	Fields []Type `json:"fields,omitempty"`
}

// Equals checks if two types are equal
func (a Type) Equals(b Type) bool {
	if a.Name != b.Name || len(a.Fields) != len(b.Fields) {
		return false
	}
	for _, v := range a.Fields {
		found := false
		for _, w := range b.Fields {
			if v.Equals(w) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
