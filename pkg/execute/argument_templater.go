package execute

import (
	"bytes"
	"text/template"
)

var argTemplate = template.New("")

// ArgumentTemplater contains all information necessary to make a argument templateable
type ArgumentTemplater struct {
	context            map[string]interface{}
	extraParameterKeys []string
}

func newArgumentTemplater(context map[string]interface{}, extraParameters map[string]interface{}) (*ArgumentTemplater, error) {
	excludedArgs := make([]string, 0)
	templateContext := make(map[string]interface{})

	// Copy context
	for key, val := range context {
		templateContext[key] = val
	}

	// Copy extra parameters
	for key, val := range extraParameters {
		templateContext[key] = val
		excludedArgs = append(excludedArgs, key)
	}

	return &ArgumentTemplater{
		context:            templateContext,
		extraParameterKeys: excludedArgs,
	}, nil
}

// Apply executes the argument template on a given string template and returns the processed result
func (at *ArgumentTemplater) Apply(template string) (string, error) {
	parsedTpl, err := argTemplate.Parse(template)
	if err != nil {
		return "", err
	}

	var stringTemplate bytes.Buffer
	if err := parsedTpl.Execute(&stringTemplate, at.context); err != nil {
		return "", err
	}

	return stringTemplate.String(), nil
}

// IsExtraParameter checks if the key provided is not part of the cli arguments but rather
// purely informational to be used in templates
func (at *ArgumentTemplater) IsExtraParameter(key string) bool {
	for _, extraParameterKey := range at.extraParameterKeys {
		if extraParameterKey == key {
			return true
		}
	}

	return false
}
