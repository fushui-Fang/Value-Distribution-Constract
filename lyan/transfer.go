package lyan

import (
	"crypto"
	"encoding/base64"
	"errors"
	"proto"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//==============================================================================================
// 价值转移
// 参数意义如下：
// 	1.价值转移的结构体
//==============================================================================================
func (s *chainCodeHub) transfer() pb.Response {

	txBase64String := s.args[0]
	txBase64, err := base64.StdEncoding.DecodeString(txBase64String)
	if err != nil {
		logger.Debug("[transfer()]" + err.Error())
		return shim.Error("[transfer()]error happens when try to decode base64String")
	}

	//获取交易结构体
	txOrgin := &TX{}
	err = proto.Unmarshal(txBase64, txOrgin)
	if err != nil {
		logger.Debug("[transfer()]" + err.Error())
		return shim.Error("[transfer()]error happens when try to decode base64String")
	}

	//获取输入的proto信息
	addr := txOrgin.Tx.InputAddr
	accountProto, err := s.stub.GetState(addr)

	//获取账户
	account := &Account{}
	err = proto.Unmarshal(accountProto, account)
	if err != nil {
		logger.Debug("[transfer()]" + err.Error())
		return shim.Error("[transfer()]error happens when try to Unmarshal account")
	}

	//获取签名以及交易的proto二进制形式
	txProto, err := proto.Marshal(txOrgin.GetTx())
	if err != nil {
		logger.Debug("[transfer()]" + err.Error())
		return shim.Error("[transfer()]error happens when try to Marshal Tx")
	}
	signText := []byte(txOrgin.Script)

	//验证签名
	err = VerifySign(account.GetPubkey(), txProto, signText, crypto.SHA1)
	if err != nil {
		logger.Debug("[transfer()]" + err.Error())
		return shim.Error("[transfer()]error happens when try to Marshal Tx")
	}
	//logger.Info(" [transfer()] success to verify the signed message")

	//验证时间戳
	if time.Now().UTC().Unix()-txOrgin.Tx.GetTimestamp() > 120 {
		logger.Error(ErrTimeOut.Error())
		return shim.Error(ErrTimeOut.Error())
	}

	//验证nonce
	if txOrgin.Tx.InputBalance != txOrgin.Tx.Nounce {
		logger.Error(errors.New("[transfer()]nounce doesn't match the imputbalance").Error())
		return shim.Error("[transfer()]nounce doesn't match the imputbalance")
	}

	//验证输入与输出是否对等
	var outbalance float32
	outbalance = 0
	for _, out := range txOrgin.Tx.GetOutput() {
		outbalance += out.OutNum
	}

	if outbalance != txOrgin.Tx.GetInputBalance() {
		logger.Error(errors.New("[transfer()]outbalance doesn't match the input balance").Error())
		return shim.Error("[transfer()]outbalance doesn't match the input balance")
	}

	//验证是否有这么多钱
	if account.GetBalance() < outbalance {
		logger.Error(errors.New("[transfer()]没那么多钱").Error())
		return shim.Error("[transfer()]没那么多钱")
	}

	//接下来是转账
	//如果是普通用户则正常转账，合约用户则随意转
	for _, out := range txOrgin.Tx.GetOutput() {
		outAddr := out.GetOutAddr()
		outAccountProto, err := s.stub.GetState(outAddr)

		outAccount := &Account{}
		err = proto.Unmarshal(outAccountProto, outAccount)
		if err != nil {
			logger.Debug("[transfer()]" + err.Error())
			return shim.Error("[transfer()]error happens when try to Unmarshal account")
		}

		//合约用户转账地址
		if outAccount.GetKind() != NormolUser {
			err = transfer2ContractAccount(account, outAccount, out.OutNum, s.stub)
			if err != nil {
				return shim.Error(err.Error())
			}
			//logger.Debug("[transfer()] 合约账户的代码还没还是进行，先空着")
			continue
		}

		account.Balance -= out.OutNum
		outAccount.Balance += out.OutNum

		outAccountProto, _ = proto.Marshal(outAccount)
		s.stub.PutState(outAddr, outAccountProto)

		accountProto, _ := proto.Marshal(account)
		s.stub.PutState(addr, accountProto)

	}

	//将本次交易信息记录下来
	TxKey, err := s.stub.CreateCompositeKey(TxCompositeKeyIndexName, []string{s.stub.GetTxID(), addr})
	if err != nil {
		return shim.Error(err.Error())
	}

	err = s.stub.PutState(TxKey, txBase64)
	if err != nil {
		return shim.Error(err.Error())
	}

	logger.Debug("transfer ID is:" + s.stub.GetTxID())
	//将转账id返回作为参考
	return shim.Success([]byte(s.stub.GetTxID()))
}

func transfer2ContractAccount(in *Account, out *Account, num float32, stub shim.ChaincodeStubInterface) error {
	//在输入账号减去这么多的金额，

	//获取分配策略
	aSgy, err := praseAsgy(out.ID, out.Addr, stub)
	if err != nil {
		return err
	}
	//遍历给out中的地址按比例增加收入
	for _, outSingleMemberInfo := range aSgy.ASgy.Ms {
		//取出账户
		outAccount, err := praseAccount(outSingleMemberInfo.Addr, stub)
		if err != nil {
			logger.Debug(err.Error())
			return err
		}

		//修改余额
		outAccount.Balance += num * outSingleMemberInfo.Ort

		//存入账户
		err = putAccount(outAccount.Addr, outAccount, stub)
		if err != nil {
			return err
		}

	}
	//logger.Debug("come here!")
	in, _ = praseAccount(in.Addr, stub)

	in.Balance -= num
	inProto, _ := proto.Marshal(in)
	stub.PutState(in.Addr, inProto)

	return nil
}
