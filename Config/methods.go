package Config

import (
	Faucet "ManageSystem/Project/go"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	// "fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	// "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

const (
	chainID       = 1337
	gasLimit      = uint64(9000000)
	FaucetAddress = "0x5FDAD3A2e1DBaB8e71e75A378E20Dfba44B68255" //合约地址

)

// GetClient 获取 ethclient.Client 对象
func GetClient() (*ethclient.Client, error) {
	rpcDial, err := rpc.Dial("http://127.0.0.1:7545")
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(rpcDial)
	return client, nil
}

// GetFacetTx 获取 Faucet.ApinameTransactor 对象
func GetFacetTx(client *ethclient.Client) (*Faucet.MainTransactor, error) {
	contract, err := Faucet.NewMainTransactor(common.HexToAddress(FaucetAddress), client)
	return contract, err
}

// GetCallFaceTx 获取 Faucet.ApinameCaller 对象
func GetCallFaceTx(client ethclient.Client) (*Faucet.MainCaller, error) {
	contract, err := Faucet.NewMainCaller(common.HexToAddress(FaucetAddress), &client)
	return contract, err
}

// 初步获取 *bind.TransactOpts 对象
func setopts(client *ethclient.Client, privateKey *ecdsa.PrivateKey, address common.Address, value int64) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		panic(err)
	}
	nonce, err := Getnonce(client, address)
	if err != nil {
		panic(err)
	}
	gasPrice, err := GetgasPrice(client)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)
	auth.GasLimit = uint64(6721975)
	auth.GasPrice = gasPrice
	return auth
}

// GetTxopts 获取最终 *bind.TransactOpts 对象
func GetTxopts(client *ethclient.Client, value int64) *bind.TransactOpts {
	privateKey, publicKey := GetPublicKeyPrivateKey()
	opts := setopts(client, privateKey, publicKey, value)
	return opts
}

// Getnonce 获取 nonce
func Getnonce(client *ethclient.Client, address common.Address) (uint64, error) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return 0, err
	} else {
		return nonce, nil
	}

}

// GetPublicKeyPrivateKey 获取公钥和私钥
func GetPublicKeyPrivateKey() (*ecdsa.PrivateKey, common.Address) {
	//PK := "7d5b2c1647d259eb62a235edde07daa67db695d355b43d72cde8762c792060cb" //用户私钥
	PK := "ad111061c1fe75d3880a5080ecbdca88a0f864c5c0b49df5cff12376d6fde716"
	//"e2298eb11d40f767c515af75cd9f4b14e6c4d35685768afde6b0c3c4ed8945de"
	privateKey, err := crypto.HexToECDSA(PK)
	if err != nil {
		panic(err)
	}
	//publicKey := common.HexToAddress("0x7fE3bC3FCc9d7d19254D4DE515eC395dA9768331") //用户地址
	publicKey := common.HexToAddress("0x5Ca22f6C1162Da1db405f4AFBD670A50c15c093a") //用户地址
	//0xE13BE5b9d25f42fC0a9f391d58962cB947F27A81
	return privateKey, publicKey
}

// GetgasPrice 获取 gasPrice
func GetgasPrice(client *ethclient.Client) (*big.Int, error) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(0), err
	} else {
		return gasPrice, nil
	}

}

// GetContractAmount 调用 charity合约合约 查看合约ETH数目
//func GetContractAmount(contract *Faucet.MainCaller) (*big.Int, error) {
//
//	res, err := contract.GetMsgBalance(nil)
//	if err != nil {
//		return big.NewInt(0), err
//	}
//	return res, nil
//}

//函数调用/////

// 调用合约 查看合约某用户所有资产情况，并在前台展示
func GetallAssets(contract *Faucet.MainCaller, initiator common.Address) ([]Faucet.AssetManagementAsset, error) {

	res, err := contract.GetpersonAssets(nil, initiator)
	if err != nil {
		return *new([]Faucet.AssetManagementAsset), err
	}
	// fmt.Println(res)
	return res, nil
}

// 调用合约 查看合约某用户所有资产领用申请情况，并在前台展示
func GetAssetReceive(contract *Faucet.MainCaller) ([]Faucet.AssetManagementreceiveAsset, error) {

	res, err := contract.GetReceiveAsset(nil)
	if err != nil {
		return *new([]Faucet.AssetManagementreceiveAsset), err
	}
	// fmt.Println(res)
	return res, nil
}

// 查看所有资产库存情况
func GetRecords(contract *Faucet.MainCaller) ([]Faucet.AssetManagementRecord, error) {

	res, err := contract.GetAllRecords(nil)
	if err != nil {
		return *new([]Faucet.AssetManagementRecord), err
	}
	// fmt.Println(res)
	return res, nil
}

// 查看所有资产库存情况
func GetAllGroup(contract *Faucet.MainCaller) ([]Faucet.AssetManagementAsset, error) {

	res, err := contract.GetAllAssets(nil)
	if err != nil {
		return *new([]Faucet.AssetManagementAsset), err
	}
	// fmt.Println(res)
	return res, nil
}

//// 调用合约 查看合约某用户某个资产情况，并在前台展示
//func GetOneAssets(contract *Faucet.MainCaller, id *big.Int, initiator common.Address) (Faucet.AssetManagementAsset, error) {
//
//	res, err := contract.GetAssetDetails(nil, id, initiator)
//	if err != nil {
//		return *new(Faucet.AssetManagementAsset), err
//	}
//	return res, nil
//}

// // contribute 调用 charity合约 用户捐助
//
//	func Contribute(client *ethclient.Client, contract *Faucet.MainTransactor, id *big.Int, value int64) (*types.Transaction, error) {
//		opts := GetTxopts(client, value)
//		res, err := contract.Contribute(opts, id)
//		if err != nil {
//			return nil, err
//		}
//		return res, nil
//	}
//
//// NewAddmin 调用合约 新增管理员, 仅管理员调用
//func NewAddmin(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address) (*types.Transaction, error) {
//	opts := GetTxopts(client, 0)
//	res, err := contract.AddAdmin(opts, initiator)
//	if err != nil {
//		return nil, err
//	}
//	return res, nil
//}

// NewAsset 调用合约 新增用户资产申请服务
func NewAsset(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, name string, description string) (*types.Transaction, error) {
	fmt.Println("新增资产")
	opts := GetTxopts(client, 0)
	res, err := contract.CreateAsset(opts, initiator, name, description)
	if err != nil {
		return nil, err
	}
	fmt.Println("新增资产,res:", res.Data())
	return res, nil
}

// 用户领用资产
func ReceiveAsset(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, id *big.Int) (*types.Transaction, error) {
	fmt.Println("用户领用资产")
	opts := GetTxopts(client, 0)
	res, err := contract.PersonReceiveAsset(opts, initiator, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("领用资产,res:", res.Data())
	return res, nil
}

// 用户归还资产
func ReturnsAsset(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, id *big.Int) (*types.Transaction, error) {
	fmt.Println("用户归还资产")
	opts := GetTxopts(client, 0)
	res, err := contract.PersonReturnAsset(opts, initiator, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("归还资产,res:", res.Data())
	return res, nil
}

// RequestAssetTransfers 调用合约 请求资产转让的函数，生成转让请求
func RequestAssetTransfer(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, id *big.Int, to common.Address) (*types.Transaction, error) {

	opts := GetTxopts(client, 0)
	res, err := contract.RequestAssetTransfer(opts, initiator, id, to)
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}

	return res, nil
}

// ApproveAssetTransfer 调用合约 批准资产转让的函数，仅允许管理员调用
func ApproveAssetTransfer(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, id *big.Int) (*types.Transaction, error) {
	fmt.Println("批准资产转让的函数")
	opts := GetTxopts(client, 0)
	res, err := contract.ApproveAssetTransfer(opts, initiator, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ApproveAssetTransfer 调用合约 批准资产转让的函数，仅允许管理员调用
func ReceiveApprove(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, id *big.Int) (*types.Transaction, error) {
	fmt.Println("批准资产领用的函数")
	opts := GetTxopts(client, 0)
	res, err := contract.ApproveAsset(opts, initiator, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ApproveAssetTransfer 调用合约 批准资产转让的函数，仅允许管理员调用
func RejectApprove(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, id *big.Int) (*types.Transaction, error) {
	fmt.Println("拒绝资产领用的函数")
	opts := GetTxopts(client, 0)
	res, err := contract.RejectAsset(opts, initiator, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RefuseAssetTransfer 调用合约 批准资产转让的函数，仅允许管理员调用
func RefuseAssetTransfer(client *ethclient.Client, contract *Faucet.MainTransactor, initiator common.Address, id *big.Int) (*types.Transaction, error) {
	fmt.Println("拒绝资产转让的函数")
	opts := GetTxopts(client, 0)
	res, err := contract.RefuseAssetRequest(opts, initiator, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAllRequests 调用合约 获取所有发起的资产转移请求 ,仅允许管理员调用
func GetAllRequests(contract *Faucet.MainCaller) ([]Faucet.AssetManagementRequest, error) {
	//opts := GetTxopts(client, 0)
	res, err := contract.GetAllRequests(nil)
	if err != nil {
		return *new([]Faucet.AssetManagementRequest), err
	}
	return res, nil
}

// GetpersonRequests 调用合约 获取某用户发起的所有资产转移请求
func GetpersonRequests(contract *Faucet.MainCaller, initiator common.Address) ([]Faucet.AssetManagementRequest, error) {

	res, err := contract.GetpersonRequests(nil, initiator)
	if err != nil {
		return *new([]Faucet.AssetManagementRequest), err
	}
	return res, nil
}
