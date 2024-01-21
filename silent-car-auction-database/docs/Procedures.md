# Silent Car Auction Database 存储过程说明

---

add_customer(first_name VARCHAR(30), last_name VARCHAR(30), location CHAR(4), email VARCHAR(64)) - 添加一个客户，即注册  
add_auto(VIN CHAR(17), location CHAR(4), type_of_auto VARCHAR(128), mileage INT, year DATE) - 添加一辆车  
add_transaction(p_auto_id INT, p_customer_id INT, p_bid DECIMAL(16, 2), p_bid_date DATE) - 添加一笔交易，即下单  