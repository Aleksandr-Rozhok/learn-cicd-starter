package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	type Tests struct {
		name     string
		inputMap map[string][]string
		expected string
	}

	tests := []Tests{
		{
			name: "Check getting api key",
			inputMap: http.Header{
				"Content-Type":    {"application/json"},
				"Authorization":   {"ApiKey 12345abcde67890fghijKLMNOPqrsTUVwxyz_9876543210"},
				"User-Agent":      {"Go-http-client/1.1"},
				"Accept-Language": {"en-US,en;q=0.9"},
			},
			expected: "12345abcde67890fghijKLMNOPqrsTUVwxyz_9876543210",
		},
		{
			name: "Check empty api key",
			inputMap: http.Header{
				"Content-Type":    {"application/json"},
				"Authorization":   {""},
				"User-Agent":      {"Go-http-client/1.1"},
				"Accept-Language": {"en-US,en;q=0.9"},
			},
			expected: "no authorization header included",
		},
		{
			name: "Check malformed authorization header",
			inputMap: http.Header{
				"Content-Type":    {"application/json"},
				"Authorization":   {"Bearer 12345abcde67890fghijKLMNOPqrsTUVwxyz_9876543210"},
				"User-Agent":      {"Go-http-client/1.1"},
				"Accept-Language": {"en-US,en;q=0.9"},
			},
			expected: "malformed authorization header",
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		actual, err := auth.GetAPIKey(tests[0].inputMap)
		if err != nil {
			t.Errorf("Test %v - '%s' FAIL: Error: %s", 1, tests[0].name, err)
		}
		if reflect.DeepEqual(actual, tests[0].expected) == false {
			t.Errorf("Test %v - '%s' FAIL: expected '%v', got '%v'", 1, tests[0].name, tests[0].expected, actual)
		}
	})

	t.Run(tests[1].name, func(t *testing.T) {
		_, err := auth.GetAPIKey(tests[1].inputMap)
		if err != nil {
			fmt.Printf("Expected: '%s', Got: '%s'\n", tests[1].expected, err.Error())
			if strings.TrimSpace(err.Error()) != strings.TrimSpace(tests[1].expected) {
				t.Errorf("Test %v - '%s' FAIL: expected '%v', got '%v'", 2, tests[1].name, tests[1].expected, err.Error())
			}
		} else {
			t.Errorf("Test %v - '%s' FAIL: expected error, got nil", 2, tests[1].name)
		}
	})

	t.Run(tests[2].name, func(t *testing.T) {
		actual, err := auth.GetAPIKey(tests[2].inputMap)
		if reflect.DeepEqual(err.Error(), tests[2].expected) == false && actual == "" {
			t.Errorf("Test %v - '%s' FAIL: expected '%v', got '%v'", 3, tests[2].name, tests[2].expected, err.Error())
		}
	})
}
