# Write your MySQL query statement below
SELECT a.Id FROM Weather AS a, Weather AS b
WHERE DATEDIFF(a.RecordDate, b.RecordDate)=1 AND a.Temperature > b.Temperature
