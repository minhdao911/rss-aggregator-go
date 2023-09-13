-- +goose Up
create table posts (
    id uuid primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    description text,
    published_at timestamp not null,
    url text not null unique,
    feed_id uuid not null references feeds(id) on delete cascade
);

-- +goose Down
drop table posts;
