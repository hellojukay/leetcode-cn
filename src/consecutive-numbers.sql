# Write your MySQL query statement below

select distinct a.Num as ConsecutiveNums  from Logs a join Logs b on
a.Num = b.Num and  a.Id = b.Id -1 join Logs c on 
c.Num = b.Num and b.Id = c.Id - 1