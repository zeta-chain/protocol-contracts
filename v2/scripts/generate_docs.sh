#!/bin/bash

forge doc

cat docs/src/README.md docs/src/SUMMARY.md > docs/src/index.md

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's|contracts/|protocol/contracts/|g' docs/src/index.md
else
    sed -i 's|contracts/|protocol/contracts/|g' docs/src/index.md
fi
