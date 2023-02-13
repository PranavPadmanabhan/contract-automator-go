// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solmate

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SolMateAutoTask is an auto generated low-level Go binding around an user-defined struct.
type SolMateAutoTask struct {
	Id               *big.Int
	TaskAddress      common.Address
	GasLimit         *big.Int
	State            uint8
	Creator          common.Address
	TotalCostForExec *big.Int
}

// SolmateMetaData contains all meta data concerning the Solmate contract.
var SolmateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DepositFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundingFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinimumBalanceRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TopUpFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AutoTaskCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"GasLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumSolMate.TaskState\",\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalCostForExec\",\"type\":\"uint256\"}],\"name\":\"NewAutoTask\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TaskDetailsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fund\",\"type\":\"uint256\"}],\"name\":\"TaskFundWithdrawSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"}],\"name\":\"TaskFundingSuccess\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"addFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"}],\"name\":\"cancelAutomation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"createAutomation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositFundsToContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"enumSolMate.TaskState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalCostForExec\",\"type\":\"uint256\"}],\"internalType\":\"structSolMate.AutoTask[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"}],\"name\":\"getExecListOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"enumSolMate.TaskState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalCostForExec\",\"type\":\"uint256\"}],\"internalType\":\"structSolMate.AutoTask\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateTaskExecDetails\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"updateTaskGasLimit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_taskAddress\",\"type\":\"address\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050608051611e9f610067600039600081816107e50152610ec40152611e9f6000f3fe6080604052600436106100a75760003560e01c80639b96eece116100645780639b96eece14610193578063a5a01a0e146101d0578063ad296f27146101ec578063c586808114610215578063d056935d1461021f578063ea25bf0e1461025c576100a7565b8063019f232a146100ac5780631e626e40146100e9578063358e9db51461010557806368742da614610121578063779900b41461013d578063893d20e814610168575b600080fd5b3480156100b857600080fd5b506100d360048036038101906100ce919061161a565b610278565b6040516100e09190611761565b60405180910390f35b61010360048036038101906100fe91906117a8565b610462565b005b61011f600480360381019061011a91906117fb565b6107e3565b005b61013b6004803603810190610136919061161a565b610a45565b005b34801561014957600080fd5b50610152610d5c565b60405161015f9190611965565b60405180910390f35b34801561017457600080fd5b5061017d610ec0565b60405161018a9190611996565b60405180910390f35b34801561019f57600080fd5b506101ba60048036038101906101b5919061161a565b610ee8565b6040516101c791906119c0565b60405180910390f35b6101ea60048036038101906101e591906119db565b610f31565b005b3480156101f857600080fd5b50610213600480360381019061020e919061161a565b6110a8565b005b61021d6112ba565b005b34801561022b57600080fd5b506102466004803603810190610241919061161a565b6112fc565b6040516102539190611aca565b60405180910390f35b610276600480360381019061027191906117fb565b611393565b005b610280611543565b610288611543565b60005b600080549050811015610458578373ffffffffffffffffffffffffffffffffffffffff16600082815481106102c3576102c2611aec565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610445576000818154811061032457610323611aec565b5b90600052602060002090600502016040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff1660018111156103cc576103cb61166f565b5b60018111156103de576103dd61166f565b5b81526020016003820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152505091505b808061045090611b4a565b91505061028b565b5080915050919050565b6601c6bf526340003410156104a3576040517fce027d0d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006040518060c0016040528060016000805490506104c29190611b92565b81526020018573ffffffffffffffffffffffffffffffffffffffff168152602001848152602001600060018111156104fd576104fc61166f565b5b81526020013373ffffffffffffffffffffffffffffffffffffffff168152602001600081525090806001815401808255809150506001900390600052602060002090600502016000909190919091506000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548160ff021916908360018111156105d2576105d161166f565b5b021790555060808201518160030160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600401555050600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002042908060018154018082558091505060019003906000526020600020016000909190919091505534600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106df9190611b92565b9250508190555060008173ffffffffffffffffffffffffffffffffffffffff163460405161070c90611bf7565b60006040518083038185875af1925050503d8060008114610749576040519150601f19603f3d011682016040523d82523d6000602084013e61074e565b606091505b5050905080610789576040517fd481fc6c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f6713326ac36e3beb35e1e8c5c819d7fbb219746e5bb858f54c540623f5f4f15560016000805490506107bc9190611b92565b858560003360006040516107d596959493929190611c60565b60405180910390a150505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610871576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086890611d1e565b60405180910390fd5b60005b600080549050811015610a40578273ffffffffffffffffffffffffffffffffffffffff16600082815481106108ac576108ab611aec565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610a2d57600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020429080600181540180825580915050600190039060005260206000200160009091909190915055816000828154811061097457610973611aec565b5b906000526020600020906005020160040160008282546109949190611b92565b9250508190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109ea9190611d3e565b925050819055507f472dbcd999ba0ec6b1178ce58f0a27fc90a830438fe449f6bf2c63baffa34cac428484604051610a2493929190611d72565b60405180910390a15b8080610a3890611b4a565b915050610874565b505050565b600080600090505b600080549050811015610bf3578273ffffffffffffffffffffffffffffffffffffffff1660008281548110610a8557610a84611aec565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610be05760008181548110610ae657610ae5611aec565b5b906000526020600020906005020160030160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141580610ba85750600180811115610b6057610b5f61166f565b5b60008281548110610b7457610b73611aec565b5b906000526020600020906005020160030160009054906101000a900460ff166001811115610ba557610ba461166f565b5b14155b15610bdf576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b8080610beb90611b4a565b915050610a4d565b50600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060003373ffffffffffffffffffffffffffffffffffffffff1682604051610c5c90611bf7565b60006040518083038185875af1925050503d8060008114610c99576040519150601f19603f3d011682016040523d82523d6000602084013e610c9e565b606091505b5050905080610cd9576040517f750b219c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507ff404fa7f55898872111a8af61ad9ea7a6123fcc288c97073fb2dc7c7c6c38ed98383604051610d4f929190611da9565b60405180910390a1505050565b60606000805480602002602001604051908101604052809291908181526020016000905b82821015610eb757838290600052602060002090600502016040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900460ff166001811115610e3257610e3161166f565b5b6001811115610e4457610e4361166f565b5b81526020016003820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152505081526020019060010190610d80565b50505050905090565b60007f0000000000000000000000000000000000000000000000000000000000000000905090565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b655af3107a4000341015610f71576040517fce027d0d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b34600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fc09190611b92565b9250508190555060008173ffffffffffffffffffffffffffffffffffffffff1634604051610fed90611bf7565b60006040518083038185875af1925050503d806000811461102a576040519150601f19603f3d011682016040523d82523d6000602084013e61102f565b606091505b505090508061106a576040517f362f189f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fea5532f7da24511b0f68fb5d9c492ae8611139a2bb849ca0291dbf6a266a318c348460405161109b929190611dd2565b60405180910390a1505050565b60005b60008054905081101561123c57600081815481106110cc576110cb611aec565b5b906000526020600020906005020160030160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611162576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166000828154811061118d5761118c611aec565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611229576001600082815481106111f0576111ef611aec565b5b906000526020600020906005020160030160006101000a81548160ff021916908360018111156112235761122261166f565b5b02179055505b808061123490611b4a565b9150506110ab565b507ff05dadadae71396048aa495562575dc275f44dcb28c57c9224cff0b70fc1a77581600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054336040516112af93929190611dfb565b60405180910390a150565b6509184e72a0003410156112fa576040517f79cacff100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561138757602002820191906000526020600020905b815481526020019060010190808311611373575b50505050509050919050565b60005b60008054905081101561153e578273ffffffffffffffffffffffffffffffffffffffff16600082815481106113ce576113cd611aec565b5b906000526020600020906005020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361152b576000818154811061142f5761142e611aec565b5b906000526020600020906005020160030160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114c5576040517fea8e4eb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81600082815481106114da576114d9611aec565b5b9060005260206000209060050201600201819055507fb13cace0335cf27f9a92f0c874a39bb19ad3863131152b467d9a0cae23e2043082843360405161152293929190611e32565b60405180910390a15b808061153690611b4a565b915050611396565b505050565b6040518060c0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000600181111561158d5761158c61166f565b5b8152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006115e7826115bc565b9050919050565b6115f7816115dc565b811461160257600080fd5b50565b600081359050611614816115ee565b92915050565b6000602082840312156116305761162f6115b7565b5b600061163e84828501611605565b91505092915050565b6000819050919050565b61165a81611647565b82525050565b611669816115dc565b82525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281106116af576116ae61166f565b5b50565b60008190506116c08261169e565b919050565b60006116d0826116b2565b9050919050565b6116e0816116c5565b82525050565b60c0820160008201516116fc6000850182611651565b50602082015161170f6020850182611660565b5060408201516117226040850182611651565b50606082015161173560608501826116d7565b5060808201516117486080850182611660565b5060a082015161175b60a0850182611651565b50505050565b600060c08201905061177660008301846116e6565b92915050565b61178581611647565b811461179057600080fd5b50565b6000813590506117a28161177c565b92915050565b6000806000606084860312156117c1576117c06115b7565b5b60006117cf86828701611605565b93505060206117e086828701611793565b92505060406117f186828701611605565b9150509250925092565b60008060408385031215611812576118116115b7565b5b600061182085828601611605565b925050602061183185828601611793565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60c08201600082015161187d6000850182611651565b5060208201516118906020850182611660565b5060408201516118a36040850182611651565b5060608201516118b660608501826116d7565b5060808201516118c96080850182611660565b5060a08201516118dc60a0850182611651565b50505050565b60006118ee8383611867565b60c08301905092915050565b6000602082019050919050565b60006119128261183b565b61191c8185611846565b935061192783611857565b8060005b8381101561195857815161193f88826118e2565b975061194a836118fa565b92505060018101905061192b565b5085935050505092915050565b6000602082019050818103600083015261197f8184611907565b905092915050565b611990816115dc565b82525050565b60006020820190506119ab6000830184611987565b92915050565b6119ba81611647565b82525050565b60006020820190506119d560008301846119b1565b92915050565b600080604083850312156119f2576119f16115b7565b5b6000611a0085828601611605565b9250506020611a1185828601611605565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000611a538383611651565b60208301905092915050565b6000602082019050919050565b6000611a7782611a1b565b611a818185611a26565b9350611a8c83611a37565b8060005b83811015611abd578151611aa48882611a47565b9750611aaf83611a5f565b925050600181019050611a90565b5085935050505092915050565b60006020820190508181036000830152611ae48184611a6c565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b5582611647565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b8757611b86611b1b565b5b600182019050919050565b6000611b9d82611647565b9150611ba883611647565b9250828201905080821115611bc057611bbf611b1b565b5b92915050565b600081905092915050565b50565b6000611be1600083611bc6565b9150611bec82611bd1565b600082019050919050565b6000611c0282611bd4565b9150819050919050565b611c15816116c5565b82525050565b6000819050919050565b6000819050919050565b6000611c4a611c45611c4084611c1b565b611c25565b611647565b9050919050565b611c5a81611c2f565b82525050565b600060c082019050611c7560008301896119b1565b611c826020830188611987565b611c8f60408301876119b1565b611c9c6060830186611c0c565b611ca96080830185611987565b611cb660a0830184611c51565b979650505050505050565b600082825260208201905092915050565b7f75736572206973206e6f74206f776e6572000000000000000000000000000000600082015250565b6000611d08601183611cc1565b9150611d1382611cd2565b602082019050919050565b60006020820190508181036000830152611d3781611cfb565b9050919050565b6000611d4982611647565b9150611d5483611647565b9250828203905081811115611d6c57611d6b611b1b565b5b92915050565b6000606082019050611d8760008301866119b1565b611d946020830185611987565b611da160408301846119b1565b949350505050565b6000604082019050611dbe6000830185611987565b611dcb60208301846119b1565b9392505050565b6000604082019050611de760008301856119b1565b611df46020830184611987565b9392505050565b6000606082019050611e106000830186611987565b611e1d60208301856119b1565b611e2a6040830184611987565b949350505050565b6000606082019050611e4760008301866119b1565b611e546020830185611987565b611e616040830184611987565b94935050505056fea26469706673582212204d822b16bc254b61098dc77ad2be4462ede4a13d489d0cbcfd7beeba398bfdc464736f6c63430008120033",
}

// SolmateABI is the input ABI used to generate the binding from.
// Deprecated: Use SolmateMetaData.ABI instead.
var SolmateABI = SolmateMetaData.ABI

// SolmateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolmateMetaData.Bin instead.
var SolmateBin = SolmateMetaData.Bin

// DeploySolmate deploys a new Ethereum contract, binding an instance of Solmate to it.
func DeploySolmate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Solmate, error) {
	parsed, err := SolmateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolmateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Solmate{SolmateCaller: SolmateCaller{contract: contract}, SolmateTransactor: SolmateTransactor{contract: contract}, SolmateFilterer: SolmateFilterer{contract: contract}}, nil
}

// Solmate is an auto generated Go binding around an Ethereum contract.
type Solmate struct {
	SolmateCaller     // Read-only binding to the contract
	SolmateTransactor // Write-only binding to the contract
	SolmateFilterer   // Log filterer for contract events
}

// SolmateCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolmateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolmateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolmateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolmateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolmateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolmateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolmateSession struct {
	Contract     *Solmate          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolmateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolmateCallerSession struct {
	Contract *SolmateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SolmateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolmateTransactorSession struct {
	Contract     *SolmateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SolmateRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolmateRaw struct {
	Contract *Solmate // Generic contract binding to access the raw methods on
}

// SolmateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolmateCallerRaw struct {
	Contract *SolmateCaller // Generic read-only contract binding to access the raw methods on
}

// SolmateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolmateTransactorRaw struct {
	Contract *SolmateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolmate creates a new instance of Solmate, bound to a specific deployed contract.
func NewSolmate(address common.Address, backend bind.ContractBackend) (*Solmate, error) {
	contract, err := bindSolmate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solmate{SolmateCaller: SolmateCaller{contract: contract}, SolmateTransactor: SolmateTransactor{contract: contract}, SolmateFilterer: SolmateFilterer{contract: contract}}, nil
}

// NewSolmateCaller creates a new read-only instance of Solmate, bound to a specific deployed contract.
func NewSolmateCaller(address common.Address, caller bind.ContractCaller) (*SolmateCaller, error) {
	contract, err := bindSolmate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolmateCaller{contract: contract}, nil
}

// NewSolmateTransactor creates a new write-only instance of Solmate, bound to a specific deployed contract.
func NewSolmateTransactor(address common.Address, transactor bind.ContractTransactor) (*SolmateTransactor, error) {
	contract, err := bindSolmate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolmateTransactor{contract: contract}, nil
}

// NewSolmateFilterer creates a new log filterer instance of Solmate, bound to a specific deployed contract.
func NewSolmateFilterer(address common.Address, filterer bind.ContractFilterer) (*SolmateFilterer, error) {
	contract, err := bindSolmate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolmateFilterer{contract: contract}, nil
}

// bindSolmate binds a generic wrapper to an already deployed contract.
func bindSolmate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolmateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solmate *SolmateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solmate.Contract.SolmateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solmate *SolmateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solmate.Contract.SolmateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solmate *SolmateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solmate.Contract.SolmateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solmate *SolmateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solmate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solmate *SolmateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solmate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solmate *SolmateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solmate.Contract.contract.Transact(opts, method, params...)
}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() view returns((uint256,address,uint256,uint8,address,uint256)[])
func (_Solmate *SolmateCaller) GetAllTasks(opts *bind.CallOpts) ([]SolMateAutoTask, error) {
	var out []interface{}
	err := _Solmate.contract.Call(opts, &out, "getAllTasks")

	if err != nil {
		return *new([]SolMateAutoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]SolMateAutoTask)).(*[]SolMateAutoTask)

	return out0, err

}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() view returns((uint256,address,uint256,uint8,address,uint256)[])
func (_Solmate *SolmateSession) GetAllTasks() ([]SolMateAutoTask, error) {
	return _Solmate.Contract.GetAllTasks(&_Solmate.CallOpts)
}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() view returns((uint256,address,uint256,uint8,address,uint256)[])
func (_Solmate *SolmateCallerSession) GetAllTasks() ([]SolMateAutoTask, error) {
	return _Solmate.Contract.GetAllTasks(&_Solmate.CallOpts)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _taskAddress) view returns(uint256)
func (_Solmate *SolmateCaller) GetBalanceOf(opts *bind.CallOpts, _taskAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Solmate.contract.Call(opts, &out, "getBalanceOf", _taskAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _taskAddress) view returns(uint256)
func (_Solmate *SolmateSession) GetBalanceOf(_taskAddress common.Address) (*big.Int, error) {
	return _Solmate.Contract.GetBalanceOf(&_Solmate.CallOpts, _taskAddress)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _taskAddress) view returns(uint256)
func (_Solmate *SolmateCallerSession) GetBalanceOf(_taskAddress common.Address) (*big.Int, error) {
	return _Solmate.Contract.GetBalanceOf(&_Solmate.CallOpts, _taskAddress)
}

// GetExecListOf is a free data retrieval call binding the contract method 0xd056935d.
//
// Solidity: function getExecListOf(address _taskAddress) view returns(uint256[])
func (_Solmate *SolmateCaller) GetExecListOf(opts *bind.CallOpts, _taskAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Solmate.contract.Call(opts, &out, "getExecListOf", _taskAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetExecListOf is a free data retrieval call binding the contract method 0xd056935d.
//
// Solidity: function getExecListOf(address _taskAddress) view returns(uint256[])
func (_Solmate *SolmateSession) GetExecListOf(_taskAddress common.Address) ([]*big.Int, error) {
	return _Solmate.Contract.GetExecListOf(&_Solmate.CallOpts, _taskAddress)
}

// GetExecListOf is a free data retrieval call binding the contract method 0xd056935d.
//
// Solidity: function getExecListOf(address _taskAddress) view returns(uint256[])
func (_Solmate *SolmateCallerSession) GetExecListOf(_taskAddress common.Address) ([]*big.Int, error) {
	return _Solmate.Contract.GetExecListOf(&_Solmate.CallOpts, _taskAddress)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Solmate *SolmateCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Solmate.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Solmate *SolmateSession) GetOwner() (common.Address, error) {
	return _Solmate.Contract.GetOwner(&_Solmate.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Solmate *SolmateCallerSession) GetOwner() (common.Address, error) {
	return _Solmate.Contract.GetOwner(&_Solmate.CallOpts)
}

// GetTask is a free data retrieval call binding the contract method 0x019f232a.
//
// Solidity: function getTask(address _taskAddress) view returns((uint256,address,uint256,uint8,address,uint256))
func (_Solmate *SolmateCaller) GetTask(opts *bind.CallOpts, _taskAddress common.Address) (SolMateAutoTask, error) {
	var out []interface{}
	err := _Solmate.contract.Call(opts, &out, "getTask", _taskAddress)

	if err != nil {
		return *new(SolMateAutoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(SolMateAutoTask)).(*SolMateAutoTask)

	return out0, err

}

// GetTask is a free data retrieval call binding the contract method 0x019f232a.
//
// Solidity: function getTask(address _taskAddress) view returns((uint256,address,uint256,uint8,address,uint256))
func (_Solmate *SolmateSession) GetTask(_taskAddress common.Address) (SolMateAutoTask, error) {
	return _Solmate.Contract.GetTask(&_Solmate.CallOpts, _taskAddress)
}

// GetTask is a free data retrieval call binding the contract method 0x019f232a.
//
// Solidity: function getTask(address _taskAddress) view returns((uint256,address,uint256,uint8,address,uint256))
func (_Solmate *SolmateCallerSession) GetTask(_taskAddress common.Address) (SolMateAutoTask, error) {
	return _Solmate.Contract.GetTask(&_Solmate.CallOpts, _taskAddress)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa5a01a0e.
//
// Solidity: function addFunds(address _taskAddress, address executor) payable returns()
func (_Solmate *SolmateTransactor) AddFunds(opts *bind.TransactOpts, _taskAddress common.Address, executor common.Address) (*types.Transaction, error) {
	return _Solmate.contract.Transact(opts, "addFunds", _taskAddress, executor)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa5a01a0e.
//
// Solidity: function addFunds(address _taskAddress, address executor) payable returns()
func (_Solmate *SolmateSession) AddFunds(_taskAddress common.Address, executor common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.AddFunds(&_Solmate.TransactOpts, _taskAddress, executor)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa5a01a0e.
//
// Solidity: function addFunds(address _taskAddress, address executor) payable returns()
func (_Solmate *SolmateTransactorSession) AddFunds(_taskAddress common.Address, executor common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.AddFunds(&_Solmate.TransactOpts, _taskAddress, executor)
}

// CancelAutomation is a paid mutator transaction binding the contract method 0xad296f27.
//
// Solidity: function cancelAutomation(address _taskAddress) returns()
func (_Solmate *SolmateTransactor) CancelAutomation(opts *bind.TransactOpts, _taskAddress common.Address) (*types.Transaction, error) {
	return _Solmate.contract.Transact(opts, "cancelAutomation", _taskAddress)
}

// CancelAutomation is a paid mutator transaction binding the contract method 0xad296f27.
//
// Solidity: function cancelAutomation(address _taskAddress) returns()
func (_Solmate *SolmateSession) CancelAutomation(_taskAddress common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.CancelAutomation(&_Solmate.TransactOpts, _taskAddress)
}

// CancelAutomation is a paid mutator transaction binding the contract method 0xad296f27.
//
// Solidity: function cancelAutomation(address _taskAddress) returns()
func (_Solmate *SolmateTransactorSession) CancelAutomation(_taskAddress common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.CancelAutomation(&_Solmate.TransactOpts, _taskAddress)
}

// CreateAutomation is a paid mutator transaction binding the contract method 0x1e626e40.
//
// Solidity: function createAutomation(address _address, uint256 _gasLimit, address executor) payable returns()
func (_Solmate *SolmateTransactor) CreateAutomation(opts *bind.TransactOpts, _address common.Address, _gasLimit *big.Int, executor common.Address) (*types.Transaction, error) {
	return _Solmate.contract.Transact(opts, "createAutomation", _address, _gasLimit, executor)
}

// CreateAutomation is a paid mutator transaction binding the contract method 0x1e626e40.
//
// Solidity: function createAutomation(address _address, uint256 _gasLimit, address executor) payable returns()
func (_Solmate *SolmateSession) CreateAutomation(_address common.Address, _gasLimit *big.Int, executor common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.CreateAutomation(&_Solmate.TransactOpts, _address, _gasLimit, executor)
}

// CreateAutomation is a paid mutator transaction binding the contract method 0x1e626e40.
//
// Solidity: function createAutomation(address _address, uint256 _gasLimit, address executor) payable returns()
func (_Solmate *SolmateTransactorSession) CreateAutomation(_address common.Address, _gasLimit *big.Int, executor common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.CreateAutomation(&_Solmate.TransactOpts, _address, _gasLimit, executor)
}

// DepositFundsToContract is a paid mutator transaction binding the contract method 0xc5868081.
//
// Solidity: function depositFundsToContract() payable returns()
func (_Solmate *SolmateTransactor) DepositFundsToContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solmate.contract.Transact(opts, "depositFundsToContract")
}

// DepositFundsToContract is a paid mutator transaction binding the contract method 0xc5868081.
//
// Solidity: function depositFundsToContract() payable returns()
func (_Solmate *SolmateSession) DepositFundsToContract() (*types.Transaction, error) {
	return _Solmate.Contract.DepositFundsToContract(&_Solmate.TransactOpts)
}

// DepositFundsToContract is a paid mutator transaction binding the contract method 0xc5868081.
//
// Solidity: function depositFundsToContract() payable returns()
func (_Solmate *SolmateTransactorSession) DepositFundsToContract() (*types.Transaction, error) {
	return _Solmate.Contract.DepositFundsToContract(&_Solmate.TransactOpts)
}

// UpdateTaskExecDetails is a paid mutator transaction binding the contract method 0x358e9db5.
//
// Solidity: function updateTaskExecDetails(address _taskAddress, uint256 amount) payable returns()
func (_Solmate *SolmateTransactor) UpdateTaskExecDetails(opts *bind.TransactOpts, _taskAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Solmate.contract.Transact(opts, "updateTaskExecDetails", _taskAddress, amount)
}

// UpdateTaskExecDetails is a paid mutator transaction binding the contract method 0x358e9db5.
//
// Solidity: function updateTaskExecDetails(address _taskAddress, uint256 amount) payable returns()
func (_Solmate *SolmateSession) UpdateTaskExecDetails(_taskAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Solmate.Contract.UpdateTaskExecDetails(&_Solmate.TransactOpts, _taskAddress, amount)
}

// UpdateTaskExecDetails is a paid mutator transaction binding the contract method 0x358e9db5.
//
// Solidity: function updateTaskExecDetails(address _taskAddress, uint256 amount) payable returns()
func (_Solmate *SolmateTransactorSession) UpdateTaskExecDetails(_taskAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Solmate.Contract.UpdateTaskExecDetails(&_Solmate.TransactOpts, _taskAddress, amount)
}

// UpdateTaskGasLimit is a paid mutator transaction binding the contract method 0xea25bf0e.
//
// Solidity: function updateTaskGasLimit(address _taskAddress, uint256 _gasLimit) payable returns()
func (_Solmate *SolmateTransactor) UpdateTaskGasLimit(opts *bind.TransactOpts, _taskAddress common.Address, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Solmate.contract.Transact(opts, "updateTaskGasLimit", _taskAddress, _gasLimit)
}

// UpdateTaskGasLimit is a paid mutator transaction binding the contract method 0xea25bf0e.
//
// Solidity: function updateTaskGasLimit(address _taskAddress, uint256 _gasLimit) payable returns()
func (_Solmate *SolmateSession) UpdateTaskGasLimit(_taskAddress common.Address, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Solmate.Contract.UpdateTaskGasLimit(&_Solmate.TransactOpts, _taskAddress, _gasLimit)
}

// UpdateTaskGasLimit is a paid mutator transaction binding the contract method 0xea25bf0e.
//
// Solidity: function updateTaskGasLimit(address _taskAddress, uint256 _gasLimit) payable returns()
func (_Solmate *SolmateTransactorSession) UpdateTaskGasLimit(_taskAddress common.Address, _gasLimit *big.Int) (*types.Transaction, error) {
	return _Solmate.Contract.UpdateTaskGasLimit(&_Solmate.TransactOpts, _taskAddress, _gasLimit)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address _taskAddress) payable returns()
func (_Solmate *SolmateTransactor) WithdrawFunds(opts *bind.TransactOpts, _taskAddress common.Address) (*types.Transaction, error) {
	return _Solmate.contract.Transact(opts, "withdrawFunds", _taskAddress)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address _taskAddress) payable returns()
func (_Solmate *SolmateSession) WithdrawFunds(_taskAddress common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.WithdrawFunds(&_Solmate.TransactOpts, _taskAddress)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address _taskAddress) payable returns()
func (_Solmate *SolmateTransactorSession) WithdrawFunds(_taskAddress common.Address) (*types.Transaction, error) {
	return _Solmate.Contract.WithdrawFunds(&_Solmate.TransactOpts, _taskAddress)
}

// SolmateAutoTaskCancelledIterator is returned from FilterAutoTaskCancelled and is used to iterate over the raw logs and unpacked data for AutoTaskCancelled events raised by the Solmate contract.
type SolmateAutoTaskCancelledIterator struct {
	Event *SolmateAutoTaskCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolmateAutoTaskCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateAutoTaskCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolmateAutoTaskCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolmateAutoTaskCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateAutoTaskCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateAutoTaskCancelled represents a AutoTaskCancelled event raised by the Solmate contract.
type SolmateAutoTaskCancelled struct {
	TaskAddress common.Address
	Balance     *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAutoTaskCancelled is a free log retrieval operation binding the contract event 0xf05dadadae71396048aa495562575dc275f44dcb28c57c9224cff0b70fc1a775.
//
// Solidity: event AutoTaskCancelled(address taskAddress, uint256 balance, address owner)
func (_Solmate *SolmateFilterer) FilterAutoTaskCancelled(opts *bind.FilterOpts) (*SolmateAutoTaskCancelledIterator, error) {

	logs, sub, err := _Solmate.contract.FilterLogs(opts, "AutoTaskCancelled")
	if err != nil {
		return nil, err
	}
	return &SolmateAutoTaskCancelledIterator{contract: _Solmate.contract, event: "AutoTaskCancelled", logs: logs, sub: sub}, nil
}

// WatchAutoTaskCancelled is a free log subscription operation binding the contract event 0xf05dadadae71396048aa495562575dc275f44dcb28c57c9224cff0b70fc1a775.
//
// Solidity: event AutoTaskCancelled(address taskAddress, uint256 balance, address owner)
func (_Solmate *SolmateFilterer) WatchAutoTaskCancelled(opts *bind.WatchOpts, sink chan<- *SolmateAutoTaskCancelled) (event.Subscription, error) {

	logs, sub, err := _Solmate.contract.WatchLogs(opts, "AutoTaskCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateAutoTaskCancelled)
				if err := _Solmate.contract.UnpackLog(event, "AutoTaskCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAutoTaskCancelled is a log parse operation binding the contract event 0xf05dadadae71396048aa495562575dc275f44dcb28c57c9224cff0b70fc1a775.
//
// Solidity: event AutoTaskCancelled(address taskAddress, uint256 balance, address owner)
func (_Solmate *SolmateFilterer) ParseAutoTaskCancelled(log types.Log) (*SolmateAutoTaskCancelled, error) {
	event := new(SolmateAutoTaskCancelled)
	if err := _Solmate.contract.UnpackLog(event, "AutoTaskCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolmateGasLimitUpdatedIterator is returned from FilterGasLimitUpdated and is used to iterate over the raw logs and unpacked data for GasLimitUpdated events raised by the Solmate contract.
type SolmateGasLimitUpdatedIterator struct {
	Event *SolmateGasLimitUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolmateGasLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateGasLimitUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolmateGasLimitUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolmateGasLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateGasLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateGasLimitUpdated represents a GasLimitUpdated event raised by the Solmate contract.
type SolmateGasLimitUpdated struct {
	GasLimit    *big.Int
	TaskAddress common.Address
	User        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGasLimitUpdated is a free log retrieval operation binding the contract event 0xb13cace0335cf27f9a92f0c874a39bb19ad3863131152b467d9a0cae23e20430.
//
// Solidity: event GasLimitUpdated(uint256 gasLimit, address taskAddress, address user)
func (_Solmate *SolmateFilterer) FilterGasLimitUpdated(opts *bind.FilterOpts) (*SolmateGasLimitUpdatedIterator, error) {

	logs, sub, err := _Solmate.contract.FilterLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &SolmateGasLimitUpdatedIterator{contract: _Solmate.contract, event: "GasLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchGasLimitUpdated is a free log subscription operation binding the contract event 0xb13cace0335cf27f9a92f0c874a39bb19ad3863131152b467d9a0cae23e20430.
//
// Solidity: event GasLimitUpdated(uint256 gasLimit, address taskAddress, address user)
func (_Solmate *SolmateFilterer) WatchGasLimitUpdated(opts *bind.WatchOpts, sink chan<- *SolmateGasLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _Solmate.contract.WatchLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateGasLimitUpdated)
				if err := _Solmate.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGasLimitUpdated is a log parse operation binding the contract event 0xb13cace0335cf27f9a92f0c874a39bb19ad3863131152b467d9a0cae23e20430.
//
// Solidity: event GasLimitUpdated(uint256 gasLimit, address taskAddress, address user)
func (_Solmate *SolmateFilterer) ParseGasLimitUpdated(log types.Log) (*SolmateGasLimitUpdated, error) {
	event := new(SolmateGasLimitUpdated)
	if err := _Solmate.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolmateNewAutoTaskIterator is returned from FilterNewAutoTask and is used to iterate over the raw logs and unpacked data for NewAutoTask events raised by the Solmate contract.
type SolmateNewAutoTaskIterator struct {
	Event *SolmateNewAutoTask // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolmateNewAutoTaskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateNewAutoTask)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolmateNewAutoTask)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolmateNewAutoTaskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateNewAutoTaskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateNewAutoTask represents a NewAutoTask event raised by the Solmate contract.
type SolmateNewAutoTask struct {
	Id               *big.Int
	TaskAddress      common.Address
	GasLimit         *big.Int
	State            uint8
	Creator          common.Address
	TotalCostForExec *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewAutoTask is a free log retrieval operation binding the contract event 0x6713326ac36e3beb35e1e8c5c819d7fbb219746e5bb858f54c540623f5f4f155.
//
// Solidity: event NewAutoTask(uint256 id, address taskAddress, uint256 gasLimit, uint8 state, address creator, uint256 totalCostForExec)
func (_Solmate *SolmateFilterer) FilterNewAutoTask(opts *bind.FilterOpts) (*SolmateNewAutoTaskIterator, error) {

	logs, sub, err := _Solmate.contract.FilterLogs(opts, "NewAutoTask")
	if err != nil {
		return nil, err
	}
	return &SolmateNewAutoTaskIterator{contract: _Solmate.contract, event: "NewAutoTask", logs: logs, sub: sub}, nil
}

// WatchNewAutoTask is a free log subscription operation binding the contract event 0x6713326ac36e3beb35e1e8c5c819d7fbb219746e5bb858f54c540623f5f4f155.
//
// Solidity: event NewAutoTask(uint256 id, address taskAddress, uint256 gasLimit, uint8 state, address creator, uint256 totalCostForExec)
func (_Solmate *SolmateFilterer) WatchNewAutoTask(opts *bind.WatchOpts, sink chan<- *SolmateNewAutoTask) (event.Subscription, error) {

	logs, sub, err := _Solmate.contract.WatchLogs(opts, "NewAutoTask")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateNewAutoTask)
				if err := _Solmate.contract.UnpackLog(event, "NewAutoTask", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAutoTask is a log parse operation binding the contract event 0x6713326ac36e3beb35e1e8c5c819d7fbb219746e5bb858f54c540623f5f4f155.
//
// Solidity: event NewAutoTask(uint256 id, address taskAddress, uint256 gasLimit, uint8 state, address creator, uint256 totalCostForExec)
func (_Solmate *SolmateFilterer) ParseNewAutoTask(log types.Log) (*SolmateNewAutoTask, error) {
	event := new(SolmateNewAutoTask)
	if err := _Solmate.contract.UnpackLog(event, "NewAutoTask", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolmateTaskDetailsUpdatedIterator is returned from FilterTaskDetailsUpdated and is used to iterate over the raw logs and unpacked data for TaskDetailsUpdated events raised by the Solmate contract.
type SolmateTaskDetailsUpdatedIterator struct {
	Event *SolmateTaskDetailsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolmateTaskDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateTaskDetailsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolmateTaskDetailsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolmateTaskDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateTaskDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateTaskDetailsUpdated represents a TaskDetailsUpdated event raised by the Solmate contract.
type SolmateTaskDetailsUpdated struct {
	Time        *big.Int
	TaskAddress common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskDetailsUpdated is a free log retrieval operation binding the contract event 0x472dbcd999ba0ec6b1178ce58f0a27fc90a830438fe449f6bf2c63baffa34cac.
//
// Solidity: event TaskDetailsUpdated(uint256 time, address taskAddress, uint256 amount)
func (_Solmate *SolmateFilterer) FilterTaskDetailsUpdated(opts *bind.FilterOpts) (*SolmateTaskDetailsUpdatedIterator, error) {

	logs, sub, err := _Solmate.contract.FilterLogs(opts, "TaskDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return &SolmateTaskDetailsUpdatedIterator{contract: _Solmate.contract, event: "TaskDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchTaskDetailsUpdated is a free log subscription operation binding the contract event 0x472dbcd999ba0ec6b1178ce58f0a27fc90a830438fe449f6bf2c63baffa34cac.
//
// Solidity: event TaskDetailsUpdated(uint256 time, address taskAddress, uint256 amount)
func (_Solmate *SolmateFilterer) WatchTaskDetailsUpdated(opts *bind.WatchOpts, sink chan<- *SolmateTaskDetailsUpdated) (event.Subscription, error) {

	logs, sub, err := _Solmate.contract.WatchLogs(opts, "TaskDetailsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateTaskDetailsUpdated)
				if err := _Solmate.contract.UnpackLog(event, "TaskDetailsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskDetailsUpdated is a log parse operation binding the contract event 0x472dbcd999ba0ec6b1178ce58f0a27fc90a830438fe449f6bf2c63baffa34cac.
//
// Solidity: event TaskDetailsUpdated(uint256 time, address taskAddress, uint256 amount)
func (_Solmate *SolmateFilterer) ParseTaskDetailsUpdated(log types.Log) (*SolmateTaskDetailsUpdated, error) {
	event := new(SolmateTaskDetailsUpdated)
	if err := _Solmate.contract.UnpackLog(event, "TaskDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolmateTaskFundWithdrawSuccessIterator is returned from FilterTaskFundWithdrawSuccess and is used to iterate over the raw logs and unpacked data for TaskFundWithdrawSuccess events raised by the Solmate contract.
type SolmateTaskFundWithdrawSuccessIterator struct {
	Event *SolmateTaskFundWithdrawSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolmateTaskFundWithdrawSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateTaskFundWithdrawSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolmateTaskFundWithdrawSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolmateTaskFundWithdrawSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateTaskFundWithdrawSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateTaskFundWithdrawSuccess represents a TaskFundWithdrawSuccess event raised by the Solmate contract.
type SolmateTaskFundWithdrawSuccess struct {
	TaskAddress common.Address
	Fund        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskFundWithdrawSuccess is a free log retrieval operation binding the contract event 0xf404fa7f55898872111a8af61ad9ea7a6123fcc288c97073fb2dc7c7c6c38ed9.
//
// Solidity: event TaskFundWithdrawSuccess(address taskAddress, uint256 fund)
func (_Solmate *SolmateFilterer) FilterTaskFundWithdrawSuccess(opts *bind.FilterOpts) (*SolmateTaskFundWithdrawSuccessIterator, error) {

	logs, sub, err := _Solmate.contract.FilterLogs(opts, "TaskFundWithdrawSuccess")
	if err != nil {
		return nil, err
	}
	return &SolmateTaskFundWithdrawSuccessIterator{contract: _Solmate.contract, event: "TaskFundWithdrawSuccess", logs: logs, sub: sub}, nil
}

// WatchTaskFundWithdrawSuccess is a free log subscription operation binding the contract event 0xf404fa7f55898872111a8af61ad9ea7a6123fcc288c97073fb2dc7c7c6c38ed9.
//
// Solidity: event TaskFundWithdrawSuccess(address taskAddress, uint256 fund)
func (_Solmate *SolmateFilterer) WatchTaskFundWithdrawSuccess(opts *bind.WatchOpts, sink chan<- *SolmateTaskFundWithdrawSuccess) (event.Subscription, error) {

	logs, sub, err := _Solmate.contract.WatchLogs(opts, "TaskFundWithdrawSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateTaskFundWithdrawSuccess)
				if err := _Solmate.contract.UnpackLog(event, "TaskFundWithdrawSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskFundWithdrawSuccess is a log parse operation binding the contract event 0xf404fa7f55898872111a8af61ad9ea7a6123fcc288c97073fb2dc7c7c6c38ed9.
//
// Solidity: event TaskFundWithdrawSuccess(address taskAddress, uint256 fund)
func (_Solmate *SolmateFilterer) ParseTaskFundWithdrawSuccess(log types.Log) (*SolmateTaskFundWithdrawSuccess, error) {
	event := new(SolmateTaskFundWithdrawSuccess)
	if err := _Solmate.contract.UnpackLog(event, "TaskFundWithdrawSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolmateTaskFundingSuccessIterator is returned from FilterTaskFundingSuccess and is used to iterate over the raw logs and unpacked data for TaskFundingSuccess events raised by the Solmate contract.
type SolmateTaskFundingSuccessIterator struct {
	Event *SolmateTaskFundingSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SolmateTaskFundingSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateTaskFundingSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SolmateTaskFundingSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SolmateTaskFundingSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateTaskFundingSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateTaskFundingSuccess represents a TaskFundingSuccess event raised by the Solmate contract.
type SolmateTaskFundingSuccess struct {
	Amount      *big.Int
	TaskAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskFundingSuccess is a free log retrieval operation binding the contract event 0xea5532f7da24511b0f68fb5d9c492ae8611139a2bb849ca0291dbf6a266a318c.
//
// Solidity: event TaskFundingSuccess(uint256 amount, address taskAddress)
func (_Solmate *SolmateFilterer) FilterTaskFundingSuccess(opts *bind.FilterOpts) (*SolmateTaskFundingSuccessIterator, error) {

	logs, sub, err := _Solmate.contract.FilterLogs(opts, "TaskFundingSuccess")
	if err != nil {
		return nil, err
	}
	return &SolmateTaskFundingSuccessIterator{contract: _Solmate.contract, event: "TaskFundingSuccess", logs: logs, sub: sub}, nil
}

// WatchTaskFundingSuccess is a free log subscription operation binding the contract event 0xea5532f7da24511b0f68fb5d9c492ae8611139a2bb849ca0291dbf6a266a318c.
//
// Solidity: event TaskFundingSuccess(uint256 amount, address taskAddress)
func (_Solmate *SolmateFilterer) WatchTaskFundingSuccess(opts *bind.WatchOpts, sink chan<- *SolmateTaskFundingSuccess) (event.Subscription, error) {

	logs, sub, err := _Solmate.contract.WatchLogs(opts, "TaskFundingSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateTaskFundingSuccess)
				if err := _Solmate.contract.UnpackLog(event, "TaskFundingSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskFundingSuccess is a log parse operation binding the contract event 0xea5532f7da24511b0f68fb5d9c492ae8611139a2bb849ca0291dbf6a266a318c.
//
// Solidity: event TaskFundingSuccess(uint256 amount, address taskAddress)
func (_Solmate *SolmateFilterer) ParseTaskFundingSuccess(log types.Log) (*SolmateTaskFundingSuccess, error) {
	event := new(SolmateTaskFundingSuccess)
	if err := _Solmate.contract.UnpackLog(event, "TaskFundingSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
