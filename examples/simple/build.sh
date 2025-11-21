#!/bin/bash

# Build simple watchface example

../../watchface-builder \
  --name "Simple Clock Example" \
  --template simple \
  --version 1.0.0 \
  --author "ziztechnology" \
  --description "A minimalist digital clock with gradient background" \
  --tags "simple,minimalist,gradient" \
  --output ./

echo ""
echo "Example built successfully!"
echo "Check the generated ZIP file in this directory."
