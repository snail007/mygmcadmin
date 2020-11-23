# 说明

基于 [GMC](https://github.com/snail007/gmc) 框架的开源管理后台.

# 使用须知

1. 安装好 [GMCT](https://snail007.github.io/gmc/zh/#/?id=gmct-%e5%b7%a5%e5%85%b7%e9%93%be) 工具链。
1. go版本要求go1.12及以上，系统配置好 `$GOPATH` 环境变量，命令行可以直接执行 `go` 命令。
1. 后台使用MySQL数据库，需要你建立一个数据库，然后导入数据文件,位于：`doc/db.sql`.
1. 修改配置文件 `conf/app.tml` 配置里面的MySQL数据库连接信息.
1. 在项目目录执行，`gmct run` ， 浏览器访问 `http://127.0.0.1:7080` 即可看到效果。

