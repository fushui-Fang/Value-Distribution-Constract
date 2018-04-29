package lyan

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"chaincode/mycc/proto"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//***************************************************************************************
//	用于实现与账本操作逻辑关联性不大的重复性操作
//	*getPubString(string)[]byte 	从某个公钥pem文件中读取文件内容
//
//***************************************************************************************

//======================================================================================
//从指定文件中获取公私钥文件
//
//=====================================================================================
func getKeyString(filepath string) ([]byte, error) {
	pubfile, filerr := os.Open(filepath)
	if filerr != nil {
		fmt.Print(filerr)
		return nil, filerr
	}
	defer pubfile.Close()

	//读取pem文件
	data := make([]byte, 1024)
	count, err := pubfile.Read(data)
	if err != nil {
		return nil, err
	}
	Org1pubKey := data[:count]

	//解析pem文件
	block, _ := pem.Decode(Org1pubKey)
	if block == nil {
		return nil, errors.New("decode key from pem file get some error")
	}

	return block.Bytes, nil
}

//=====================================================================================
//
// 将in的数据根据制定hash算法来处理，返回string类型
//
//=====================================================================================
func hashData(in []byte, hash crypto.Hash) string {
	h := hash.New()
	h.Write(in)
	hashed := h.Sum(nil)
	result := base64.StdEncoding.EncodeToString(hashed)
	return result
}

func VerifySign(publicKeyString string, src []byte, sign []byte, hash crypto.Hash) error {
	pubInterface, err := x509.ParsePKIXPublicKey([]byte(publicKeyString))
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}

//签名
func signMessage(message []byte, privateKey []byte, hash crypto.Hash) ([]byte, error) {
	priv, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	h := hash.New()
	h.Write(message)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, priv, hash, hashed)
}

func VerifyHashdata(in []byte, hash crypto.Hash, res string) error {
	h := hash.New()
	h.Write(in)
	hashed := h.Sum(nil)
	result := base64.StdEncoding.EncodeToString(hashed)

	if res != result {
		return errors.New("VerifyHashdata:hash 校验没有通过")
	}

	return nil
}

//===============================================================================================
// 验证分配策略的签名情况
//
//===============================================================================================
func InitVerifySgySigned(stub shim.ChaincodeStubInterface, m *AllocateSgy, signStructBase64String string) error {
	//解码得到签名信息列表
	signStructProto, err := praseBase64String(signStructBase64String)
	if err != nil {
		return err
	}

	signStruct := &SigeForSgyStub{}
	err = proto.Unmarshal(signStructProto, signStruct)

	//验证签名数是否一致
	if len(m.ASgy.Ms) != len(signStruct.Set) {
		return errors.New("[InitVerifySgySigned]:签名个数不对")
	}

	//开始验证签名,
	for _, singleSigned := range signStruct.Set {
		err = VerifySgySingleSigned(stub, m, singleSigned)
		if err != nil {
			return err
		}
	}
	return nil
}

//===============================================================================================
//
//===============================================================================================
func VerifySgySingleSigned(stub shim.ChaincodeStubInterface, m *AllocateSgy, singleSigned *SigeForSgy) error {
	seq := singleSigned.Si.Seq
	//logger.Debug("序号是")
	//logger.Debug(seq)

	//基本信息确认
	if singleSigned.Si.Ort != m.ASgy.Ms[seq].Ort ||
		seq != m.ASgy.Ms[seq].Sequence ||
		singleSigned.Si.Addr != m.ASgy.Ms[seq].Addr ||
		singleSigned.Si.ID != m.GetID() {
		logger.Debug("签名信息中:")
		logger.Debug(singleSigned.Si.Ort)
		logger.Debug(singleSigned.Si.Addr)
		logger.Debug(singleSigned.Si.ID)

		logger.Debug("分配策略中:")
		logger.Debug(m.ASgy.Ms[seq].Ort)
		logger.Debug(m.ASgy.Ms[seq].Addr)
		logger.Debug(m.ID)

		return errors.New("[VerifySgySigned]:基本信息没有对上")
	}

	//取出账户公钥，验证签名
	account := &Account{}
	accountProto, err := stub.GetState(singleSigned.Si.Addr)
	if err != nil {
		return err
	}

	proto.Unmarshal(accountProto, account)
	publicKey := account.Pubkey

	Siproto, _ := proto.Marshal(singleSigned.Si)

	err = VerifySign(publicKey, Siproto, []byte(singleSigned.Script), crypto.SHA1)
	if err != nil {
		return err
	}

	//验证时间戳
	if time.Now().UTC().Unix()-singleSigned.Si.GetTimestamp() > 3600 {
		logger.Error(ErrTimeOut.Error())
		return errors.New("VerifySgySigned:签名过时")
	}

	//验证正确，在此设置相应的标志位,并将签名信息写进区块链中
	m.HasSigned[seq] = true
	compositeKye, err := stub.CreateCompositeKey(UserSiganedindexName, []string{singleSigned.Si.Addr, string(singleSigned.Si.ID)})
	if err != nil {
		return errors.New("VerifySgySigned:create compositeKye fail!")
	}
	singleSignedProto, _ := proto.Marshal(singleSigned)

	stub.PutState(compositeKye, singleSignedProto)
	return nil
}
