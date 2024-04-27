package Hander

import (
	mid "ManageSystem/Config"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/big"
	"strconv"
)

//
//// 获得当前账户余额,从合约
//func CountAmount(c *gin.Context) {
//	client, err := mid.GetClient()
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	contract, err := mid.GetCallFaceTx(*client)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	res, err := mid.GetContractAmount(contract)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	respOK(c, res)
//	client.Close()
//}

// 查看合约某用户所有资产情况，并在前台展示
func GetallAssets(c *gin.Context) {
	initiator := c.PostForm("initiator")
	fmt.Println("initiator:", initiator)
	initiator1 := common.HexToAddress(initiator)
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetallAssets(contract, initiator1)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 查看合约用户所有资产领用申请，并在前台展示
func GetAssetReceive(c *gin.Context) {
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetAssetReceive(contract)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 管理员批准资产领用申请，并在前台展示
func ReceiveApprove(c *gin.Context) {
	initiator := c.PostForm("initiator")
	fmt.Println("批准者账号:", initiator)
	initiator1 := common.HexToAddress(initiator)

	id := c.PostForm("id")
	//string转int
	i_id, err := strconv.Atoi(id)
	//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.ReceiveApprove(client, contract, initiator1, big1)
	if err != nil {
		respError(c, err)
		fmt.Println("res:", err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 管理员拒绝资产领用申请，并在前台展示
func RejectApprove(c *gin.Context) {
	initiator := c.PostForm("initiator")
	fmt.Println("拒绝者账号:", initiator)
	initiator1 := common.HexToAddress(initiator)

	id := c.PostForm("id")
	//string转int
	i_id, err := strconv.Atoi(id)
	//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.RejectApprove(client, contract, initiator1, big1)
	if err != nil {
		respError(c, err)
		fmt.Println("res:", err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 查看合约某用户所有资产操作记录
func GetRecords(c *gin.Context) {

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetRecords(contract)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 查看所有资产库存情况
func GetAllGroup(c *gin.Context) {

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetAllGroup(contract)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

//
//// 查看合约某用户某个资产情况，并在前台展示  中间件
//func GetOneAssets(c *gin.Context) {
//	id := c.PostForm("id")
//
//	//string转int
//	i_id, err := strconv.Atoi(id)
//
//	//int先转uint64 再转*big.Int类型
//	big1 := new(big.Int).SetUint64(uint64(i_id))
//	initiator := c.PostForm("initiator")
//	initiator1 := common.HexToAddress(initiator)
//	client, err := mid.GetClient()
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	contract, err := mid.GetCallFaceTx(*client)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	res, err := mid.GetOneAssets(contract, big1, initiator1)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	respOK(c, res)
//	client.Close()
//}
//
//// 调用合约 新增管理员, 仅管理员调用 中间件
//func NewAddmin(c *gin.Context) {
//
//	initiator := c.PostForm("initiator")
//
//	add_i_initiator := common.HexToAddress(initiator)
//	client, err := mid.GetClient()
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	contract, err := mid.GetFacetTx(client)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	res, err := mid.NewAddmin(client, contract, add_i_initiator)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	respOK(c, res)
//	client.Close()
//}

func SaveAddress(c *gin.Context) {
	adminAddress := c.PostForm("adminAddress")

	// 获取 session
	session := sessions.Default(c)

	// 存储数据到 session
	session.Set("adminAddress", adminAddress)
	err := session.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从 session 中检索数据
	storedValue1 := session.Get("adminAddress")
	fmt.Println("获取的 session:", storedValue1)

	respOK(c, adminAddress)
}
func GetSaveAddress(c *gin.Context) {
	//adminAddress := c.PostForm("adminAddress")
	// 获取session
	session := sessions.Default(c)

	// 从session中检索数据
	storedValue := session.Get("adminAddress")

	if storedValue == nil {
		respError(c, "no key")
	} else {
		fmt.Println("成功获取保存的session:", storedValue)
		respOK(c, storedValue)
	}

}

// 新增用户资产  中间件
func NewAsset(c *gin.Context) {

	initiator := c.PostForm("initiator")
	fmt.Println("initiator", initiator)
	add_i_initiator := common.HexToAddress(initiator)
	name := c.PostForm("name")
	fmt.Println("name", name)
	description := c.PostForm("description")
	fmt.Println("description", description)
	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(contract)
	res, err := mid.NewAsset(client, contract, add_i_initiator, name, description)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	respOK(c, res)
	client.Close()
}

// 用户领用资产  中间件
func ReceiveAsset(c *gin.Context) {

	initiator := c.PostForm("initiator")
	fmt.Println("initiator", initiator)
	add_i_initiator := common.HexToAddress(initiator)
	id := c.PostForm("id")
	fmt.Println("id", id)
	//string转int
	i_id, err := strconv.Atoi(id)
	//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(contract)
	res, err := mid.ReceiveAsset(client, contract, add_i_initiator, big1)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	respOK(c, res)
	client.Close()
}

// 用户领用资产  中间件
func ReturnsAsset(c *gin.Context) {

	initiator := c.PostForm("initiator")
	fmt.Println("initiator", initiator)
	add_i_initiator := common.HexToAddress(initiator)
	id := c.PostForm("id")
	fmt.Println("id", id)
	//string转int
	i_id, err := strconv.Atoi(id)
	//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(contract)
	res, err := mid.ReturnsAsset(client, contract, add_i_initiator, big1)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	respOK(c, res)
	client.Close()
}

// 请求资产转让的函数，生成转让请求 中间件
func RequestAssetTransfer(c *gin.Context) {
	id := c.PostForm("id")
	//string转int
	i_id, err := strconv.Atoi(id)
	//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	initiator := c.PostForm("initiator")
	initiator1 := common.HexToAddress(initiator)

	to := c.PostForm("to")
	to1 := common.HexToAddress(to)

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(contract)
	res, err := mid.RequestAssetTransfer(client, contract, initiator1, big1, to1)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	respOK(c, res)
	client.Close()
}

// 批准资产转让的函数，仅允许管理员调用  中间件
func ApproveAssetTransfer(c *gin.Context) {
	initiator := c.PostForm("initiator")
	fmt.Println("批准者账号:", initiator)
	initiator1 := common.HexToAddress(initiator)

	id := c.PostForm("id")
	//string转int
	i_id, err := strconv.Atoi(id)
	//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.ApproveAssetTransfer(client, contract, initiator1, big1)
	if err != nil {
		respError(c, err)
		fmt.Println("res:", err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 拒绝资产转让的函数，仅允许管理员调用  中间件
func RefuseAssetTransfer(c *gin.Context) {
	initiator := c.PostForm("initiator")
	fmt.Println("拒绝者账号:", initiator)
	initiator1 := common.HexToAddress(initiator)

	id := c.PostForm("id")
	//string转int
	i_id, err := strconv.Atoi(id)
	//int先转uint64 再转*big.Int类型
	big1 := new(big.Int).SetUint64(uint64(i_id))

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetFacetTx(client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.RefuseAssetTransfer(client, contract, initiator1, big1)
	if err != nil {
		respError(c, err)
		fmt.Println("res:", err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 获取所有发起的资产转移请求 ,管理员查看  中间件
func GetAllRequests(c *gin.Context) {

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetAllRequests(contract)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

// 获取某用户发起的所有资产转移请求  中间件
func GetpersonRequests(c *gin.Context) {
	initiator := c.PostForm("initiator")
	add_i_initiator := common.HexToAddress(initiator)

	client, err := mid.GetClient()
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	contract, err := mid.GetCallFaceTx(*client)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	res, err := mid.GetpersonRequests(contract, add_i_initiator)
	if err != nil {
		respError(c, err)
		fmt.Println(err)
		return
	}
	respOK(c, res)
	client.Close()
}

//
//// 删除用户收藏的某个公益项目  中间件
//func DelInscollect(c *gin.Context) {
//
//	id := c.PostForm("id")
//	a, err := strconv.ParseInt(id, 10, 64)
//
//	g_goal := big.NewInt(a)
//
//	client, err := mid.GetClient()
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	contract, err := mid.GetFacetTx(client)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	res, err := mid.DelInscollect(client, contract, g_goal)
//	if err != nil {
//		respError(c, err)
//		fmt.Println(err)
//		return
//	}
//	respOK(c, res)
//	client.Close()
//}
