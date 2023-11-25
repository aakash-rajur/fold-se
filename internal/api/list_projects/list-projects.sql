select
count(*) over () as record_count,
p.id as id,
p.name as name,
p.slug as slug,
p.description as description,
p.created_at as created_at,
(
  select
  jsonb_agg(
    jsonb_build_object(
      'id', u.id,
      'name', u.name
    ) order by u.id asc
  )
  from user_projects up
  inner join users u on up.user_id = u.id
  where true
  and up.project_id = p.id
) as users,
(
  select
  jsonb_agg(
    jsonb_build_object(
      'id', h.id,
      'name', h.name
    ) order by h.id asc
  )
  from project_hashtags ph
  inner join hashtags h on ph.hashtag_id = h.id
  where true
  and ph.project_id = p.id
) as hashtags
from projects p
where true
order by p.id
offset :offset -- :offset type: bigint
limit :limit; -- :limit type: bigint