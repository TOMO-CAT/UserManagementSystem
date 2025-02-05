# UserManagerSystem

## 功能

 简单但完整的 golang 后端项目框架

* 常用的 rpc 框架，包括 grpc、http 和 thrift
* 监控报警
* mysql 数据库的简单使用
* redis 缓存的简单使用
* 基于 http 的基础前端搭建方法
* 基于 vue 的简单前端搭建方法（[https://www.gin-vue-admin.com](https://www.gin-vue-admin.com)）
* 单元测试写法
* 性能测试写法
* 压测工具
* docker 镜像使用方法
* 基于动态配置更新的降级方案

## TODO

> [https://www.kancloud.cn/sliver_horn/gorm/1861152](https://www.kancloud.cn/sliver_horn/gorm/1861152)

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

# redis 和 mysql 数据库
docker-compose up --build -d

# 运行 docker 容器
bash docker.sh run
```

## 编译

```bash
go mod tidy
go mod vendor

```
