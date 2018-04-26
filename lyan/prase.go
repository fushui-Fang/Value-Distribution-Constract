package lyan

import (
	"encoding/base64"
	"errors"
	"proto"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//===============================================================================
//主要将账本中的序列化过的数据进行解析操作
//
//===============================================================================
func praseAccount(key string, stub shim.ChaincodeStubInterface) (*Account, error) {

	if key == "" {
		return nil, errors.New("参数不正确")
	}

	userAccount := &Account{}
	userAccountProto, err := stub.GetState(key)
	if err != nil {
		return nil, errors.New("praseAccount:error hanppened when try to get accout ")
	}

	if err = proto.Unmarshal(userAccountProto, userAccount); err != nil {
		return nil, err
	}

	return userAccount, nil
}

func putAccount(key string, account *Account, stub shim.ChaincodeStubInterface) error {

	if key == "" || account == nil {
		return errors.New("putAccount:参数不正确")
	}

	accountProto, err := proto.Marshal(account)
	if err != nil {
		return errors.New("putAccount:error hanppened when try to marshal accout:" + err.Error())
	}

	err = stub.PutState(key, accountProto)
	if err != nil {
		return errors.New("putAccount:error hanppened when try to putState:" + err.Error())
	}

	return nil
}

//===============================================================================
//解析base64
//目的：为了避免将来因为编码处理不同的问题，在这里统一处理，方便修改
//===============================================================================
func praseBase64String(data string) ([]byte, error) {
	if data == "" {
		return nil, errors.New("[praseBase64String]传入参数为空,无法解析")
	}
	return base64.StdEncoding.DecodeString(data)

}
