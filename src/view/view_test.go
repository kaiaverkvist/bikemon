package view

import (
	"github.com/kaiaverkvist/bikemon/src/config"
	"github.com/yarf-framework/yarf"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestView_Render(t *testing.T) {

	variableToInclude := "isaac newton was born in 1642"

	// Create a fake request.
	var fakeRequest, err = http.NewRequest(http.MethodGet, "/", nil)
	fakeResponse := httptest.NewRecorder()

	// Give a test error if the fake request doesn't work.
	if err != nil {
		t.Error("Unable to return a fake request from http.NewRequest.")
	}

	// Let's use the default config since we're going to be modifying the config slightly.
	viewConfig := config.GetDefaultConfig().ViewConfig

	viewConfig.Folder = "../../testdata"

	// We create a view and populate it as we would during production code.
	testView := New(yarf.NewContext(fakeRequest, fakeResponse), viewConfig)
	testView.Name = "template_test"
	testView.Variables["testVariable"] = variableToInclude

	// Attempts to render the template, and catch any potential errors related to file io.
	err = testView.Render()

	if err != nil {
		t.Errorf("Template rendering error: %s", err)
	}

	// We now test whether the variable inclusion worked successfully.
	responseBody := string(fakeResponse.Body.Bytes())
	if responseBody != "This template has a variable which equals: isaac newton was born in 1642" {
		t.Errorf("Template variable inclusion did not work correctly (expected: %s, actual: %s)", variableToInclude, responseBody)
	}
}