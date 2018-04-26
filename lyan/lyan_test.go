package lyan

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"mypro/hyperledger-coin/proto"
	"testing"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkAcount(t *testing.T, stub *shim.MockStub, name string, balance string, kind int) {

	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get account info")
		t.FailNow()
	}

	userdata := &Account{}
	err := proto.Unmarshal(bytes, userdata)
	if err != nil {
		fmt.Println("State", name, "failed to Unmarshal account info")
		t.FailNow()
	}

	fmt.Println("\n/*****************check user********************/")
	fmt.Println("get:" + userdata.GetAddr())
	fmt.Printf("balance: %f\n", userdata.GetBalance())
	fmt.Println("balance wanted : " + balance)
	fmt.Printf("get kind: %d\n", userdata.GetKind())
	fmt.Printf("kind wanted: %d\n", kind)
	fmt.Printf("ID %s \n", userdata.GetID())
	fmt.Println("/***********************************************/\n")
}

func checkQueryAccount(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Query", args, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", args, "failed to get result")
		t.FailNow()
	}

	queryBase64String := string(res.Payload)

	data, _ := base64.StdEncoding.DecodeString(queryBase64String)

	userdata := &Account{}
	err := proto.Unmarshal(data, userdata)
	if err != nil {
		panic(err)
		//t.FailNow()
	}
	fmt.Println("Get addr :" + userdata.GetAddr())

}

func checckTransfer(t *testing.T, stub *shim.MockStub, args [][]byte, addr string) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("transfer", args, "failed", string(res.Message))
		t.FailNow()
	}

	ID := string(res.Payload)
	fmt.Print("result :" + ID + "\n")

	name, _ := stub.CreateCompositeKey(TxCompositeKeyIndexName, []string{ID, addr})

	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get account info")
		t.FailNow()
	}
	userdata := &TX{}
	err := proto.Unmarshal(bytes, userdata)
	if err != nil {
		fmt.Println("State", name, "failed to Unmarshal account info")
		t.FailNow()
	}

	fmt.Printf("%v", userdata)

}

func chechContractAcountCreating(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("transfer", args, "failed", string(res.Message))
		t.FailNow()
	}
	logger.Debugf("chechContractAcountCreating :%s", args[1])
	checkAcount(t, stub, string(args[1]), string(0), 1)

}

//====================================================================================
//	从上一级目录中读取公钥文件，做hash，以及
//
//====================================================================================

/*func TestInit(t *testing.T) {
	scc := new(MyChaincode)
	stub := shim.NewMockStub("ex02", scc)
	Org1pubKey, _ := getKeyString("../ORG1pub.pem")
	Org2pubKey, _ := getKeyString("../ORG2pub.pem")

	//fmt.Print(Org1pubKey)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), Org1pubKey, []byte("1000"), Org2pubKey, []byte("234")})

	Org1Addr := hashData(Org1pubKey, crypto.SHA1)
	Org2Addr := hashData(Org2pubKey, crypto.SHA1)

	//fmt.Println(Org1Addr)
	//fmt.Println(Org2Addr)

	checkAcount(t, stub, Org1Addr, "1000", 0)
	checkAcount(t, stub, Org2Addr, "234", 0)
	//检查刚刚创建的两个用户
	checkQueryAccount(t, stub, [][]byte{[]byte("queryAcccount"), []byte(Org2Addr)})

}

*/

func TestTransfer(t *testing.T) {

	//初始化以及获取公私钥
	scc := new(MyChaincode)
	stub := shim.NewMockStub("ex02", scc)
	Org1pubKey, _ := getKeyString("../ORG1pub.pem")
	Org2pubKey, _ := getKeyString("../ORG2pub.pem")
	Org1priKey, _ := getKeyString("../ORG1pri.pem")
	Org2priKey, _ := getKeyString("../ORG2pri.pem")

	//生成地址
	Org1Addr := hashData(Org1pubKey, crypto.SHA1)
	Org2Addr := hashData(Org2pubKey, crypto.SHA1)

	out := []*TranferOutput{{Org2Addr, 100}}

	txInfo := &TxInfo{
		Timestamp:    time.Now().UTC().Unix(),
		InputAddr:    Org1Addr,
		InputBalance: 100,
		Nounce:       100,
		Output:       out,
		Info:         "test transfer",
	}

	//获取交易的proto编码
	txInfoProto, err := proto.Marshal(txInfo)
	if err != nil {
		fmt.Println("[TestTransfer]" + err.Error())
		t.FailNow()
	}

	//进行签名
	txInfoProtoScript, err := signMessage(txInfoProto, Org1priKey, crypto.SHA1)
	if err != nil {
		fmt.Println("[TestTransfer]" + err.Error())
		t.FailNow()
	}

	err = VerifySign(string(Org1pubKey), txInfoProto, txInfoProtoScript, crypto.SHA1)
	if err != nil {
		fmt.Println("[TestTransfer]" + err.Error())
		t.FailNow()
	}

	txExample := &TX{
		Tx:     txInfo,
		Script: string(txInfoProtoScript),
	}

	txExampleProto, _ := proto.Marshal(txExample)

	txExampleProtoString := base64.StdEncoding.EncodeToString(txExampleProto)

	// Init A=1000 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), Org1pubKey, []byte("1000"), Org2pubKey, []byte("234")})
	checckTransfer(t, stub, [][]byte{[]byte("transfer"), []byte(txExampleProtoString)}, Org1Addr)
	checkAcount(t, stub, Org1Addr, "900", 0)
	// After A=900 B=334

	conAddr := hashData([]byte("nicaicia"), crypto.SHA1)

	//检测是否能创建合约账户
	//创建分配策略
	tOrtinfo := &MemberInfo{
		[]*MemberSingle{
			&MemberSingle{
				Sequence: 0,
				Addr:     Org1Addr,
				Ort:      0.9,
			},
			&MemberSingle{
				Sequence: 1,
				Addr:     Org2Addr,
				Ort:      0.1,
			},
		},
	}

	testSgy := &AllocateSgy{
		Addr:      conAddr,
		ASgy:      tOrtinfo,
		ID:        "0",
		HasSigned: []bool{false, false},
	}

	//对分配策略求hash
	zhongjian, _ := proto.Marshal(testSgy.ASgy)
	testSgy.Hash = hashData(zhongjian, crypto.SHA1)

	testSgyProto, _ := proto.Marshal(testSgy)

	testSgyProtoBase64String := base64.StdEncoding.EncodeToString(testSgyProto)

	testSgyOrg1SignedInfo := &SigeForSgyInfo{
		Addr:      Org1Addr,
		ID:        "0",
		Seq:       0,
		Timestamp: time.Now().UTC().Unix(),
		Ort:       0.9,
	}
	testSgyOrg1SignedInfoProto, _ := proto.Marshal(testSgyOrg1SignedInfo)
	testSgyOrg1Signedmessage, _ := signMessage(testSgyOrg1SignedInfoProto, Org1priKey, crypto.SHA1)

	testSgyOrg2SignedInfo := &SigeForSgyInfo{
		Addr:      Org2Addr,
		ID:        "0",
		Seq:       1,
		Timestamp: time.Now().UTC().Unix(),
		Ort:       0.1,
	}
	testSgyOrg2SignedInfoProto, _ := proto.Marshal(testSgyOrg2SignedInfo)
	testSgyOrg2Signedmessage, _ := signMessage(testSgyOrg2SignedInfoProto, Org2priKey, crypto.SHA1)

	testSgySignedStub := &SigeForSgyStub{
		Set: []*SigeForSgy{
			{
				testSgyOrg1SignedInfo,
				string(testSgyOrg1Signedmessage),
			},
			{
				testSgyOrg2SignedInfo,
				string(testSgyOrg2Signedmessage),
			},
		},
	}

	testSgySignedStubProto, _ := proto.Marshal(testSgySignedStub)
	testSgySignedStubProtoBase64 := base64.StdEncoding.EncodeToString(testSgySignedStubProto)

	chechContractAcountCreating(t, stub, [][]byte{[]byte("createContractAcount"), []byte(conAddr), []byte(testSgyProtoBase64String), []byte(testSgySignedStubProtoBase64)})
	checkAcount(t, stub, conAddr, "0", 1)

	//对合约地址进行转账

	out = []*TranferOutput{{conAddr, 100}}
	txInfo = &TxInfo{
		Timestamp:    time.Now().UTC().Unix(),
		InputAddr:    Org1Addr,
		InputBalance: 100,
		Nounce:       100,
		Output:       out,
		Info:         "test transfer",
	}

	//获取交易的proto编码
	txInfoProto, err = proto.Marshal(txInfo)
	if err != nil {
		fmt.Println("[TestTransfer]" + err.Error())
		t.FailNow()
	}

	//进行签名
	txInfoProtoScript, err = signMessage(txInfoProto, Org1priKey, crypto.SHA1)
	if err != nil {
		fmt.Println("[TestTransfer]" + err.Error())
		t.FailNow()
	}

	err = VerifySign(string(Org1pubKey), txInfoProto, txInfoProtoScript, crypto.SHA1)
	if err != nil {
		fmt.Println("[TestTransfer]" + err.Error())
		t.FailNow()
	}

	txExample = &TX{
		Tx:     txInfo,
		Script: string(txInfoProtoScript),
	}

	txExampleProto, _ = proto.Marshal(txExample)

	txExampleProtoString = base64.StdEncoding.EncodeToString(txExampleProto)
	checkAcount(t, stub, Org1Addr, "900", 0)
	checkAcount(t, stub, Org2Addr, "334", 0)
	checckTransfer(t, stub, [][]byte{[]byte("transfer"), []byte(txExampleProtoString)}, Org1Addr)
	checkAcount(t, stub, Org1Addr, "890", 0)
	checkAcount(t, stub, Org2Addr, "344", 0)

}
