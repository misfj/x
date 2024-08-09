## 二进制部署
###
1.执行下面的3(make conf),修改数据库配置
###
2../main.exe

## 源码编译

###

1.安装make构建工具(linux 执行搜索 windows安装mingw和git,可查看https://go-sponge.com/quick-start当前链接进行教程安装)
### 
2.执行install-go.sh安装go环境(只适合linux,windows自行搜索)

### 
3.make conf 生成默认配置文件

###
4.修改配置文件的数据库即可,其他组件无需修改(当前版本无需其他组件)

### 
5.执行make run 程序就启动了

