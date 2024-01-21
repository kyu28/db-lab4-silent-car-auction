# Silent Car Auction Server
## 数据库系统 实验4 选题6 后端

---

# 语言和技术栈
Go - 语言  
Gin - 后端框架  
pgx - PostgreSQL驱动  

---

# 文件结构
```
controller/ - 路由控制逻辑层
    utils.go - 工具函数
docs/ - 文档
model/ - 数据库实体层
    function/ - 数据库函数返回表实体层
        FunctionResult.go - 数据库函数返回表实体需实现的统一接口
    Table.go - 数据库实体需实现的统一接口
repository/ - 数据持久化仓库层
    CRUDRepository.go - 数据库表增删改查统一接口
    ProcedureRepository.go - 调用存储过程的统一接口
    FunctionRepository.go - 调用函数的统一接口
    Procedure.go - 注册的数据库存储过程
    Functions.go - 注册的数据库函数
main.go - 服务器启动入口
```

---

# 各包注释及函数注释
使用`go doc <包名/函数名/结构体名/接口名>`指令查看详细文档  