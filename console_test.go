package bmx_test

import (
	"testing"

	"github.com/rtkwlf/bmx"
	"github.com/rtkwlf/bmx/saml/identityProviders/okta"
)

func TestFindAppByLabels(t *testing.T) {
	name := "nameToFind"
	expected := okta.OktaAppLink{Label: name}
	applinks := []okta.OktaAppLink{
		okta.OktaAppLink{Label: "myname"},
		okta.OktaAppLink{Label: "myname"},
		expected,
	}
	actual, ok := bmx.FindAppByLabel(name, applinks)
	if !ok {
		t.Errorf("Failed to find app, got: %v, expected %v", ok, true)
	}

	if name != actual.Label {
		t.Errorf("Label is incorrect, got: %s, expected %s", actual.Label, expected.Label)
	}
}
