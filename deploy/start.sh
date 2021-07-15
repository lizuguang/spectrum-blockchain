#!/bin/bash

# 根据需求保留，这里相当于使用fabric-samples_v1.4.7中的bin
if [[ `uname` == 'Darwin' ]]; then
    echo "Mac OS"
    export PATH=${PWD}/fabric/mac/bin:${PWD}:$PATH
fi
if [[ `uname` == 'Linux' ]]; then
    echo "Linux"
    export PATH=${PWD}/fabric/linux/bin:${PWD}:$PATH
fi

echo "1. 清理环境"
mkdir -p config
mkdir -p crypto-config
rm -rf config/*
rm -rf crypto-config/*
./stop.sh
echo "清理完毕"

echo "2. 生成证书和创世区块"
cryptogen generate --config=./crypto-config.yaml
configtxgen -profile OneOrgOrdererGenesis -outputBlock ./config/genesis.block

echo "区块链 ： 启动"
docker-compose up -d
echo "正在等待节点的启动完成，等待10秒"
sleep 10

echo "3. 生成channel配置区块"
configtxgen -profile TwoOrgChannel -outputCreateChannelTx ./config/assetschannel.tx -channelID assetschannel

echo "4. 创建channel"
# -o 为orderer_address
# -c 为channel名
# -f 为创建的channel地址
docker exec cli peer channel create -o orderer.spectrumblockchain.com:7050 -c assetschannel -f /etc/hyperledger/config/assetschannel.tx

echo "5. 节点加入channel"
docker exec cli peer channel join -b assetschannel.block

# -n 是链码的名字，可以自己随便设置
# -v 就是版本号，就是composer的bna版本
# -p 是目录，目录是基于cli这个docker里面的$GOPATH相对的
echo "6. 安装链代码"
docker exec cli peer chaincode install -n spectrum-blockchain -v 1.0.0 -l golang -p github.com/lizuguang/spectrum-blockchain/chaincode/spectrum-blockchain

#-n 对应前文安装链码的名字 其实就是composer network start bna名字
#-v 为版本号，相当于composer network start bna名字@版本号
#-C 是通道，在fabric的世界，一个通道就是一条不同的链，composer并没有很多提现这点，composer提现channel也就在于多组织时候的数据隔离和沟通使用
#-c 为传参，传入init参数
echo "7. 实例化链码"
if [[ "$(docker images -q hyperledger/fabric-ccenv:1.4.4 2> /dev/null)" == "" ]]; then
  docker pull hyperledger/fabric-ccenv:1.4.4
fi
if [[ "$(docker images -q hyperledger/fabric-ccenv:latest 2> /dev/null)" == "" ]]; then
  docker tag hyperledger/fabric-ccenv:1.4.4 hyperledger/fabric-ccenv:latest
fi
docker exec cli peer chaincode instantiate -o orderer.spectrumblockchain.com:7050 -C assetschannel -n spectrum-blockchain -l golang -v 1.0.0 -c '{"Args":["init"]}'

echo "正在等待链码实例化完成，等待5秒"
sleep 5

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
echo "8. 验证查询账户信息"
docker exec cli peer chaincode invoke -C assetschannel -n spectrum-blockchain -c '{"Args":["queryAccountList"]}'