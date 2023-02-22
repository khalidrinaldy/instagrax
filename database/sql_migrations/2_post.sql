-- +migrate Up
-- +migrate StatementBegin

create table post (
    id uuid default uuid_generate_v4(),
    image_url varchar(256) not null,
    caption text not null,
    user_id uuid not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    primary key (id)
)

-- +migrate StatementEnd