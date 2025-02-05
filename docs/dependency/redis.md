# Redis

docker 中安装 systemctl 命令比较麻烦，因为 docker 的设计理念是一个容器只运行一个服务，我们暂时在宿主机中安装 redis 和 mysql。

```
sudo apt install redis-server

sudo systemctl start redis-server

sudo systemctl status redis-server
```
