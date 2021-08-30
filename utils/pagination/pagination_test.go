package pagination_test

import (
	"errors"
	"kardashian_api/utils/pagination"
	"testing"
)

func TestValidate_SkipLessThanTotal(t *testing.T) {
	var expectation error
	actual := pagination.Validate(0, 10)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestValidate_SkipEqualTotal(t *testing.T) {
	var expectation error
	actual := pagination.Validate(10, 10)
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestValidate_SkipGreaterThanTotal(t *testing.T) {
	var expectation error = errors.New("page out of range")
	actual := pagination.Validate(15, 10)
	if actual.Error() != expectation.Error() {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestParams_0(t *testing.T) {
	var expected = []int{1, 0, 10}
	page, skip, limit := pagination.Params("0")
	if page != expected[0] {
		t.Errorf("Expected page %v but got %v", expected[0], page)
	}
	if skip != expected[1] {
		t.Errorf("Expected skip %v but got %v", expected[1], page)
	}
	if limit != expected[2] {
		t.Errorf("Expected limit %v but got %v", expected[2], page)
	}
}

func TestParams_1(t *testing.T) {
	var expected = []int{1, 0, 10}
	page, skip, limit := pagination.Params("1")
	if page != expected[0] {
		t.Errorf("Expected page %v but got %v", expected[0], page)
	}
	if skip != expected[1] {
		t.Errorf("Expected skip %v but got %v", expected[1], page)
	}
	if limit != expected[2] {
		t.Errorf("Expected limit %v but got %v", expected[2], page)
	}
}

func TestParams_5(t *testing.T) {
	var expected = []int{5, 40, 10}
	page, skip, limit := pagination.Params("5")
	if page != expected[0] {
		t.Errorf("Expected page %v but got %v", expected[0], page)
	}
	if skip != expected[1] {
		t.Errorf("Expected skip %v but got %v", expected[1], skip)
	}
	if limit != expected[2] {
		t.Errorf("Expected limit %v but got %v", expected[2], limit)
	}
}
