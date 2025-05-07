# zero-qp
# 适合棋牌 包网的项目框架 用的是go-zero框架 websocket 也写好了

## 项目结构介绍
### panguService pangu服务相关 只有rpc 里面只对接第三方的sms， 广告sdk 等
### userService 用户服务相关 只有rpc 也可以自己在userService 新建个api 做登录注册等
### hallService 游戏大厅相关 比如接入其他三方游戏 。只有rpc
### gameService 自己的游戏逻辑和三方游戏回调逻辑处理  api/rpc（api 只对三方游戏回调有用,如果不对接三方游戏就可以直接删掉）
### payService 充值服务 rpc/api（因为充值会不停的接入三四方充值，所以充值服务的api 独立起来 当然主要看你们自己）
### gateService 网关入口 注册或者公告类 只有api。
### wsService websocket 连接 只有api




### 启动顺序
#### 1.先启动 panguService/userService 里面的userrpc
#### 2.启动 gameService/hallService/payService 里面的rpc
#### 3.启动  gateService/wsService/callBackService 里面的rpc
#### 4.最后才启动 各个服务的api
#### 配置参考 user 


### 支付/三方数据回调思路

在api中  支付的回调 做一个区分代收代付的rpc 这个回调rpc 只做一个 就是就是组装好参数后发mq，rpc监听 异步处理回调成功还是失败  