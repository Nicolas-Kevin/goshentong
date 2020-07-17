# goshentong

## 介绍

Golang Shentong database driver conforming to the Go database/sql interface

## 安装
go get github.com/wb253/goshentong

新建文件 aci.pc, 内容

prefix=/usr
includedir=/opt/ShenTong/aci/include
libdir=/opt/ShenTong/aci/linux64

Name: aci
Description: Shentong Client
Version: 1.0
Cflags: -I${includedir}
Libs: -L${libdir} -laci


设置环境变量
export LD_LIBRARY_PATH=/opt/ShenTong/aci/linux64
export PKG_CONFIG_PATH=api.pc 所在目录


