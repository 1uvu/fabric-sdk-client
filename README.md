# fabric-sdk-client

A simple wrapper for Hyperledger Fabric SDK

## Features

- [x] [pkg/client](./pkg/client)：提供 App 和 Admin Client 的并发单例访问功能
- [x] [pkg/config](./pkg/config)：解析 Client 配置文件
- [x] [pkg/sdk](./pkg/sdk)：对 Fabric-SDK Client 的包装
- [x] [pkg/types](./pkg/types)：类型定义

## Todo

进行更加细致的包装 (基于上面的两种 Client)，包括：

- [ ] 链码相关的包装 (打包、安装、批准、提交、初始化、调用、升级、修改背书策略等等)
- [ ] 通道相关的包装 (创建、加入、锚节点等等)
- [ ] MSP 管理相关操作 (增加 Peer、User 等)
- [ ] 操作账本的相关操作 (查询和解析 TX、Block 等)
