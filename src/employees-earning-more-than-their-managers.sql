# Write your MySQL query statement below

select a.name Employee  from Employee a  join Employee b
  on 
  a.managerId  = b.id and a.salary  > b.salary 