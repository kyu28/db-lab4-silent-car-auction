# Silent Car Auction 后端表添加指南

---

# 添加一张数据库中存在的表
1. 在model中添加实体  
    1.1 新建一个存储表中各字段的结构体  
    1.2 为该结构体实现Table接口  
2. 在controller中添加路由  
    2.1 编写路由逻辑和路由注册函数  
    2.2 在main.go中注册路由  