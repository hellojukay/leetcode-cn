select ifNull((select  a.salary   from Employee  a where a.salary < (select max(salary) from Employee) order by a.salary desc limit 1),null) as SecondHighestSalary