# redis

## 安装

### 1. [不推荐] 在宿主机安装

> docker 中安装 systemctl 命令比较麻烦，因为 docker 的设计理念是一个容器只运行一个服务，我们暂时在宿主机中安装 redis。

```bash
sudo apt install redis-server

sudo systemctl start redis-server

sudo systemctl status redis-server
```

### 2. [推荐] 启动 redis 容器

```bash
# 后台启动 redis server container
# @note:
#   --restart always: 退出后总是自动重启
#   ---p 6379:6379: 映射 redis 端口
$ docker run -d --name redis-server --restart always -p 6379:6379 redis


# 检查是否在运行中
$ docker container ls
CONTAINER ID   IMAGE                    COMMAND                  CREATED              STATUS              PORTS      NAMES
855464acca89   redis                    "docker-entrypoint.s…"   About a minute ago   Up About a minute   6379/tcp   redis-server

# 由于是 `--net host` 模式, 因此不需要创建 docker 网络
# # 创建一个 docker 网络
# $ docker network create cat-network

# # 将我们的 dispatch-engine container 连接到这个共享网络
# $ docker network connect cat-network cat_dispatch-engine
```
