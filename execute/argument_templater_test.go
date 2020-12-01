package execute

import "testing"

func TestArgumentTemplater_IsExtraParameter(t *testing.T) {
	templater, _ := newArgumentTemplater(
		map[string]interface{}{
			"regular_value": "foo",
		}, map[string]interface{}{
			"extra_value": "foo",
		},
	)

	if ! templater.IsExtraParameter("extra_value") {
		t.Error("Expected extra paramater to be detected")
	}

	if templater.IsExtraParameter("regular_value") {
		t.Error("Expected regular paramater not to be detected")
	}
}
