package lyan

import (
	"crypto"
	"strconv"

	proto "github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//======================================================================================
// get the 4 args: org1Pubkey  org1's balance init num
//				   org2Pubkey  org2's balance init num
//	将公钥从文件中读取出来并进行hash存储，设定相关的初始账号值
//  hash算法采用sha1
//=====================================================================================
func (t *MyChaincode) initEnv(stub shim.ChaincodeStubInterface, args []string) error {
	if len(args) != 4 {
		logger.Error(ErrInitArgNum.Error())
		return ErrInitArgNum
	}

	org1Pubkey := args[0]
	org1InitNum, err := strconv.ParseInt(args[1], 10, 64)
	org1Addr := hashData([]byte(org1Pubkey), crypto.SHA1)
	//由公钥采用SHA1hash得到地址，注意，是全部公钥pem文件的hash

	org1Account := &Account{
		Addr:    org1Addr,
		Kind:    NormolUser,
		Balance: float32(org1InitNum),
		Pubkey:  org1Pubkey,
		ID:      "0",
	}

	x, err := proto.Marshal(org1Account)
	if err != nil {
		logger.Debugf("can't  marshal the org1Account")
		return ErrInit
	}
	stub.PutState(org1Addr, x)

	//设置ORG1的结构体,并存入ledger中

	org2Pubkey := args[2]
	org2InitNum, err := strconv.ParseInt(args[3], 10, 64)
	org2Addr := hashData([]byte(org2Pubkey), crypto.SHA1)

	org2Account := &Account{
		Addr:    org2Addr,
		Kind:    NormolUser,
		Balance: float32(org2InitNum),
		Pubkey:  org2Pubkey,
		ID:      "0",
	}

	x, err = proto.Marshal(org2Account)
	if err != nil {
		logger.Debugf("can't  marshal the org1Account")
		return ErrInit
	}
	stub.PutState(org2Addr, x)

	return nil
}
