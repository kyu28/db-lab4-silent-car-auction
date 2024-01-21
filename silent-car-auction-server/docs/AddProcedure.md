# Silent Car Auction 后端数据库存储过程添加指南

---

# 添加一个存储过程
1. 在controller中添加路由  
    1.1 编写路由调用存储过程逻辑和路由注册函数  
    1.2 在main.go中注册路由  
2. 在repository/Procedures.go中注册数据库存储过程  
    2.1 编写函数通过callProcedure统一接口调用函数  