-- +migrate Up
-- +migrate StatementBegin

create table users (
    id uuid default uuid_generate_v4(),
    username varchar(25) not null unique,
    name varchar(256) not null,
    email varchar(256) not null unique,
    password varchar(72) not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    primary key (id)
)

-- +migrate StatementEnd