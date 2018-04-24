package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/mycc/lyan"
)

func main() {
	
	if err := shim.Start(&lyan.MyChaincode{}); err != nil {
		panic(err)
	}
}