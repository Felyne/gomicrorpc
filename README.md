### 基于go-micro和etcd服务中心的demo

#### 准备
1.安装以下工具  
```
go get github.com/golang/protobuf/{proto,protoc-gen-go}
go get github.com/micro/protoc-gen-micro
go get github.com/Felyne/admin-tool
```
2.启动etcd，这里是`localhost:2379`  

3.上传配置  
```shell
# say是proto文件定义的服务名
./admin-tool config set dev say ./config/server.toml localhost:2379
```

#### 服务端
```shell
cd proto && bash gen.sh
cd ../
cd srv && make
./say dev 0 localhost:2379
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
  - [go-micro文档](https://micro.mu/docs/)