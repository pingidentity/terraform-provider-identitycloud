import re
providerSchemaPattern = re.compile(r'"(.*)": schema\.')
markdownSchemaPattern = re.compile(r'^- `(.*)` \(')

# Find all provider schema attributes
schemaAttrs = []
with open("internal/provider/provider.go") as providerGo:
  for line in providerGo:
    match = providerSchemaPattern.search(line)
    if match:
      schemaAttrs.append(match.group(1))

# Find all schema attributes included in the index template
schemaLabelFound = False
foundSchemaAttrs = []
with open("templates/index.md.tmpl") as indexTmpl:
  for line in indexTmpl:
    if line.strip() == "## Schema":
      schemaLabelFound = True
    if not schemaLabelFound:
      continue
    # Search for the matching lines in the Schema section
    match = markdownSchemaPattern.search(line)
    if match:
      foundSchemaAttrs.append(match.group(1))

# Ensure the two lists are the same
if set(schemaAttrs) != set(foundSchemaAttrs):
  print("Schema attributes in provider.go and index.md.tmpl do not match. Every schema attribute in provider.go must be documented in index.md.tmpl.")
  print("provider.go: ", schemaAttrs)
  print("index.md.tmpl: ", foundSchemaAttrs)
  exit(1)
