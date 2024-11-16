const { PluginLedgerConnectorEthereum } = require("@hyperledger/cactus-plugin-ledger-connector-ethereum");
const { Contract } = require("ethers");
const SimpleStorageABI = require("../SimpleStorageABI.json");

const pluginEthereum = new PluginLedgerConnectorEthereum({
  rpcApiHttpHost: "https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID",  // Replace with your Ethereum RPC endpoint
  instanceId: "ethereum-connector",
  logLevel: "INFO",
});

const contractAddress = "YOUR_SIMPLE_STORAGE_CONTRACT_ADDRESS";

async function transactEthereum(value) {
  try {
    const req = {
      functionName: "store",
      functionArguments: [value],
      contractAddress: contractAddress,
      abi: SimpleStorageABI,
    };

    const res = await pluginEthereum.transact(req);
    console.log("Ethereum transaction result:", res);
    return res;
  } catch (ex) {
    console.error("Error interacting with Ethereum:", ex);
    throw ex;
  }
}

module.exports = { transactEthereum };
