package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type SimpleChaincode struct{}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "storeData" {
		return t.storeData(stub, args)
	} else if function == "addPerson" {
		return t.addPerson(stub, args)
	} else if function == "getPrivateData" {
		return t.getPrivateData(stub, args)
	}
	return shim.Error("Invalid function name")
}

// Store data on Ethereum via middleware
func (t *SimpleChaincode) storeData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	data := args[0]

	// Save private data in Fabric
	err := stub.PutPrivateData("collectionPrivateData", "dataKey", []byte(data))
	if err != nil {
		return shim.Error("Failed to store private data: " + err.Error())
	}

	// Send data to Ethereum
	err = invokeEthereum("store", []interface{}{data})
	if err != nil {
		return shim.Error("Failed to invoke Ethereum contract: " + err.Error())
	}

	return shim.Success(nil)
}

// Add a person to Ethereum via middleware
func (t *SimpleChaincode) addPerson(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: name, favoriteNumber")
	}
	name := args[0]
	favoriteNumber := args[1]

	// Save private data in Fabric
	privateData := map[string]string{
		"name":           name,
		"favoriteNumber": favoriteNumber,
	}
	privateDataJSON, _ := json.Marshal(privateData)
	err := stub.PutPrivateData("collectionPrivateData", name, privateDataJSON)
	if err != nil {
		return shim.Error("Failed to store private data: " + err.Error())
	}

	// Send data to Ethereum
	err = invokeEthereum("addPerson", []interface{}{name, favoriteNumber})
	if err != nil {
		return shim.Error("Failed to invoke Ethereum contract: " + err.Error())
	}

	return shim.Success(nil)
}

// Retrieve private data from Fabric
func (t *SimpleChaincode) getPrivateData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: name")
	}
	name := args[0]

	data, err := stub.GetPrivateData("collectionPrivateData", name)
	if err != nil {
		return shim.Error("Failed to get private data: " + err.Error())
	}
	return shim.Success(data)
}

// Middleware to interact with Ethereum
func invokeEthereum(functionName string, args []interface{}) error {
	apiURL := "http://localhost:3000/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-ethereum/invokeContract"
	reqBody := map[string]interface{}{
		"functionName":     functionName,
		"functionArguments": args,
		"contractAddress":  "0xEF0d6002DaF7CA2163A9ed2399AefaEcf6fC22Dd",
		"abi":              getSimpleStorageABI(),
	}
	jsonData, _ := json.Marshal(reqBody)

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to invoke Ethereum contract, status: %d", resp.StatusCode)
	}
	return nil
}

// ABI of the Ethereum contract
func getSimpleStorageABI() string {
	return `[{"inputs":[{"internalType":"uint256","name":"_favoriteNumber","type":"uint256"}],"name":"store","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"_name","type":"string"},{"internalType":"uint256","name":"_favoriteNumber","type":"uint256"}],"name":"addPerson","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting SimpleChaincode: %s", err)
	}
}
