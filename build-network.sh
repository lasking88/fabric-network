#/bin/bash

rm -r ./crypto-config
rm -r ./channel-artifacts
mkdir ./channel-artifacts

export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=tnochannel

./bin/cryptogen generate --config=./crypto-config.yaml
./bin/configtxgen -profile OrdererGenesis -outputBlock ./channel-artifacts/genesis.block
./bin/configtxgen -profile MyChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
./bin/configtxgen -profile MyChannel -outputAnchorPeersUpdate ./channel-artifacts/Comp1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Comp1MSP
./bin/configtxgen -profile MyChannel -outputAnchorPeersUpdate ./channel-artifacts/Comp2MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Comp2MSP
