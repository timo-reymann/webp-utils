package execute

import (
	"reflect"
	"testing"
)

func TestArgument_IsFlag(t *testing.T) {
	bool := newArgument("isBool", true)
	if !bool.IsFlag() {
		t.Error("Flag not detected")
	}

	noBool := newArgument("isNoBool", "definitelyNot")
	if noBool.IsFlag() {
		t.Error("Flag detected, but is none")
	}
}

func TestArgument_IsFlagEnabled(t *testing.T) {
	bool := newArgument("isBool", true)
	if !bool.IsFlagEnabled() {
		t.Error("Flag enabled not detected correctly")
	}

	bool.Value = false
	if bool.IsFlagEnabled() {
		t.Error("Flag disabled not detected correctly")
	}
}

func TestArgument_BuildArg(t *testing.T) {
	testCases := []struct {
		argumentName  string
		argumentValue interface{}
		expectedArgs  []string
		expectedError string
	}{
		{
			argumentName:  "disabled-flag",
			argumentValue: false,
			expectedArgs:  []string{},
			expectedError: "",
		},
		{
			argumentName:  "enabled-flag",
			argumentValue: true,
			expectedArgs:  []string{"-enabled-flag"},
			expectedError: "",
		},
		{
			argumentName:  "value-flag",
			argumentValue: "value",
			expectedArgs:  []string{"-value-flag", "value"},
			expectedError: "",
		},
		{
			argumentName:  "templated-flag",
			argumentValue: "{{ .important_value }}",
			expectedArgs:  []string{"-templated-flag", "foo"},
			expectedError: "",
		},
	}

	for _, tc := range testCases {
		arg := newArgument(tc.argumentName, tc.argumentValue)
		templater, _ := newArgumentTemplater(
			map[string]interface{}{
				"important_value": "foo",
			}, map[string]interface{}{},
		)
		args, err := arg.BuildArg(templater)

		if !reflect.DeepEqual(tc.expectedArgs, args) {
			t.Errorf("Expected arguments %v, but got %v", tc.expectedArgs, args)
		}

		if tc.expectedError == "" && err != nil {
			t.Errorf("Expected no error, gut %v", err)
		}

		if err == nil && tc.expectedError != "" {
			t.Errorf("Expected error %s but got nothing", tc.expectedError)
		}
	}
}
