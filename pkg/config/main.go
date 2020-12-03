package config

import (
	"errors"
	"github.com/timo-reymann/webp-utils/pkg/json"
	"io/ioutil"
)

// Load reads the file and validates it against the schema
func Load(path string, schema string) (map[string]interface{}, error) {
	// Parse configuration
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Validate configuration
	errs := json.Validate(content, schema)
	for _, err := range errs {
		return nil, errors.New(err)
	}

	// Parse json
	return json.Parse(content)
}
