package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"strings"
	"encoding/xml"
	"bytes"
)

type MetaDataChainCode struct {
}

<<<<<<< HEAD
// definition of data request message
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
type dataMessage struct {
	Id			 	string		`xml:"id"`
	Action		 	string		`xml:"action"`
	MetadataType 	string		`xml:"metadataType"`
	Payload 		struct {
		Element		string 		`xml:",innerxml"`
	}	`xml:"payload"`
}

<<<<<<< HEAD

// definition of query request message
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
type queryMessage struct {
	Id 				string		`xml:"id"`
	Action 			string		`xml:"action"`
	Payload			struct {
		Element		string		`xml:",innerxml"`
	}	`xml:"payload"`
	Scope 			string		`xml:"scope"`
}

type IdsElementInterface interface {
	GetKey()	string
	Activate()
	Passivate()
}

<<<<<<< HEAD
// common features as an IDS element
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
type IdsElement struct {
	Key			string		`xml:"key"`
	Active		string		`xml:"active"`
	Description string		`xml:"description"`
}

func (e *IdsElement) GetKey() string {
	return e.Key
}

func (e *IdsElement) Activate() {
	e.Active = "true"
}

func (e *IdsElement) Passivate() {
	e.Active = "false"
}

<<<<<<< HEAD
// Participant type which has many Connector entities
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
type participant struct {
	IdsElement
	IdsConnectors		[]string	`xml:">idsConnector"`
}

<<<<<<< HEAD
// Connector type which has a participant as a parent entity and many data endpoints entities
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
type connector struct {
	IdsElement
	Participant			string		`xml:"idsParticipant"`
	Dataendpoints		[]string	`xml:">idsDataEndpoint"`
}

<<<<<<< HEAD
// Data-endpoint type which has a parent Connector
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
type dataendpoint struct {
	IdsElement
	Connector			string		`xml:"idsConnector"`
}

func main() {
	if err := shim.Start(new(MetaDataChainCode)); err != nil {
		fmt.Printf("Error starting chaincode %s", err)
	}
}

func (t *MetaDataChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

<<<<<<< HEAD
// Invoke method implementation to invoke a transaction or a query
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
func (t *MetaDataChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running...")

	if fn == "dataRequest" {
		return t.dataRequest(stub, args)
	} else if fn == "queryRequest" {
		return t.queryRequest(stub, args)
	}

	fmt.Println("invoke did not find func: " + fn) //error
	return shim.Error("Received unknown function invocation")
}

<<<<<<< HEAD
// Auxilary function to publish a metadata
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
func (t *MetaDataChainCode) dataRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Expecting a dataRequestMessage")
	}

	fmt.Println("processing data request message..")
	if len(args[0]) <= 0 {
		err := fmt.Sprintf("Message must be non-empty string")
		return shim.Error(err)
	}

	var message dataMessage
	err := xml.Unmarshal([]byte(args[0]), &message)
	if err != nil {
		return shim.Error("message unmarshal failed")
	}

	id := message.Id
	if len(id) <= 0 {
		err := fmt.Sprintf("Id must be non-empty string")
		return shim.Error(err)
	}

	action := strings.ToLower(message.Action)
	metadataType := strings.ToLower(message.MetadataType)
	payload := message.Payload.Element

	switch action {
	case "register", "update":
		return t.updateElement(stub, action, metadataType, payload)
	case "activate", "passivate":
		return t.activateElement(stub, action, metadataType, payload)
	case "remove":
		return t.removeElement(stub, metadataType, payload)
	default:
		return shim.Error("No action exist: " + action)
	}

	return shim.Error("Unknown action")
}

<<<<<<< HEAD
// Auxiliary function to retrieve a metadata or a metadata set
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
func (t *MetaDataChainCode) queryRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Expecting a queryRequestMessage")
	}

	fmt.Println("processing query request message..")
	if len(args[0]) <= 0 {
		err := fmt.Sprintf("Message must be non-empty string")
		return shim.Error(err)
	}

	var message queryMessage
	err := xml.Unmarshal([]byte(args[0]), &message)
	if err != nil {
		return shim.Error("message unmarshal failed")
	}

	id := message.Id
	if len(id) <= 0 {
		err := fmt.Sprintf("Id must be non-empty string")
		return shim.Error(err)
	}

	action := strings.ToLower(message.Action)
	scope := strings.ToLower(message.Scope)
	if scope != "all" && scope != "active" && scope != "access" {
		err := fmt.Sprintf("There is no matching scope")
		return shim.Error(err)
	}

	payloadBytes := []byte(message.Payload.Element)
	switch action {
	case "get": return t.fetchElementByKey(stub, action, payloadBytes, scope)
	case "list": return t.fetchElementsByType(stub, action, payloadBytes, scope)
	//case "query": // not yet supported
	default:
		return shim.Error("No action exist: " + action)
	}

	return shim.Success(nil)
}

func (t *MetaDataChainCode) updateElement(stub shim.ChaincodeStubInterface, action string, metadataType string, payload string) peer.Response {
	var err error
	var element IdsElement
	var par participant
	var con connector
	var dat dataendpoint
	switch metadataType {
	case "participant":
		err = xml.Unmarshal([]byte(payload), &par)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		element = par.IdsElement
	case "connector":
		err = xml.Unmarshal([]byte(payload), &con)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		element = con.IdsElement
		parBytes, err := stub.GetState(con.Participant)
		if err != nil {
			return shim.Error("Get state failed")
		} else if parBytes == nil {
			return shim.Error("Associated participant do not exist")
		}
		err = xml.Unmarshal(parBytes, &par)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		existed := false
		for _, element := range par.IdsConnectors {
			if element == con.GetKey() {
				existed = true
				break
			}
		}
		if !existed {
			par.IdsConnectors = append(par.IdsConnectors, con.GetKey())
		}
		parBytes, err = xml.Marshal(par)
		if err != nil {
			return shim.Error("Marshal failed")
		}
		response := t.updateElement(stub, "update", "participant", string(parBytes))
		if response.Status != shim.OK {
			return shim.Error("Update participant failed")
		}
	case "dataendpoint":
		err = xml.Unmarshal([]byte(payload), &dat)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		element = dat.IdsElement
		conBytes, err := stub.GetState(dat.Connector)
		if err != nil {
			return shim.Error("Get state failed")
		} else if conBytes == nil {
			return shim.Error("Associated connector do not exist")
		}
		err = xml.Unmarshal(conBytes, &con)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		existed := false
		for _, element := range con.Dataendpoints {
			if element == dat.GetKey() {
				existed = true
				break
			}
		}
		if !existed {
			con.Dataendpoints = append(con.Dataendpoints, dat.GetKey())
		}

		conBytes, err = xml.Marshal(con)
		if err != nil {
			return shim.Error("Marshal failed")
		}
		response := t.updateElement(stub, "update", "connector", string(conBytes))
		if response.Status != shim.OK {
			return shim.Error("Update connector failed")
		}
	default:
		err := fmt.Sprintf("There is no matching metadata type")
		return shim.Error(err)
	}
	payloadKey := element.GetKey()
	if len(payloadKey) <= 0 {
		return shim.Error("Key must be non-empty string")
	}

	elementBytes, err := stub.GetState(payloadKey)
	if err != nil {
		return shim.Error("Failed to get element" + err.Error())
	} else if action == "register" && elementBytes != nil {
		return shim.Error("Key already exists : " + payloadKey)
	} else if action == "update" && elementBytes == nil {
		return shim.Error("Element does not exist : " + payloadKey)
	}

	err = stub.PutState(payloadKey, []byte(payload))
	if err != nil {
		return shim.Error(err.Error())
	}

	indexName := "key~type"
	keyActiveIndexKey, err := stub.CreateCompositeKey(indexName, []string{metadataType, element.GetKey()})
	if err != nil {
		return shim.Error(err.Error())
	}
	value := []byte{0x00}
	stub.PutState(keyActiveIndexKey, value)

	return shim.Success(nil)
}

func (t *MetaDataChainCode) activateElement(stub shim.ChaincodeStubInterface, action string, metadataType string, payload string) peer.Response {
	elementBytes, err := stub.GetState(payload)
	if err != nil {
		return shim.Error("Failed to get element" + err.Error())
	} else if elementBytes == nil {
		return shim.Error("Element does not exist : " + payload)
	}

	var par *participant
	var con *connector
	var dat *dataendpoint
	var element *IdsElement
	switch metadataType {
	case "participant":
		err = xml.Unmarshal(elementBytes, &par)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		element = &par.IdsElement
		activation := activateHelper(action, &par.IdsElement)
		if activation.Status != shim.OK {
			return activation
		}
		valueBytes, err := xml.Marshal(par)
		if err != nil {
			return shim.Error("Marshal failed")
		}
		err = stub.PutState(element.GetKey(), valueBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	case "connector":
		err = xml.Unmarshal(elementBytes, &con)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		element = &con.IdsElement
		activation := activateHelper(action, &con.IdsElement)
		if activation.Status != shim.OK {
			return activation
		}
		valueBytes, err := xml.Marshal(con)
		if err != nil {
			return shim.Error("Marshal failed")
		}
		err = stub.PutState(element.GetKey(), valueBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	case "dataendpoint":
		err = xml.Unmarshal(elementBytes, &dat)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		element = &dat.IdsElement
		activation := activateHelper(action, &dat.IdsElement)
		if activation.Status != shim.OK {
			return activation
		}
		valueBytes, err := xml.Marshal(dat)
		if err != nil {
			return shim.Error("Marshal failed")
		}
		err = stub.PutState(element.GetKey(), valueBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	default:
		err := fmt.Sprintf("There is no matching metadata type")
		return shim.Error(err)
	}

	return shim.Success(nil)
}

func activateHelper(action string, element *IdsElement) peer.Response {
	if action == "activate" {
		element.Activate()
		return shim.Success(nil)
	} else if action == "passivate" {
		element.Passivate()
		return shim.Success(nil)
	} else {
		return shim.Error("No action exist")
	}
}

func (t *MetaDataChainCode) removeElement(stub shim.ChaincodeStubInterface, metadataType string, payload string) peer.Response {
	elementBytes, err := stub.GetState(payload)
	if err != nil {
		return shim.Error("Get state failed")
	} else if elementBytes == nil {
		return shim.Error("Element does not exist: " + payload)
	}
	switch metadataType {
	case "participant" :
		var par participant
		err = xml.Unmarshal(elementBytes, &par)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		children := par.IdsConnectors
		for _, child := range children {
			response := t.removeElement(stub, "connector", child)
			if response.Status != shim.OK {
				return shim.Error("Failed remove one of the children")
			}
		}
	case "connector" :
		var con connector
		err = xml.Unmarshal(elementBytes, &con)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
		children := con.Dataendpoints
		for _, child := range children {
			response := t.removeElement(stub, "dataendpoint", child)
			if response.Status != shim.OK {
				return shim.Error("Failed remove one of the children")
			}
		}
	case "dataendpoint" :
		var dat dataendpoint
		err = xml.Unmarshal(elementBytes, &dat)
		if err != nil {
			return shim.Error("Unmarshal failed")
		}
	default:
		return shim.Error("No metadataType exist: " + metadataType)
	}

	err = stub.DelState(payload)
	if err != nil {
		return shim.Error("Failed to delete element " + err.Error())
	}

	indexName := "key~type"
	compositeKey, err := stub.CreateCompositeKey(indexName, []string{metadataType, payload})
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.DelState(compositeKey)
	if err != nil {
		return shim.Error("Failed to delete state: " + err.Error())
	}

	return shim.Success(nil)
}

<<<<<<< HEAD
// fetch an IDS element filtered by a key
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
func (t *MetaDataChainCode) fetchElementByKey(stub shim.ChaincodeStubInterface, action string, payloadBytes []byte, scope string) peer.Response {
	var element IdsElement
	err := xml.Unmarshal(payloadBytes, &element)
	if err != nil {
		return shim.Error("Unmarshal failed")
	}

	elementBytes, err := stub.GetState(element.GetKey())
	if err != nil {
		return shim.Error("Get state is failed: " + element.GetKey())
	} else if elementBytes == nil {
		return shim.Error("No element exist: " + element.GetKey())
	}
	return scopeHelper(stub, elementBytes, scope)
}

<<<<<<< HEAD
// fetch an IDS element filtered by a type
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
func (t *MetaDataChainCode) fetchElementsByType(stub shim.ChaincodeStubInterface, action string, payloadBytes []byte, scope string) peer.Response {
	metadataType := string(payloadBytes)

	if metadataType != "dataendpoint" && metadataType != "connector" && metadataType != "participant" {
		return shim.Error("No metadata exist : " + metadataType)
	}

	typedElementIterator, err := stub.GetStateByPartialCompositeKey("key~type", []string{metadataType})
	if err != nil {
		return shim.Error(err.Error())
	}
	defer typedElementIterator.Close()

	var i int
	var buffer bytes.Buffer
	buffer.WriteString("<elements>")
	for i = 0; typedElementIterator.HasNext(); i++ {
		responseRange, err := typedElementIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedType := compositeKeyParts[0]
		returnedKey := compositeKeyParts[1]
		fmt.Printf("- found a data from index:%s type:%s key:%s\n", objectType, returnedType, returnedKey)

		elementBytes, err := stub.GetState(returnedKey)
		if err != nil {
			return shim.Error("Get state error")
		} else if elementBytes == nil {
			return shim.Error("Element does not exist : " + returnedKey)
		}
		response := scopeHelper(stub, elementBytes, scope)
		if response.Status == shim.OK {
			buffer.WriteString("<value>")
			buffer.Write(elementBytes)
			buffer.WriteString("</value>")
		}
	}

	buffer.WriteString("</elements>")
	return shim.Success(buffer.Bytes())
}

<<<<<<< HEAD
// Auxiliary function to deal with scope of a message
=======
>>>>>>> cff2ef9b487cfb9d1267040269dde7cb93d62c8d
func scopeHelper(stub shim.ChaincodeStubInterface, elementBytes []byte, scope string) peer.Response {
	var element IdsElement
	err := xml.Unmarshal(elementBytes, &element)
	if err != nil {
		return shim.Error("Unmarshal error")
	}
	switch scope {
	case "all": return shim.Success(elementBytes)
	case "active":
		if element.Active == "true" {
			return shim.Success(elementBytes)
		} else {
			return shim.Error("The element is not active : " + element.GetKey())
		}
	case "access":
		// not yet supported.
	}
	return shim.Error("Unexpecte error")
}