package types

import "time"

// only for admin client's app client
type TransactionInfo struct {
	// 自己包装的一个结构
	TransactionID string
	OrgName       string
	Date          string
}

func NewTransactionInfo(txid, orgname string) *TransactionInfo {
	return &TransactionInfo{
		TransactionID: txid,
		OrgName:       orgname,
		Date:          time.Now().Format("2006/1/2 15:04:05"),
	}
}

type BlockInfo struct {
	// lc.QueryBlockByTxID(txid) (如果 tx 没有提交到账本则这里查询不到与 tx 相关联的 block)
	// https://github.com/hyperledger/fabric-protos-go/blob/main/common/common.pb.go#L655
	BlockNumber  uint64
	DataHash     []byte
	PreviousHash []byte
}

type ChannelInfo struct {
	// lc.QueryConfig()
	// https://github.com/SWU-Blockchain/fabric-sdk-go/blob/main/pkg/common/providers/fab/channel.go#L53
	ChannelID         string
	Height            uint64
	CurrentBlockHash  []byte
	PreviousBlockHash []byte
	Orderers          []string
	Endorser          string
	Status            int32
}
