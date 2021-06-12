# fabric-sdk-client

A simple wrapper for Hyperledger Fabric SDK

## Features

- [x] [admin](./client/admin.go)：提供 `GetAdminClient` 和 `admin.GetAppClient` 接口
- [x] [app](./client/app.go)：提供 `GetAppClient` 接口

## Todo 

进行更加细致的包装 (基于上面的两种 Client)，包括：

- [ ] 链码相关的包装 (打包、安装、批准、提交、初始化、调用、升级、修改背书策略等等)
- [ ] 通道相关的包装 (创建、加入、锚节点等等)
- [ ] MSP 管理相关操作 (增加 Peer、User 等)
- [ ] 操作账本的相关操作 (查询和解析 TX、Block 等)
