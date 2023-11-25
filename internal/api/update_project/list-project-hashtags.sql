select
ph.project_id as project_id,
ph.hashtag_id as hashtag_id,
h.name as hashtag_name,
h.created_at as hashtag_created_at
from project_hashtags ph
inner join hashtags h on ph.hashtag_id = h.id
where true
and ph.project_id = :project_id; -- :project_id type: bigint
