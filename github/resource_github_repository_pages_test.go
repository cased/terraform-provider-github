package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryPages_EnableOnExisting(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("adds pages to an existing repository", func(t *testing.T) {
		// Step 1: Create repo without pages
		configWithoutPages := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-pages-%s"
				auto_init    = true
				visibility   = "public"
			}
		`, randomID)

		// Step 2: Add pages to existing repo
		configWithPages := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-pages-%s"
				auto_init    = true
				visibility   = "public"
				pages {
					source {
						branch = "main"
					}
				}
			}
		`, randomID)

		// Step 3: Update pages configuration
		configWithUpdatedPages := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-pages-%s"
				auto_init    = true  
				visibility   = "public"
				pages {
					source {
						branch = "main"
						path   = "/docs"
					}
				}
			}
		`, randomID)

		// Step 4: Remove pages
		configWithoutPagesAgain := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-pages-%s"
				auto_init    = true
				visibility   = "public"
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configWithoutPages,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_repository.test", "name",
								fmt.Sprintf("tf-acc-pages-%s", randomID),
							),
						),
					},
					{
						Config: configWithPages,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_repository.test", "pages.0.source.0.branch",
								"main",
							),
							resource.TestCheckResourceAttr(
								"github_repository.test", "pages.0.source.0.path",
								"/",
							),
						),
					},
					{
						Config: configWithUpdatedPages,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_repository.test", "pages.0.source.0.branch",
								"main",
							),
							resource.TestCheckResourceAttr(
								"github_repository.test", "pages.0.source.0.path",
								"/docs",
							),
						),
					},
					{
						Config: configWithoutPagesAgain,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckNoResourceAttr(
								"github_repository.test", "pages.0.source.0.branch",
							),
						),
					},
				},
			})
		}

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}