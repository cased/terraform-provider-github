# GitHub Terraform Provider - Priority Issues Report

*Generated automatically by scripts/prioritize_issues.sh*

## Priority Scoring Methodology

Each issue is scored based on:
- **Community Engagement** (40%): Reactions (2x weight) + Comments
- **Issue Type** (20%): Bugs get 20 points, Features get 10 points
- **Labels** (20%): Critical labels like "Status: Pinned", authentication, security
- **Age Factor** (20%): Older high-engagement issues indicate persistent problems

### Priority Levels
- **P0 (Critical)**: Score 70+ - Immediate attention needed
- **P1 (High)**: Score 40-69 - Should be addressed soon  
- **P2 (Medium)**: Score 20-39 - Important but not urgent
- **P3 (Low)**: Score <20 - Nice to have

---

## P0 - Critical Priority Issues

### 294. [[Feature Request] Manage GitHub Apps via Terraform](https://github.com/integrations/terraform-provider-github/issues/509)
- **Issue #**: 509
- **Type**: Feature
- **Score**: 294 (👍 317 reactions, 💬 28 comments)
- **Labels**: Type: Feature, New resource
- **Created**: 2020-07-07

### ~~146. [continual changes detected due to changing etag](https://github.com/integrations/terraform-provider-github/issues/796)~~ ❌ BLOCKED
- **Issue #**: 796
- **Type**: Bug
- **Score**: 146 (👍 113 reactions, 💬 41 comments)
- **Labels**: Type: Bug, Status: Up for grabs, r/repository, hacktoberfest
- **Created**: 2021-05-20
- **Note**: Requires migration to Plugin Framework for private state support

### ~~145. [Slow performance when managing dozens of repositories](https://github.com/integrations/terraform-provider-github/issues/567)~~ ✅ FIXED
- **Issue #**: 567
- **Type**: Bug
- **Score**: 145 (👍 93 reactions, 💬 53 comments)
- **Labels**: Type: Bug, r/branch_protection, Status: Pinned
- **Created**: 2020-10-14
- **Fix**: Enabled parallel_requests for github.com (~10x speedup)

### 106. [[FEAT]: Copilot support in rulesets](https://github.com/integrations/terraform-provider-github/issues/2583)
- **Issue #**: 2583
- **Type**: Feature
- **Score**: 106 (👍 105 reactions, 💬 6 comments)
- **Labels**: Type: Feature, Status: Triage
- **Created**: 2025-03-05

### 106. [[FEAT]: Support GitHub repository custom properties](https://github.com/integrations/terraform-provider-github/issues/1956)
- **Issue #**: 1956
- **Type**: Feature
- **Score**: 106 (👍 88 reactions, 💬 27 comments)
- **Labels**: Type: Feature, Status: Blocked
- **Created**: 2023-10-13

### ~~97. [[BUG]: 422 Secret scanning is not available for this repository.](https://github.com/integrations/terraform-provider-github/issues/2145)~~ ✅ FIXED
- **Issue #**: 2145
- **Type**: Bug
- **Score**: 97 (👍 49 reactions, 💬 45 comments)
- **Labels**: Type: Bug, Status: Up for grabs
- **Created**: 2024-02-13
- **Fix**: Handle nil security_and_analysis fields from GitHub API

### 88. [create new fork repo resource](https://github.com/integrations/terraform-provider-github/issues/152)
- **Issue #**: 152
- **Type**: Feature
- **Score**: 88 (👍 59 reactions, 💬 3 comments)
- **Labels**: Type: Feature, New resource, Status: Pinned
- **Created**: 2018-09-10

### 83. [[NEW FEATURE] Allow forking other github repositories to init new repositories.](https://github.com/integrations/terraform-provider-github/issues/43)
- **Issue #**: 43
- **Type**: Feature
- **Score**: 83 (👍 44 reactions, 💬 21 comments)
- **Labels**: Type: Feature, New resource, Status: Pinned
- **Created**: 2017-08-01

### 83. [Update to Support Scheduled Reminders](https://github.com/integrations/terraform-provider-github/issues/722)
- **Issue #**: 722
- **Type**: Feature
- **Score**: 83 (👍 57 reactions, 💬 7 comments)
- **Labels**: Type: Feature, New resource, Status: Blocked, Status: Pinned
- **Created**: 2021-03-09

### 81. [[FEAT]: Create Github App](https://github.com/integrations/terraform-provider-github/issues/2167)
- **Issue #**: 2167
- **Type**: Feature
- **Score**: 81 (👍 63 reactions, 💬 3 comments)
- **Labels**: Type: Feature
- **Created**: 2024-02-23

### 81. [[FEAT]: Support Org Copilot settings](https://github.com/integrations/terraform-provider-github/issues/1629)
- **Issue #**: 1629
- **Type**: Feature
- **Score**: 81 (👍 58 reactions, 💬 12 comments)
- **Labels**: Type: Feature
- **Created**: 2023-03-31

### 77. [Add ignoring feature for archived repositories](https://github.com/integrations/terraform-provider-github/issues/737)
- **Issue #**: 737
- **Type**: Feature
- **Score**: 77 (👍 43 reactions, 💬 7 comments)
- **Labels**: Type: Feature, Status: Up for grabs, r/label, Status: Pinned
- **Created**: 2021-03-24

### 73. [Add ability to opt out of listing branches in github_repository](https://github.com/integrations/terraform-provider-github/issues/1010)
- **Issue #**: 1010
- **Type**: Feature
- **Score**: 73 (👍 36 reactions, 💬 12 comments)
- **Labels**: Type: Feature, Status: Up for grabs, Status: Pinned
- **Created**: 2021-12-14

### 73. [[BUG]: 5.43 upgrade failing due to signoff issues](https://github.com/integrations/terraform-provider-github/issues/2077)
- **Issue #**: 2077
- **Type**: Bug
- **Score**: 73 (👍 36 reactions, 💬 12 comments)
- **Labels**: Type: Bug, Status: Up for grabs
- **Created**: 2024-01-04

### 71. [Cannot enable pages on an existing repository](https://github.com/integrations/terraform-provider-github/issues/777)
- **Issue #**: 777
- **Type**: Bug
- **Score**: 71 (👍 23 reactions, 💬 7 comments)
- **Labels**: Type: Bug, Status: Up for grabs, r/repository, Status: Pinned, hacktoberfest
- **Created**: 2021-05-05

---

## P1 - High Priority Issues

### 69. [oauth app creation resource for github](https://github.com/integrations/terraform-provider-github/issues/786)
- **Issue #**: 786
- **Type**: Feature
- **Score**: 69 (👍 50 reactions, 💬 12 comments)
- **Labels**: Type: Feature, New resource, Status: Blocked

### 69. [Provider config not inherited](https://github.com/integrations/terraform-provider-github/issues/846)
- **Issue #**: 846
- **Type**: Bug
- **Score**: 69 (👍 20 reactions, 💬 9 comments)
- **Labels**: Type: Bug, Provider, Status: Pinned

### 69. [`vulnerability_alerts` set to true does not enable "Dependabot security updates"](https://github.com/integrations/terraform-provider-github/issues/588)
- **Issue #**: 588
- **Type**: Bug
- **Score**: 69 (👍 28 reactions, 💬 17 comments)
- **Labels**: Type: Bug, r/repository, Status: Needs info

### 69. [Data source github_actions_public_key returns a 404](https://github.com/integrations/terraform-provider-github/issues/667)
- **Issue #**: 667
- **Type**: Bug
- **Score**: 69 (👍 23 reactions, 💬 27 comments)
- **Labels**: Type: Bug, Status: Up for grabs, d/actions_public_key

### 67. [Editing github_actions_secret from TF apply fails](https://github.com/integrations/terraform-provider-github/issues/810)
- **Issue #**: 810
- **Type**: Bug
- **Score**: 67 (👍 18 reactions, 💬 8 comments)
- **Labels**: Type: Bug, Status: Up for grabs, r/actions_secret, Status: Pinned

### 67. [[BUG] github_branch_protection_v3 tries to delete the contexts every time](https://github.com/integrations/terraform-provider-github/issues/1701)
- **Issue #**: 1701
- **Type**: Bug
- **Score**: 67 (👍 12 reactions, 💬 19 comments)
- **Labels**: Type: Bug, Status: Up for grabs, Status: Pinned, vNext, hacktoberfest

### 66. [Push restrictions aren't applied properly on branches](https://github.com/integrations/terraform-provider-github/issues/670)
- **Issue #**: 670
- **Type**: Bug
- **Score**: 66 (👍 16 reactions, 💬 10 comments)
- **Labels**: Type: Bug, Status: Up for grabs, r/branch_protection, Status: Pinned

### 66. [[BUG]: unable to add github apps to bypassers](https://github.com/integrations/terraform-provider-github/issues/2236)
- **Issue #**: 2236
- **Type**: Bug
- **Score**: 66 (👍 27 reactions, 💬 12 comments)
- **Labels**: Type: Bug, Status: Up for grabs

### 65. [app_auth credentials expire after an hour](https://github.com/integrations/terraform-provider-github/issues/977)
- **Issue #**: 977
- **Type**: Bug
- **Score**: 65 (👍 6 reactions, 💬 15 comments)
- **Labels**: Type: Bug, Status: Up for grabs, Provider, Authentication, p/app_auth, Status: Pinned

### 65. [Failed setting visibility=internal on a newly created repo from template](https://github.com/integrations/terraform-provider-github/issues/925)
- **Issue #**: 925
- **Type**: Bug
- **Score**: 65 (👍 15 reactions, 💬 8 comments)
- **Labels**: Type: Bug, Status: Up for grabs, Visibility, Status: Pinned

---

## P2 - Medium Priority Issues

### [[BUG]: Plan incorrectly proposes change for repository ruleset code scanning results](https://github.com/integrations/terraform-provider-github/issues/2556)
- **Issue #**: 2556 | **Type**: Bug | **Score**: 39

### [feat: Enable custom protection rule on GH environment](https://github.com/integrations/terraform-provider-github/pull/2352)
- **Issue #**: 2352 | **Type**: Other | **Score**: 38

### [Error: This resource can only be used in the context of an organization, "foo" is a user](https://github.com/integrations/terraform-provider-github/issues/769)
- **Issue #**: 769 | **Type**: Bug | **Score**: 38

### [Require all status checks to pass without specifying their names](https://github.com/integrations/terraform-provider-github/issues/1142)
- **Issue #**: 1142 | **Type**: Feature | **Score**: 38

### [[FEAT]: Support filter parameter on github_organization_team_sync_groups data source](https://github.com/integrations/terraform-provider-github/issues/1809)
- **Issue #**: 1809 | **Type**: Feature | **Score**: 38

### [[BUG]: Terraform plan hangs/freezes when used in a GitHub Actions workflow AND auth is based on GitHub App](https://github.com/integrations/terraform-provider-github/issues/2241)
- **Issue #**: 2241 | **Type**: Bug | **Score**: 38

### [[BUG]: Allow setting review notifications without delegation](https://github.com/integrations/terraform-provider-github/issues/2273)
- **Issue #**: 2273 | **Type**: Bug | **Score**: 38

### [[BUG]: github_repository_collaborators - Permission replacement where in-place modification would be sufficient](https://github.com/integrations/terraform-provider-github/issues/1646)
- **Issue #**: 1646 | **Type**: Bug | **Score**: 37

### [[BUG]: `github_issue_labels` can't clobber existing labels](https://github.com/integrations/terraform-provider-github/issues/2089)
- **Issue #**: 2089 | **Type**: Bug | **Score**: 37

### [[BUG]: integration_id value returns difference that can't be ignored](https://github.com/integrations/terraform-provider-github/issues/2317)
- **Issue #**: 2317 | **Type**: Bug | **Score**: 37

### [[FEAT]:  GitHub App creation via terraform](https://github.com/integrations/terraform-provider-github/issues/2389)
- **Issue #**: 2389 | **Type**: Feature | **Score**: 37

### [[BUG]: The `github_branch_protection_v3` resource churns when using `checks` or `contexts` in `required_status_checks`](https://github.com/integrations/terraform-provider-github/issues/2493)
- **Issue #**: 2493 | **Type**: Bug | **Score**: 37

### [[BUG]: 422 Validation Failed [] on api.github.com/organizations/XXX/team/XXX/repos/XXX/YYY](https://github.com/integrations/terraform-provider-github/issues/2198)
- **Issue #**: 2198 | **Type**: Bug | **Score**: 36

### [feat: Add file path protection to rulesets](https://github.com/integrations/terraform-provider-github/pull/2415)
---

## Statistics

### By Priority
- **P0 (Critical)**: 15 issues
- **P1 (High)**: 50 issues
- **P2 (Medium)**: 157 issues
- **P3 (Low)**: 160 issues
- **Total Open Issues**: 382

### By Type
- **Bugs**: 147
- **Features**: 114
- **Documentation**: 17

### Most Active Issues (by engagement)

- #509: 345 interactions
- #796: 154 interactions
- #567: 146 interactions
- #1956: 115 interactions
- #2583: 111 interactions

---

*Report generated: 2025-09-04 00:26:29 PDT*
*Updated: 2025-09-04 01:08:53 PDT - Marked completed issues*
