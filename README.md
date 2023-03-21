## Project to demonstrate microservice design (inter-service communication) secured by mTLS

Simple backend implementation including create user, login and authenticate user.

### System design

![image info](resources/system_design.svg)

### Tech specs

* **Go**
* **gRPC-go**: Core implementation of RPC for microservice design
* **Docker**: Build image and deploy application with docker compose 
* **mongo-go**: Driver of Mongodb for Golang
* **Redis**(on development): Cached revocational list of JWT tokens

### Feature

Create independent service that can deploy and scale with docker compose.

Secure internal service-service communication by mTLS.

