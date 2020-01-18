package graph

import (
	"fmt"
	"strings"
)

// SchemaData from schema response
type SchemaData struct {
	Schema []Predicate `json:"schema,omitempty"`
	Types  []Type      `json:"types,omitempty"`
}

func (s SchemaData) formSchema() string {
	var b strings.Builder
	for _, p := range s.Schema {
		if p.Name != "dgraph.type" {
			fmt.Fprintf(&b, "%s: ", p.Name)
			if p.IsList {
				fmt.Fprintf(&b, "[%s]", p.Type)
			} else {
				fmt.Fprintf(&b, "%s", p.Type)
			}
			if p.IsIndex {
				fmt.Fprintf(&b, " @index(%s)", strings.Join(p.Tokenizer, ","))
			}
			if p.HasReverse {
				fmt.Fprint(&b, " @reverse")
			}
			fmt.Fprintln(&b, " .")
		}
	}
	for _, t := range s.Types {
		fmt.Fprintln(&b)
		fmt.Fprintf(&b, "type %s {\n", t.Name)
		for _, f := range t.Fields {
			fmt.Fprintf(&b, "    %s\n", f.Name)
		}
		fmt.Fprintln(&b, "}")
	}
	return b.String()
}
