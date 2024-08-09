#!/bin/bash

# 定义Go的版本和安装路径
GO_VERSION="1.18.1" # 请根据需要替换为最新版本
GO_ARCH="amd64" # 根据你的系统架构替换，例如：386, arm, arm64 等
GO_TARBALL="go${GO_VERSION}.linux-${GO_ARCH}.tar.gz"
GO_URL="https://golang.org/dl/${GO_TARBALL}"
GO_PATH="/usr/local/go"

# 检查脚本是否以root权限运行
if [ "$EUID" -ne 0 ]; then
  echo "请以root权限运行此脚本"
  exit 1
fi

# 检查系统类型并下载相应的Go版本
if [ "$(uname -m)" = "x86_64" ]; then
  GO_ARCH="amd64"
elif [ "$(uname -m)" = "aarch64" ]; then
  GO_ARCH="arm64"
else
  echo "不支持的架构"
  exit 1
fi

# 下载Go
echo "正在下载Go ${GO_VERSION}..."
wget -q "$GO_URL" -O "$GO_TARBALL"

# 检查下载是否成功
if [ $? -ne 0 ]; then
  echo "下载失败，请检查网络连接或URL是否正确"
  exit 1
fi

# 解压Go
echo "正在解压Go到 ${GO_PATH}..."
tar -C /usr/local -xzf "$GO_TARBALL"

# 设置环境变量
echo "正在设置环境变量..."
echo "export PATH=\$PATH:$GO_PATH/bin" > /etc/profile.d/go.sh

# 清理
rm -f "$GO_TARBALL"

# 验证安装
if go version | grep -q "go version go${GO_VERSION}"; then
  echo "Go ${GO_VERSION} 安装成功"
else
  echo "Go 安装失败"
  exit 1
fi