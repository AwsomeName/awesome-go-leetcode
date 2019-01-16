SELECT max(Salary) as SecondHighestSalary
FROM Employee
WHERE Salary < (SELECT max(Salary) FROM Employee)

select (
  select distinct Salary from Employee order by Salary Desc limit 1 offset 1
)as SecondHighestSalary
