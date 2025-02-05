# MySQL

```
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
