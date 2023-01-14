# GBlog博客后端(Golang)

## 快速开始

- 1.推荐 `Golang 18版本`
- 2.在 `config`目录下创建`config.toml`文件
- 3.配置 `config/config.toml` 配置信息
- 4.执行 `go build main.go`
- 5.执行 `./main`
- 6.在数据库user表里新增一条记录,作为管理员账号

## 配置config.toml文件如下
```
[Server]
Mode = "debug"
Port = "9000"
[Mysql]
DBName = "数据库名"
DBPort = "数据库端口"
DBUser = "数据库用户名"
DBHost = "数据库主机地址"
DBPassWord = "数据库密码"
[Jwt]
JwtSecret = "Jwt秘钥"
JwtAuthUser = "Jwt授权用户名"
[Redis]
RedisHost = "Redis主机"
RedisPort = "Redis端口"
RedisPassWord = "Redis密码"
```

