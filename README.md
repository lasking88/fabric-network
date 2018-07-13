# Hyperledger Fabric Broker Service Network

This repository is to build hyperledger fabric network based on [Build Your First Network](http://hyperledger-fabric.readthedocs.io/en/latest/build_network.html) tutorial.

## Requirement
Please follow prerequisites of Hyperledger [Fabric v1.1-released](https://hyperledger-fabric.readthedocs.io/en/release-1.1/build_network.html#install-prerequisites).

## Description
This Fabric network is for a simulation of broker service in [International Data Space (IDS)](https://www.fraunhofer.de/content/dam/zv/en/fields-of-research/industrial-data-space/whitepaper-industrial-data-space-eng.pdf) using blockchain technology. Broker service in IDS is a registry of metadata, that is publication or retrieval of matadata. The functions can be executed by invoking a transaction included in a message as an argument. There are two types of messages :
1. data message
2. query message

A data message is for a publication of metadata which contains an action, a metadata type, and a payload. On the other hand, a query meassage is for a retrieval of metadata which contains an action, a scope and a payload.

An action for a data message can be
1. Register
2. Update
3. Remove
4. Activate
5. Passivate

An action for a query message can be
1. Key
2. List
3. Query - not yet supported

A scope for a query message can be
1. All
2. Active
3. Access - not yet supported

A metadata type can be
1. Participant
2. Connector
3. Dataendpoint
4. DataApp - not yet supported

These types have a hierarchy such that a participatn can have many connectors each of which can have many data endpoints.

The examples of transactions are attached on a file named ```cli_script.sh```

## Execution
```

```
