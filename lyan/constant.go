package lyan

import logging "mypro/mycc/go-logging"

var (
	logger = logging.MustGetLogger("lyan")
)

const (
	//InitOrg1AcountNum 表示初始化价值数量
	InitOrg1AcountNum = 500000

	//InitOrg2AcountNum 表示初始化价值数量
	InitOrg2AcountNum = 300000

	//NormolUser 表示账户类别为一般账户
	NormolUser = 0

	//ContractUser 表示账户类别为合约账户
	ContractUser = 1

	//UserSiganedindexName 表示创建签名存档复合键时的indexName
	UserSiganedindexName      = "userSignedMessage"
	ASgyCompositeKeyIndexName = "ASgyCompositeKeyIndexName"
	TxCompositeKeyIndexName   = "TX"
)
