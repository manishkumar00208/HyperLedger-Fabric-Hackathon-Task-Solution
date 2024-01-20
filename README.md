# HyperLedger-Fabric-Hackathon-Task-Solution


https://hyperledger-fabric.readthedocs.io/en/latest/whatis.html

Level 1

https://hyperledger-fabric.readthedocs.io/en/latest/getting_started.html
1.	Develop Hyperledger Fabric Chaincode for 
a.	Creating Asset
b.	Updating Asset values
c.	Query World State (Read Asset)
d.	Asset Transaction History
Asset structure
DEALERID
MSISDN
MPIN
BALANCE
STATUS
TRANSAMOUNT
TRANSTYPE
REMARKS

https://hyperledger-fabric.readthedocs.io/en/latest/smartcontract/smartcontract.html
Sample Chaincode 
https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-basic/chaincode-go
2.	Setup Hyperledger Fabric Test network
https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html
3.	Package and Deploy the chain code into fabric test network 
https://hyperledger-fabric.readthedocs.io/en/release-2.4/deploy_chaincode.html
4.	Test the chain code functionality using Fabric Peer CLI commands
Create Assset
Update Asset
Query World State(Read Asset)
Asset Transaction History
https://hyperledger-fabric.readthedocs.io/en/release-2.4/deploy_chaincode.html#invoking-the-chaincode


Level 2
1.	Develop REST API wrapper connecting to FABRIC GATEWAY API for invoking the above created chaincode methods
https://hyperledger-fabric.readthedocs.io/en/latest/gateway.html

GIN GOLANG Framework ( For REST API) Reference
		https://gin-gonic.com/
    
2.	Produce OPENAPI/Swagger based REST API Documentation
https://github.com/swaggo/gin-swagger
https://openapi-generator.tech/docs/generators/go-gin-server/
https://github.com/openapitools/openapi-generator

3.	Publish the REST API to Github
4.	Create a docker image for REST API
5.	Run Docker Image
6.	Test the rest API

