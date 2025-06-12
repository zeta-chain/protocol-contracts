#!/bin/bash

set -euo pipefail

forge doc

rm -rf index.md

cat docs/src/README.md docs/src/SUMMARY.md > docs/src/index.md

echo -e "---\ntitle: \"Protocol contracts\"\n---\n" | cat - docs/src/index.md > temp && mv temp docs/src/index.md

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's|contracts/|protocol/contracts/|g' docs/src/index.md
else
    sed -i 's|contracts/|protocol/contracts/|g' docs/src/index.md
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    find docs/src/ -type f -name "*.md" -exec sed -i '' -E 's|(https://github.com/zeta-chain/protocol-contracts/blob/)[^/]+/|\1main/|g' {} +
else
    find docs/src/ -type f -name "*.md" -exec sed -i -E 's|(https://github.com/zeta-chain/protocol-contracts/blob/)[^/]+/|\1main/|g' {} +
fi

# Create temporary files for prioritized and other docs
touch prioritized_docs.md other_docs.md

# First, process the two specific gateway files
for file in "docs/src/contracts/evm/GatewayEVM.sol/contract.GatewayEVM.md" "docs/src/contracts/zevm/GatewayZEVM.sol/contract.GatewayZEVM.md"; do
    if [ -f "$file" ]; then
        cat "$file" >> prioritized_docs.md
        rm "$file"
    fi
done

# Then process all other files
find docs/src/contracts -type f -name "*.md" -print0 | sort -z | while IFS= read -r -d '' file; do
    cat "$file" >> other_docs.md
done

cat prioritized_docs.md other_docs.md > index.md

rm prioritized_docs.md other_docs.md

rm -rf docs

mkdir -p docs

echo -e "# ZetaChain and EVM Protocol Contracts\n" > docs/index.md

npx markdown-toc index.md >> docs/index.md

cat index.md >> docs/index.md

rm index.md