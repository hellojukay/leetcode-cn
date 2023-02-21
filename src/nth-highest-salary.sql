CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  set N=N - 1;
  RETURN (
      # Write your MySQL query statement below.
      select distinct(salary) as a from Employee order by a desc limit N, 1
  );
END