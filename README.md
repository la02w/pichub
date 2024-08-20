# PicHub 个人图床

本仓库基本代码来自[kaf.im](https://kaf.im)，使用 GO 语言还原。

> 代办

- [x] 创建配置文件，一键配置
- [ ] 创建数据库保存图片数据，能按日期展示图片
- [x] 复制链接时，可选复制单链接、Markdown 链接和 HTML 链接
- [x] 第三方存储库配置-只添加了腾讯 COS 对象存储

## 手动使用 make 打包

```bash
# 安装编译打包依赖
# ubuntu
sudo apt install build-essential zip
# 编译打包
make build
# 打包后的文件位于build/xxx.tar.gz    build/xxx.zip
# 清除build文件夹，也可以使用rm -rf build/
make clean
```

## 使用方法

进入 config/ 配置 config.ini 文件

| 配置项                | 描述             | 值                |
| --------------------- | ---------------- | ----------------- |
| PICK_SERVICE          | 选择服务项       | "local"&"tencent" |
| SERVER_PORT           | 服务端口         | ":2356"           |
| LOCAL_BASE_FOLDER     | 本地路径         | "./upload/"       |
| TENCENT_COS_URL       | 用户存储桶和地区 | null              |
| TENCENT_COS_SECRETID  | SECRETID         | null              |
| TENCENT_COS_SECRETKEY | SECRETKEY        | null              |
