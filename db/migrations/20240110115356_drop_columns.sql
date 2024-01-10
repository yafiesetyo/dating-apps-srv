-- migrate:up
alter table user_swipes drop column passed_user_id;

-- migrate:down

