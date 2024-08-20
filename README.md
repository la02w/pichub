# PicHub 个人图床

本仓库基本代码来自[kaf.im](https://kaf.im)，使用 GO 语言还原。

> 代办

- [x] 创建配置文件，一键配置
- [ ] 创建数据库保存图片数据，能按日期展示图片
- [ ] 复制链接时，可选复制单链接、Markdown 链接和 HTML 链接
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
