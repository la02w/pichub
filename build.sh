#!/bin/bash

# 构建Linux版本
GOOS=linux GOARCH=amd64 go build -o pichub

# 构建Windows版本
GOOS=windows GOARCH=amd64 go build -o pichub.exe

# 如果需要其他架构或操作系统，可以继续添加
