-- migrate:up
alter table user_swipes add column passed_user_id bigint;
alter table user_swipes alter column liked_user_id drop not null;
alter table user_swipes add constraint fk_passed_user_id foreign key (passed_user_id) references users(id);


-- migrate:down

