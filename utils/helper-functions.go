package utils

import (
	"automator/constants"
	"automator/contracts/solmate"
	"automator/contracts/web3"
	"automator/controllers"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Data struct {
	Status bool
	gas    uint64
}

func Automate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	RPC_URL := constants.RPC_URL
	fmt.Println("Automation Initiating..")
	client, err := ethclient.DialContext(context.Background(), RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(10 * time.Second)
	<-ticker.C
	getTasks(client)
	Automate()

}

func execute(client *ethclient.Client, address common.Address, abi string, privateKey *ecdsa.PrivateKey) *Data {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	contract, err := web3.NewContract(address, client, abi)
	if err != nil {
		log.Fatal(err)
	}
	status := false
	gasUsed := 0

	automationCheck, err := contract.CheckAutomationStatus(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status : %v", automationCheck)
	if automationCheck {
		tx, err := contract.Automate(auth)
		if err != nil {
			log.Fatal(err)
		}
		receipt, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			log.Fatal(err)
		}

		if receipt.Status == 1 {
			status = true
			gasUsed = int(receipt.GasUsed)
		} else {
			status = false
			gasUsed = 0
		}

		fmt.Printf("%v\n", auth.GasPrice.Uint64()*receipt.GasUsed)

	}
	return &Data{Status: status, gas: uint64(gasUsed)}

}

func getTasks(client *ethclient.Client) {
	key := constants.KEY_ENCRYPTED
	contractAddress := constants.CONTRACT_ADDRESS_ENCRYPTED
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	contract, err := solmate.NewSolmate(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	data, _ := contract.GetAllTasks(nil)
	// data, _ := contract.GetAllTasks(nil)
	fmt.Println(len(data))
	for i := 0; i < len(data); i++ {
		task := controllers.GetTask(data[i].TaskAddress.String())
		if data[i].State == 0 {
			response := execute(client, common.HexToAddress(data[i].TaskAddress.Hex()), task.Abi, privateKey)
			if response.Status {
				ok := updateDetails(contract, client, common.HexToAddress(data[i].TaskAddress.Hex()), big.NewInt(auth.GasPrice.Int64()*int64(response.gas)), privateKey)
				fmt.Printf("\n%v\n", ok)
				if ok {
					fmt.Println("done")
				}
			}
		}
	}
}

func updateDetails(contract *solmate.Solmate, client *ethclient.Client, address common.Address, amount *big.Int, privateKey *ecdsa.PrivateKey) bool {
	fmt.Println("updating Details")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tx, err := contract.UpdateTaskExecDetails(auth, address, amount)
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	status := false
	fmt.Printf("\nhttps://goerli.etherscan.io/tx/%v\n", receipt.TxHash)
	if receipt.Status == 1 {
		status = true
	} else {
		status = false
	}
	return status
}
