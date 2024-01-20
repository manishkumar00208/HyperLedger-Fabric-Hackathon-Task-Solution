package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/swagger"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type AssetContract struct {
	contractapi.Contract
}

type Asset struct {
	DealerID      string `json:"dealerID"`
	MSISDN        string `json:"msisdn"`
	MPIN          string `json:"mpin"`
	Balance       int    `json:"balance"`
	Status        string `json:"status"`
	TransAmount   int    `json:"transAmount"`
	TransType     string `json:"transType"`
	Remarks       string `json:"remarks"`
}

func (ac *AssetContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerID, msisdn, mpin string, balance int, status, transType, remarks string) error {
	asset := Asset{
		DealerID:      dealerID,
		MSISDN:        msisdn,
		MPIN:          mpin,
		Balance:       balance,
		Status:        status,
		TransAmount:   0,
		TransType:     transType,
		Remarks:       remarks,
	}

	return ctx.GetStub().PutState(msisdn, asset.ToJSON())
}

func (ac *AssetContract) UpdateAsset(ctx contractapi.TransactionContextInterface, msisdn string, newBalance int, newStatus, newTransType, newRemarks string) error {
	asset, err := ac.ReadAsset(ctx, msisdn)
	if err != nil {
		return err
	}

	asset.Balance = newBalance
	asset.Status = newStatus
	asset.TransType = newTransType
	asset.Remarks = newRemarks

	return ctx.GetStub().PutState(msisdn, asset.ToJSON())
}

func (ac *AssetContract) ReadAsset(ctx contractapi.TransactionContextInterface, msisdn string) (*Asset, error) {
	assetBytes, err := ctx.GetStub().GetState(msisdn)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	if assetBytes == nil {
		return nil, fmt.Errorf("asset not found: %s", msisdn)
	}

	var asset Asset
	err = asset.FromJSON(assetBytes)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling asset JSON: %v", err)
	}

	return &asset, nil
}

func main() {
	// Set up Fabric Gateway
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile("connection-profile.yaml")),
		gateway.WithIdentity("user", "mspID"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to Fabric gateway: %s", err.Error())
		return
	}
	defer gw.Close()

	network, err := gw.GetNetwork("channel1")
	if err != nil {
		fmt.Printf("Failed to get network: %s", err.Error())
		return
	}

	assetContract := new(AssetContract)

	// Set up GIN Golang framework
	router := gin.Default()

	// Load TLS CA certificates
	tlsCACertOrg1, err := ioutil.ReadFile("certificates/org1/tlsca.org1.example.com-cert.pem")
	if err != nil {
		fmt.Printf("Failed to load TLS CA certificate for Org1: %s", err.Error())
		return
	}

	tlsCACertOrderer, err := ioutil.ReadFile("certificates/orderer/tlsca.example.com-cert.pem")
	if err != nil {
		fmt.Printf("Failed to load TLS CA certificate for Orderer: %s", err.Error())
		return
	}

	// Define REST API routes
	router.POST("/createAsset", func(c *gin.Context) {
		// Extract parameters from the request and invoke the chaincode method
		// Example: assetContract.CreateAsset(...)
		c.JSON(http.StatusOK, gin.H{"message": "Asset created successfully"})
	})

	router.PUT("/updateAsset", func(c *gin.Context) {
		// Similar to createAsset, handle the updateAsset route
		c.JSON(http.StatusOK, gin.H{"message": "Asset updated successfully"})
	})

	router.GET("/readAsset/:msisdn", func(c *gin.Context) {
		// Extract MSISDN from the request and invoke the chaincode method to read asset
		// Example: assetContract.ReadAsset(...)
		c.JSON(http.StatusOK, gin.H{"message": "Asset read successfully"})
	})


	// Produce OpenAPI/Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the GIN server
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Failed to start GIN server: %s", err.Error())
	}
}
