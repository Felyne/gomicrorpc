#### 基于go-micro和etcd服务发现的demo

#### 准备
1.启动etcd，这里是`localhost:2379`

2.利用[管理工具](https://github.com/Felyne/admin_tool)上传服务端的配置
```shell
./admin_tool config set dev SAY_SERVICE ./config/server.conf localhost:2379
```

#### 服务端
```shell

bash proto/gen.sh

cd srv
make
./srv dev 0 localhost:2379
```

#### 客户端
```shell
cd cli
go build
./cli dev localhost:2379
```

#### 参考文档
  - [gRPC 官方文档中文版](https://doc.oschina.net/grpc)
  - [Protobuf3语言指南](https://blog.csdn.net/u011518120/article/details/54604615)