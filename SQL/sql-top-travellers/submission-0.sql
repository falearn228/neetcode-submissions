-- Write your query below
select u.name, coalesce(sum(r.distance), 0) as travelled_distance
from users u
left join rides as r on r.user_id = u.id
group by u.name
order by travelled_distance desc ,u.name asc