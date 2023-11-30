package tests

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

var testError = errors.New("test error")

func getEmptyResult() driver.Result {
	return sqlmock.NewResult(0, 0)
}

func checkErrorIsNotNil(t *testing.T, err error) {
	if err == nil {
		t.Error("expected error, got", err)
	}
}

func checkErrorIsNil(t *testing.T, err error) {
	if err != nil {
		t.Error("expected nil, got", err)
	}
}

func checkArrayIsEmpty[T any](t *testing.T, arr []T) {
	if len(arr) != 0 {
		t.Errorf("expected array to be empty, got %v\n", arr)
	}
}

func checkIsFalse(t *testing.T, value bool) {
	if value {
		t.Error("expected false, got true")
	}
}

func checkIsTrue(t *testing.T, value bool) {
	if !value {
		t.Error("expected true, got false")
	}
}
