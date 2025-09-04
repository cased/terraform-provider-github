package github

import (
	"reflect"
	"testing"

	"github.com/google/go-github/v66/github"
)

func TestFlattenSecurityAndAnalysis(t *testing.T) {
	tests := []struct {
		name     string
		input    *github.SecurityAndAnalysis
		expected []interface{}
	}{
		{
			name:     "nil input returns empty",
			input:    nil,
			expected: []interface{}{},
		},
		{
			name: "empty SecurityAndAnalysis returns empty",
			input: &github.SecurityAndAnalysis{},
			expected: []interface{}{},
		},
		{
			name: "only advanced security present",
			input: &github.SecurityAndAnalysis{
				AdvancedSecurity: &github.AdvancedSecurity{
					Status: github.String("enabled"),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"advanced_security": []interface{}{
						map[string]interface{}{
							"status": "enabled",
						},
					},
				},
			},
		},
		{
			name: "all fields present",
			input: &github.SecurityAndAnalysis{
				AdvancedSecurity: &github.AdvancedSecurity{
					Status: github.String("enabled"),
				},
				SecretScanning: &github.SecretScanning{
					Status: github.String("enabled"),
				},
				SecretScanningPushProtection: &github.SecretScanningPushProtection{
					Status: github.String("enabled"),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					"advanced_security": []interface{}{
						map[string]interface{}{
							"status": "enabled",
						},
					},
					"secret_scanning": []interface{}{
						map[string]interface{}{
							"status": "enabled",
						},
					},
					"secret_scanning_push_protection": []interface{}{
						map[string]interface{}{
							"status": "enabled",
						},
					},
				},
			},
		},
		{
			name: "secret scanning fields with nil status are omitted",
			input: &github.SecurityAndAnalysis{
				AdvancedSecurity: &github.AdvancedSecurity{
					Status: github.String("disabled"),
				},
				SecretScanning:               &github.SecretScanning{},
				SecretScanningPushProtection: &github.SecretScanningPushProtection{},
			},
			expected: []interface{}{
				map[string]interface{}{
					"advanced_security": []interface{}{
						map[string]interface{}{
							"status": "disabled",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenSecurityAndAnalysis(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("flattenSecurityAndAnalysis() = %v, want %v", result, tt.expected)
			}
		})
	}
}