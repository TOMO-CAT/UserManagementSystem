# UserManagerSystem

## 功能

基于 Golang 的用户管理系统，面试必备项目。在这个项目中你可以学到：

* 简单但完整的 golang 后端项目框架
* 常用的 rpc 框架，包括 grpc、http 和 thrift
* 监控报警
* mysql 数据库的简单使用
* redis 缓存的简单使用
* 基于 http 的基础前端搭建方法
* 基于 vue 的简单前端搭建方法（<https://www.gin-vue-admin.com>）
* 单元测试写法
* 性能测试写法
* 压测工具
* docker 镜像使用方法
* 基于动态配置更新的降级方案

## TODO

> <https://www.kancloud.cn/sliver_horn/gorm/1861152>

基于gin+vue搭建的后台管理系统框架，集成:

* jwt鉴权
* 权限管理
* 动态路由
* 分页封装
* 多点登录拦截
* 资源权限
* 上传下载
* 代码生成器
* 表单生成器等基础功能
* 五分钟一套CURD前后端代码

## 搭建 docker 环境

```bash
# 搭建 docker 容器
bash docker.sh build

# 运行 docker 容器
bash docker.sh run
```

## 依赖

### 1. protoc

已经在 Dockerfile 中增加了，如果使用 docker 的话可以跳过这一步：

```bash
sudo apt install protobuf-compiler
```

还需要安装 `protoc-gen-go`，用于生成 `xx.pb.go` 文件：

```bash
# 网络不通时使用代理:
# export GOPROXY=https://proxy.golang.com.cn,direct

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

再安装 `protoc-gen-go-grpc` 插件，用于生成 `xx_grpc.pb.go` 文件：

```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

此时这两个插件会安装在 `~/go/bin` 下，我们可以在 `~/.bashrc` 中加入一行：

```bash
export PATH=~/go/bin:$PATH
```

检查插件是否安装成功：

```bash
$ protoc-gen-go --version
protoc-gen-go v1.31.0
```

### 2. redis

> docker 中安装 systemctl 命令比较麻烦，因为 docker 的设计理念是一个容器只运行一个服务，我们暂时在宿主机中安装 redis 和 mysql。

```bash
sudo apt install redis-server

sudo systemctl start redis-server

sudo systemctl status redis-server
```

### 3. MySql

```bash
# 安装
sudo apt install mysql-server

# 启动 MySql 服务
sudo systemctl start mysql

# 查看启动状态
sudo systemctl status mysql

# 查看默认账户密码
cd /etc/mysql
sudo cat debian.cnf

# 使用默认账户登陆 mysql
# 密码: W80yHnUinnJBHTlQ
mysql -u debian-sys-maint -p


# 修改 root 密码 (MySQL 8.0 版本以上): https://techvblogs.com/blog/how-to-reset-mysql-root-password-ubuntu
# 1. 停止服务
sudo systemctl stop mysql.service
# 2. 跳过权限验证
sudo systemctl set-environment MYSQLD_OPTS="--skip-networking --skip-grant-tables"
# 3. 重启服务
sudo systemctl start mysql.service
# 4. 登陆
sudo mysql -u root
# 5. 刷新权限表
mysql> flush privileges;
# 6. 修改密码
mysql> use mysql;
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '12345678';
# 7. 退出
mysql> quit;
# 8. 重启服务
sudo killall -u mysql
sudo systemctl restart mysql.service

# 登陆 root 账号 (必须带上 sudo)
sudo mysql -u root -p

# 创建数据库
mysql> CREATE DATABASE db_ums;
```

## 编译

```bash
go mod tidy
go mod vendor

```
