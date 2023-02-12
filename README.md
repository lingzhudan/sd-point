# 自律点
采用go-kratos搭建的后端系统 对外接口提供http 对内服务使用grpc调用

目前正在开发中 功能尚不完善

## 使用技术

### 服务发现注册
etcd
### 参数校验
github.com/envoyproxy/protoc-gen-validate
### 认证
内部服务调用grpc采用jwt认证

外部服务调用计划采用sessionID方式
### 日志
kratos/log
### 监控
待添加
### 链路追踪
待添加
### 熔断机制
待添加
### 限流器
待添加
### 负载均衡
待添加
### 错误处理
待添加
### 运维
计划采用docker+k8s部署微服务

## 待优化事项

### （一）/sd-point/app_makefile
makefile中docker部分内容待更新

### （二）/sd-point/app/sd-point/interface/internal/service/sd-point_interface
service实现方法待实现

## 项目初始化
```
# 先安装go环境后再执行此命令
make init
```

## 常用快捷命令
```
# 新增proto模板
kratos proto add api/xxx{服务名}/v1{版本}/xxx.proto
# 生成pb.go文件
kratos proto client api/xxx{服务名}/v1{版本}/xxx.proto 
    # 或
make api

# 生成service文件
kratos proto server api/xxx{服务名}/v1{版本}/xxx.proto -t \
 app/xxx{服务名}/internal/service

# 生成二进制执行文件
cd app/xxx{服务名}/service && make all
# 运行
./bin/server -conf ./configs
```

## Docker 此项待更新 不具有效性
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

