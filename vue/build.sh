#!/bin/bash

echo "一、清理环境、删除旧容器"
rm -rf dist
docker rm -f spectrum-blockchain-vue

echo "二、开始打包编译"
docker build -t lizuguang/spectrum-blockchain-vue:v1 .

echo "三、运行编译容器"
docker run -it -d --name spectrum-blockchain-vue lizuguang/spectrum-blockchain-vue:v1

echo "四、拷贝容器中编译后的dist资源并放到application目录下"
docker cp spectrum-blockchain-vue:/root/vue/dist ./../application/dist
