## 二进制部署
###
1.执行下面的3(make conf),修改数据库配置
###
2../main.exe

## 源码编译(推荐方式)

###
1.安装make构建工具(linux[ubutun:sudo apt-get install make centos:yum install make ] 
windows安装mingw和git,可查看https://go-sponge.com/quick-start当前链接进行教程安装)😀

### 
2.执行install-go.sh(给文件加上可执行权限,chmod +x install-go.sh) 安装go环境(只适合linux,windows自行搜索)

### 
3.make conf 生成默认配置文件

### 
4.make swagger 生成swagger文档

###
4.修改配置文件的数据库即可,其他组件无需修改(当前版本无需其他组件)

### 
5.执行make run 程序就启动了

###
6.打开浏览器 访问 http://localhost:8080/swagger/index.html  查看swagger文档

###
7.注意, 使用make clean 的时候,会删除本地存储的文件,造成数据库记录的数据被删除的现象,如果不想要删除,请修改下makefile(rm  -rf ./encrypt-files
rm  -rf ./uploads)把这两行注释掉