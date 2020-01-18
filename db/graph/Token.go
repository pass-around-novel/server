package graph

import "time"

// Token node type
type Token struct {
	DType []string   `json:"dgraph.type,omitempty" tokenizer:"exact"`
	Token string     `json:"token,omitempty" tokenizer:"hash"`
	Since *time.Time `json:"since,omitempty"`
	Until *time.Time `json:"until,omitempty"`
	User  []User     `json:"user,omitempty" reverse:"true"`
}
