package graph

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/dgraph-io/dgo/protos/api"
)

var types = []reflect.Type{
	reflect.TypeOf(Novel{}),
	reflect.TypeOf(Permission{}),
	reflect.TypeOf(Role{}),
	reflect.TypeOf(Token{}),
	reflect.TypeOf(User{}),
}

func expectedSchema() *SchemaData {
	schema := SchemaData{
		Schema: []Predicate{},
		Types:  []Type{},
	}
	predicates := map[string]Predicate{}
	for _, t := range types {
		fields := t.NumField()
		typeObj := Type{
			Name:   t.Name(),
			Fields: []Type{},
		}
		for i := 0; i < fields; i++ {
			field := t.Field(i)
			tokenizerName, isIndex := field.Tag.Lookup("tokenizer")
			var tokenizer []string
			if isIndex {
				tokenizer = []string{tokenizerName}
			} else {
				tokenizer = []string{}
			}
			_, hasReverse := field.Tag.Lookup("reverse")
			var typename string
			isList := false
			switch field.Type.Kind() {
			case reflect.Bool:
				typename = "bool"
				break
			case reflect.Int:
				typename = "int"
				break
			case reflect.String:
				typename = "string"
				break
			case reflect.Slice:
				isList = true
				if field.Type.Elem().Kind() == reflect.String {
					typename = "string"
				} else {
					typename = "uid"
				}
				break
			case reflect.Ptr:
				switch field.Type.Elem().Name() {
				case "Time":
					typename = "datetime"
					break
				default:
					l.Errorf("Unknown predicate pointer type '%s'.", field.Type)
					return nil
				}
				break
			default:
				l.Errorf("Unknown predicate type '%s' of kind %s.", field.Type, field.Type.Kind().String())
				return nil
			}
			predicate := Predicate{
				Name:       strings.SplitN(field.Tag.Get("json"), ",", 2)[0],
				Type:       typename,
				IsIndex:    isIndex,
				IsList:     isList,
				HasReverse: hasReverse,
				Tokenizer:  tokenizer,
			}
			if other, ok := predicates[predicate.Name]; ok {
				if !predicate.Equals(other) {
					l.Errorf("Predicate '%s' varies between types!", predicate.Name)
				}
			} else {
				schema.Schema = append(schema.Schema, predicate)
				predicates[predicate.Name] = predicate
			}
			if predicate.Name != "dgraph.type" {
				typeObj.Fields = append(typeObj.Fields, Type{
					Name: predicate.Name,
				})
			}
		}
		schema.Types = append(schema.Types, typeObj)
	}
	return &schema
}

// VerifySchema checks to see if the database schema matches the model objects
func VerifySchema(quiet bool) bool {
	resp, err := dg.NewReadOnlyTxn().Query(ctx, "schema{}")
	if err != nil {
		l.Errorf("Unable to query schema: %s", err)
		return false
	}
	var schema SchemaData
	err = json.Unmarshal(resp.Json, &schema)
	if err != nil {
		l.Errorf("Invalid schema JSON: %s", err)
	}
	l.Debugf("Database schema: %s", schema)
	if len(schema.Schema) == 1 && len(schema.Types) == 0 {
		if !quiet {
			l.Error("Database is uninitialized.")
		}
		return false
	}
	expected := expectedSchema()
	if expected == nil {
		l.Error("Unable to verify schema since the reference schema could not be generated.")
		return false
	}
	schemaValid := true
	for _, pd := range schema.Schema {
		found := false
		for _, pe := range expected.Schema {
			if pd.Name == pe.Name {
				found = true
				if !pd.Equals(pe) {
					if !quiet {
						l.Errorf("Schema mismatch for predicate '%s'", pd.Name)
						l.Debugf("Predicate '%s' schema in database: %s", pd.Name, pd)
						l.Debugf("Predicate '%s' schema in model: %s", pe.Name, pe)
					}
					schemaValid = false
				}
				break
			}
		}
		if !found {
			if !quiet {
				l.Errorf("Extra predicate '%s' found in database not in model", pd.Name)
				l.Debugf("Predicate '%s': %s", pd.Name, pd)
			}
			schemaValid = false
		}
	}
	for _, pe := range expected.Schema {
		found := false
		for _, pd := range schema.Schema {
			if pd.Name == pe.Name {
				found = true
				break
			}
		}
		if !found {
			if !quiet {
				l.Errorf("Predicate '%s' from model missing in database", pe.Name)
				l.Debugf("Predicate '%s': %s", pe.Name, pe)
			}
			schemaValid = false
		}
	}
	for _, td := range schema.Types {
		found := false
		for _, te := range expected.Types {
			if td.Name == te.Name {
				found = true
				if !td.Equals(te) {
					if !quiet {
						l.Errorf("Schema mismatch for type '%s'", td.Name)
						l.Debugf("Type '%s' schema in database: %s", td.Name, td)
						l.Debugf("Type '%s' schema in model: %s", te.Name, te)
					}
					schemaValid = false
				}
				break
			}
		}
		if !found {
			if !quiet {
				l.Errorf("Extra type '%s' found in database not in model", td.Name)
				l.Debugf("Type '%s': %s", td.Name, td)
			}
			schemaValid = false
		}
	}
	for _, te := range expected.Types {
		found := false
		for _, td := range schema.Types {
			if td.Name == te.Name {
				found = true
				break
			}
		}
		if !found {
			if !quiet {
				l.Errorf("Type '%s' from model missing in database", te.Name)
				l.Debugf("Type '%s': %s", te.Name, te)
			}
			schemaValid = false
		}
	}
	return schemaValid
}

// ApplySchema creates the schema in the database based on the model classes
func ApplySchema() bool {
	schema := expectedSchema()
	if schema == nil {
		l.Error("Unable to apply schema since the reference schema could not be generated.")
		return false
	}
	err := dg.Alter(ctx, &api.Operation{
		Schema: schema.formSchema(),
	})
	if err != nil {
		l.Errorf("Unable to apply schema: %s", err)
		return false
	}
	return true
}
