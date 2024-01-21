# Silent Car Auction 后端数据库函数添加指南

---

# 添加一个数据库中函数返回的临时表
1. 在model/function中添加函数返回表的实体  
    1.1 新建一个存储函数返回表各字段的结构体  
    1.2 为该结构体实现FunctionResult接口  
2. 在controller中添加路由  
    2.1 编写路由调用函数逻辑和路由注册函数  
    2.2 在main.go中注册路由  
3. 在repository/Functions.go中注册数据库函数  
    3.1 编写函数通过callFunction统一接口调用函数  