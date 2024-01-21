# Silent Car Auction 前端数据库函数添加指南

---

# 添加一个数据库中函数返回的临时表
1. 在main.js中的const lists对象中添加新的键值对  
```
"your-function-name": {
    table: "yourServerFunctionRoute",
    cols: [ "col1", "col2" ],
    url: "http://example.com/your-table/your-func",
    hasArgs: true
}
```
2. 在App.vue中添加菜单项
```
<el-sub-menu>
    <!-- Add new menu item under any el-sub-menu -->
    <el-menu-item index="yourIndex">
        <RouterLink to="/table/new-function">New function</RouterLink>
    </el-menu-item>
</el-sub-menu>
```