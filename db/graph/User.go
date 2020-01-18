package graph

import "time"

// User node type
type User struct {
	DType         []string   `json:"dgraph.type,omitempty" tokenizer:"exact"`
	Username      string     `json:"username,omitempty" tokenizer:"hash"`
	UUID          string     `json:"uuid,omitempty" tokenizer:"hash"`
	Password      string     `json:"password,omitempty"`
	PasswordToken string     `json:"password-token,omitempty" tokenizer:"hash"`
	Email         string     `json:"email,omitempty" tokenizer:"hash"`
	EmailToken    string     `json:"email-token,omitempty" tokenizer:"hash"`
	Phone         string     `json:"phone,omitempty" tokenizer:"hash"`
	PhoneToken    string     `json:"phone-token,omitempty" tokenizer:"hash"`
	Since         *time.Time `json:"since,omitempty"`
	Picture       string     `json:"picture,omitempty"`
	Friends       []User     `json:"friends,omitempty"`               // nickname string, since *time.Time
	Roles         []Role     `json:"roles,omitempty"`                 // priority int
	Novels        []Novel    `json:"novels,omitempty" reverse:"true"` // starred, notify bool, chapter, pos, turn int
}
