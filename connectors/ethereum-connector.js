import abi from "../SimpleStorageABI.json";


const axios = require("axios");

const API_SERVER_URL = "http://localhost:3000/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-ethereum";

async function transactEthereum(value) {
  try {
    const req = {
      functionName: "store", // Ethereum contract method
      functionArguments: [value],
      contractAddress: "0xEF0d6002DaF7CA2163A9ed2399AefaEcf6fC22Dd",
      abi: abi,
    };    

    const res = await axios.post(`${API_SERVER_URL}/invokeContract`, req);
    console.log("Ethereum transaction result:", res.data);
    return res.data;
  } catch (error) {
    console.error("Error interacting with Ethereum via API:", error.message);
    throw error;
  }
}

module.exports = { transactEthereum };
