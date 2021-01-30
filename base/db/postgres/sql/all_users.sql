create table all_users (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
login_id text,
full_name text,
state integer DEFAULT 1,
version text default 'v1',
uts timestamp default current_timestamp
);