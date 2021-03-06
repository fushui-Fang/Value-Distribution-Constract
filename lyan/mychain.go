package lyan

import (
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	"chaincode/mycc/proto"
)

type chainCodeHub struct {
	stub shim.ChaincodeStubInterface
	args []string
}

type store interface {
	createContractAcount() pb.Response
	createAccount() pb.Response
	queryAccount() pb.Response
	transfer() pb.Response
	querySgyByID() pb.Response
	queryCurrentSgyID() pb.Response
	modifyAllocateByUser() pb.Response
	signForAsgy() pb.Response
}

//============================================================================================
//  将stub 和 args 封装进结构体里面，
//	减少传值的繁琐代码
//=============================================================================================
func makeChainCodeHub(stub shim.ChaincodeStubInterface, args []string) store {
	t := &chainCodeHub{}
	t.stub = stub
	t.args = args
	return t
}

//==============================================================================================
// 创建合约账户
// 参数意义如下：
// 	1. 账户地址 2.策略初始集合 3合约签名集合
// 备注：为了避免创建合约账户之后，被攻击者中途指定新的合约账户，因此在最开始的时候，传入参与者签名
//		即，这个合约账户的创建是得到所有参与者的同意的
//==============================================================================================

func (s *chainCodeHub) createContractAcount() pb.Response {
	addr := s.args[0]
	res, _ := s.stub.GetState(addr)

	fmt.Printf("\nconaddr  : %v\n", addr)

	//账户是否已经存在以及账户地址是否是按照SHA1标准得来的
	if res != nil {
		logger.Error("[createContractAcount]:欲创建的合约账号已存在")
		return shim.Error("[createContractAcount]:欲创建的合约账号已存在")
	}
	if len(addr) != 28 {
		logger.Error("[createContractAcount]:欲创建的合约账户地址长度不对")
		logger.Error("长度为" + addr)
		return shim.Error("[createContractAcount]:欲创建的合约账户地址长度不对")

	}

	//获取当前txID
	aSgyID := "0"
	//logger.Debug("[createContractAcount]:TxID is " + aSgyID)

	//为了方便后面的签名检查，在这里把合约的初始策略设置为0
	account := &Account{
		Addr:    addr,
		Kind:    ContractUser,
		Balance: 0,
		Pubkey:  "0",
		ID:      "0",
	}

	//接收传来的策略
	aSgy := &AllocateSgy{}

	//传来的数据先解码
	aSgyProto, err := praseBase64String(s.args[1])
	if err != nil {
		logger.Error("[createContractAcount]:base64 解码错误:" + err.Error())
		return shim.Error("[createContractAcount]:base64 解码错误" + err.Error())
	}

	err = proto.Unmarshal(aSgyProto, aSgy)
	if err != nil {
		logger.Error("[createContractAcount]:proto 解码错误:" + err.Error())
		return shim.Error("[createContractAcount]:proto 解码错误:" + err.Error())
	}

	//验证价值分配策略的总和是否为1
	var sum float32
	for _, r := range aSgy.ASgy.Ms {
		sum += r.Ort
	}
	if sum != 1 {
		logger.Error("[createContractAcount]:不合法的分配策略，总和不为1")
		return shim.Error("[createContractAcount]:不合法的分配策略，总和不为1")
	}

	/*
		//验证合约账户的完整性
		sgyInfo, _ := proto.Marshal(aSgy.ASgy)
		err = VerifyHashdata(sgyInfo, crypto.SHA1, aSgy.Hash)
		if err != nil {
			logger.Error("[createContractAcount]:" + err.Error())
			return shim.Error("[createContractAcount]:" + err.Error())
		}
	*/

	//验证账户签名
	err = InitVerifySgySigned(s.stub, aSgy, s.args[2])
	if err != nil {
		logger.Error("[createContractAcount]:" + err.Error())
		return shim.Error("[createContractAcount]:" + err.Error())
	}

	//建立合约账户
	err = putAccount(addr, account, s.stub)
	if err != nil {
		logger.Error("[createContractAcount]:" + err.Error())
		return shim.Error("[createContractAcount]:" + err.Error())
	}

	//保存分配策略
	aSgyKey, err := s.stub.CreateCompositeKey(ASgyCompositeKeyIndexName, []string{aSgyID, addr})
	if err != nil {
		return shim.Error(err.Error())
	}
	aSgy.ID = "0"
	aSgy.PriorID = "0"

	aSgyProto, _ = proto.Marshal(aSgy)
	err = s.stub.PutState(aSgyKey, aSgyProto)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("createContractAcount: successful to create the contract account "))
}

//==============================================================================================
// 创建nomal 账户
// 参数意义如下：
// 	1.公钥
//  默认一般用户的hash为0
//==============================================================================================

func (s *chainCodeHub) createAccount() pb.Response {
	if len(s.args) != 1 {
		logger.Error("[createAccount]:创建一般账户的参数数量不对")
		return shim.Error("[createAccount]:创建一般账户的参数数量不对")
	}
	nUserPubKey := s.args[0]
	nUserAddr := hashData([]byte(nUserPubKey), crypto.SHA1)

	nUserAccount := &Account{
		Addr:    nUserAddr,
		Kind:    NormolUser,
		Balance: 0,
		Pubkey:  nUserPubKey,
		ID:      "0",
	}

	x, err := proto.Marshal(nUserAccount)
	if err != nil {
		logger.Debugf("[createAccount]:can't  marshal the nUserAccount")
		return shim.Error("[createAccount]:can't  marshal the nUserAccount")
	}
	s.stub.PutState(nUserAddr, x)

	return shim.Success([]byte("successful to create the normol Account!"))
}

//==============================================================================================
// 根据地址查询创建账户
//	只有账户地址一个参数
//==============================================================================================
func (s *chainCodeHub) queryAccount() pb.Response {

	if len(s.args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	qAcountAddr := s.args[0]
	data, err := s.stub.GetState(qAcountAddr)
	if err != nil {
		return shim.Error(errors.New("获取账户信息失败").Error())
	}

	//将账户信息转换为base64编码形式
	qAcount := &Account{}
	queryBase64String := base64.StdEncoding.EncodeToString(data)
	proto.Unmarshal(data, qAcount)

	logger.Debug(" [queryAcount base64 Encode ] addr " + queryBase64String)
	logger.Debug(" [queryAcount get by data ] addr [%v]", qAcount.GetAddr())

	return shim.Success([]byte(queryBase64String))
}

//==============================================================================================
// 根据地址和ID查询分配策略
//	参数  1 id 2 addr
//	返回参与结果的
//==============================================================================================
func (s *chainCodeHub) querySgyByID() pb.Response {
	//获取 分配策略的KEY
	id := s.args[0]
	addr := s.args[1]

	aSgyKey, err := s.stub.CreateCompositeKey(ASgyCompositeKeyIndexName, []string{id, addr})
	if err != nil {
		return shim.Error(err.Error())
	}

	aSgyProto, err := s.stub.GetState(aSgyKey)
	if err != nil {
		return shim.Error(err.Error())
	}

	aSgyBase64 := base64.StdEncoding.EncodeToString(aSgyProto)

	return shim.Success([]byte(aSgyBase64))
}

//==============================================================================================
// 根据地址查询分配策略最新ID
//	参数  1  addr
//	返回策略ID
//==============================================================================================
func (s *chainCodeHub) queryCurrentSgyID() pb.Response {

	account, err := praseAccount(s.args[0], s.stub)
	if err != nil {
		logger.Error("[queryCurrentSgyID]:" + err.Error())
		return shim.Error("[queryCurrentSgyID]:" + err.Error())
	}

	return shim.Success([]byte(account.ID))
}

//==============================================================================================
// 提供新的未签名的策略集
//	参数 1 提出的策略的结构体
//	返回处理结果
//==============================================================================================

func (s *chainCodeHub) modifyAllocateByUser() pb.Response {
	//数据解码
	base64, err := praseBase64String(s.args[0])
	if err != nil {
		logger.Debug("[modifyAllocateByUser]" + err.Error())
		return shim.Error("[modifyAllocateByUser]error happens when try to decode base64String")
	}

	asgyp, err := praseAsgyPropose(base64)
	if err != nil {
		logger.Debug("[modifyAllocateByUser]" + err.Error())
		return shim.Error("[modifyAllocateByUser]error happens when try to decode proto")
	}

	//获取用户账号信息
	account, err := praseAccount(asgyp.Asgip.UserAddr, s.stub)
	if err != nil {
		logger.Debug("[modifyAllocateByUser]" + err.Error())
		return shim.Error("[modifyAllocateByUser]error happens when try to decode proto")
	}

	//是否在最新的策略上更改
	conAccount, err := praseAccount(asgyp.Asgip.ContractAddr, s.stub)
	if err != nil {
		logger.Debug("[modifyAllocateByUser]" + err.Error())
		return shim.Error("[modifyAllocateByUser]error happens when try to decode proto")
	}
	if conAccount.ID != asgyp.Asgip.PriorID {
		fmt.Print(conAccount.ID)
		fmt.Print(asgyp.Asgip.PriorID)
		logger.Debug("[modifyAllocateByUser]:没有在最新策略上更改")
		return shim.Error("[modifyAllocateByUser]没有在最新策略上更改")
	}

	//验证账号签名
	asgyInfoProto, err := proto.Marshal(asgyp.Asgip)
	if err != nil {
		logger.Debug("[modifyAllocateByUser]" + err.Error())
		return shim.Error("[modifyAllocateByUser]error happens when try to marshal proto")
	}

	err = VerifySign(account.Pubkey, asgyInfoProto, []byte(asgyp.Script), crypto.SHA1)
	if err != nil {
		logger.Debug("[modifyAllocateByUser]" + err.Error())
		return shim.Error("[modifyAllocateByUser]策略签名未通过")
	}

	//校验是否存在这个用户在所添加的策略集里
	asgy, err := praseAsgy(asgyp.Asgip.PriorID, asgyp.Asgip.ContractAddr, s.stub)
	if err != nil {
		logger.Debug("[modifyAllocateByUser]" + err.Error())
		return shim.Error("[modifyAllocateByUser]获取前置策略失败")
	}

	//验证成员是否更改以及ort总和是否为1
	hashSigned := []bool{}
	var sum float32
	for _, r := range asgy.ASgy.Ms {
		sum += r.Ort
		seq := r.Sequence
		mem := asgyp.Asgip.Mi.Ms[seq]
		if mem.Addr != r.Addr {
			logger.Error("[createContractAcount]:成员地址与原来的不符")
			return shim.Error("[createContractAcount]:成员地址与原来的不符")
		}
		hashSigned = append(hashSigned, false)
	}
	if sum != 1 {
		logger.Error("[createContractAcount]:不合法的分配策略，总和不为1")
		return shim.Error("[createContractAcount]:不合法的分配策略，总和不为1")
	}

	//验证成员是否有资格创建相关的策略
	seq := asgyp.Asgip.Seq
	if asgy.ASgy.Ms[seq].Addr != asgyp.Asgip.UserAddr {
		logger.Error("[modifyAllocateByUser]:签名账户不在对应的序列表中或者序列号错误")
		return shim.Error("[modifyAllocateByUser]签名账户不在对应的序列表中或者序列号错误")
	}

	//验证时间戳是否过期
	if time.Now().UTC().Unix()-asgyp.Asgip.Timestamp > 120 {
		logger.Error("[modifyAllocateByUser]:时间戳超时")
		return shim.Error("[modifyAllocateByUser]时间戳超时")
	}

	//接下来要完成的部分：讲这个提议存下来，创建新的未签名的分配策略

	id := s.stub.GetTxID()
	asgy = &AllocateSgy{
		ID:        id,
		PriorID:   asgyp.Asgip.PriorID,
		Addr:      asgyp.Asgip.ContractAddr,
		ASgy:      asgyp.Asgip.Mi,
		HasSigned: hashSigned,
	}

	err = putAsgy(asgy, s.stub)
	if err != nil {
		logger.Error("[modifyAllocateByUser]:" + err.Error())
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(id))
}

//==============================================================================================
//	为新编写的策略签名确认
//	参数 1	签名集合的复合体 2合约地址　３合约当前ID
//	返回处理结果
//==============================================================================================

func (s *chainCodeHub) signForAsgy() pb.Response {
	//取出数据
	sFAGBase64, err := praseBase64String(s.args[0])
	if err != nil {
		logger.Debug("[signForAsgy]" + err.Error())
		return shim.Error("[signForAsgy]error happens when try to decode base64String")
	}

	sFAG, err := praseSigeForSgy(sFAGBase64)
	if err != nil {
		logger.Debug("[signForAsgy]" + err.Error())
		return shim.Error("[signForAsgy]error happens when try to decode base64String")
	}

	//取出要改变的合约的信息
	ID := s.args[2]
	Addr := s.args[1]

	//取出策略集
	aSgy, err := praseAsgy(ID, Addr, s.stub)
	if err != nil {
		logger.Debug("[signForAsgy]" + err.Error())
		return shim.Error("[signForAsgy]error happens when try to decode base64String")
	}
	logger.Debug("signForAsgy小伙子在哪呢--去策略有问题")
	logger.Debug(aSgy.Addr)

	err = VerifySgySingleSigned(s.stub, aSgy, sFAG)
	if err != nil {
		logger.Debug("[signForAsgy]" + err.Error())
		return shim.Error("[signForAsgy]error happens when try to decode base64String")
	}

	logger.Debug("[signForAsgy] come here!\n\n\n")
	isChange := true

	//aSgy.HasSigned[sFAG.Si.Seq] = true

	for _, r := range aSgy.HasSigned {

		isChange = isChange && r
	}

	//将修改后的分配策略存回
	putAsgy(aSgy, s.stub)

	//如果所有的用户都已经签字,则修改用户策略
	if isChange == true {
		conAccount, err := praseAccount(aSgy.Addr, s.stub)
		if err != nil {
			logger.Debug("[signForAsgy]" + err.Error())
			return shim.Error("[signForAsgy]error happens when try to decode base64String")
		}

		//是否是修改的最新的用户策略
		if conAccount.ID != aSgy.PriorID {
			logger.Debug("[signForAsgy]提交的策略不是最新的修改策略")
			return shim.Success([]byte("Success to submit teh user's sign message "))
		}

		//修改合约用户的最新合约
		conAccount.ID = aSgy.ID
		putAccount(conAccount.Addr, conAccount, s.stub)
	}

	return shim.Success([]byte("Success to submit teh user's sign message "))
}

