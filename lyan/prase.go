package lyan

import (
	"encoding/base64"
	"errors"

	"chaincode/mycc/proto"

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

func praseAsgy(id string, addr string, stub shim.ChaincodeStubInterface) (*AllocateSgy, error) {
	aSgyKey, err := stub.CreateCompositeKey(ASgyCompositeKeyIndexName, []string{id, addr})
	if err != nil {
		logger.Debug("praseAsgy: somthing wrong in CreateCompositeKey" + err.Error())
		return nil, err
	}

	//获取分配策略
	aSgyProto, err := stub.GetState(aSgyKey)
	aSgy := &AllocateSgy{}
	err = proto.Unmarshal(aSgyProto, aSgy)
	if err != nil {
		logger.Debug("praseAsgy: somthing wrong in proto.Unmarshal" + err.Error())
		return nil, err
	}

	return aSgy, nil
}

func putAsgy(aSgy *AllocateSgy, stub shim.ChaincodeStubInterface) error {
	aSgyKey, err := stub.CreateCompositeKey(ASgyCompositeKeyIndexName, []string{aSgy.ID, aSgy.Addr})
	if err != nil {
		logger.Debug("putAsgy: somthing wrong in CreateCompositeKey" + err.Error())
		return err
	}
	aSgyProto, _ := proto.Marshal(aSgy)

	err = stub.PutState(aSgyKey, aSgyProto)
	if err != nil {
		return err
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

//===============================================================================
//解析AsgyProposeProto
//目的：为了避免将来因为编码处理不同的问题，在这里统一处理，方便修改
//===============================================================================
func praseAsgyPropose(data []byte) (*ASgyPropose, error) {
	if data == nil {
		return nil, errors.New("[praseAsgyPropose]参数不正确")
	}
	x := &ASgyPropose{}
	err := proto.Unmarshal(data, x)
	if err != nil {
		return nil, err
	}
	return x, nil
}

func praseSigeForSgy(data []byte) (*SigeForSgy, error) {
	if data == nil {
		return nil, errors.New("[praseSigeForSgy]参数不正确")
	}
	x := &SigeForSgy{}
	err := proto.Unmarshal(data, x)
	if err != nil {
		return nil, err
	}
	return x, nil
}
