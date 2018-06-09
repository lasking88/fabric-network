package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"strings"
	"encoding/xml"
)

type MetaDataChainCode struct {
}

type metadata struct {
	Key			string
	Type 		string
	Description string
}

func main() {
	if err := shim.Start(new(MetaDataChainCode)); err != nil {
		fmt.Printf("Error starting chaincode %s", err)
	}
}

func (t *MetaDataChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 0 {
		return shim.Error("Expecting no parameter")
	}

	return shim.Success(nil)
}

func (t *MetaDataChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running...")

	if fn == "publish" {
		return t.publish(stub, args)
	} else if fn == "retrieve" {
		return t.retrieve(stub, args)
	}

	fmt.Println("invoke did not find func: " + fn) //error
	return shim.Error("Received unknown function invocation")
}

func (t *MetaDataChainCode) publish(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments expecting 3")
	}
	fmt.Println("start publishing metadata..")
	for i := 0; i < 3; i++ {
		if len(args[i]) <= 0 {
			err := fmt.Sprintf("Arguments must be non-empty strings (%d)", i+1)
			shim.Error(err)
		}
	}
	key := strings.ToLower(args[0])
	mtype := strings.ToLower(args[1])
	description := args[2]

	// Check if metadata already exists
	data, err := stub.GetState(key)
	if err != nil {
		shim.Error("Failed to get metadata: "+ err.Error())
	} else if data != nil {
		shim.Error("This metadata already exists : "+ key)
	}

	metadata := &metadata{key, mtype, description}
	metadatabytes, err := xml.Marshal(metadata)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(key, metadatabytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *MetaDataChainCode) retrieve(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting the key of the metadata to retrieve")
	}
	fmt.Println("start retrieving data ")

	key := args[0]
	data, err := stub.GetState(key)
	if err != nil {
		shim.Error("Failed to get metadata: "+ err.Error())
	} else if data != nil {
		shim.Error("This metadata already exists : "+ key)
	}

	return shim.Success(data)
}
