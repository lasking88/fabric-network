#/bin/bash

rm -r ./crypto-config
rm -r ./channel-artifacts
mkdir ./channel-artifacts

export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=mychannel

./bin/cryptogen generate --config=./crypto-config.yaml
./bin/configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
./bin/configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
./bin/configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org2MSP

docker-compose -f docker-compose-cli.yaml up -d
