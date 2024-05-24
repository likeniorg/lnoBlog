#!/bin/bash
echo "正在设定数据库密码，请输入"
read dbpass
yum -y update
yum remove docker  docker-common docker-selinux docker-engine -y
yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-manager --add-repo http://download.docker.com/linux/centos/docker-ce.repo
yum -y install docker-ce
systemctl start docker
systemctl enable docker
docker run --detach --name lnodb -p3306:3306 --env MARIADB_ROOT_PASSWORD=$dbpass  mariadb:latest
  

MYSQL_HOST="127.0.0.1"
MYSQL_DATABASE="lnoblog"
MYSQL_USER="root"
MYSQL_PASSWORD=$dbpass
SQL_FILE="install/init.sql"

mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD -e "create database lnoblog;" >/dev/null 2>&1

# 尝试连接到 MySQL
mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE -e "SELECT 1;" >/dev/null 2>&1

# 检查连接是否成功
if [ $? -eq 0 ]; then
  echo "连接到数据库成功，正在导入数据库"
  
  # 导入数据
  mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE < $SQL_FILE
  
  echo "数据库导入成功"
else
  echo "连接到数据库失败"
fi

go build .
./lnoblog