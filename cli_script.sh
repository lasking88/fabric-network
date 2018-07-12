export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export CHANNEL_NAME=mychannel
# register participant
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>0233c7b3-aa91-4897-b381-6a0e23c1e5ed</id><action>register</action><metadataType>participant</metadataType><payload><participant><key>www.example.com/participant</key><active>true</active><description>This is a participant</description></participant></payload></dataMessage>"]}'
# register
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>3bb50a0c-c811-46de-a1f5-462cd5364ae6</id><action>register</action><metadataType>connector</metadataType><payload><connector><key>www.example.com/connector1</key><active>true</active><description>This is connector1</description><idsParticipant>www.example.com/participant</idsParticipant></connector></payload></dataMessage>"]}'
# register 3
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>03b02375-b78f-4903-b66b-33d943cbe7ae</id><action>register</action><metadataType>connector</metadataType><payload><connector><key>www.example.com/connector2</key><active>true</active><description>This is connector2</description><idsParticipant>www.example.com/participant</idsParticipant></connector></payload></dataMessage>"]}'
# register 4
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>d8fe44b8-8b64-4d41-9842-c86d50564a7e</id><action>register</action><metadataType>dataendpoint</metadataType><payload><dataendpoint><key>www.example.com/dataendpoint1</key><active>true</active><description>This is dataendpoint1</description><idsConnector>www.example.com/connector1</idsConnector></dataendpoint></payload></dataMessage>"]}'
# register 5
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>53c37e87-e5ac-4c48-b956-4489fe2baaa5</id><action>register</action><metadataType>dataendpoint</metadataType><payload><dataendpoint><key>www.example.com/dataendpoint2</key><active>true</active><description>This is dataendpoint2</description><idsConnector>www.example.com/connector2</idsConnector></dataendpoint></payload></dataMessage>"]}'
# update
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>0233c7b3-aa91-4897-b381-6a0e23c1e5ed</id><action>update</action><metadataType>participant</metadataType><payload><participant><key>www.example.com/participant</key><active>true</active><description>This is a modified participant</description><IdsConnectors></IdsConnectors></participant></payload></dataMessage>"]}'
# activate
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>c70d2d19-1486-4d0b-a4ec-48bdcf9fdb71</id><action>activate</action><metadataType>participant</metadataType><payload>www.example.com/participant</payload></dataMessage>"]}'
# passivate
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>f5fb0cb5-37c9-4fb9-9bde-62c8b58b7ee3</id><action>passivate</action><metadataType>participant</metadataType><payload>www.example.com/participant</payload></dataMessage>"]}'
# remove
peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["dataRequest","<dataMessage><id>0233c7b3-aa91-4897-b381-6a0e23c1e5ed</id><action>remove</action><metadataType>participant</metadataType><payload>www.example.com/participant</payload></dataMessage>"]}'
# query by key
peer chaincode query -C mychannel -n mycc -v 1.0 -c '{"Args":["queryRequest","<queryMessage><id>tmp</id><action>get</action><scope>active</scope><payload><dataendpoint><key>www.example.com/participant</key></dataendpoint></payload></queryMessage>"]}'
# query by type
peer chaincode query -C mychannel -n mycc -v 1.0 -c '{"Args":["queryRequest","<queryMessage><id>tmp</id><action>list</action><scope>active</scope><payload>connector</payload></queryMessage>"]}'
