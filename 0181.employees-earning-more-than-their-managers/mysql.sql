select A.Name as Employee from Employee A 
where A.Salary > (select Salary from Employee B where B.id = A.ManagerId)
