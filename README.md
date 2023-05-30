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
