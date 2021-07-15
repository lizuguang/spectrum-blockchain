#!/bin/bash


echo "1. 清理docker中images, ccontainers, volumes"

docker system prune --volumes

docker image prune # 移除未使用的docker
systemctl stop docker
rm -rf /var/lib/docker
systemctl start docker
echo "清理docker完毕"
