import { keccak256, toUtf8Bytes } from "ethers";
import * as fs from "fs";
import * as path from "path";

// Get CLI arg
const [, , abiPath] = process.argv;

if (!abiPath) {
  console.error("Usage: ts-node index.ts <path-to-abi.json>");
  process.exit(1);
}

const resolvedPath = path.resolve(abiPath);

let abiJson: any;
let abi: any[];

try {
  abiJson = JSON.parse(fs.readFileSync(resolvedPath, "utf-8"));

  // Use .abi if available (Foundry/Hardhat), otherwise assume file *is* the ABI
  abi = Array.isArray(abiJson) ? abiJson : abiJson.abi;

  if (!Array.isArray(abi)) {
    throw new Error("ABI is not an array");
  }
} catch (err) {
  console.error(`❌ Failed to load ABI from ${resolvedPath}:`, err);
  process.exit(1);
}

// Build function signature
function getFunctionSignature(item: any): string {
  const inputs = item.inputs.map((i: any) => i.type).join(",");
  return `${item.name}(${inputs})`;
}

// Compute selector
function getSelector(signature: string): Buffer {
  const hash = keccak256(toUtf8Bytes(signature));
  return Buffer.from(hash.slice(2, 10), "hex");
}

// Compute interface ID
let interfaceId = 0;
const selectors: { sig: string; selector: string }[] = [];

for (const item of abi) {
  if (item.type !== "function") continue;

  const signature = getFunctionSignature(item);
  const selectorBuf = getSelector(signature);
  const selectorInt = selectorBuf.readUInt32BE(0);

  selectors.push({ sig: signature, selector: `0x${selectorBuf.toString("hex")}` });
  interfaceId ^= selectorInt;
}

console.log(`0x${interfaceId.toString(16).padStart(8, "0")}`);
