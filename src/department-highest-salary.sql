# Write your MySQL query statement below


select Department.name 'Department',Employee.name 'Employee', tmp.salary 'Salary' from Department join 
(select departmentId   ,max(salary ) salary from Employee group by departmentId) tmp 
on Department.Id = tmp.departmentId join Employee  on 
Employee.salary  = tmp.salary and Employee.departmentId = tmp.departmentId