Analyze the schema package for:

1. **Empty/Unused Types**: Find all struct types with no fields or no usages
2. **Redundant Types**: Find structs with identical or nearly identical fields
3. **Similar Types**: Find structs with different names but overlapping attributes (>70% similar)
4. **Type Usage Map**: Show where each type is used

Output a markdown report with:
- Table of all types with field count and usage count
- Groups of similar/redundant types
- Recommended consolidations
