package graph

import "time"

// Novel node type
type Novel struct {
	DType       []string   `json:"dgraph.type,omitempty" tokenizer:"exact"`
	Name        string     `json:"name,omitempty"`
	UUID        string     `json:"uuid,omitempty" tokenizer:"hash"`
	NumChapters int        `json:"num-chapters,omitempty"`
	Since       *time.Time `json:"since,omitempty"`
	LastTurn    *time.Time `json:"last-turn,omitempty"`
	Done        bool       `json:"done,omitempty"`
	Picture     string     `json:"picture,omitempty"`
	Public      bool       `json:"public,omitempty"`
	TurnMin     int        `json:"turn-min,omitempty"`
	TurnMax     int        `json:"turn-max,omitempty"`
	TurnType    int        `json:"turn-type,omitempty"`
}
