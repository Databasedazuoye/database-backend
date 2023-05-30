# database-backend

## 启动项目
```
go mod tidy
go run main.go
```

在企业里，使用Gin框架来写后端时，项目结构通常会遵循MVC（Model-View-Controller）的设计模式。这种设计模式将应用程序分为三个部分：模型、视图和控制器。每个部分都有自己的职责，以实现代码的高内聚性和低耦合性。
下面是一个典型的Gin框架项目结构：
```
├── config
│   ├── config.go #配置文件读取
│   └── config.yaml #配置文件
├── controller #控制器层
│   ├── bill_controller.go #订单控制器
│   └── user_controller.go #用户控制器
├── middleware #中间件层
│   ├── auth_middleware.go #认证中间件
│   └── cors_middleware.go #跨域中间件
├── model #模型层
│   ├── bill_model.go #订单模型
│   └── user_model.go #用户模型
├── dao #数据访问层
│   ├── bill_dao.go #订单数据访问层
│   └── user_dao.go #用户数据访问层
├── router #路由层
│   ├── router.go #路由注册和初始化
├── service #服务层
│   ├── bill_service.go #订单服务层
│   └── user_service.go #用户服务层
├── util #工具类库
│   ├── jwt.go #JWT工具类库
├── main.go #应用程序入口文件，负责Gin框架的初始化和启动。
└── go.mod #Go模块文件，用于管理依赖项。
```

在本次项目中，为了简化架构，对项目做出了调整，以下为项目结构：
```
├── config
│   |── config.yml #配置文件
├── middleware #中间件层
│   ├── interceptor.go #拦截器 对请求携带的token进行校验 若token不合法会拒绝请求
├── model #模型层
│   ├── goods.go #商品模型
│   ├── user.go #用户模型
│   └── supplier.go #供应商模型
├── dao #数据访问层 一般来说 sql语句放在这一层写
│   ├── goodsDao.go #商品数据访问层
│   ├── supplierDao.go #供应商数据访问层
│   └── userDao.go #用户数据访问层
├── router #路由层
│   ├── router.go #路由注册和初始化 在这一层设置每一个url对应响应的服务
├── service #服务层
│   ├── goodsService.go #商品服务层
│   ├── supplierService.go #供应商服务层
│   └── userService.go #用户服务层
├── util #工具类库
│   ├── jwt.go #JWT工具类库
│   ├── cors.go #跨域工具类库
│   └── systemInit.go #系统启动工具类（连接数据库 读取yml配置文件加载到内存)
├── main.go #应用程序入口文件，负责Gin框架的初始化和启动。
└── go.mod #Go模块文件，用于管理依赖项。
```

接下来用goods模型来演示一遍写接口的过程
### goods
```
type Goods struct {
	Id              int  `xorm:"pk autoincr"`  //xorm是本次使用的数据库框架 pk设置Id属性为主键属性 autoincr设置主键自增
	Name            string
	Category        string
	Unit            string
	ManufactureDate string
	ExpirationDate  string
	PurchaseDate    string
	PurchasingPrice float64
	Num             int
}
```

### goodsDao
```
func GoodsInsert(goods model.Goods) int64 { //插入
	db := utils.GetDb() // 获取数据库连接
	goods.Id = 0 // 将商品的ID设为0，表示插入新的商品
	insert, err := db.Insert(goods) // 在数据库中插入商品
	if err != nil {
		panic(err) // 如果插入过程中出现错误，抛出异常
	}
	return insert // 返回插入操作影响的行数
}

func GoodsUpdate(goods model.Goods) int64 {
	db := utils.GetDb() // 获取数据库连接
	update, err := db.Id(goods.Id).Update(goods) // 根据商品的ID更新数据库中的商品信息
	if err != nil {
		panic(err) // 如果更新过程中出现错误，抛出异常
	}
	return update // 返回更新操作影响的行数
}

func GoodsDeleteById(id int64) int64 {
	db := utils.GetDb() // 获取数据库连接
	goods := model.Goods{} // 创建一个空的Goods对象
	i, err := db.Id(id).Delete(goods) // 根据商品的ID从数据库中删除商品
	if err != nil {
		panic(err) // 如果删除过程中出现错误，抛出异常
	}
	return i // 返回删除操作影响的行数
}

func GoodsSelectAll() []model.Goods {
	db := utils.GetDb() // 获取数据库连接
	goodsList := make([]model.Goods, 0) // 创建一个空的Goods对象切片
	db.Find(&goodsList) // 查询数据库中的所有商品，并将结果存储到goodsList切片中
	return goodsList // 返回查询到的所有商品
}
```

### goodsService
```
func GoodsInsert(c *gin.Context) {
	var goods model.Goods // 声明一个Goods结构体变量
	c.BindJSON(&goods) // 从请求中获取JSON数据并绑定到goods变量上
	insertNum := dao.GoodsInsert(goods) // 调用dao层的函数将商品信息插入到数据库中，并获取插入影响的行数
	if insertNum == 0 { // 如果插入影响的行数为0
		c.JSON(400, gin.H{ // 返回状态码400和错误消息JSON
			"msg": "插入失败",
		})
	} else {
		c.JSON(200, gin.H{ // 返回状态码200和成功消息JSON
			"msg": "插入成功",
		})
	}
}

func GoodsUpdate(c *gin.Context) {
	var goods model.Goods // 声明一个Goods结构体变量
	c.BindJSON(&goods) // 从请求中获取JSON数据并绑定到goods变量上
	updateNum := dao.GoodsUpdate(goods) // 调用dao层的函数更新数据库中的商品信息，并获取更新影响的行数
	if updateNum == 0 { // 如果更新影响的行数为0
		c.JSON(400, gin.H{ // 返回状态码400和错误消息JSON
			"msg": "更新失败",
		})
	} else {
		c.JSON(200, gin.H{ // 返回状态码200和成功消息JSON
			"msg": "更新成功",
		})
	}
}

func GoodsSelectAll(c *gin.Context) {
	all := dao.GoodsSelectAll() // 调用dao层的函数查询数据库中的所有商品信息，并将结果存储到all变量中
	for idx, item := range all { // 遍历all变量中的每个商品信息
		all[idx].ExpirationDate = strings.Split(item.ExpirationDate, "T")[0] // 将过期日期字段的时间部分去除，只保留日期部分
		all[idx].ManufactureDate = strings.Split(item.ManufactureDate, "T")[0] // 将生产日期字段的时间部分去除，只保留日期部分
		all[idx].PurchaseDate = strings.Split(item.PurchaseDate, "T")[0] // 将购买日期字段的时间部分去除，只保留日期部分
	}

	c.JSON(200, gin.H{ // 返回状态码200和包含所有商品信息的数据JSON
		"data": all,
	})
}

func GoodsDeleteById(c *gin.Context) {
	param := c.Param("id") // 获取URL中的id参数
	i, _ := strconv.ParseInt(param, 10, 64) // 将id参数转换为int64类型
	num := dao.GoodsDeleteById(i) // 调用dao层的函数根据商品ID从数据库中删除商品，并获取删除影响的行数
	if num == 0 { // 如果删除影响的行数为0
		c.JSON(400, gin.H{ // 返回状态码400和错误消息JSON
			"msg": "删除失败",
		})
	} else {
		c.JSON(200, gin.H{ // 返回状态码200和成功消息JSON
			"msg": "删除成功",
			})
	}
}
```

### router
```
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.Cors())
	r.Use(gin.Logger())


	goods := r.Group("/goods")
	{
		goods.POST("", service.GoodsInsert)
		goods.PUT("", service.GoodsUpdate)
		goods.DELETE("/:id", service.GoodsDeleteById)
		goods.GET("", service.GoodsSelectAll)
	}
```

本代码创建了一个路由分组 /goods，并且在该分组下面注册了四个 HTTP 请求处理函数。以下是对每个函数的解释：
```
goods.POST("", service.GoodsInsert)
```
作用：接受 POST 请求，将商品插入到数据库中；
参数：第一个参数 "" 表示路由路径为空，即 /goods；

```
goods.PUT("", service.GoodsUpdate)
```
作用：接受 PUT 请求，更新商品信息；
参数：第一个参数 "" 表示路由路径为空，即 /goods；

```
goods.DELETE("/:id", service.GoodsDeleteById)
```
作用： 接受 DELETE 请求，删除指定 ID 的商品；
参数：:id 是一个动态参数，表示请求路径中的商品 ID。例如，如果请求路径为 /goods/123，则 :id 的值为 "123"。service.GoodsDeleteById 是一个处理函数，用于通过 ID 删除商品；

```
goods.GET("", service.GoodsSelectAll)
```
作用：接受 GET 请求，查询所有商品信息；
参数：第一个参数 "" 表示路由路径为空，即 /goods；




