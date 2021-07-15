#!/bin/bash

priv_sk_path=$(ls ../crypto-config/peerOrganizations/org1.spectrumblockchain.com/users/Admin\@org1.spectrumblockchain.com/msp/keystore/)

cp -rf ./connection-profile/network_temp.json ./connection-profile/network.json

sed -i "s/priv_sk/$priv_sk_path/" ./connection-profile/network.json

docker-compose down -v
docker-compose up -d
echo "区块链浏览器，请访问http://localhost:8080，用户名admin，密码 123456"