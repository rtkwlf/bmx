package bmx_test

import (
	"testing"

	"github.com/rtkwlf/bmx"
	"github.com/rtkwlf/bmx/saml/identityProviders/okta"
)

func TestFindAppByLabels(t *testing.T) {
	// Need 2 things, name of label and name of OktaApp
	var expectedLabels = []struct {
		testName    string
		searchName  string
		targetLabel string
		isFound     bool
	}{
		{"samecase", "TestApp3", "TestApp3", true},
		{"lowercase", "testapp3", "TestApp3", true},
		{"uppercase", "TESTAPP3", "TestApp3", true},
		{"closebutnot", "TestAppp3", "TestApp3", false},
		{"differentnumber", "TestApp2", "TestApp3", false},
		{"emptystring", "", "TestApp3", false},
	}

	dataset := []okta.OktaAppLink{
		okta.OktaAppLink{Label: "awsAppTesting1"},
		okta.OktaAppLink{Label: "myTestApp2"},
		okta.OktaAppLink{Label: "testingForAws"},
		okta.OktaAppLink{Label: "MyAwsApp"},
		okta.OktaAppLink{Label: "NotAnAwsApp"},
	}

	for _, test := range expectedLabels {
		t.Run(test.testName, func(t *testing.T) {
			expected := okta.OktaAppLink{Label: test.targetLabel}
			applinks := append(dataset, expected)

			actual, ok := bmx.FindAppByLabel(test.searchName, applinks)
			if ok && !test.isFound {
				t.Errorf("Found result of: %s, when expected nothing to be found", actual.Label)
			} else if !ok && test.isFound {
				t.Errorf("Did not find: %s, when result: %s was expected", test.searchName, test.targetLabel)
			}

			if ok && actual.Label != expected.Label {
				t.Errorf("Label is incorrect, got: %s, expected %s", actual.Label, expected.Label)
			}
		})
	}
}
