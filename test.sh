#!/bin/bash
export GO111MODULE=on
# 可以启动多个service
go run srv/main.go &
# 客户端
go run cli/main.go
