package graph

// Role node type
type Role struct {
	DType       []string     `json:"dgraph.type,omitempty" tokenizer:"exact"`
	Name        string       `json:"name,omitempty"`
	Permissions []Permission `json:"perms,omitempty"` // value interface{}
}
