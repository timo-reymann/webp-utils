package json

import (
	"context"
	"encoding/json"
	"github.com/gobuffalo/packr"
	"github.com/qri-io/jsonschema"
)

// Schemas contains all available schemas, bundled with packer, the key is the file name with extension
// e. g. cwebp.json
var Schemas = make(map[string]jsonschema.Schema, 0)

func parseSchema() error {
	box := packr.NewBox("../schema")

	for _, schemaFile := range box.List() {
		schemaData, err := box.Find(schemaFile)
		if err != nil {
			return err
		}

		schema := &jsonschema.Schema{}
		if err := json.Unmarshal(schemaData, schema); err != nil {
			return err
		}

		Schemas[schemaFile] = *schema
	}

	return nil
}

// Validate the fileContent as raw json and return all human readable error messages
func Validate(fileContent []byte, schemaName string) []string {
	ctx := context.Background()
	schema := Schemas[schemaName+".json"]
	keyErrors, err := schema.ValidateBytes(ctx, fileContent)
	if err != nil {
		return []string{
			err.Error(),
		}
	}

	errors := make([]string, len(keyErrors))
	for i, keyError := range keyErrors {
		errors[i] = keyError.Error()
	}

	return errors
}
