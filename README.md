### 基于go-micro和etcd服务中心的demo

#### 准备
1.启动etcd，这里是`localhost:2379`

2.利用[管理工具](https://github.com/Felyne/admin_tool)上传服务端的配置
```shell
./admin_tool config set dev SAY_SERVICE ./config/server.conf localhost:2379
```

#### 服务端
```shell

# 先装好grpc用到的东西
bash proto/gen.sh
cd srv && make
./srv dev 0 localhost:2379
```

#### 客户端
```shell
cd cli
go build
./cli dev localhost:2379
```

#### 参考文档
  - [etcd官方文档](https://etcd.io/docs/v3.4.0/)
  - [gRPC官方文档中文版](https://doc.oschina.net/grpc)
  - [Protobuf3语言指南](https://blog.csdn.net/u011518120/article/details/54604615)