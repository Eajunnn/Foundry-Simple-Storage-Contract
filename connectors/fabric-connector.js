const { PluginLedgerConnectorFabric } = require("@hyperledger/cactus-plugin-ledger-connector-fabric");
const { RunTransactionRequest } = require("@hyperledger/cactus-core-api");

const pluginFabric = new PluginLedgerConnectorFabric({
  connectionProfile: require("../fabric-network/connection-org1.json"),
  instanceId: "fabric-connector",
  logLevel: "INFO",
  fabricWalletDir: "./wallet", // Path to your Fabric wallet
  discoveryOptions: { enabled: true, asLocalhost: true },
});

async function transactFabric(data) {
  try {
    const req = {
      channelName: "mychannel",
      chaincodeId: "mychaincode",
      fcn: "storeData",
      args: [data],
    };

    const res = await pluginFabric.transact(req);
    console.log("Fabric transaction result:", res);
    return res;
  } catch (ex) {
    console.error("Error interacting with Fabric:", ex);
    throw ex;
  }
}

module.exports = { transactFabric };
