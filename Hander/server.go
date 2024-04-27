package Hander

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/ugorji/go"
	"net/http"
)

// Index 前端登入注册页面
func Index() {
	engine := gin.Default()                  //返回默认引擎
	engine.Static("/static", "views/static") //加载静态文件
	engine.LoadHTMLGlob("views/html/*")      //加载view/html/所有的html文件
	engine.Use(cors.Default())               //跨域 插件

	// 创建 Cookie 存储
	store := cookie.NewStore([]byte("your-secret-key"))

	// 初始化 session 中间件
	engine.Use(sessions.Sessions("my-session", store))

	//注册页面
	engine.GET("/register", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "register.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件
	//登入页面
	engine.GET("/login", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "login.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件

	//管理页面
	engine.GET("/manage", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "manage.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件
	//书籍主页
	engine.GET("/index", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "index.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件
	engine.GET("/table", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "tables.html", gin.H{"title": "资产后台管理"})
	}) //触发get请求 加载html文件
	engine.GET("/assetManage", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "AssetManage.html", gin.H{"title": "资产出库入库"})
	}) //触发get请求 加载html文件
	engine.GET("/assetLedger", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "AssetLedger.html", gin.H{"title": "资产出库入库"})
	}) //触发get请求 加载html文件
	engine.GET("/b", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "profile.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件

	engine.GET("/c", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "table.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件
	engine.GET("/e", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "project_classify.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件
	engine.GET("/f", func(c *gin.Context) {
		//根据文件名渲染
		//加载模板是加载的路径，替换的是文件中的某个变量
		c.HTML(200, "inscollect.html", gin.H{"title": "标题"})
	}) //触发get请求 加载html文件

	dai := engine.Group("dai") //创建 请求组
	{
		dai.POST("/saveAddress", SaveAddress)
		dai.GET("/getSaveAddress", GetSaveAddress)
		dai.POST("/login", Login)       //POST>>>>登入
		dai.POST("/register", Register) //POST>>>>注册
		//dai.POST("/countAmount", CountAmount)   //POST>>>>连接合约获得合约余额
		dai.POST("/getallAssets", GetallAssets) //POST>>>>连接合约获得所有公益项目  (已改)
		dai.POST("/getAssetReceive", GetAssetReceive)
		dai.POST("/receiveApprove", ReceiveApprove)
		dai.POST("/rejectApprove", RejectApprove)

		dai.POST("/getRecords", GetRecords) //POST>>>>连接合约获得所有公益项目  (已改)

		dai.POST("/getAllGroup", GetAllGroup)
		//dai.POST("/getOneAssets", GetOneAssets) //POST>>>>连接合约获得某个公益项目 (已改)
		//dai.POST("/contribute", Contribute)             //POST>>>>连接合约 用户捐助
		dai.POST("/newAsset", NewAsset) //POST>>>>连接合约  发布一个新项目  (已改)
		dai.POST("/receiveAsset", ReceiveAsset)
		dai.POST("/returnsAsset", ReturnsAsset)
		//dai.POST("/newAddmin", NewAddmin)                       //POST>>>>连接合约  发布一个新项目  (已改)
		dai.POST("/requestAssetTransfer", RequestAssetTransfer) //POST>>>>连接合约  删除一个项目 (已改)
		dai.POST("/approveAssetTransfer", ApproveAssetTransfer) //POST>>>>连接合约  把一个项目加入某个用户的收藏 (已改)
		dai.POST("/refuseAssetTransfer", RefuseAssetTransfer)
		dai.POST("/getAllRequests", GetAllRequests)       //POST>>>>连接合约  把一个项目加入某个用户的收藏 (已改)
		dai.POST("/getpersonRequests", GetpersonRequests) //POST>>>>连接合约  把一个项目加入某个用户的收藏
		//dai.POST("/delInscollect", DelInscollect)               //POST>>>>连接合约  把一个项目加入某个用户的收藏
		dai.POST("/getUser", GetUser)
		dai.POST("/changeUser", ChangeUser)

	}
	err := engine.Run("localhost:8080") //启动端口127.0.0.1:8080
	if err != nil {
		return
	} //处理错误机制
}
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}
func tohtml(c *gin.Context, data interface{}) { //渲染html 返回
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}
func tohtml1(c *gin.Context, data interface{}, data1 interface{}) { //渲染html 返回
	c.JSON(200, gin.H{
		"code":  0,
		"data":  data,
		"data1": data1,
	})
}

// 封装函数，用于向客户端返回正确消息
func respOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}
