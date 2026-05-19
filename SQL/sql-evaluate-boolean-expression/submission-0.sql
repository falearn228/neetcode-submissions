-- Write your query below
select e.left_operand, e.operator, e.right_operand, 
case when e.operator = '=' and l.value = r.value then 'true'
 when e.operator = '>' and l.value > r.value then 'true'
 when e.operator = '<' and l.value < r.value then 'true'
else 'false' end as value
from expressions e
join variables l on l.name = e.left_operand
join variables r on r.name = e.right_operand
