-- +migrate Up
-- +migrate StatementBegin

create table likes (
    id uuid default uuid_generate_v4(),
    user_id uuid not null,
    post_id uuid not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    primary key (id)
)

-- +migrate StatementEnd