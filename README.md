# Redis-Sentinel Verification

This repository provides a Docker Compose setup for verifying Redis Sentinel with master/slave replication and a Golang pub/sub demo.

## Overview

1. Redis master/slave setup with Docker Compose (3 instances)
2. Redis Sentinel setup with Docker Compose (3 instances)
3. Golang pub/sub demo code

### 1.Run Redis Docker Compose
To start the Redis master/slave setup, use the following command:
```
 docker-compose up
```
 ![mole](https://github.com/bowwowxx/redis-sentinel/blob/main/redis-cluster.jpg)  


### 2.Run Redis Sentinel Docker Compose  
To start the Redis Sentinel setup, use the following command:  
**docker run** 
```
 docker-compose up
```
  ![mole](https://github.com/bowwowxx/redis-sentinel/blob/main/sentinel.jpg)  

### 3.Check All Containers and IPs
To list all running containers and their IP addresses, execute:  

 ```
 docker ps 
./redis.sh
 ```
  ![mole](https://github.com/bowwowxx/redis-sentinel/blob/main/container.jpg)  

### 4. Run Golang Publish and Check Sentinel Failover  
To run the Golang demo code that checks the Sentinel failover, follow these steps:  

 ```
 docker run -it -d --name goredis --network redis_default -v $(pwd):/app -w /app golang:latest
 docker exec -ti goredis bash
 go run main.go
 ```
  ![mole](https://github.com/bowwowxx/redis-sentinel/blob/main/Failover01.jpg) 
  ![mole](https://github.com/bowwowxx/redis-sentinel/blob/main/Failover02.jpg)  

### 5.Enjoy It!
Explore and enjoy the setup!