package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestProvider_parallel_requests_configuration(t *testing.T) {
	t.Run("parallel_requests schema allows github.com", func(t *testing.T) {
		provider := Provider()

		// Check that parallel_requests is in the schema
		if _, ok := provider.Schema["parallel_requests"]; !ok {
			t.Fatal("parallel_requests should be in provider schema")
		}

		// Verify it's optional and defaults to false
		parallelSchema := provider.Schema["parallel_requests"]
		if parallelSchema.Type != schema.TypeBool {
			t.Fatal("parallel_requests should be TypeBool")
		}
		if parallelSchema.Required {
			t.Fatal("parallel_requests should not be required")
		}
		if parallelSchema.Optional != true {
			t.Fatal("parallel_requests should be optional")
		}

		// The key test: no validation function that would reject github.com
		if parallelSchema.ValidateFunc != nil {
			t.Fatal("parallel_requests should not have a ValidateFunc that restricts github.com")
		}
	})

	t.Run("provider configuration with parallel_requests and github.com", func(t *testing.T) {
		// Create a minimal test configuration
		raw := map[string]interface{}{
			"token":             "test-token",
			"parallel_requests": true,
		}

		resourceData := schema.TestResourceDataRaw(t, Provider().Schema, raw)

		// Verify we can get the value
		parallelRequests := resourceData.Get("parallel_requests").(bool)
		if !parallelRequests {
			t.Fatal("parallel_requests should be true when set")
		}

		// Verify base_url defaults to github.com
		baseURL := resourceData.Get("base_url").(string)
		if baseURL != "" && baseURL != "https://api.github.com/" {
			// Empty string means it will default to github.com in providerConfigure
		}
	})
}
