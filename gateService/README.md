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


# API 规范
1. 根目录`cd api`
2. 创建api的模板`gateService.api`
3. 生成api文件夹路径
```
goctl api go -api gateService.api -dir . --style goZero
```
也可以根目录下，使用makefile命令
```
make goctl-api
```


# idea配置debug
1. mac的idea
2. 点击右上角的add configuration打开界面
3. 点击左上角"+"
4. 选择go build

## 以api为例（rpc一样的）

5. working directory 项目的绝对路径
   eg:
   working directory: 绝对路径../gateservice/api
6. 点击 play 或者 debug, 及运行