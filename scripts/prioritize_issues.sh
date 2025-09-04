#!/bin/bash

# GitHub Issue Prioritizer for terraform-provider-github
# This script fetches issues from the original repository and generates a prioritized list
# based on community engagement, issue type, and other factors

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
REPO="integrations/terraform-provider-github"
OUTPUT_DIR="issue_analysis"
OUTPUT_FILE="PRIORITY_ISSUES.md"  # Main output in base directory
JSON_FILE="${OUTPUT_DIR}/issues_data.json"

echo -e "${BLUE}GitHub Issue Prioritizer${NC}"
echo "=========================="
echo ""

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo -e "${RED}Error: GitHub CLI (gh) is not installed${NC}"
    echo "Please install it from: https://cli.github.com/"
    exit 1
fi

# Check if jq is installed
if ! command -v jq &> /dev/null; then
    echo -e "${RED}Error: jq is not installed${NC}"
    echo "Please install it: brew install jq"
    exit 1
fi

echo -e "${YELLOW}Fetching issues from ${REPO}...${NC}"

# Fetch all open issues (up to 500)
gh api "repos/${REPO}/issues?state=open&per_page=100" > "${JSON_FILE}.1"
gh api "repos/${REPO}/issues?state=open&per_page=100&page=2" > "${JSON_FILE}.2"
gh api "repos/${REPO}/issues?state=open&per_page=100&page=3" > "${JSON_FILE}.3"
gh api "repos/${REPO}/issues?state=open&per_page=100&page=4" > "${JSON_FILE}.4"
gh api "repos/${REPO}/issues?state=open&per_page=100&page=5" > "${JSON_FILE}.5"

# Combine all pages
jq -s 'add' "${JSON_FILE}".* > "$JSON_FILE"
rm -f "${JSON_FILE}".*

# Count total issues
TOTAL_ISSUES=$(jq 'length' "$JSON_FILE")
echo -e "${GREEN}Fetched ${TOTAL_ISSUES} open issues${NC}"

# Process and score issues
echo -e "${YELLOW}Analyzing and scoring issues...${NC}"

# Create the prioritized markdown file
cat > "$OUTPUT_FILE" << 'EOF'
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

EOF

# Process issues and calculate scores
jq '
  def calculate_score:
    # Base engagement score (40% of total)
    ((.reactions.total_count * 2 + .comments) * 0.4) as $engagement |
    
    # Type score (20% of total)
    (if (.labels | map(.name) | any(test("bug|Bug"; "i"))) then 20
     elif (.labels | map(.name) | any(test("feat|feature|Feature"; "i"))) then 10
     else 5 end) as $type_score |
    
    # Label priority score (20% of total)
    (if (.labels | map(.name) | any(test("Status: Pinned"))) then 10 else 0 end +
     if (.labels | map(.name) | any(test("auth|Authentication"))) then 5 else 0 end +
     if (.labels | map(.name) | any(test("security|Security"))) then 5 else 0 end +
     if (.labels | map(.name) | any(test("Status: Blocked"))) then -5 else 0 end) as $label_score |
    
    # Age factor (20% of total) - older issues with engagement get bonus
    (if (now - (.created_at | fromdateiso8601)) > (365 * 24 * 60 * 60) and (.reactions.total_count + .comments) > 20 
     then 20 
     elif (now - (.created_at | fromdateiso8601)) > (180 * 24 * 60 * 60) and (.reactions.total_count + .comments) > 10
     then 10
     else 0 end) as $age_score |
    
    ($engagement + $type_score + $label_score + $age_score);

  # Sort by calculated score
  sort_by(calculate_score) | reverse | 
  
  # Map each issue to formatted output
  [.[] |
  calculate_score as $score |
  
  # Determine priority level
  (if $score >= 70 then "P0"
   elif $score >= 40 then "P1"
   elif $score >= 20 then "P2"
   else "P3" end) as $priority |
   
  # Extract issue type
  (if (.labels | map(.name) | any(test("bug|Bug"; "i"))) then "Bug"
   elif (.labels | map(.name) | any(test("feat|feature|Feature"; "i"))) then "Feature"
   elif (.labels | map(.name) | any(test("doc|Doc"; "i"))) then "Documentation"
   else "Other" end) as $issue_type |
   
  # Format the output
  {
    priority: $priority,
    score: $score,
    number: .number,
    title: .title,
    type: $issue_type,
    reactions: .reactions.total_count,
    comments: .comments,
    labels: (.labels | map(.name) | join(", ")),
    created: .created_at,
    url: .html_url
  }]
' "$JSON_FILE" > "${OUTPUT_DIR}/scored_issues.json"

# Group issues by priority and write to markdown
echo "## P0 - Critical Priority Issues" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

P0_COUNT=$(jq '[.[] | select(.priority == "P0")] | length' "${OUTPUT_DIR}/scored_issues.json")
if [ "$P0_COUNT" -eq "0" ]; then
    echo "*No P0 issues identified*" >> "$OUTPUT_FILE"
else
    jq -r '.[] | select(.priority == "P0") | 
      "### \(.score | floor). [\(.title)](\(.url))\n" +
      "- **Issue #**: \(.number)\n" +
      "- **Type**: \(.type)\n" +
      "- **Score**: \(.score | floor) (👍 \(.reactions) reactions, 💬 \(.comments) comments)\n" +
      "- **Labels**: \(.labels)\n" +
      "- **Created**: \(.created | split("T")[0])\n"
    ' "${OUTPUT_DIR}/scored_issues.json" >> "$OUTPUT_FILE"
fi

echo "---" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"
echo "## P1 - High Priority Issues" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

jq -r '.[] | select(.priority == "P1") | 
  "### \(.score | floor). [\(.title)](\(.url))\n" +
  "- **Issue #**: \(.number)\n" +
  "- **Type**: \(.type)\n" +
  "- **Score**: \(.score | floor) (👍 \(.reactions) reactions, 💬 \(.comments) comments)\n" +
  "- **Labels**: \(.labels)\n"
' "${OUTPUT_DIR}/scored_issues.json" | head -60 >> "$OUTPUT_FILE"

echo "---" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"
echo "## P2 - Medium Priority Issues" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

jq -r '.[] | select(.priority == "P2") | 
  "### [\(.title)](\(.url))\n" +
  "- **Issue #**: \(.number) | **Type**: \(.type) | **Score**: \(.score | floor)\n"
' "${OUTPUT_DIR}/scored_issues.json" | head -40 >> "$OUTPUT_FILE"

# Generate statistics
echo "---" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"
echo "## Statistics" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

P1_COUNT=$(jq '[.[] | select(.priority == "P1")] | length' "${OUTPUT_DIR}/scored_issues.json")
P2_COUNT=$(jq '[.[] | select(.priority == "P2")] | length' "${OUTPUT_DIR}/scored_issues.json")
P3_COUNT=$(jq '[.[] | select(.priority == "P3")] | length' "${OUTPUT_DIR}/scored_issues.json")

BUG_COUNT=$(jq '[.[] | select(.type == "Bug")] | length' "${OUTPUT_DIR}/scored_issues.json")
FEATURE_COUNT=$(jq '[.[] | select(.type == "Feature")] | length' "${OUTPUT_DIR}/scored_issues.json")
DOC_COUNT=$(jq '[.[] | select(.type == "Documentation")] | length' "${OUTPUT_DIR}/scored_issues.json")

echo "### By Priority" >> "$OUTPUT_FILE"
echo "- **P0 (Critical)**: $P0_COUNT issues" >> "$OUTPUT_FILE"
echo "- **P1 (High)**: $P1_COUNT issues" >> "$OUTPUT_FILE"
echo "- **P2 (Medium)**: $P2_COUNT issues" >> "$OUTPUT_FILE"
echo "- **P3 (Low)**: $P3_COUNT issues" >> "$OUTPUT_FILE"
echo "- **Total Open Issues**: $TOTAL_ISSUES" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

echo "### By Type" >> "$OUTPUT_FILE"
echo "- **Bugs**: $BUG_COUNT" >> "$OUTPUT_FILE"
echo "- **Features**: $FEATURE_COUNT" >> "$OUTPUT_FILE"
echo "- **Documentation**: $DOC_COUNT" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

# Top contributors (most active issue reporters)
echo "### Most Active Issues (by engagement)" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"
jq -r 'sort_by(.reactions + .comments) | reverse | .[0:5] | 
  .[] | "- #\(.number): \(.reactions + .comments) interactions"
' "${OUTPUT_DIR}/scored_issues.json" >> "$OUTPUT_FILE"

echo "" >> "$OUTPUT_FILE"
echo "---" >> "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"
echo "*Report generated: $(date '+%Y-%m-%d %H:%M:%S %Z')*" >> "$OUTPUT_FILE"

# Print summary
echo ""
echo -e "${GREEN}✅ Priority analysis complete!${NC}"
echo ""
echo "Summary:"
echo "--------"
echo "P0 (Critical): $P0_COUNT issues"
echo "P1 (High): $P1_COUNT issues"
echo "P2 (Medium): $P2_COUNT issues"
echo "P3 (Low): $P3_COUNT issues"
echo ""
echo -e "📊 Full report saved to: ${BLUE}${OUTPUT_FILE}${NC}"
echo -e "📁 Raw data saved to: ${BLUE}${JSON_FILE}${NC} (gitignored)"