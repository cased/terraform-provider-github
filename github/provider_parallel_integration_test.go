package github

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccProvider_parallel_requests_performance(t *testing.T) {
	// Skip if not running acceptance tests
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set to 1")
	}

	// This test requires a GitHub token
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		t.Skip("GITHUB_TOKEN must be set for parallel performance test")
	}

	t.Run("performance comparison", func(t *testing.T) {
		// Test fetching multiple repositories to simulate real workload
		// We'll use public repos that definitely exist
		repoSlugs := []string{
			"golang/go",
			"kubernetes/kubernetes",
			"docker/docker",
			"hashicorp/terraform",
			"prometheus/prometheus",
			"grafana/grafana",
			"elastic/elasticsearch",
			"apache/kafka",
			"redis/redis",
			"postgresql/postgresql",
		}

		// Test with parallel_requests = false (serial)
		serialDuration := measureFetchTime(t, token, repoSlugs, false)
		t.Logf("Serial fetch time: %v", serialDuration)

		// Test with parallel_requests = true (parallel)
		parallelDuration := measureFetchTime(t, token, repoSlugs, true)
		t.Logf("Parallel fetch time: %v", parallelDuration)

		// Calculate speedup
		speedup := float64(serialDuration) / float64(parallelDuration)
		t.Logf("Speedup: %.2fx faster with parallel requests", speedup)

		// Assert that parallel is actually faster
		if parallelDuration >= serialDuration {
			t.Errorf("Expected parallel requests to be faster. Serial: %v, Parallel: %v",
				serialDuration, parallelDuration)
		}

		// Assert meaningful speedup (at least 1.5x faster)
		if speedup < 1.5 {
			t.Errorf("Expected at least 1.5x speedup, got %.2fx", speedup)
		}
	})
}

func measureFetchTime(t *testing.T, token string, repoSlugs []string, parallel bool) time.Duration {
	// Create config
	config := Config{
		Token:            token,
		BaseURL:          "https://api.github.com/",
		ParallelRequests: parallel,
	}

	meta, err := config.Meta()
	if err != nil {
		t.Fatalf("Failed to create meta: %v", err)
	}

	client := meta.(*Owner).v3client
	ctx := context.Background()

	start := time.Now()

	if parallel {
		// Parallel fetching
		var wg sync.WaitGroup
		for _, slug := range repoSlugs {
			wg.Add(1)
			go func(repoSlug string) {
				defer wg.Done()
				owner, name := parseRepoSlug(repoSlug)
				_, _, err := client.Repositories.Get(ctx, owner, name)
				if err != nil {
					t.Logf("Error fetching %s: %v", repoSlug, err)
				}
			}(slug)
		}
		wg.Wait()
	} else {
		// Serial fetching
		for _, slug := range repoSlugs {
			owner, name := parseRepoSlug(slug)
			_, _, err := client.Repositories.Get(ctx, owner, name)
			if err != nil {
				t.Logf("Error fetching %s: %v", slug, err)
			}
		}
	}

	return time.Since(start)
}

func parseRepoSlug(slug string) (string, string) {
	// Simple parser for "owner/repo" format
	var owner, name string
	fmt.Sscanf(slug, "%s/%s", &owner, &name)
	if owner == "" || name == "" {
		// Fallback parsing
		for i := 0; i < len(slug); i++ {
			if slug[i] == '/' {
				return slug[:i], slug[i+1:]
			}
		}
	}
	return owner, name
}

func TestAccProvider_parallel_terraform_resources(t *testing.T) {
	// Skip if not running acceptance tests
	if os.Getenv("TF_ACC") != "1" {
		t.Skip("Acceptance tests skipped unless env 'TF_ACC' set to 1")
	}

	// Create multiple resources to test parallel performance
	const configSerial = `
provider "github" {
  parallel_requests = false
}

data "github_repository" "test1" {
  name = "terraform"
  owner = "hashicorp"
}

data "github_repository" "test2" {
  name = "go"
  owner = "golang"
}

data "github_repository" "test3" {
  name = "kubernetes"
  owner = "kubernetes"
}
`

	const configParallel = `
provider "github" {
  parallel_requests = true
}

data "github_repository" "test1" {
  name = "terraform"
  owner = "hashicorp"
}

data "github_repository" "test2" {
  name = "go"
  owner = "golang"
}

data "github_repository" "test3" {
  name = "kubernetes"
  owner = "kubernetes"
}
`

	t.Run("serial performance", func(t *testing.T) {
		start := time.Now()
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: configSerial,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_repository.test1", "name", "terraform"),
					),
				},
			},
		})
		serialDuration := time.Since(start)
		t.Logf("Serial Terraform execution time: %v", serialDuration)
	})

	t.Run("parallel performance", func(t *testing.T) {
		start := time.Now()
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { testAccPreCheck(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: configParallel,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_repository.test1", "name", "terraform"),
					),
				},
			},
		})
		parallelDuration := time.Since(start)
		t.Logf("Parallel Terraform execution time: %v", parallelDuration)
	})
}
