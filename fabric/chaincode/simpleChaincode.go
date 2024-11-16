package main

import (
	"fmt"
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
	}
	return shim.Error("Invalid function name")
}

func (t *SimpleChaincode) storeData(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	data := args[0]

	// Here, you would invoke the middleware to send data to Ethereum
	// You can use an external call to your Node.js middleware or even directly interact with the Ethereum network
	err := invokeEthereum(data)
	if err != nil {
		return shim.Error("Failed to invoke Ethereum contract: " + err.Error())
	}

	return shim.Success(nil)
}

func invokeEthereum(data string) error {
	// This function would call your middleware (Node.js script) to send data to Ethereum
	// For simplicity, we're just printing the data here
	fmt.Printf("Sending data to Ethereum: %s\n", data)
	return nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting SimpleChaincode: %s", err)
	}
}
