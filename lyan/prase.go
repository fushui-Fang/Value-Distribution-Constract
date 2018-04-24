package lyan

import (
	"encoding/base64"
	"errors"
	"proto"
)

//===============================================================================
//主要将账本中的序列化过的数据进行解析操作
//
//===============================================================================
func praseAccount(data []byte) (*Account, error) {

	if data == nil {
		return nil, errors.New("参数不正确")
	}

	userAccount := &Account{}
	if err := proto.Unmarshal(data, userAccount); err != nil {
		return nil, err
	}

	return userAccount, nil
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
