# AGGRA FORUM

一款基于Go Iris Gorm 的轻量级论坛

## 项目依赖

1. Mysql 5.7
2. Gorm
3. Redis

## 配置
```
conf/db.go // 数据库配置
conf/sysconf.go // 系统配置
```

## 编译 & 运行

```bash
go build web/main2.go

./main2
```

## Features

1. 登录/注册
1. 发帖/回帖/收藏/关注
1. 个人资料修改
1. 消息中心
1. 后台管理
