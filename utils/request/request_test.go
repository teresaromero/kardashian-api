package request_test

import (
	"kardashian_api/utils/request"
	"net/http/httptest"
	"testing"
)

func TestAddToContext(t *testing.T) {
	var reqTest = httptest.NewRequest("GET", "https://example.com", nil)
	expectation := 3
	updatedRequest := request.AddToContext(reqTest, "test", expectation)
	actual := updatedRequest.Context().Value("test")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestGetContextValue_EmptyKey(t *testing.T) {
	var reqTest = httptest.NewRequest("GET", "https://example.com", nil)
	actual := request.GetContextValue(reqTest, "test")
	if actual != nil {
		t.Errorf("Expected %v but got %v", nil, actual)
	}
}

func TestGetContextValue_ExistKey(t *testing.T) {
	var reqTest = httptest.NewRequest("GET", "https://example.com", nil)
	expectation := 3
	reqTest = request.AddToContext(reqTest, "test", expectation)
	actual := request.GetContextValue(reqTest, "test")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestGetURI_NoQueryParams(t *testing.T) {
	var reqTest = httptest.NewRequest("GET", "https://example.com/", nil)
	expectation := "https://example.com"
	actual := request.GetURI(reqTest)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestGetURI_QueryParams(t *testing.T) {
	var reqTest = httptest.NewRequest("GET", "https://example.com/?test=1", nil)
	expectation := "https://example.com"
	actual := request.GetURI(reqTest)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
