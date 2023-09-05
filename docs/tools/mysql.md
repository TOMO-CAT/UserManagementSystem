# mysql

## 安装

### 1. [不推荐] 在宿主机安装

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

### 2. [推荐] 启动 mysql 容器

```bash
# -d: 在后台运行容器
# --name mysql-server: 为容器指定一个名称
# --restart always: 退出后总是自动重启
# -e MYSQL_ROOT_PASSWORD=my-secret-pw: 设置 root 用户的密码
# -p 3306:3306: 映射 3306 端口
# mysql:5.7: 指定使用 mysql 容器镜像的版本
# docker run -d --name mysql-server --restart always -e MYSQL_ROOT_PASSWORD=cat123456 -p 3306:3306 mysql:5.7

docker run --name mysql-server --restart always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -v /home/vagrant/mysql5.7/data:/var/lib/mysql -d mysql:5.7


# 检查是否在运行中
$ docker container ls
CONTAINER ID   IMAGE                    COMMAND                  CREATED         STATUS         PORTS                                                  NAMES
de29e4b31061   mysql:5.7                "docker-entrypoint.s…"   2 minutes ago   Up 2 minutes   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   mysql-server

# 进入容器
$ docker exec -it de29e4b31061 /bin/bash
# [容器内] 查看 mysql 进程
bash-4.2# ps aux | grep mysql
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
mysql          1  0.2  1.2 1163032 198616 ?      Ssl  08:19   0:00 mysqld
# [容器内]

# 如果有问题的话可以查看日志
$ docker logs mysql-server
```

此时在宿主机就可以使用 mysql 了：

```bash
sudo apt install mysql-client
mysql -h 127.0.0.1 -P 3306 -u root -p
```
