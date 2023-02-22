# Write your MySQL query statement below

select distinct a.Email from Person a, Person b where a.Email = b.Email and a.Id != b.Id
