package reforger

import (
	"context"
	_ "embed"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

//go:embed config.schema.json
var CONFIG_JSON_SCHEMA string

type config struct {
	schema string
}

type ConfigValidator interface {
	Validate(context.Context, interface{}) error
}

// NewConfigValidator returns a new instance of a ConfigValidator
func NewConfigValidator() ConfigValidator {
	return config{schema: CONFIG_JSON_SCHEMA}
}

// Validate validates the given Reforger config against the schema
func (s config) Validate(ctx context.Context, v interface{}) error {
	sch := jsonschema.MustCompileString("schema.json", s.schema)
	return sch.Validate(v)
}
