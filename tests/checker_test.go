package tests

import (
	"TalkHub/pkg/checker"
	"testing"
)

type TestCaseBool struct {
	inputData string
	expected  bool
}

func TestIsValidEmail(t *testing.T) {
	tests := []TestCaseBool{
		{"", false},
		{"a", false},
		{"abcdef", false},
		{"1234567890", false},
		{"qwerty", false},
		{"abc123", false},
		{"a@123", false},
		{"email@.", false},
		{"email123@.2", false},
		{"@mail.ru", false},
		{"@.ru", false},
		{"test@.com", false},
		{"mail@mail.", false},
		{"test_email.ru", false},
		{"test.email.ru", false},
		{"test_email@mail.ru", true},
		{"email@gmail.com", true},
		{"123abcd@mail.ru", true},
	}
	var valid bool
	for _, test := range tests {
		valid = checker.IsValidEmail(test.inputData)
		if valid != test.expected {
			t.Errorf("input data: %s, expected %t, got %t\n", test.inputData, test.expected, valid)
		}
	}
}

func TestIsValidPassword(t *testing.T) {
	tests := []TestCaseBool{
		{"", false},
		{"abcdef", false},
		{"123qwerty", false},
		{"testCase_", false},
		{"_userTest", false},
		{"password333", true},
		{"testPass", true},
		{"a1v2s3f4", true},
		{"successfulPassword", true},
		{"WORKANDNormal", true},
		{"ESSENCES", true},
		{"SSS1234567890", true},
	}
	var valid bool
	for _, test := range tests {
		valid = checker.IsValidPassword(test.inputData)
		if valid != test.expected {
			t.Errorf("input data: %s, expected %t, got %t\n", test.inputData, test.expected, valid)
		}
	}
}
