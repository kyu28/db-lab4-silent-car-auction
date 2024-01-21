# Silent Car Auction 前端表添加指南

---

# 添加一张数据库中存在的表
1. 在main.js中的const tables对象中添加新的键值对  
```
"your-table-name": {
    table: "yourServerTableRoute",
    immutableCols: [ "auto_increment_col", "foreign_reference_col" ],
    mutableCols: [ "col1", "col2" ],
    baseURL: "http://example.com"
}
```
2. 在App.vue中的tableRoutes中添加列表信息对象
```
const tableRoutes = ref([
...
  {
    "display": "Your Table Name",
    "path": "your-table-name"
  }
]);
```
