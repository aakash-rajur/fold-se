select
up.project_id as project_id,
up.user_id as user_id,
u.name as user_name,
u.created_at as user_created_at
from user_projects up
inner join users u on up.user_id = u.id
where true
and up.project_id = :project_id; -- :project_id type: bigint