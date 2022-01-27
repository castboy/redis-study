#!/bin/bash

# 创建bridge网络
#docker network create testnet
# 创建master节点
docker run -it --name master --network testnet --network-alias master -v /Users/kite/PopMart/go/redis-study/master_slave:/etc/redis -d redis
# 创建slave节点1
#docker run -it --name slave1 --network testnet --network-alias centos-2 -v /Users/kite/PopMart/go/redis-study/cluster:/etc/redis -d redis /bin/bash redis-server /etc/redis/redis.slave1.conf
# 创建slave节点2
#docker run -it --name slave2 --network testnet --network-alias centos-3 -v /Users/kite/PopMart/go/redis-study/cluster:/etc/redis -d redis /bin/bash redis-server /etc/redis/redis.slave1.conf