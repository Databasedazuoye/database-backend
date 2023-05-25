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
