package main

import (
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/gendataset/fff"
)

var (
	source_lvdb_path    = ""       //原数据库lvdb的地址
	source_ancient_path = ""       //原数据库ancient的地址
	tmp_lvdb_path       = ""       //临时leveldb的位置 随便填
	tmp_cfg_path        = ""       //临时配置文件的位置	随便填
	startBlock_number   = 16000000 //开始区块号码 16000000
	endBlock_number     = 16150000 //结束区块号码 16150000
	output_CSV_path     = ""       //输出文件位置
)

func main() {
	dbSSD, _ := rawdb.NewLevelDBDatabaseWithFreezer(source_lvdb_path, 1024, 256, source_ancient_path, "", true)

	defer dbSSD.Close()
	read_db := dbSSD

	writeTmp_db_lvdb1, _ := rawdb.NewLevelDBDatabase(tmp_lvdb_path, 1024, 256, "", false)
	var (
		startC = uint64(startBlock_number)
		endC   = uint64(endBlock_number)
	)

	go fff.Gendata(startC, endC, &writeTmp_db_lvdb1, &read_db, output_CSV_path, tmp_cfg_path)
	for {
	}

}
