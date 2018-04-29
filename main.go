package main

import (
	"chaincode/mycc/lyan"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {

	if err := shim.Start(&lyan.MyChaincode{}); err != nil {
		panic(err)
	}
}
