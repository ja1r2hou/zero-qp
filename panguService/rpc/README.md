# readme
# 文件夹说明
+ `go.mod`
+ `rpc` rpc服务
+ `api` api http调用
+ `cronjob` task任务
+ `script` 脚本
+ `rmp` 队列
+ `model` 数据库

文件夹tree说明
```
xxx-service
|____go.mod
|____README.md
|____rpc
| |____xxxxx
|____api
| |____xxxxx
|model
| |____xxxxx
|cronjob
| |____xxxxx
|script
| |____xxxxx
|rmp
| |____xxxxx
```

# idea配置
## 自动引用mac系统的mod库
1. 点击菜单栏 Goland
2. Preferences
3. Go -> Go Modules
4. 打钩✅Enable Go modules integration
5. 配置GOPROXY 选择 direct保存


# RPC 规范
1. 根目录`cd rpc`
2. 生成模板
```
goctl rpc protoc panguService.proto --style goZero --go_out=. --go-grpc_out=. --zrpc_out=.
```
3. 编辑`panguService.proto`定义
    1. package 规范
    ```
    # 使用全小写
    package buybuybuycrpc
    ```
    2. service 规范
    ```
    service BuybuybuyRPC {}
    ```
4. 生成rpc文件夹路径
```
goctl rpc proto -src panguService.proto -dir . -style goZero
```
也可以根目录下，使用makefile命令
```
make goctl-rpc
```

# 数据库规范
1. 根目录`cd model`
```
goctl model mysql datasource -url="account:password(db_ip:3306)/db_name" -table="table_name"  -dir="." -style goZero 
```
+ 生成对应的table_name的数据库文件.go
+ 切记代码风格`-style goZero`

----

# 文件夹说明
无论是api/rpc生成都有对应的internal文件夹
## etc配置
1. `./etc/xxxService.yaml` 配置数据库,redis,mq位置信息
## 内部的文件夹管理
1. `./internal/config/config.go` 配置数据库，redis，其他外部组件信息
2. `./internal/svc/serviceContext.go` 配置数据库关系，外部rpc调用关系
3. `./internal/logic` 编写代码逻辑


# idea配置debug
1. mac的idea
2. 点击右上角的add configuration打开界面
3. 点击左上角"+"
4. 选择go build

## 以api为例（rpc一样的）

5. working directory 项目的绝对路径
   eg:
   working directory: 绝对路径../user-service/api
6. 点击 play 或者 debug, 及运行