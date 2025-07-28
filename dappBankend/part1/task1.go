package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"

	//"github.com/ethereum/go-ethereum/log"
	"golang.org/x/net/context"
	"math/big"
	"os"
)

/*
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。

	具体任务

环境搭建
安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
查询区块
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
输出查询结果到控制台。
发送交易
准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
对交易进行签名，并将签名后的交易发送到网络。
输出交易的哈希值。
*/
func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AKVVIRU9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to connect to the Ethereum network")
	}

	_block, err := client.BlockByNumber(context.Background(), big.NewInt(8850373))
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("block hash:%v \n", _block.Hash().Hex())
	fmt.Printf("block timestamp:%v \n", _block.Time())
	fmt.Printf("block transaction count:%v \n", len(_block.Transactions()))

	private_key := os.Getenv("SEPOLIA_PRIVATE_KEY_THREE")
	fmt.Printf("privateKey:%v \n", private_key)

	privateEcdsa, err := crypto.HexToECDSA(private_key)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	publickEcdsa := privateEcdsa.PublicKey

	fromAddress := crypto.PubkeyToAddress(publickEcdsa)
	toAddress := common.HexToAddress("0x4FC907Ecf5908Cf304424743A197cfe755292252")
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	gasLimit := uint64(800000)

	chainId, err := client.NetworkID(context.Background())
	tx := types.NewTx(&types.DynamicFeeTx{
		To:        &toAddress,
		Gas:       gasLimit,
		GasFeeCap: price,
		GasTipCap: price,
		Value:     big.NewInt(1e14),
		Nonce:     nonce,
	})

	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateEcdsa)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signTx)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("transaction hash:%v \n", signTx.Hash().Hex())
}
