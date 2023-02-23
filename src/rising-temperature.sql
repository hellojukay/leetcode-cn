# Write your MySQL query statement below

select a.id from Weather a join Weather b on

a.recordDate  = AddDate(b.recordDate , INTERVAL 1 DAY) and a.temperature    > b.temperature   
