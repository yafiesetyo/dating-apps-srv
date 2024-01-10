-- migrate:up
create table "users" (
    id bigserial primary key,
    "name" text not null,
    "gender" text not null,
    "username" text not null,
    "password" text not null,
    dob timestamp not null,
    pob text not null,
    "description" text,
    hobby text,
    image_url text[],
    religion text not null,
    verified_user boolean default false,
    unlimited_swipe boolean default false,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz
);

create table "user_swipes" (
    id bigserial primary key,
    user_id bigint not null,
    liked_user_id bigint not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz,

    constraint fk_user_id foreign key (user_id) references "users" (id),
    constraint fk_liked_user_id foreign key (liked_user_id) references "users" (id)
);


-- migrate:down
drop table "user_swipes";
drop table "users";

