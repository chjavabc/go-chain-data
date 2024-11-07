package main

//获取链上数据 https://learnblockchain.cn/article/6307

import (
	"go-chain-data/config"
	"go-chain-data/global"
	"go-chain-data/pkg/blockchain"
	"log"
)

func init() {
	config.SetupConfig()
	config.SetupDBEngine()

	err := config.MigrateDb()
	if err != nil {
		log.Panic("config.MigrateDb error : ", err)
	}
	config.SetupEthClient()
}

func main() {
	log.Println(global.BlockChainConfig.RpcUrl)

	//初始化
	blockchain.InitBlock()

	//监听区块,并数据持久化
	blockchain.SyncTask()

	//数据库 新增加一条记录
	//block := models.Blocks{
	//	BlockHeight:       1,
	//	BlockHash:         "hash",
	//	ParentHash:        "parentHash",
	//	LatestBlockHeight: 2,
	//}
	//
	//err := block.Insert()
	//if err != nil {
	//	log.Panic("block.Insert error : ", err)
	//}

	//获取当前区块
	//blockNumber, err := global.EthRpcClient.BlockNumber(context.Background())
	//if err != nil {
	//	log.Panic("EthRpcClient.BlockNumber error : ", err)
	//}
	//log.Println("blockNumber is : ", blockNumber)

	//根据区块获取当前区块信息
	//lastBlock, err := global.EthRpcClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
	//
	//if err != nil {
	//	log.Panic("EthRpcClient.BlockByNumber error : ", err)
	//}
	//log.Println("lastBlock is : ", lastBlock)
	//
	//for i := range lastBlock.Transactions() {
	//	tx := lastBlock.Transactions()[i]
	//
	//	//交易回执信息
	//	receipt, err := global.EthRpcClient.TransactionReceipt(context.Background(), tx.Hash())
	//
	//	if err != nil {
	//		log.Panic("EthRpcClient.TransactionReceipt error : ", err)
	//	}
	//log.Println("receipt is : ", receipt)

	//}
}
