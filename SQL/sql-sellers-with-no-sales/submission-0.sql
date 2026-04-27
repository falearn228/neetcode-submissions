-- Write your query below
select seller_name
from seller as s
where s.seller_id not in (
    select seller_id
    from orders
    where sale_date >= '2020-01-01' AND sale_date <= '2020-12-31'
)
order by seller_name asc