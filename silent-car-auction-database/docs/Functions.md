# Silent Car Auction Database 函数说明

---

available_fords() - 返回所有FORD牌汽车列表  
available_cheverolets() - 返回所有CHEVEROLET牌汽车列表  
select_by_brand(brand VARCHAR(128)) - 返回给定品牌的汽车列表  
max_bid() - 返回各车辆最高出价表  
winners_and_losers(p_auto_id INT) - 根据给定车辆ID返回赢家和输家列表  
  
set_winners() - 触发器用函数，将上次出价长于1月的赢家加入赢家表，表示拍出  