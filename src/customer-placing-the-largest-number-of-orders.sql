# Write your MySQL query statement below

select b customer_number from (select count(order_number) a, customer_number b  from Orders group by customer_number order by a desc limit 1) c