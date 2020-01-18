package graph

// Permission node type
type Permission struct {
	DType []string `json:"dgraph.type,omitempty" tokenizer:"exact"`
	Name  string   `json:"name,omitempty"`
}
