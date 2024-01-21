# Silent Car Auction Server

# 技术栈
Gin  

# Installation
```
go get github.com/jackc/pgx  
go get github.com/gin-contrib/cors  
go get github.com/gin-gonic/gin  
go build -o silent-car-auction-server  
```

# Run
`./silent-car-auction-server [port]`  

# API
/表名 GET 参数为筛选条件的表单 SELECT  
/表名 POST 参数为新行的数据的表单 INSERT  
/表名/列名/值 PUT 参数为筛选条件的表单 UPDATE  
/表名 DELETE 参数为筛选条件的表单 DELETE  
  
示例：  
`/auto` GET即`SELECT * FROM tb_auto;`  

`/auto` POST  
数据`VIN=1FMDU74W02UB72009&location=Z095&year=2002&type_of_auto=FORD EXPLORER EB 4X4 8A ASBW&mileage=82`  
即`INSERT INTO tb_auto (VIN, location, type_of_auto, mileage, year) VALUES ('1FMDU74W02UB72009', 'Z095', 'FORD EXPLORER EB 4X4 8A ASBW', 82, TO_DATE('2002', 'YYYY'));`  

`/auto/mileage/114` PUT  
数据`VIN=1FMDU74W02UB72009`  
即`UPDATE tb_auto SET mileage=114 WHERE VIN='1FMDU74W02UB72009';`  

`/auto` DELETE  
数据`VIN=1FMDU74W02UB72009`  
即`DELETE FROM tb_auto WHERE VIN='1FMDU74W02UB72009';`  

# 新表创建
在model中创建新实体实现Table接口  
在controller中创建新controller实现路由逻辑  
在main.go中注册新的controller  