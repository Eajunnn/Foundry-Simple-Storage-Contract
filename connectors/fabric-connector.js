const axios = require("axios");

const API_SERVER_URL = "http://localhost:3000/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-fabric";

async function transactFabric(data) {
  try {
    const req = {
      channelName: "mychannel",
      chaincodeId: "mychaincode",
      fcn: "storeData", // Fabric chaincode function
      args: [data],
    };

    const res = await axios.post(`${API_SERVER_URL}/runTransaction`, req);
    console.log("Fabric transaction result:", res.data);
    return res.data;
  } catch (error) {
    console.error("Error interacting with Fabric via API:", error.message);
    throw error;
  }
}

module.exports = { transactFabric };
