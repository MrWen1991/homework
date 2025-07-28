package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"homework/counter"
	"log"
	"os"
)

/*
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
 具体任务
编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。
使用 abigen 生成 Go 绑定代码
安装 abigen 工具。
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。
*/

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AKVVIRU9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Failed to connect to the Ethereum network")
	}
	contractAddress := common.HexToAddress("0xB5D2d8FAC1ed6Cca036e4e3863b070c189fe2638")
	CounterContract, err := counter.NewCounter(contractAddress, client)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// privateKey
	privateKey := os.Getenv("SEPOLIA_PRIVATE_KEY")
	privateEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())

	opts, err := bind.NewKeyedTransactorWithChainID(privateEcdsa, chainId)

	tx, err := CounterContract.IncrementAndGet(opts)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	res, err := CounterContract.Get(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fmt.Printf("res:%v \n", res)

}
