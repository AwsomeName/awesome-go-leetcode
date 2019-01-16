# Write your MySQL query statement below
# SELECT A.Name from Customers A
# WHERE NOT EXISTS (SELECT 1 FROM Orders B WHERE A.Id = B.CustomerId)

# SELECT A.Name from Customers A
# LEFT JOIN Orders B on  a.Id = B.CustomerId
# WHERE b.CustomerId is NULL

SELECT A.Name as Customers from Customers A
WHERE A.Id NOT IN (SELECT B.CustomerId from Orders B)
