const { transactEthereum } = require("../connectors/ethereum-connector");
const { transactFabric } = require("../connectors/fabric-connector");

async function syncData(payload) {
  try {
    // Send data to Fabric
    const fabricRes = await transactFabric(payload.value);
    console.log("Data stored on Fabric:", fabricRes);

    // Send the same data to Ethereum
    const ethereumRes = await transactEthereum(payload.value);
    console.log("Data stored on Ethereum:", ethereumRes);
  } catch (error) {
    console.error("Error in syncData:", error);
  }
}

// If running as a standalone script (e.g., via `node sync.js "someData"`)
if (require.main === module) {
  const data = process.argv[2];  // Get data from command-line argument
  syncData({ value: data }).catch((err) => {
    console.error("Error syncing data:", err);
  });
}

module.exports = { syncData };
