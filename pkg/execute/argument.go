package execute

import (
	"fmt"
)

// Argument represents a high level command line argument to webp tools
type Argument struct {
	Name  string
	Value interface{}
}

func newArgument(name string, value interface{}) Argument {
	return Argument{
		Name:  name,
		Value: value,
	}
}

// IsFlag checks if the argument is a flag (type boolean)
func (a *Argument) IsFlag() bool {
	_, isBool := a.Value.(bool)
	return isBool
}

// IsFlagEnabled checks if the argument is a flag of type enable
// it relies that you previously checked that the argument is actually boolean using the
// IsFlag method
func (a *Argument) IsFlagEnabled() bool {
	value, _ := a.Value.(bool)
	return value
}

func (a *Argument) getStringValue() string {
	return fmt.Sprintf("%v", a.Value)
}

func (a *Argument) processValueTemplate(templater *ArgumentTemplater) (string, error) {
	strVal := a.getStringValue()
	return templater.Apply(strVal)
}

// BuildArg creates a array containing the parameter name and optional the templated processed value
func (a *Argument) BuildArg(templater *ArgumentTemplater) ([]string, error) {
	arg := []string{
		"-" + a.Name,
	}

	// Set flag accordingly without variable or omit
	if a.IsFlag() {
		if a.IsFlagEnabled() {
			return arg, nil
		}

		// omit
		return []string{}, nil
	}

	// Try to process parameter value template
	processedValue, err := a.processValueTemplate(templater)
	if err != nil {
		return nil, err
	}

	return append(arg, processedValue), nil
}
